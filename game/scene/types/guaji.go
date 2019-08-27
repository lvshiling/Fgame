package types

type GuaJiType int32

const (
	GuaJiTypeWorld GuaJiType = iota
	GuaJiTypeQuest
	GuaJiTypeActivity
	GuaJiTypeDailyQuest
	GuaJiTypeFuBen
	GuaJiTypeXianFuSilver
	GuaJiTypeXianFuExp
	GuaJiTypeXianShuangXiu
	GuaJiTypeSoulRuins
	GuaJiTypeBiaoChe
	GuaJiTypeWorldBoss
	GuaJiTypeMoonLove
	GuaJiTypeChengWai
	GuaJiTypeHuangGong
	GuaJiTypeFourGodGate
	GuaJiTypeFourGodWar
	GuaJiTypeMarry
	GuaJiTypeTower
	GuaJiTypeTowerScene
	GuaJiTypeMyBoss
	GuaJiTypeVipBoss
	GuaJiTypeUnrealBoss
	GuaJiTypeOutlandBoss
	GuaJiTypeGodSiegeScene
	GuaJiTypeLianYuScene
	GuaJiTypeArena
	GuaJiTypeShengShou
	GuaJiTypeCrossWorldBoss
	GuaJiTypeCrossTuLong
	GuaJiTypeWuJianLianYu
	GuaJiTypeCross
	GuaJiTypeMaterial
	GuaJiTypeTianJieTa
	GuaJiTypeBaGua
	GuaJiTypeTeamFuBen
	GuaJiTypeCangJingGe
	GuaJiTypeYuXi
	GuaJiTypeGuideReplicaRescue
	GuaJiTypeShenMo
	GuaJiTypeDenseWat
	GuaJiTypeShengTan
	GuaJiTypeXianTao
	GuaJiTypeLongGong
	GuaJiTypeTulong
)

var (
	guaJiTypeMap = map[GuaJiType]string{
		GuaJiTypeWorld:              "世界挂机",
		GuaJiTypeQuest:              "主线挂机",
		GuaJiTypeDailyQuest:         "日常任务挂机",
		GuaJiTypeActivity:           "活动挂机",
		GuaJiTypeFuBen:              "副本挂机",
		GuaJiTypeXianFuSilver:       "仙府银两挂机",
		GuaJiTypeXianFuExp:          "仙府经验副本挂机",
		GuaJiTypeSoulRuins:          "帝魂遗迹挂机",
		GuaJiTypeXianShuangXiu:      "双休挂机",
		GuaJiTypeBiaoChe:            "镖车挂机",
		GuaJiTypeWorldBoss:          "世界boss挂机",
		GuaJiTypeMoonLove:           "月下情缘",
		GuaJiTypeChengWai:           "城外",
		GuaJiTypeHuangGong:          "皇宫",
		GuaJiTypeFourGodGate:        "四神遗迹城门",
		GuaJiTypeFourGodWar:         "四神遗迹",
		GuaJiTypeMarry:              "结婚挂机",
		GuaJiTypeTower:              "打宝塔挂机",
		GuaJiTypeTowerScene:         "打宝塔场景挂机",
		GuaJiTypeMyBoss:             "个人boss挂机",
		GuaJiTypeVipBoss:            "vipBoss挂机",
		GuaJiTypeUnrealBoss:         "幻境BOSS",
		GuaJiTypeOutlandBoss:        "外域boss",
		GuaJiTypeGodSiegeScene:      "神兽攻城",
		GuaJiTypeLianYuScene:        "无间炼狱",
		GuaJiTypeArena:              "3v3竞技场挂机",
		GuaJiTypeShengShou:          "四圣兽挂机",
		GuaJiTypeCrossWorldBoss:     "跨服世界boss挂机",
		GuaJiTypeCrossTuLong:        "跨服屠龙挂机",
		GuaJiTypeWuJianLianYu:       "无间炼狱挂机",
		GuaJiTypeCross:              "跨服挂机",
		GuaJiTypeMaterial:           "材料副本",
		GuaJiTypeTianJieTa:          "天劫塔",
		GuaJiTypeBaGua:              "八卦秘境",
		GuaJiTypeTeamFuBen:          "组队副本",
		GuaJiTypeCangJingGe:         "藏经阁",
		GuaJiTypeYuXi:               "玉玺之战",
		GuaJiTypeGuideReplicaRescue: "救援副本",
		GuaJiTypeShenMo:             "神魔战场",
		GuaJiTypeDenseWat:           "金银密窟",
		GuaJiTypeShengTan:           "圣坛",
		GuaJiTypeXianTao:            "仙桃大会",
		GuaJiTypeLongGong:           "龙宫",
		GuaJiTypeTulong:             "屠龙",
	}
)

func (t GuaJiType) String() string {
	return guaJiTypeMap[t]
}
