package types

import (
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
)

type SkillFirstType int32

const (
	//普通技能
	SkillFirstTypeNormal SkillFirstType = iota
	//职业技能
	SkillFirstTypeRole
	//坐骑技能
	SkillFirstTypeMount
	//异兽技能
	SkillFirstTypeYiShou
	//古魂技能
	SkillFirstTypeGuHun
	//跳跃技能
	SkillFirstTypeJump
	//绝学技能
	SkillFirstTypeJueXue
	//境界技能
	SkillFirstTypeRealm
	//心法技能
	SkillFirstTypeXinFa
	//仙盟技能
	SkillFirstTypeAlliance
	//龙神技能-------10
	SkillFirstTypeLongShen
	//怪物被动技能
	SkillFirstTypeMonsterPassive
	//身法
	SkillFirstTypeShenfa
	//领域
	SkillFirstTypeLingyu
	//元神金装套装技能
	SkillFirstTypeGoldEquipSuit
	//暗器
	SkillFirstTypeAnQi
	//骑装
	SkillFirstTypeMountEquip
	//战翼符石
	SkillFirstTypeWingStone
	//暗器机关
	SkillFirstTypeAnQiJiGuan
	//法宝技能
	SkillFirstTypeFaBao
	//仙体技能-----20
	SkillFirstTypeXianTi
	//血盾技能
	SkillFirstTypeXueDun
	//衣橱技能
	SkillFirstTypeWardrobe
	//天魔体技能（自动激活）
	SkillFirstTypeTianMoAdvancedSkill
	//噬魂幡技能（自动激活）
	SkillFirstTypeShiHunFanAdvancedSkill
	//天魔体技能（手动激活)
	SkillFirstTypeTianMoSystemSkill
	//噬魂幡技能（手动激活)
	SkillFirstTypeShiHunFanSystemSkill
	//领域技能（手动激活)
	SkillFirstTypeLingYuSystemSkill
	//身法技能（手动激活)
	SkillFirstTypeShenFaSystemSkill
	//灵童兵魂技能(手动激活
	SkillFirstTypeLingTongWeaponSystemSkill
	//灵童坐骑技能(手动激活----30-
	SkillFirstTypeLingTongMountSystemSkill
	//灵童战翼技能(手动激活
	SkillFirstTypeLingTongWingSystemSkill
	//灵童身法技能(手动激活
	SkillFirstTypeLingTongShenFaSystemSkill
	//灵童领域技能(手动激活
	SkillFirstTypeLingTongLingYuSystemSkill
	//灵童法宝技能(手动激活
	SkillFirstTypeLingTongFaBaoSystemSkill
	//灵童仙体技能(手动激活
	SkillFirstTypeLingTongXianTiSystemSkill
	//灵童普通攻击
	SkillFirstTypeLingTongNormal
	//灵童技能
	SkillFirstTypeLingTongSkill
	//子技能
	SkillFirstTypeSubSkill
	//圣痕技能青龙
	SkillFirstTypeShengHenQingLong
	//圣痕技能白虎---------40
	SkillFirstTypeShengHenBaiHu
	//圣痕技能朱雀
	SkillFirstTypeShengHenZhuQue
	//圣痕技能玄武
	SkillFirstTypeShengHenXuanWu
	//神器技能
	SkillFirstTypeShenQi
	//屠龙装备技能
	SkillFirstTypeTuLongEquip
	//屠龙装备套装技能
	SkillFirstTypeTuLongEquipSuit
	//宝宝天赋技能
	SkillFirstTypeBabyTalent
	//宝宝玩具套装技能
	SkillFirstTypeBabyToySuit
	//八卦符石技能
	SkillFirstTypeFuShi
	//神铸-锻魂技能
	SkillFirstTypeCastingSoul
	//五行灵珠技能------------50
	SkillFirstTypeWuXingLingZhu
	//特戒技能
	SkillFirstTypeRing
)

func (sft SkillFirstType) Valid() bool {
	switch sft {
	case SkillFirstTypeNormal,
		SkillFirstTypeRole,
		SkillFirstTypeMount,
		SkillFirstTypeYiShou,
		SkillFirstTypeGuHun,
		SkillFirstTypeJump,
		SkillFirstTypeJueXue,
		SkillFirstTypeRealm,
		SkillFirstTypeXinFa,
		SkillFirstTypeAlliance,
		SkillFirstTypeLongShen,
		SkillFirstTypeMonsterPassive,
		SkillFirstTypeShenfa,
		SkillFirstTypeLingyu,
		SkillFirstTypeGoldEquipSuit,
		SkillFirstTypeAnQi,
		SkillFirstTypeMountEquip,
		SkillFirstTypeWingStone,
		SkillFirstTypeAnQiJiGuan,
		SkillFirstTypeFaBao,
		SkillFirstTypeXueDun,
		SkillFirstTypeXianTi,
		SkillFirstTypeWardrobe,
		SkillFirstTypeTianMoAdvancedSkill,
		SkillFirstTypeShiHunFanAdvancedSkill,
		SkillFirstTypeTianMoSystemSkill,
		SkillFirstTypeShiHunFanSystemSkill,
		SkillFirstTypeLingYuSystemSkill,
		SkillFirstTypeShenFaSystemSkill,
		SkillFirstTypeLingTongWeaponSystemSkill,
		SkillFirstTypeLingTongMountSystemSkill,
		SkillFirstTypeLingTongWingSystemSkill,
		SkillFirstTypeLingTongShenFaSystemSkill,
		SkillFirstTypeLingTongLingYuSystemSkill,
		SkillFirstTypeLingTongFaBaoSystemSkill,
		SkillFirstTypeLingTongXianTiSystemSkill,
		SkillFirstTypeLingTongNormal,
		SkillFirstTypeLingTongSkill,
		SkillFirstTypeSubSkill,
		SkillFirstTypeShengHenQingLong,
		SkillFirstTypeShengHenBaiHu,
		SkillFirstTypeShengHenZhuQue,
		SkillFirstTypeShengHenXuanWu,
		SkillFirstTypeShenQi,
		SkillFirstTypeTuLongEquip,
		SkillFirstTypeTuLongEquipSuit,
		SkillFirstTypeBabyTalent,
		SkillFirstTypeBabyToySuit,
		SkillFirstTypeFuShi,
		SkillFirstTypeCastingSoul,
		SkillFirstTypeWuXingLingZhu,
		SkillFirstTypeRing:
		return true
	default:
		return false
	}
}

var (
	//用于收集skill模块属性计算的
	skillModulePropertyEffectoryMap = map[SkillFirstType]string{
		SkillFirstTypeNormal:         "普通技能",
		SkillFirstTypeRole:           "职业技能",
		SkillFirstTypeYiShou:         "异兽技能",
		SkillFirstTypeJump:           "跳跃技能",
		SkillFirstTypeJueXue:         "绝学技能",
		SkillFirstTypeXinFa:          "心法技能",
		SkillFirstTypeAlliance:       "仙盟技能",
		SkillFirstTypeLongShen:       "龙神技能",
		SkillFirstTypeMonsterPassive: "怪物被动技能",
		SkillFirstTypeSubSkill:       "子技能",
		SkillFirstTypeFuShi:          "八卦符石技能",
	}
)

func GetSkillModulePropertyEffectoryMap() map[SkillFirstType]string {
	return skillModulePropertyEffectoryMap
}

type SkillSecondType int32

const (
	//主动
	SkillSecondTypePositive SkillSecondType = iota + 1
	//攻击概率
	SkillSecondTypeAttackProbability
	//被攻击概率
	SkillSecondTypeAttackedProbability
	//被动
	SkillSecondTypePassive
	//血量低于一定值触发
	SkillSecondTypeHp
	//复活
	SkillSecondTypeReborn
)

func (sst SkillSecondType) Valid() bool {
	switch sst {
	case SkillSecondTypePositive,
		SkillSecondTypeAttackProbability,
		SkillSecondTypeAttackedProbability,
		SkillSecondTypePassive,
		SkillSecondTypeHp,
		SkillSecondTypeReborn:
		return true
	default:
		return false
	}
}

type SkillThirdType int32

const (
	//静态
	SkillThirdTypeStatic SkillThirdType = iota
	//动态
	SkillThirdTypeDynamic
)

func (stt SkillThirdType) Valid() bool {
	switch stt {
	case SkillThirdTypeStatic,
		SkillThirdTypeDynamic:
		return true
	default:
		return false
	}
}

type SkillFourthType int32

const (
	//伤害技能
	SkillFourthTypeAttack SkillFourthType = iota
	//辅助技能
	SkillFourthTypeHelp
)

func (stt SkillFourthType) Valid() bool {
	switch stt {
	case SkillFourthTypeAttack,
		SkillFourthTypeHelp:
		return true
	default:
		return false
	}
}

//技能需求目标
type SkillTargetNeedType int32

const (
	//自己
	SkillTargetNeedTypeSelf SkillTargetNeedType = iota
	//友方
	SkillTargetNeedTypeAlliance
	//敌方
	SkillTargetNeedTypeEnemy
	//全体
	SkillTargetNeedTypeAll
	//范围
	SkillTargetNeedTypeRange
)

func (stnt SkillTargetNeedType) Valid() bool {
	switch stnt {
	case SkillTargetNeedTypeSelf,
		SkillTargetNeedTypeAlliance,
		SkillTargetNeedTypeEnemy,
		SkillTargetNeedTypeAll,
		SkillTargetNeedTypeRange:
		return true
	default:
		return false
	}
}

//技能生效目标
type SkillTargetSelectType int32

const (
	SkillTargetSelectTypeSelf SkillTargetSelectType = iota
	SkillTargetSelectTypeAlliance
	SkillTargetSelectTypeEnemy
)

func (stst SkillTargetSelectType) Mask() int32 {
	return int32(1 << uint(stst))
}

//技能作用目标
type SkillTargetActionType int32

const (
	SkillTargetActionTypePlayer SkillTargetSelectType = iota
	SkillTargetActionTypeMonster
)

func (stst SkillTargetActionType) Mask() int32 {
	return int32(1 << uint(stst))
}

//无视安全区施法
type SkillRuleBreakType int32

const (
	//安全区无法施法
	SkillRuleBreakTypeNo SkillRuleBreakType = iota
	//安全区施法
	SkillRuleBreakTypeYes
)

func (srbt SkillRuleBreakType) Valid() bool {
	switch srbt {
	case SkillRuleBreakTypeNo,
		SkillRuleBreakTypeYes:
		return true
	default:
		return false
	}
}

//技能朝向
type SkillFaceNeedType int32

const (
	//无需朝向
	SkillFaceNeedTypeNo SkillFaceNeedType = iota
	//需要朝向
	SkillFaceNeedTypeYes
)

func (sfnt SkillFaceNeedType) Valid() bool {
	switch sfnt {
	case SkillFaceNeedTypeNo,
		SkillFaceNeedTypeYes:
		return true
	default:
		return false
	}
}

//技能范围
type SkillAreaType int32

const (
	SkillAreaTypeDefault SkillAreaType = iota
	SkillAreaTypeSingle
	SkillAreaTypeLine
	SkillAreaTypeFan
	SkillAreaTypeRound
)

func (sat SkillAreaType) Valid() bool {
	switch sat {
	case SkillAreaTypeDefault,
		SkillAreaTypeSingle,
		SkillAreaTypeLine,
		SkillAreaTypeFan,
		SkillAreaTypeRound:
		return true
	default:
		return false
	}
}

type SkillArea interface {
	PositionInArea(pos coretypes.Position, angle float64, targetPos coretypes.Position, targetAngle float64, targetLength float64, targetWidth float64) bool
}

//技能直线范围
type skillAreaRectangle struct {
	Length float64
	Width  float64
}

const (
	defaultWidth = 0.1
)

//TODO 验证
func (sar *skillAreaRectangle) PositionInArea(pos coretypes.Position, angle float64, targetPos coretypes.Position, targetAngle float64, targetLength float64, targetWidth float64) bool {
	//生成多边形

	targetPolygons := coreutils.GetRectangle(targetPos, targetLength, targetWidth, targetAngle)

	skillPolygons := coreutils.GetRectangleByButtom(pos, sar.Length, sar.Width, angle)
	if coreutils.PolygonIntersectPolygon(targetPolygons, skillPolygons) {
		return true
	}
	return false
	// deltaX := targetPos.X - pos.X
	// deltaZ := targetPos.Z - pos.Z
	// radian := mathutils.AngleToRadian(90 - angle)
	// forwardX := math.Cos(radian)
	// forwardZ := math.Sin(radian)
	// forwardDistance := deltaX*forwardX + deltaZ*forwardZ
	// if forwardDistance > sar.Length || forwardDistance <= 0 {
	// 	return false
	// }

	// rightDistance := deltaX*forwardZ - deltaZ*forwardX
	// if math.Abs(rightDistance) > sar.Width/2 {
	// 	return false
	// }
	// return true
}

func NewSkillAreaRectangle(length float64, width float64) SkillArea {
	return &skillAreaRectangle{
		Length: length,
		Width:  width,
	}
}

//技能扇形
type skillAreaFan struct {
	//角度
	Angle float64
	//半径
	Radius       float64
	radiusSquare float64
}

func (saf *skillAreaFan) PositionInArea(pos coretypes.Position, angle float64, targetPos coretypes.Position, targetAngle float64, targetLength float64, targetWidth float64) bool {

	targetPolygons := coreutils.GetRectangle(targetPos, targetLength, targetWidth, targetAngle)

	if coreutils.PolygonIntersectFan(targetPolygons, pos, saf.Radius, angle, saf.Angle) {
		return true
	}
	return false
	// deltaX := targetPos.X - pos.X
	// deltaZ := targetPos.Z - pos.Z
	// minAngle := angle - saf.Angle/2
	// maxAngle := angle + saf.Angle/2

	// xyAngle := utils.GetAngle(pos, targetPos)
	// isAngleIn := mathutils.BetweenAngles(xyAngle, minAngle, maxAngle)

	// if !isAngleIn {
	// 	return false
	// }
	// distance := deltaX*deltaX + deltaZ*deltaZ

	// return distance <= saf.radiusSquare
}

func NewSkillAreaFan(angle float64, radius float64) SkillArea {
	return &skillAreaFan{
		Angle:        angle,
		Radius:       radius,
		radiusSquare: radius * radius,
	}
}

//技能圆形
type skillAreaRound struct {
	Radius       float64
	radiusSquare float64
}

func (saf *skillAreaRound) PositionInArea(pos coretypes.Position, angle float64, targetPos coretypes.Position, targetAngle float64, targetLength float64, targetWidth float64) bool {
	targetPolygons := coreutils.GetRectangle(targetPos, targetLength, targetWidth, targetAngle)

	if coreutils.PolygonIntersectRound(targetPolygons, pos, saf.Radius) {
		return true
	}
	return false
	// deltaX := pos.X - targetPos.X
	// deltaZ := pos.Z - targetPos.Z
	// distance := deltaX*deltaX + deltaZ*deltaZ

	// return distance <= saf.radiusSquare
}

func NewSkillAreaRound(radius float64) SkillArea {
	return &skillAreaRound{
		Radius:       radius,
		radiusSquare: radius * radius,
	}
}

//特效效果标识
type SkillSpecialEffectType int32

const (
	SkillSpecialEffectTypeNone SkillSpecialEffectType = iota
	//拉近
	SkillSpecialEffectTypeClose
	//击退
	SkillSpecialEffectTypeRepel
	//冲刺
	SkillSpecialEffectTypeSprint
)

func (sset SkillSpecialEffectType) Valid() bool {
	switch sset {
	case SkillSpecialEffectTypeNone,
		SkillSpecialEffectTypeClose,
		SkillSpecialEffectTypeRepel,
		SkillSpecialEffectTypeSprint:
		return true
	default:
		return false
	}
}

//限制被动技能触发
type SkillLimitTouchType int32

const (
	SkillLimitTouchTypeNo SkillLimitTouchType = iota
	SkillLimitTouchTypeYes
)

func (sltt SkillLimitTouchType) Valid() bool {
	switch sltt {
	case SkillLimitTouchTypeNo,
		SkillLimitTouchTypeYes:
		return true
	default:
		return false
	}
}

type SkillImmuneType int32

const (
	SkillImmuneTypeClose SkillImmuneType = 1 << iota
	SkillImmuneTypeRepel
)

//被打触发
type SkillBeTriggerType int32

const (
	//被玩家打
	SkillBeTriggerTypePlayer SkillBeTriggerType = iota
	//被怪物打
	SkillBeTriggerTypeMonster
)

func (t SkillBeTriggerType) Mask() int32 {
	return int32(1 << uint(t))
}

//特殊作用目标
type SkillSpecialTarget int32

const (
	//怪物
	SkillSpecialTargetMonster SkillSpecialTarget = iota
	//玩家
	SkillSpecialTargetPlayer
)

func (t SkillSpecialTarget) Mask() int32 {
	return int32(1 << uint(t))
}
