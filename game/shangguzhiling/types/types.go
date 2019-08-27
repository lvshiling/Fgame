package types

// 灵兽类型
type LingshouType int32

const (
	LingshouTypeBaize     LingshouType = iota //白泽
	LingshouTypeKuiniu                        //夔牛
	LingshouTypeFenghuang                     //凤凰
	LingshouTypeQilin                         //麒麟
	LingshouTypeTaowu                         //梼杌
	LingshouTypeXiezhi                        //獬豸
	LingshouTypeBifang                        //毕方
	LingshouTypeTaotie                        //饕餮
	LingshouTypeZhulong                       //烛龙
	LingshouTypeBian                          //狴犴
)

func (t LingshouType) Valid() bool {
	switch t {
	case LingshouTypeBaize,
		LingshouTypeKuiniu,
		LingshouTypeFenghuang,
		LingshouTypeQilin,
		LingshouTypeTaowu,
		LingshouTypeXiezhi,
		LingshouTypeBifang,
		LingshouTypeTaotie,
		LingshouTypeZhulong,
		LingshouTypeBian:
		return true
	default:
		return false
	}
}

var lingshouMap = map[LingshouType]string{
	LingshouTypeBaize:     "白泽",
	LingshouTypeKuiniu:    "夔牛",
	LingshouTypeFenghuang: "凤凰",
	LingshouTypeQilin:     "麒麟",
	LingshouTypeTaowu:     "梼杌",
	LingshouTypeXiezhi:    "獬豸",
	LingshouTypeBifang:    "毕方",
	LingshouTypeTaotie:    "饕餮",
	LingshouTypeZhulong:   "烛龙",
	LingshouTypeBian:      "狴犴",
}

func (t LingshouType) String() string {
	return lingshouMap[t]
}

const (
	MinLingshouType = LingshouTypeBaize
	MaxLingshouType = LingshouTypeBian
)

// 灵纹类型
type LingwenType int32

const (
	LingwenTypeFeng LingwenType = iota
	LingwenTypeHuo
	LingwenTypeLei
	LingwenTypeDian
	LingwenTypeYu
	LingwenTypeGuang
	LingwenTypeAn
	LingwenTypeBing
	LingwenTypeTu
	LingwenTypeJin
)

func (t LingwenType) Valid() bool {
	switch t {
	case LingwenTypeFeng,
		LingwenTypeHuo,
		LingwenTypeLei,
		LingwenTypeDian,
		LingwenTypeYu,
		LingwenTypeGuang,
		LingwenTypeAn,
		LingwenTypeBing,
		LingwenTypeTu,
		LingwenTypeJin:
		return true
	default:
		return false
	}
}

var lingwenMap = map[LingwenType]string{
	LingwenTypeFeng:  "风",
	LingwenTypeHuo:   "火",
	LingwenTypeLei:   "雷",
	LingwenTypeDian:  "电",
	LingwenTypeYu:    "雨",
	LingwenTypeGuang: "光",
	LingwenTypeAn:    "暗",
	LingwenTypeBing:  "冰",
	LingwenTypeTu:    "土",
	LingwenTypeJin:   "金",
}

func (t LingwenType) String() string {
	return lingwenMap[t]
}

const (
	MinLingwenType = LingwenTypeFeng
	MaxLingwenType = LingwenTypeJin
)

//灵炼部位类型
type LinglianPosType int32

const (
	LinglianPosTypeOne LinglianPosType = iota
	LinglianPosTypeTwo
	LinglianPosTypeThree
	LinglianPosTypeFour
	LinglianPosTypeFive
	LinglianPosTypeSix
	LinglianPosTypeSeven
	LinglianPosTypeEight
	LinglianPosTypeNine
	LinglianPosTypeTen
)

func (t LinglianPosType) Valid() bool {
	switch t {
	case LinglianPosTypeOne,
		LinglianPosTypeTwo,
		LinglianPosTypeThree,
		LinglianPosTypeFour,
		LinglianPosTypeFive,
		LinglianPosTypeSix,
		LinglianPosTypeSeven,
		LinglianPosTypeEight,
		LinglianPosTypeNine,
		LinglianPosTypeTen:
		return true
	default:
		return false
	}
}

const (
	MinLinglianPosType = LinglianPosTypeOne
	MaxLinglianPosType = LinglianPosTypeTen
)
