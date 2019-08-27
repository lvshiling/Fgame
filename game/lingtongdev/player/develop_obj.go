package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtongdev/dao"
	"fgame/fgame/game/lingtongdev/entity"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sort"
)

//灵童养成类对象
type PlayerLingTongDevObject struct {
	player        player.Player
	id            int64
	playerId      int64
	classType     lingtongdevtypes.LingTongDevSysType
	advanceId     int
	seqId         int32
	unrealLevel   int32
	unrealNum     int32
	unrealPro     int32
	unrealList    []int
	timesNum      int32
	bless         int32
	blessTime     int64
	culLevel      int32
	culNum        int32
	culPro        int32
	tongLingLevel int32
	tongLingNum   int32
	tongLingPro   int32
	hidden        int32
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerLingTongDevObject(pl player.Player) *PlayerLingTongDevObject {
	pwo := &PlayerLingTongDevObject{
		player: pl,
	}
	return pwo
}

func convertNewPlayerLingTongDevObjectToEntity(pwo *PlayerLingTongDevObject) (*entity.PlayerLingTongDevEntity, error) {
	unrealInfoBytes, err := json.Marshal(pwo.unrealList)
	if err != nil {
		return nil, err
	}
	e := &entity.PlayerLingTongDevEntity{
		Id:            pwo.id,
		PlayerId:      pwo.playerId,
		ClassType:     int32(pwo.classType),
		AdvancedId:    pwo.advanceId,
		SeqId:         pwo.seqId,
		UnrealLevel:   pwo.unrealLevel,
		UnrealNum:     pwo.unrealNum,
		UnrealPro:     pwo.unrealPro,
		UnrealInfo:    string(unrealInfoBytes),
		TimesNum:      pwo.timesNum,
		Bless:         pwo.bless,
		BlessTime:     pwo.blessTime,
		CulLevel:      pwo.culLevel,
		CulNum:        pwo.culNum,
		CulPro:        pwo.culPro,
		TongLingLevel: pwo.tongLingLevel,
		TongLingNum:   pwo.tongLingNum,
		TongLingPro:   pwo.tongLingPro,
		Hidden:        pwo.hidden,
		UpdateTime:    pwo.updateTime,
		CreateTime:    pwo.createTime,
		DeleteTime:    pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongDevObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongDevObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongDevObject) GetClassType() lingtongdevtypes.LingTongDevSysType {
	return pwo.classType
}

func (pwo *PlayerLingTongDevObject) GetAdvancedId() int32 {
	return int32(pwo.advanceId)
}

func (pwo *PlayerLingTongDevObject) GetSeqId() int32 {
	return pwo.seqId
}

func (pwo *PlayerLingTongDevObject) GetUnrealLevel() int32 {
	return pwo.unrealLevel
}

func (pwo *PlayerLingTongDevObject) GetUnrealNum() int32 {
	return pwo.unrealNum
}

func (pwo *PlayerLingTongDevObject) GetUnrealPro() int32 {
	return pwo.unrealPro
}

func (pwo *PlayerLingTongDevObject) GetUnrealList() []int {
	return pwo.unrealList
}

func (pwo *PlayerLingTongDevObject) GetTimesNum() int32 {
	return pwo.timesNum
}

func (pwo *PlayerLingTongDevObject) GetBless() int32 {
	return pwo.bless
}

func (pwo *PlayerLingTongDevObject) GetBlessTime() int64 {
	return pwo.blessTime
}

func (pwo *PlayerLingTongDevObject) GetCulLevel() int32 {
	return pwo.culLevel
}

func (pwo *PlayerLingTongDevObject) GetCulNum() int32 {
	return pwo.culNum
}

func (pwo *PlayerLingTongDevObject) GetCulPro() int32 {
	return pwo.culPro
}

func (pwo *PlayerLingTongDevObject) GetTongLingLevel() int32 {
	return pwo.tongLingLevel
}

func (pwo *PlayerLingTongDevObject) GetTongLingNum() int32 {
	return pwo.tongLingNum
}

func (pwo *PlayerLingTongDevObject) GetTongLingPro() int32 {
	return pwo.tongLingPro
}

func (pwo *PlayerLingTongDevObject) GetHidden() int32 {
	return pwo.hidden
}

func (pwo *PlayerLingTongDevObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerLingTongDevObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongDevObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongDevEntity)

	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}
	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.classType = lingtongdevtypes.LingTongDevSysType(pse.ClassType)
	pwo.advanceId = pse.AdvancedId
	pwo.seqId = pse.SeqId
	pwo.unrealLevel = pse.UnrealLevel
	pwo.unrealNum = pse.UnrealNum
	pwo.unrealPro = pse.UnrealPro
	pwo.unrealList = unrealList
	pwo.timesNum = pse.TimesNum
	pwo.bless = pse.Bless
	pwo.blessTime = pse.BlessTime
	pwo.culLevel = pse.CulLevel
	pwo.culNum = pse.CulNum
	pwo.culPro = pse.CulPro
	pwo.hidden = pse.Hidden
	pwo.tongLingLevel = pse.TongLingLevel
	pwo.tongLingNum = pse.TongLingNum
	pwo.tongLingPro = pse.TongLingPro
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongDevObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtongdev: err[%s]", err.Error()))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

func (m *PlayerLingTongDevDataManager) loadDevelop() (err error) {
	m.playerLingTongDevMap = make(map[types.LingTongDevSysType]*PlayerLingTongDevObject)
	//加载玩家灵童养成信息
	lingTongDevList, err := dao.GetLingTongDevDao().GetLingTongDevList(m.p.GetId())
	if err != nil {
		return
	}
	for _, lingTongDevObj := range lingTongDevList {
		pwo := NewPlayerLingTongDevObject(m.p)
		pwo.FromEntity(lingTongDevObj)

		classType := pwo.GetClassType()
		m.playerLingTongDevMap[classType] = pwo
	}
	return
}

func (m *PlayerLingTongDevDataManager) refreshAllBless() (err error) {
	for classType, _ := range m.playerLingTongDevMap {
		err = m.refreshBless(classType)
		if err != nil {
			return
		}
	}
	return
}

func (m *PlayerLingTongDevDataManager) refreshBless(classType types.LingTongDevSysType) (err error) {
	obj, ok := m.playerLingTongDevMap[classType]
	if !ok {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	number := obj.GetAdvancedId()
	nextNumber := number + 1
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, nextNumber)
	if lingTongDevTemplate == nil {
		return
	}
	if !lingTongDevTemplate.GetIsClear() {
		return
	}
	lastTime := obj.GetBlessTime()
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return err
		}
		if !flag {
			obj.bless = 0
			obj.blessTime = 0
			obj.timesNum = 0
			obj.SetModified()
		}
	}
	return
}

func (m *PlayerLingTongDevDataManager) initPlayerLingTongDevObj(classType types.LingTongDevSysType) (obj *PlayerLingTongDevObject) {
	obj = m.getLingTongInfo(classType)
	if obj != nil {
		return
	}
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj = NewPlayerLingTongDevObject(m.p)
	obj.playerId = m.p.GetId()
	obj.id = id
	obj.classType = classType
	obj.advanceId = 0
	obj.seqId = 0
	obj.unrealList = make([]int, 0, 8)
	obj.unrealLevel = 0
	obj.unrealNum = 0
	obj.unrealPro = 0
	obj.culLevel = 0
	obj.culNum = 0
	obj.culPro = 0
	obj.timesNum = 0
	obj.bless = 0
	obj.blessTime = 0
	obj.createTime = now
	//默认不隐藏
	obj.hidden = 0
	// if classType != types.LingTongDevSysTypeLingBing {
	// 	obj.hidden = 1
	// } else {
	// 	obj.hidden = 0
	// }
	obj.SetModified()
	m.playerLingTongDevMap[classType] = obj
	return obj
}

func (m *PlayerLingTongDevDataManager) getLingTongInfo(classType types.LingTongDevSysType) *PlayerLingTongDevObject {
	obj, ok := m.playerLingTongDevMap[classType]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerLingTongDevDataManager) AdvancedInit(classType types.LingTongDevSysType) (obj *PlayerLingTongDevObject) {
	return m.initPlayerLingTongDevObj(classType)
}

//灵童养成信息对象
func (m *PlayerLingTongDevDataManager) GetLingTongDevInfo(classType types.LingTongDevSysType) *PlayerLingTongDevObject {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return m.AdvancedInit(classType)
	}
	m.refreshBless(classType)
	return obj
}

func (m *PlayerLingTongDevDataManager) GetLingTongDevAdvancedId(classType types.LingTongDevSysType) int32 {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return 0
	}
	return obj.GetAdvancedId()
}

func (m *PlayerLingTongDevDataManager) GetLingTongDevSeqId(classType types.LingTongDevSysType) int32 {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return 0
	}
	switch classType {
	case types.LingTongDevSysTypeLingBing:
		break
	default:
		if obj.GetHidden() != 0 {
			return 0
		}
		break
	}

	if obj.GetSeqId() != 0 {
		return obj.GetSeqId()
	}

	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, obj.GetAdvancedId())
	if lingTongDevTemplate == nil {
		return 0
	}
	return int32(lingTongDevTemplate.TemplateId())
}

func (m *PlayerLingTongDevDataManager) EatUnrealDan(classType types.LingTongDevSysType, level int32) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	if obj.GetUnrealLevel() == level || level <= 0 {
		return
	}
	lingTongDevHuanHuaTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevHuanHuaTemplate(classType, level)
	if lingTongDevHuanHuaTemplate == nil {
		return
	}
	obj.unrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	return
}

func (m *PlayerLingTongDevDataManager) IsLingTongDevHidden(classType types.LingTongDevSysType) bool {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return true
	}
	return obj.GetHidden() == 1
}

//进阶
func (m *PlayerLingTongDevDataManager) LingTongDevAdvanced(classType types.LingTongDevSysType, pro, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextAdvancedId := obj.GetAdvancedId() + 1
		lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, nextAdvancedId)
		if lingTongDevTemplate == nil {
			return
		}
		obj.advanceId += 1
		obj.timesNum = 0
		obj.bless = 0
		obj.blessTime = 0
		obj.seqId = 0
		gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
		gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, m.p, obj)
	} else {
		obj.timesNum += addTimes
		if obj.GetBless() == 0 {
			obj.blessTime = now
		}
		obj.bless += pro
	}
	obj.updateTime = now
	obj.SetModified()
	return
}

//直升券进阶
func (m *PlayerLingTongDevDataManager) LingTongDevAdvancedTicket(classType types.LingTongDevSysType, addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}

	canAddNum := 0
	nextAdvancedId := obj.GetAdvancedId() + 1
	for addAdvancedNum > 0 {
		lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, nextAdvancedId)
		if lingTongDevTemplate == nil {
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
	obj.advanceId += canAddNum
	obj.timesNum = 0
	obj.bless = 0
	obj.blessTime = 0
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, m.p, obj)
	return
}

//展示隐藏灵童养成
func (m *PlayerLingTongDevDataManager) Hidden(classType types.LingTongDevSysType, hiddenFlag bool) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		obj.hidden = 1
	} else {
		obj.hidden = 0
	}
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
	if !hiddenFlag {
		gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevUse, m.p, obj)
	}
	return
}

//设置幻化信息
func (m *PlayerLingTongDevDataManager) AddUnrealInfo(classType types.LingTongDevSysType, seqId int) {
	if seqId <= 0 {
		return
	}
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, seqId)
	if lingTongDevTemplate == nil {
		return
	}
	obj.unrealList = append(obj.unrealList, seqId)
	sort.Ints(obj.unrealList)
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	m.LingTongDevOtherGet(classType, int32(seqId))

	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevUnrealActivate, m.p, obj)
	return
}

//是否能幻化
func (m *PlayerLingTongDevDataManager) IsCanUnreal(classType types.LingTongDevSysType, seqId int) (flag bool) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, seqId)
	if lingTongDevTemplate == nil {
		return
	}
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	curAdvancedId := obj.GetAdvancedId()
	//食用幻化丹等级
	curUnrealLevel := obj.GetUnrealLevel()
	//幻化条件
	for condType, cond := range lingTongDevTemplate.GetMagicParamXUMap() {
		scondType := types.LingTongDevUCondType(condType)
		switch scondType {
		//灵童养成阶别
		case types.LingTongDevUCondTypeX:
			if curAdvancedId < cond {
				return
			}
		//食用幻化丹等级条件
		case types.LingTongDevUCondTypeU:
			if curUnrealLevel < cond {
				return
			}
		default:
			break
		}
	}
	return true
}

//是否已幻化
func (m *PlayerLingTongDevDataManager) IsUnrealed(classType types.LingTongDevSysType, seqId int) (flag bool) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	uList := obj.GetUnrealList()
	for _, v := range uList {
		if v == seqId {
			//容错处理
			m.LingTongDevOtherGet(classType, int32(seqId))
			return true
		}
	}
	return false
}

//幻化
func (m *PlayerLingTongDevDataManager) Unreal(classType types.LingTongDevSysType, seqId int) (flag bool) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, seqId)
	if lingTongDevTemplate == nil {
		return
	}
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	if !m.IsUnrealed(classType, seqId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj.seqId = int32(seqId)
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
	flag = true
	return
}

//卸下
func (m *PlayerLingTongDevDataManager) Unload(classType types.LingTongDevSysType) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	obj.seqId = 0
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
	return
}

//非进阶灵童养成激活
func (m *PlayerLingTongDevDataManager) LingTongDevOtherGet(classType types.LingTongDevSysType, seqId int32) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))
	if lingTongDevTemplate == nil {
		return
	}
	typ := lingTongDevTemplate.GetType()
	if typ != types.LingTongDevTypeSkin {
		return
	}
	m.initPlayerLingTongOtherObject(classType, seqId)
	return
}

func (m *PlayerLingTongDevDataManager) TongLing(classType types.LingTongDevSysType, pro int32, sucess bool) (flag bool) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	tongLingLevel := obj.GetTongLingLevel()
	nextLevel := tongLingLevel + 1
	lingTongDevTongLingTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTongLingTemplate(classType, nextLevel)
	if lingTongDevTongLingTemplate == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.tongLingLevel += 1
		obj.tongLingNum = 0
		obj.tongLingPro = pro
	} else {
		obj.tongLingNum += 1
		obj.tongLingPro += pro
	}
	obj.updateTime = now
	obj.SetModified()
	return true
}

func (m *PlayerLingTongDevDataManager) EatCulDan(classType types.LingTongDevSysType, level int32) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		return
	}
	if obj.GetCulLevel() == level || level <= 0 {
		return
	}
	lingTongDevHuanHuaTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevHuanHuaTemplate(classType, level)
	if lingTongDevHuanHuaTemplate == nil {
		return
	}
	obj.culLevel = level
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	return
}
