package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/weapon/dao"
	weapontypes "fgame/fgame/game/weapon/types"
	"fgame/fgame/game/weapon/weapon"
	"fmt"

	gameevent "fgame/fgame/game/event"
	weaponentity "fgame/fgame/game/weapon/entity"
	weaponeventtypes "fgame/fgame/game/weapon/event/types"
	"fgame/fgame/pkg/idutil"
)

//玩家兵魂数据
type PlayerWeaponObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	WeaponId   int32
	ActiveFlag int32
	Level      int32
	UpNum      int32
	UpPro      int32
	CulLevel   int32
	CulNum     int32
	CulPro     int32
	State      weapontypes.WeaponAwakenStatusType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerWeaponObject(pl player.Player) *PlayerWeaponObject {
	pso := &PlayerWeaponObject{
		player: pl,
	}
	return pso
}

func convertPlayerWeaponObjectToEntity(pwo *PlayerWeaponObject) (pse *weaponentity.PlayerWeaponEntity, err error) {

	e := &weaponentity.PlayerWeaponEntity{
		Id:         pwo.Id,
		PlayerId:   pwo.PlayerId,
		WeaponId:   pwo.WeaponId,
		ActiveFlag: pwo.ActiveFlag,
		Level:      pwo.Level,
		UpNum:      pwo.UpNum,
		UpPro:      pwo.UpPro,
		CulLevel:   pwo.CulLevel,
		CulNum:     pwo.CulNum,
		CulPro:     pwo.CulPro,
		State:      int32(pwo.State),
		UpdateTime: pwo.UpdateTime,
		CreateTime: pwo.CreateTime,
		DeleteTime: pwo.DeleteTime,
	}
	return e, nil
}

func (pwo *PlayerWeaponObject) GetPlayerId() int64 {
	return pwo.PlayerId
}

func (pwo *PlayerWeaponObject) GetDBId() int64 {
	return pwo.Id
}

func (pwo *PlayerWeaponObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertPlayerWeaponObjectToEntity(pwo)
	if err != nil {
		return
	}
	return
}

func (pwo *PlayerWeaponObject) FromEntity(e storage.Entity) (err error) {
	pse, _ := e.(*weaponentity.PlayerWeaponEntity)
	pwo.Id = pse.Id
	pwo.WeaponId = pse.WeaponId
	pwo.ActiveFlag = pse.ActiveFlag
	pwo.Level = pse.Level
	pwo.UpNum = pse.UpNum
	pwo.UpPro = pse.UpPro
	pwo.CulLevel = pse.CulLevel
	pwo.CulNum = pse.CulNum
	pwo.CulPro = pse.CulPro
	pwo.State = weapontypes.WeaponAwakenStatusType(pse.State)
	pwo.PlayerId = pse.PlayerId
	pwo.UpdateTime = pse.UpdateTime
	pwo.CreateTime = pse.CreateTime
	pwo.DeleteTime = pse.DeleteTime
	return
}

func (pwo *PlayerWeaponObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		return
	}

	obj, _ := e.(types.PlayerDataEntity)
	if obj == nil {
		panic("never reach here")
	}
	pwo.player.AddChangedObject(obj)
	return
}

//兵魂信息对象
type PlayerWeaponInfoObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	WeaponWear int32
	Star       int32
	Power      int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerWeaponInfoObject(pl player.Player) *PlayerWeaponInfoObject {
	pwwo := &PlayerWeaponInfoObject{
		player: pl,
	}
	return pwwo
}

func convertNewPlayerWeaponInfoObjectToEntity(pwwo *PlayerWeaponInfoObject) (*weaponentity.PlayerWeaponInfoEntity, error) {
	e := &weaponentity.PlayerWeaponInfoEntity{
		Id:         pwwo.Id,
		PlayerId:   pwwo.PlayerId,
		WeaponWear: pwwo.WeaponWear,
		Star:       pwwo.Star,
		Power:      pwwo.Power,
		UpdateTime: pwwo.UpdateTime,
		CreateTime: pwwo.CreateTime,
		DeleteTime: pwwo.DeleteTime,
	}
	return e, nil
}

func (pwwo *PlayerWeaponInfoObject) GetPlayerId() int64 {
	return pwwo.PlayerId
}

func (pwwo *PlayerWeaponInfoObject) GetDBId() int64 {
	return pwwo.Id
}

func (pwwo *PlayerWeaponInfoObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerWeaponInfoObjectToEntity(pwwo)
	return e, err
}

func (pwwo *PlayerWeaponInfoObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*weaponentity.PlayerWeaponInfoEntity)

	pwwo.Id = pse.Id
	pwwo.PlayerId = pse.PlayerId
	pwwo.WeaponWear = pse.WeaponWear
	pwwo.Star = pse.Star
	pwwo.Power = pse.Power
	pwwo.UpdateTime = pse.UpdateTime
	pwwo.CreateTime = pse.CreateTime
	pwwo.DeleteTime = pse.DeleteTime
	return nil
}

func (pwwo *PlayerWeaponInfoObject) SetModified() {
	e, err := pwwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("weapon: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwwo.player.AddChangedObject(obj)
	return
}

//玩家兵魂数据管理器
type PlayerWeaponDataManager struct {
	p player.Player
	//兵魂数据
	playerWeaponObjectMap map[int32]*PlayerWeaponObject
	//兵魂信息
	playerWeaponInfoObject *PlayerWeaponInfoObject
}

//获取所有兵魂
func (pwdm *PlayerWeaponDataManager) GetAllWeapon() map[int32]*PlayerWeaponObject {
	return pwdm.playerWeaponObjectMap
}

//根据兵魂id获取兵魂信息
func (pwdm *PlayerWeaponDataManager) GetWeapon(weaponId int32) (pwo *PlayerWeaponObject) {
	if v, ok := pwdm.playerWeaponObjectMap[weaponId]; ok {
		return v
	}
	return nil
}

//根据兵魂id获取兵魂等级
func (pwdm *PlayerWeaponDataManager) GetWeaponLevel(weaponId int32) int32 {
	obj, ok := pwdm.playerWeaponObjectMap[weaponId]
	if !ok {
		return 0
	}
	return obj.Level
}

//根据兵魂id获取兵魂培养等级
func (pwdm *PlayerWeaponDataManager) GetWeaponPeiYangLevel(weaponId int32) int32 {
	obj, ok := pwdm.playerWeaponObjectMap[weaponId]
	if !ok {
		return 0
	}
	return obj.CulLevel
}

//根据兵魂id获取兵魂觉醒状态
func (pwdm *PlayerWeaponDataManager) GetWeaponState(weaponId int32) weapontypes.WeaponAwakenStatusType {
	obj, ok := pwdm.playerWeaponObjectMap[weaponId]
	if !ok {
		return 0
	}
	return obj.State
}

//获取穿戴兵魂
func (pwdm *PlayerWeaponDataManager) GetWeaponWear() int32 {
	return pwdm.playerWeaponInfoObject.WeaponWear
}

func (pwdm *PlayerWeaponDataManager) Player() player.Player {
	return pwdm.p
}

//加载
func (pwdm *PlayerWeaponDataManager) Load() (err error) {
	//TODO 数据加载封装
	pwdm.playerWeaponObjectMap = make(map[int32]*PlayerWeaponObject)
	pseList, err := dao.GetWeaponDao().GetWeaponList(pwdm.p.GetId())
	if err != nil {
		return
	}
	for _, pse := range pseList {
		pwo := NewPlayerWeaponObject(pwdm.p)
		pwo.FromEntity(pse)
		pwdm.playerWeaponObjectMap[pwo.WeaponId] = pwo
	}

	//加载玩家穿戴兵魂信息
	weaponInfoEntity, err := dao.GetWeaponDao().GetWeaponInfoEntity(pwdm.p.GetId())
	if err != nil {
		return
	}
	if weaponInfoEntity == nil {
		pwdm.initPlayerWeaponInfoObject()
	} else {
		pwdm.playerWeaponInfoObject = NewPlayerWeaponInfoObject(pwdm.p)
		pwdm.playerWeaponInfoObject.FromEntity(weaponInfoEntity)
	}
	return
}

//第一次初始化
func (pwdm *PlayerWeaponDataManager) initPlayerWeaponInfoObject() {
	bornWeapon := pwdm.getBornWeapon()

	//初始化第一套冰魂
	if !pwdm.IfWeaponExist(bornWeapon) {
		flag := pwdm.WeaponActive(bornWeapon, false)
		if !flag {
			panic(fmt.Errorf("fashion:初始化第一套冰魂应该成功"))
		}
	}

	pwwo := NewPlayerWeaponInfoObject(pwdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwwo.Id = id
	//生成id
	pwwo.PlayerId = pwdm.p.GetId()
	pwwo.WeaponWear = bornWeapon
	pwwo.Star = 0
	pwwo.Power = 0
	pwwo.CreateTime = now
	pwdm.playerWeaponInfoObject = pwwo
	pwwo.SetModified()
}

func (pwdm *PlayerWeaponDataManager) AfterLoad() (err error) {
	return nil
}

func (pwdm *PlayerWeaponDataManager) Heartbeat() {

}

//参数有效性
func (pwdm *PlayerWeaponDataManager) IsValid(weaponId int32) bool {
	if weaponId <= 0 {
		return false
	}
	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if weaponTemplate == nil {
		return false
	}
	return true
}

// func (pwdm *PlayerWeaponDataManager) EatCulDan(weaponId int32, pro int32, sucess bool) {
// 	if pro < 0 {
// 		return
// 	}
// 	weaponInfo := pwdm.GetWeapon(weaponId)
// 	if weaponInfo == nil {
// 		return
// 	}
// 	if sucess {
// 		weaponInfo.CulLevel += 1
// 		weaponInfo.CulNum = 0
// 		weaponInfo.CulPro = pro
// 	} else {
// 		weaponInfo.CulNum += 1
// 		weaponInfo.CulPro += pro
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	weaponInfo.UpdateTime = now
// 	weaponInfo.SetModified()
// 	return
// }

func (pwdm *PlayerWeaponDataManager) EatCulDan(weaponId int32, level int32) {
	if level <= 0 {
		return
	}
	weaponInfo := pwdm.GetWeapon(weaponId)
	if weaponInfo == nil {
		return
	}
	if weaponInfo.CulLevel == level {
		return
	}
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to == nil {
		return
	}
	culTemplate := to.GetWeaponPeiYangByLevel(level)
	if culTemplate == nil {
		return
	}
	weaponInfo.CulLevel = level
	now := global.GetGame().GetTimeService().Now()
	weaponInfo.UpdateTime = now
	weaponInfo.SetModified()
	return
}

//是否配置觉醒
func (pwdm *PlayerWeaponDataManager) IfIsAwaken(weaponId int32) bool {
	flag := pwdm.IsValid(weaponId)
	if !flag {
		return false
	}
	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if weaponTemplate.IsAwaken == int32(weapontypes.WeaponAwakenTypeNo) {
		return false
	}
	return true
}

//是否已拥有该兵魂
func (pwdm *PlayerWeaponDataManager) IfWeaponExist(weaponId int32) bool {
	//该兵魂是否已获得
	weapon := pwdm.GetWeapon(weaponId)
	if weapon == nil {
		return false
	}
	if weapon.ActiveFlag == 0 {
		return false
	}
	return true
}

//是否能升星
func (pwdm *PlayerWeaponDataManager) IfCanUpStar(weaponId int32) bool {
	flag := pwdm.IfWeaponExist(weaponId)
	if !flag {
		return false
	}
	level := pwdm.GetWeaponLevel(weaponId)
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to.WeaponUpgradeBeginId == 0 {
		return false
	}
	if level <= 0 {
		return true
	}
	nextTo := to.GetWeaponUpstarByLevel(level)
	if nextTo.NextId != 0 {
		return true
	}
	return false
}

//是否能培养
func (pwdm *PlayerWeaponDataManager) IfCanPeiYang(weaponId int32) bool {
	flag := pwdm.IfWeaponExist(weaponId)
	if !flag {
		return false
	}
	level := pwdm.GetWeaponPeiYangLevel(weaponId)
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to.WeaponPeiYangBeginId == 0 {
		return false
	}
	if level <= 0 {
		return true
	}
	nextTo := to.GetWeaponPeiYangByLevel(level)
	if nextTo.NextId != 0 {
		return true
	}
	return false
}

func (pwdm *PlayerWeaponDataManager) weaponActiveInit(weaponId int32) (flag bool) {
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	now := global.GetGame().GetTimeService().Now()
	pwo := NewPlayerWeaponObject(pwdm.p)

	id, err := idutil.GetId()
	if err != nil {
		return false
	}
	pwo.Id = id
	pwo.PlayerId = pwdm.p.GetId()
	pwo.WeaponId = weaponId
	pwo.ActiveFlag = 1
	pwo.Level = int32(0)
	pwo.CulLevel = int32(0)
	pwo.CulNum = 0
	pwo.CulPro = 0
	pwo.State = weapontypes.WeaponAwakenStatusTypeNo
	if to.IsAwaken == int32(weapontypes.WeaponAwakenTypeOk) &&
		to.NeedStar == 0 {
		pwo.State = weapontypes.WeaponAwakenStatusTypeOk
	}
	pwo.CreateTime = now
	pwo.SetModified()
	pwdm.playerWeaponObjectMap[weaponId] = pwo
	return true
}

//永久兵魂激活
func (pwdm *PlayerWeaponDataManager) WeaponActive(weaponId int32, sendEvent bool) bool {
	flag := pwdm.IsValid(weaponId)
	if !flag {
		return false
	}
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to == nil {
		return false
	}
	if to.GetWeaponTag() != weapontypes.WeaponTagTypePermanent {
		return false
	}
	flag = pwdm.IfWeaponExist(weaponId)
	if flag {
		return false
	}
	pwdm.weaponActiveInit(weaponId)
	if sendEvent {
		gameevent.Emit(weaponeventtypes.EventTypeWeaponActivate, pwdm.p, weaponId)
	}
	return true
}

//临时兵魂激活
func (pwdm *PlayerWeaponDataManager) WeaponActiveTemp(weaponId int32) bool {
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to == nil {
		return false
	}
	if to.GetWeaponTag() != weapontypes.WeaponTagTypeTemp {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	weaponObject, exist := pwdm.playerWeaponObjectMap[weaponId]
	if exist {
		if weaponObject.ActiveFlag != 0 {
			return false
		}
		weaponObject.ActiveFlag = 1
		weaponObject.UpdateTime = now
		weaponObject.SetModified()
	} else {
		pwdm.weaponActiveInit(weaponId)
	}
	gameevent.Emit(weaponeventtypes.EventTypeWeaponActivate, pwdm.p, weaponId)
	return true
}

//临时兵魂移除
func (pwdm *PlayerWeaponDataManager) WeaponRemoveTemp(weaponId int32) bool {
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if to == nil {
		return false
	}
	if to.GetWeaponTag() != weapontypes.WeaponTagTypeTemp {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	weaponObject, exist := pwdm.playerWeaponObjectMap[weaponId]
	if !exist {
		return false
	}
	wearWeaponId := pwdm.playerWeaponInfoObject.WeaponWear
	if weaponObject.WeaponId == wearWeaponId {
		bornWeapon := pwdm.getBornWeapon()
		pwdm.playerWeaponInfoObject.WeaponWear = bornWeapon
		pwdm.playerWeaponInfoObject.UpdateTime = now
		pwdm.playerWeaponInfoObject.SetModified()
		gameevent.Emit(weaponeventtypes.EventTypeWeaponChanged, pwdm.p, nil)
	}
	weaponObject.ActiveFlag = 0
	weaponObject.UpdateTime = now
	weaponObject.SetModified()
	gameevent.Emit(weaponeventtypes.EventTypeWeaponRemove, pwdm.p, weaponId)
	return true
}

//兵魂升星
func (pwdm *PlayerWeaponDataManager) Upstar(weaponId int32, pro int32, sucess bool) bool {
	flag := pwdm.IfCanUpStar(weaponId)
	if !flag {
		return false
	}
	obj := pwdm.GetWeapon(weaponId)
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
		if to == nil {
			return false
		}
		weaponUpstarTemplate := to.GetWeaponUpstarByLevel(obj.Level + 1)
		if weaponUpstarTemplate == nil {
			return false
		}
		obj.Level += 1
		obj.UpNum = 0
		obj.UpPro = pro
		pwdm.playerWeaponInfoObject.Star += 1
		pwdm.playerWeaponInfoObject.SetModified()
	} else {
		obj.UpNum += 1
		obj.UpPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return true
}

//觉醒星数是否ok
func (pwdm *PlayerWeaponDataManager) IfAwakenStar(weaponId int32) bool {
	flag := pwdm.IfIsAwaken(weaponId)
	if !flag {
		return false
	}
	flag = pwdm.IfWeaponExist(weaponId)
	if !flag {
		return false
	}
	obj := pwdm.GetWeapon(weaponId)
	if obj.State == weapontypes.WeaponAwakenStatusTypeOk {
		return false
	}
	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if obj.Level < to.NeedStar {
		return false
	}
	return true
}

//兵魂觉醒
func (pwdm *PlayerWeaponDataManager) Awaken(weaponId int32) bool {
	obj := pwdm.GetWeapon(weaponId)
	if obj == nil {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	obj.State = weapontypes.WeaponAwakenStatusTypeOk
	obj.UpdateTime = now
	obj.SetModified()
	gameevent.Emit(weaponeventtypes.EventTypeWeaponChanged, pwdm.p, nil)
	gameevent.Emit(weaponeventtypes.EventTypeWeaponAwaken, pwdm.p, weaponId)
	return true
}

//兵魂穿戴
func (pwdm *PlayerWeaponDataManager) Wear(weaponId int32) bool {
	flag := pwdm.IsValid(weaponId)
	if !flag {
		return false
	}
	flag = pwdm.IfWeaponExist(weaponId)
	if !flag {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerWeaponInfoObject.WeaponWear = weaponId
	pwdm.playerWeaponInfoObject.UpdateTime = now
	pwdm.playerWeaponInfoObject.SetModified()

	gameevent.Emit(weaponeventtypes.EventTypeWeaponChanged, pwdm.p, nil)
	return true
}

//获取出生武器
func (pwdm *PlayerWeaponDataManager) getBornWeapon() int32 {
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(pwdm.p.GetRole(), pwdm.p.GetSex())
	return playerCreateTemplate.WeaponId
}

//兵魂卸下
func (pwdm *PlayerWeaponDataManager) Unload() {
	bornWeapon := pwdm.getBornWeapon()
	if pwdm.playerWeaponInfoObject.WeaponWear == bornWeapon {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerWeaponInfoObject.WeaponWear = bornWeapon
	pwdm.playerWeaponInfoObject.UpdateTime = now
	pwdm.playerWeaponInfoObject.SetModified()

	gameevent.Emit(weaponeventtypes.EventTypeWeaponChanged, pwdm.p, nil)
	return
}

//兵魂战斗力
func (pwdm *PlayerWeaponDataManager) WeaponPower(power int64) {
	if power <= 0 {
		return
	}
	if pwdm.playerWeaponInfoObject.Power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pwdm.playerWeaponInfoObject.Power = power
	pwdm.playerWeaponInfoObject.UpdateTime = now
	pwdm.playerWeaponInfoObject.SetModified()
	return
}

//兵魂信息
func (pwdm *PlayerWeaponDataManager) ToAllWeaponInfo() (allWeaponInfo *weapontypes.AllWeaponInfo) {
	allWeaponInfo = &weapontypes.AllWeaponInfo{}
	allWeaponInfo.Wear = pwdm.playerWeaponInfoObject.WeaponWear
	for _, tempWeapon := range pwdm.playerWeaponObjectMap {
		weaponInfo := &weapontypes.WeaponInfo{
			WeaponId: tempWeapon.WeaponId,
			Level:    tempWeapon.Level,
			CulLevel: tempWeapon.CulLevel,
			CulPro:   tempWeapon.CulPro,
			State:    int32(tempWeapon.State),
		}
		allWeaponInfo.WeaponList = append(allWeaponInfo.WeaponList, weaponInfo)
	}
	return
}

//仅gm使用 兵魂激活
func (pwdm *PlayerWeaponDataManager) GmWeaponActive(weaponId int32) {
	flag := pwdm.IsValid(weaponId)
	if !flag {
		return
	}
	flag = pwdm.IfWeaponExist(weaponId)
	if flag {
		return
	}
	id, err := idutil.GetId()
	if err != nil {
		return
	}

	to := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	now := global.GetGame().GetTimeService().Now()
	pwo := NewPlayerWeaponObject(pwdm.p)
	pwo.Id = id
	pwo.PlayerId = pwdm.p.GetId()
	pwo.WeaponId = weaponId
	pwo.Level = int32(0)
	pwo.CulLevel = int32(0)
	pwo.CulNum = 0
	pwo.CulPro = 0
	pwo.State = weapontypes.WeaponAwakenStatusTypeNo
	if to.IsAwaken == int32(weapontypes.WeaponAwakenTypeOk) &&
		to.NeedStar == 0 {
		pwo.State = weapontypes.WeaponAwakenStatusTypeOk
	}
	pwo.CreateTime = now
	pwo.SetModified()
	pwdm.playerWeaponObjectMap[weaponId] = pwo
	gameevent.Emit(weaponeventtypes.EventTypeWeaponActivate, pwdm.p, weaponId)
}

//仅gm使用 兵魂培养
func (pwdm *PlayerWeaponDataManager) GmWeaponPeiYang(weaponId int32, level int32) {
	weaponInfo := pwdm.GetWeapon(weaponId)
	if weaponInfo == nil {
		return
	}
	weaponInfo.CulLevel = level
	weaponInfo.CulNum = 0
	weaponInfo.CulPro = 0

	now := global.GetGame().GetTimeService().Now()
	weaponInfo.UpdateTime = now
	weaponInfo.SetModified()
}

//仅gm使用 兵魂升星
func (pwdm *PlayerWeaponDataManager) GmWeaponUpstar(weaponId int32, upStarLevel int32) {
	flag := pwdm.IfCanUpStar(weaponId)
	if !flag {
		return
	}
	obj := pwdm.GetWeapon(weaponId)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	oldLevel := obj.Level
	obj.Level = upStarLevel
	obj.UpNum = 0
	obj.UpPro = 0
	pwdm.playerWeaponInfoObject.Star += (upStarLevel - oldLevel)
	pwdm.playerWeaponInfoObject.SetModified()
	obj.UpdateTime = now
	obj.SetModified()
}

func CreatePlayerWeaponDataManager(p player.Player) player.PlayerDataManager {
	pwdm := &PlayerWeaponDataManager{}
	pwdm.p = p
	return pwdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerWeaponDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerWeaponDataManager))
}
