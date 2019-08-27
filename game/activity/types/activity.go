package types

type ActivityType int32

const (
	ActivityTypeNone               ActivityType = 0
	ActivityTypeMoonLove                        = 1
	ActivityTypeAlliance                        = 2
	ActivityTypeFourGod                         = 3
	ActivityTypeArena                           = 4
	ActivityTypeBiaoChe                         = 5
	ActivityTypeWorldBoss                       = 6
	ActivityTypeCoressTuLong                    = 7
	ActivityTypeXueKuangCollect                 = 8
	ActivityTypeShangCheng                      = 9
	ActivityTypeLianYu                          = 10
	ActivityTypeDaBaoTower                      = 11
	ActivityTypeGodSiegeQiLin                   = 12
	ActivityTypeGodSiegeHuoFeng                 = 13
	ActivityTypeGodSiegeDuLong                  = 14
	ActivityTypeDenseWat                        = 15
	ActivityTypeShenMoWar                       = 16
	ActivityTypeAllianceShengTan                = 19
	ActivityTypeXianTaoDaHui                    = 20
	ActivityTypeShenYu                          = 21
	ActivityTypeLocalLianYu                     = 22
	ActivityTypeLocalGodSiegeQiLin              = 23
	ActivityTypeLocalShenMoWar                  = 24
	ActivityTypeLongGong                        = 25
	ActivityTypeYuXi                            = 26
	ActivityTypeArenapvp                        = 27
	ActivityTypeChuangShiZhiZhan                = 28
)

var (
	activityMap = map[ActivityType]string{
		ActivityTypeMoonLove:           "月下情缘",
		ActivityTypeAlliance:           "九霄城战",
		ActivityTypeFourGod:            "四神",
		ActivityTypeArena:              "3v3竞技场",
		ActivityTypeBiaoChe:            "护送镖车",
		ActivityTypeWorldBoss:          "魔界巢穴",
		ActivityTypeCoressTuLong:       "跨服屠龙",
		ActivityTypeXueKuangCollect:    "血矿采集",
		ActivityTypeShangCheng:         "活动商城",
		ActivityTypeLianYu:             "无间炼狱",
		ActivityTypeDaBaoTower:         "打宝塔",
		ActivityTypeGodSiegeQiLin:      "神兽攻城-麒麟来袭",
		ActivityTypeGodSiegeHuoFeng:    "神兽攻城-火凤来袭",
		ActivityTypeGodSiegeDuLong:     "神兽攻城-毒龙来袭",
		ActivityTypeDenseWat:           "金银密窟",
		ActivityTypeShenMoWar:          "神魔战场",
		ActivityTypeAllianceShengTan:   "仙盟圣坛",
		ActivityTypeXianTaoDaHui:       "仙桃大会",
		ActivityTypeShenYu:             "神域之战",
		ActivityTypeLocalLianYu:        "无间炼狱（本服）",
		ActivityTypeLocalGodSiegeQiLin: "神兽攻城-麒麟来袭（本服）",
		ActivityTypeLocalShenMoWar:     "神魔战场（本服）",
		ActivityTypeLongGong:           "龙宫探宝",
		ActivityTypeYuXi:               "玉玺之战",
		ActivityTypeArenapvp:           "比武大会",
		ActivityTypeChuangShiZhiZhan:   "创世之战",
	}
)

func (t ActivityType) Valid() bool {
	switch t {
	case ActivityTypeAlliance,
		ActivityTypeMoonLove,
		ActivityTypeFourGod,
		ActivityTypeArena,
		ActivityTypeBiaoChe,
		ActivityTypeWorldBoss,
		ActivityTypeCoressTuLong,
		ActivityTypeXueKuangCollect,
		ActivityTypeShangCheng,
		ActivityTypeLianYu,
		ActivityTypeDaBaoTower,
		ActivityTypeGodSiegeQiLin,
		ActivityTypeGodSiegeHuoFeng,
		ActivityTypeGodSiegeDuLong,
		ActivityTypeDenseWat,
		ActivityTypeShenMoWar,
		ActivityTypeAllianceShengTan,
		ActivityTypeXianTaoDaHui,
		ActivityTypeShenYu,
		ActivityTypeLocalLianYu,
		ActivityTypeLocalGodSiegeQiLin,
		ActivityTypeLocalShenMoWar,
		ActivityTypeLongGong,
		ActivityTypeYuXi,
		ActivityTypeArenapvp,
		ActivityTypeChuangShiZhiZhan:
		return true
	}
	return false
}

func (t ActivityType) String() string {
	return activityMap[t]
}
