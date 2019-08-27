package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/lucky/dao"
	luckyeventtypes "fgame/fgame/game/lucky/event/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"fmt"
)

//玩家幸运符管理器
type PlayerLuckyDataManager struct {
	p player.Player
	//玩家幸运符对象
	playerLuckyObjectMap map[itemtypes.ItemType]map[itemtypes.ItemSubType]*PlayerLuckyObject
}

func (m *PlayerLuckyDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerLuckyDataManager) Load() (err error) {
	//加载玩家幸运符信息
	luckyEntityList, err := dao.GetLuckyDao().GetLuckyEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range luckyEntityList {
		obj := NewPlayerLuckyObject(m.p)
		obj.FromEntity(entity)

		err = m.addLuckyObj(obj)
		if err != nil {
			return
		}
	}

	return
}

//加载后
func (m *PlayerLuckyDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerLuckyDataManager) Heartbeat() {}

//添加幸运类型
func (m *PlayerLuckyDataManager) AddLuckyType(itemId, num int32) (err error) {
	itemTemp := item.GetItemService().GetItem(int(itemId))
	typ := itemTemp.GetItemType()
	subType := itemTemp.GetItemSubType()
	now := global.GetGame().GetTimeService().Now()

	obj := m.getLuckObj(typ, subType)
	if obj == nil {
		obj = NewPlayerLuckyObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.createTime = now
		obj.itemId = itemId
		obj.typ = typ
		obj.subType = subType

		err = m.addLuckyObj(obj)
		if err != nil {
			return
		}
	}

	// 物品id不同，时间重置
	expireTime := int64(itemTemp.TypeFlag2) * int64(num)
	curExpire := obj.expireTime
	if now > curExpire || itemId != obj.itemId {
		curExpire = now + expireTime
	} else {
		curExpire += expireTime
	}

	obj.itemId = itemId
	obj.expireTime = curExpire
	obj.updateTime = now
	obj.SetModified()

	gameevent.Emit(luckyeventtypes.EventTypeLuckyAdd, m.p, obj)

	return
}

//获取幸运率加成
func (m *PlayerLuckyDataManager) GetIncrSuccessRate(typ itemtypes.ItemType, subType itemtypes.ItemSubType) (rate int32) {
	obj := m.getLuckObj(typ, subType)
	if obj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if now > obj.expireTime {
		return
	}

	temp := item.GetItemService().GetItem(int(obj.itemId))
	return temp.TypeFlag1
}

func (m *PlayerLuckyDataManager) GetLuckyExpireTime(typ itemtypes.ItemType, subType itemtypes.ItemSubType) int64 {
	obj := m.getLuckObj(typ, subType)
	if obj == nil {
		return 0
	}

	return obj.expireTime
}

func (m *PlayerLuckyDataManager) GetAllLuckyObj() map[itemtypes.ItemType]map[itemtypes.ItemSubType]*PlayerLuckyObject {
	return m.playerLuckyObjectMap
}

func (m *PlayerLuckyDataManager) GetLuckyInfoAll() map[int32]int64 {
	now := global.GetGame().GetTimeService().Now()
	newMap := make(map[int32]int64)
	for _, subMap := range m.playerLuckyObjectMap {
		for _, obj := range subMap {
			expire := obj.GetExpireTime()
			itemId := obj.GetItemId()

			if now > expire {
				continue
			}

			newMap[itemId] = expire
		}
	}
	return newMap
}

func (m *PlayerLuckyDataManager) getLuckObj(typ itemtypes.ItemType, subType itemtypes.ItemSubType) *PlayerLuckyObject {
	subMap, ok := m.playerLuckyObjectMap[typ]
	if !ok {
		return nil
	}
	obj, ok := subMap[subType]
	if !ok {
		return nil
	}

	return obj
}

func (m *PlayerLuckyDataManager) addLuckyObj(obj *PlayerLuckyObject) (err error) {
	subMap, ok := m.playerLuckyObjectMap[obj.typ]
	if !ok {
		subMap = make(map[itemtypes.ItemSubType]*PlayerLuckyObject)
		m.playerLuckyObjectMap[obj.typ] = subMap
	}

	_, ok = subMap[obj.subType]
	if ok {
		return fmt.Errorf("幸运类型重复, type:%d,subType:%d", obj.typ, obj.subType)
	}
	subMap[obj.subType] = obj
	return
}

func CreatePlayerLuckyDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerLuckyDataManager{}
	m.p = p
	m.playerLuckyObjectMap = make(map[itemtypes.ItemType]map[itemtypes.ItemSubType]*PlayerLuckyObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerLuckyDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerLuckyDataManager))
}
