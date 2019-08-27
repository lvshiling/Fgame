package types

type DevicePlatformType int32

const (
	DevicePlatformTypeWindows DevicePlatformType = 1 + iota
	DevicePlatformTypeAndroid
	DevicePlatformTypeIOS
)

func (t DevicePlatformType) Valid() bool {
	switch t {
	case DevicePlatformTypeWindows,
		DevicePlatformTypeAndroid,
		DevicePlatformTypeIOS:
		return true
	}
	return false
}

var (
	deviceMap = map[DevicePlatformType]string{
		DevicePlatformTypeWindows: "windows",
		DevicePlatformTypeAndroid: "安卓",
		DevicePlatformTypeIOS:     "ios",
	}
)

func (t DevicePlatformType) String() string {
	return deviceMap[t]
}

type SDKType int32

const (
	SDKTypePC SDKType = 1 + iota
	SDKTypeHengGeWan
	SDKTypeFeiYang
	SDKTypeZhengFu
	SDKTypeFeiFan
	SDKTypeShunYou
	SDKTypeJuDu
	SDKTypeMengYuanWenXian
	SDKTypeLuoLiWan
	SDKTypeXinFeng //10
	SDKTypeXingYue
	SDKTypeBoCai
	SDKTypeXiXiYou
	SDKTypeQiA
	SDKTypeZuoWan
	SDKTypeQiLing
	SDKTypeTaiFeng
	SDKTypeBaJi
	SDKTypeZhuTianXing
	SDKTypeJiuLing // 20
	SDKTypeAoTian
	SDKTypePiaoMiao
	SDKTypeYeYuShengGe
	SDKTypeZhangJian
	SDKTypeXiongWei
	SDKTypeDiSui
	SDKTypeTianShen
	SDKTypeShenYu
	SDKTypeWenJian
	SDKTypeRuoZi //30
	SDKTypeJianDao
	SDKTypeChenXi
	SDKTypeTianXing
	SDKTypeTianJi
	SDKTypeYaoJing
	SDKTypeXiaKeXing
	SDKTypeSanJie
	SDKTypeYouMeng
	SDKTypeJiuMeng
	SDKTypeLongYu //40
	SDKTypeChenXi2
	SDKTypeXianFan
	SDKTypeJiangHu
	SDKTypeFengYuan
	SDKTypeQiLing2
	SDKTypeQiLing3
	SDKTypeMengHuan
	SDKTypeXiaoYao
	SDKTypeMengHuan2
	SDKTypeXiaoYaoJiangHu //50
	SDKTypeTaiGu
	SDKTypeYiYun
	SDKTypeWanJie
	SDKTypeLingMeng
	SDKTypeLieYan
	SDKTypeMingJian
	SDKTypeJiuTian
	SDKTypeQiYiYouZhangJian
	SDKTypeNiuChaYouFuTu
	SDKTypeMoFangYouXiFengMo //60
	SDKTypePaiQuYouXiXianLu
)

const (
	SDKTypeSelf SDKType = 999
)

var (
	SdkMap = map[SDKType]string{
		SDKTypePC:                "pc登陆",
		SDKTypeHengGeWan:         "亨哥玩",
		SDKTypeFeiYang:           "飞扬",
		SDKTypeZhengFu:           "征服",
		SDKTypeFeiFan:            "非凡",
		SDKTypeShunYou:           "顺游",
		SDKTypeJuDu:              "巨都",
		SDKTypeMengYuanWenXian:   "梦缘问仙",
		SDKTypeLuoLiWan:          "萝莉玩",
		SDKTypeXinFeng:           "新蜂",
		SDKTypeXingYue:           "星月",
		SDKTypeBoCai:             "菠菜",
		SDKTypeXiXiYou:           "嘻嘻游",
		SDKTypeQiA:               "7A",
		SDKTypeZuoWan:            "佐玩",
		SDKTypeSelf:              "自己平台",
		SDKTypeQiLing:            "启灵",
		SDKTypeTaiFeng:           "泰逢",
		SDKTypeBaJi:              "吧唧",
		SDKTypeZhuTianXing:       "起灵",
		SDKTypeJiuLing:           "九零",
		SDKTypeAoTian:            "傲天",
		SDKTypePiaoMiao:          "缥缈",
		SDKTypeYeYuShengGe:       "夜雨笙歌",
		SDKTypeZhangJian:         "仗剑",
		SDKTypeXiongWei:          "雄威大世",
		SDKTypeDiSui:             "地随世渊",
		SDKTypeTianShen:          "天神决游",
		SDKTypeShenYu:            "神域",
		SDKTypeWenJian:           "问剑",
		SDKTypeRuoZi:             "若紫言诺",
		SDKTypeJianDao:           "剑道问情",
		SDKTypeChenXi:            "晨曦仙道",
		SDKTypeTianXing:          "天行战歌",
		SDKTypeTianJi:            "天机",
		SDKTypeYaoJing:           "妖精仙魔录",
		SDKTypeXiaKeXing:         "侠客行",
		SDKTypeSanJie:            "三界天书",
		SDKTypeYouMeng:           "游梦江湖",
		SDKTypeJiuMeng:           "九梦天书",
		SDKTypeLongYu:            "龙语大陆",
		SDKTypeChenXi2:           "晨曦2",
		SDKTypeXianFan:           "仙凡",
		SDKTypeJiangHu:           "江湖",
		SDKTypeFengYuan:          "风渊",
		SDKTypeQiLing2:           "至昊网络-启灵",
		SDKTypeQiLing3:           "萌萌网络-启灵",
		SDKTypeMengHuan:          "6pgame-梦幻情缘",
		SDKTypeXiaoYao:           "好玩-逍遥无双",
		SDKTypeMengHuan2:         "26580-梦幻情缘",
		SDKTypeXiaoYaoJiangHu:    "火鸟游-逍遥江湖",
		SDKTypeTaiGu:             "i5sy-太古纪元",
		SDKTypeYiYun:             "齐齐乐-亦云间",
		SDKTypeWanJie:            "6亿游-万界奇谭",
		SDKTypeLingMeng:          "天启-灵梦仙界",
		SDKTypeLieYan:            "乐趣游-烈焰神尊",
		SDKTypeMingJian:          "枭娱游戏-鸣剑仙录",
		SDKTypeJiuTian:           "起灵-九天凌云",
		SDKTypeQiYiYouZhangJian:  "7亿游-仗剑凡尘",
		SDKTypeNiuChaYouFuTu:     "牛叉游-浮屠幻境",
		SDKTypeMoFangYouXiFengMo: "魔方游戏-封魔录",
		SDKTypePaiQuYouXiXianLu:  "派趣游戏-仙路争锋",
	}
)

func (t SDKType) String() string {
	return SdkMap[t]
}

func (t SDKType) Valid() bool {
	switch t {
	case SDKTypePC,
		SDKTypeHengGeWan,
		SDKTypeFeiYang,
		SDKTypeZhengFu,
		SDKTypeFeiFan,
		SDKTypeShunYou,
		SDKTypeJuDu,
		SDKTypeMengYuanWenXian,
		SDKTypeLuoLiWan,
		SDKTypeXinFeng,
		SDKTypeXingYue,
		SDKTypeBoCai,
		SDKTypeXiXiYou,
		SDKTypeQiA,
		SDKTypeZuoWan,
		SDKTypeQiLing,
		SDKTypeTaiFeng,
		SDKTypeBaJi,
		SDKTypeZhuTianXing,
		SDKTypeJiuLing,
		SDKTypeAoTian,
		SDKTypePiaoMiao,
		SDKTypeYeYuShengGe,
		SDKTypeZhangJian,
		SDKTypeTianShen,
		SDKTypeShenYu,
		SDKTypeWenJian,
		SDKTypeXiongWei,
		SDKTypeDiSui,
		SDKTypeRuoZi,
		SDKTypeJianDao,
		SDKTypeChenXi,
		SDKTypeTianXing,
		SDKTypeTianJi,
		SDKTypeYaoJing,
		SDKTypeXiaKeXing,
		SDKTypeSanJie,
		SDKTypeYouMeng,
		SDKTypeJiuMeng,
		SDKTypeLongYu,
		SDKTypeChenXi2,
		SDKTypeXianFan,
		SDKTypeJiangHu,
		SDKTypeFengYuan,
		SDKTypeQiLing2,
		SDKTypeQiLing3,
		SDKTypeMengHuan,
		SDKTypeXiaoYao,
		SDKTypeMengHuan2,
		SDKTypeXiaoYaoJiangHu,
		SDKTypeTaiGu,
		SDKTypeYiYun,
		SDKTypeWanJie,
		SDKTypeLingMeng,
		SDKTypeLieYan,
		SDKTypeMingJian,
		SDKTypeJiuTian,
		SDKTypeQiYiYouZhangJian,
		SDKTypeNiuChaYouFuTu,
		SDKTypeMoFangYouXiFengMo,
		SDKTypePaiQuYouXiXianLu:
		return true
	}
	return false
}
