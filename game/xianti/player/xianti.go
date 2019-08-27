package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	xianticommon "fgame/fgame/game/xianti/common"
	"fgame/fgame/game/xianti/dao"
	xiantientity "fgame/fgame/game/xianti/entity"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	xiantitypes "fgame/fgame/game/xianti/types"
	"fgame/fgame/game/xianti/xianti"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"

	"github.com/pkg/errors"
)

//仙体对象
type PlayerXianTiObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	AdvanceId   int   //进阶id
	XianTiId    int32 //幻化id
	UnrealLevel int32
	UnrealNum   int32
	UnrealPro   int32
	UnrealList  []int
	TimesNum    int32
	Bless       int32
	Hidden      int32
	BlessTime   int64
	Power       int64
	UpdateTime  int64
	CreateTime  int64
	DeleteTime  int64
}

func NewPlayerXianTiObject(pl player.Player) *PlayerXianTiObject {
	pmo := &PlayerXianTiObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerXianTiObjectToEntity(pmo *PlayerXianTiObject) (*xiantientity.PlayerXianTiEntity, error) {
	unrealInfoBytes, err := json.Marshal(pmo.UnrealList)
	if err != nil {
		return nil, err
	}

	e := &xiantientity.PlayerXianTiEntity{
		Id:          pmo.Id,
		PlayerId:    pmo.PlayerId,
		AdvancedId:  pmo.AdvanceId,
		XianTiId:    pmo.XianTiId,
		UnrealLevel: pmo.UnrealLevel,
		UnrealNum:   pmo.UnrealNum,
		UnrealPro:   pmo.UnrealPro,
		UnrealInfo:  string(unrealInfoBytes),
		TimesNum:    pmo.TimesNum,
		Bless:       pmo.Bless,
		BlessTime:   pmo.BlessTime,
		Hidden:      pmo.Hidden,
		Power:       pmo.Power,
		UpdateTime:  pmo.UpdateTime,
		CreateTime:  pmo.CreateTime,
		DeleteTime:  pmo.DeleteTime,
	}
	return e, nil
}

func (pmo *PlayerXianTiObject) GetPlayerId() int64 {
	return pmo.PlayerId
}

func (pmo *PlayerXianTiObject) GetDBId() int64 {
	return pmo.Id
}

func (pmo *PlayerXianTiObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerXianTiObjectToEntity(pmo)
	return e, err
}

func (pmo *PlayerXianTiObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*xiantientity.PlayerXianTiEntity)
	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}

	pmo.Id = pse.Id
	pmo.PlayerId = pse.PlayerId
	pmo.AdvanceId = pse.AdvancedId
	pmo.XianTiId = pse.XianTiId
	pmo.UnrealLevel = pse.UnrealLevel
	pmo.UnrealNum = pse.UnrealNum
	pmo.UnrealPro = pse.UnrealPro
	pmo.UnrealList = unrealList
	pmo.TimesNum = pse.TimesNum
	pmo.Bless = pse.Bless
	pmo.BlessTime = pse.BlessTime
	pmo.Hidden = pse.Hidden
	pmo.Power = pse.Power
	pmo.UpdateTime = pse.UpdateTime
	pmo.CreateTime = pse.CreateTime
	pmo.DeleteTime = pse.DeleteTime
	return nil
}

func (pmo *PlayerXianTiObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "XianTi"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}

//仙体非进阶对象
type PlayerXianTiOtherObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Typ        xiantitypes.XianTiType
	XianTiId   int32
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerXianTiOtherObject(pl player.Player) *PlayerXianTiOtherObject {
	pmo := &PlayerXianTiOtherObject{
		player: pl,
	}
	return pmo
}

func convertXianTiOtherObjectToEntity(pmo *PlayerXianTiOtherObject) (*xiantientity.PlayerXianTiOtherEntity, error) {

	e := &xiantientity.PlayerXianTiOtherEntity{
		Id:         pmo.Id,
		PlayerId:   pmo.PlayerId,
		Typ:        int32(pmo.Typ),
		XianTiId:   pmo.XianTiId,
		Level:      pmo.Level,
		UpNum:      pmo.UpNum,
		UpPro:      pmo.UpPro,
		UpdateTime: pmo.UpdateTime,
		CreateTime: pmo.CreateTime,
		DeleteTime: pmo.DeleteTime,
	}
	return e, nil
}

func (pmo *PlayerXianTiOtherObject) GetPlayerId() int64 {
	return pmo.PlayerId
}

func (pmo *PlayerXianTiOtherObject) GetDBId() int64 {
	return pmo.Id
}

func (pmo *PlayerXianTiOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertXianTiOtherObjectToEntity(pmo)
	return e, err
}

func (pmo *PlayerXianTiOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*xiantientity.PlayerXianTiOtherEntity)

	pmo.Id = pse.Id
	pmo.PlayerId = pse.PlayerId
	pmo.Typ = xiantitypes.XianTiType(pse.Typ)
	pmo.XianTiId = pse.XianTiId
	pmo.Level = pse.Level
	pmo.UpNum = pse.UpNum
	pmo.UpPro = pse.UpPro
	pmo.UpdateTime = pse.UpdateTime
	pmo.CreateTime = pse.CreateTime
	pmo.DeleteTime = pse.DeleteTime
	return nil
}

func (pmo *PlayerXianTiOtherObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "XianTiOther"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}

//玩家仙体管理器
type PlayerXianTiDataManager struct {
	p player.Player
	//玩家仙体对象
	playerXianTiObject *PlayerXianTiObject
	//玩家非进阶仙体对象
	playerOtherMap map[xiantitypes.XianTiType]map[int32]*PlayerXianTiOtherObject
}

func (pmdm *PlayerXianTiDataManager) Player() player.Player {
	return pmdm.p
}

//加载
func (pmdm *PlayerXianTiDataManager) Load() (err error) {
	pmdm.playerOtherMap = make(map[xiantitypes.XianTiType]map[int32]*PlayerXianTiOtherObject)
	//加载玩家仙体信息
	xianTiEntity, err := dao.GetXianTiDao().GetXianTiEntity(pmdm.p.GetId())
	if err != nil {
		return
	}
	if xianTiEntity == nil {
		pmdm.initPlayerXianTiObject()
	} else {
		pmdm.playerXianTiObject = NewPlayerXianTiObject(pmdm.p)
		pmdm.playerXianTiObject.FromEntity(xianTiEntity)
	}

	//加载玩家非进阶仙体信息
	xianTiOtherList, err := dao.GetXianTiDao().GetXianTiOtherList(pmdm.p.GetId())
	if err != nil {
		return
	}

	//非进阶仙体信息
	for _, xianTiOther := range xianTiOtherList {
		pmo := NewPlayerXianTiOtherObject(pmdm.p)
		pmo.FromEntity(xianTiOther)

		typ := xiantitypes.XianTiType(xianTiOther.Typ)
		playerOtherMap, exist := pmdm.playerOtherMap[typ]
		if !exist {
			playerOtherMap = make(map[int32]*PlayerXianTiOtherObject)
			pmdm.playerOtherMap[typ] = playerOtherMap
		}
		playerOtherMap[pmo.XianTiId] = pmo
	}

	return nil
}

//第一次初始化
func (pmdm *PlayerXianTiDataManager) initPlayerXianTiObject() {
	pmo := NewPlayerXianTiObject(pmdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmo.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(pmdm.p.GetRole(), pmdm.p.GetSex())
	advanceId := playerCreateTemplate.XianTi
	//生成id
	pmo.PlayerId = pmdm.p.GetId()
	pmo.AdvanceId = int(advanceId)
	pmo.XianTiId = 0
	pmo.Hidden = 0
	pmo.UnrealList = make([]int, 0, 8)
	pmo.UnrealLevel = 0
	pmo.UnrealNum = 0
	pmo.UnrealPro = 0
	pmo.TimesNum = int32(0)
	pmo.Bless = int32(0)
	pmo.BlessTime = int64(0)
	pmo.Power = int64(0)
	pmo.CreateTime = now
	pmdm.playerXianTiObject = pmo
	pmo.SetModified()
}

//增加非进阶仙体
func (pmdm *PlayerXianTiDataManager) newPlayerXianTiOtherObject(typ xiantitypes.XianTiType, xianTiId int32) {

	playerOtherMap, exist := pmdm.playerOtherMap[typ]
	if !exist {
		playerOtherMap = make(map[int32]*PlayerXianTiOtherObject)
		pmdm.playerOtherMap[typ] = playerOtherMap
	}
	_, exist = playerOtherMap[xianTiId]
	if exist {
		return
	}
	pmo := NewPlayerXianTiOtherObject(pmdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmo.Id = id
	//生成id
	pmo.PlayerId = pmdm.p.GetId()
	pmo.Typ = typ
	pmo.XianTiId = xianTiId
	pmo.Level = 0
	pmo.UpNum = 0
	pmo.UpPro = 0
	pmo.CreateTime = now
	playerOtherMap[xianTiId] = pmo
	pmo.SetModified()
	return
}

func (pmdm *PlayerXianTiDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(pmdm.playerXianTiObject.AdvanceId)
	nextNumber := number + 1
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(nextNumber)
	if xianTiTemplate == nil {
		return
	}
	if !xianTiTemplate.GetIsClear() {
		return
	}
	lastTime := pmdm.playerXianTiObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			pmdm.playerXianTiObject.Bless = 0
			pmdm.playerXianTiObject.BlessTime = 0
			pmdm.playerXianTiObject.TimesNum = 0
			pmdm.playerXianTiObject.SetModified()
		}
	}
	return
}

//加载后
func (pmdm *PlayerXianTiDataManager) AfterLoad() (err error) {
	pmdm.refreshBless()
	return nil
}

//仙体信息对象
func (pmdm *PlayerXianTiDataManager) GetXianTiInfo() *PlayerXianTiObject {
	pmdm.refreshBless()
	return pmdm.playerXianTiObject
}

//获取当前阶数
func (pmdm *PlayerXianTiDataManager) GetXianTiAdvancedId() int32 {
	return int32(pmdm.playerXianTiObject.AdvanceId)
}

//获取玩家非进阶仙体对象
func (pmdm *PlayerXianTiDataManager) GetXianTiOtherMap() map[xiantitypes.XianTiType]map[int32]*PlayerXianTiOtherObject {
	return pmdm.playerOtherMap
}

func (pmdm *PlayerXianTiDataManager) EatUnrealDan(level int32) {
	if pmdm.playerXianTiObject.UnrealLevel == level || level <= 0 {
		return
	}
	hunaHuaTemplate := xianti.GetXianTiService().GetXianTiHuanHuaTemplate(level)
	if hunaHuaTemplate == nil {
		return
	}

	pmdm.playerXianTiObject.UnrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	return
}

//卸下
func (pmdm *PlayerXianTiDataManager) Unload() {
	pmdm.playerXianTiObject.XianTiId = 0
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
	return
}

func (pmdm *PlayerXianTiDataManager) IsHidden() bool {
	return pmdm.playerXianTiObject.Hidden == 1
}

//展示隐藏仙体
func (pmdm *PlayerXianTiDataManager) Hidden(hiddenFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		pmdm.playerXianTiObject.Hidden = 1
	} else {
		pmdm.playerXianTiObject.Hidden = 0
	}
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
	// if !hiddenFlag {
	// 	gameevent.Emit(xiantieventtypes.EventTypeXianTiUse, pmdm.p, nil)
	// }
	return
}

//进阶
func (pmdm *PlayerXianTiDataManager) XianTiAdvanced(pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextAdvancedId := pmdm.playerXianTiObject.AdvanceId + 1
		xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(nextAdvancedId))
		if xianTiTemplate == nil {
			return
		}
		pmdm.playerXianTiObject.AdvanceId += 1
		pmdm.playerXianTiObject.TimesNum = 0
		pmdm.playerXianTiObject.Bless = 0
		pmdm.playerXianTiObject.BlessTime = 0
		pmdm.playerXianTiObject.XianTiId = int32(0)
		gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
		gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvanced, pmdm.p, pmdm.playerXianTiObject.AdvanceId)
	} else {
		pmdm.playerXianTiObject.TimesNum += addTimes
		if pmdm.playerXianTiObject.Bless == 0 {
			pmdm.playerXianTiObject.BlessTime = now
		}
		pmdm.playerXianTiObject.Bless += pro
	}
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	return
}

//直升券进阶
func (pmdm *PlayerXianTiDataManager) XianTiAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	canAddNum := 0
	nextAdvancedId := pmdm.playerXianTiObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(nextAdvancedId))
		if xianTiTemplate == nil {
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
	pmdm.playerXianTiObject.AdvanceId += canAddNum
	pmdm.playerXianTiObject.TimesNum = 0
	pmdm.playerXianTiObject.Bless = 0
	pmdm.playerXianTiObject.BlessTime = 0
	pmdm.playerXianTiObject.XianTiId = int32(0)
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
	gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvanced, pmdm.p, pmdm.playerXianTiObject.AdvanceId)
	return
}

//设置幻化信息
func (pmdm *PlayerXianTiDataManager) AddUnrealInfo(xianTiId int) {
	if xianTiId <= 0 {
		return
	}
	pmdm.playerXianTiObject.UnrealList = append(pmdm.playerXianTiObject.UnrealList, xianTiId)
	sort.Ints(pmdm.playerXianTiObject.UnrealList)
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.XianTiId = int32(xianTiId)
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	pmdm.XianTiOtherGet(int32(xianTiId))

	gameevent.Emit(xiantieventtypes.EventTypeXianTiUnrealActivate, pmdm.p, xianTiId)
	return
}

//是否能幻化
func (pmdm *PlayerXianTiDataManager) IsCanUnreal(xianTiId int) bool {
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(xianTiId)
	if xianTiTemplate == nil {
		return false
	}
	curAdvancedId := pmdm.playerXianTiObject.AdvanceId
	//食用幻化丹等级
	curUrealDanLevel := pmdm.playerXianTiObject.UnrealLevel
	//幻化条件
	for condType, cond := range xianTiTemplate.GetMagicParamXUMap() {
		scondType := xiantitypes.XianTiUCondType(condType)
		switch scondType {
		//仙体阶别
		case xiantitypes.XianTiUCondTypeX:
			if curAdvancedId < int(cond) {
				return false
			}
		//食用幻化丹数量条件
		case xiantitypes.XianTiUCondTypeU:
			if int32(curUrealDanLevel) < cond {
				return false
			}
			//关联其他功能等级
		case xiantitypes.XianTiUCondTypeW:
			return false
		default:
			break
		}
	}
	return true
}

//心跳
func (pmdm *PlayerXianTiDataManager) Heartbeat() {

}

//是否已幻化
func (pmdm *PlayerXianTiDataManager) IsUnrealed(xianTiId int) bool {
	uList := pmdm.playerXianTiObject.UnrealList
	for _, v := range uList {
		if v == xianTiId {
			//容错处理
			pmdm.XianTiOtherGet(int32(xianTiId))
			return true
		}
	}
	return false
}

//幻化
func (pmdm *PlayerXianTiDataManager) Unreal(xianTiId int) (flag bool) {
	if !pmdm.IsUnrealed(xianTiId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.XianTiId = int32(xianTiId)
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()

	gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
	flag = true
	return
}

//仙体战斗力
func (pmdm *PlayerXianTiDataManager) XianTiPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := pmdm.playerXianTiObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.Power = power
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	gameevent.Emit(xiantieventtypes.EventTypeXianTiPowerChanged, pmdm.p, power)
	return
}

//非进阶仙体激活
func (pmdm *PlayerXianTiDataManager) XianTiOtherGet(xianTiId int32) {
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(int(xianTiId))
	if xianTiTemplate == nil {
		return
	}
	typ := xianTiTemplate.GetTyp()
	if typ == xiantitypes.XianTiTypeAdvanced {
		return
	}
	pmdm.newPlayerXianTiOtherObject(typ, xianTiId)
	return
}

//是否已拥有该仙体皮肤
func (pmdm *PlayerXianTiDataManager) IfXianTiSkinExist(xianTiId int32) (*PlayerXianTiOtherObject, bool) {
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(int(xianTiId))
	if xianTiTemplate == nil {
		return nil, false
	}
	typ := xianTiTemplate.GetTyp()
	playerOtherMap, exist := pmdm.playerOtherMap[typ]
	if !exist {
		return nil, false
	}
	xianTiOtherObj, exist := playerOtherMap[xianTiId]
	if !exist {
		return nil, false
	}
	return xianTiOtherObj, true
}

//是否能升星
func (pmdm *PlayerXianTiDataManager) IfCanUpStar(xianTiId int32) (*PlayerXianTiOtherObject, bool) {
	xianTiTemplate := xianti.GetXianTiService().GetXianTi(int(xianTiId))
	if xianTiTemplate == nil {
		return nil, false
	}

	xianTiOtherObj, flag := pmdm.IfXianTiSkinExist(xianTiId)
	if !flag {
		return nil, false
	}

	if xianTiTemplate.XianTiUpstarBeginId == 0 {
		return nil, false
	}

	level := xianTiOtherObj.Level
	if level <= 0 {
		return xianTiOtherObj, true
	}
	nextTo := xianTiTemplate.GetXianTiUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return xianTiOtherObj, true
	}
	return nil, false
}

//仙体皮肤升星
func (pmdm *PlayerXianTiDataManager) Upstar(xianTiId int32, pro int32, sucess bool) bool {
	obj, flag := pmdm.IfCanUpStar(xianTiId)
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

func (pmdm *PlayerXianTiDataManager) ToXianTiInfo() *xianticommon.XianTiInfo {
	info := &xianticommon.XianTiInfo{
		AdvanceId:   pmdm.playerXianTiObject.AdvanceId,
		XianTiId:    pmdm.playerXianTiObject.XianTiId,
		UnrealLevel: pmdm.playerXianTiObject.UnrealLevel,
		UnrealPro:   pmdm.playerXianTiObject.UnrealPro,
	}
	for _, typM := range pmdm.playerOtherMap {
		for _, otherObj := range typM {
			skinInfo := &xianticommon.XianTiSkinInfo{
				XianTiId: otherObj.XianTiId,
				Level:    otherObj.Level,
				UpPro:    otherObj.UpPro,
			}
			info.SkinList = append(info.SkinList, skinInfo)
		}
	}
	return info
}

func (pmdm *PlayerXianTiDataManager) GetXianTiId() int32 {
	if pmdm.playerXianTiObject.Hidden != 0 {
		return 0
	}
	if pmdm.playerXianTiObject.XianTiId != 0 {
		return pmdm.playerXianTiObject.XianTiId
	}
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(pmdm.playerXianTiObject.AdvanceId))
	if xianTiTemplate == nil {
		return 0
	}
	return int32(xianTiTemplate.TemplateId())
}

func (pmdm *PlayerXianTiDataManager) IfFullAdvanced() (flag bool) {
	if pmdm.playerXianTiObject.AdvanceId == 0 {
		return
	}
	xianTiTemplate := xianti.GetXianTiService().GetXianTiNumber(int32(pmdm.playerXianTiObject.AdvanceId))
	if xianTiTemplate == nil {
		return
	}
	if xianTiTemplate.NextId == 0 {
		return true
	}
	return
}

//仅gm使用 仙体进阶
func (pmdm *PlayerXianTiDataManager) GmSetXianTiAdvanced(advancedId int) {
	pmdm.playerXianTiObject.AdvanceId = advancedId
	pmdm.playerXianTiObject.TimesNum = int32(0)
	pmdm.playerXianTiObject.Bless = int32(0)
	pmdm.playerXianTiObject.BlessTime = int64(0)
	pmdm.playerXianTiObject.XianTiId = int32(0)
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
	gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvanced, pmdm.p, pmdm.playerXianTiObject.AdvanceId)
	return
}

//仅gm使用 仙体幻化
func (pmdm *PlayerXianTiDataManager) GmSetXianTiUnreal(xianTiId int) {
	now := global.GetGame().GetTimeService().Now()
	if !pmdm.IsUnrealed(xianTiId) {
		pmdm.playerXianTiObject.UnrealList = append(pmdm.playerXianTiObject.UnrealList, xianTiId)
		sort.Ints(pmdm.playerXianTiObject.UnrealList)
		pmdm.playerXianTiObject.UpdateTime = now
		pmdm.playerXianTiObject.SetModified()
	}

	pmdm.playerXianTiObject.XianTiId = int32(xianTiId)
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	gameevent.Emit(xiantieventtypes.EventTypeXianTiChanged, pmdm.p, nil)
	gameevent.Emit(xiantieventtypes.EventTypeXianTiUnrealActivate, pmdm.p, xianTiId)
}

//仅gm使用 仙体食幻化丹等级
func (pmdm *PlayerXianTiDataManager) GmSetXianTiUnrealDanLevel(level int32) {
	pmdm.playerXianTiObject.UnrealLevel = level
	pmdm.playerXianTiObject.UnrealNum = 0
	pmdm.playerXianTiObject.UnrealPro = 0

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXianTiObject.UpdateTime = now
	pmdm.playerXianTiObject.SetModified()
	return
}

func CreatePlayerXianTiDataManager(p player.Player) player.PlayerDataManager {
	pmdm := &PlayerXianTiDataManager{}
	pmdm.p = p
	return pmdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXianTiDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXianTiDataManager))
}
