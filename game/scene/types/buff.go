package types

import (
	propertytyeps "fgame/fgame/game/property/types"
)

//buff类型
type BuffType int32

const (
	//通用buff
	BuffTypeCommon BuffType = iota
	//被动特殊光效
	BuffTypePassive
	//定制称号buff
	BuffTypeTitleDingZhi
)

func (bt BuffType) Valid() bool {
	switch bt {
	case BuffTypeCommon,
		BuffTypePassive,
		BuffTypeTitleDingZhi:
		return true
	}
	return false
}

type BuffRemoveType int32

const (
	//走路打断
	BuffRemoveTypeWalk BuffRemoveType = iota
	//骑乘打断
	BuffRemoveTypeMount
	//死亡打断
	BuffRemoveTypeDead
	//攻击打断
	BuffRemoveTypeAttack
	//受击打断
	BuffRemoveTypeAttacked
	//治疗打断
	BuffRemoveTypeCure
	//切换地图打断
	BuffRemoveTypeChangeScene
	//脱战打断
	BuffRemoveTypeExitBattle
	//切换pk模式打断
	BuffRemoveTypeSwitchPK
	//采集打断
	BuffRemoveTypeCollect
)

func (t BuffRemoveType) Mask() int32 {
	return 1 << uint(t)
}

//战斗限制类型
type BattleLimitType int32

const (
	//限制使用物品
	BattleLimitTypeItemUse BattleLimitType = iota
	//不能移动
	BattleLimitTypeMove
	//不能使用技能
	BattleLimitTypeSkill
	//不能被技能攻击
	BattleLimitTypeAttacked
	//不能显示(隐身)
	BattleLimitTypeHidden
	//不能传送
	BattleLimitTypeCrossScene
	//禁止轻功
	BattleLimitTypeQingGong
	//禁止回血
	BattleLimitTypeRecovery
	//不能打坐
	BattleLimitTypeDaZuo
	//不能采集
	BattleLimitTypeCaiJi
	//不能回体
	BattleLimitTypeRecoveryTi
	//过场保护不能被攻击
	BattleLimitTypeAttackedChangeScene
	//复活保护不能被攻击
	BattleLimitTypeAttackedRelive
	//不是使用所有技能
	BattleLimitTypeSkillAll
	//pk保护不能被攻击
	BattleLimitTypeAttackedPKProtect
	//不能上马
	BattleLimitTypeMount
	//无敌状态不能被攻击
	BattleLimitTypeNoAttacked
)

func (t BattleLimitType) Mask() int64 {
	return 1 << uint(t)
}

type BuffTouchType int32

const (
	//获得状态触发
	BuffTouchTypeImmediate BuffTouchType = iota
	//定时触发
	BuffTouchTypeTimer
	//获得伤害触发
	BuffTouchTypeHurted
	//获得控制类状态触发
	BuffTouchTypeBuff
	//死亡触发
	BuffTouchTypeDead
	//对目标造成伤害触发（子状态给受击者，其他给自己）
	BuffTouchTypeObjectDamage
	//对目标造成伤害触发（子状态以及其他状态均给自己）
	BuffTouchTypeObjectDamageSelf
	//跳闪成功触发
	BuffTouchTypeJump
	//被攻击时触发(子状态给攻击者，其他给自己)
	BuffTouchTypeHurtedOther
)

func (btt BuffTouchType) Valid() bool {
	switch btt {
	case BuffTouchTypeImmediate,
		BuffTouchTypeTimer,
		BuffTouchTypeHurted,
		BuffTouchTypeBuff,
		BuffTouchTypeDead,
		BuffTouchTypeObjectDamage,
		BuffTouchTypeObjectDamageSelf,
		BuffTouchTypeJump,
		BuffTouchTypeHurtedOther:
		return true
	}
	return false
}

type BuffImmuneType int32

const (
	//玩家免疫
	BuffImmuneTypePlayer BuffImmuneType = 1 << iota
	//怪物免疫
	BuffImmuneTypeMonster
	//BOSS免疫
	BuffImmuneTypeBoss
	//世界boss
	BuffImmuneTypeWorldBoss
	//天劫塔boss
	BuffImmuneTypeRealmBoss
	//非生物类型(镖车)
	BuffImmuneTypeSpecialBiologyBiaoChe
	//非生物类型(城门)
	BuffImmuneTypeSpecialBiologyChengMen
	//怪物免疫(邪将传令官)
	BuffImmuneTypeSpecialMonster
	//藏经阁(BOSS)
	BuffImmuneTypeCangJingGeBoss
)

//下线保存
type BuffOfflineSaveType int32

const (
	BuffOfflineSaveTypeNone = iota
	BuffOfflineSaveTypeTimeStop
	BuffOfflineSaveTypeTimeContinue
)

func (bost BuffOfflineSaveType) Valid() bool {
	switch bost {
	case BuffOfflineSaveTypeNone,
		BuffOfflineSaveTypeTimeStop,
		BuffOfflineSaveTypeTimeContinue:
		return true
	}
	return false
}

type BuffStackType int32

const (
	//时间叠加
	BuffStackTypeTime BuffStackType = 1 << iota
	//效果叠加
	BuffStackTypeEffect
)

type BuffPiaoZi int32

const (
	BuffPiaoZiZhuoShao BuffPiaoZi = iota + 1
	BuffPiaoZiHuiChun
	BuffPiaoZiZhongDu
)

var (
	buffDamangeTypeMap = map[BuffPiaoZi]DamageType{
		BuffPiaoZiZhuoShao: DamageTypeZhuoShao,
		BuffPiaoZiHuiChun:  DamageTypeRecovery,
		BuffPiaoZiZhongDu:  DamageTypeZhongDu,
	}
)

func (z BuffPiaoZi) DamageType() (flag bool, damageType DamageType) {
	damageType, ok := buffDamangeTypeMap[z]
	if !ok {
		return
	}
	flag = true
	return
}

type BuffKangXing int32

const (
	BuffKangXingZhuoShao BuffKangXing = iota + 1
	BuffKangXingJianSu
	BuffKangXingShiMing
	BuffKangXingDingShen
	BuffKangXingBingDong
	BuffKangXingHunMi
	BuffKangXingZhongDu
	BuffKangXingKuiLei
	BuffKangXingPoJia
	BuffKangXingKuJie
	BuffKangXingXuRuo
	BuffKangXingJiaoXie
)

var (
	buffKangXingMap = map[BuffKangXing]propertytyeps.BattlePropertyType{
		BuffKangXingZhuoShao: propertytyeps.BattlePropertyTypeZhuoShaoRes,
		BuffKangXingJianSu:   propertytyeps.BattlePropertyTypeJianSuRes,
		BuffKangXingShiMing:  propertytyeps.BattlePropertyTypeShiMingRes,
		BuffKangXingDingShen: propertytyeps.BattlePropertyTypeDingShenRes,
		BuffKangXingBingDong: propertytyeps.BattlePropertyTypeBingDongRes,
		BuffKangXingHunMi:    propertytyeps.BattlePropertyTypeHunMiRes,
		BuffKangXingZhongDu:  propertytyeps.BattlePropertyTypeZhongDuRes,
		BuffKangXingKuiLei:   propertytyeps.BattlePropertyTypeKuiLeiRes,
		BuffKangXingPoJia:    propertytyeps.BattlePropertyTypePoJiaRes,
		BuffKangXingKuJie:    propertytyeps.BattlePropertyTypeKuJieRes,
		BuffKangXingXuRuo:    propertytyeps.BattlePropertyTypeXuRuoRes,
		BuffKangXingJiaoXie:  propertytyeps.BattlePropertyTypeJiaoXieRes,
	}
)

func (x BuffKangXing) PropertyType() (flag bool, typ propertytyeps.BattlePropertyType) {
	typ, ok := buffKangXingMap[x]
	if !ok {
		return
	}
	flag = true
	return
}

type BuffEffectType int32

const (
	BuffEffectTypeNone         BuffEffectType = 0
	BuffEffectTypeHuDun                       = 2 //护盾
	BuffEffectTypeXueChiDeBuff                = 4 //血池回复效果削弱
)

func (t BuffEffectType) Valid() bool {
	switch t {
	case BuffEffectTypeHuDun,
		BuffEffectTypeNone,
		BuffEffectTypeXueChiDeBuff:
		return true
	}
	return false
}
