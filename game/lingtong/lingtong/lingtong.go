package lingtong

import (
	"fgame/fgame/core/heartbeat"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/buff/buff"
	cdcommon "fgame/fgame/game/cd/common"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skillcommon "fgame/fgame/game/skill/common"
	"fgame/fgame/game/skill/skill"
	skilltemplate "fgame/fgame/game/skill/template"
	gametemplate "fgame/fgame/game/template"
	"runtime/debug"

	log "github.com/Sirupsen/logrus"
)

type LingTong struct {
	id               int64
	owner            scene.Player
	lingTongTemplate *gametemplate.LingTongTemplate
	*LingTongShowManager
	//cd组管理器
	cdGroupManager *cdcommon.CDGroupManager
	//场景对象
	*LingTongSceneManager
	//buff管理器
	*buff.BuffDataManager
	//技能管理器
	*skill.SkillManager
	*battle.TeShuSkillManager
	//系统属性管理器
	*battle.SystemPropertyManager
	//系统属性管理器
	*battle.StateDataManager
	//战斗属性管理器
	*battle.BattlePropertyManager
	//基础战斗管理器
	*battle.BattleManager
	*battle.MoveAction
	*LingTongStateManager
	//定时器
	hbRunner heartbeat.HeartbeatTaskRunner
	name     string
}

func (l *LingTong) GetId() int64 {
	return l.id
}

func (l *LingTong) GetName() string {
	return l.name
}

func (l *LingTong) UpdateName(name string) {
	l.name = name
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongRename, l, nil)
}

func (l *LingTong) GetOwner() scene.Player {
	return l.owner
}

func (l *LingTong) GetSceneObjectSetType() scenetypes.BiologySetType {
	return scenetypes.BiologySetTypePlayer
}

func (l *LingTong) GetCDGroupManager() *cdcommon.CDGroupManager {
	return l.cdGroupManager
}

func (l *LingTong) Heartbeat() {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"error": terr,
					"stack": string(debug.Stack()),
				}).Error("lingtong:Heartbeat,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
	}()
	l.BuffDataManager.Heartbeat()
	l.SkillManager.Heartbeat()
	l.MoveAction.Heartbeat()
	l.hbRunner.Heartbeat()
}

func (l *LingTong) IsDead() bool {
	return false
}

//是否是敌人
func (l *LingTong) IsEnemy(bo scene.BattleObject) bool {
	if l.owner == nil {
		return false
	}
	return l.owner.IsEnemy(bo)
}

//对象移动
func (l *LingTong) OnMove(bo scene.BattleObject, pos coretypes.Position, angle float64) {

}

//死亡
func (l *LingTong) OnDead(bo scene.BattleObject) {

}

//重生
func (l *LingTong) OnReborn(bo scene.BattleObject) {

}

//重生
func (l *LingTong) Reborn(pos coretypes.Position) {

}
func (l *LingTong) GetLingTongTemplate() *gametemplate.LingTongTemplate {
	return l.lingTongTemplate
}

func (l *LingTong) GetLingTongId() int32 {
	return int32(l.lingTongTemplate.TemplateId())
}

func (l *LingTong) UpdateLingTongTemplate(lingTongTemplate *gametemplate.LingTongTemplate, name string) {
	if lingTongTemplate.TemplateId() == l.lingTongTemplate.TemplateId() {
		return
	}
	l.name = name
	l.lingTongTemplate = lingTongTemplate
	//技能列表
	skillList := make([]skillcommon.SkillObject, 0, 2)
	for _, tempSkillId := range lingTongTemplate.GetSkillIdList() {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(tempSkillId)
		skillObj := skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil)
		skillList = append(skillList, skillObj)
	}

	l.SkillManager = skill.CreateSkillManager(l, l.cdGroupManager, skillList)
	gameevent.Emit(lingtongeventtypes.EventTypeBattleLingTongChanged, l, nil)

}

//使用技能
func (l *LingTong) UseSkill(skillId int32) bool {
	flag := l.SkillManager.UseSkill(skillId)
	if !flag {
		return false
	}

	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)

	//技能动作
	if skillTemplate.IsPositive() {
		l.SetSkillActionTime(int64(skillTemplate.ActionTime))
		l.Attack()
	}
	return true
}

//创建灵童
func CreateLingTong(owner scene.Player,
	id int64,
	name string,
	pos coretypes.Position,
	angle float64,
	lingTongTemplate *gametemplate.LingTongTemplate,
	showObj *LingTongShowObject,
	battleProperties map[int32]int64) scene.LingTong {
	l := &LingTong{}
	l.id = id
	l.owner = owner
	l.name = name
	l.lingTongTemplate = lingTongTemplate

	l.LingTongShowManager = CreateLingTongShowManagerWithObject(l, showObj)
	l.cdGroupManager = cdcommon.NewCDGroupManager()
	l.BuffDataManager = buff.CreateBuffDataManager(l)

	//技能列表
	skillList := make([]skillcommon.SkillObject, 0, 2)
	for _, tempSkillId := range lingTongTemplate.GetSkillIdList() {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(tempSkillId)
		skillObj := skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil)
		skillList = append(skillList, skillObj)
	}

	l.SkillManager = skill.CreateSkillManager(l, l.cdGroupManager, skillList)

	l.SystemPropertyManager = battle.CreateSystemPropertyManagerWithData(l, battleProperties)
	l.StateDataManager = battle.CreateStateDateManager(l)
	l.BattlePropertyManager = battle.CreateBattlePropertyManager(l, 0, 0)
	l.BattleManager = battle.CreateBattleManager(l, scenetypes.FactionTypePlayer)
	l.LingTongSceneManager = CreateLingTongSceneManager(l, pos, angle)
	l.LingTongStateManager = CreateLingTongStateManager(l)
	l.MoveAction = battle.CreateMoveAction(l)
	l.TeShuSkillManager = battle.CreateTeShuSkillManager(l, nil)
	l.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	lingTongTask := CreateLingTongTask(l)
	l.hbRunner.AddTask(lingTongTask)
	l.Calculate()
	return l
}
