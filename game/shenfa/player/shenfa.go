package player

import (
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/shenfa/dao"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	shenfatypes "fgame/fgame/game/shenfa/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"
)

//玩家身法管理器
type PlayerShenfaDataManager struct {
	p player.Player
	//玩家身法对象
	playerShenfaObject *PlayerShenfaObject
	//玩家非进阶身法对象
	playerOtherMap map[shenfatypes.ShenfaType]map[int32]*PlayerShenfaOtherObject
}

func (m *PlayerShenfaDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerShenfaDataManager) Load() (err error) {

	m.playerOtherMap = make(map[shenfatypes.ShenfaType]map[int32]*PlayerShenfaOtherObject)
	//加载玩家身法信息
	shenfaEntity, err := dao.GetShenfaDao().GetShenfaEntity(m.p.GetId())
	if err != nil {
		return
	}
	if shenfaEntity == nil {
		m.initPlayerShenfaObject()
	} else {
		m.playerShenfaObject = NewPlayerShenfaObject(m.p)
		m.playerShenfaObject.FromEntity(shenfaEntity)
	}

	//加载玩家非进阶身法信息
	shenfaOtherList, err := dao.GetShenfaDao().GetShenfaOtherList(m.p.GetId())
	if err != nil {
		return
	}

	//非进阶身法信息
	for _, shenfaOther := range shenfaOtherList {
		pwo := NewPlayerShenfaOtherObject(m.p)
		pwo.FromEntity(shenfaOther)

		typ := shenfatypes.ShenfaType(shenfaOther.Typ)
		playerOtherMap, exist := m.playerOtherMap[typ]
		if !exist {
			playerOtherMap = make(map[int32]*PlayerShenfaOtherObject)
			m.playerOtherMap[typ] = playerOtherMap
		}
		playerOtherMap[pwo.ShenFaId] = pwo

	}

	return
}

//第一次初始化
func (m *PlayerShenfaDataManager) initPlayerShenfaObject() {
	pwo := NewPlayerShenfaObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	advanceId := playerCreateTemplate.Shenfa
	//生成id
	pwo.AdvanceId = int(advanceId)
	pwo.ShenfaId = int32(0)
	pwo.UnrealLevel = int32(0)
	pwo.UnrealNum = int32(0)
	pwo.UnrealPro = int32(0)
	pwo.UnrealList = make([]int, 0, 8)
	pwo.TimesNum = int32(0)
	pwo.Bless = int32(0)
	pwo.BlessTime = int64(0)
	pwo.Hidden = 0
	pwo.Power = int64(0)
	pwo.CreateTime = now
	m.playerShenfaObject = pwo
	pwo.SetModified()
}

//增加非进阶身法
func (m *PlayerShenfaDataManager) newPlayerShenfaOtherObject(typ shenfatypes.ShenfaType, shenfaId int32) (err error) {

	playerOtherMap, exist := m.playerOtherMap[typ]
	if !exist {
		playerOtherMap = make(map[int32]*PlayerShenfaOtherObject)
		m.playerOtherMap[typ] = playerOtherMap
	}
	_, exist = playerOtherMap[shenfaId]
	if exist {
		return
	}
	pwo := NewPlayerShenfaOtherObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.Id = id
	//生成id
	pwo.Typ = typ
	pwo.ShenFaId = shenfaId
	pwo.Level = 0
	pwo.UpNum = 0
	pwo.UpPro = 0
	pwo.CreateTime = now
	playerOtherMap[shenfaId] = pwo
	pwo.SetModified()
	return
}

//加载后
func (m *PlayerShenfaDataManager) AfterLoad() (err error) {
	m.refreshBless()
	return
}

func (m *PlayerShenfaDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(m.playerShenfaObject.AdvanceId)
	nextNumber := number + 1
	shenFaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(nextNumber)
	if shenFaTemplate == nil {
		return
	}
	if !shenFaTemplate.GetIsClear() {
		return
	}
	lastTime := m.playerShenfaObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			m.playerShenfaObject.Bless = 0
			m.playerShenfaObject.BlessTime = 0
			m.playerShenfaObject.TimesNum = 0
			m.playerShenfaObject.SetModified()
		}
	}
}

//身法信息对象
func (m *PlayerShenfaDataManager) GetShenfaInfo() *PlayerShenfaObject {
	m.refreshBless()
	return m.playerShenfaObject
}

func (m *PlayerShenfaDataManager) GetShenfaAdvanced() int32 {
	return int32(m.playerShenfaObject.AdvanceId)
}

func (m *PlayerShenfaDataManager) GetShenFaId() int32 {
	if m.playerShenfaObject.Hidden != 0 {
		return 0
	}
	if m.playerShenfaObject.ShenfaId != 0 {
		return m.playerShenfaObject.ShenfaId
	}
	shenFaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(m.playerShenfaObject.AdvanceId))
	if shenFaTemplate == nil {
		return 0
	}
	return int32(shenFaTemplate.TemplateId())
}

//获取玩家非进阶身法对象
func (m *PlayerShenfaDataManager) GetShenfaOtherMap() map[shenfatypes.ShenfaType]map[int32]*PlayerShenfaOtherObject {
	return m.playerOtherMap
}

//能吃几个幻化丹
func (m *PlayerShenfaDataManager) CanEatUnrealDanNum() int32 {
	advancedId := m.playerShenfaObject.AdvanceId
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(advancedId)
	shidanLimit := shenfaTemplate.ShidanLimit
	currentNum := m.playerShenfaObject.UnrealNum
	if currentNum >= shidanLimit {
		return 0
	}
	totalNum := shidanLimit - currentNum
	return totalNum
}

// //食幻化丹
// func (m *PlayerShenfaDataManager) EatUnrealDan(pro int32, sucess bool) (err error) {
// 	if pro < 0 {
// 		return
// 	}

// 	if sucess {
// 		m.playerShenfaObject.UnrealLevel += 1
// 		m.playerShenfaObject.UnrealNum = 0
// 		m.playerShenfaObject.UnrealPro = pro
// 	} else {
// 		m.playerShenfaObject.UnrealNum += 1
// 		m.playerShenfaObject.UnrealPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerShenfaObject.UpdateTime = now
// 	m.playerShenfaObject.SetModified()
// 	return
// }

//食幻化丹
func (m *PlayerShenfaDataManager) EatUnrealDan(level int32) (err error) {
	if m.playerShenfaObject.UnrealLevel == level || level <= 0 {
		return
	}

	hunaHuaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaHuanHuaTemplate(level)
	if hunaHuaTemplate == nil {
		return
	}
	m.playerShenfaObject.UnrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	return
}

//心跳
func (m *PlayerShenfaDataManager) Heartbeat() {

}

func (m *PlayerShenfaDataManager) IsWingHidden() bool {
	return m.playerShenfaObject.Hidden == 1
}

//进阶
func (m *PlayerShenfaDataManager) ShenfaAdvanced(pro, addTimes int32, sucess bool) (err error) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextAdvancedId := m.playerShenfaObject.AdvanceId + 1
		shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(nextAdvancedId)
		if shenfaTemplate == nil {
			return
		}
		m.playerShenfaObject.AdvanceId += 1
		m.playerShenfaObject.TimesNum = 0
		m.playerShenfaObject.Bless = 0
		m.playerShenfaObject.BlessTime = 0
		m.playerShenfaObject.ShenfaId = int32(0)

		gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)
		gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvanced, m.p, int32(m.playerShenfaObject.AdvanceId))
	} else {
		m.playerShenfaObject.TimesNum += addTimes
		if m.playerShenfaObject.Bless == 0 {
			m.playerShenfaObject.BlessTime = now
		}
		m.playerShenfaObject.Bless += pro
	}
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	return
}

//直升券进阶
func (m *PlayerShenfaDataManager) ShenfaAdvancedTicket(addAdvancedNum int32) (err error) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := m.playerShenfaObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(nextAdvancedId)
		if shenfaTemplate == nil {
			return
		}
		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.AdvanceId += canAddNum
	m.playerShenfaObject.TimesNum = 0
	m.playerShenfaObject.Bless = 0
	m.playerShenfaObject.BlessTime = 0
	m.playerShenfaObject.ShenfaId = int32(0)
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvanced, m.p, int32(m.playerShenfaObject.AdvanceId))
	return
}

//非进阶领域激活
func (m *PlayerShenfaDataManager) ShenFaOtherGet(shenfaId int32) {
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenfaId))
	if shenfaTemplate == nil {
		return
	}
	typ := shenfaTemplate.GetTyp()
	if typ == shenfatypes.ShenfaTypeAdvanced {
		return
	}
	m.newPlayerShenfaOtherObject(typ, shenfaId)
	return
}

//设置幻化信息
func (m *PlayerShenfaDataManager) AddUnrealInfo(shenfaId int) (err error) {
	if shenfaId <= 0 {
		return
	}
	m.playerShenfaObject.UnrealList = append(m.playerShenfaObject.UnrealList, shenfaId)
	sort.Ints(m.playerShenfaObject.UnrealList)
	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	m.ShenFaOtherGet(int32(shenfaId))
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaUnrealActivate, m.p, shenfaId)
	return
}

//是否能幻化
func (m *PlayerShenfaDataManager) IsCanUnreal(advancedId int) bool {
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(advancedId))
	if shenfaTemplate == nil {
		return false
	}
	curAdvancedId := m.playerShenfaObject.AdvanceId
	//食用幻化丹数量
	curUrealDanLevel := m.playerShenfaObject.UnrealLevel
	//幻化条件
	for condType, cond := range shenfaTemplate.GetMagicParamXUMap() {
		scondType := shenfatypes.ShenfaUCondType(condType)
		switch scondType {
		//身法阶别
		case shenfatypes.ShenfaUCondTypeX:
			if curAdvancedId < int(cond) {
				return false
			}
		//食用幻化丹数量条件
		case shenfatypes.ShenfaUCondTypeU:
			if int32(curUrealDanLevel) < cond {
				return false
			}
		default:
			break
		}
	}
	return true
}

//是否已幻化
func (m *PlayerShenfaDataManager) IsUnrealed(advancedId int) bool {
	uList := m.playerShenfaObject.UnrealList
	for _, v := range uList {
		if v == advancedId {
			return true
		}
	}
	return false
}

//幻化
func (m *PlayerShenfaDataManager) Unreal(advancedId int) (flag bool) {
	if !m.IsUnrealed(advancedId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.ShenfaId = int32(advancedId)
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	flag = true

	gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)
	return
}

//卸下
func (m *PlayerShenfaDataManager) Unload() {
	m.playerShenfaObject.ShenfaId = 0
	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)

}

//身法战斗力
func (m *PlayerShenfaDataManager) ShenfaPower(power int64) (err error) {
	if power <= 0 {
		return
	}
	curPower := m.playerShenfaObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.Power = power
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaPowerChanged, m.p, power)
	return
}

//非进阶身法激活
func (m *PlayerShenfaDataManager) ShenfaOtherGet(shenfaId int32) (err error) {
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenfaId))
	if shenfaTemplate == nil {
		return
	}
	typ := shenfaTemplate.GetTyp()
	if typ == shenfatypes.ShenfaTypeAdvanced {
		return
	}
	m.newPlayerShenfaOtherObject(typ, shenfaId)
	return
}

//展示隐藏坐骑
func (m *PlayerShenfaDataManager) Hidden(hiddenFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		m.playerShenfaObject.Hidden = 1
	} else {
		m.playerShenfaObject.Hidden = 0
	}
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()

	gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)
	if !hiddenFlag {
		gameevent.Emit(shenfaeventtypes.EventTypeShenfaUse, m.p, nil)
	}

}

//是否已拥有该身法皮肤
func (m *PlayerShenfaDataManager) IfShenFaSkinExist(shenFaId int32) (*PlayerShenfaOtherObject, bool) {
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenFaId))
	if shenfaTemplate == nil {
		return nil, false
	}
	typ := shenfaTemplate.GetTyp()
	playerOtherMap, exist := m.playerOtherMap[typ]
	if !exist {
		return nil, false
	}
	shenfaOtherObj, exist := playerOtherMap[shenFaId]
	if !exist {
		return nil, false
	}
	return shenfaOtherObj, true
}

//是否能升星
func (m *PlayerShenfaDataManager) IfCanUpStar(shenFaId int32) (*PlayerShenfaOtherObject, bool) {
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfa(int(shenFaId))
	if shenfaTemplate == nil {
		return nil, false
	}

	shenfaOtherObj, flag := m.IfShenFaSkinExist(shenFaId)
	if !flag {
		return nil, false
	}

	if shenfaTemplate.ShenfaUpstarBeginId == 0 {
		return nil, false
	}

	level := shenfaOtherObj.Level
	if level <= 0 {
		return shenfaOtherObj, true
	}
	nextTo := shenfaTemplate.GetShenFaUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return shenfaOtherObj, true
	}
	return nil, false
}

//身法皮肤升星
func (m *PlayerShenfaDataManager) Upstar(shenfaId int32, pro int32, sucess bool) bool {
	obj, flag := m.IfCanUpStar(shenfaId)
	if !flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.Level += 1
		obj.UpNum = 0
		obj.UpPro = pro
	} else {
		obj.UpNum += 1
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return true
}

func (m *PlayerShenfaDataManager) IfFullAdvanced() (flag bool) {
	if m.playerShenfaObject.AdvanceId == 0 {
		return
	}
	shenfaTemplate := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(m.playerShenfaObject.AdvanceId))
	if shenfaTemplate == nil {
		return
	}
	if shenfaTemplate.NextId == 0 {
		return true
	}
	return
}

//仅gm使用
func (m *PlayerShenfaDataManager) GmSetShenfaAdvanced(advancedId int) {
	now := global.GetGame().GetTimeService().Now()
	m.playerShenfaObject.AdvanceId = advancedId
	m.playerShenfaObject.TimesNum = int32(0)
	m.playerShenfaObject.Bless = int32(0)
	m.playerShenfaObject.BlessTime = int64(0)
	m.playerShenfaObject.ShenfaId = int32(0)
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()

	gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvanced, m.p, int32(m.playerShenfaObject.AdvanceId))

}

//仅gm使用 身法幻化
func (m *PlayerShenfaDataManager) GmSetShenFaUnreal(shenFaId int) {
	now := global.GetGame().GetTimeService().Now()
	if !m.IsUnrealed(shenFaId) {
		m.playerShenfaObject.UnrealList = append(m.playerShenfaObject.UnrealList, shenFaId)
		sort.Ints(m.playerShenfaObject.UnrealList)
		m.playerShenfaObject.UpdateTime = now
		m.playerShenfaObject.SetModified()
	}

	m.playerShenfaObject.ShenfaId = int32(shenFaId)
	m.playerShenfaObject.UpdateTime = now
	m.playerShenfaObject.SetModified()
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaChanged, m.p, nil)
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaUnrealActivate, m.p, shenFaId)
}

func (m *PlayerShenfaDataManager) ToShenfaInfo() *shenfatypes.ShenfaInfo {
	info := &shenfatypes.ShenfaInfo{
		AdvanceId:   int32(m.playerShenfaObject.AdvanceId),
		ShenfaId:    m.playerShenfaObject.ShenfaId,
		UnrealLevel: m.playerShenfaObject.UnrealLevel,
		UnrealPro:   m.playerShenfaObject.UnrealPro,
	}
	for _, typM := range m.playerOtherMap {
		for _, otherObj := range typM {
			skinInfo := &shenfatypes.ShenfaSkinInfo{
				ShenfaId: otherObj.ShenFaId,
				Level:    otherObj.Level,
				UpPro:    otherObj.UpPro,
			}
			info.SkinList = append(info.SkinList, skinInfo)
		}
	}
	return info
}

func CreatePlayerShenfaDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerShenfaDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerShenfaDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerShenfaDataManager))
}
