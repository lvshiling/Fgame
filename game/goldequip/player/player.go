package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/goldequip/dao"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"

	log "github.com/Sirupsen/logrus"
)

const (
	maxLogLen = 50
)

//玩家元神金装管理器
type PlayerGoldEquipDataManager struct {
	p player.Player
	//元神金装背包
	goldEquipBag *BodyBag
	//分解日志
	logList []*PlayerGoldEquipLogObject
	//金装设置
	equipSettingObj *PlayerGoldEquipSettingObject
	//元神金装数据
	goldEquipObject *PlayerGoldEquipObject
}

func (m *PlayerGoldEquipDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerGoldEquipDataManager) Load() (err error) {
	//加载装备数据
	err = m.loadGoldEquipSlot()
	if err != nil {
		return
	}

	err = m.loadLog()
	if err != nil {
		return
	}
	err = m.loadSetting()
	if err != nil {
		return
	}
	err = m.loadGoldEquipObject()
	if err != nil {
		return
	}
	return nil
}

//加载后
func (m *PlayerGoldEquipDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (m *PlayerGoldEquipDataManager) Heartbeat() {

}

//加载金装日志
func (m *PlayerGoldEquipDataManager) loadLog() (err error) {
	entityList, err := dao.GetGoldEquipDao().GetPlayerGoldEquipLogEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range entityList {
		logObj := NewPlayerGoldEquipLogObject(m.p)
		logObj.FromEntity(entity)
		m.logList = append(m.logList, logObj)
	}

	return
}

//加载金装设置
func (m *PlayerGoldEquipDataManager) loadSetting() (err error) {
	entity, err := dao.GetGoldEquipDao().GetPlayerGoldEquipSettingEntity(m.p.GetId())
	if err != nil {
		return
	}

	if entity != nil {
		obj := NewPlayerGoldEquipSettingObject(m.p)
		obj.FromEntity(entity)
		m.equipSettingObj = obj
	} else {
		m.initEquipSeting()
	}

	return
}

//加载金装设置
func (m *PlayerGoldEquipDataManager) loadGoldEquipObject() (err error) {
	entity, err := dao.GetGoldEquipDao().GetPlayerGoldEquipEntity(m.p.GetId())
	if err != nil {
		return
	}

	if entity != nil {
		obj := NewPlayerGoldEquipObject(m.p)
		obj.FromEntity(entity)
		m.goldEquipObject = obj
	} else {
		m.initGoldEquipObject()
	}

	return
}

// 初始化设置
func (m *PlayerGoldEquipDataManager) initEquipSeting() {
	obj := NewPlayerGoldEquipSettingObject(m.p)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.fenJieIsAuto = 0
	obj.fenJieQuality = 0
	//zrc: 修改过的
	//TODO:cjb 默认是检测过的,看完删除注释
	obj.isCheckOldSt = int32(0)
	obj.createTime = now
	obj.SetModified()

	m.equipSettingObj = obj
	return
}

// 初始化设置
func (m *PlayerGoldEquipDataManager) initGoldEquipObject() {
	obj := NewPlayerGoldEquipObject(m.p)
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.power = 0
	obj.createTime = now
	obj.SetModified()

	m.goldEquipObject = obj
	return
}

//获取金装背包
func (m *PlayerGoldEquipDataManager) GetGoldEquipBag() *BodyBag {
	return m.goldEquipBag
}

//加载身上金装
func (m *PlayerGoldEquipDataManager) loadGoldEquipSlot() (err error) {
	//加载金装槽位
	goldEquipSlotList, err := dao.GetGoldEquipDao().GetGoldEquipSlotList(m.p.GetId())
	if err != nil {
		return
	}
	slotList := make([]*PlayerGoldEquipSlotObject, 0, len(goldEquipSlotList))
	for _, slot := range goldEquipSlotList {
		pio := NewPlayerGoldEquipSlotObject(m.p)
		err := pio.FromEntity(slot)
		if err != nil {
			return err
		}
		slotList = append(slotList, pio)
	}

	m.fixUpstarLevel(slotList)
	m.goldEquipBag = createBodyBag(m.p, slotList)
	return
}

// 修正升星强化等级
func (m *PlayerGoldEquipDataManager) fixUpstarLevel(itemObjList []*PlayerGoldEquipSlotObject) {
	for _, itemObj := range itemObjList {
		if itemObj.IsEmpty() {
			continue
		}

		goldequipData, ok := itemObj.propertyData.(*goldequiptypes.GoldEquipPropertyData)
		if !ok {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(itemObj.itemId))
		if itemTemp.GetGoldEquipTemplate() == nil {
			log.Info("itemid:", itemObj.itemId)
			continue
		}
		maxLeve := itemTemp.GetGoldEquipTemplate().GetMaxUpstarLevel()
		goldequipData.FixUpstarLevel(maxLeve)
		itemObj.SetModified()
	}
}

//获取装备
func (m *PlayerGoldEquipDataManager) GetGoldEquipByPos(pos inventorytypes.BodyPositionType) *PlayerGoldEquipSlotObject {
	item := m.goldEquipBag.GetByPosition(pos)
	if item == nil {
		return nil
	}

	return item
}

//使用装备
func (m *PlayerGoldEquipDataManager) PutOn(pos inventorytypes.BodyPositionType, itemId int32, level int32, bind itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) (flag bool) {
	flag = m.goldEquipBag.PutOn(pos, itemId, level, bind, propertyData)
	if flag {
		gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipPutOn, m.p, itemId)
	}

	return
}

//脱下装备
func (m *PlayerGoldEquipDataManager) TakeOff(pos inventorytypes.BodyPositionType) (itemId int32) {
	//判断是否可以脱下
	flag := m.IfCanTakeOff(pos)
	if !flag {
		return
	}

	slot := m.goldEquipBag.GetByPosition(pos)
	data, _ := slot.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	openlightlevel := data.OpenLightLevel
	strengthlevel := slot.newStLevel
	upstarlevel := slot.level

	itemId = m.goldEquipBag.TakeOff(pos)
	if itemId > 0 {
		gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipTakeOff, m.p, itemId)

		eventData := goldequipeventtypes.CreatePlayerGoldEquipStatusEventData(pos, openlightlevel, strengthlevel, upstarlevel)
		gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipStatusWhenTakeOff, m.p, eventData)
	}
	return
}

//获取套装数量
func (m *PlayerGoldEquipDataManager) GetGoldEquipGroupNum() map[int32]int32 {
	curGroupMap := make(map[int32]int32)
	for _, slot := range m.goldEquipBag.GetAll() {
		if slot.IsEmpty() {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(slot.GetItemId()))
		groupId := itemTemp.GetGoldEquipTemplate().SuitGroup
		if groupId == 0 {
			continue
		}

		_, ok := curGroupMap[groupId]
		if ok {
			curGroupMap[groupId] += int32(1)
		} else {
			curGroupMap[groupId] = int32(1)
		}
	}

	return curGroupMap
}

//装备改变
func (pidm *PlayerGoldEquipDataManager) GetChangedEquipmentSlotAndReset() (itemList []*PlayerGoldEquipSlotObject) {
	return pidm.goldEquipBag.GetChangedSlotAndReset()
}

//是否可以卸下
func (m *PlayerGoldEquipDataManager) IfCanTakeOff(pos inventorytypes.BodyPositionType) bool {
	item := m.GetGoldEquipByPos(pos)
	if item == nil {
		return false
	}
	if item.IsEmpty() {
		return false
	}
	return true
}

//开光
func (m *PlayerGoldEquipDataManager) OpenLight(pos inventorytypes.BodyPositionType, isSuccess bool) bool {
	item := m.GetGoldEquipByPos(pos)
	if item == nil {
		return false
	}
	if item.IsEmpty() {
		return false
	}

	propertyData := item.propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if isSuccess {
		propertyData.OpenLightLevel += 1
		propertyData.OpenTimes = 0
	} else {
		propertyData.OpenTimes += 1
	}

	now := global.GetGame().GetTimeService().Now()
	item.updateTime = now
	item.SetModified()

	return true
}

//获取强化总等级
func (m *PlayerGoldEquipDataManager) CountTotalUpstarLevel() int32 {
	slotList := m.goldEquipBag.GetAll()

	totalLevel := int32(0)
	for _, slot := range slotList {
		totalLevel += slot.newStLevel
	}

	return totalLevel
}

//获取镶嵌宝石总等级
func (m *PlayerGoldEquipDataManager) CountTotalGemLevel() int32 {
	slotList := m.goldEquipBag.GetAll()

	totalLevel := int32(0)
	for _, slot := range slotList {
		for _, itemId := range slot.GemInfo {
			itemTemp := item.GetItemService().GetItem(int(itemId))
			totalLevel += itemTemp.TypeFlag2
		}
	}

	return totalLevel
}

func (m *PlayerGoldEquipDataManager) ToGoldEquipSlotList() (slotInfoList []*goldequiptypes.GoldEquipSlotInfo) {
	for _, slot := range m.goldEquipBag.GetAll() {
		slotInfo := &goldequiptypes.GoldEquipSlotInfo{}
		slotInfo.SlotId = int32(slot.GetSlotId())
		slotInfo.Level = slot.GetLevel()
		slotInfo.NewStLevel = slot.GetNewStLevel()
		slotInfo.ItemId = slot.GetItemId()
		slotInfo.GemUnlockInfo = slot.GemUnlockInfo
		slotInfo.Gems = slot.GemInfo
		slotInfo.CastingSpiritInfo = slot.CastingSpiritInfo
		slotInfo.ForgeSoulInfo = slot.ForgeSoulInfo
		slotInfoList = append(slotInfoList, slotInfo)
		data, ok := slot.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)
		if !ok {
			//TODO xzk:临时处理bug
			slotInfo.PropertyData = goldequiptypes.NewGoldEquipPropertyData()
			slotInfo.PropertyData.InitBase()
		} else {
			slotInfo.PropertyData = data
		}
	}
	return
}

//
func (m *PlayerGoldEquipDataManager) AddGoldEquipLog(fenJieItemIdList []int32, rewItemStr string) {
	now := global.GetGame().GetTimeService().Now()
	var obj *PlayerGoldEquipLogObject
	if len(m.logList) >= int(maxLogLen) {
		obj = m.logList[0]
		m.logList = m.logList[1:]
	} else {
		obj = NewPlayerGoldEquipLogObject(m.p)
		id, _ := idutil.GetId()
		obj.id = id
		obj.createTime = now
	}

	obj.fenJieItemIdList = fenJieItemIdList
	obj.rewItemStr = rewItemStr
	obj.updateTime = now
	obj.SetModified()

	m.logList = append(m.logList, obj)
}

// 获取金装日志列表
func (m *PlayerGoldEquipDataManager) GetLogList() []*PlayerGoldEquipLogObject {
	return m.logList
}

//设置自动分解
func (m *PlayerGoldEquipDataManager) SetAutoFenJie(isAuto int32, quality itemtypes.ItemQualityType, zhuanShu int32) {
	now := global.GetGame().GetTimeService().Now()
	m.equipSettingObj.fenJieIsAuto = isAuto
	m.equipSettingObj.fenJieQuality = quality
	m.equipSettingObj.fenJieZhuanShu = zhuanShu
	m.equipSettingObj.updateTime = now
	m.equipSettingObj.SetModified()

	// TODO: xzk25 后台日志
}

//设置自动分解
func (m *PlayerGoldEquipDataManager) GetGoldEquipSetting() *PlayerGoldEquipSettingObject {
	return m.equipSettingObj
}

//获取特殊技能
func (m *PlayerGoldEquipDataManager) GetTeShuSkillList() (skillList []*scene.TeshuSkillObject) {
	for _, obj := range m.goldEquipBag.GetAll() {
		if obj.IsEmpty() {
			continue
		}
		itemTemplate := item.GetItemService().GetItem(int(obj.itemId))
		if itemTemplate == nil {
			continue
		}
		goldequipTemplate := itemTemplate.GetGoldEquipTemplate()
		if goldequipTemplate == nil {
			continue
		}
		if !goldequipTemplate.IsGodCastingEquip() {
			continue
		}
	Loop:
		for soulType, info := range obj.ForgeSoulInfo {
			forgeSoulTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetForgeSoulTemplate(obj.GetSlotId(), soulType)
			if forgeSoulTemplate == nil {
				continue
			}
			soulLevelTemplate := forgeSoulTemplate.GetLevelTemplate(info.Level)
			if soulLevelTemplate == nil {
				continue
			}

			for _, skillObj := range skillList {
				if skillObj.GetSkillId() == forgeSoulTemplate.GetTeshuSkillTemp().SkillId {
					skillObj.AddRate(soulLevelTemplate.ChufaRate, soulLevelTemplate.DikangRate)
					continue Loop
				}
			}

			skillObj := scene.CreateTeshuSkillObject(forgeSoulTemplate.GetTeshuSkillTemp().SkillId, soulLevelTemplate.ChufaRate, soulLevelTemplate.DikangRate)
			skillList = append(skillList, skillObj)
		}
	}
	return skillList
}

//获取特殊技能
func (m *PlayerGoldEquipDataManager) UplevelSoul(bodyPos inventorytypes.BodyPositionType, soulType goldequiptypes.ForgeSoulType, sucess bool) {
	m.goldEquipBag.UplevelSoul(bodyPos, soulType, sucess)
}

// 获取元神金装战力
func (m *PlayerGoldEquipDataManager) GetPower() int64 {
	return m.goldEquipObject.power
}

// 设置元神金装战力
func (m *PlayerGoldEquipDataManager) SetPower(power int64) {
	if power <= 0 {
		return
	}
	if m.goldEquipObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.goldEquipObject.power = power
	m.goldEquipObject.updateTime = now
	m.goldEquipObject.SetModified()
}

func CreatePlayerGoldEquipDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerGoldEquipDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerGoldEquipDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerGoldEquipDataManager))
}
