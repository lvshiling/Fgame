package player

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shopdiscount/dao"
	shopdiscounttypes "fgame/fgame/game/shopdiscount/types"
	"fgame/fgame/pkg/idutil"
)

//玩家商城促销管理器
type PlayerShopDiscountDataManager struct {
	p player.Player
	//玩家商城促销对象
	playerShopDiscountObject *PlayerShopDiscountObject
}

func (m *PlayerShopDiscountDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerShopDiscountDataManager) Load() (err error) {
	//加载玩家商城促销信息
	entity, err := dao.GetShopDiscountDao().GetShopDiscountEntity(m.p.GetId())
	if err != nil {
		return
	}
	if entity == nil {
		m.initPlayerShopDiscountObject()
	} else {
		m.playerShopDiscountObject = NewPlayerShopDiscountObject(m.p)
		m.playerShopDiscountObject.FromEntity(entity)
	}

	return nil
}

//第一次初始化
func (m *PlayerShopDiscountDataManager) initPlayerShopDiscountObject() {
	obj := NewPlayerShopDiscountObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.Id = id
	obj.Typ = shopdiscounttypes.ShopDiscountTypeDefault
	obj.StartTime = 0
	obj.EndTime = 0
	obj.CreateTime = now
	obj.SetModified()

	m.playerShopDiscountObject = obj
}

func (m *PlayerShopDiscountDataManager) GetCurShopDiscountType() shopdiscounttypes.ShopDiscountType {

	now := global.GetGame().GetTimeService().Now()
	if m.playerShopDiscountObject.StartTime < now && now < m.playerShopDiscountObject.EndTime {
		return m.playerShopDiscountObject.Typ
	}

	return shopdiscounttypes.ShopDiscountTypeDefault
}

func (m *PlayerShopDiscountDataManager) SetCurShopDiscountType(typ shopdiscounttypes.ShopDiscountType, startTime int64, endTime int64) {
	if !typ.Valid() {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerShopDiscountObject.Typ = typ
	m.playerShopDiscountObject.StartTime = startTime
	m.playerShopDiscountObject.EndTime = endTime
	m.playerShopDiscountObject.UpdateTime = now
	m.playerShopDiscountObject.SetModified()
	return
}

//加载后
func (m *PlayerShopDiscountDataManager) AfterLoad() (err error) {
	return nil
}

//商城促销信息对象
func (m *PlayerShopDiscountDataManager) GetShopDiscountObj() *PlayerShopDiscountObject {
	return m.playerShopDiscountObject
}

//心跳
func (m *PlayerShopDiscountDataManager) Heartbeat() {

}

func CreatePlayerShopDiscountDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShopDiscountDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerShopDiscountDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShopDiscountDataManager))
}
