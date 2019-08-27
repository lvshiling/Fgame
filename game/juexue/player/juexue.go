package player

import (
	"fgame/fgame/core/storage"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/juexue/dao"
	jxentity "fgame/fgame/game/juexue/entity"
	juexueeventtypes "fgame/fgame/game/juexue/event/types"
	"fgame/fgame/game/juexue/juexue"
	jxtypes "fgame/fgame/game/juexue/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"github.com/pkg/errors"
)

//绝学对象
type PlayerJueXueObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Type       jxtypes.JueXueType
	Level      int32
	Insight    jxtypes.JueXueStageType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerJueXueObject(pl player.Player) *PlayerJueXueObject {
	pso := &PlayerJueXueObject{
		player: pl,
	}
	return pso
}

func (pjxo *PlayerJueXueObject) GetPlayerId() int64 {
	return pjxo.PlayerId
}

func (pjxo *PlayerJueXueObject) GetDBId() int64 {
	return pjxo.Id
}

func (pjxo *PlayerJueXueObject) ToEntity() (e storage.Entity, err error) {
	e = &jxentity.PlayerJueXueEntity{
		Id:         pjxo.Id,
		PlayerId:   pjxo.PlayerId,
		Type:       int32(pjxo.Type),
		Level:      pjxo.Level,
		Insight:    int32(pjxo.Insight),
		UpdateTime: pjxo.UpdateTime,
		CreateTime: pjxo.CreateTime,
		DeleteTime: pjxo.DeleteTime,
	}
	return e, err
}

func (pjxo *PlayerJueXueObject) FromEntity(e storage.Entity) error {
	pjxe, _ := e.(*jxentity.PlayerJueXueEntity)
	pjxo.Id = pjxe.Id
	pjxo.PlayerId = pjxe.PlayerId
	pjxo.Type = jxtypes.JueXueType(pjxe.Type)
	pjxo.Level = pjxe.Level
	pjxo.Insight = jxtypes.JueXueStageType(pjxe.Insight)
	pjxo.UpdateTime = pjxe.UpdateTime
	pjxo.CreateTime = pjxe.CreateTime
	pjxo.DeleteTime = pjxe.DeleteTime
	return nil
}

func (pjxo *PlayerJueXueObject) SetModified() {
	e, err := pjxo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "JueXue"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pjxo.player.AddChangedObject(obj)
	return
}

//绝学使用对象
type PlayerJueXueUseObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	Type       jxtypes.JueXueType
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerJueXueUseObject(pl player.Player) *PlayerJueXueUseObject {
	pso := &PlayerJueXueUseObject{
		player: pl,
	}
	return pso
}

func (pjxuo *PlayerJueXueUseObject) GetPlayerId() int64 {
	return pjxuo.PlayerId
}

func (pjxuo *PlayerJueXueUseObject) GetDBId() int64 {
	return pjxuo.Id
}

func (pjxuo *PlayerJueXueUseObject) ToEntity() (e storage.Entity, err error) {
	e = &jxentity.PlayerJueXueUseEntity{
		Id:         pjxuo.Id,
		PlayerId:   pjxuo.PlayerId,
		Type:       int32(pjxuo.Type),
		UpdateTime: pjxuo.UpdateTime,
		CreateTime: pjxuo.CreateTime,
		DeleteTime: pjxuo.DeleteTime,
	}
	return e, err
}

func (pjxuo *PlayerJueXueUseObject) FromEntity(e storage.Entity) error {
	pjxue, _ := e.(*jxentity.PlayerJueXueUseEntity)
	pjxuo.Id = pjxue.Id
	pjxuo.PlayerId = pjxue.PlayerId
	pjxuo.Type = jxtypes.JueXueType(pjxue.Type)
	pjxuo.UpdateTime = pjxue.UpdateTime
	pjxuo.CreateTime = pjxue.CreateTime
	pjxuo.DeleteTime = pjxue.DeleteTime
	return nil
}

func (pjxuo *PlayerJueXueUseObject) SetModified() {
	e, err := pjxuo.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "JueXueUse"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pjxuo.player.AddChangedObject(obj)
	return
}

//玩家绝学管理器
type PlayerJueXueDataManager struct {
	p player.Player
	//玩家绝学使用
	playerJueXueUseObject *PlayerJueXueUseObject
	//玩家绝学map
	playerJueXueObjectMap map[jxtypes.JueXueType]*PlayerJueXueObject
}

func (pjxdm *PlayerJueXueDataManager) Player() player.Player {
	return pjxdm.p
}

//加载
func (pjxdm *PlayerJueXueDataManager) Load() (err error) {
	pjxdm.playerJueXueObjectMap = make(map[jxtypes.JueXueType]*PlayerJueXueObject)
	//加载玩家绝学
	jueXues, err := dao.GetJueXueDao().GetJueXueList(pjxdm.p.GetId())
	if err != nil {
		return
	}
	//绝学信息
	for _, juexue := range jueXues {
		pjxo := NewPlayerJueXueObject(pjxdm.p)
		pjxo.FromEntity(juexue)
		pjxdm.playerJueXueObjectMap[pjxo.Type] = pjxo
	}

	//加载玩家绝学使用信息
	jueXueUseEntity, err := dao.GetJueXueDao().GetJueXueUseEntity(pjxdm.p.GetId())
	if err != nil {
		return
	}
	if jueXueUseEntity == nil {
		pjxdm.initPlayerJueXueUseObject()
	} else {
		pjxdm.playerJueXueUseObject = NewPlayerJueXueUseObject(pjxdm.p)
		pjxdm.playerJueXueUseObject.FromEntity(jueXueUseEntity)
	}
	return nil
}

//第一次初始化
func (pjxdm *PlayerJueXueDataManager) initPlayerJueXueUseObject() {
	// TODO 初始化出生自带绝学
	pjxuo := NewPlayerJueXueUseObject(pjxdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pjxuo.Id = id
	//生成id
	pjxuo.PlayerId = pjxdm.p.GetId()
	pjxuo.Type = 0
	pjxuo.CreateTime = now
	pjxdm.playerJueXueUseObject = pjxuo
	pjxuo.SetModified()
}

//加载后
func (pjxdm *PlayerJueXueDataManager) AfterLoad() (err error) {
	pjxdm.resetMaxLevel()
	return nil
}

func (pjxdm *PlayerJueXueDataManager) resetMaxLevel() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for jueXueType, jueXueObj := range pjxdm.playerJueXueObjectMap {
		jueXueTemplate := juexue.GetJueXueService().GetJueXueMaxLevel(jueXueType, jueXueObj.Insight)
		if jueXueObj.Level > jueXueTemplate.Level {
			jueXueObj.Level = jueXueTemplate.Level
			jueXueObj.UpdateTime = now
			jueXueObj.SetModified()
		}
	}
	return
}

//心跳
func (pjxdm *PlayerJueXueDataManager) Heartbeat() {

}

//获取玩家使用绝学
func (pjxdm *PlayerJueXueDataManager) GetJueXueUseId() int32 {
	typ := pjxdm.playerJueXueUseObject.Type
	stage, flag := pjxdm.getJueXueStageByTyp(typ)
	if !flag {
		return 0
	}
	_, level := pjxdm.GetJueXueLevelByTyp(typ)
	to := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, stage, level)
	return int32(to.TemplateId())
}

//使用绝学类型
func (pjxdm *PlayerJueXueDataManager) GetJueXueUseTyp() jxtypes.JueXueType {
	return pjxdm.playerJueXueUseObject.Type
}

//获取玩家绝学列表
func (pjxdm *PlayerJueXueDataManager) GetJueXueMap() map[jxtypes.JueXueType]*PlayerJueXueObject {
	return pjxdm.playerJueXueObjectMap
}

//获取绝学通过绝学类型
func (pjxdm *PlayerJueXueDataManager) GetJueXueByTyp(typ jxtypes.JueXueType) *PlayerJueXueObject {
	jueXueObj, exist := pjxdm.playerJueXueObjectMap[typ]
	if !exist {
		return nil
	}
	return jueXueObj
}

//获取绝学状态通过绝学类型
func (pjxdm *PlayerJueXueDataManager) getJueXueStageByTyp(typ jxtypes.JueXueType) (jxtypes.JueXueStageType, bool) {
	obj := pjxdm.GetJueXueByTyp(typ)
	if obj == nil {
		return 0, false
	}
	return obj.Insight, true
}

//获取等级通过绝学类型
func (pjxdm *PlayerJueXueDataManager) GetJueXueLevelByTyp(typ jxtypes.JueXueType) (jxtypes.JueXueStageType, int32) {
	obj := pjxdm.GetJueXueByTyp(typ)
	if obj == nil {
		return 0, 0
	}
	return obj.Insight, obj.Level
}

//是否已激活
func (pjxdm *PlayerJueXueDataManager) IfJueXueExist(typ jxtypes.JueXueType) bool {
	obj := pjxdm.GetJueXueByTyp(typ)
	if obj == nil {
		return false
	}
	return true
}

//是否已使用
func (pjxdm *PlayerJueXueDataManager) IfUseExist(typ jxtypes.JueXueType) bool {
	curUseTyp := pjxdm.GetJueXueUseTyp()
	return curUseTyp == typ
}

//是否达到满级
func (pjxdm *PlayerJueXueDataManager) ifFullLevel(typ jxtypes.JueXueType) bool {
	stage, level := pjxdm.GetJueXueLevelByTyp(typ)
	to := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, stage, level)
	if to.NextId == 0 {
		return true
	}
	return false
}

//能否升级
func (pjxdm *PlayerJueXueDataManager) IfCanUpgrade(typ jxtypes.JueXueType) bool {
	flag := pjxdm.IfJueXueExist(typ)
	if !flag {
		return false
	}
	flag = pjxdm.ifFullLevel(typ)
	if flag {
		return false
	}
	return true
}

//能够顿悟
func (pjxdm *PlayerJueXueDataManager) IfCanInsight(typ jxtypes.JueXueType) bool {
	flag := pjxdm.IfJueXueExist(typ)
	if !flag {
		return false
	}
	state, level := pjxdm.GetJueXueLevelByTyp(typ)
	maxLevelTemp := juexue.GetJueXueService().GetJueXueMaxLevel(typ, state)
	maxLevel := maxLevelTemp.Level

	if level < maxLevel && state == jxtypes.JueXueStageTypeAorU {
		return false
	}

	if state == jxtypes.JueXueStageTypeInsight {
		to := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, jxtypes.JueXueStageTypeInsight, level+1)
		if to == nil {
			return false
		}
	}

	return true
}

//绝学激活
func (pjxdm *PlayerJueXueDataManager) JueXueActive(typ jxtypes.JueXueType) bool {
	if !typ.Valid() {
		return false
	}

	flag := pjxdm.IfJueXueExist(typ)
	if flag {
		return false
	}
	id, err := idutil.GetId()
	if err != nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	pjxo := NewPlayerJueXueObject(pjxdm.p)
	pjxo.Id = id
	pjxo.PlayerId = pjxdm.p.GetId()
	pjxo.Type = typ
	pjxo.Level = int32(1)
	pjxo.Insight = jxtypes.JueXueStageTypeAorU
	pjxo.CreateTime = now
	pjxo.SetModified()

	pjxdm.playerJueXueObjectMap[typ] = pjxo

	gameevent.Emit(juexueeventtypes.EventTypeJueXueAcitivate, pjxdm.p, typ)
	return true
}

//绝学升级
func (pjxdm *PlayerJueXueDataManager) Upgrade(typ jxtypes.JueXueType) bool {
	flag := pjxdm.IfCanUpgrade(typ)
	if !flag {
		return false
	}

	obj := pjxdm.GetJueXueByTyp(typ)
	now := global.GetGame().GetTimeService().Now()
	obj.UpdateTime = now
	obj.Level += 1
	obj.SetModified()

	gameevent.Emit(juexueeventtypes.EventTypeJueXueUpgrade, pjxdm.p, typ)
	return true
}

//能否使用
func (pjxdm *PlayerJueXueDataManager) IfCanUse(typ jxtypes.JueXueType) bool {
	flag := pjxdm.IfJueXueExist(typ)
	if !flag {
		return false
	}
	flag = pjxdm.IfUseExist(typ)
	if flag {
		return false
	}
	return true
}

//绝学顿悟
func (pjxdm *PlayerJueXueDataManager) Insight(typ jxtypes.JueXueType) bool {
	flag := pjxdm.IfCanInsight(typ)
	if !flag {
		return false
	}

	oldSkillId := int32(0)
	curTyp := pjxdm.GetJueXueUseTyp()
	jueXueObj := pjxdm.GetJueXueByTyp(curTyp)
	if jueXueObj != nil {
		oldSkillId = juexue.GetJueXueService().GetSkillId(jueXueObj.Insight, curTyp, jueXueObj.Level)
	}

	obj := pjxdm.GetJueXueByTyp(typ)
	now := global.GetGame().GetTimeService().Now()
	obj.UpdateTime = now
	if obj.Insight == jxtypes.JueXueStageTypeAorU {
		obj.Insight = jxtypes.JueXueStageTypeInsight
		obj.Level = 1
	} else {
		obj.Level++
	}
	obj.SetModified()
	eventData := juexueeventtypes.CreateJueXueInsightEventData(typ, oldSkillId)
	gameevent.Emit(juexueeventtypes.EventTypeJueXueInsight, pjxdm.p, eventData)
	return true
}

//绝学使用
func (pjxdm *PlayerJueXueDataManager) JueXueUse(typ jxtypes.JueXueType) bool {
	flag := pjxdm.IfCanUse(typ)
	if !flag {
		return false
	}

	oldSkillId := int32(0)
	//获取老的技能
	curTyp := pjxdm.GetJueXueUseTyp()
	juexueObj := pjxdm.GetJueXueByTyp(curTyp)
	if juexueObj != nil {
		oldSkillId = juexue.GetJueXueService().GetSkillId(juexueObj.Insight, curTyp, juexueObj.Level)
	}

	now := global.GetGame().GetTimeService().Now()
	pjxdm.playerJueXueUseObject.Type = typ
	pjxdm.playerJueXueUseObject.UpdateTime = now
	pjxdm.playerJueXueUseObject.SetModified()

	gameevent.Emit(juexueeventtypes.EventTypeJueXueUse, pjxdm.p, oldSkillId)
	return true
}

//绝学卸下
func (pjxdm *PlayerJueXueDataManager) Unload() bool {
	curTyp := pjxdm.GetJueXueUseTyp()
	juexueObj := pjxdm.GetJueXueByTyp(curTyp)
	oldSkillId := juexue.GetJueXueService().GetSkillId(juexueObj.Insight, curTyp, juexueObj.Level)

	now := global.GetGame().GetTimeService().Now()
	pjxdm.playerJueXueUseObject.Type = 0
	pjxdm.playerJueXueUseObject.UpdateTime = now
	pjxdm.playerJueXueUseObject.SetModified()

	gameevent.Emit(juexueeventtypes.EventTypeJueXueUnload, pjxdm.p, oldSkillId)

	return true
}

func CreatePlayerJueXueDataManager(p player.Player) player.PlayerDataManager {
	pjxdm := &PlayerJueXueDataManager{}
	pjxdm.p = p
	return pjxdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerJueXueDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerJueXueDataManager))
}
