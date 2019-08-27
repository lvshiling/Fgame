package remote

import (
	"context"
	centerclient "fgame/fgame/center/client"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	corerunner "fgame/fgame/core/runner"
	"fgame/fgame/coupon_server/exchange"
	"fgame/fgame/coupon_server/types"
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
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
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
	Exchange(obj *exchange.FeedbackfeeExchangeObject) bool
	Start()
	Stop()
}

type RemoteConfig struct {
	Center *centerclient.Config `json:"center"`
}

type remoteService struct {
	rwm            sync.RWMutex
	cfg            *RemoteConfig
	exchangeObjMap map[int64]*exchange.FeedbackfeeExchangeObject
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
	s.exchangeObjMap = make(map[int64]*exchange.FeedbackfeeExchangeObject)
	s.runner = corerunner.NewGoRunner("remote", s.Heartbeat, refreshTime)

	err = s.loadUnfinishExchange()
	if err != nil {
		return err
	}
	return nil
}

var (
	maxDelay = 5 * time.Second
)

func (s *remoteService) getRemoteClient(ctx context.Context, platform int32, serverId int32) (c remoteclient.RemoteClient, err error) {
	resp, err := s.centerClient.GetServerInfoByPlatform(ctx, platform, serverId)
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

func (s *remoteService) loadUnfinishExchange() (err error) {
	exchangeEntityList := make([]*exchange.FeedbackfeeExchangeEntity, 0, 4)
	if err = s.db.DB().Find(&exchangeEntityList, "status=?", int32(types.ExchangeStatusFinish)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	for _, exchangeEntity := range exchangeEntityList {
		exchangeObj := exchange.NewFeedbackfeeExchangeObject()
		exchangeObj.FromEntity(exchangeEntity)
		s.exchangeObjMap[exchangeObj.GetId()] = exchangeObj
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
	for _, exchangeObj := range s.exchangeObjMap {
		s.retry(exchangeObj)
	}
}

func (s *remoteService) retry(obj *exchange.FeedbackfeeExchangeObject) {
	go func(obj *exchange.FeedbackfeeExchangeObject) {
		playerId := obj.GetPlayerId()
		exchangeId := obj.GetExchangeId()
		money := obj.GetMoney()
		code := obj.GetCode()
		platform := obj.GetPlatform()
		success, err := s.exchange(obj)
		if err != nil {
			log.WithFields(
				log.Fields{
					"playerId":   playerId,
					"exchangeId": exchangeId,
					"money":      money,
					"platform":   platform,
					"code":       code,
					"error":      err,
				}).Error("remote:兑换回调,重试失败")
			return
		}
		if !success {
			return
		}
		s.finishExchangeObj(obj)
	}(obj)
}

func (s *remoteService) addFailedExchangeObject(exchangeObj *exchange.FeedbackfeeExchangeObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.exchangeObjMap[exchangeObj.GetId()] = exchangeObj
}

func (s *remoteService) finishExchangeObj(obj *exchange.FeedbackfeeExchangeObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	delete(s.exchangeObjMap, obj.GetId())

	now := timeutils.TimeToMillisecond(time.Now())
	flag := obj.Notify(now)
	if !flag {
		return
	}
	e, err := obj.ToEntity()
	if err != nil {
		return
	}
	playerId := obj.GetPlayerId()
	exchangeId := obj.GetExchangeId()
	money := obj.GetMoney()
	code := obj.GetCode()
	platform := obj.GetPlatform()
	if err := s.db.DB().Save(e).Error; err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"exchangeId": exchangeId,
				"money":      money,
				"platform":   platform,
				"code":       code,
				"error":      err,
			}).Error("remote:完成发货,保存失败")
		return
	}

}

func (s *remoteService) Exchange(obj *exchange.FeedbackfeeExchangeObject) bool {

	if obj.GetStatus() != types.ExchangeStatusFinish {
		return false
	}
	playerId := obj.GetPlayerId()
	exchangeId := obj.GetExchangeId()
	money := obj.GetMoney()
	code := obj.GetCode()
	platform := obj.GetPlatform()
	success, err := s.exchange(obj)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"exchangeId": exchangeId,
				"money":      money,
				"code":       code,
				"platform":   platform,
				"error":      err,
			}).Error("remote:兑换回调,错误")
		s.addFailedExchangeObject(obj)
		return true
	}
	if !success {
		s.addFailedExchangeObject(obj)
		return true
	}
	now := timeutils.TimeToMillisecond(time.Now())
	flag := obj.Notify(now)
	if !flag {
		return false
	}

	log.WithFields(
		log.Fields{
			"playerId":   playerId,
			"exchangeId": exchangeId,
			"money":      money,
			"code":       code,
			"platform":   platform,
		}).Info("remote:兑换回调,成功")

	s.finishExchangeObj(obj)
	return true
}

const (
	exchangeTimeout = time.Second * 5
)

func (s *remoteService) exchange(obj *exchange.FeedbackfeeExchangeObject) (success bool, err error) {
	client, err := s.getRemoteClient(context.Background(), obj.GetPlatform(), obj.GetServerId())
	if err != nil {
		return
	}
	if client == nil {
		return false, fmt.Errorf("remote:找不到远程链接")
	}
	cmdExchange := &cmdpb.CmdExchange{}
	playerId := obj.GetPlayerId()
	cmdExchange.PlayerId = &playerId
	exchangeId := obj.GetExchangeId()
	cmdExchange.ExchangeId = &exchangeId
	code := obj.GetCode()
	cmdExchange.Code = &code

	money := obj.GetMoney()
	cmdExchange.Money = &money
	log.WithFields(
		log.Fields{
			"playerId":   playerId,
			"exchangeId": exchangeId,
			"code":       code,
			"money":      money,
		}).Info("remote:兑换回调中")
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, exchangeTimeout)
	defer cancel()
	resp, err := client.DoCmd(timeoutCtx, cmdExchange)
	if err != nil {
		return false, err
	}
	defer client.Close()
	if resp.GetErrorCode() != 0 {
		log.WithFields(
			log.Fields{
				"playerId":   playerId,
				"exchangeId": exchangeId,
				"code":       code,
				"money":      money,
				"errorMsg":   resp.GetErrorMsg(),
			}).Warn("remote:兑换回调,失败")
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
