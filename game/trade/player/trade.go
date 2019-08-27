package player

import (
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/trade/dao"
	tradetemplate "fgame/fgame/game/trade/template"
	tradetypes "fgame/fgame/game/trade/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家交易管理器
type PlayerTradeManager struct {
	p                        player.Player
	playerTradeLogObjectList []*PlayerTradeLogObject
	playerTradeRecycleObject *PlayerTradeRecycleObject
}

func (m *PlayerTradeManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerTradeManager) Load() (err error) {
	//加载玩家交易日志信息
	tradeLogEntityList, err := dao.GetTradeDao().GetTradeLogList(m.p.GetId())
	if err != nil {
		return
	}
	for _, tradeLogEntity := range tradeLogEntityList {
		tradeLog := createPlayerTradeLogObject(m.p)
		err = tradeLog.FromEntity(tradeLogEntity)
		if err != nil {
			return
		}
		m.playerTradeLogObjectList = append(m.playerTradeLogObjectList, tradeLog)
	}
	playerTradeRecycleEntity, err := dao.GetTradeDao().GetPlayerTradeRecycle(m.p.GetId())
	if err != nil {
		return
	}
	if playerTradeRecycleEntity == nil {
		m.initPlayerTradeRecycleObject()
	} else {
		m.playerTradeRecycleObject = createPlayerTradeRecycleObject(m.p)
		m.playerTradeRecycleObject.FromEntity(playerTradeRecycleEntity)
	}
	return nil
}

func (m *PlayerTradeManager) initPlayerTradeRecycleObject() {
	o := createPlayerTradeRecycleObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.createTime = now
	o.SetModified()
	m.playerTradeRecycleObject = o
}

//加载后
func (m *PlayerTradeManager) AfterLoad() (err error) {
	m.refreshTradeRecycle()
	return nil
}

//心跳
func (m *PlayerTradeManager) Heartbeat() {}

//获取交易日志
func (m *PlayerTradeManager) GetTradeLogList() []*PlayerTradeLogObject {
	return m.playerTradeLogObjectList
}

func (m *PlayerTradeManager) AddSellLog(
	tradeId int64,
	sellServerId int32,
	sellPlayerId int64,
	sellPlayerName string,
	buyServerId int32,
	buyPlayerId int64,
	buyPlayerName string,
	gold int32,
	getGold int32,
	fee int32,
	itemId int32,
	itemNum int32,
	propertyData inventorytypes.ItemPropertyData,
	level int32,
	sellTime int64,
	feeRate int32,
) {
	logObj := createPlayerTradeLogObject(m.p)
	logObj.id, _ = idutil.GetId()
	logObj.logType = tradetypes.TradeLogTypeSell
	logObj.tradeId = tradeId
	logObj.sellPlayerId = sellPlayerId
	logObj.sellServerId = sellServerId
	logObj.sellPlayerName = sellPlayerName
	logObj.buyServerId = buyServerId
	logObj.buyPlayerId = buyPlayerId
	logObj.buyPlayerName = buyPlayerName
	logObj.gold = gold
	logObj.getGold = getGold
	logObj.fee = fee
	logObj.itemId = itemId
	logObj.itemNum = itemNum
	logObj.propertyData = propertyData
	logObj.level = level
	logObj.feeRate = feeRate
	logObj.createTime = sellTime
	logObj.SetModified()
	m.playerTradeLogObjectList = append(m.playerTradeLogObjectList, logObj)
}

func (m *PlayerTradeManager) AddBuyLog(
	tradeId int64,
	sellServerId int32,
	sellPlayerId int64,
	sellPlayerName string,
	buyServerId int32,
	buyPlayerId int64,
	buyPlayerName string,
	gold int32,
	getGold int32,
	fee int32,
	itemId int32,
	itemNum int32,
	propertyData inventorytypes.ItemPropertyData,
	level int32,
	sellTime int64,
) {
	logObj := createPlayerTradeLogObject(m.p)
	logObj.id, _ = idutil.GetId()
	logObj.logType = tradetypes.TradeLogTypeBuy
	logObj.tradeId = tradeId
	logObj.sellPlayerId = sellPlayerId
	logObj.sellServerId = sellServerId
	logObj.sellPlayerName = sellPlayerName
	logObj.buyServerId = buyServerId
	logObj.buyPlayerId = buyPlayerId
	logObj.buyPlayerName = buyPlayerName
	logObj.gold = gold
	logObj.getGold = getGold
	logObj.fee = fee
	logObj.itemId = itemId
	logObj.itemNum = itemNum
	logObj.level = level
	logObj.propertyData = propertyData
	logObj.createTime = sellTime
	logObj.SetModified()
	m.playerTradeLogObjectList = append(m.playerTradeLogObjectList, logObj)
}

func (m *PlayerTradeManager) refreshTradeRecycle() {
	now := global.GetGame().GetTimeService().Now()
	flag, _ := timeutils.IsSameDay(now, m.playerTradeRecycleObject.updateTime)
	if flag {
		return
	}
	m.playerTradeRecycleObject.recycleGold = 0
	m.playerTradeRecycleObject.updateTime = now
	m.playerTradeRecycleObject.SetModified()
}

func (m *PlayerTradeManager) IfCanRecycle(gold int64) bool {
	m.refreshTradeRecycle()
	maxPersonalGold := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().PlayerMoneyMax
	if m.playerTradeRecycleObject.recycleGold+gold > int64(maxPersonalGold) {
		return false
	}
	return true
}

func (m *PlayerTradeManager) AddRecycleGold(gold int64) bool {
	m.refreshTradeRecycle()
	currentGold := m.playerTradeRecycleObject.recycleGold + gold
	maxPersonalGold := tradetemplate.GetTradeTemplateService().GetTradeConstantTemplate().PlayerMoneyMax
	if currentGold > int64(maxPersonalGold) {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerTradeRecycleObject.recycleGold = currentGold
	m.playerTradeRecycleObject.updateTime = now
	m.playerTradeRecycleObject.SetModified()
	return true
}

func (m *PlayerTradeManager) GMSetRecycle(recycle int64) {
	now := global.GetGame().GetTimeService().Now()
	m.playerTradeRecycleObject.recycleGold = recycle
	m.playerTradeRecycleObject.updateTime = now
	m.playerTradeRecycleObject.SetModified()

}

func CreatePlayerTradeDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerTradeManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTradeDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTradeDataManager))
}
