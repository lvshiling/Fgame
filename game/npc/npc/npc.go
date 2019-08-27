package npc

import (
	"fgame/fgame/core/heartbeat"
	coretypes "fgame/fgame/core/types"
	"runtime/debug"

	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/buff/buff"
	buffcommon "fgame/fgame/game/buff/common"
	bufftemplate "fgame/fgame/game/buff/template"
	cdcommon "fgame/fgame/game/cd/common"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/global"
	npceventtypes "fgame/fgame/game/npc/event/types"
	propertytypes "fgame/fgame/game/property/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skillcommon "fgame/fgame/game/skill/common"
	"fgame/fgame/game/skill/skill"
	skilltemplate "fgame/fgame/game/skill/template"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//npc
type NPCBase struct {
	n scene.NPC
	//id
	id int64
	//场景的索引
	idInScene  int32
	createTime int64
	//出生地点
	bornPos coretypes.Position
	//出生角度
	bornAngle float64
	//生物模板
	biologyTemplate *gametemplate.BiologyTemplate
	//当前攻击目标
	attackTarget *scene.Enemy

	//主人类型
	ownerType scenetypes.OwnerType
	//主人id
	ownerId int64
	//主人仙盟
	ownerAllianceId int64
	//cd组管理器
	cdGroupManager *cdcommon.CDGroupManager
	//场景对象
	*scene.SceneObjectBase
	//buff管理器
	*buff.BuffDataManager
	//技能管理器
	*skill.SkillManager
	*battle.TeShuSkillManager
	//系统属性管理器
	*battle.SystemPropertyManager
	//状态管理器
	*NPCStateManager
	//战斗属性管理器
	*battle.BattlePropertyManager
	//基础战斗管理器
	*battle.BattleManager
	//状态数据管理器
	*battle.StateDataManager
	//移动动作
	*battle.MoveAction
	//上次心跳时间
	lastHeartBeatTime int64
	//死亡
	isDead bool
	//死亡时间
	deadTime int64
	//定时器
	hbRunner heartbeat.HeartbeatTaskRunner
}

func (n *NPCBase) GetId() int64 {
	return n.id
}

func (n *NPCBase) GetIdInScene() int32 {
	return n.idInScene
}

func (n *NPCBase) GetBornPosition() coretypes.Position {
	return n.bornPos
}

func (n *NPCBase) GetBornAngle() float64 {
	return n.bornAngle
}

func (n *NPCBase) GetCDGroupManager() *cdcommon.CDGroupManager {
	return n.cdGroupManager
}

func (n *NPCBase) GetTempId() int {
	return n.biologyTemplate.Id
}

func (n *NPCBase) GetBiologyTemplate() *gametemplate.BiologyTemplate {
	return n.biologyTemplate
}

//重载
func (n *NPCBase) UseSkill(skillId int32) bool {
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		return false
	}
	//判断cd
	if n.SkillManager.IsSkillInCd(skillTemplate.TypeId) {
		return false
	}

	//技能动作
	if skillTemplate.IsPositive() {
		flag := n.NPCStateManager.Attack()
		if !flag {
			return false
		}
		//设置技能时间
		n.SetSkillActionTime(int64(skillTemplate.ActionTime))
	}

	//使用技能
	flag := n.SkillManager.UseSkill(skillTemplate.TypeId)
	if !flag {
		panic(fmt.Errorf("npc:使用技能应该成功"))
	}

	return true
}

//被攻击时位移
func (n *NPCBase) AttackedMove(pos coretypes.Position, angle float64, moveSpeed float64, stopTime float64) {
	flag := n.NPCStateManager.AttackedMove()
	if !flag {
		return
	}
	n.SkilledStop(pos, int64(stopTime*float64(common.SECOND)), moveSpeed)
	return
}

//更新战斗属性
func (n *NPCBase) UpdateBattleProperty(mask uint32) {
	//不需要更新
}

//心跳
func (n *NPCBase) Heartbeat() {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"error": terr,
					"stack": string(debug.Stack()),
				}).Error("npc:Heartbeat,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
	}()
	n.hbRunner.Heartbeat()
	n.BuffDataManager.Heartbeat()
	n.SkillManager.Heartbeat()
	n.MoveAction.Heartbeat()
	n.NPCStateManager.Heartbeat()
	n.lastHeartBeatTime = global.GetGame().GetTimeService().Now()
	return
}

//空闲
func (n *NPCBase) Idle() bool {
	flag := n.NPCStateManager.Idle()
	if !flag {
		return false
	}
	n.clear()
	return true
}

func (n *NPCBase) Trace() bool {
	flag := n.NPCStateManager.Trace()
	if !flag {
		return false
	}
	return true
}

//复活
func (n *NPCBase) Reborn(rebornPos coretypes.Position) {
	flag := n.NPCStateManager.Idle()
	if !flag {
		return
	}

	n.ownerType = scenetypes.OwnerTypeNone
	n.SetFactionType(n.biologyTemplate.GetFactionType())
	n.ownerId = 0
	n.ownerAllianceId = 0
	n.SceneObjectBase.ResetNeighbor()
	//重新设置主人
	n.isDead = false

	n.SetAngle(n.bornAngle)
	//重新回复血量
	n.BattlePropertyManager.Reborn()
	//发送事件
	gameevent.Emit(sceneeventtypes.EventTypeBattleObjectReborn, n.n, rebornPos)

	n.clear()

}

func (n *NPCBase) AddHP(hp int64) int64 {
	oldHp := n.GetHP()
	add := n.BattlePropertyManager.AddHP(hp)
	newHp := n.GetHP()
	eventData := CreateNPCHPChangedEventData(oldHp, newHp, 0)
	gameevent.Emit(npceventtypes.EventTypeNPCHPChanged, n.n, eventData)
	return add
}

func (n *NPCBase) Recycle(playerId int64) bool {
	flag := n.CostHP(n.GetHP(), playerId)
	if !flag {
		return false
	}
	return n.Dead(playerId)
}

//重载
//返回是否死亡
func (n *NPCBase) CostHP(hp int64, attackId int64) bool {
	if hp <= 0 {
		return false
	}

	oldHp := n.GetHP()
	//扣血
	dead := n.BattlePropertyManager.CostHP(hp, attackId)
	currentHp := n.GetHP()
	eventData := CreateNPCHPChangedEventData(oldHp, currentHp, attackId)
	if n.GetBiologyTemplate().GetDropType() == scenetypes.DropTypePercent {
		maxHp := n.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
		//算出万分比
		oldPercent := float64(oldHp) / float64(maxHp) * common.MAX_RATE
		currentPercent := float64(currentHp) / float64(maxHp) * common.MAX_RATE
		dropPercent := n.GetBiologyTemplate().DropFlag
		for float64(dropPercent) < currentPercent {
			dropPercent += n.GetBiologyTemplate().DropFlag
		}
		if float64(dropPercent) < oldPercent {
			num := int32(math.Floor(oldPercent-float64(dropPercent)))/n.GetBiologyTemplate().DropFlag + 1
			//掉落
			eventData := sceneeventtypes.CreateMonsterDropData(attackId, num)
			//掉落
			gameevent.Emit(sceneeventtypes.EventTypeMonsterDrop, n.n, eventData)
		}
	}
	gameevent.Emit(npceventtypes.EventTypeNPCHPChanged, n.n, eventData)
	s := n.n.GetScene()
	if s != nil {
		//在场景内
		if s.GetSceneObject(attackId) != nil {
			//添加伤害
			n.n.AddDamage(attackId, oldHp-currentHp)
		}
	}
	if dead {
		flag := n.NPCStateManager.Dead()
		if !flag {
			panic(fmt.Errorf("npc：状态改变应该成功"))
		}
		n.ClearAllSkillAction()
		n.deadTime = global.GetGame().GetTimeService().Now()
		n.isDead = true

		return true
	}

	return dead
}

func (n *NPCBase) Dead(attackId int64) bool {
	if !n.isDead {
		return false
	}

	gameevent.Emit(sceneeventtypes.EventTypeBattleObjectDead, n.n, attackId)

	//状态改变
	if !n.biologyTemplate.CanReborn() {
		n.GetScene().RemoveSceneObject(n.n, true)
	}
	return true
}

func (n *NPCBase) DeadInTime(deadTime int64) bool {
	oldHp := n.GetHP()
	//扣血
	dead := n.BattlePropertyManager.CostHP(oldHp, 0)

	if dead {
		flag := n.NPCStateManager.Dead()
		if !flag {
			panic(fmt.Errorf("npc：状态改变应该成功"))
		}
		n.deadTime = deadTime
		n.isDead = true
		return true
	}

	return true
}

//是否死亡
func (n *NPCBase) IsDead() bool {
	return n.isDead
}

//是否死亡
func (n *NPCBase) GetDeadTime() int64 {
	return n.deadTime
}

//是否死亡
func (n *NPCBase) GetName() string {
	return n.biologyTemplate.Name
}

func (n *NPCBase) GetExtraSpeed() int64 {
	return 0
}

func (n *NPCBase) ExitScene(active bool) {
	n.n.ClearAllSkillAction()
	//暂停移动
	n.n.PauseMove()
	n.SceneObjectBase.ExitScene(active)
}

func (n *NPCBase) ShouldRemove() bool {
	if n.biologyTemplate.XiaoshiTime == 0 {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	if now-n.createTime > int64(n.biologyTemplate.XiaoshiTime) {
		return true
	}
	return false
}

func NewNPCBase(np scene.NPC, ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64) *NPCBase {
	n := &NPCBase{}
	n.n = np
	if id == 0 {
		id, _ = idutil.GetId()
	}
	now := global.GetGame().GetTimeService().Now()
	n.createTime = now
	n.id = id
	n.idInScene = idInScene
	n.bornPos = pos
	n.bornAngle = angle
	n.biologyTemplate = biologyTemplate
	n.cdGroupManager = cdcommon.NewCDGroupManager()

	buffList := make([]buffcommon.BuffObject, 0, len(biologyTemplate.GetBuffIdList()))
	for _, buffId := range biologyTemplate.GetBuffIdList() {
		buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
		if buffTemplate == nil {
			continue
		}
		buffObj := buffcommon.NewBuffObject(0, buffId, buffTemplate.Group, now, 0, 1, 0, buffTemplate.TimeDuration, nil)
		buffList = append(buffList, buffObj)
	}
	n.BuffDataManager = buff.CreateBuffDataManagerWithBuffList(np, buffList)

	n.SystemPropertyManager = battle.CreateSystemPropertyManagerWithData(np, biologyTemplate.GetBattlePropertyMap())

	n.lastHeartBeatTime = global.GetGame().GetTimeService().Now()

	skillList := make([]skillcommon.SkillObject, 0, 4)
	for _, skillId := range n.biologyTemplate.GetAllSkills() {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(int32(skillId))
		so := skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil)
		skillList = append(skillList, so)
	}

	n.SkillManager = skill.CreateSkillManager(np, n.cdGroupManager, skillList)
	n.NPCStateManager = NewNPCStateManager(np, n.biologyTemplate.GetBiologyScriptType())

	n.ownerId = ownerId
	n.ownerAllianceId = ownerAllianceId
	n.ownerType = ownerType
	factionType := biologyTemplate.GetFactionType()
	biologyType := biologyTemplate.GetBiologyType()
	n.BattleManager = battle.CreateBattleManager(np, factionType)
	n.SceneObjectBase = scene.NewSceneObjectBase(np, pos, angle, biologyType)
	n.StateDataManager = battle.CreateStateDateManager(np)
	hp := n.GetSystemBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	tp := n.GetSystemBattleProperty(propertytypes.BattlePropertyTypeMaxTP)
	n.BattlePropertyManager = battle.CreateBattlePropertyManager(np, hp, tp)
	n.MoveAction = battle.CreateMoveAction(np)
	n.TeShuSkillManager = battle.CreateTeShuSkillManager(np, nil)
	n.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	//恢复任务
	recoverTask := CreateNPCRecoverTask(np)
	n.hbRunner.AddTask(recoverTask)

	return n
}

type npc struct {
	*NPCBase
}

func CreateNPC(ownerType scenetypes.OwnerType, ownerId int64, owberAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) scene.NPC {
	n := &npc{}
	b := NewNPCBase(n, ownerType, ownerId, owberAllianceId, id, idInScene, biologyTemplate, pos, angle)
	n.NPCBase = b
	n.Calculate()
	if deadTime != 0 {
		n.DeadInTime(deadTime)
	}
	return n
}
