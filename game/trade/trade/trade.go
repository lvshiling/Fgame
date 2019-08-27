package trade

import (
	"context"
	"encoding/json"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/center/center"
	dummytemplate "fgame/fgame/game/dummy/template"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/trade/dao"
	tradeeventtypes "fgame/fgame/game/trade/event/types"
	playertrade "fgame/fgame/game/trade/player"
	tradetemplate "fgame/fgame/game/trade/template"
	tradetypes "fgame/fgame/game/trade/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	trade_serverclient "fgame/fgame/trade_server/client"
	"fmt"
	"runtime/debug"
	"sort"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	rpcTimeout = 5 * time.Second
)

type TradeService interface {
	Start()
	Stop()
	Heartbeat()
	GetGlobalTradeList() []*GlobalTradeItemObject
	GetNumOfGlobalTradeItems() int32
	GetGlobalTradeItem(globalTradeId int64) *GlobalTradeItemObject
	GetTradeList(pl player.Player) []*TradeItemObject
	GetCanRecycleTradeList(pl player.Player) []*TradeItemObject
	UploadItem(pl player.Player, itemId int32, itemNum int32, propertyData inventorytypes.ItemPropertyData, level int32, gold int32) (err error)
	WithdrawItem(pl player.Player, tradeId int64) (err error)
	TradeItem(pl player.Player, tradeId int64) (orderObj *TradeOrderObject, err error)
	SellItem(playerId int64, tradeId int64, buyPlayerPlatform int32, buyPlayerServerId int32, buyPlayerId int64, buyPlayerName string)
	EndSellItem(pl player.Player, tradeItemObj *TradeItemObject)
	EndTradeItem(pl player.Player, orderObj *TradeOrderObject)
	GetUnfinishOrderList(pl player.Player) []*TradeOrderObject
	GetSellList(pl player.Player) []*TradeItemObject
	SyncGlobalTradeList()
	SyncRetryTradeList()
	SyncRetryOrderList()
	SystemWithdrawTradeList()
	SystemRecycle()
	SystemRecycleTrade(tradeId int64) bool
	GMSetRecycle(recycle int64)
	GMSetCustomRecycle(recycle int64)
}

type TradeOptions struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type RecyleTimesObject struct {
	playerId int64
	num      int64
}

func newRecycleTimeObject(playerId int64) *RecyleTimesObject {
	o := &RecyleTimesObject{}
	o.playerId = playerId
	return o
}

type recycleTimesObjectList []*RecyleTimesObject

func (p recycleTimesObjectList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p recycleTimesObjectList) Len() int           { return len(p) }
func (p recycleTimesObjectList) Less(i, j int) bool { return p[i].num < p[j].num }

type tradeService struct {
	rwm         sync.RWMutex
	options     *TradeOptions
	tradeClient *trade_serverclient.Client
	//玩家map
	tradeItemListMap map[int64][]*TradeItemObject
	//全球交易列表
	globalTradeItemMap map[int64]*GlobalTradeItemObject

	//付款订单
	orderMap map[int64]*TradeOrderObject
	//已经完成但是未发货的
	unfinishOrderMap map[int64][]*TradeOrderObject
	//已经出售的商品
	sellItemMap map[int64][]*TradeItemObject
	//回购总额
	tradeRecyleObj *TradeRecycleObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner

	//回收数量
	systemRecycleMap map[int64]*RecyleTimesObject
}

//初始化
func (s *tradeService) init(options *TradeOptions) (err error) {
	s.options = options
	err = s.initTradeClient()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
			}).Error("trade:初始化客户端,错误")
		err = nil
	}
	s.tradeItemListMap = make(map[int64][]*TradeItemObject)
	s.globalTradeItemMap = make(map[int64]*GlobalTradeItemObject)

	s.orderMap = make(map[int64]*TradeOrderObject)

	s.unfinishOrderMap = make(map[int64][]*TradeOrderObject)
	s.sellItemMap = make(map[int64][]*TradeItemObject)
	s.systemRecycleMap = make(map[int64]*RecyleTimesObject)
	//初始化所有交易对象
	err = s.initTradeItems()
	if err != nil {
		return
	}
	//初始化出售过的
	err = s.initSellTradeItems()
	if err != nil {
		return
	}
	//初始化未完成的订单
	err = s.initOrder()
	if err != nil {
		return
	}
	//初始化未完成的订单
	err = s.initUnfinishTradeOrders()
	if err != nil {
		return
	}
	//初始化回购
	err = s.initRecycleObj()
	if err != nil {
		return
	}
	s.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	s.heartbeatRunner.AddTask(createGlobalTradeSyncTask(s))
	s.heartbeatRunner.AddTask(createTradeRetryTask(s))
	s.heartbeatRunner.AddTask(createSystemWithDrawTask(s))
	s.heartbeatRunner.AddTask(createTradeOrderRetryTask(s))
	s.heartbeatRunner.AddTask(createTradeRecycleTask(s))
	s.heartbeatRunner.AddTask(createTradeRecycleGoldTask(s))

	return
}

//初始化交易客户端
func (s *tradeService) initTradeClient() (err error) {
	cfg := &trade_serverclient.Config{
		Host: s.options.Host,
		Port: s.options.Port,
	}
	s.tradeClient, err = trade_serverclient.NewClient(cfg)
	if err != nil {
		return
	}
	return nil
}

//获取交易客户端
func (s *tradeService) getTradeClient() (c *trade_serverclient.Client, err error) {
	if s.tradeClient == nil {
		err = s.initTradeClient()
		if err != nil {
			return
		}
	}
	c = s.tradeClient
	return
}

//初始化上传物品
func (s *tradeService) initTradeItems() (err error) {
	serverId := global.GetGame().GetServerIndex()
	//上架中
	initList, err := dao.GetTradeDao().GetTradeItemList(serverId, tradetypes.TradeStatusInit)
	if err != nil {
		return
	}

	for _, entity := range initList {
		obj := createTradeItemObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		s.addTradeItem(obj)
	}
	//上架的
	uploadList, err := dao.GetTradeDao().GetTradeItemList(serverId, tradetypes.TradeStatusUpload)
	if err != nil {
		return
	}
	for _, entity := range uploadList {
		obj := createTradeItemObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		s.addTradeItem(obj)
	}
	//下架中
	withdrawingList, err := dao.GetTradeDao().GetTradeItemList(serverId, tradetypes.TradeStatusWithDrawing)
	if err != nil {
		return
	}
	for _, entity := range withdrawingList {
		obj := createTradeItemObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		s.addTradeItem(obj)
	}

	return
}

//初始化出售的
func (s *tradeService) initSellTradeItems() (err error) {
	serverId := global.GetGame().GetServerIndex()
	//出售的
	soldList, err := dao.GetTradeDao().GetTradeItemList(serverId, tradetypes.TradeStatusSold)
	if err != nil {
		return
	}
	for _, entity := range soldList {
		obj := createTradeItemObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		s.addSellTradeItem(obj)
	}
	return
}

//初始化付款的订单
func (s *tradeService) initOrder() (err error) {
	serverId := global.GetGame().GetServerIndex()
	entityList, err := dao.GetTradeDao().GetTradeOrderList(serverId, tradetypes.TradeOrderStatusInit)
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := createTradeOrderObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		s.addOrder(obj)
	}
	return
}

//初始化未完成的订单
func (s *tradeService) initUnfinishTradeOrders() (err error) {
	serverId := global.GetGame().GetServerIndex()
	entityList, err := dao.GetTradeDao().GetTradeOrderList(serverId, tradetypes.TradeOrderStatusFinish)
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := createTradeOrderObject()
		err = obj.FromEntity(entity)
		if err != nil {
			return
		}
		s.addFinishOrder(obj)
	}
	return
}

//初始化回购
func (s *tradeService) initRecycleObj() (err error) {
	serverId := global.GetGame().GetServerIndex()
	tradeRecyleEntity, err := dao.GetTradeDao().GetTradeRecycle(serverId)
	if err != nil {
		return
	}
	if tradeRecyleEntity == nil {
		now := global.GetGame().GetTimeService().Now()
		s.tradeRecyleObj = createTradeRecycleObject()
		s.tradeRecyleObj.id, _ = idutil.GetId()
		s.tradeRecyleObj.serverId = global.GetGame().GetServerIndex()
		s.tradeRecyleObj.createTime = now
		s.tradeRecyleObj.SetModified()
	} else {
		s.tradeRecyleObj = createTradeRecycleObject()
		err = s.tradeRecyleObj.FromEntity(tradeRecyleEntity)
		if err != nil {
			return
		}
	}
	s.refreshTradeRecycle()
	return
}

func (s *tradeService) refreshTradeRecycle() {
	now := global.GetGame().GetTimeService().Now()
	flag, _ := timeutils.IsSameDay(now, s.tradeRecyleObj.updateTime)
	if flag {
		return
	}

	log.WithFields(
		log.Fields{
			"recycleGold": s.tradeRecyleObj.recycleGold,
		}).Info("trade:交易服务,触发刷新交易回收")

	s.tradeRecyleObj.recycleGold = 0
	s.tradeRecyleObj.updateTime = now
	s.tradeRecyleObj.SetModified()
	s.systemRecycleMap = make(map[int64]*RecyleTimesObject)
	return
}

func (s *tradeService) updateRecycleTime() {
	now := global.GetGame().GetTimeService().Now()
	s.tradeRecyleObj.recycleTime = now
	s.tradeRecyleObj.SetModified()
}

//添加交易商品
func (s *tradeService) addTradeItem(obj *TradeItemObject) {
	_, ok := s.tradeItemListMap[obj.GetPlayerId()]
	if !ok {
		s.tradeItemListMap[obj.GetPlayerId()] = make([]*TradeItemObject, 0, 8)
	}
	s.tradeItemListMap[obj.GetPlayerId()] = append(s.tradeItemListMap[obj.GetPlayerId()], obj)

}

func (s *tradeService) getTradeItemByTradeId(playerId int64, tradeId int64) (index int32, obj *TradeItemObject) {
	itemList, ok := s.tradeItemListMap[playerId]
	if !ok {
		return -1, nil
	}
	for index, it := range itemList {
		if it.GetId() == tradeId {
			return int32(index), it
		}
	}
	return -1, nil
}

//移除交易商品
func (s *tradeService) removeTradeItem(playerId int64, tradeId int64) {
	tradeItemList, ok := s.tradeItemListMap[playerId]
	if !ok {
		return
	}
	index, _ := s.getTradeItemByTradeId(playerId, tradeId)
	if index != -1 {
		s.tradeItemListMap[playerId] = append(tradeItemList[:index], tradeItemList[index+1:]...)
	}
}

func (s *tradeService) addSellTradeItem(obj *TradeItemObject) {
	if obj.GetStatus() != tradetypes.TradeStatusSold {
		return
	}
	s.sellItemMap[obj.GetPlayerId()] = append(s.sellItemMap[obj.GetPlayerId()], obj)
}

func (s *tradeService) getSellTradeItemByTradeId(playerId int64, tradeId int64) (index int32, obj *TradeItemObject) {
	itemList, ok := s.sellItemMap[playerId]
	if !ok {
		return -1, nil
	}
	for index, it := range itemList {
		if it.GetId() == tradeId {
			return int32(index), it
		}
	}
	return -1, nil
}

func (s *tradeService) removeSellTradeItem(playerId, tradeId int64) {
	itemList, ok := s.sellItemMap[playerId]
	if !ok {
		return
	}
	index, _ := s.getSellTradeItemByTradeId(playerId, tradeId)
	if index >= 0 {
		s.sellItemMap[playerId] = append(itemList[:index], itemList[index+1:]...)
	}
}

//获取上架的物品列表
func (s *tradeService) GetTradeList(pl player.Player) []*TradeItemObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.tradeItemListMap[pl.GetId()]
}

//获取可以回收的物品列表
func (s *tradeService) GetCanRecycleTradeList(pl player.Player) []*TradeItemObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.getCanRecycleTradeList(pl)
}

//获取可以回收的物品列表
func (s *tradeService) getCanRecycleTradeList(pl player.Player) []*TradeItemObject {
	now := global.GetGame().GetTimeService().Now()
	recycleTradeItemList := make([]*TradeItemObject, 0, 8)
	tradeItemList := s.tradeItemListMap[pl.GetId()]
	systemRecycleTime := int64(tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().ShangjiaTime)
	for _, tradeObj := range tradeItemList {
		if tradeObj.status != tradetypes.TradeStatusUpload {
			continue
		}
		elapse := now - tradeObj.GetCreateTime()
		//上架时间太少
		if elapse < systemRecycleTime {
			continue
		}
		itemTemplate := item.GetItemService().GetItem(int(tradeObj.GetItemId()))
		if itemTemplate == nil {
			continue
		}
		//没有回购价格
		if itemTemplate.HuigouPrice <= 0 {
			continue
		}
		//回购价格小于交易价格
		if itemTemplate.HuigouPrice < tradeObj.GetGold() {
			continue
		}
		//价格大于系统回购价格最大值
		if tradeObj.GetGold() > tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().HuigouPriceMax {
			continue
		}

		recycleTradeItemList = append(recycleTradeItemList, tradeObj)
	}
	return recycleTradeItemList
}

func (s *tradeService) Start() {

}

func (s *tradeService) Stop() {

}

func (s *tradeService) Heartbeat() {
	s.heartbeatRunner.Heartbeat()
}

//同步商品列表
func (s *tradeService) SyncGlobalTradeList() {
	if !center.GetCenterService().IsTradeOpen() {
		return
	}
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	tradeClient, err := s.getTradeClient()
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
			}).Error("trade:获取商品列表,错误")
		return
	}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	res, err := tradeClient.GetTradeList(timeoutCtx, platform, serverId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
			}).Error("trade:获取商品列表,错误")
		return
	}
	globalTradeItemList, err := ConvertFromGlobalTradeItemList(res.GetTradeItemList())
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err.Error(),
			}).Error("trade:获取商品列表,错误")
		return
	}

	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.resetAllGlobalTradeItemList(globalTradeItemList)
}

//重置
func (s *tradeService) resetAllGlobalTradeItemList(globalTradeItemList []*GlobalTradeItemObject) {
	for _, globalTradeItem := range s.globalTradeItemMap {
		s.removeGlobalTradeItem(globalTradeItem.GetId())
	}
	for _, globalTradeItem := range globalTradeItemList {
		s.addGlobalTradeItem(globalTradeItem)
	}
}

func (s *tradeService) removeGlobalTradeItem(globalTradeId int64) {
	delete(s.globalTradeItemMap, globalTradeId)
}

func (s *tradeService) addGlobalTradeItem(globalTradeItemObject *GlobalTradeItemObject) {
	s.globalTradeItemMap[globalTradeItemObject.GetId()] = globalTradeItemObject
}

const (
	maxTradeItems = 1000
)

//获取所有物品列表
func (s *tradeService) GetGlobalTradeList() []*GlobalTradeItemObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	tradeItemList := make([]*GlobalTradeItemObject, 0, len(s.globalTradeItemMap))
	for _, tradeItem := range s.globalTradeItemMap {
		tradeItemList = append(tradeItemList, tradeItem)
	}
	if len(tradeItemList) >= maxTradeItems {
		tradeItemList = tradeItemList[:maxTradeItems]
	}
	return tradeItemList
}

//获取所有物品列表
func (s *tradeService) GetNumOfGlobalTradeItems() int32 {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return int32(len(s.globalTradeItemMap))
}

//获取所有物品列表
func (s *tradeService) GetGlobalTradeItem(globalTradeId int64) *GlobalTradeItemObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	globalTradeItemObj, ok := s.globalTradeItemMap[globalTradeId]
	if !ok {
		return nil
	}
	return globalTradeItemObj
}

//重试本地商品状态
func (s *tradeService) SyncRetryTradeList() {
	if !center.GetCenterService().IsTradeOpen() {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	for _, tradeItemObjList := range s.tradeItemListMap {
		for _, tradeObj := range tradeItemObjList {
			switch tradeObj.status {
			case tradetypes.TradeStatusInit:
				s.asyncUploadItem(tradeObj)
				break
			case tradetypes.TradeStatusWithDrawing:
				s.asyncWithdrawItem(tradeObj)
				break
			}
		}
	}
}

//重试本地商品状态
func (s *tradeService) SyncRetryOrderList() {
	if !center.GetCenterService().IsTradeOpen() {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	for _, orderItemObj := range s.orderMap {
		s.asyncTradeItem(orderItemObj)
	}
}

const (
//	systemWithdrawTime = int64(3 * common.DAY)
)

//系统下架
func (s *tradeService) SystemWithdrawTradeList() {
	if !center.GetCenterService().IsTradeOpen() {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	systemWithdrawTime := int64(tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().XiajiaTime)
	now := global.GetGame().GetTimeService().Now()
	for _, tradeItemObjList := range s.tradeItemListMap {
		for _, tradeObj := range tradeItemObjList {
			//不是上架
			if tradeObj.GetStatus() != tradetypes.TradeStatusUpload {
				continue
			}

			elapse := now - tradeObj.GetCreateTime()
			if elapse < systemWithdrawTime {
				continue
			}
			s.withdrawItem(tradeObj, true)
		}
	}
}

const (
	maxRecyleNum = 10
)

//系统回购
func (s *tradeService) SystemRecycle() {
	if !center.GetCenterService().IsTradeOpen() {
		return
	}
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//更新回收时间
	defer s.updateRecycleTime()
	tempRecycleTimesObjectList := make([]*RecyleTimesObject, 0, 16)

	tradeConstantTemplate := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate()
	for playerId, _ := range s.tradeItemListMap {
		p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if p == nil {
			continue
		}
		//小于需要的充值数
		if p.GetChargeGoldNum() > int64(tradeConstantTemplate.NeedChongzhi) {
			continue
		}
		canRecycleTradeList := s.getCanRecycleTradeList(p)
		//有没有上架物品
		if len(canRecycleTradeList) == 0 {
			continue
		}
		found := false
		tradeManager := p.GetPlayerDataManager(playertypes.PlayerTradeDataManagerType).(*playertrade.PlayerTradeManager)
		for _, recycleTradeObj := range canRecycleTradeList {
			flag := tradeManager.IfCanRecycle(int64(recycleTradeObj.GetGold()))
			if flag {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		recyleTimesObject, ok := s.systemRecycleMap[playerId]
		if !ok {
			recyleTimesObject = newRecycleTimeObject(playerId)
		}
		tempRecycleTimesObjectList = append(tempRecycleTimesObjectList, recyleTimesObject)

	}
	//优先回购权重高的
	sort.Sort(recycleTimesObjectList(tempRecycleTimesObjectList))
	currentRecycleNum := 0

	for _, recycleTimesObject := range tempRecycleTimesObjectList {
		p := player.GetOnlinePlayerManager().GetPlayerById(recycleTimesObject.playerId)
		if p == nil {
			continue
		}

		gameevent.Emit(tradeeventtypes.TradeEventTypeTradeRecycle, 0, p)
		currentRecycleNum += 1
		if currentRecycleNum >= maxRecyleNum {
			break
		}
	}
}

//系统回购
func (s *tradeService) SystemRecycleTrade(tradeId int64) bool {

	s.rwm.Lock()
	defer s.rwm.Unlock()

	globalTradeItem, ok := s.globalTradeItemMap[tradeId]
	if !ok {
		return false
	}
	s.refreshTradeRecycle()

	//超过回购上限
	totalRecycleGold := s.tradeRecyleObj.recycleGold + int64(globalTradeItem.GetGold())
	gmMaxMoney := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().GmMoneyMax
	if s.tradeRecyleObj.customRecycleGold > 0 {
		gmMaxMoney = int32(s.tradeRecyleObj.customRecycleGold)
	}
	if totalRecycleGold > int64(gmMaxMoney) {
		return false
	}
	//回购
	order, ok := s.orderMap[tradeId]
	if ok {
		return false
	}
	serverId := global.GetGame().GetServerIndex()
	now := global.GetGame().GetTimeService().Now()
	order = createTradeOrderObject()
	order.id, _ = idutil.GetId()
	order.itemId = globalTradeItem.itemId
	order.itemNum = globalTradeItem.num
	order.gold = globalTradeItem.gold
	//系统回购
	order.playerId = 0
	order.playerName = dummytemplate.GetDummyTemplateService().GetRandomDummyName()
	order.serverId = serverId
	order.buyServerId = serverId
	order.tradeId = globalTradeItem.id
	order.gold = globalTradeItem.gold
	order.sellPlatform = globalTradeItem.platform
	order.sellServerId = globalTradeItem.serverId
	order.sellPlayerId = globalTradeItem.playerId
	order.sellPlayerName = globalTradeItem.playerName
	order.propertyData = globalTradeItem.propertyData.Copy()
	order.level = globalTradeItem.level
	order.createTime = now
	order.SetModified()
	s.addOrder(order)
	s.asyncTradeItem(order)

	//添加回购金额
	s.tradeRecyleObj.recycleGold += int64(globalTradeItem.GetGold())
	s.tradeRecyleObj.updateTime = now

	//更新回收次数
	recycleTimesObj, ok := s.systemRecycleMap[globalTradeItem.playerId]
	if !ok {
		recycleTimesObj = newRecycleTimeObject(globalTradeItem.playerId)
	}
	recycleTimesObj.num += 1
	return true
}

//上传商品
func (s *tradeService) UploadItem(pl player.Player, itemId int32, itemNum int32, propertyData inventorytypes.ItemPropertyData, level int32, gold int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//判断上传上限
	tradeItemNumPersonLimit := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().PersonalCountMax
	currentNum := s.getTradeItemNum(pl.GetId())
	if tradeItemNumPersonLimit <= currentNum {
		err = errorTradeItemPersonalNumLimit
		return
	}
	playerId := pl.GetId()
	serverId := global.GetGame().GetServerIndex()
	originServerId := pl.GetServerId()
	playerName := pl.GetName()
	now := global.GetGame().GetTimeService().Now()
	//创建交易
	tradeItemObject := createTradeItemObject()
	tradeItemObject.id, _ = idutil.GetId()
	tradeItemObject.serverId = serverId
	tradeItemObject.originServerId = originServerId
	tradeItemObject.status = tradetypes.TradeStatusInit
	tradeItemObject.itemId = itemId
	tradeItemObject.num = itemNum
	tradeItemObject.propertyData = propertyData
	tradeItemObject.level = level
	tradeItemObject.gold = gold
	tradeItemObject.playerId = playerId
	tradeItemObject.playerName = playerName
	tradeItemObject.createTime = now
	tradeItemObject.SetModified()
	s.addTradeItem(tradeItemObject)
	s.asyncUploadItem(tradeItemObject)
	return
}

//异步上传交易物品
func (s *tradeService) asyncUploadItem(tradeItemObject *TradeItemObject) {
	platform := global.GetGame().GetPlatform()
	serverId := tradeItemObject.GetOriginServerId()
	playerId := tradeItemObject.GetPlayerId()
	playerName := tradeItemObject.GetPlayerName()
	tradeId := tradeItemObject.GetId()
	itemId := tradeItemObject.GetItemId()
	itemNum := tradeItemObject.GetNum()
	propertyDataBytes, err := json.Marshal(tradeItemObject.GetPropertyData())
	if err != nil {
		return
	}
	propertyData := string(propertyDataBytes)
	gold := tradeItemObject.GetGold()
	level := tradeItemObject.GetLevel()
	go func(tradeItemObject *TradeItemObject) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"platform":   platform,
						"playerId":   playerId,
						"playerName": playerName,
						"serverId":   serverId,
						"tradeId":    tradeId,
						"itemId":     itemId,
						"itemNum":    itemNum,
						"gold":       gold,
						"error":      r,
						"stack":      exceptionContent,
					}).Error("trade:上传交易物品,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()
		tradeClient, err := s.getTradeClient()
		if err != nil {
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
					"err":      err,
				}).Warn("trade:上传交易物品失败,错误")
			return
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		res, err := tradeClient.Upload(timeoutCtx, platform, serverId, tradeId, playerId, playerName, itemId, itemNum, gold, propertyData, level)
		if err != nil {
			sta, ok := status.FromError(err)
			if !ok {
				log.WithFields(
					log.Fields{
						"tradeId":  tradeId,
						"playerId": playerId,
						"err":      err,
					}).Warn("trade:上传交易物品失败,错误")
				return
			}
			if sta.Code() != codes.FailedPrecondition {
				log.WithFields(
					log.Fields{
						"tradeId":  tradeId,
						"playerId": playerId,
						"err":      err,
					}).Warn("trade:上传交易物品失败,错误")
				return
			}
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
				}).Info("trade:上架物品,超过最大")
			//交易失败
			s.uploadItemFailed(tradeItemObject)
			return
		}
		globalTradeItem, err := ConvertFromGlobalTradeItem(res.GetTradeItem())
		if err != nil {
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
					"err":      err,
				}).Warn("trade:上传交易物品失败,错误")
			return
		}
		s.uploadItemSuccess(globalTradeItem)
	}(tradeItemObject)
	return
}

func (s *tradeService) uploadItemFailed(tradeItemObj *TradeItemObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := tradeItemObj.GetPlayerId()
	tradeId := tradeItemObj.GetId()
	//上传成功
	flag := tradeItemObj.Refund()
	if !flag {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
				"state":    tradeItemObj.status,
			}).Warn("trade:上传交易物品失败,返还失败")
		return
	}
	s.removeTradeItem(playerId, tradeId)
	//TODO:zrc 做成delegate
	//TODO 发送事件
	gameevent.Emit(tradeeventtypes.TradeEventTypeTradeUploadRefund, tradeItemObj, nil)
}

func (s *tradeService) uploadItemSuccess(globalTradeItem *GlobalTradeItemObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := globalTradeItem.GetPlayerId()
	tradeId := globalTradeItem.GetTradeId()
	globalTradeId := globalTradeItem.GetId()
	_, tradeItemObj := s.getTradeItemByTradeId(playerId, tradeId)
	if tradeItemObj == nil {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
			}).Warn("trade:上传交易物品成功,物品不存在")
		return
	}
	//上传成功
	flag := tradeItemObj.Upload(globalTradeId)
	if !flag {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
				"state":    tradeItemObj.status,
			}).Warn("trade:上传交易物品失败,上传失败")
		return
	}
	s.addGlobalTradeItem(globalTradeItem)
	//TODO:zrc 做成delegate
	//TODO 发送事件
	gameevent.Emit(tradeeventtypes.TradeEventTypeTradeUpload, tradeItemObj, nil)

	//后台日志
	reason := commonlog.TradeLogReasonUpload
	logEventData := tradeeventtypes.CreatePlayerTradeLogEventData(tradeItemObj.playerId, tradeItemObj.itemId, tradeItemObj.num, tradeItemObj.gold, reason, reason.String())
	gameevent.Emit(tradeeventtypes.EventTypeTradeLogUpload, nil, logEventData)
}

//下架商品
func (s *tradeService) WithdrawItem(pl player.Player, tradeId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	_, tradeItemObj := s.getTradeItemByTradeId(playerId, tradeId)
	//交易物品不存在
	if tradeItemObj == nil {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
			}).Warn("trade:下架交易物品,物品不存在")
		err = errorTradeUploadItemNoExist
		return
	}
	s.withdrawItem(tradeItemObj, false)
	return
}

func (s *tradeService) withdrawItem(tradeItemObj *TradeItemObject, system bool) (err error) {
	//下架
	flag := tradeItemObj.WithDrawing(system)
	if !flag {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeItemObj.GetId(),
				"playerId": tradeItemObj.GetPlayerId(),
			}).Warn("trade:下架交易物品,下架失败")
		err = errorTradeUploadItemNoUpload
		return
	}
	s.asyncWithdrawItem(tradeItemObj)
	return
}

//异步下架商品
func (s *tradeService) asyncWithdrawItem(tradeItemObj *TradeItemObject) {
	platform := global.GetGame().GetPlatform()
	playerId := tradeItemObj.GetPlayerId()
	serverId := tradeItemObj.GetServerId()
	tradeId := tradeItemObj.GetId()
	globalTradeId := tradeItemObj.GetGlobalTradeId()

	go func(tradeItemObj *TradeItemObject) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"platform":      platform,
						"serverId":      serverId,
						"tradeId":       tradeId,
						"globalTradeId": globalTradeId,
						"error":         r,
						"stack":         exceptionContent,
					}).Error("trade:下架交易物品,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()
		tradeClient, err := s.getTradeClient()
		if err != nil {
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
					"err":      err,
				}).Warn("trade:下架交易物品,错误")
			//下架失败
			return
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		_, err = tradeClient.Withdraw(timeoutCtx, platform, serverId, globalTradeId)
		if err != nil {
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
					"err":      err,
				}).Warn("trade:下架交易物品,错误")
			//下架失败
			return
		}
		s.withdrawItemSuccess(playerId, tradeId)
	}(tradeItemObj)
	return
}

func (s *tradeService) withdrawItemSuccess(playerId int64, tradeId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, tradeItemObj := s.getTradeItemByTradeId(playerId, tradeId)
	if tradeItemObj == nil {
		return
	}
	//上传成功
	flag := tradeItemObj.WithDraw()
	if !flag {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
			}).Warn("trade:下架交易物品,下架失败")
		return
	}
	//TODO 发送事件
	s.removeTradeItem(playerId, tradeId)
	gameevent.Emit(tradeeventtypes.TradeEventTypeTradeWithdraw, tradeItemObj, nil)

	//后台日志
	reason := commonlog.TradeLogReasonWithdraw
	reasonText := fmt.Sprintf(reason.String(), tradeItemObj.IsSystem())
	logEventData := tradeeventtypes.CreatePlayerTradeLogEventData(tradeItemObj.playerId, tradeItemObj.itemId, tradeItemObj.num, tradeItemObj.gold, reason, reasonText)
	gameevent.Emit(tradeeventtypes.EventTypeTradeLogWithdraw, nil, logEventData)
}

//交易商品
func (s *tradeService) TradeItem(pl player.Player, tradeId int64) (order *TradeOrderObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	globalTradeItem, ok := s.globalTradeItemMap[tradeId]
	if !ok {
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
			}).Warn("trade:交易物品,商品不存在")
		err = errorTradeItemNoExist
		return
	}

	order, ok = s.orderMap[tradeId]
	if ok {
		if order.playerId != pl.GetId() {
			//别人下单了
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
				}).Warn("trade:交易物品,已经被下单了")
			err = errorTradeItemAlreadyOrderOther
		} else {
			//自己已经购买
			log.WithFields(
				log.Fields{
					"tradeId":  tradeId,
					"playerId": playerId,
				}).Warn("trade:交易物品,自己已经下单了")
			err = errorTradeItemAlreadyOrderSelf
		}
		return
	}
	serverId := global.GetGame().GetServerIndex()
	now := global.GetGame().GetTimeService().Now()
	order = createTradeOrderObject()
	order.id, _ = idutil.GetId()
	order.itemId = globalTradeItem.itemId
	order.itemNum = globalTradeItem.num
	order.gold = globalTradeItem.gold
	order.playerId = pl.GetId()
	order.playerName = pl.GetName()
	order.serverId = serverId
	order.buyServerId = pl.GetServerId()
	order.tradeId = globalTradeItem.id
	order.gold = globalTradeItem.gold
	order.sellPlatform = globalTradeItem.platform
	order.sellServerId = globalTradeItem.serverId
	order.sellPlayerId = globalTradeItem.playerId
	order.sellPlayerName = globalTradeItem.playerName
	order.propertyData = globalTradeItem.propertyData.Copy()
	order.level = globalTradeItem.level
	order.createTime = now
	order.SetModified()
	s.addOrder(order)
	s.asyncTradeItem(order)
	return
}

func (s *tradeService) asyncTradeItem(orderObj *TradeOrderObject) {
	platform := global.GetGame().GetPlatform()
	playerId := orderObj.GetPlayerId()
	serverId := orderObj.GetServerId()
	playerName := orderObj.GetPlayerName()
	globalTradeId := orderObj.GetTradeId()
	go func(orderObj *TradeOrderObject) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"platform":      platform,
						"serverId":      serverId,
						"globalTradeId": globalTradeId,
						"playerId":      playerId,
						"error":         r,
						"stack":         exceptionContent,
					}).Error("trade:交易物品,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()
		tradeClient, err := s.getTradeClient()
		if err != nil {
			log.WithFields(
				log.Fields{
					"platform":      platform,
					"serverId":      serverId,
					"globalTradeId": globalTradeId,
					"playerId":      playerId,
				}).Error("trade:交易物品,错误")
			return
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()
		_, err = tradeClient.Trade(timeoutCtx, platform, serverId, playerId, playerName, globalTradeId)
		if err != nil {
			sta, ok := status.FromError(err)
			if !ok {
				log.WithFields(
					log.Fields{
						"platform":      platform,
						"serverId":      serverId,
						"globalTradeId": globalTradeId,
						"playerId":      playerId,
						"err":           err,
					}).Warn("trade:交易物品,错误")
				return
			}
			if sta.Code() != codes.FailedPrecondition {
				log.WithFields(
					log.Fields{
						"platform":      platform,
						"serverId":      serverId,
						"globalTradeId": globalTradeId,
						"playerId":      playerId,
						"err":           err,
					}).Warn("trade:交易物品,错误")
				return
			}
			log.WithFields(
				log.Fields{
					"platform":      platform,
					"serverId":      serverId,
					"globalTradeId": globalTradeId,
					"playerId":      playerId,
				}).Info("trade:交易物品,已经被人购买或者下架了")
			//交易失败
			s.tradeItemFailed(orderObj)
			return
		}
		s.tradeItemSuccess(orderObj)
	}(orderObj)
	return
}

//交易失败
func (s *tradeService) tradeItemFailed(tradeOrderObject *TradeOrderObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	globalTradeId := tradeOrderObject.GetTradeId()
	orderObj := s.getOrder(globalTradeId)
	flag := orderObj.Refund()
	if !flag {
		return
	}
	s.removeOrder(globalTradeId)
	//TODO:退还
	gameevent.Emit(tradeeventtypes.TradeEventTypeTradeItemRefund, orderObj, nil)
	return
}

//交易成功
func (s *tradeService) tradeItemSuccess(tradeOrderObject *TradeOrderObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	globalTradeId := tradeOrderObject.GetTradeId()
	orderObj := s.getOrder(globalTradeId)
	flag := orderObj.Finish()
	if !flag {
		return
	}
	s.removeGlobalTradeItem(globalTradeId)
	s.removeOrder(globalTradeId)
	s.addFinishOrder(orderObj)
	//TODO:发货
	gameevent.Emit(tradeeventtypes.TradeEventTypeTradeItem, orderObj, nil)

	//购买日志
	reason := commonlog.TradeLogReasonTradeItem
	logEventData := tradeeventtypes.CreatePlayerTradeLogEventData(tradeOrderObject.playerId, tradeOrderObject.itemId, tradeOrderObject.itemNum, tradeOrderObject.gold, reason, reason.String())
	gameevent.Emit(tradeeventtypes.EventTypeTradeLogBuy, nil, logEventData)
	return
}

//结束
func (s *tradeService) EndTradeItem(pl player.Player, orderObj *TradeOrderObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	flag := orderObj.End()
	if !flag {
		return
	}
	s.removeUnfinishOrder(pl, orderObj.GetTradeId())
	return
}

//获取未发货的订单
func (s *tradeService) GetUnfinishOrderList(pl player.Player) []*TradeOrderObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	orderList, ok := s.unfinishOrderMap[pl.GetId()]
	if !ok {
		return nil
	}
	return orderList
}

func (s *tradeService) addFinishOrder(obj *TradeOrderObject) {
	s.unfinishOrderMap[obj.GetPlayerId()] = append(s.unfinishOrderMap[obj.GetPlayerId()], obj)
}

func (s *tradeService) removeUnfinishOrder(pl player.Player, globalTradeId int64) {
	unfinishOrderList, ok := s.unfinishOrderMap[pl.GetId()]
	if !ok {
		return
	}
	findIndex := -1
	for index, unfinishOrder := range unfinishOrderList {
		if unfinishOrder.GetTradeId() == globalTradeId {
			findIndex = index
			break
		}
	}
	if findIndex < 0 {
		return
	}
	var remainOrderList []*TradeOrderObject
	if findIndex >= 0 {
		remainOrderList := append(remainOrderList, unfinishOrderList[:findIndex]...)
		remainOrderList = append(remainOrderList, unfinishOrderList[findIndex+1:]...)
		s.unfinishOrderMap[pl.GetId()] = remainOrderList
	}
}

func (s *tradeService) getOrder(tradeId int64) *TradeOrderObject {
	obj, ok := s.orderMap[tradeId]
	if !ok {
		return nil
	}
	return obj
}

func (s *tradeService) addOrder(obj *TradeOrderObject) {
	s.orderMap[obj.GetTradeId()] = obj
}

func (s *tradeService) removeOrder(tradeId int64) {
	delete(s.orderMap, tradeId)
}

//获取交易列表
func (s *tradeService) getTradeItemList(playerId int64) []*TradeItemObject {
	itemList, ok := s.tradeItemListMap[playerId]
	if !ok {
		return nil
	}
	return itemList
}

//获取已上架数量
func (s *tradeService) getTradeItemNum(playerId int64) int32 {
	itemList := s.getTradeItemList(playerId)
	return int32(len(itemList))
}

//卖出物品
func (s *tradeService) SellItem(playerId int64, tradeId int64, buyPlatform int32, buyPlayerServerId int32, buyPlayerId int64, buyPlayerName string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, tradeItem := s.getTradeItemByTradeId(playerId, tradeId)
	if tradeItem == nil {
		//已经卖出去了
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
			}).Warn("trade:出售,已经出售了")
		return
	}

	flag := tradeItem.Sell(buyPlatform, buyPlayerServerId, buyPlayerId, buyPlayerName)
	if !flag {
		//TODO 日志
		log.WithFields(
			log.Fields{
				"tradeId":  tradeId,
				"playerId": playerId,
			}).Warn("trade:出售,出售失败")
		return
	}
	s.removeTradeItem(playerId, tradeId)

	s.addSellTradeItem(tradeItem)
	//发送事件
	gameevent.Emit(tradeeventtypes.TradeEventTypeTradeSellItem, tradeItem, nil)

	//购买日志
	reason := commonlog.TradeLogReasonSellItem
	logEventData := tradeeventtypes.CreatePlayerTradeLogEventData(tradeItem.playerId, tradeItem.itemId, tradeItem.num, tradeItem.gold, reason, reason.String())
	gameevent.Emit(tradeeventtypes.EventTypeTradeLogSell, nil, logEventData)

	return
}

//卖出物品通知
func (s *tradeService) EndSellItem(pl player.Player, tradeItem *TradeItemObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()

	flag := tradeItem.SellNotice()
	if !flag {
		return
	}
	tradeId := tradeItem.GetId()
	s.removeSellTradeItem(playerId, tradeId)
	return
}

//获取未发货的订单
func (s *tradeService) GetSellList(pl player.Player) []*TradeItemObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	sellList, ok := s.sellItemMap[pl.GetId()]
	if !ok {
		return nil
	}
	return sellList
}

//获取未发货的订单
func (s *tradeService) GMSetRecycle(recycle int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	s.tradeRecyleObj.recycleGold = recycle
	s.tradeRecyleObj.updateTime = now
	s.tradeRecyleObj.SetModified()
}

//自定义回购池
func (s *tradeService) GMSetCustomRecycle(recycle int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	s.tradeRecyleObj.customRecycleGold = recycle
	s.tradeRecyleObj.updateTime = now
	s.tradeRecyleObj.SetModified()
}

//自定义回购池
func (s *tradeService) sendRecycleGoldLog() {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	reason := commonlog.TradeLogReasonRecycleGold
	gmMaxMoney := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().GmMoneyMax
	if s.tradeRecyleObj.customRecycleGold > 0 {
		gmMaxMoney = int32(s.tradeRecyleObj.customRecycleGold)
	}
	remainRecycleGold := int64(gmMaxMoney) - s.tradeRecyleObj.recycleGold
	logEventData := tradeeventtypes.CreateTradeLogRecyclGoldEventData(remainRecycleGold, reason, reason.String())
	gameevent.Emit(tradeeventtypes.EventTypeTradeLogRecycleGold, nil, logEventData)
}

var (
	once sync.Once
	cs   *tradeService
)

func Init(cfg *TradeOptions) (err error) {
	once.Do(func() {
		cs = &tradeService{}
		err = cs.init(cfg)
	})
	return err
}

func GetTradeService() TradeService {
	return cs
}
