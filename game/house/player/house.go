package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/house/dao"
	houseeventtypes "fgame/fgame/game/house/event/types"
	housetemplate "fgame/fgame/game/house/template"
	housetypes "fgame/fgame/game/house/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

//玩家房子管理器
type PlayerHouseDataManager struct {
	p player.Player
	//玩家房子对象
	playerHouseMap map[int32]*PlayerHouseObject
}

func (m *PlayerHouseDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerHouseDataManager) Load() (err error) {
	//加载玩家房子信息
	entityList, err := dao.GetHouseDao().GetHouseEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range entityList {
		obj := NewPlayerHouseObject(m.p)
		obj.FromEntity(entity)
		m.playerHouseMap[obj.houseIndex] = obj
	}

	return nil
}

//第一次初始化
func (m *PlayerHouseDataManager) initPlayerHouse() *PlayerHouseObject {
	obj := NewPlayerHouseObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId() //随机主键生成工具
	obj.id = id
	obj.level = 0
	obj.maxLevel = 0
	obj.isBroken = 0
	obj.dayTimes = 0
	obj.lastBrokenTime = now
	obj.refreshUpdateTime = now
	obj.createTime = now
	obj.SetModified() //更新到数据库
	return obj
}

//加载后
func (m *PlayerHouseDataManager) AfterLoad() (err error) {
	m.refreshHouse()
	return nil
}

//心跳
func (m *PlayerHouseDataManager) Heartbeat() {
}

func (m *PlayerHouseDataManager) refreshHouse() {
	// 获取当前时间
	now := global.GetGame().GetTimeService().Now()
	for _, house := range m.playerHouseMap {
		// 过滤激活过，但是已经出售的房子对象
		if !house.IsActivate() {
			continue
		}

		// 判断是否同一天
		flag, err := timeutils.IsSameDay(house.refreshUpdateTime, now)
		if err != nil {
			return
		}
		if !flag {
			// 房子跨天更新前的处理事件
			gameevent.Emit(houseeventtypes.EventTypeHouseBeforeCrossDay, m.p, house)

			// 房子是否损坏
			if house.IfCanBroken(now) {
				house.isBroken = 1
				house.lastBrokenTime = now
			}

			// 重置数据
			house.dayTimes = 0
			house.isRent = 0
			house.refreshUpdateTime = now
			house.SetModified()
		}
	}
	return
}

// 房子是否允许激活
func (m *PlayerHouseDataManager) IsCanActivate(houseIndex int32) (flag bool) {
	if houseIndex == housetypes.InitHouseIndex {
		flag = true
		return
	}

	//房子数量
	maxHouseNum := housetemplate.GetHouseTemplateService().GetHouseConstantTemplate().FangZiCountMax
	if m.GetCurHouseNum() >= maxHouseNum {
		return
	}

	house := m.getHouse(houseIndex)
	if house != nil && house.IsActivate() {
		return
	}

	// 上一个房子满级
	preHouseIndex := houseIndex - 1
	preHous := m.getHouse(preHouseIndex)
	if preHous == nil {
		return
	}

	preHouseTemplate := housetemplate.GetHouseTemplateService().GetHouseTemplate(preHouseIndex, preHous.houseType, preHous.maxLevel)
	if preHouseTemplate == nil {
		return
	}
	if preHouseTemplate.NextId != 0 {
		return
	}

	flag = true
	return
}

//房子激活
func (m *PlayerHouseDataManager) HouseActivate(houseIndex int32, houseType housetypes.HouseType) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	if !m.IsCanActivate(houseIndex) {
		return
	}

	// 房子可重复激活（对象复用）
	house := m.getHouse(houseIndex)
	if house == nil {
		house = m.initPlayerHouse()
		house.houseType = houseType
		house.houseIndex = houseIndex
		house.maxLevel = 1
		m.playerHouseMap[houseIndex] = house
	} else {
		house.houseType = houseType
		house.updateTime = now
	}

	house.level = 1
	house.isBroken = 0
	house.dayTimes = 0
	house.isRent = 0
	house.lastBrokenTime = now
	house.refreshUpdateTime = now
	house.SetModified()

	//房子激活事件
	gameevent.Emit(houseeventtypes.EventTypeHouseActivate, m.p, house)

	flag = true
	return
}

//房子升级
func (m *PlayerHouseDataManager) HouseUplevel(houseIndex int32) (flag bool) {
	// 房子是否存在
	house := m.getHouse(houseIndex)
	if house == nil {
		return
	}

	// 升级次数是否存在
	if !m.IsEnoughUplevelTimes(houseIndex) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	house.level += 1

	// 判断历史最高房子等级
	if house.level > house.maxLevel {
		house.maxLevel = house.level
	}

	//升级次数+1
	house.dayTimes += 1
	house.updateTime = now
	house.SetModified()

	//升级事件：需要在房子升级后做处理的，监听此事件
	gameevent.Emit(houseeventtypes.EventTypeHouseUplevel, m.p, house)
	flag = true
	return
}

//是否有升级次数
func (m *PlayerHouseDataManager) IsEnoughUplevelTimes(houseIndex int32) (flag bool) {
	house := m.getHouse(houseIndex)
	if house == nil {
		return
	}

	maxTimes := housetemplate.GetHouseTemplateService().GetHouseConstantTemplate().UplevLimitCount
	if house.dayTimes >= maxTimes {
		return
	}

	flag = true
	return
}

//房子出售
func (m *PlayerHouseDataManager) HouseSell(houseIndex int32) (flag bool) {
	houseObj := m.getHouse(houseIndex)
	if houseObj == nil {
		return
	}

	houseLevel := houseObj.level
	now := global.GetGame().GetTimeService().Now()
	houseObj.level = 0
	houseObj.updateTime = now
	houseObj.SetModified()

	// 房子出售时间：需要房子出售后做处理的，监听此事件
	// eventData：因为这个时候level等级已经修改了，不能传递houseObj对象
	eventData := houseeventtypes.CreatePlayerHouseSellEventData(houseObj.houseType, houseLevel, houseIndex)
	gameevent.Emit(houseeventtypes.EventTypeHouseSell, m.p, eventData)

	flag = true
	return
}

//房子领取租金
func (m *PlayerHouseDataManager) HouseReceiveRent(houseIndex int32) (flag bool) {
	houseObj := m.getHouse(houseIndex)

	if houseObj == nil {
		return
	}

	if houseObj.IsBroken() {
		return
	}

	if houseObj.IsRent() {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	houseObj.isRent = 1
	houseObj.updateTime = now
	houseObj.SetModified()

	flag = true
	return
}

//房子维修
func (m *PlayerHouseDataManager) HouseRepair(houseIndex int32) (flag bool) {
	house := m.getHouse(houseIndex)
	if house == nil {
		return
	}

	if !house.IsBroken() {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	house.isBroken = 0
	house.updateTime = now
	house.SetModified()

	flag = true
	return
}

//房子信息对象
func (m *PlayerHouseDataManager) GetHouseMap() map[int32]*PlayerHouseObject {
	m.refreshHouse()
	return m.playerHouseMap
}

//当前房子数量
func (m *PlayerHouseDataManager) GetCurHouseNum() int32 {
	num := int32(0)
	for _, house := range m.playerHouseMap {
		if !house.IsActivate() {
			continue
		}
		num += 1
	}

	return num
}

//房子信息
func (m *PlayerHouseDataManager) GetHouse(houseIndex int32) *PlayerHouseObject {
	return m.getHouse(houseIndex)
}

//房子信息
func (m *PlayerHouseDataManager) getHouse(houseIndex int32) *PlayerHouseObject {
	return m.playerHouseMap[houseIndex]
}

//仅gm 房子等级
func (m *PlayerHouseDataManager) GmSetHouseLevel(houseIndex, level int32) {
	house := m.getHouse(houseIndex)
	if house == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if house.maxLevel < level {
		house.maxLevel = level
	}
	house.level = level
	house.updateTime = now
	house.SetModified()

	gameevent.Emit(houseeventtypes.EventTypeHouseUplevel, m.p, house)
}

//仅gm 房子损坏
func (m *PlayerHouseDataManager) GmSetHouseBroken(houseIndex int32) {
	houseObj := m.getHouse(houseIndex)
	if houseObj == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	houseObj.isBroken = 1
	houseObj.lastBrokenTime = now
	houseObj.updateTime = now
	houseObj.SetModified()
}

func CreatePlayerHouseDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerHouseDataManager{}
	m.p = p
	m.playerHouseMap = make(map[int32]*PlayerHouseObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerHouseDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerHouseDataManager))
}
