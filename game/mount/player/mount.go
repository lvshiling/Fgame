package player

import (
	"encoding/json"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"sort"

	"fgame/fgame/game/mount/dao"
	mountentity "fgame/fgame/game/mount/entity"
	"fgame/fgame/game/mount/mount"

	gameevent "fgame/fgame/game/event"
	mountcommon "fgame/fgame/game/mount/common"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	mounttypes "fgame/fgame/game/mount/types"

	"github.com/pkg/errors"
)

//坐骑对象
type PlayerMountObject struct {
	player      player.Player
	Id          int64
	PlayerId    int64
	AdvanceId   int   //进阶id
	MountId     int32 //幻化id
	UnrealLevel int32
	UnrealNum   int32
	UnrealPro   int32
	CulLevel    int32
	CulNum      int32
	CulPro      int32
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

func NewPlayerMountObject(pl player.Player) *PlayerMountObject {
	pmo := &PlayerMountObject{
		player: pl,
	}
	return pmo
}

func convertNewPlayerMountObjectToEntity(pmo *PlayerMountObject) (*mountentity.PlayerMountEntity, error) {
	unrealInfoBytes, err := json.Marshal(pmo.UnrealList)
	if err != nil {
		return nil, err
	}

	e := &mountentity.PlayerMountEntity{
		Id:          pmo.Id,
		PlayerId:    pmo.PlayerId,
		AdvancedId:  pmo.AdvanceId,
		MountId:     pmo.MountId,
		UnrealLevel: pmo.UnrealLevel,
		UnrealNum:   pmo.UnrealNum,
		UnrealPro:   pmo.UnrealPro,
		CulLevel:    pmo.CulLevel,
		CulNum:      pmo.CulNum,
		CulPro:      pmo.CulPro,
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

func (pmo *PlayerMountObject) GetPlayerId() int64 {
	return pmo.PlayerId
}

func (pmo *PlayerMountObject) GetDBId() int64 {
	return pmo.Id
}

func (pmo *PlayerMountObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerMountObjectToEntity(pmo)
	return e, err
}

func (pmo *PlayerMountObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*mountentity.PlayerMountEntity)
	var unrealList = make([]int, 0, 8)
	if err := json.Unmarshal([]byte(pse.UnrealInfo), &unrealList); err != nil {
		return err
	}

	pmo.Id = pse.Id
	pmo.PlayerId = pse.PlayerId
	pmo.AdvanceId = pse.AdvancedId
	pmo.MountId = pse.MountId
	pmo.UnrealLevel = pse.UnrealLevel
	pmo.UnrealNum = pse.UnrealNum
	pmo.UnrealPro = pse.UnrealPro
	pmo.CulLevel = pse.CulLevel
	pmo.CulNum = pse.CulNum
	pmo.CulPro = pse.CulPro
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

func (pmo *PlayerMountObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Mount"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}

//坐骑非进阶对象
type PlayerMountOtherObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Typ        mounttypes.MountType
	MountId    int32
	Level      int32
	UpNum      int32
	UpPro      int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerMountOtherObject(pl player.Player) *PlayerMountOtherObject {
	pmo := &PlayerMountOtherObject{
		player: pl,
	}
	return pmo
}

func convertMountOtherObjectToEntity(pmo *PlayerMountOtherObject) (*mountentity.PlayerMountOtherEntity, error) {

	e := &mountentity.PlayerMountOtherEntity{
		Id:         pmo.Id,
		PlayerId:   pmo.PlayerId,
		Typ:        int32(pmo.Typ),
		MountId:    pmo.MountId,
		Level:      pmo.Level,
		UpNum:      pmo.UpNum,
		UpPro:      pmo.UpPro,
		UpdateTime: pmo.UpdateTime,
		CreateTime: pmo.CreateTime,
		DeleteTime: pmo.DeleteTime,
	}
	return e, nil
}

func (pmo *PlayerMountOtherObject) GetPlayerId() int64 {
	return pmo.PlayerId
}

func (pmo *PlayerMountOtherObject) GetDBId() int64 {
	return pmo.Id
}

func (pmo *PlayerMountOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertMountOtherObjectToEntity(pmo)
	return e, err
}

func (pmo *PlayerMountOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*mountentity.PlayerMountOtherEntity)

	pmo.Id = pse.Id
	pmo.PlayerId = pse.PlayerId
	pmo.Typ = mounttypes.MountType(pse.Typ)
	pmo.MountId = pse.MountId
	pmo.Level = pse.Level
	pmo.UpNum = pse.UpNum
	pmo.UpPro = pse.UpPro
	pmo.UpdateTime = pse.UpdateTime
	pmo.CreateTime = pse.CreateTime
	pmo.DeleteTime = pse.DeleteTime
	return nil
}

func (pmo *PlayerMountOtherObject) SetModified() {
	e, err := pmo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "MountOther"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pmo.player.AddChangedObject(obj)
	return
}

//玩家坐骑管理器
type PlayerMountDataManager struct {
	p player.Player
	//玩家坐骑对象
	playerMountObject *PlayerMountObject
	//玩家非进阶坐骑对象
	playerOtherMap map[mounttypes.MountType]map[int32]*PlayerMountOtherObject
}

func (pmdm *PlayerMountDataManager) Player() player.Player {
	return pmdm.p
}

//加载
func (pmdm *PlayerMountDataManager) Load() (err error) {
	pmdm.playerOtherMap = make(map[mounttypes.MountType]map[int32]*PlayerMountOtherObject)
	//加载玩家坐骑信息
	mountEntity, err := dao.GetMountDao().GetMountEntity(pmdm.p.GetId())
	if err != nil {
		return
	}
	if mountEntity == nil {
		pmdm.initPlayerMountObject()
	} else {
		pmdm.playerMountObject = NewPlayerMountObject(pmdm.p)
		pmdm.playerMountObject.FromEntity(mountEntity)
	}

	//加载玩家非进阶坐骑信息
	mountOtherList, err := dao.GetMountDao().GetMountOtherList(pmdm.p.GetId())
	if err != nil {
		return
	}

	//非进阶坐骑信息
	for _, mountOther := range mountOtherList {
		pmo := NewPlayerMountOtherObject(pmdm.p)
		pmo.FromEntity(mountOther)

		typ := mounttypes.MountType(mountOther.Typ)
		playerOtherMap, exist := pmdm.playerOtherMap[typ]
		if !exist {
			playerOtherMap = make(map[int32]*PlayerMountOtherObject)
			pmdm.playerOtherMap[typ] = playerOtherMap
		}
		playerOtherMap[pmo.MountId] = pmo
	}

	return nil
}

//第一次初始化
func (pmdm *PlayerMountDataManager) initPlayerMountObject() {
	pmo := NewPlayerMountObject(pmdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmo.Id = id
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(pmdm.p.GetRole(), pmdm.p.GetSex())
	advanceId := playerCreateTemplate.Mount
	//生成id
	pmo.PlayerId = pmdm.p.GetId()
	pmo.AdvanceId = int(advanceId)
	pmo.MountId = 0
	pmo.Hidden = 1
	pmo.UnrealList = make([]int, 0, 8)
	pmo.UnrealLevel = 0
	pmo.UnrealNum = 0
	pmo.UnrealPro = 0
	pmo.CulLevel = 0
	pmo.CulNum = 0
	pmo.CulPro = 0
	pmo.TimesNum = int32(0)
	pmo.Bless = int32(0)
	pmo.BlessTime = int64(0)
	pmo.Power = int64(0)
	pmo.CreateTime = now
	pmdm.playerMountObject = pmo
	pmo.SetModified()
}

//增加非进阶坐骑
func (pmdm *PlayerMountDataManager) newPlayerMountOtherObject(typ mounttypes.MountType, mountId int32) {

	playerOtherMap, exist := pmdm.playerOtherMap[typ]
	if !exist {
		playerOtherMap = make(map[int32]*PlayerMountOtherObject)
		pmdm.playerOtherMap[typ] = playerOtherMap
	}
	_, exist = playerOtherMap[mountId]
	if exist {
		return
	}
	pmo := NewPlayerMountOtherObject(pmdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmo.Id = id
	//生成id
	pmo.PlayerId = pmdm.p.GetId()
	pmo.Typ = typ
	pmo.MountId = mountId
	pmo.Level = 0
	pmo.UpNum = 0
	pmo.UpPro = 0
	pmo.CreateTime = now
	playerOtherMap[mountId] = pmo
	pmo.SetModified()
	return
}

func (pmdm *PlayerMountDataManager) refreshBless() {
	now := global.GetGame().GetTimeService().Now()
	number := int32(pmdm.playerMountObject.AdvanceId)
	nextNumber := number + 1
	mountTemplate := mount.GetMountService().GetMountNumber(nextNumber)
	if mountTemplate == nil {
		return
	}
	if !mountTemplate.GetIsClear() {
		return
	}
	lastTime := pmdm.playerMountObject.BlessTime
	if lastTime != 0 {
		flag, err := timeutils.IsSameDay(lastTime, now)
		if err != nil {
			return
		}
		if !flag {
			pmdm.playerMountObject.Bless = 0
			pmdm.playerMountObject.BlessTime = 0
			pmdm.playerMountObject.TimesNum = 0
			pmdm.playerMountObject.SetModified()
		}
	}
	return
}

//加载后
func (pmdm *PlayerMountDataManager) AfterLoad() (err error) {
	pmdm.refreshBless()
	return nil
}

//坐骑信息对象
func (pmdm *PlayerMountDataManager) GetMountInfo() *PlayerMountObject {
	pmdm.refreshBless()
	return pmdm.playerMountObject
}

//获取当前阶数
func (pmdm *PlayerMountDataManager) GetMountAdvancedId() int32 {
	return int32(pmdm.playerMountObject.AdvanceId)
}

//获取玩家非进阶坐骑对象
func (pmdm *PlayerMountDataManager) GetMountOtherMap() map[mounttypes.MountType]map[int32]*PlayerMountOtherObject {
	return pmdm.playerOtherMap
}

// func (pmdm *PlayerMountDataManager) EatCulDan(pro int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}

// 	if sucess {
// 		pmdm.playerMountObject.CulLevel += 1
// 		pmdm.playerMountObject.CulNum = 0
// 		pmdm.playerMountObject.CulPro = pro
// 	} else {
// 		pmdm.playerMountObject.CulNum += 1
// 		pmdm.playerMountObject.CulPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	pmdm.playerMountObject.UpdateTime = now
// 	pmdm.playerMountObject.SetModified()
// 	return
// }

func (pmdm *PlayerMountDataManager) EatCulDan(level int32) {
	if pmdm.playerMountObject.CulLevel == level || level <= 0 {
		return
	}
	caoLiaoTemplate := mount.GetMountService().GetMountCaoLiaoTemplate(level)
	if caoLiaoTemplate == nil {
		return
	}
	pmdm.playerMountObject.CulLevel = level
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	return
}

// func (pmdm *PlayerMountDataManager) EatUnrealDan(pro int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}

// 	if sucess {
// 		pmdm.playerMountObject.UnrealLevel += 1
// 		pmdm.playerMountObject.UnrealNum = 0
// 		pmdm.playerMountObject.UnrealPro = pro
// 	} else {
// 		pmdm.playerMountObject.UnrealNum += 1
// 		pmdm.playerMountObject.UnrealPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	pmdm.playerMountObject.UpdateTime = now
// 	pmdm.playerMountObject.SetModified()
// 	return
// }

func (pmdm *PlayerMountDataManager) EatUnrealDan(level int32) {
	if pmdm.playerMountObject.UnrealLevel == level || level <= 0 {
		return
	}
	hunaHuaTemplate := mount.GetMountService().GetMountHuanHuaTemplate(level)
	if hunaHuaTemplate == nil {
		return
	}

	pmdm.playerMountObject.UnrealLevel = level
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	return
}

//卸下
func (pmdm *PlayerMountDataManager) Unload() {
	pmdm.playerMountObject.MountId = 0
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	gameevent.Emit(mounteventtypes.EventTypeMountChanged, pmdm.p, nil)
	return
}

func (pmdm *PlayerMountDataManager) IsHidden() bool {
	return pmdm.playerMountObject.Hidden == 1
}

//展示隐藏坐骑
func (pmdm *PlayerMountDataManager) Hidden(hiddenFlag bool) {
	now := global.GetGame().GetTimeService().Now()
	if hiddenFlag {
		pmdm.playerMountObject.Hidden = 1
	} else {
		pmdm.playerMountObject.Hidden = 0
	}
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	return
}

//进阶
func (pmdm *PlayerMountDataManager) MountAdvanced(pro int32, addTimes int32, sucess bool) {
	if pro < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		nextAdvancedId := pmdm.playerMountObject.AdvanceId + 1
		mountTemplate := mount.GetMountService().GetMountNumber(int32(nextAdvancedId))
		if mountTemplate == nil {
			return
		}
		pmdm.playerMountObject.AdvanceId += 1
		pmdm.playerMountObject.TimesNum = 0
		pmdm.playerMountObject.Bless = 0
		pmdm.playerMountObject.BlessTime = 0
		pmdm.playerMountObject.MountId = int32(0)
		gameevent.Emit(mounteventtypes.EventTypeMountChanged, pmdm.p, nil)
		gameevent.Emit(mounteventtypes.EventTypeMountAdvanced, pmdm.p, pmdm.playerMountObject.AdvanceId)
	} else {
		pmdm.playerMountObject.TimesNum += addTimes
		if pmdm.playerMountObject.Bless == 0 {
			pmdm.playerMountObject.BlessTime = now
		}
		pmdm.playerMountObject.Bless += pro
	}
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	return
}

//直升券进阶
func (pmdm *PlayerMountDataManager) MountAdvancedTicket(addAdvancedNum int32) {
	if addAdvancedNum <= 0 {
		return
	}

	//TODO:xzk:修改逻辑，取最高等级对比
	canAddNum := 0
	nextAdvancedId := pmdm.playerMountObject.AdvanceId + 1
	for addAdvancedNum > 0 {
		mountTemplate := mount.GetMountService().GetMountNumber(int32(nextAdvancedId))
		if mountTemplate == nil {
			break
		}

		canAddNum += 1
		nextAdvancedId += 1
		addAdvancedNum -= 1
	}

	if canAddNum == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.AdvanceId += canAddNum
	pmdm.playerMountObject.TimesNum = 0
	pmdm.playerMountObject.Bless = 0
	pmdm.playerMountObject.BlessTime = 0
	pmdm.playerMountObject.MountId = int32(0)
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	gameevent.Emit(mounteventtypes.EventTypeMountChanged, pmdm.p, nil)
	gameevent.Emit(mounteventtypes.EventTypeMountAdvanced, pmdm.p, pmdm.playerMountObject.AdvanceId)
	return
}

//设置幻化信息
func (pmdm *PlayerMountDataManager) AddUnrealInfo(mountId int) {
	if mountId <= 0 {
		return
	}
	pmdm.playerMountObject.UnrealList = append(pmdm.playerMountObject.UnrealList, mountId)
	sort.Ints(pmdm.playerMountObject.UnrealList)
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.MountId = int32(mountId)
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	pmdm.MountOtherGet(int32(mountId))

	gameevent.Emit(mounteventtypes.EventTypeMountUnrealActivate, pmdm.p, mountId)
	return
}

//是否能幻化
func (pmdm *PlayerMountDataManager) IsCanUnreal(mountId int) bool {
	mountTemplate := mount.GetMountService().GetMount(mountId)
	if mountTemplate == nil {
		return false
	}
	curAdvancedId := pmdm.playerMountObject.AdvanceId
	//食用幻化丹等级
	curUrealDanLevel := pmdm.playerMountObject.UnrealLevel
	//幻化条件
	for condType, cond := range mountTemplate.GetMagicParamXUMap() {
		scondType := mounttypes.MountUCondType(condType)
		switch scondType {
		//坐骑阶别
		case mounttypes.MountUCondTypeX:
			if curAdvancedId < int(cond) {
				return false
			}
		//食用幻化丹数量条件
		case mounttypes.MountUCondTypeU:
			if int32(curUrealDanLevel) < cond {
				return false
			}
		default:
			break
		}
	}
	return true
}

//心跳
func (pmdm *PlayerMountDataManager) Heartbeat() {

}

//是否已幻化
func (pmdm *PlayerMountDataManager) IsUnrealed(mountId int) bool {
	uList := pmdm.playerMountObject.UnrealList
	for _, v := range uList {
		if v == mountId {
			//容错处理
			pmdm.MountOtherGet(int32(mountId))
			return true
		}
	}
	return false
}

//幻化
func (pmdm *PlayerMountDataManager) Unreal(mountId int) (flag bool) {
	if !pmdm.IsUnrealed(mountId) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.MountId = int32(mountId)
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()

	gameevent.Emit(mounteventtypes.EventTypeMountChanged, pmdm.p, nil)
	flag = true
	return
}

//坐骑战斗力
func (pmdm *PlayerMountDataManager) MountPower(power int64) {
	if power <= 0 {
		return
	}
	curPower := pmdm.playerMountObject.Power
	if curPower == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.Power = power
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()

	gameevent.Emit(mounteventtypes.EventTypeMountPowerChanged, pmdm.p, power)
	return
}

//非进阶坐骑激活
func (pmdm *PlayerMountDataManager) MountOtherGet(mountId int32) {
	mountTemplate := mount.GetMountService().GetMount(int(mountId))
	if mountTemplate == nil {
		return
	}
	typ := mountTemplate.GetTyp()
	if typ == mounttypes.MountTypeAdvanced {
		return
	}
	pmdm.newPlayerMountOtherObject(typ, mountId)
	return
}

//是否已拥有该坐骑皮肤
func (pmdm *PlayerMountDataManager) IfMountSkinExist(mountId int32) (*PlayerMountOtherObject, bool) {
	mountTemplate := mount.GetMountService().GetMount(int(mountId))
	if mountTemplate == nil {
		return nil, false
	}
	typ := mountTemplate.GetTyp()
	playerOtherMap, exist := pmdm.playerOtherMap[typ]
	if !exist {
		return nil, false
	}
	mountOtherObj, exist := playerOtherMap[mountId]
	if !exist {
		return nil, false
	}
	return mountOtherObj, true
}

//是否能升星
func (pmdm *PlayerMountDataManager) IfCanUpStar(mountId int32) (*PlayerMountOtherObject, bool) {
	mountTemplate := mount.GetMountService().GetMount(int(mountId))
	if mountTemplate == nil {
		return nil, false
	}

	mountOtherObj, flag := pmdm.IfMountSkinExist(mountId)
	if !flag {
		return nil, false
	}

	if mountTemplate.MountUpstarBeginId == 0 {
		return nil, false
	}

	level := mountOtherObj.Level
	if level <= 0 {
		return mountOtherObj, true
	}
	nextTo := mountTemplate.GetMountUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return mountOtherObj, true
	}
	return nil, false
}

//坐骑皮肤升星
func (pmdm *PlayerMountDataManager) Upstar(mountId int32, pro int32, sucess bool) bool {
	obj, flag := pmdm.IfCanUpStar(mountId)
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

func (pmdm *PlayerMountDataManager) ToMountInfo() *mountcommon.MountInfo {
	info := &mountcommon.MountInfo{
		AdvanceId:   pmdm.playerMountObject.AdvanceId,
		MountId:     pmdm.playerMountObject.MountId,
		UnrealLevel: pmdm.playerMountObject.UnrealLevel,
		UnrealPro:   pmdm.playerMountObject.UnrealPro,
		CulLevel:    pmdm.playerMountObject.CulLevel,
		CulPro:      pmdm.playerMountObject.CulPro,
	}
	for _, typM := range pmdm.playerOtherMap {
		for _, otherObj := range typM {
			skinInfo := &mountcommon.MountSkinInfo{
				MountId: otherObj.MountId,
				Level:   otherObj.Level,
				UpPro:   otherObj.UpPro,
			}
			info.SkinList = append(info.SkinList, skinInfo)
		}
	}
	return info
}

func (pmdm *PlayerMountDataManager) GetMountId() int32 {

	if pmdm.playerMountObject.MountId != 0 {
		return pmdm.playerMountObject.MountId
	}
	mountTemplate := mount.GetMountService().GetMountNumber(int32(pmdm.playerMountObject.AdvanceId))
	if mountTemplate == nil {
		return 0
	}
	return int32(mountTemplate.TemplateId())
}

func (pmdm *PlayerMountDataManager) IfFullAdvanced() (flag bool) {
	if pmdm.playerMountObject.AdvanceId == 0 {
		return
	}
	mountTemplate := mount.GetMountService().GetMountNumber(int32(pmdm.playerMountObject.AdvanceId))
	if mountTemplate == nil {
		return
	}
	if mountTemplate.NextId == 0 {
		return true
	}
	return
}

//仅gm使用 坐骑进阶
func (pmdm *PlayerMountDataManager) GmSetMountAdvanced(advancedId int) {
	pmdm.playerMountObject.AdvanceId = advancedId
	pmdm.playerMountObject.TimesNum = int32(0)
	pmdm.playerMountObject.Bless = int32(0)
	pmdm.playerMountObject.BlessTime = int64(0)
	pmdm.playerMountObject.MountId = int32(0)
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()

	gameevent.Emit(mounteventtypes.EventTypeMountChanged, pmdm.p, nil)
	gameevent.Emit(mounteventtypes.EventTypeMountAdvanced, pmdm.p, pmdm.playerMountObject.AdvanceId)
	return
}

//仅gm使用 坐骑幻化
func (pmdm *PlayerMountDataManager) GmSetMountUnreal(mountId int) {
	now := global.GetGame().GetTimeService().Now()
	if !pmdm.IsUnrealed(mountId) {
		pmdm.playerMountObject.UnrealList = append(pmdm.playerMountObject.UnrealList, mountId)
		sort.Ints(pmdm.playerMountObject.UnrealList)
		pmdm.playerMountObject.UpdateTime = now
		pmdm.playerMountObject.SetModified()
	}
	pmdm.playerMountObject.MountId = int32(mountId)
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	gameevent.Emit(mounteventtypes.EventTypeMountChanged, pmdm.p, nil)
	gameevent.Emit(mounteventtypes.EventTypeMountUnrealActivate, pmdm.p, mountId)
}

//仅gm使用 坐骑食幻化丹等级
func (pmdm *PlayerMountDataManager) GmSetMountUnrealDanLevel(level int32) {
	pmdm.playerMountObject.UnrealLevel = level
	pmdm.playerMountObject.UnrealNum = 0
	pmdm.playerMountObject.UnrealPro = 0

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
	return
}

//仅gm使用 坐骑食草料等级
func (pmdm *PlayerMountDataManager) GmSetMountCaoLiaoLevel(level int32) {

	pmdm.playerMountObject.CulLevel = level
	pmdm.playerMountObject.CulNum = 0
	pmdm.playerMountObject.CulPro = 0

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerMountObject.UpdateTime = now
	pmdm.playerMountObject.SetModified()
}

func CreatePlayerMountDataManager(p player.Player) player.PlayerDataManager {
	pmdm := &PlayerMountDataManager{}
	pmdm.p = p
	return pmdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerMountDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerMountDataManager))
}
