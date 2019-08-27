package player

import (
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingyu/dao"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	lingyutypes "fgame/fgame/game/lingyu/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"
)

//玩家领域管理器
type PlayerLingyuDataManager struct {
	p player.Player
	//玩家领域对象
	playerLingyuObject *PlayerLingyuObject
	//玩家非进阶领域对象
	playerOtherMap map[lingyutypes.LingyuType]map[int32]*PlayerLingyuOtherObject
}

func (m *PlayerLingyuDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerLingyuDataManager) Load() (err error) {
	m.playerOtherMap = make(map[lingyutypes.LingyuType]map[int32]*PlayerLingyuOtherObject)
	//加载玩家领域信息
	lingyuEntity, err := dao.GetLingyuDao().GetLingyuEntity(m.p.GetId())
	if err != nil {
		return
	}
	if lingyuEntity == nil {
		m.initPlayerLingyuObject()
	} else {
		m.playerLingyuObject = NewPlayerLingyuObject(m.p)
		m.playerLingyuObject.FromEntity(lingyuEntity)
	}

	//加载玩家非进阶领域信息
	lingyuOtherList, err := dao.GetLingyuDao().GetLingyuOtherList(m.p.GetId())
	if err != nil {
		return
	}

	//非进阶领域信息
	for _, lingyuOther := range lingyuOtherList {
		pwo := NewPlayerLingyuOtherObject(m.p)
		pwo.FromEntity(lingyuOther)

		typ := lingyutypes.LingyuType(lingyuOther.Typ)
		playerOtherMap, exist := m.playerOtherMap[typ]
		if !exist {
			playerOtherMap = make(map[int32]*PlayerLingyuOtherObject)
			m.playerOtherMap[typ] = playerOtherMap
		}
		playerOtherMap[pwo.LingYuId] = pwo
	}

	return nil
}

//第一次初始化
func (m *PlayerLingyuDataManager) initPlayerLingyuObject() {
	pwo := NewPlayerLingyuObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	advanceId := playerCreateTemplate.Field
	//生成id
	pwo.AdvanceId = int(advanceId)
	pwo.LingyuId = int32(0)
	pwo.UnrealLevel = int32(0)
	pwo.UnrealNum = int32(0)
	pwo.UnrealPro = int32(0)
	pwo.UnrealList = make([]int, 0, 8)
	pwo.TimesNum = int32(0)
	pwo.Bless = int32(0)
	pwo.BlessTime = int64(0)
	pwo.Power = int64(0)
	pwo.Hidden = 0
	pwo.CreateTime = now
	m.playerLingyuObject = pwo
	pwo.SetModified()
}

//增加非进阶领域
func (m *PlayerLingyuDataManager) newPlayerLingyuOtherObject(typ lingyutypes.LingyuType, lingyuId int32) (err error) {

	playerOtherMap, exist := m.playerOtherMap[typ]
	if !exist {
		playerOtherMap = make(map[int32]*PlayerLingyuOtherObject)
		m.playerOtherMap[typ] = playerOtherMap
	}
	_, exist = playerOtherMap[lingyuId]
	if exist {
		return
	}

	pwo := NewPlayerLingyuOtherObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	//生成id
	pwo.Id = id
	pwo.Typ = typ
	pwo.LingYuId = lingyuId
	pwo.Level = 0
	pwo.UpNum = 0
	pwo.UpPro = 0
	pwo.CreateTime = now
	playerOtherMap[lingyuId] = pwo
	pwo.SetModified()
	return
}

//加载后
func (m *PlayerLingyuDataManager) AfterLoad() (err error) {
	m.refreshBless()

	return nil
}

func (m *PlayerLingyuDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(m.playerLingyuObject.AdvanceId)
	nextNumber := number + 1
	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(nextNumber)
	if lingYuTemplate == nil {
		return
	}
	if !lingYuTemplate.GetIsClear() {
		return
	}
	lastTime := m.playerLingyuObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			m.playerLingyuObject.Bless = 0
			m.playerLingyuObject.BlessTime = 0
			m.playerLingyuObject.TimesNum = 0
			m.playerLingyuObject.SetModified()
		}
	}
}

//领域信息对象
func (m *PlayerLingyuDataManager) GetLingyuInfo() *PlayerLingyuObject {
	m.refreshBless()

	return m.playerLingyuObject
}

func (m *PlayerLingyuDataManager) GetLingyuAdvanced() int32 {
	return int32(m.playerLingyuObject.AdvanceId)
}

//获取玩家非进阶领域对象
func (m *PlayerLingyuDataManager) GetLingyuOtherMap() map[lingyutypes.LingyuType]map[int32]*PlayerLingyuOtherObject {
	return m.playerOtherMap
}

//能吃几个幻化丹
func (m *PlayerLingyuDataManager) CanEatUnrealDanNum() int32 {
	advancedId := m.playerLingyuObject.AdvanceId
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(advancedId)
	shidanLimit := lingyuTemplate.ShidanLimit
	currentNum := m.playerLingyuObject.UnrealNum
	if currentNum >= shidanLimit {
		return 0
	}
	totalNum := shidanLimit - currentNum
	return totalNum
}

// //食幻化丹
// func (m *PlayerLingyuDataManager) EatUnrealDan(pro int32, sucess bool) (err error) {
// 	if pro < 0 {
// 		return nil
// 	}

// 	if sucess {
// 		m.playerLingyuObject.UnrealLevel += 1
// 		m.playerLingyuObject.UnrealNum = 0
// 		m.playerLingyuObject.UnrealPro = pro
// 	} else {
// 		m.playerLingyuObject.UnrealNum += 1
// 		m.playerLingyuObject.UnrealPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	m.playerLingyuObject.UpdateTime = now
// 	m.playerLingyuObject.SetModified()
// 	return
// }

//食幻化丹
func (m *PlayerLingyuDataManager) EatUnrealDan(level int32) (err error) {
	if m.playerLingyuObject.UnrealLevel == level || level <= 0 {
		return
	}
	hunaHuaTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuHuanHuaTemplate(level)
	if hunaHuaTemplate == nil {
		return
	}
	m.playerLingyuObject.UnrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	return
}

//心跳
func (m *PlayerLingyuDataManager) Heartbeat() {

}

func (m *PlayerLingyuDataManager) IsWingHidden() bool {
	return m.playerLingyuObject.Hidden == 1
}

//进阶
func (m *PlayerLingyuDataManager) LingyuAdvanced(pro, addTimes int32, sucess bool) (err error) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(m.playerLingyuObject.AdvanceId + 1)
		if lingyuTemplate == nil {
			return
		}
		m.playerLingyuObject.AdvanceId += 1
		m.playerLingyuObject.TimesNum = 0
		m.playerLingyuObject.Bless = 0
		m.playerLingyuObject.BlessTime = 0
		m.playerLingyuObject.LingyuId = int32(0)

		gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)
		gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvanced, m.p, int32(m.playerLingyuObject.AdvanceId))
	} else {
		m.playerLingyuObject.TimesNum += addTimes
		if m.playerLingyuObject.Bless == 0 {
			m.playerLingyuObject.BlessTime = now
		}
		m.playerLingyuObject.Bless += pro
	}
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	return
}

//直升券进阶
func (m *PlayerLingyuDataManager) LingyuAdvancedTicket(addAdvancedNum int32) (err error) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := m.playerLingyuObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(nextAdvancedId)
		if lingyuTemplate == nil {
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
	m.playerLingyuObject.AdvanceId += canAddNum
	m.playerLingyuObject.TimesNum = 0
	m.playerLingyuObject.Bless = 0
	m.playerLingyuObject.BlessTime = 0
	m.playerLingyuObject.LingyuId = int32(0)
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)
	gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvanced, m.p, int32(m.playerLingyuObject.AdvanceId))
	return
}

//非进阶领域激活
func (m *PlayerLingyuDataManager) LingYuOtherGet(lingYuId int32) {
	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuId))
	if lingYuTemplate == nil {
		return
	}
	typ := lingYuTemplate.GetTyp()
	if typ == lingyutypes.LingyuTypeAdvanced {
		return
	}
	m.newPlayerLingyuOtherObject(typ, lingYuId)
	return
}

//设置幻化信息
func (m *PlayerLingyuDataManager) AddUnrealInfo(lingyuId int) (err error) {
	if lingyuId <= 0 {
		return
	}
	m.playerLingyuObject.UnrealList = append(m.playerLingyuObject.UnrealList, lingyuId)
	sort.Ints(m.playerLingyuObject.UnrealList)
	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	m.LingYuOtherGet(int32(lingyuId))
	gameevent.Emit(lingyueventtypes.EventTypeLingyuUnrealActivate, m.p, lingyuId)
	return
}

//是否能幻化
func (m *PlayerLingyuDataManager) IsCanUnreal(advancedId int) bool {
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(advancedId))
	if lingyuTemplate == nil {
		return false
	}
	curAdvancedId := m.playerLingyuObject.AdvanceId
	//食用幻化丹数量
	curUrealDanLevel := m.playerLingyuObject.UnrealLevel
	//幻化条件
	for condType, cond := range lingyuTemplate.GetMagicParamXUMap() {
		scondType := lingyutypes.LingyuUCondType(condType)
		switch scondType {
		//领域阶别
		case lingyutypes.LingyuUCondTypeX:
			if curAdvancedId < int(cond) {
				return false
			}
		//食用幻化丹数量条件
		case lingyutypes.LingyuUCondTypeU:
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
func (m *PlayerLingyuDataManager) IsUnrealed(advancedId int) bool {
	uList := m.playerLingyuObject.UnrealList
	for _, v := range uList {
		if v == advancedId {
			return true
		}
	}
	return false
}

//幻化
func (m *PlayerLingyuDataManager) Unreal(advancedId int) (flag bool) {
	if !m.IsUnrealed(advancedId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.LingyuId = int32(advancedId)
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	flag = true

	gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)
	return
}

//卸下
func (m *PlayerLingyuDataManager) Unload() {
	m.playerLingyuObject.LingyuId = 0
	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()

	gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)

}

//领域战斗力
func (m *PlayerLingyuDataManager) LingyuPower(power int64) (err error) {
	if power <= 0 {
		return
	}
	curPower := m.playerLingyuObject.Power
	if curPower == power {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.Power = power
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	gameevent.Emit(lingyueventtypes.EventTypeLingyuPowerChanged, m.p, power)
	return
}

//非进阶领域激活
func (m *PlayerLingyuDataManager) LingyuOtherGet(lingyuId int32) (err error) {
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingyuId))
	if lingyuTemplate == nil {
		return
	}
	typ := lingyuTemplate.GetTyp()
	if typ == lingyutypes.LingyuTypeAdvanced {
		return
	}
	m.newPlayerLingyuOtherObject(typ, lingyuId)
	return
}

//展示隐藏坐骑
func (m *PlayerLingyuDataManager) Hidden(hiddenFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		m.playerLingyuObject.Hidden = 1
	} else {
		m.playerLingyuObject.Hidden = 0
	}
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()

	gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)
	if !hiddenFlag {
		gameevent.Emit(lingyueventtypes.EventTypeLingyuUse, m.p, nil)
	}
}

//是否已拥有该领域皮肤
func (m *PlayerLingyuDataManager) IfLingYuSkinExist(lingYuId int32) (*PlayerLingyuOtherObject, bool) {
	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuId))
	if lingYuTemplate == nil {
		return nil, false
	}
	typ := lingYuTemplate.GetTyp()
	playerOtherMap, exist := m.playerOtherMap[typ]
	if !exist {
		return nil, false
	}
	lingYuOtherObj, exist := playerOtherMap[lingYuId]
	if !exist {
		return nil, false
	}
	return lingYuOtherObj, true
}

//是否能升星
func (m *PlayerLingyuDataManager) IfCanUpStar(lingYuId int32) (*PlayerLingyuOtherObject, bool) {
	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyu(int(lingYuId))
	if lingYuTemplate == nil {
		return nil, false
	}

	lingYuOtherObj, flag := m.IfLingYuSkinExist(lingYuId)
	if !flag {
		return nil, false
	}

	if lingYuTemplate.FieldUpstarBeginId == 0 {
		return nil, false
	}

	level := lingYuOtherObj.Level
	if level <= 0 {
		return lingYuOtherObj, true
	}
	nextTo := lingYuTemplate.GetLingYuUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return lingYuOtherObj, true
	}
	return nil, false
}

//领域皮肤升星
func (m *PlayerLingyuDataManager) Upstar(lingYuId int32, pro int32, sucess bool) bool {
	obj, flag := m.IfCanUpStar(lingYuId)
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

//领域充值数
func (m *PlayerLingyuDataManager) AddChargeNum(chargeNum int32) {
	if chargeNum <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.ChargeVal += int64(chargeNum)
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	return
}

//仅gm使用
func (m *PlayerLingyuDataManager) GmSetLingyuAdvanced(advancedId int) {
	m.playerLingyuObject.AdvanceId = advancedId
	m.playerLingyuObject.TimesNum = int32(0)
	m.playerLingyuObject.Bless = int32(0)
	m.playerLingyuObject.BlessTime = int64(0)
	m.playerLingyuObject.LingyuId = int32(0)
	now := global.GetGame().GetTimeService().Now()
	m.playerLingyuObject.UpdateTime = now

	gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)
	gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvanced, m.p, int32(m.playerLingyuObject.AdvanceId))

}

//仅gm使用 领域幻化
func (m *PlayerLingyuDataManager) GmSetLingYuUnreal(lingYuId int) {
	now := global.GetGame().GetTimeService().Now()
	if !m.IsUnrealed(lingYuId) {
		m.playerLingyuObject.UnrealList = append(m.playerLingyuObject.UnrealList, lingYuId)
		sort.Ints(m.playerLingyuObject.UnrealList)
		m.playerLingyuObject.UpdateTime = now
		m.playerLingyuObject.SetModified()
	}

	m.playerLingyuObject.LingyuId = int32(lingYuId)
	m.playerLingyuObject.UpdateTime = now
	m.playerLingyuObject.SetModified()
	gameevent.Emit(lingyueventtypes.EventTypeLingyuChanged, m.p, nil)
	gameevent.Emit(lingyueventtypes.EventTypeLingyuUnrealActivate, m.p, lingYuId)
}

func (m *PlayerLingyuDataManager) ToLingyuInfo() *lingyutypes.LingyuInfo {
	info := &lingyutypes.LingyuInfo{
		AdvanceId:   int32(m.playerLingyuObject.AdvanceId),
		LingyuId:    m.playerLingyuObject.LingyuId,
		UnrealLevel: m.playerLingyuObject.UnrealLevel,
		UnrealPro:   m.playerLingyuObject.UnrealPro,
	}
	for _, typM := range m.playerOtherMap {
		for _, otherObj := range typM {
			skinInfo := &lingyutypes.LingyuSkinInfo{
				LingyuId: otherObj.LingYuId,
				Level:    otherObj.Level,
				UpPro:    otherObj.UpPro,
			}
			info.SkinList = append(info.SkinList, skinInfo)
		}
	}
	return info
}

func (m *PlayerLingyuDataManager) GetLingYuId() int32 {
	if m.playerLingyuObject.Hidden != 0 {
		return 0
	}
	if m.playerLingyuObject.LingyuId != 0 {
		return m.playerLingyuObject.LingyuId
	}
	lingyuTemplate := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(int32(m.playerLingyuObject.AdvanceId))
	if lingyuTemplate == nil {
		return 0
	}
	return int32(lingyuTemplate.TemplateId())
}

func CreatePlayerLingyuDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerLingyuDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerLingyuDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerLingyuDataManager))
}
