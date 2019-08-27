package types

import (
	activitytypes "fgame/fgame/game/activity/types"
	majortypes "fgame/fgame/game/major/types"
	materialtypes "fgame/fgame/game/material/types"
	questtypes "fgame/fgame/game/quest/types"
	teamtypes "fgame/fgame/game/team/types"
)

//回收类型
type FoundType int32

const (
	FoundTypeFree FoundType = iota
	FoundTypeGold
)

func (r FoundType) Valid() bool {
	switch r {
	case FoundTypeGold, FoundTypeFree:
		return true
	default:
		return false
	}
}

//玩法类型
type PlayModeType int32

const (
	PlayModeTypeDailyTasks             PlayModeType = iota //日常类
	PlayModeTypeJoinTimesLimitActivity                     //次数限制活动
	PlayModeTypeFreeJoinActivity                           //无次数限制活动
)

// 资源类型
type FoundResourceType int32

const (
	FoundResourceTypeExp                    FoundResourceType = 1  //经验副本
	FoundResourceTypeSilver                                   = 2  //银两副本
	FoundResourceTypeDiHunYiJi                                = 3  //帝魂遗迹
	FoundResourceTypeTuMo                                     = 4  //屠魔任务
	FoundResourceTypeTianJiPai                                = 5  //天机牌
	FoundResourceTypeZhuoYao                                  = 6  //捉妖记
	FoundResourceTypeMoonLove                                 = 7  //月下情缘
	FoundResourceTypeJiuXiaoChengZhan                         = 8  //九霄城战
	FoundResourceTypeDailyQuest                               = 9  //日环任务
	FoundResourceTypeMaterialMount                            = 10 //坐骑副本
	FoundResourceTypeMaterialWing                             = 11 //战翼副本
	FoundResourceTypeMaterialShenFa                           = 12 //身法副本
	FoundResourceTypeMaterialFaBao                            = 13 //法宝副本
	FoundResourceTypeMaterialXianTi                           = 14 //仙体副本
	FoundResourceTypeMaterialTianMo                           = 15 //天魔材料副本
	FoundResourceTypeMaterialLingTong                         = 16 //灵童材料副本
	FoundResourceTypeMaterialLingTongWeapon                   = 17 //灵兵材料副本
	FoundResourceTypeMaterialLingTongLingYu                   = 18 //灵域材料副本
	FoundResourceTypeMaterialTeamSilver                       = 19 //银两组队副本
	FoundResourceTypeMaterialTeamLingTong                     = 20 //灵童组队副本
	FoundResourceTypeMaterialTeamEquip                        = 21 //装备组队副本
	FoundResourceTypeMaterialTeamStrengthen                   = 22 //强化组队副本
	FoundResourceTypeMaterialTeamWeapon                       = 23 //兵魂组队副本
	FoundResourceTypeMaterialTeamXueMo                        = 24 //血魔组队副本
	FoundResourceTypeMaterialTeamUpStar                       = 25 //升星组队副本
	FoundResourceTypeMaterialTeamXingChen                     = 26 //星尘组队副本
	FoundResourceTypeMajorShuangXiu                           = 27 //双修副本
	FoundResourceTypeMajorFuQi                                = 28 //夫妻副本
	FoundResourceTypeGodSiege                                 = 29 //神兽攻城
	FoundResourceTypeDenseWat                                 = 30 //金银密窟
	FoundResourceTypeShenMo                                   = 31 //神魔战场
	FoundResourceTypeFourGod                                  = 32 //四神遗迹
	FoundResourceTypeTuLong                                   = 33 //仙盟屠龙
	FoundResourceTypeLianYu                                   = 34 //无间炼狱
	FoundResourceTypeShengTan                                 = 35 //仙盟圣坛
	FoundResourceTypePeachParty                               = 36 //仙桃大会
	FoundResourceTypeShenYu                                   = 37 //神域之战
	FoundResourceTypeAllianceDaily                            = 38 //仙盟日常
)

func (t FoundResourceType) Valid() bool {
	switch t {
	case FoundResourceTypeExp,
		FoundResourceTypeSilver,
		FoundResourceTypeDiHunYiJi,
		FoundResourceTypeTuMo,
		FoundResourceTypeTianJiPai,
		FoundResourceTypeZhuoYao,
		FoundResourceTypeMoonLove,
		FoundResourceTypeJiuXiaoChengZhan,
		FoundResourceTypeDailyQuest,
		FoundResourceTypeMaterialMount,
		FoundResourceTypeMaterialWing,
		FoundResourceTypeMaterialShenFa,
		FoundResourceTypeMaterialFaBao,
		FoundResourceTypeMaterialXianTi,
		FoundResourceTypeMaterialTianMo,
		FoundResourceTypeMaterialLingTong,
		FoundResourceTypeMaterialLingTongWeapon,
		FoundResourceTypeMaterialLingTongLingYu,
		FoundResourceTypeMaterialTeamSilver,
		FoundResourceTypeMaterialTeamLingTong,
		FoundResourceTypeMaterialTeamEquip,
		FoundResourceTypeMaterialTeamStrengthen,
		FoundResourceTypeMaterialTeamWeapon,
		FoundResourceTypeMaterialTeamXueMo,
		FoundResourceTypeMaterialTeamUpStar,
		FoundResourceTypeMaterialTeamXingChen,
		FoundResourceTypeMajorShuangXiu,
		FoundResourceTypeMajorFuQi,
		FoundResourceTypeGodSiege,
		FoundResourceTypeDenseWat,
		FoundResourceTypeShenMo,
		FoundResourceTypeFourGod,
		FoundResourceTypeTuLong,
		FoundResourceTypeLianYu,
		FoundResourceTypeShengTan,
		FoundResourceTypePeachParty,
		FoundResourceTypeShenYu,
		FoundResourceTypeAllianceDaily:
		return true

	default:
		return false
	}
}

//资源类型To玩法类型
var (
	playModeTypeMap = map[FoundResourceType]PlayModeType{
		FoundResourceTypeExp:                    PlayModeTypeDailyTasks,
		FoundResourceTypeSilver:                 PlayModeTypeDailyTasks,
		FoundResourceTypeDiHunYiJi:              PlayModeTypeDailyTasks,
		FoundResourceTypeTuMo:                   PlayModeTypeDailyTasks,
		FoundResourceTypeTianJiPai:              PlayModeTypeDailyTasks,
		FoundResourceTypeZhuoYao:                PlayModeTypeJoinTimesLimitActivity,
		FoundResourceTypeMoonLove:               PlayModeTypeFreeJoinActivity,
		FoundResourceTypeJiuXiaoChengZhan:       PlayModeTypeFreeJoinActivity,
		FoundResourceTypeDailyQuest:             PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialMount:          PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialWing:           PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialShenFa:         PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialFaBao:          PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialXianTi:         PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTianMo:         PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialLingTong:       PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialLingTongWeapon: PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialLingTongLingYu: PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamSilver:     PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamLingTong:   PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamEquip:      PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamStrengthen: PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamWeapon:     PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamXueMo:      PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamUpStar:     PlayModeTypeDailyTasks,
		FoundResourceTypeMaterialTeamXingChen:   PlayModeTypeDailyTasks,
		FoundResourceTypeMajorShuangXiu:         PlayModeTypeDailyTasks,
		FoundResourceTypeMajorFuQi:              PlayModeTypeDailyTasks,
		FoundResourceTypeGodSiege:               PlayModeTypeFreeJoinActivity,
		FoundResourceTypeDenseWat:               PlayModeTypeFreeJoinActivity,
		FoundResourceTypeShenMo:                 PlayModeTypeFreeJoinActivity,
		FoundResourceTypeFourGod:                PlayModeTypeFreeJoinActivity,
		FoundResourceTypeTuLong:                 PlayModeTypeFreeJoinActivity,
		FoundResourceTypeLianYu:                 PlayModeTypeFreeJoinActivity,
		FoundResourceTypeShengTan:               PlayModeTypeFreeJoinActivity,
		FoundResourceTypePeachParty:             PlayModeTypeFreeJoinActivity,
		FoundResourceTypeShenYu:                 PlayModeTypeFreeJoinActivity,
		FoundResourceTypeAllianceDaily:          PlayModeTypeDailyTasks,
	}
)

func (r FoundResourceType) GetPlayModeType() PlayModeType {
	return playModeTypeMap[r]
}

func GetPlayModeTypeMap() map[FoundResourceType]PlayModeType {
	return playModeTypeMap
}

// 活动类型To资源类型
var (
	activityToResTypeMap = map[activitytypes.ActivityType]FoundResourceType{
		activitytypes.ActivityTypeMoonLove:         FoundResourceTypeMoonLove,
		activitytypes.ActivityTypeAlliance:         FoundResourceTypeJiuXiaoChengZhan,
		activitytypes.ActivityTypeGodSiegeQiLin:    FoundResourceTypeGodSiege,
		activitytypes.ActivityTypeDenseWat:         FoundResourceTypeDenseWat,
		activitytypes.ActivityTypeShenMoWar:        FoundResourceTypeShenMo,
		activitytypes.ActivityTypeFourGod:          FoundResourceTypeFourGod,
		activitytypes.ActivityTypeCoressTuLong:     FoundResourceTypeTuLong,
		activitytypes.ActivityTypeLianYu:           FoundResourceTypeLianYu,
		activitytypes.ActivityTypeAllianceShengTan: FoundResourceTypeShengTan,
		activitytypes.ActivityTypeXianTaoDaHui:     FoundResourceTypePeachParty,
		activitytypes.ActivityTypeShenYu:           FoundResourceTypeShenYu,
	}
)

func ActivityTypeToFoundResType(acType activitytypes.ActivityType) (resType FoundResourceType, flag bool) {
	resType, flag = activityToResTypeMap[acType]
	return
}

// 任务类型To资源类型
var (
	questToResTypeMap = map[questtypes.QuestType]FoundResourceType{
		questtypes.QuestTypeDaily:         FoundResourceTypeDailyQuest,
		questtypes.QuestTypeTuMo:          FoundResourceTypeTuMo,
		questtypes.QuestTypeTianJiPai:     FoundResourceTypeTianJiPai,
		questtypes.QuestTypeDailyAlliance: FoundResourceTypeAllianceDaily,
	}
)

func QuestTypeToFoundResType(questType questtypes.QuestType) (resType FoundResourceType, flag bool) {
	resType, flag = questToResTypeMap[questType]
	return
}

// 夫妻副本To资源类型
var (
	majorToResTypeMap = map[majortypes.MajorType]FoundResourceType{
		majortypes.MajorTypeShuangXiu: FoundResourceTypeMajorShuangXiu,
		majortypes.MajorTypeFuQi:      FoundResourceTypeMajorFuQi,
	}
)

func MajorTypeToFoundResType(materialType majortypes.MajorType) (resType FoundResourceType, flag bool) {
	resType, flag = majorToResTypeMap[materialType]
	return
}

// 材料副本类型To资源类型
var (
	materialToResTypeMap = map[materialtypes.MaterialType]FoundResourceType{
		materialtypes.MaterialTypeMount:    FoundResourceTypeMaterialMount,
		materialtypes.MaterialTypeWing:     FoundResourceTypeMaterialWing,
		materialtypes.MaterialTypeShenfa:   FoundResourceTypeMaterialShenFa,
		materialtypes.MaterialTypeFabao:    FoundResourceTypeMaterialFaBao,
		materialtypes.MaterialTypeXianti:   FoundResourceTypeMaterialXianTi,
		materialtypes.MaterialTypeTianMo:   FoundResourceTypeMaterialTianMo,
		materialtypes.MaterialTypeLingTong: FoundResourceTypeMaterialLingTong,
		materialtypes.MaterialTypeLingBing: FoundResourceTypeMaterialLingTongWeapon,
		materialtypes.MaterialTypeLingYu:   FoundResourceTypeMaterialLingTongLingYu,
	}
)

func MaterialTypeToFoundResType(materialType materialtypes.MaterialType) (resType FoundResourceType, flag bool) {
	resType, flag = materialToResTypeMap[materialType]
	return
}

// 组队副本类型To资源类型
var (
	materialTeamToResTypeMap = map[teamtypes.TeamPurposeType]FoundResourceType{
		teamtypes.TeamPurposeTypeFuBenSilver:           FoundResourceTypeMaterialTeamSilver,
		teamtypes.TeamPurposeTypeFuBenXingChen:         FoundResourceTypeMaterialTeamXingChen,
		teamtypes.TeamPurposeTypeFuBenXueMo:            FoundResourceTypeMaterialTeamXueMo,
		teamtypes.TeamPurposeTypeFuBenZhuangShengEquip: FoundResourceTypeMaterialTeamEquip,
		teamtypes.TeamPurposeTypeFuBenWeapon:           FoundResourceTypeMaterialTeamWeapon,
		teamtypes.TeamPurposeTypeFuBenStrength:         FoundResourceTypeMaterialTeamStrengthen,
		teamtypes.TeamPurposeTypeFuBenLingTong:         FoundResourceTypeMaterialTeamLingTong,
		teamtypes.TeamPurposeTypeFuBenUpstar:           FoundResourceTypeMaterialTeamUpStar,
	}
)

func MaterialTeamTypeToFoundResType(materialType teamtypes.TeamPurposeType) (resType FoundResourceType, flag bool) {
	resType, flag = materialTeamToResTypeMap[materialType]
	return
}

//找回状态
type FoundBackStatus int32

const (
	FoundBackStatusWaitReceive FoundBackStatus = iota //待领取
	FoundBackStatusHadReceive                         //已领取
	FoundBackStatusNotFound                           //不可找回
)
