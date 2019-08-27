package remote

import (
	"context"
	logintypes "fgame/fgame/account/login/types"
	centerclient "fgame/fgame/center/client"
	"fgame/fgame/charge_server/charge"
	"fgame/fgame/charge_server/types"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	corerunner "fgame/fgame/core/runner"
	remoteclient "fgame/fgame/game/remote/client"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func SetupRemoteServiceHandler(s RemoteService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithRemoteService(ctx, s)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

type contextKey string

const (
	remoteServiceKey contextKey = "fgame.remote"
)

func WithRemoteService(parent context.Context, s RemoteService) context.Context {
	ctx := context.WithValue(parent, remoteServiceKey, s)
	return ctx
}

func RemoteServiceInContext(parent context.Context) RemoteService {
	s := parent.Value(remoteServiceKey)
	if s == nil {
		return nil
	}
	ts, ok := s.(RemoteService)
	if !ok {
		return nil
	}
	return ts
}

type RemoteService interface {
	Heartbeat()
	Charge(obj *charge.OrderObject) bool
	Start()
	Stop()
}

type RemoteConfig struct {
	Center *centerclient.Config `json:"center"`
}

type remoteService struct {
	rwm sync.RWMutex
	cfg *RemoteConfig
	// client         remoteclient.RemoteClient
	orderObjectMap map[string]*charge.OrderObject
	runner         corerunner.GoRunner
	db             coredb.DBService
	rs             coreredis.RedisService
	centerClient   *centerclient.Client
}

func (s *remoteService) init() error {
	centerClient, err := centerclient.NewClient(s.cfg.Center)
	if err != nil {
		return err
	}
	s.centerClient = centerClient
	s.orderObjectMap = make(map[string]*charge.OrderObject)
	s.runner = corerunner.NewGoRunner("remote", s.Heartbeat, refreshTime)

	err = s.loadUnfinishOrder()
	if err != nil {
		return err
	}
	return nil
}

var (
	maxDelay = 5 * time.Second
)

func (s *remoteService) getRemoteClient(ctx context.Context, sdkType logintypes.SDKType, serverId int32) (c remoteclient.RemoteClient, err error) {
	resp, err := s.centerClient.GetServerInfo(ctx, int32(sdkType), serverId)
	if err != nil {
		return
	}
	if resp.ServerInfo == nil {
		return
	}
	ip := resp.ServerInfo.RemoteIp
	var options []grpc.DialOption
	options = append(options, grpc.WithInsecure())

	callOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(waitBetween)),
		grpc_retry.WithCodes(codes.Unavailable, codes.Internal, codes.Aborted),
		grpc_retry.WithMax(maxRetry),
	}
	options = append(options, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(callOpts...)))

	maxDelayOption := grpc.WithBackoffMaxDelay(5 * time.Second)
	options = append(options, maxDelayOption)

	conn, err := grpc.Dial(ip, options...)
	if err != nil {
		return nil, err
	}
	c = remoteclient.NewRemoteClient(conn)

	return
}

func (s *remoteService) loadUnfinishOrder() (err error) {
	orderEntityList := make([]*charge.OrderEntity, 0, 4)
	if err = s.db.DB().Find(&orderEntityList, "status=?", int32(types.OrderStatusPay)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	for _, orderEntity := range orderEntityList {
		orderObj := charge.NewOrderObject()
		orderObj.FromEntity(orderEntity)
		s.orderObjectMap[orderObj.GetOrderId()] = orderObj
	}
	return
}

func (s *remoteService) Start() {
	s.runner.Start()
}

func (s *remoteService) Stop() {
	s.runner.Stop()
}

func (s *remoteService) Heartbeat() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	for _, orderObject := range s.orderObjectMap {
		s.retry(orderObject)
	}
}

func (s *remoteService) retry(obj *charge.OrderObject) {
	go func() {
		playerId := obj.GetPlayerId()
		orderId := obj.GetOrderId()
		chargeId := obj.GetChargeId()
		skdType := obj.GetSdkType()
		success, err := s.charge(obj)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"orderId":  orderId,
					"chargeId": chargeId,
					"sdkType":  skdType.String(),
					"error":    err,
				}).Error("remote:充值回调,重试失败")
			return
		}
		if !success {
			return
		}
		s.finishOrderObject(obj)
	}()
}

func (s *remoteService) addFailedOrderObject(orderObject *charge.OrderObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.orderObjectMap[orderObject.GetOrderId()] = orderObject
}

func (s *remoteService) finishOrderObject(obj *charge.OrderObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	delete(s.orderObjectMap, obj.GetOrderId())

	now := timeutils.TimeToMillisecond(time.Now())
	obj.SetStatus(types.OrderStatusFinish)
	obj.SetUpdateTime(now)
	e := obj.ToEntity()
	playerId := obj.GetPlayerId()
	orderId := obj.GetOrderId()
	chargeId := obj.GetChargeId()
	skdType := obj.GetSdkType()
	if err := s.db.DB().Save(e).Error; err != nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"orderId":  orderId,
				"chargeId": chargeId,
				"sdkType":  skdType.String(),
				"error":    err,
			}).Error("remote:完成充值,保存失败")
		return
	}

}

func (s *remoteService) Charge(obj *charge.OrderObject) bool {
	if obj.GetStatus() != types.OrderStatusPay {
		return false
	}
	playerId := obj.GetPlayerId()
	orderId := obj.GetOrderId()
	chargeId := obj.GetChargeId()
	skdType := obj.GetSdkType()
	success, err := s.charge(obj)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"orderId":  orderId,
				"chargeId": chargeId,
				"sdkType":  skdType.String(),
				"error":    err,
			}).Error("remote:充值回调,错误")
		s.addFailedOrderObject(obj)
		return true
	}
	if !success {
		s.addFailedOrderObject(obj)
		return true
	}
	now := timeutils.TimeToMillisecond(time.Now())
	obj.SetUpdateTime(now)
	obj.SetStatus(types.OrderStatusFinish)
	log.WithFields(
		log.Fields{
			"playerId": playerId,
			"orderId":  orderId,
			"chargeId": chargeId,
			"sdkType":  skdType.String(),
		}).Info("remote:充值回调,成功")
	e := obj.ToEntity()
	if err := s.db.DB().Save(e).Error; err != nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"orderId":  orderId,
				"chargeId": chargeId,
				"sdkType":  skdType.String(),
				"error":    err,
			}).Error("remote:充值回调,保存失败")
	}
	return true
}

const (
	chargeTimeout = time.Second * 5
)

func (s *remoteService) charge(obj *charge.OrderObject) (success bool, err error) {
	client, err := s.getRemoteClient(context.Background(), obj.GetSdkType(), obj.GetServerId())
	if err != nil {
		return
	}
	if client == nil {
		return false, fmt.Errorf("remote:找不到远程链接")
	}
	cmdCharge := &cmdpb.CmdCharge{}
	playerId := obj.GetPlayerId()
	cmdCharge.PlayerId = &playerId
	orderId := obj.GetOrderId()
	cmdCharge.OrderId = &orderId
	chargeId := obj.GetChargeId()
	cmdCharge.ChargeId = &chargeId
	skdType := obj.GetSdkType()
	money := obj.GetMoney()
	cmdCharge.Money = &money
	log.WithFields(
		log.Fields{
			"playerId": playerId,
			"orderId":  orderId,
			"chargeId": chargeId,
			"sdkType":  skdType.String(),
		}).Info("remote:充值回调中")
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, chargeTimeout)
	defer cancel()
	resp, err := client.DoCmd(timeoutCtx, cmdCharge)
	if err != nil {
		return false, err
	}
	defer client.Close()
	if resp.GetErrorCode() != 0 {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"orderId":  orderId,
				"chargeId": chargeId,
				"sdkType":  skdType.String(),
				"errorMsg": resp.GetErrorMsg(),
			}).Warn("remote:充值回调,失败")
		return
	}
	return true, nil
}

const (
	waitBetween = time.Second
	maxRetry    = 3
	refreshTime = time.Second * 10
)

func NewRemoteService(cfg *RemoteConfig, db coredb.DBService, rs coreredis.RedisService) (ts RemoteService, err error) {
	s := &remoteService{}
	s.cfg = cfg
	s.db = db
	s.rs = rs

	err = s.init()
	if err != nil {
		return
	}
	ts = s
	return
}
