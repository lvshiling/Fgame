package trade

import (
	"context"
	centerclient "fgame/fgame/center/client"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	corerunner "fgame/fgame/core/runner"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fgame/fgame/trade_server/store"
	tradeservertypes "fgame/fgame/trade_server/types"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	lru "github.com/hashicorp/golang-lru"
)

var cacheSize = 500

type TradeServer struct {
	rwm               sync.RWMutex
	db                coredb.DBService
	rs                coreredis.RedisService
	tradeStore        store.TradeStore
	tradeObjMap       map[int64]*TradeObject
	serverTradeObjMap map[int32]map[int32]map[int64]*TradeObject
	//缓存
	tradeObjCache      *lru.TwoQueueCache
	localTradeObjCache *lru.TwoQueueCache
	runner             corerunner.GoRunner
	centerClient       *centerclient.Client
	tradeServerMap     map[int32]map[int32]int32
}

const (
	refreshTime = time.Second * 10
)

func (s *TradeServer) init() (err error) {

	s.runner = corerunner.NewGoRunner("trade", s.Heartbeat, refreshTime)

	s.tradeStore = store.NewTradeStore(s.db)

	s.tradeObjCache, err = lru.New2Q(cacheSize)
	if err != nil {
		return
	}
	s.localTradeObjCache, err = lru.New2Q(cacheSize)
	if err != nil {
		return
	}

	s.tradeObjMap = make(map[int64]*TradeObject)
	s.serverTradeObjMap = make(map[int32]map[int32]map[int64]*TradeObject)
	tradeItemEntityList, err := s.tradeStore.GetAll(tradeservertypes.TradeStatusInit)
	if err != nil {
		return
	}
	for _, tradeItemEntity := range tradeItemEntityList {
		tradeObj := NewTradeObject()
		err = tradeObj.FromEntity(tradeItemEntity)
		if err != nil {
			return
		}
		s.addTradeItem(tradeObj)
	}
	return
}

func (s *TradeServer) Start() {
	s.runner.Start()
}

func (s *TradeServer) Stop() {
	s.runner.Stop()
}

func (s *TradeServer) addTradeItem(tradeObj *TradeObject) {
	_, ok := s.tradeObjMap[tradeObj.id]
	if !ok {
		s.tradeObjMap[tradeObj.id] = tradeObj
	}
	//平台和服务器索引
	platformTradeObjMap, ok := s.serverTradeObjMap[tradeObj.platform]
	if !ok {
		platformTradeObjMap = make(map[int32]map[int64]*TradeObject)
		s.serverTradeObjMap[tradeObj.platform] = platformTradeObjMap
	}
	serverTradeObjMap, ok := platformTradeObjMap[tradeObj.serverId]
	if !ok {
		serverTradeObjMap = make(map[int64]*TradeObject)
		platformTradeObjMap[tradeObj.serverId] = serverTradeObjMap
	}
	_, ok = serverTradeObjMap[tradeObj.id]
	if !ok {
		serverTradeObjMap[tradeObj.id] = tradeObj
	}
	s.tradeObjCache.Add(tradeObj.GetId(), tradeObj)
	s.localTradeObjCache.Add(tradeObj.GetTradeId(), tradeObj)
}

func (s *TradeServer) getTradeItem(tradeId int64) (tradeObj *TradeObject, err error) {
	tradeObjIface, ok := s.tradeObjCache.Get(tradeId)
	if !ok {
		goto DB
	}
	tradeObj, ok = tradeObjIface.(*TradeObject)
	if ok {
		return
	}
DB:
	tradeEntity, err := s.tradeStore.GetTradeItemByTradeId(tradeId)
	if err != nil {
		return
	}
	if tradeEntity == nil {
		return
	}
	tradeObj = NewTradeObject()
	err = tradeObj.FromEntity(tradeEntity)
	if err != nil {
		return
	}
	s.tradeObjCache.Add(tradeId, tradeObj)
	return
}

func (s *TradeServer) getLocalTradeItem(tradeId int64) (tradeObj *TradeObject, err error) {
	tradeObjIface, ok := s.localTradeObjCache.Get(tradeId)
	if !ok {
		goto DB
	}
	tradeObj, ok = tradeObjIface.(*TradeObject)
	if ok {
		return
	}
DB:
	tradeEntity, err := s.tradeStore.GetTradeItemByLocalTradeId(tradeId)
	if err != nil {
		return
	}
	if tradeEntity == nil {
		return
	}
	tradeObj = NewTradeObject()
	err = tradeObj.FromEntity(tradeEntity)
	if err != nil {
		return
	}
	s.localTradeObjCache.Add(tradeId, tradeObj)
	return
}

func (s *TradeServer) removeTradeItem(tradeObj *TradeObject) {
	delete(s.tradeObjMap, tradeObj.id)
	platformTradeObjMap, ok := s.serverTradeObjMap[tradeObj.platform]
	if ok {
		serverTradeObjMap, ok := platformTradeObjMap[tradeObj.serverId]
		if ok {
			delete(serverTradeObjMap, tradeObj.id)
		}
	}
	s.tradeObjCache.Remove(tradeObj.GetId())
	s.localTradeObjCache.Remove(tradeObj.GetTradeId())
}

//交易商品
func (s *TradeServer) Heartbeat() {
	ctx := context.Background()
	resp, err := s.centerClient.GetTradeServerList(ctx)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("trade:同步交易服务器列表,失败")
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()
	log.WithFields(
		log.Fields{}).Info("trade:同步交易服务器列表")
	s.tradeServerMap = make(map[int32]map[int32]int32)
	for _, serverInfo := range resp.TradeServerInfoList {
		platformTradeServerMap, ok := s.tradeServerMap[serverInfo.Platform]
		if !ok {
			platformTradeServerMap = make(map[int32]int32)
			s.tradeServerMap[serverInfo.Platform] = platformTradeServerMap
		}
		platformTradeServerMap[serverInfo.ServerId] = serverInfo.RegionId
	}
}

//获取区id
func (s *TradeServer) getTradeRegionId(platform int32, serverId int32) (regionId int32, ok bool) {
	tradeServerMap, ok := s.tradeServerMap[platform]
	if !ok {
		return 0, false
	}
	regionId, ok = tradeServerMap[serverId]
	if !ok {
		return 0, false
	}
	return regionId, true
}

const (
	maxShowNum = 1000
)

//获取商品列表
func (s *TradeServer) GetTradeList(platform int32, serverId int32) (tradeObjList []*TradeObject) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	regionId, ok := s.getTradeRegionId(platform, serverId)
	if !ok {
		return nil
	}
	platformTradeObjMap, ok := s.serverTradeObjMap[platform]
	if !ok {
		return nil
	}
	for tempServerId, serverTradeObjMap := range platformTradeObjMap {
		tempRegionId, ok := s.getTradeRegionId(platform, tempServerId)
		if !ok {
			continue
		}
		if regionId != tempRegionId {
			continue
		}
		for _, tradeObj := range serverTradeObjMap {
			tradeObjList = append(tradeObjList, tradeObj)
		}
	}
	if len(tradeObjList) >= maxShowNum {
		tradeObjList = tradeObjList[:maxShowNum]
	}
	return
}

//获取商品列表数量
func (s *TradeServer) getTradeListNum(platform int32, serverId int32) int32 {
	regionId, ok := s.getTradeRegionId(platform, serverId)
	if !ok {
		return 0
	}
	platformTradeObjMap, ok := s.serverTradeObjMap[platform]
	if !ok {
		return 0
	}
	numOfItems := int32(0)
	for tempServerId, serverTradeObjMap := range platformTradeObjMap {
		tempRegionId, ok := s.getTradeRegionId(platform, tempServerId)
		if !ok {
			continue
		}
		if regionId != tempRegionId {
			continue
		}
		numOfItems += int32(len(serverTradeObjMap))
	}
	return numOfItems
}

const (
	maxTradeNum = int32(1000)
)

//下单物品
func (s *TradeServer) Upload(platform int32, serverId int32, tradeId int64, playerId int64, playerName string, itemId int32, itemNum int32, porpertyData string, gold int32, level int32) (tradeObj *TradeObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	tradeObj, err = s.getLocalTradeItem(tradeId)
	if err != nil {
		return
	}
	if tradeObj != nil {
		log.WithFields(
			log.Fields{
				"platform":     platform,
				"serverId":     serverId,
				"tradeId":      tradeId,
				"playerId":     playerId,
				"playerName":   playerName,
				"itemId":       itemId,
				"itemNum":      itemNum,
				"gold":         gold,
				"propertyData": porpertyData,
				"err":          err,
			}).Info("trade:已经上传过了")
		return
	}
	currentNum := s.getTradeListNum(platform, serverId)
	if currentNum >= maxTradeNum {
		log.WithFields(
			log.Fields{
				"platform":   platform,
				"serverId":   serverId,
				"currentNum": currentNum,
				"maxNum":     maxTradeNum,
			}).Warn("trade:超过最大数量")
		err = statusTradeUploadMax.Err()
		return
	}
	now := timeutils.TimeToMillisecond(time.Now())
	tradeObj = NewTradeObject()
	tradeObj.id, _ = idutil.GetId()
	tradeObj.platform = platform
	tradeObj.serverId = serverId
	tradeObj.tradeId = tradeId
	tradeObj.playerId = playerId
	tradeObj.playerName = playerName
	tradeObj.itemId = itemId
	tradeObj.itemNum = itemNum
	tradeObj.propertyData = porpertyData
	tradeObj.level = level
	tradeObj.gold = gold
	tradeObj.status = tradeservertypes.TradeStatusInit
	tradeObj.buyPlayerId = 0
	tradeObj.buyPlayerName = ""
	tradeObj.updateTime = 0
	tradeObj.createTime = now
	e, err := tradeObj.ToEntity()
	if err != nil {
		return
	}
	//保存数据
	if err = s.db.DB().Save(e).Error; err != nil {
		return
	}
	s.addTradeItem(tradeObj)

	return
}

//下架商品
func (s *TradeServer) Withdraw(platform int32, serverId int32, tradeId int64) (tradeObj *TradeObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	tradeObj, err = s.getTradeItem(tradeId)
	if err != nil {
		return
	}
	if tradeObj == nil {
		log.WithFields(
			log.Fields{
				"platform": platform,
				"serverId": serverId,
				"tradeId":  tradeId,
			}).Warn("trade:下架,已经不存在")
		err = statusTradeItemAreadyNoExistOrSell.Err()
		return
	}

	if tradeObj.GetPlatform() != platform {
		log.WithFields(
			log.Fields{
				"platform":       platform,
				"serverId":       serverId,
				"tradeId":        tradeId,
				"actualPlatform": tradeObj.GetPlatform(),
				"actualServerId": tradeObj.GetServerId(),
			}).Warn("trade:下架商品,平台或服务器对不上")
		err = statusTradeItemAreadyNoExistOrSell.Err()
		return
	}

	switch tradeObj.status {
	case tradeservertypes.TradeStatusInit:
		{
			tradeObj.status = tradeservertypes.TradeStatusWithdraw
			e, err := tradeObj.ToEntity()
			if err != nil {
				return nil, err
			}
			//保存数据
			if err = s.db.DB().Save(e).Error; err != nil {
				return nil, err
			}
			s.removeTradeItem(tradeObj)
			return tradeObj, nil
		}
	case tradeservertypes.TradeStatusWithdraw:
		{
			return
		}
	default:
		{
			//已经出售
			log.WithFields(
				log.Fields{
					"platform": platform,
					"serverId": serverId,
					"tradeId":  tradeId,
					"status":   tradeObj.status,
				}).Warn("trade:下架商品,物品已经出售,或者状态不对")
			err = statusTradeItemAreadyNoExistOrSell.Err()
			return
		}
	}

	return
}

//交易商品
func (s *TradeServer) TradeItem(globalTradeId int64, buyPlatform int32, buyServerId int32, buyPlayerId int64, buyPlayerName string) (tradeObj *TradeObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	tradeObj, err = s.getTradeItem(globalTradeId)
	if err != nil {
		return
	}
	if tradeObj == nil {
		log.WithFields(
			log.Fields{
				"tradeId":       globalTradeId,
				"buyPlatform":   buyPlatform,
				"buyServerId":   buyServerId,
				"buyPlayerId":   buyPlayerId,
				"buyPlayerName": buyPlayerName,
			}).Warn("trade:物品交易,不存在")
		err = statusTradeItemAreadyNoExistOrSell.Err()
		return
	}
	switch tradeObj.status {
	case tradeservertypes.TradeStatusInit:
		{
			now := timeutils.TimeToMillisecond(time.Now())
			tradeObj.status = tradeservertypes.TradeStatusSell
			tradeObj.updateTime = now
			tradeObj.buyPlayerPlatform = buyPlatform
			tradeObj.buyPlayerServerId = buyServerId
			tradeObj.buyPlayerName = buyPlayerName
			tradeObj.buyPlayerId = buyPlayerId

			e, err := tradeObj.ToEntity()
			if err != nil {
				return nil, err
			}
			//保存数据
			if err = s.db.DB().Save(e).Error; err != nil {
				return nil, err
			}
			s.removeTradeItem(tradeObj)
			return tradeObj, nil

		}
	case tradeservertypes.TradeStatusSell,
		tradeservertypes.TradeStatusSellNotice:
		{
			if buyPlayerId == tradeObj.GetBuyPlayerId() {
				return
			}
			//已经出售
			log.WithFields(
				log.Fields{
					"tradeId":                 globalTradeId,
					"actualBuyPlayerId":       tradeObj.GetBuyPlayerId(),
					"actualBuyServerId":       tradeObj.GetBuyPlayerServerId(),
					"actualBuyPlayerPlatform": tradeObj.GetBuyPlayerPlatform(),
					"status":                  tradeObj.status,
					"buyPlatform":             buyPlatform,
					"buyServerId":             buyServerId,
					"buyPlayerId":             buyPlayerId,
					"buyPlayerName":           buyPlayerName,
				}).Warn("trade:物品已经出售")
			err = statusTradeItemAreadyNoExistOrSell.Err()
			return
		}
	default:
		{
			//已经出售
			log.WithFields(
				log.Fields{
					"tradeId":       globalTradeId,
					"buyPlatform":   buyPlatform,
					"buyServerId":   buyServerId,
					"buyPlayerId":   buyPlayerId,
					"buyPlayerName": buyPlayerName,
					"status":        tradeObj.status,
				}).Warn("trade:物品已经下架,或者状态不对")
			err = statusTradeItemAreadyNoExistOrSell.Err()
			return
		}
	}

	return
}

func NewTradeServer(db coredb.DBService, rs coreredis.RedisService, centerClient *centerclient.Client) (ss *TradeServer, err error) {
	ss = &TradeServer{}
	ss.db = db
	ss.rs = rs
	ss.centerClient = centerClient
	err = ss.init()
	if err != nil {
		return
	}
	return
}
