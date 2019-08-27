package types

import (
	centertypes "fgame/fgame/center/types"
	activitytypes "fgame/fgame/game/activity/types"
	crosstypes "fgame/fgame/game/cross/types"
)

//地图大类型
type MapType int32

const (
	//默认
	MapTypeNone MapType = iota
	//世界
	MapTypeWorld
	//副本
	MapTypeFuBen
	//活动
	MapTypeActivity
	//活动子场景
	MapTypeActivitySub
	//结婚场景
	MapTypeMarry
	//竞技场景
	MapTypeArena
	//世界boss
	// MapTypeWorldBoss
	//跨服世界boss
	// MapTypeCrossWorldBoss
	//打宝塔
	MapTypeTower
	//幻境BOSS
	// MapTypeUnrealBoss
	//外域BOSS
	// MapTypeOutlandBoss
	//藏经阁
	// MapTypeCangJingGe
	//活动副本
	MapTypeActivityFuBen
	//创世之战（世界）
	MapTypeChuangShiWorld
	//珍惜boss
	// MapTypeZhenXiBoss
	//定时boss
	// MapTypeDingShiBoss
	//boss
	MapTypeBoss
)

//场景类型
type SceneType int32

//统一处理
const (
	//世界场景
	SceneTypeWorld SceneType = 1
	//帝陵遗迹副本场景
	SceneTypeFuBenSoulRuins = 2
	//八点半
	SceneTypeEightClock = 3
	//天劫塔
	SceneTypeTianJieTa = 4
	//材料副本
	SceneTypeMaterial = 6
	//世界boss
	SceneTypeWorldBoss = 7
	//皇城争霸
	SceneTypeChengZhan = 10
	//银两副本
	SceneTypeYinLiang = 11
	//月下情缘
	SceneTypeYueXiaQingYuan = 12
	//经验副本
	SceneTypeExperience = 13
	//元宝副本
	SceneTypeGold = 14
	//皇宫地图
	SceneTypeHuangGong = 15
	//竞技场
	SceneTypeArena = 16
	//圣兽
	SceneTypeArenaShengShou = 17
	//四神遗迹出生地图
	SceneTypeFourGodGate = 18
	//四神遗迹地图
	SceneTypeFourGodWar = 19
	//结婚场景
	SceneTypeMarry = 20
	//灵池战斗场景
	SceneTypeLingChiFighting = 21
	//双修副本地图
	SceneTypeMajor = 22
	//跨服世界boss
	SceneTypeCrossWorldBoss = 23
	//跨服屠龙
	SceneTypeCrossTuLong = 24
	//无间炼狱（跨服）
	SceneTypeCrossLianYu = 25
	//反击魔界
	SceneTypeBeatBackMoJie = 26
	//打宝塔
	SceneTypeTower = 27
	//个人BOSS
	SceneTypeMyBoss = 28
	//VIP个人BOSS
	SceneTypeMyBossVip = 29
	//神兽攻城
	SceneTypeCrossGodSiege = 30
	//幻境BOSS
	SceneTypeUnrealBoss = 31
	//外域BOSS
	SceneTypeOutlandBoss = 32
	//八卦秘境
	SceneTypeBaGuaMiJing = 33
	//引导副本
	SceneTypeGuideReplica = 34
	//组队副本
	SceneTypeCrossTeamCopy = 35
	//金银密窟
	SceneTypeCrossDenseWat = 36
	//神魔战场
	SceneTypeCrossShenMo = 37
	//藏经阁
	SceneTypeCangJingGe = 38
	//仙盟boss地图
	SceneTypeAllianceBoss = 39
	//仙盟圣坛
	SceneTypeAllianceShengTan = 40
	//夫妻副本
	SceneTypeFuQiFuBen = 41
	//仙桃大会
	SceneTypeXianTaoDaHui = 42
	//神域之战
	SceneTypeShenYu = 43
	//龙宫探宝
	SceneTypeLongGong = 44
	//玉玺之战
	SceneTypeYuXi = 45
	//魔剑副本
	SceneTypeGuideReplicaMoJian = 46
	//猫狗副本
	SceneTypeGuideReplicaCatDog = 47
	//救援副本
	SceneTypeGuideReplicaRescue = 48
	//运营活动-奇遇岛
	SceneTypeWelfareQiYu = 49
	//本服-无间炼狱
	SceneTypeLocalLianYu = 50
	//pvp
	SceneTypeArenapvpHaiXuan = 51
	//中立城
	SceneTypeChuangShiZhiZhanZhongLi = 52
	//比武大会
	SceneTypeArenapvp = 53
	//主城
	SceneTypeChuangShiZhiZhanMain = 54
	//附属
	SceneTypeChuangShiZhiZhanFuShu = 55
	//珍惜boss
	SceneTypeZhenXi  = 56
	SceneTypeDingShi = 57
)

var (
	sceneTypeMap = map[SceneType]MapType{
		SceneTypeWorld:                   MapTypeWorld,
		SceneTypeFuBenSoulRuins:          MapTypeFuBen,
		SceneTypeEightClock:              MapTypeActivity,
		SceneTypeTianJieTa:               MapTypeFuBen,
		SceneTypeWorldBoss:               MapTypeBoss,
		SceneTypeChengZhan:               MapTypeActivity,
		SceneTypeYinLiang:                MapTypeFuBen,
		SceneTypeYueXiaQingYuan:          MapTypeActivity,
		SceneTypeExperience:              MapTypeFuBen,
		SceneTypeGold:                    MapTypeFuBen,
		SceneTypeArena:                   MapTypeArena,
		SceneTypeArenaShengShou:          MapTypeBoss,
		SceneTypeFourGodGate:             MapTypeActivitySub,
		SceneTypeFourGodWar:              MapTypeActivity,
		SceneTypeMarry:                   MapTypeMarry,
		SceneTypeLingChiFighting:         MapTypeFuBen,
		SceneTypeMajor:                   MapTypeFuBen,
		SceneTypeCrossWorldBoss:          MapTypeBoss,
		SceneTypeCrossTuLong:             MapTypeActivity,
		SceneTypeCrossLianYu:             MapTypeActivity,
		SceneTypeBeatBackMoJie:           MapTypeActivity,
		SceneTypeTower:                   MapTypeTower,
		SceneTypeMyBoss:                  MapTypeFuBen,
		SceneTypeMyBossVip:               MapTypeFuBen,
		SceneTypeCrossGodSiege:           MapTypeActivity,
		SceneTypeUnrealBoss:              MapTypeBoss,
		SceneTypeMaterial:                MapTypeFuBen,
		SceneTypeOutlandBoss:             MapTypeBoss,
		SceneTypeBaGuaMiJing:             MapTypeFuBen,
		SceneTypeGuideReplica:            MapTypeFuBen,
		SceneTypeCrossTeamCopy:           MapTypeFuBen,
		SceneTypeCrossDenseWat:           MapTypeActivity,
		SceneTypeCrossShenMo:             MapTypeActivity,
		SceneTypeCangJingGe:              MapTypeBoss,
		SceneTypeAllianceBoss:            MapTypeFuBen,
		SceneTypeAllianceShengTan:        MapTypeActivityFuBen,
		SceneTypeFuQiFuBen:               MapTypeFuBen,
		SceneTypeShenYu:                  MapTypeActivity,
		SceneTypeXianTaoDaHui:            MapTypeActivity,
		SceneTypeLongGong:                MapTypeActivity,
		SceneTypeYuXi:                    MapTypeActivity,
		SceneTypeGuideReplicaMoJian:      MapTypeFuBen,
		SceneTypeGuideReplicaCatDog:      MapTypeFuBen,
		SceneTypeGuideReplicaRescue:      MapTypeFuBen,
		SceneTypeWelfareQiYu:             MapTypeActivityFuBen,
		SceneTypeLocalLianYu:             MapTypeActivity,
		SceneTypeArenapvpHaiXuan:         MapTypeArena,
		SceneTypeChuangShiZhiZhanZhongLi: MapTypeChuangShiWorld,
		SceneTypeArenapvp:                MapTypeArena,
		SceneTypeChuangShiZhiZhanMain:    MapTypeChuangShiWorld,
		SceneTypeChuangShiZhiZhanFuShu:   MapTypeActivityFuBen,
		SceneTypeZhenXi:                  MapTypeBoss,
		SceneTypeDingShi:                 MapTypeBoss,
	}
)

func (t SceneType) MapType() MapType {
	return sceneTypeMap[t]
}

var (
	sceneTypeServerTypeMap = map[SceneType]centertypes.GameServerType{
		SceneTypeWorld:                   centertypes.GameServerTypeSingle,
		SceneTypeFuBenSoulRuins:          centertypes.GameServerTypeSingle,
		SceneTypeEightClock:              centertypes.GameServerTypeSingle,
		SceneTypeTianJieTa:               centertypes.GameServerTypeSingle,
		SceneTypeChengZhan:               centertypes.GameServerTypeSingle,
		SceneTypeYinLiang:                centertypes.GameServerTypeSingle,
		SceneTypeYueXiaQingYuan:          centertypes.GameServerTypeSingle,
		SceneTypeExperience:              centertypes.GameServerTypeSingle,
		SceneTypeGold:                    centertypes.GameServerTypeSingle,
		SceneTypeHuangGong:               centertypes.GameServerTypeSingle,
		SceneTypeArena:                   centertypes.GameServerTypeRegion,
		SceneTypeArenaShengShou:          centertypes.GameServerTypeRegion,
		SceneTypeFourGodGate:             centertypes.GameServerTypeSingle,
		SceneTypeFourGodWar:              centertypes.GameServerTypeSingle,
		SceneTypeMajor:                   centertypes.GameServerTypeSingle,
		SceneTypeCrossWorldBoss:          centertypes.GameServerTypeRegion,
		SceneTypeCrossTuLong:             centertypes.GameServerTypeRegion,
		SceneTypeCrossLianYu:             centertypes.GameServerTypeRegion,
		SceneTypeTower:                   centertypes.GameServerTypeSingle,
		SceneTypeMyBoss:                  centertypes.GameServerTypeSingle,
		SceneTypeMyBossVip:               centertypes.GameServerTypeSingle,
		SceneTypeCrossGodSiege:           centertypes.GameServerTypeRegion,
		SceneTypeMaterial:                centertypes.GameServerTypeSingle,
		SceneTypeBaGuaMiJing:             centertypes.GameServerTypeSingle,
		SceneTypeGuideReplica:            centertypes.GameServerTypeSingle,
		SceneTypeCrossTeamCopy:           centertypes.GameServerTypeSingle,
		SceneTypeCrossDenseWat:           centertypes.GameServerTypeRegion,
		SceneTypeCrossShenMo:             centertypes.GameServerTypePlatform,
		SceneTypeAllianceBoss:            centertypes.GameServerTypeSingle,
		SceneTypeAllianceShengTan:        centertypes.GameServerTypeSingle,
		SceneTypeFuQiFuBen:               centertypes.GameServerTypeSingle,
		SceneTypeShenYu:                  centertypes.GameServerTypeSingle,
		SceneTypeXianTaoDaHui:            centertypes.GameServerTypeSingle,
		SceneTypeLongGong:                centertypes.GameServerTypeSingle,
		SceneTypeYuXi:                    centertypes.GameServerTypeSingle,
		SceneTypeGuideReplicaMoJian:      centertypes.GameServerTypeSingle,
		SceneTypeGuideReplicaCatDog:      centertypes.GameServerTypeSingle,
		SceneTypeGuideReplicaRescue:      centertypes.GameServerTypeSingle,
		SceneTypeWelfareQiYu:             centertypes.GameServerTypeSingle,
		SceneTypeLocalLianYu:             centertypes.GameServerTypeSingle,
		SceneTypeArenapvpHaiXuan:         centertypes.GameServerTypeAll,
		SceneTypeArenapvp:                centertypes.GameServerTypeAll,
		SceneTypeChuangShiZhiZhanZhongLi: centertypes.GameServerTypeAll,
		SceneTypeChuangShiZhiZhanMain:    centertypes.GameServerTypeAll,
		SceneTypeChuangShiZhiZhanFuShu:   centertypes.GameServerTypeAll,
		SceneTypeZhenXi:                  centertypes.GameServerTypeRegion,
		SceneTypeDingShi:                 centertypes.GameServerTypeSingle,
	}
)

func (t SceneType) GameServerType() centertypes.GameServerType {
	return sceneTypeServerTypeMap[t]
}

var (
	sceneTypeStringMap = map[SceneType]string{
		SceneTypeWorld:                   "世界",
		SceneTypeFuBenSoulRuins:          "帝陵遗迹副本",
		SceneTypeEightClock:              "8点半",
		SceneTypeTianJieTa:               "天劫塔",
		SceneTypeWorldBoss:               "世界boss",
		SceneTypeMaterial:                "材料副本",
		SceneTypeChengZhan:               "城战",
		SceneTypeYinLiang:                "银两副本",
		SceneTypeYueXiaQingYuan:          "月下情缘",
		SceneTypeExperience:              "经验副本",
		SceneTypeGold:                    "元宝副本",
		SceneTypeHuangGong:               "城战-皇宫",
		SceneTypeArena:                   "3v3",
		SceneTypeArenaShengShou:          "圣兽",
		SceneTypeFourGodGate:             "四神-入口",
		SceneTypeFourGodWar:              "四神",
		SceneTypeMarry:                   "结婚",
		SceneTypeLingChiFighting:         "1v1",
		SceneTypeCrossWorldBoss:          "跨服世界boss",
		SceneTypeCrossTuLong:             "跨服屠龙",
		SceneTypeCrossLianYu:             "杀戮之都",
		SceneTypeBeatBackMoJie:           "反击魔界",
		SceneTypeTower:                   "打宝塔",
		SceneTypeMyBoss:                  "个人BOSS",
		SceneTypeMyBossVip:               "VIP个人BOSS",
		SceneTypeCrossGodSiege:           "神兽攻城",
		SceneTypeUnrealBoss:              "幻境BOSS",
		SceneTypeOutlandBoss:             "外域BOSS",
		SceneTypeBaGuaMiJing:             "八卦秘境",
		SceneTypeGuideReplica:            "引导副本",
		SceneTypeCrossTeamCopy:           "组队副本",
		SceneTypeCrossDenseWat:           "金银密窟",
		SceneTypeCrossShenMo:             "神魔战场",
		SceneTypeCangJingGe:              "藏经阁",
		SceneTypeAllianceBoss:            "仙盟boss圣坛",
		SceneTypeAllianceShengTan:        "仙盟圣坛",
		SceneTypeFuQiFuBen:               "夫妻副本",
		SceneTypeShenYu:                  "神域之战",
		SceneTypeXianTaoDaHui:            "仙桃大会",
		SceneTypeLongGong:                "龙宫探宝",
		SceneTypeYuXi:                    "玉玺之战",
		SceneTypeGuideReplicaMoJian:      "魔剑副本",
		SceneTypeGuideReplicaCatDog:      "猫狗副本",
		SceneTypeGuideReplicaRescue:      "救援副本",
		SceneTypeWelfareQiYu:             "奇遇岛副本",
		SceneTypeLocalLianYu:             "无间炼狱",
		SceneTypeArenapvpHaiXuan:         "比武大会海选",
		SceneTypeArenapvp:                "比武大会",
		SceneTypeChuangShiZhiZhanZhongLi: "创世之战中立城",
		SceneTypeChuangShiZhiZhanMain:    "创世之战主城",
		SceneTypeChuangShiZhiZhanFuShu:   "创世之战附属城",
		SceneTypeZhenXi:                  "珍惜",
	}
)

func (t SceneType) String() string {
	return sceneTypeStringMap[t]
}

var (
	sceneEnterDistanceMap = map[SceneType]float64{
		SceneTypeFuBenSoulRuins:     10000,
		SceneTypeTianJieTa:          10000,
		SceneTypeYinLiang:           10000,
		SceneTypeMaterial:           10000,
		SceneTypeExperience:         10000,
		SceneTypeGold:               10000,
		SceneTypeArena:              18,
		SceneTypeLingChiFighting:    10000,
		SceneTypeMajor:              10000,
		SceneTypeCrossTuLong:        18,
		SceneTypeCrossTeamCopy:      10000,
		SceneTypeAllianceShengTan:   10000,
		SceneTypeFuQiFuBen:          10000,
		SceneTypeGuideReplicaMoJian: 10000,
		SceneTypeGuideReplicaCatDog: 10000,
		SceneTypeGuideReplicaRescue: 10000,
	}
	sceneExitDistanceMap = map[SceneType]float64{
		SceneTypeFuBenSoulRuins:     10000,
		SceneTypeTianJieTa:          10000,
		SceneTypeYinLiang:           10000,
		SceneTypeMaterial:           10000,
		SceneTypeExperience:         10000,
		SceneTypeGold:               10000,
		SceneTypeArena:              20,
		SceneTypeLingChiFighting:    10000,
		SceneTypeMajor:              10000,
		SceneTypeCrossTuLong:        20,
		SceneTypeCrossTeamCopy:      10000,
		SceneTypeAllianceShengTan:   10000,
		SceneTypeFuQiFuBen:          10000,
		SceneTypeGuideReplicaMoJian: 10000,
		SceneTypeGuideReplicaCatDog: 10000,
		SceneTypeGuideReplicaRescue: 10000,
	}
	defaultEnterDistance = float64(13)
	defaultExitDistance  = float64(15)
)

func (t SceneType) GetEnterDistance() float64 {
	enterDistance, ok := sceneEnterDistanceMap[t]
	if !ok {
		return defaultEnterDistance
	}
	return enterDistance
}

func (t SceneType) GetExitDistance() float64 {
	exitDistance, ok := sceneExitDistanceMap[t]
	if !ok {
		return defaultExitDistance
	}
	return exitDistance
}

var (
	//断线重连进入场景
	sceneReConnectMap = map[SceneType]bool{
		SceneTypeYuXi: false,
	}

	// 默认进入
	defaultConnect = true
)

func (t SceneType) IsReConnect() bool {
	isReConnect, ok := sceneReConnectMap[t]
	if !ok {
		return defaultConnect
	}
	return isReConnect
}

func (t SceneType) ToActivityType() (act activitytypes.ActivityType, flag bool) {
	act, flag = sceneActiveTypeMap[t]
	return
}

var (
	//本服活动映射
	sceneActiveTypeMap = map[SceneType]activitytypes.ActivityType{
		SceneTypeYueXiaQingYuan:        activitytypes.ActivityTypeMoonLove,
		SceneTypeFourGodGate:           activitytypes.ActivityTypeFourGod,
		SceneTypeFourGodWar:            activitytypes.ActivityTypeFourGod,
		SceneTypeChengZhan:             activitytypes.ActivityTypeAlliance,
		SceneTypeHuangGong:             activitytypes.ActivityTypeAlliance,
		SceneTypeAllianceShengTan:      activitytypes.ActivityTypeAllianceShengTan,
		SceneTypeShenYu:                activitytypes.ActivityTypeShenYu,
		SceneTypeXianTaoDaHui:          activitytypes.ActivityTypeXianTaoDaHui,
		SceneTypeLongGong:              activitytypes.ActivityTypeLongGong,
		SceneTypeCrossLianYu:           activitytypes.ActivityTypeLianYu,
		SceneTypeLocalLianYu:           activitytypes.ActivityTypeLocalLianYu,
		SceneTypeCrossGodSiege:         activitytypes.ActivityTypeGodSiegeQiLin, // TODO:xzk25 临时处理，修改 支持本服活动
		SceneTypeCrossShenMo:           activitytypes.ActivityTypeShenMoWar,     // TODO:xzk25 临时处理，修改 支持本服活动
		SceneTypeYuXi:                  activitytypes.ActivityTypeYuXi,
		SceneTypeArenapvpHaiXuan:       activitytypes.ActivityTypeArenapvp,
		SceneTypeChuangShiZhiZhanFuShu: activitytypes.ActivityTypeChuangShiZhiZhan,
	}
)

//
func (t SceneType) ToCrossType() (corssType crosstypes.CrossType, flag bool) {
	corssType, flag = toCrossTypeMap[t]
	return
}

var (
	toCrossTypeMap = map[SceneType]crosstypes.CrossType{
		SceneTypeArenapvpHaiXuan:         crosstypes.CrossTypeArenapvp,
		SceneTypeChuangShiZhiZhanZhongLi: crosstypes.CrossTypeChuangShi,
		SceneTypeChuangShiZhiZhanMain:    crosstypes.CrossTypeChuangShi,
		SceneTypeChuangShiZhiZhanFuShu:   crosstypes.CrossTypeChuangShi,
	}
)

//

var (
	//挂机类型
	sceneGuaJiTypeMap = map[SceneType]GuaJiType{
		SceneTypeFuBenSoulRuins:     GuaJiTypeFuBen,
		SceneTypeTianJieTa:          GuaJiTypeFuBen,
		SceneTypeMaterial:           GuaJiTypeFuBen,
		SceneTypeYinLiang:           GuaJiTypeFuBen,
		SceneTypeYueXiaQingYuan:     GuaJiTypeMoonLove,
		SceneTypeExperience:         GuaJiTypeFuBen,
		SceneTypeGold:               GuaJiTypeFuBen,
		SceneTypeFourGodGate:        GuaJiTypeFourGodGate,
		SceneTypeFourGodWar:         GuaJiTypeFourGodWar,
		SceneTypeMajor:              GuaJiTypeFuBen,
		SceneTypeBaGuaMiJing:        GuaJiTypeFuBen,
		SceneTypeGuideReplica:       GuaJiTypeFuBen,
		SceneTypeTower:              GuaJiTypeTowerScene,
		SceneTypeCrossLianYu:        GuaJiTypeLianYuScene,
		SceneTypeCrossGodSiege:      GuaJiTypeGodSiegeScene,
		SceneTypeChengZhan:          GuaJiTypeChengWai,
		SceneTypeHuangGong:          GuaJiTypeHuangGong,
		SceneTypeFuQiFuBen:          GuaJiTypeFuBen,
		SceneTypeYuXi:               GuaJiTypeYuXi,
		SceneTypeGuideReplicaMoJian: GuaJiTypeFuBen,
		SceneTypeGuideReplicaCatDog: GuaJiTypeFuBen,
		SceneTypeGuideReplicaRescue: GuaJiTypeFuBen,
		SceneTypeWelfareQiYu:        GuaJiTypeFuBen,
		SceneTypeCrossShenMo:        GuaJiTypeShenMo,
		SceneTypeCrossDenseWat:      GuaJiTypeDenseWat,
		SceneTypeAllianceShengTan:   GuaJiTypeShengTan,
		SceneTypeXianTaoDaHui:       GuaJiTypeXianTao,
		SceneTypeLongGong:           GuaJiTypeLongGong,
		SceneTypeCrossTuLong:        GuaJiTypeTulong,
	}
)

func (t SceneType) GetGuaJiType() (typ GuaJiType, flag bool) {
	typ, flag = sceneGuaJiTypeMap[t]
	return
}

//进入场景方式
type SceneEnterType int32

const (
	SceneEnterTypeCommon SceneEnterType = iota //通用
	SceneEnterTypeTrac                         //追踪
	SceneEnterTypePortal                       //传送阵
)

func (t SceneEnterType) Mask() int64 {
	return 1 << uint(t)
}
