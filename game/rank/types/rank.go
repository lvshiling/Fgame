package types

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/lingtongdev/types"
)

const (
	//每页显示10条
	PageLimit = 10
	//显示排行前100
	ListLimit = 100
)

type RankClassType int32

const (
	//本服排行
	RankClassTypeLocal RankClassType = iota
	//本区排行
	RankClassTypeArea
	//本服活动排行榜
	RankClassTypeLocalActivity
)

func (r RankClassType) RankRefreshTime() int64 {
	return rankRefreshTimeMap[r]
}

func (r RankClassType) RankRedisExpireTime() int64 {
	return rankRedisExpireTimeMap[r]
}

func (r RankClassType) GetRankTypeMap() map[RankType]string {
	return rankClassTypeMap[r]
}

type RankType int32

const (
	//战力排行榜
	RankTypeForce RankType = 1
	//帮派排行榜
	RankTypeGang RankType = 2
	//坐骑排行榜
	RankTypeMount RankType = 3
	//战翼排行榜
	RankTypeWing RankType = 4
	//兵魂排行榜
	RankTypeWeapon RankType = 5
	//护体盾排行榜
	RankTypeBodyShield RankType = 6
	//身法排行榜
	RankTypeShenFa RankType = 7
	//领域排行榜
	RankTypeLingYu RankType = 8
	//护体仙羽排行榜
	RankTypeFeather RankType = 9
	//神盾尖刺排行榜
	RankTypeShield RankType = 10
	//暗器排行榜
	RankTypeAnQi RankType = 11
	//魅力排行榜
	RankTypeCharm RankType = 12
	//次数排行榜
	RankTypeCount RankType = 13
	//法宝排行榜
	RankTypeFaBao RankType = 14
	//仙体排行榜
	RankTypeXianTi RankType = 15
	//等级排行榜
	RankTypeLevel RankType = 16
	//天魔体
	RankTypeTianMoTi RankType = 17
	//噬魂幡
	RankTypeShiHunFan RankType = 18
	//灵骑
	RankTypeLingQi RankType = 19
	//灵兵
	RankTypeLingBing RankType = 20
	//灵翼
	RankTypeLingYi RankType = 21
	//灵宝
	RankTypeLingBao RankType = 22
	//灵身
	RankTypeLingShen RankType = 23
	//灵域
	RankTypeLingTongYu RankType = 24
	//灵体
	RankTypeLingTi RankType = 25
	//灵童等级
	RankTypeLingTongLevel RankType = 26
	//飞升排行榜
	RankTypeFeiSheng RankType = 27
	//充值排行榜
	RankTypeCharge RankType = 50
	//消费排行榜
	RankTypeCost RankType = 51
	//表白排行榜
	RankTypeMarryDevelop RankType = 52
	//灵童战力排行榜
	RankTypeLingTongForce RankType = 53
	//元神金装战力排行榜
	RankTypeGoldEquipForce RankType = 54
	//点星战力排行榜
	RankTypeDianXingForce RankType = 55
	//神器战力排行榜
	RankTypeShenQiForce RankType = 56
	//命格战力排行榜
	RankTypeMingGeForce RankType = 57
	//圣痕战力排行榜
	RankTypeShengHenForce RankType = 58
	//阵法战力排行榜
	RankTypeZhenFaForce RankType = 59
	//屠龙装战力排行榜
	RankTypeTuLongEquipForce RankType = 60
	//宝宝战力排行榜
	RankTypeBabyForce RankType = 61
	//转生等级
	RankTypeZhuanSheng RankType = 62
)

func (rt RankType) Valid() bool {
	switch rt {
	case RankTypeForce,
		RankTypeGang,
		RankTypeMount,
		RankTypeWing,
		RankTypeWeapon,
		RankTypeBodyShield,
		RankTypeShenFa,
		RankTypeLingYu,
		RankTypeFeather,
		RankTypeShield,
		RankTypeAnQi,
		RankTypeCharm,
		RankTypeCount,
		RankTypeFaBao,
		RankTypeXianTi,
		RankTypeCharge,
		RankTypeCost,
		RankTypeLevel,
		RankTypeTianMoTi,
		RankTypeShiHunFan,
		RankTypeLingQi,
		RankTypeLingBing,
		RankTypeLingYi,
		RankTypeLingBao,
		RankTypeLingShen,
		RankTypeLingTongYu,
		RankTypeLingTi,
		RankTypeLingTongLevel,
		RankTypeFeiSheng,
		RankTypeMarryDevelop,
		RankTypeLingTongForce,
		RankTypeGoldEquipForce,
		RankTypeDianXingForce,
		RankTypeShenQiForce,
		RankTypeMingGeForce,
		RankTypeShengHenForce,
		RankTypeZhenFaForce,
		RankTypeTuLongEquipForce,
		RankTypeBabyForce,
		RankTypeZhuanSheng:
		return true
	}
	return false
}

func (rt RankType) String() string {
	return rankTypeMap[rt]
}

//排行榜刷新时间
const (
	rankRefreshTimeNormal   = int64(5 * common.MINUTE)
	rankRefreshTimeActivity = int64(1 * common.MINUTE)
)

var (
	//排行榜刷新时间
	rankRefreshTimeMap = map[RankClassType]int64{
		RankClassTypeArea:          rankRefreshTimeNormal,
		RankClassTypeLocal:         rankRefreshTimeNormal,
		RankClassTypeLocalActivity: rankRefreshTimeActivity,
	}
)

//redis 失效时间(单位是秒)
const (
	rankExprieTimeNormal   = int64(25 * common.MINUTE / 1000)
	rankExprieTimeActivity = int64(7 * common.MINUTE / 1000)
)

var (
	rankRedisExpireTimeMap = map[RankClassType]int64{
		RankClassTypeArea:          rankExprieTimeNormal,
		RankClassTypeLocal:         rankExprieTimeNormal,
		RankClassTypeLocalActivity: rankExprieTimeActivity,
	}
)

var (
	rankTypeMap = map[RankType]string{
		RankTypeForce:            "战力排行榜",
		RankTypeGang:             "帮派排行榜",
		RankTypeMount:            "坐骑排行榜",
		RankTypeWing:             "战翼排行榜",
		RankTypeWeapon:           "兵魂排行榜",
		RankTypeBodyShield:       "护体盾排行榜",
		RankTypeShenFa:           "身法排行榜",
		RankTypeLingYu:           "领域排行榜",
		RankTypeFeather:          "护体仙羽排行榜",
		RankTypeShield:           "神盾尖刺排行榜",
		RankTypeAnQi:             "暗器排行榜",
		RankTypeCharm:            "魅力排行榜",
		RankTypeCount:            "抽奖次数排行榜",
		RankTypeFaBao:            "法宝排行榜",
		RankTypeXianTi:           "仙体排行榜",
		RankTypeCharge:           "充值排行榜",
		RankTypeCost:             "消费排行榜",
		RankTypeLevel:            "等级排行榜",
		RankTypeTianMoTi:         "天魔体排行榜",
		RankTypeShiHunFan:        "噬魂幡排行榜",
		RankTypeLingQi:           "灵骑排行榜",
		RankTypeLingBing:         "灵兵排行榜",
		RankTypeLingYi:           "灵翼排行榜",
		RankTypeLingBao:          "灵宝排行榜",
		RankTypeLingShen:         "灵身排行榜",
		RankTypeLingTongYu:       "灵域排行榜",
		RankTypeLingTi:           "灵体排行榜",
		RankTypeLingTongLevel:    "灵童等级排行榜",
		RankTypeFeiSheng:         "飞升等级排行榜",
		RankTypeMarryDevelop:     "表白排行榜",
		RankTypeLingTongForce:    "灵童战力排行榜",
		RankTypeGoldEquipForce:   "元神金装战力排行榜",
		RankTypeDianXingForce:    "点星战力排行榜",
		RankTypeShenQiForce:      "神器战力排行榜",
		RankTypeMingGeForce:      "命格战力排行榜",
		RankTypeShengHenForce:    "圣痕战力排行榜",
		RankTypeZhenFaForce:      "阵法战力排行榜",
		RankTypeTuLongEquipForce: "屠龙装战力排行榜",
		RankTypeBabyForce:        "宝宝战力排行榜",
		RankTypeZhuanSheng:       "转生排行榜",
	}

	rankTypeLocalMap = map[RankType]string{
		RankTypeForce:         "战力排行榜",
		RankTypeGang:          "帮派排行榜",
		RankTypeMount:         "坐骑排行榜",
		RankTypeWing:          "战翼排行榜",
		RankTypeWeapon:        "兵魂排行榜",
		RankTypeBodyShield:    "护体盾排行榜",
		RankTypeShenFa:        "身法排行榜",
		RankTypeLingYu:        "领域排行榜",
		RankTypeFeather:       "护体仙羽排行榜",
		RankTypeShield:        "神盾尖刺排行榜",
		RankTypeAnQi:          "暗器排行榜",
		RankTypeFaBao:         "法宝排行榜",
		RankTypeXianTi:        "仙体排行榜",
		RankTypeTianMoTi:      "天魔体排行榜",
		RankTypeShiHunFan:     "噬魂幡排行榜",
		RankTypeLingQi:        "灵骑排行榜",
		RankTypeLingBing:      "灵兵排行榜",
		RankTypeLingYi:        "灵翼排行榜",
		RankTypeLingBao:       "灵宝排行榜",
		RankTypeLingShen:      "灵身排行榜",
		RankTypeLingTongYu:    "灵域排行榜",
		RankTypeLingTi:        "灵体排行榜",
		RankTypeLingTongLevel: "灵童等级排行榜",
		RankTypeFeiSheng:      "飞升等级排行榜",
	}
	rankTypeAreaMap = map[RankType]string{
		RankTypeForce: "战力排行榜",
		RankTypeGang:  "帮派排行榜",
	}

	rankClassTypeMap = map[RankClassType]map[RankType]string{
		RankClassTypeLocal:         rankTypeLocalMap,
		RankClassTypeArea:          rankTypeAreaMap,
		RankClassTypeLocalActivity: map[RankType]string{},
	}
)

var (
	rankTypeClassTypeMap = map[RankType]types.LingTongDevSysType{
		RankTypeForce:            types.LingTongDevSysTypeDefault,
		RankTypeGang:             types.LingTongDevSysTypeDefault,
		RankTypeMount:            types.LingTongDevSysTypeDefault,
		RankTypeWing:             types.LingTongDevSysTypeDefault,
		RankTypeWeapon:           types.LingTongDevSysTypeDefault,
		RankTypeBodyShield:       types.LingTongDevSysTypeDefault,
		RankTypeShenFa:           types.LingTongDevSysTypeDefault,
		RankTypeLingYu:           types.LingTongDevSysTypeDefault,
		RankTypeFeather:          types.LingTongDevSysTypeDefault,
		RankTypeShield:           types.LingTongDevSysTypeDefault,
		RankTypeAnQi:             types.LingTongDevSysTypeDefault,
		RankTypeCharm:            types.LingTongDevSysTypeDefault,
		RankTypeCount:            types.LingTongDevSysTypeDefault,
		RankTypeFaBao:            types.LingTongDevSysTypeDefault,
		RankTypeXianTi:           types.LingTongDevSysTypeDefault,
		RankTypeCharge:           types.LingTongDevSysTypeDefault,
		RankTypeCost:             types.LingTongDevSysTypeDefault,
		RankTypeLevel:            types.LingTongDevSysTypeDefault,
		RankTypeTianMoTi:         types.LingTongDevSysTypeDefault,
		RankTypeShiHunFan:        types.LingTongDevSysTypeDefault,
		RankTypeLingQi:           types.LingTongDevSysTypeLingQi,
		RankTypeLingBing:         types.LingTongDevSysTypeLingBing,
		RankTypeLingYi:           types.LingTongDevSysTypeLingYi,
		RankTypeLingBao:          types.LingTongDevSysTypeLingBao,
		RankTypeLingShen:         types.LingTongDevSysTypeLingShen,
		RankTypeLingTongYu:       types.LingTongDevSysTypeLingYu,
		RankTypeLingTi:           types.LingTongDevSysTypeLingTi,
		RankTypeLingTongLevel:    types.LingTongDevSysTypeDefault,
		RankTypeFeiSheng:         types.LingTongDevSysTypeDefault,
		RankTypeMarryDevelop:     types.LingTongDevSysTypeDefault,
		RankTypeLingTongForce:    types.LingTongDevSysTypeDefault,
		RankTypeGoldEquipForce:   types.LingTongDevSysTypeDefault,
		RankTypeDianXingForce:    types.LingTongDevSysTypeDefault,
		RankTypeShenQiForce:      types.LingTongDevSysTypeDefault,
		RankTypeMingGeForce:      types.LingTongDevSysTypeDefault,
		RankTypeShengHenForce:    types.LingTongDevSysTypeDefault,
		RankTypeZhenFaForce:      types.LingTongDevSysTypeDefault,
		RankTypeTuLongEquipForce: types.LingTongDevSysTypeDefault,
		RankTypeBabyForce:        types.LingTongDevSysTypeDefault,
		RankTypeZhuanSheng:       types.LingTongDevSysTypeDefault,
	}
)

func (rt RankType) GetLongTongDevType() types.LingTongDevSysType {
	return rankTypeClassTypeMap[rt]
}

var (
	lingTongDevTypeRankTypeMap = map[types.LingTongDevSysType]RankType{
		types.LingTongDevSysTypeLingQi:   RankTypeLingQi,
		types.LingTongDevSysTypeLingBing: RankTypeLingBing,
		types.LingTongDevSysTypeLingYi:   RankTypeLingYi,
		types.LingTongDevSysTypeLingBao:  RankTypeLingBao,
		types.LingTongDevSysTypeLingShen: RankTypeLingShen,
		types.LingTongDevSysTypeLingYu:   RankTypeLingTongYu,
		types.LingTongDevSysTypeLingTi:   RankTypeLingTi,
	}
)

func LongTongDevTypeToRankType(devType types.LingTongDevSysType) RankType {
	return lingTongDevTypeRankTypeMap[devType]
}

type MyRankReqType int32

const (
	//战力排名
	MyRankReqTypeForce MyRankReqType = 1 + iota
	//帮派排名
	MyRankReqTypeGang
	//坐骑排名
	MyRankReqTypeMount
	//战翼排名
	MyRankReqTypeWing
	//兵魂排名
	MyRankReqTypeWeapon
	//护体盾排名
	MyRankReqTypeBodyShield
	//身法排名
	MyRankReqTypeShenFa
	//领域排名
	MyRankReqTypeLingYu
	//护体仙羽排名
	MyRankReqTypeFeather
	//神盾尖刺排名
	MyRankReqTypeShield
	//暗器排名
	MyRankReqTypeAnQi
	//法宝排名
	MyRankReqTypeFaBao
	//仙体排名
	MyRankReqTypeXianTi
	//噬魂幡排名
	MyRankReqTypeShiHunFan
	//天魔体排名
	MyRankReqTypeTianMoTi
	//灵童等级排名
	MyRankReqTypeLingTongLevel
	//灵兵排名
	MyRankReqTypeLingTongWeapon
	//灵骑排名
	MyRankReqTypeLingTongMount
	//灵翼排名
	MyRankReqTypeLingTongWing
	//灵身排名
	MyRankReqTypeLingTongShenFa
	//灵域排名
	MyRankReqTypeLingTongLingYu
	//灵宝排名
	MyRankReqTypeLingTongFaBao
	//灵体排名
	MyRankReqTypeLingTongXianTi
	//飞升排名
	MyRankReqTypeFeiSheng
	//3v3连胜排行
	MyRankReqTypeArena
)

func (t MyRankReqType) Valid() bool {
	switch t {
	case MyRankReqTypeForce,
		MyRankReqTypeGang,
		MyRankReqTypeMount,
		MyRankReqTypeWing,
		MyRankReqTypeWeapon,
		MyRankReqTypeBodyShield,
		MyRankReqTypeShenFa,
		MyRankReqTypeLingYu,
		MyRankReqTypeFeather,
		MyRankReqTypeShield,
		MyRankReqTypeAnQi,
		MyRankReqTypeFaBao,
		MyRankReqTypeXianTi,
		MyRankReqTypeShiHunFan,
		MyRankReqTypeTianMoTi,
		MyRankReqTypeLingTongLevel,
		MyRankReqTypeLingTongWeapon,
		MyRankReqTypeLingTongMount,
		MyRankReqTypeLingTongWing,
		MyRankReqTypeLingTongShenFa,
		MyRankReqTypeLingTongLingYu,
		MyRankReqTypeLingTongFaBao,
		MyRankReqTypeLingTongXianTi,
		MyRankReqTypeFeiSheng,
		MyRankReqTypeArena:

		return true
	}
	return false
}

func (t MyRankReqType) GetRankType() RankType {
	return myRankReqTypeMap[t]
}

var (
	myRankReqStrMap = map[MyRankReqType]string{
		MyRankReqTypeForce:          "我的战力排名",
		MyRankReqTypeGang:           "我的帮派排名",
		MyRankReqTypeMount:          "我的坐骑排名",
		MyRankReqTypeWing:           "我的战翼排名",
		MyRankReqTypeWeapon:         "我的兵魂排名",
		MyRankReqTypeBodyShield:     "我的护体盾排名",
		MyRankReqTypeShenFa:         "我的身法排名",
		MyRankReqTypeLingYu:         "我的领域排名",
		MyRankReqTypeFeather:        "我的护体仙羽排名",
		MyRankReqTypeShield:         "我的神盾尖刺排名",
		MyRankReqTypeAnQi:           "我的暗器排名",
		MyRankReqTypeFaBao:          "我的法宝排名",
		MyRankReqTypeXianTi:         "我的仙体排名",
		MyRankReqTypeShiHunFan:      "我的噬魂幡排名",
		MyRankReqTypeTianMoTi:       "我的天魔体排名",
		MyRankReqTypeLingTongLevel:  "我的灵童等级排名",
		MyRankReqTypeLingTongWeapon: "我的灵兵排名",
		MyRankReqTypeLingTongMount:  "我的灵骑排名",
		MyRankReqTypeLingTongWing:   "我的灵翼排名",
		MyRankReqTypeLingTongShenFa: "我的灵身排名",
		MyRankReqTypeLingTongLingYu: "我的灵域排名",
		MyRankReqTypeLingTongFaBao:  "我的灵宝排名",
		MyRankReqTypeLingTongXianTi: "我的灵体排名",
		MyRankReqTypeFeiSheng:       "我的飞升排名",
		MyRankReqTypeArena:          "我的3V3连胜排名",
	}
)

var (
	myRankReqTypeMap = map[MyRankReqType]RankType{
		MyRankReqTypeForce:          RankTypeForce,
		MyRankReqTypeGang:           RankTypeGang,
		MyRankReqTypeMount:          RankTypeMount,
		MyRankReqTypeWing:           RankTypeWing,
		MyRankReqTypeWeapon:         RankTypeWeapon,
		MyRankReqTypeBodyShield:     RankTypeBodyShield,
		MyRankReqTypeShenFa:         RankTypeShenFa,
		MyRankReqTypeLingYu:         RankTypeLingYu,
		MyRankReqTypeFeather:        RankTypeFeather,
		MyRankReqTypeShield:         RankTypeShield,
		MyRankReqTypeAnQi:           RankTypeAnQi,
		MyRankReqTypeFaBao:          RankTypeFaBao,
		MyRankReqTypeXianTi:         RankTypeXianTi,
		MyRankReqTypeShiHunFan:      RankTypeShiHunFan,
		MyRankReqTypeTianMoTi:       RankTypeTianMoTi,
		MyRankReqTypeLingTongLevel:  RankTypeLingTongLevel,
		MyRankReqTypeLingTongWeapon: RankTypeLingBing,
		MyRankReqTypeLingTongMount:  RankTypeLingQi,
		MyRankReqTypeLingTongWing:   RankTypeLingYi,
		MyRankReqTypeLingTongShenFa: RankTypeLingShen,
		MyRankReqTypeLingTongLingYu: RankTypeLingYu,
		MyRankReqTypeLingTongFaBao:  RankTypeLingBao,
		MyRankReqTypeLingTongXianTi: RankTypeLingTi,
		MyRankReqTypeFeiSheng:       RankTypeFeiSheng,
	}
)
