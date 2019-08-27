package remote

import (
	"context"
	centerclient "fgame/fgame/center/client"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	corerunner "fgame/fgame/core/runner"
	remoteclient "fgame/fgame/game/remote/client"
	cmdpb "fgame/fgame/game/remote/cmd/pb"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/trade_server/store"
	"fgame/fgame/trade_server/trade"
	"fgame/fgame/trade_server/types"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	log "github.com/Sirupsen/logrus"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type RemoteService interface {
	Heartbeat()
	Sell(obj *trade.TradeObject) bool
	Start()
	Stop()
}

type remoteService struct {
	rwm          sync.RWMutex
	tradeItemMap map[int64]*trade.TradeObject
	runner       corerunner.GoRunner
	db           coredb.DBService
	rs           coreredis.RedisService
	centerClient *centerclient.Client
}

func (s *remoteService) init() error {

	s.tradeItemMap = make(map[int64]*trade.TradeObject)
	s.runner = corerunner.NewGoRunner("remote", s.Heartbeat, refreshTime)

	err := s.loadUnfinishTradeItem()
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

func (s *remoteService) loadUnfinishTradeItem() (err error) {
	tradeItemEntityList := make([]*store.TradeItemEntity, 0, 4)
	if err = s.db.DB().Find(&tradeItemEntityList, "status=?", int32(types.TradeStatusSell)).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	for _, tradeItemEntity := range tradeItemEntityList {
		tradeObj := trade.NewTradeObject()
		tradeObj.FromEntity(tradeItemEntity)
		s.tradeItemMap[tradeObj.GetId()] = tradeObj
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
	for _, tradeItem := range s.tradeItemMap {
		s.retry(tradeItem)
	}
}

func (s *remoteService) retry(obj *trade.TradeObject) {
	go func() {
		playerId := obj.GetPlayerId()
		tradeId := obj.GetId()
		platform := obj.GetPlatform()
		serverId := obj.GetServerId()
		playerName := obj.GetPlayerName()
		itemId := obj.GetItemId()
		itemNum := obj.GetItemNum()
		gold := obj.GetGold()
		propertyData := obj.GetPropertyData()
		status := obj.GetStatus()
		buyPlatform := obj.GetBuyPlayerPlatform()
		buyServerId := obj.GetBuyPlayerServerId()
		buyPlayerId := obj.GetBuyPlayerId()
		buyPlayerName := obj.GetBuyPlayerName()
		success, err := s.sell(obj)
		if err != nil {
			log.WithFields(
				log.Fields{
					"tradeId":       tradeId,
					"playerId":      playerId,
					"platform":      platform,
					"serverId":      serverId,
					"playerName":    playerName,
					"itemId":        itemId,
					"itemNum":       itemNum,
					"gold":          gold,
					"propertyData":  propertyData,
					"status":        status,
					"buyPlatform":   buyPlatform,
					"buyServerId":   buyServerId,
					"buyPlayerId":   buyPlayerId,
					"buyPlayerName": buyPlayerName,
					"error":         err,
				}).Error("remote:交易回调,重试失败")
			return
		}
		if !success {
			return
		}
		s.finishTradeObject(obj)
	}()
}

func (s *remoteService) addFailedTradeObject(tradeObject *trade.TradeObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.tradeItemMap[tradeObject.GetId()] = tradeObject
}

func (s *remoteService) finishTradeObject(obj *trade.TradeObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := obj.GetPlayerId()
	tradeId := obj.GetId()
	platform := obj.GetPlatform()
	serverId := obj.GetServerId()
	playerName := obj.GetPlayerName()
	itemId := obj.GetItemId()
	itemNum := obj.GetItemNum()
	gold := obj.GetGold()
	propertyData := obj.GetPropertyData()
	status := obj.GetStatus()
	buyPlatform := obj.GetBuyPlayerPlatform()
	buyServerId := obj.GetBuyPlayerServerId()
	buyPlayerId := obj.GetBuyPlayerId()
	buyPlayerName := obj.GetBuyPlayerName()
	now := timeutils.TimeToMillisecond(time.Now())
	flag := obj.SellNotice(now)
	if !flag {
		log.WithFields(
			log.Fields{
				"tradeId":       tradeId,
				"playerId":      playerId,
				"platform":      platform,
				"serverId":      serverId,
				"playerName":    playerName,
				"itemId":        itemId,
				"itemNum":       itemNum,
				"gold":          gold,
				"propertyData":  propertyData,
				"status":        status,
				"buyPlatform":   buyPlatform,
				"buyServerId":   buyServerId,
				"buyPlayerId":   buyPlayerId,
				"buyPlayerName": buyPlayerName,
			}).Warn("remote:完成交易,失败")
		return
	}

	delete(s.tradeItemMap, obj.GetId())
	e, _ := obj.ToEntity()
	if err := s.db.DB().Save(e).Error; err != nil {
		log.WithFields(
			log.Fields{
				"tradeId":       tradeId,
				"playerId":      playerId,
				"platform":      platform,
				"serverId":      serverId,
				"playerName":    playerName,
				"itemId":        itemId,
				"itemNum":       itemNum,
				"gold":          gold,
				"propertyData":  propertyData,
				"status":        status,
				"buyPlatform":   buyPlatform,
				"buyServerId":   buyServerId,
				"buyPlayerId":   buyPlayerId,
				"buyPlayerName": buyPlayerName,
				"error":         err,
			}).Error("remote:完成交易,保存失败")
		return
	}

}

func (s *remoteService) Sell(obj *trade.TradeObject) bool {
	if obj.GetStatus() != types.TradeStatusSell {
		return false
	}
	playerId := obj.GetPlayerId()
	tradeId := obj.GetId()
	platform := obj.GetPlatform()
	serverId := obj.GetServerId()
	playerName := obj.GetPlayerName()
	itemId := obj.GetItemId()
	itemNum := obj.GetItemNum()
	gold := obj.GetGold()
	propertyData := obj.GetPropertyData()
	status := obj.GetStatus()
	buyPlatform := obj.GetBuyPlayerPlatform()
	buyServerId := obj.GetBuyPlayerServerId()
	buyPlayerId := obj.GetBuyPlayerId()
	buyPlayerName := obj.GetBuyPlayerName()
	success, err := s.sell(obj)
	if err != nil {
		log.WithFields(
			log.Fields{
				"tradeId":       tradeId,
				"playerId":      playerId,
				"platform":      platform,
				"serverId":      serverId,
				"playerName":    playerName,
				"itemId":        itemId,
				"itemNum":       itemNum,
				"gold":          gold,
				"propertyData":  propertyData,
				"status":        status,
				"buyPlatform":   buyPlatform,
				"buyServerId":   buyServerId,
				"buyPlayerId":   buyPlayerId,
				"buyPlayerName": buyPlayerName,
				"error":         err,
			}).Error("remote:交易回调,错误")
		s.addFailedTradeObject(obj)
		return true
	}
	if !success {
		s.addFailedTradeObject(obj)
		return true
	}
	s.finishTradeObject(obj)
	return true
}

const (
	chargeTimeout = time.Second * 5
)

func (s *remoteService) sell(obj *trade.TradeObject) (success bool, err error) {
	client, err := s.getRemoteClient(context.Background(), obj.GetPlatform(), obj.GetServerId())
	if err != nil {
		return
	}
	if client == nil {
		return false, fmt.Errorf("remote:找不到远程链接")
	}
	cmdTradeSell := &cmdpb.CmdTradeSell{}
	playerId := obj.GetPlayerId()
	cmdTradeSell.PlayerId = &playerId
	serverId := obj.GetServerId()
	cmdTradeSell.ServerId = &serverId
	platform := obj.GetPlatform()
	cmdTradeSell.Platform = &platform
	tradeId := obj.GetTradeId()
	cmdTradeSell.TradeId = &tradeId
	gold := int32(obj.GetGold())
	cmdTradeSell.Gold = &gold
	globalTradeId := obj.GetId()
	cmdTradeSell.GlobalTradeId = &globalTradeId
	buyPlatform := obj.GetBuyPlayerPlatform()
	cmdTradeSell.BuyPlatform = &buyPlatform
	buyServerId := obj.GetBuyPlayerServerId()
	cmdTradeSell.BuyServerId = &buyServerId
	buyPlayerId := obj.GetBuyPlayerId()
	cmdTradeSell.BuyPlayerId = &buyPlayerId
	buyPlayerName := obj.GetBuyPlayerName()
	cmdTradeSell.BuyPlayerName = &buyPlayerName
	log.WithFields(
		log.Fields{
			"playerId":      playerId,
			"serverId":      serverId,
			"platform":      platform,
			"tradeId":       tradeId,
			"gold":          gold,
			"globalTradeId": globalTradeId,
			"buyPlatform":   buyPlatform,
			"buyServerId":   buyServerId,
			"buyPlayerId":   buyPlayerId,
			"buyPlayerName": buyPlayerName,
		}).Info("remote:交易回调中")
	ctx := context.Background()
	defer client.Close()
	timeoutCtx, cancel := context.WithTimeout(ctx, chargeTimeout)
	defer cancel()
	resp, err := client.DoCmd(timeoutCtx, cmdTradeSell)
	if err != nil {
		return false, err
	}

	if resp.GetErrorCode() != 0 {
		log.WithFields(
			log.Fields{
				"playerId":      playerId,
				"serverId":      serverId,
				"platform":      platform,
				"tradeId":       tradeId,
				"gold":          gold,
				"globalTradeId": globalTradeId,
				"buyPlatform":   buyPlatform,
				"buyServerId":   buyServerId,
				"buyPlayerId":   buyPlayerId,
				"buyPlayerName": buyPlayerName,
				"errorCode":     resp.GetErrorCode(),
			}).Warn("remote:交易回调,失败")
		return
	}
	return true, nil
}

const (
	waitBetween = time.Second
	maxRetry    = 3
	refreshTime = time.Second * 10
)

func NewRemoteService(db coredb.DBService, rs coreredis.RedisService, centerClient *centerclient.Client) (ts RemoteService, err error) {
	s := &remoteService{}

	s.db = db
	s.rs = rs
	s.centerClient = centerClient
	err = s.init()
	if err != nil {
		return
	}
	ts = s
	return
}
