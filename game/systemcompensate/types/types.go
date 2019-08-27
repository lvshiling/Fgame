package types

type SystemCompensateType int32

const (
	SystemCompensateTypeMount          SystemCompensateType = iota //坐骑
	SystemCompensateTypeWing                                       //战翼
	SystemCompensateTypeAnQi                                       //暗器
	SystemCompensateTypeFaBao                                      //法宝
	SystemCompensateTypeXianTi                                     //仙体
	SystemCompensateTypeLingYu                                     //领域
	SystemCompensateTypeShenFa                                     //身法
	SystemCompensateTypeShiHunFan                                  //噬魂幡
	SystemCompensateTypeTianMo                                     //天魔体
	SystemCompensateTypeBodyShield                                 //护体盾
	SystemCompensateTypeFeather                                    //仙羽
	SystemCompensateTypeShield                                     //盾刺
	SystemCompensateTypeLingTongWeapon                             //灵兵
	SystemCompensateTypeLingTongMount                              //灵骑
	SystemCompensateTypeLingTongWing                               //灵翼
	SystemCompensateTypeLingTongShenFa                             //灵身
	SystemCompensateTypeLingTongLingYu                             //灵域
	SystemCompensateTypeLingTongFaBao                              //灵宝
	SystemCompensateTypeLingTongXianTi                             //灵体
)

func (s SystemCompensateType) Valid() bool {
	switch s {
	case SystemCompensateTypeMount,
		SystemCompensateTypeWing,
		SystemCompensateTypeAnQi,
		SystemCompensateTypeFaBao,
		SystemCompensateTypeXianTi,
		SystemCompensateTypeLingYu,
		SystemCompensateTypeShenFa,
		SystemCompensateTypeShiHunFan,
		SystemCompensateTypeTianMo,
		SystemCompensateTypeBodyShield,
		SystemCompensateTypeFeather,
		SystemCompensateTypeShield,
		SystemCompensateTypeLingTongWeapon,
		SystemCompensateTypeLingTongMount,
		SystemCompensateTypeLingTongWing,
		SystemCompensateTypeLingTongShenFa,
		SystemCompensateTypeLingTongLingYu,
		SystemCompensateTypeLingTongFaBao,
		SystemCompensateTypeLingTongXianTi:
		return true
	default:
		return false
	}
}

const (
	MinSysCompensate = SystemCompensateTypeMount
	MaxSysCompensate = SystemCompensateTypeLingTongXianTi
)

var (
	systemCompensateTypeStringMap = map[SystemCompensateType]string{
		SystemCompensateTypeMount:          "坐骑系统",
		SystemCompensateTypeWing:           "战翼系统",
		SystemCompensateTypeAnQi:           "暗器系统",
		SystemCompensateTypeFaBao:          "法宝系统",
		SystemCompensateTypeXianTi:         "仙体系统",
		SystemCompensateTypeLingYu:         "领域系统",
		SystemCompensateTypeShenFa:         "身法系统",
		SystemCompensateTypeShiHunFan:      "噬魂幡系统",
		SystemCompensateTypeFeather:        "免爆仙羽系统",
		SystemCompensateTypeBodyShield:     "护体盾系统",
		SystemCompensateTypeShield:         "破格盾刺系统",
		SystemCompensateTypeTianMo:         "天魔体系统",
		SystemCompensateTypeLingTongWeapon: "灵兵系统",
		SystemCompensateTypeLingTongMount:  "灵骑系统",
		SystemCompensateTypeLingTongWing:   "灵翼系统",
		SystemCompensateTypeLingTongShenFa: "灵身系统",
		SystemCompensateTypeLingTongLingYu: "灵域系统",
		SystemCompensateTypeLingTongFaBao:  "灵宝系统",
		SystemCompensateTypeLingTongXianTi: "灵体系统",
	}
)

func (s SystemCompensateType) String() string {
	return systemCompensateTypeStringMap[s]
}
