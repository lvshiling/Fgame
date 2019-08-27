package types

//属性作用器效果
type PropertyEffectorType uint32

const (
	//初始化
	PlayerPropertyEffectorTypeInit PropertyEffectorType = iota
	//等级
	PlayerPropertyEffectorTypeLevel
	//技能
	PlayerPropertyEffectorTypeSkill
	//丹药
	PlayerPropertyEffectorTypeDan
	//坐骑
	PlayerPropertyEffectorTypeMount
	//战翼
	PlayerPropertyEffectorTypeWing
	//护体盾
	PlayerPropertyEffectorTypeBodyShield
	//称号
	PlayerPropertyEffectorTypeTitle
	//时装
	PlayerPropertyEffectorTypeFashion
	//兵魂
	PlayerPropertyEffectorTypeWeapon
	//装备
	PlayerPropertyEffectorTypeEquipment //10
	//帝魂
	PlayerPropertyEffectorTypeSoul
	//境界
	PlayerPropertyEffectorTypeRealm
	//帝王加成
	PlayerPropertyEffectorTypeEmperor
	//仙盟虎符加成
	PlayerPropertyEffectorTypeHuFu
	//神龙现世
	PlayerPropertyEffectorTypeDragon
	//领域
	PlayerPropertyEffectorTypeLingyu
	//身法
	PlayerPropertyEffectorTypeShenfa
	//斗神领域技能光环
	PlayerPropertyEffectorTypeLingyuAura
	//结婚
	PlayerPropertyEffectorTypeMarry
	//斗神金装固定属性
	PlayerPropertyEffectorTypeGoldequip //20
	//神盾尖刺
	PlayerPropertyEffectorTypeShield
	//护体仙羽
	PlayerPropertyEffectorTypeFeather
	//暗器
	PlayerPropertyEffectorTypeAnqi
	//转生
	PlayerPropertyEffectorTypeZhuanSheng
	//元神等级
	PlayerPropertyEffectorTypeGoldYuan
	//vip等级
	PlayerPropertyEffectorTypeVipLevel
	//戮仙刃
	PlayerPropertyEffectorTypeMassacre
	//天书
	PlayerPropertyEffectorTypeTianShu
	//法宝
	PlayerPropertyEffectorTypeFaBao
	//血盾
	PlayerPropertyEffectorTypeXueDun //30
	//仙体
	PlayerPropertyEffectorTypeXianTi
	//点星系统
	PlayerPropertyEffectorTypeDianXing
	//衣橱
	PlayerPropertyEffectorTypeWardrobe
	//噬魂幡系统
	PlayerPropertyEffectorTypeShiHunFan
	//天魔体
	PlayerPropertyEffectorTypeTianMoTi
	//玩家灵童
	PlayerPropertyEffectorTypeLingTong
	//玩家灵童时装
	PlayerPropertyEffectorTypeLingTongFashion
	//玩家灵童_灵宝
	PlayerPropertyEffectorTypeLingTongFaBao
	//玩家灵童_灵兵
	PlayerPropertyEffectorTypeLingTongWeapon
	//玩家灵童_灵骑
	PlayerPropertyEffectorTypeLingTongMount ////40
	//玩家灵童_灵身
	PlayerPropertyEffectorTypeLingTongShenFa
	//玩家灵童_灵体
	PlayerPropertyEffectorTypeLingTongXianTi
	//玩家灵童_灵翼
	PlayerPropertyEffectorTypeLingTongWing
	//玩家灵童_灵域
	PlayerPropertyEffectorTypeLingTongLingYu
	//玩家飞升
	PlayerPropertyEffectorTypeFeiShen
	//玩家至尊称号
	PlayerPropertyEffectorTypeSupremeTitle
	//圣痕
	PlayerPropertyEffectorTypeShengHen
	//屠龙装备
	PlayerPropertyEffectorTypeTuLongEquip
	//命格
	PlayerPropertyEffectorTypeMingGe
	//神器
	PlayerPropertyEffectorTypeShenQi //50
	//阵法
	PlayerPropertyEffectorTypeZhenFa
	//英灵谱
	PlayerPropertyEffectorTypeYingLingPu
	//宝宝
	PlayerPropertyEffectorTypeBaby
	// 泣血枪
	PlayerPropertyEffectorTypeQiXue
	// 结义-------------55
	PlayerPropertyEffectorTypeJieYi
	//无双神器
	PlayerPropertyEffectorTypeWushuangWeapon
	//大力丸
	PlayerPropertyEffectorTypeDaLiWan
	//创世系统
	PlayerPropertyEffectorTypeChuangShi
	//特戒
	PlayerPropertyEffectorTypeRing
	//上古之灵
	PlayerPropertyEffectorTypeShangGuZhiLing
)

func (pet PropertyEffectorType) EffectorType() uint32 {
	return uint32(pet)
}

var (
	propertyEffectoryTypeMap = map[PropertyEffectorType]string{
		PlayerPropertyEffectorTypeInit:            "初始化",
		PlayerPropertyEffectorTypeLevel:           "等级",
		PlayerPropertyEffectorTypeSkill:           "技能",
		PlayerPropertyEffectorTypeDan:             "丹药",
		PlayerPropertyEffectorTypeMount:           "坐骑",
		PlayerPropertyEffectorTypeWing:            "战翼",
		PlayerPropertyEffectorTypeBodyShield:      "护体盾",
		PlayerPropertyEffectorTypeTitle:           "称号",
		PlayerPropertyEffectorTypeFashion:         "时装",
		PlayerPropertyEffectorTypeWeapon:          "兵魂",
		PlayerPropertyEffectorTypeEquipment:       "装备",
		PlayerPropertyEffectorTypeSoul:            "帝魂",
		PlayerPropertyEffectorTypeRealm:           "境界",
		PlayerPropertyEffectorTypeEmperor:         "帝王加成",
		PlayerPropertyEffectorTypeHuFu:            "虎符加成",
		PlayerPropertyEffectorTypeDragon:          "神龙现世",
		PlayerPropertyEffectorTypeShenfa:          "身法",
		PlayerPropertyEffectorTypeLingyu:          "领域",
		PlayerPropertyEffectorTypeLingyuAura:      "斗神领域技能光环",
		PlayerPropertyEffectorTypeMarry:           "结婚",
		PlayerPropertyEffectorTypeGoldequip:       "斗神金装固定属性",
		PlayerPropertyEffectorTypeShield:          "神盾尖刺",
		PlayerPropertyEffectorTypeFeather:         "护体仙羽",
		PlayerPropertyEffectorTypeAnqi:            "暗器",
		PlayerPropertyEffectorTypeZhuanSheng:      "转生",
		PlayerPropertyEffectorTypeGoldYuan:        "元神等级",
		PlayerPropertyEffectorTypeVipLevel:        "vip等级",
		PlayerPropertyEffectorTypeMassacre:        "戮仙刃",
		PlayerPropertyEffectorTypeTianShu:         "天书",
		PlayerPropertyEffectorTypeFaBao:           "法宝",
		PlayerPropertyEffectorTypeXueDun:          "血盾",
		PlayerPropertyEffectorTypeXianTi:          "仙体",
		PlayerPropertyEffectorTypeDianXing:        "点星系统",
		PlayerPropertyEffectorTypeWardrobe:        "衣橱",
		PlayerPropertyEffectorTypeShiHunFan:       "噬魂幡系统",
		PlayerPropertyEffectorTypeTianMoTi:        "天魔体",
		PlayerPropertyEffectorTypeLingTong:        "玩家灵童",
		PlayerPropertyEffectorTypeLingTongFashion: "玩家灵童时装",
		PlayerPropertyEffectorTypeLingTongFaBao:   "玩家灵童_灵宝",
		PlayerPropertyEffectorTypeLingTongWeapon:  "玩家灵童_灵兵",
		PlayerPropertyEffectorTypeLingTongMount:   "玩家灵童_灵骑",
		PlayerPropertyEffectorTypeLingTongShenFa:  "玩家灵童_灵身",
		PlayerPropertyEffectorTypeLingTongXianTi:  "玩家灵童_灵体",
		PlayerPropertyEffectorTypeLingTongWing:    "玩家灵童_灵翼",
		PlayerPropertyEffectorTypeLingTongLingYu:  "玩家灵童_领域",
		PlayerPropertyEffectorTypeFeiShen:         "玩家飞升",
		PlayerPropertyEffectorTypeSupremeTitle:    "玩家至尊称号",
		PlayerPropertyEffectorTypeShengHen:        "玩家圣痕",
		PlayerPropertyEffectorTypeMingGe:          "玩家命格",
		PlayerPropertyEffectorTypeShenQi:          "玩家神器",
		PlayerPropertyEffectorTypeZhenFa:          "玩家阵法",
		PlayerPropertyEffectorTypeYingLingPu:      "英灵谱",
		PlayerPropertyEffectorTypeBaby:            "宝宝",
		PlayerPropertyEffectorTypeQiXue:           "泣血枪",
		PlayerPropertyEffectorTypeJieYi:           "结义",
		PlayerPropertyEffectorTypeWushuangWeapon:  "无双神器",
		PlayerPropertyEffectorTypeDaLiWan:         "大力丸",
		PlayerPropertyEffectorTypeChuangShi:       "创世系统",
		PlayerPropertyEffectorTypeRing:            "特戒",
		PlayerPropertyEffectorTypeShangGuZhiLing:  "上古之灵",
	}
)

func (pet PropertyEffectorType) String() string {
	return propertyEffectoryTypeMap[pet]
}

func (pet PropertyEffectorType) Mask() uint64 {
	return 1 << uint(pet)
}

//所有作用器
const (
	PropertyEffectorTypeMaskAll = 1<<63 - 1
)

func GetPropertyEffectoryTypeMap() map[PropertyEffectorType]string {
	return propertyEffectoryTypeMap
}

func StringForMask(mask uint64) string {
	if mask == PropertyEffectorTypeMaskAll {
		return "重新计算"
	}
	s := ""
	for t := PlayerPropertyEffectorTypeInit; t <= PropertyEffectorType(63); t++ {
		if t.Mask()&mask != 0 {
			s += t.String()
			s += ","
		}
	}
	return s
}
