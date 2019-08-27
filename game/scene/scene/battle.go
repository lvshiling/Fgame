package scene

import (
	coretypes "fgame/fgame/core/types"
	buffcommon "fgame/fgame/game/buff/common"
	cdcommon "fgame/fgame/game/cd/common"
	"fgame/fgame/game/common/common"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	skillcommon "fgame/fgame/game/skill/common"
	skilltypes "fgame/fgame/game/skill/types"
	"fmt"
	"math"
)

//通用
type SkillActionManager interface {
	AddSkillAction(skillId int32)
	ClearAllSkillAction()
}

//通用
//buff管理器
type BuffManager interface {
	UpdateBuff(o buffcommon.BuffObject)
	//添加buff
	AddBuff(buffId int32, ownerId int64, times int32, tianFuList []int32) (flag bool)
	//移除buff
	RemoveBuff(buffId int32)
	//更新buff
	RefreshBuff()
	//触发buff
	TouchBuff(buffId int32)
	//获取buff列表
	GetBuffs() map[int32]buffcommon.BuffObject
	//获取buff
	GetBuff(buffId int32) buffcommon.BuffObject
	//战斗限制
	GetBattleLimit() int64
	//扣除buff特殊效果值
	CostEffectNum(effectType scenetypes.BuffEffectType, num int64) int64
	GetEffectNum(effectType scenetypes.BuffEffectType) int64
}

type TeshuSkillObject struct {
	skillId    int32
	chuFaRate  int32
	diKangRate int32
}

func (o *TeshuSkillObject) GetSkillId() int32 {
	return o.skillId
}

func (o *TeshuSkillObject) GetChuFaRate() int32 {
	return o.chuFaRate
}

func (o *TeshuSkillObject) GetDiKangRate() int32 {
	return o.diKangRate
}

func (o *TeshuSkillObject) AddRate(chuFaRate int32, diKangRate int32) {
	o.chuFaRate += chuFaRate
	o.diKangRate += diKangRate
}

func CreateTeshuSkillObject(skillId int32, chuFaRate int32, diKangRate int32) *TeshuSkillObject {
	o := &TeshuSkillObject{
		skillId:    skillId,
		chuFaRate:  chuFaRate,
		diKangRate: diKangRate,
	}
	return o
}

type TeShuSkillManager interface {
	ResetTeShuSkills(oList []*TeshuSkillObject)
	GetTeShuSkills() []*TeshuSkillObject
	GetTeShuSkill(skillId int32) *TeshuSkillObject
}

//通用
//技能管理器
type SkillManager interface {
	//是否在cd中
	IsSkillInCd(skillId int32) bool
	//使用技能
	UseSkill(skillId int32) bool
	//获取所有技能
	GetAllSkills() map[int32]skillcommon.SkillObject
	//获取技能
	GetSkills(skilltypes.SkillSecondType) map[int32]skillcommon.SkillObject
	//获取技能
	GetSkill(skillId int32) skillcommon.SkillObject
	//改变静态技能
	ChangeSkill(oldSkillId int32, newSkillId int32) bool
	//改变动态技能
	ChangeDynamicSkill(skillId int32, level int32)
}

//状态数据管理器
type StateDataManager interface {
	//设置技能动作事件
	SetSkillActionTime(skillActionTime int64)
	//攻击时间
	GetSkillTime() int64
	//受击
	SkilledStop(pos coretypes.Position, skilledStopTime int64, attackedMoveSpeed float64)
	//攻击动作时间
	GetSkillActionTime() int64
	//受击时间
	GetSkilledTime() int64
	//受击停顿时间
	GetSkilledStopTime() int64
	//获取被攻击速度
	GetAttackedMoveSpeed() float64
	GetDestPosition() coretypes.Position
}

//通用
//系统属性
type SystemPropertyManager interface {
	GetSystemBattleProperty(pt propertytypes.BattlePropertyType) int64
	//获取所有系统属性
	GetAllSystemBattleProperties() map[int32]int64
	//更新战斗属性
	UpdateSystemBattleProperty(properties map[int32]int64)
	//获取系统变更的
	GetSystemBattlePropertyChangedTypesAndReset() (battleChanged map[int32]int64)
}

//战斗属性管理器
type BattlePropertyManager interface {
	//获取战斗属性
	GetBattleProperty(propertytypes.BattlePropertyType) int64
	//更新buff属性
	UpdateBuffProperty()
	//重新计算属性
	Calculate()
	//获取当前hp
	GetHP() int64
	//返回是否死亡
	CostHP(hp int64, attackId int64) bool
	Dead(attackId int64) bool
	//加血
	AddHP(hp int64) int64
	//重置血量
	ResetHP()
	//扣除tp
	CostTP(tp int64) bool
	AddTP(tp int64) int64
	GetTP() int64
	//是否死亡
	IsDead() bool
	GetAllBattleProperties() map[int32]int64

	GetMaxHP() int64
	GetMaxTP() int64
	//重生
	Reborn(pos coretypes.Position)
}

//战斗管理器
type BattleManager interface {
	//获取场景对象类型
	GetSceneObjectSetType() scenetypes.BiologySetType
	//获取阵营
	GetFactionType() scenetypes.FactionType
	//设置阵营
	SetFactionType(scenetypes.FactionType)

	//被攻击位移
	AttackedMove(pos coretypes.Position, angle float64, moveSpeed float64, stopTime float64)
	//是否是敌人
	IsEnemy(bo BattleObject) bool
	//添加敌人
	AddEnemy(bo *Enemy) bool
	//移除敌人
	RemoveEnemy(e *Enemy) bool
	//获取敌人
	GetEnemy(bo BattleObject) *Enemy
	//进入战斗
	EnterBattle(bo BattleObject)
	ExitBattle(bo BattleObject)
	ResetEnemy()
	//伤害统计
	AddDamage(attackId int64, damage int64)
	ClearDamage(attackId int64)
	ClearAllDamages()
	GetAllDamages() map[int64]int64
	//对象移动
	OnMove(bo BattleObject, pos coretypes.Position, angle float64)
	//死亡
	OnDead(bo BattleObject)
	//重生
	OnReborn(bo BattleObject)
	FindHatestEnemy() (e *Enemy)
	GetEnemies() map[int64]*Enemy

	SetAttackTarget(bo BattleObject)
	GetAttackTarget() BattleObject
	SetDefaultAttackTarget(bo BattleObject)
	GetDefaultAttackTarget() BattleObject
	SetForeverAttackTarget(bo BattleObject)
	GetForeverAttackTarget() BattleObject
	//添加仇恨值
	AddHate(target BattleObject, hate int)
	//获取仇恨值
	GetHate(target BattleObject) int
}

type MoveManager interface {
	//设置目的地
	SetDestPosition(destPos coretypes.Position) bool
	PauseMove()
	IsMove() bool
}

//TODO 组件化
//战斗对象
type BattleObject interface {
	//获取cd组
	GetCDGroupManager() *cdcommon.CDGroupManager
	GetExtraSpeed() int64
	//场景
	SceneObject
	//buff
	BuffManager
	//技能
	SkillManager
	TeShuSkillManager
	//系统属性管理器
	SystemPropertyManager
	//战斗属性管理器
	BattlePropertyManager
	//战斗管理器
	BattleManager
	//技能动作管理器
	SkillActionManager
	//状态数据管理器
	StateDataManager
	//移动管理器
	MoveManager
}

//获取技能伤害
func GetBasicDamage(attackObj BattleObject, defendObj BattleObject, skillDamageAttack float64, extraDamage float64) float64 {
	return GetDamage(attackObj, 0, defendObj, skillDamageAttack, extraDamage)
}

func GetDamage(attackObj BattleObject, extraAttack int64, defendObj BattleObject, skillDamageAttack float64, extraDamage float64) float64 {
	attack := attackObj.GetBattleProperty(propertytypes.BattlePropertyTypeAttack)
	attack += extraAttack
	defend := defendObj.GetBattleProperty(propertytypes.BattlePropertyTypeDefend)

	defendAttackRatio := float64(defend) / (4 * float64(attack))
	basicDamage := float64(attack) / math.Exp(defendAttackRatio+math.Pow(defendAttackRatio, 1.5)/2)
	basicDamage *= skillDamageAttack
	basicDamage += extraDamage

	//修正伤害加成
	damageAttack := float64(attackObj.GetBattleProperty(propertytypes.BattlePropertyTypeDamageAddPercent)) / common.MAX_RATE
	damageDefend := float64(defendObj.GetBattleProperty(propertytypes.BattlePropertyTypeDamageDefendPercent)) / common.MAX_RATE
	damage := basicDamage * (1 + damageAttack) * (1 - damageDefend)

	damageAttackValue := float64(attackObj.GetBattleProperty(propertytypes.BattlePropertyTypeDamageAdd))
	damageDefendValue := float64(defendObj.GetBattleProperty(propertytypes.BattlePropertyTypeDamageDefend))
	damage += damageAttackValue
	damage -= damageDefendValue

	return damage
}

//敌人
type Enemy struct {
	hate         int
	BattleObject BattleObject
}

func (e *Enemy) String() string {
	return fmt.Sprintf("%s,hate:%d", e.BattleObject.String(), e.hate)
}

func (e *Enemy) GetHate() int {
	return e.hate
}

func (e *Enemy) AddHate(hate int) {
	e.hate += hate
}

const (
	defaultHate = 1
	taFangHate  = 2
)

//TODO 复用对象池
func CreateDefaultEnemy(BattleObject BattleObject) *Enemy {
	return &Enemy{
		BattleObject: BattleObject,
		hate:         defaultHate,
	}
}

func CreateEnemy(BattleObject BattleObject, hate int) *Enemy {
	return &Enemy{
		BattleObject: BattleObject,
		hate:         hate,
	}
}

func CreateTaFangEnemy(BattleObject BattleObject) *Enemy {
	return &Enemy{
		BattleObject: BattleObject,
		hate:         taFangHate,
	}
}
