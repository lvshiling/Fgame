package player

import (
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"

	gameevent "fgame/fgame/game/event"
	fabaocommon "fgame/fgame/game/fabao/common"
	"fgame/fgame/game/fabao/dao"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaotemplate "fgame/fgame/game/fabao/template"
	fabaotypes "fgame/fgame/game/fabao/types"
)

//玩家法宝管理器
type PlayerFaBaoDataManager struct {
	p player.Player
	//玩家法宝对象
	playerFaBaoObject *PlayerFaBaoObject
	//玩家非进阶法宝对象
	playerOtherMap map[fabaotypes.FaBaoType]map[int32]*PlayerFaBaoOtherObject
}

func (pwdm *PlayerFaBaoDataManager) Player() player.Player {
	return pwdm.p
}

//加载
func (pwdm *PlayerFaBaoDataManager) Load() (err error) {
	pwdm.playerOtherMap = make(map[fabaotypes.FaBaoType]map[int32]*PlayerFaBaoOtherObject)

	//加载玩家法宝信息
	faBaoEntity, err := dao.GetFaBaoDao().GetFaBaoEntity(pwdm.p.GetId())
	if err != nil {
		return
	}
	if faBaoEntity == nil {
		pwdm.initPlayerFaBaoObject()
	} else {
		pwdm.playerFaBaoObject = NewPlayerFaBaoObject(pwdm.p)
		pwdm.playerFaBaoObject.FromEntity(faBaoEntity)
	}

	//加载玩家非进阶法宝信息
	faBaoOtherList, err := dao.GetFaBaoDao().GetFaBaoOtherList(pwdm.p.GetId())
	if err != nil {
		return
	}
	//非进阶法宝信息
	for _, faBaoOther := range faBaoOtherList {
		pwo := NewPlayerFaBaoOtherObject(pwdm.p)
		pwo.FromEntity(faBaoOther)

		typ := fabaotypes.FaBaoType(faBaoOther.Typ)
		playerOtherMap, exist := pwdm.playerOtherMap[typ]
		if !exist {
			playerOtherMap = make(map[int32]*PlayerFaBaoOtherObject)
			pwdm.playerOtherMap[typ] = playerOtherMap
		}
		playerOtherMap[pwo.faBaoId] = pwo
	}

	return nil
}

//第一次初始化
func (pwdm *PlayerFaBaoDataManager) initPlayerFaBaoObject() {
	pwo := NewPlayerFaBaoObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(pwdm.p.GetRole(), pwdm.p.GetSex())
	advanceId := playerCreateTemplate.FaBao
	//生成id
	pwo.playerId = pwdm.p.GetId()
	pwo.advanceId = int(advanceId)
	pwo.faBaoId = int32(0)
	pwo.unrealLevel = int32(0)
	pwo.unrealNum = int32(0)
	pwo.unrealPro = int32(0)
	pwo.unrealList = make([]int, 0, 8)
	pwo.timesNum = int32(0)
	pwo.bless = int32(0)
	pwo.blessTime = int64(0)
	pwo.tongLingLevel = 0
	pwo.tongLingNum = 0
	pwo.tongLingPro = 0
	pwo.hidden = 0
	pwo.power = int64(0)
	pwo.createTime = now
	pwdm.playerFaBaoObject = pwo
	pwo.SetModified()
}

//增加非进阶法宝
func (pwdm *PlayerFaBaoDataManager) newPlayerFaBaoOtherObject(typ fabaotypes.FaBaoType, faBaoId int32) (err error) {

	playerOtherMap, exist := pwdm.playerOtherMap[typ]
	if !exist {
		playerOtherMap = make(map[int32]*PlayerFaBaoOtherObject)
		pwdm.playerOtherMap[typ] = playerOtherMap
	}
	_, exist = playerOtherMap[faBaoId]
	if exist {
		return
	}

	pwo := NewPlayerFaBaoOtherObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.id = id
	//生成id
	pwo.playerId = pwdm.p.GetId()
	pwo.typ = typ
	pwo.faBaoId = faBaoId
	pwo.level = 0
	pwo.upNum = 0
	pwo.upPro = 0
	pwo.createTime = now
	playerOtherMap[faBaoId] = pwo
	pwo.SetModified()
	return
}

func (pwdm *PlayerFaBaoDataManager) refreshBless() (err error) {
	now := global.GetGame().GetTimeService().Now()
	number := int32(pwdm.playerFaBaoObject.advanceId)
	nextNumber := number + 1
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(nextNumber)
	if faBaoTemplate == nil {
		return
	}
	if !faBaoTemplate.GetIsClear() {
		return
	}
	lastTime := pwdm.playerFaBaoObject.blessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			pwdm.playerFaBaoObject.bless = 0
			pwdm.playerFaBaoObject.blessTime = 0
			pwdm.playerFaBaoObject.timesNum = 0
			pwdm.playerFaBaoObject.SetModified()
		}
	}
	return
}

//加载后
func (pwdm *PlayerFaBaoDataManager) AfterLoad() (err error) {
	err = pwdm.refreshBless()
	return nil
}

//法宝信息对象
func (pwdm *PlayerFaBaoDataManager) GetFaBaoInfo() *PlayerFaBaoObject {
	pwdm.refreshBless()
	return pwdm.playerFaBaoObject
}

func (pwdm *PlayerFaBaoDataManager) GetFaBaoAdvancedId() int32 {
	return int32(pwdm.playerFaBaoObject.advanceId)
}

func (m *PlayerFaBaoDataManager) GetFaBaoId() int32 {
	if m.playerFaBaoObject.hidden != 0 {
		return 0
	}

	if m.playerFaBaoObject.faBaoId != 0 {
		return m.playerFaBaoObject.faBaoId
	}
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(m.playerFaBaoObject.advanceId))
	if faBaoTemplate == nil {
		return 0
	}
	return int32(faBaoTemplate.TemplateId())
}

//获取玩家非进阶法宝对象
func (pwdm *PlayerFaBaoDataManager) GetFaBaoOtherMap() map[fabaotypes.FaBaoType]map[int32]*PlayerFaBaoOtherObject {
	return pwdm.playerOtherMap
}

func (pwdm *PlayerFaBaoDataManager) EatUnrealDan(level int32) {
	if pwdm.playerFaBaoObject.unrealLevel == level || level <= 0 {
		return
	}
	huanHuaTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoHuanHuaTemplate(level)
	if huanHuaTemplate == nil {
		return
	}
	pwdm.playerFaBaoObject.unrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	return
}

//心跳
func (pwdm *PlayerFaBaoDataManager) Heartbeat() {

}

func (pwdm *PlayerFaBaoDataManager) IsFaBaoHidden() bool {
	return pwdm.playerFaBaoObject.hidden == 1
}

//进阶
func (pwdm *PlayerFaBaoDataManager) FaBaoAdvanced(pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextAdvancedId := pwdm.playerFaBaoObject.advanceId + 1
		faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(nextAdvancedId))
		if faBaoTemplate == nil {
			return
		}
		pwdm.playerFaBaoObject.advanceId += 1
		pwdm.playerFaBaoObject.timesNum = 0
		pwdm.playerFaBaoObject.bless = 0
		pwdm.playerFaBaoObject.blessTime = 0
		pwdm.playerFaBaoObject.faBaoId = int32(0)
		gameevent.Emit(fabaoeventtypes.EventTypeFaBaoChanged, pwdm.p, nil)
		gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvanced, pwdm.p, int32(pwdm.playerFaBaoObject.advanceId))
	} else {
		pwdm.playerFaBaoObject.timesNum += addTimes
		if pwdm.playerFaBaoObject.bless == 0 {
			pwdm.playerFaBaoObject.blessTime = now
		}
		pwdm.playerFaBaoObject.bless += pro
	}
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	return
}

//直升券进阶
func (pwdm *PlayerFaBaoDataManager) FaBaoAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := pwdm.playerFaBaoObject.advanceId + 1
	for addAdvancedNum > 0 {
		fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(nextAdvancedId))
		if fabaoTemplate == nil {
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
	pwdm.playerFaBaoObject.advanceId += canAddNum
	pwdm.playerFaBaoObject.timesNum = 0
	pwdm.playerFaBaoObject.bless = 0
	pwdm.playerFaBaoObject.blessTime = 0
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvanced, pwdm.p, int32(pwdm.playerFaBaoObject.advanceId))
	return
}

//展示隐藏法宝
func (pwdm *PlayerFaBaoDataManager) Hidden(hiddenFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		pwdm.playerFaBaoObject.hidden = 1
	} else {
		pwdm.playerFaBaoObject.hidden = 0
	}
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoChanged, pwdm.p, nil)
	if !hiddenFlag {
		gameevent.Emit(fabaoeventtypes.EventTypeFaBaoUse, pwdm.p, nil)
	}
	return
}

//设置幻化信息
func (pwdm *PlayerFaBaoDataManager) AddUnrealInfo(faBaoId int) {
	if faBaoId <= 0 {
		return
	}
	pwdm.playerFaBaoObject.unrealList = append(pwdm.playerFaBaoObject.unrealList, faBaoId)
	sort.Ints(pwdm.playerFaBaoObject.unrealList)
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	pwdm.FaBaoOtherGet(int32(faBaoId))

	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoUnrealActivate, pwdm.p, faBaoId)
	return
}

//是否能幻化
func (pwdm *PlayerFaBaoDataManager) IsCanUnreal(faBaoId int) bool {
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(faBaoId)
	if faBaoTemplate == nil {
		return false
	}
	curAdvancedId := pwdm.playerFaBaoObject.advanceId
	//食用幻化丹等级
	curUnrealLevel := pwdm.playerFaBaoObject.unrealLevel
	//幻化条件
	for condType, cond := range faBaoTemplate.GetMagicParamXUMap() {
		scondType := fabaotypes.FaBaoUCondType(condType)
		switch scondType {
		//法宝阶别
		case fabaotypes.FaBaoUCondTypeX:
			if curAdvancedId < int(cond) {
				return false
			}
		//食用幻化丹等级条件
		case fabaotypes.FaBaoUCondTypeU:
			if int32(curUnrealLevel) < cond {
				return false
			}
		//关联其他功能等级
		case fabaotypes.FaBaoUCondTypeW:
			return false
		default:
			break
		}
	}
	return true
}

//是否已幻化
func (pwdm *PlayerFaBaoDataManager) IsUnrealed(faBaoId int) bool {
	uList := pwdm.playerFaBaoObject.unrealList
	for _, v := range uList {
		if v == faBaoId {
			//容错处理
			pwdm.FaBaoOtherGet(int32(faBaoId))
			return true
		}
	}
	return false
}

//幻化
func (pwdm *PlayerFaBaoDataManager) Unreal(faBaoId int) (flag bool) {
	if !pwdm.IsUnrealed(faBaoId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.faBaoId = int32(faBaoId)
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoChanged, pwdm.p, nil)
	flag = true
	return
}

//卸下
func (pwdm *PlayerFaBaoDataManager) Unload() {
	pwdm.playerFaBaoObject.faBaoId = 0
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoChanged, pwdm.p, nil)
	return
}

//法宝战斗力
func (pwdm *PlayerFaBaoDataManager) FaBaoPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := pwdm.playerFaBaoObject.power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.power = power
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoPowerChanged, pwdm.p, power)
	return
}

//非进阶法宝激活
func (pwdm *PlayerFaBaoDataManager) FaBaoOtherGet(faBaoId int32) {
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
	if faBaoTemplate == nil {
		return
	}
	typ := faBaoTemplate.GetTyp()
	if typ == fabaotypes.FaBaoTypeAdvanced {
		return
	}
	pwdm.newPlayerFaBaoOtherObject(typ, faBaoId)
	return
}

//是否已拥有该法宝皮肤
func (pwdm *PlayerFaBaoDataManager) IfFaBaoSkinExist(faBaoId int32) (*PlayerFaBaoOtherObject, bool) {
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
	if faBaoTemplate == nil {
		return nil, false
	}
	typ := faBaoTemplate.GetTyp()
	playerOtherMap, exist := pwdm.playerOtherMap[typ]
	if !exist {
		return nil, false
	}
	faBaoOtherObj, exist := playerOtherMap[faBaoId]
	if !exist {
		return nil, false
	}
	return faBaoOtherObj, true
}

//是否能升星
func (pwdm *PlayerFaBaoDataManager) IfCanUpStar(faBaoId int32) (*PlayerFaBaoOtherObject, bool) {
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
	if faBaoTemplate == nil {
		return nil, false
	}

	faBaoOtherObj, flag := pwdm.IfFaBaoSkinExist(faBaoId)
	if !flag {
		return nil, false
	}

	if faBaoTemplate.FaBaoUpstarBeginId == 0 {
		return nil, false
	}

	level := faBaoOtherObj.level
	if level <= 0 {
		return faBaoOtherObj, true
	}
	nextTo := faBaoTemplate.GetFaBaoUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return faBaoOtherObj, true
	}
	return nil, false
}

//法宝皮肤升星
func (pwdm *PlayerFaBaoDataManager) Upstar(faBaoId int32, pro int32, sucess bool) bool {
	obj, flag := pwdm.IfCanUpStar(faBaoId)
	if !flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		to := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
		if to == nil {
			return false
		}
		faBaoUpstarTemplate := to.GetFaBaoUpstarByLevel(obj.level + 1)
		if faBaoUpstarTemplate == nil {
			return false
		}
		obj.level += 1
		obj.upNum = 0
		obj.upPro = pro
	} else {
		obj.upNum += 1
		obj.upPro += pro
	}
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (pwdm *PlayerFaBaoDataManager) TongLing(pro int32, sucess bool) (flag bool) {
	tongLingLevel := pwdm.playerFaBaoObject.tongLingLevel
	nextLevel := tongLingLevel + 1
	tongLingTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoTongLingTemplate(nextLevel)
	if tongLingTemplate == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		pwdm.playerFaBaoObject.tongLingLevel += 1
		pwdm.playerFaBaoObject.tongLingNum = 0
		pwdm.playerFaBaoObject.tongLingPro = pro
	} else {
		pwdm.playerFaBaoObject.tongLingNum += 1
		pwdm.playerFaBaoObject.tongLingPro += pro
	}
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	return true
}

func (pwdm *PlayerFaBaoDataManager) ToFaBaoInfo() *fabaocommon.FaBaoInfo {
	info := &fabaocommon.FaBaoInfo{
		AdvanceId:     int32(pwdm.playerFaBaoObject.advanceId),
		FaBaoId:       pwdm.playerFaBaoObject.faBaoId,
		UnrealLevel:   pwdm.playerFaBaoObject.unrealLevel,
		UnrealPro:     pwdm.playerFaBaoObject.unrealPro,
		TongLingLevel: pwdm.playerFaBaoObject.tongLingLevel,
		TongLingPro:   pwdm.playerFaBaoObject.tongLingPro,
	}
	for _, typM := range pwdm.playerOtherMap {
		for _, otherObj := range typM {
			skinInfo := &fabaocommon.FaBaoSkinInfo{
				FaBaoId: otherObj.faBaoId,
				Level:   otherObj.level,
				UpPro:   otherObj.upPro,
			}
			info.SkinList = append(info.SkinList, skinInfo)
		}
	}
	return info
}

func (pwdm *PlayerFaBaoDataManager) IfFullAdvanced() (flag bool) {
	if pwdm.playerFaBaoObject.advanceId == 0 {
		return
	}
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(pwdm.playerFaBaoObject.advanceId))
	if faBaoTemplate == nil {
		return
	}
	if faBaoTemplate.NextId == 0 {
		return true
	}
	return
}

//仅gm使用 法宝进阶
func (pwdm *PlayerFaBaoDataManager) GmSetFaBaoAdvanced(advancedId int) {
	pwdm.playerFaBaoObject.advanceId = advancedId
	pwdm.playerFaBaoObject.timesNum = int32(0)
	pwdm.playerFaBaoObject.bless = int32(0)
	pwdm.playerFaBaoObject.blessTime = int64(0)
	pwdm.playerFaBaoObject.faBaoId = int32(0)
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoChanged, pwdm.p, nil)
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvanced, pwdm.p, int32(pwdm.playerFaBaoObject.advanceId))
	return
}

//仅gm使用 法宝幻化
func (pwdm *PlayerFaBaoDataManager) GmSetFaBaoUnreal(faBaoId int) {
	now := global.GetGame().GetTimeService().Now()
	if !pwdm.IsUnrealed(faBaoId) {
		pwdm.playerFaBaoObject.unrealList = append(pwdm.playerFaBaoObject.unrealList, faBaoId)
		sort.Ints(pwdm.playerFaBaoObject.unrealList)
		pwdm.playerFaBaoObject.updateTime = now
		pwdm.playerFaBaoObject.SetModified()
	}

	pwdm.playerFaBaoObject.faBaoId = int32(faBaoId)
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoChanged, pwdm.p, nil)
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoUnrealActivate, pwdm.p, faBaoId)
}

//仅gm使用 法宝食幻化丹
func (pwdm *PlayerFaBaoDataManager) GmSetFaBaoUnrealDanLevel(level int32) {
	pwdm.playerFaBaoObject.unrealLevel = level
	pwdm.playerFaBaoObject.unrealNum = 0
	pwdm.playerFaBaoObject.unrealPro = 0

	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
}

//仅gm使用 法宝通灵
func (pwdm *PlayerFaBaoDataManager) GmSetFaBaoTongLing(level int32) {
	pwdm.playerFaBaoObject.tongLingLevel = level
	pwdm.playerFaBaoObject.tongLingNum = 0
	pwdm.playerFaBaoObject.tongLingPro = 0

	now := global.GetGame().GetTimeService().Now()
	pwdm.playerFaBaoObject.updateTime = now
	pwdm.playerFaBaoObject.SetModified()
}

func CreatePlayerFaBaoDataManager(p player.Player) player.PlayerDataManager {
	pwdm := &PlayerFaBaoDataManager{}
	pwdm.p = p
	return pwdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFaBaoDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFaBaoDataManager))
}
