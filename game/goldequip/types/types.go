package types

//元神金装存放的位置
type GoldEquipPositionType int32

const (
	GoldEquipPositionTypeBag   GoldEquipPositionType = iota //背包
	GoldEquipPositionTypeSlote                              //元神金装装备槽
)

func (t GoldEquipPositionType) Valid() bool {
	switch t {
	case GoldEquipPositionTypeBag, GoldEquipPositionTypeSlote:
		return true
	default:
		return false
	}
}

type UpstarResultType int32

const (
	UpstarResultTypeSuccess UpstarResultType = iota //强化成功
	UpstarResultTypeFailed                          //强化失败
	UpstarResultTypeBack                            //强化回退
)

// 等级套装类型
type TaoZhuangMuBiaoType int32

const (
	TaoZhuangMuBiaoTypeGem        TaoZhuangMuBiaoType = iota //宝石
	TaoZhuangMuBiaoTypeUpstar                                //强化
	TaoZhuangMuBiaoTypeOpen                                  //开光
	TaoZhuangMuBiaoTypeLevel                                 //升星
	TaoZhuangMuBiaoTypeGodCasting                            //神铸
	TaoZhuangMuBiaoTypeTemp                                  //临时
)

func (t TaoZhuangMuBiaoType) Valid() bool {
	switch t {
	case TaoZhuangMuBiaoTypeGem,
		TaoZhuangMuBiaoTypeOpen,
		TaoZhuangMuBiaoTypeLevel,
		TaoZhuangMuBiaoTypeUpstar,
		TaoZhuangMuBiaoTypeGodCasting,
		TaoZhuangMuBiaoTypeTemp:
		return true
	default:
		return false
	}
}

type SpiritType int32

const (
	SpiritTypeOne   SpiritType = iota //铸灵A
	SpiritTypeTwo                     //铸灵B
	SpiritTypeThree                   //铸灵C
	SpiritTypeFour                    //铸灵D
	SpiritTypeFive                    //铸灵E
)

const (
	MinSpiritType = SpiritTypeOne
	MaxSpiritType = SpiritTypeFive
)

func (t SpiritType) Valid() bool {
	switch t {
	case SpiritTypeOne,
		SpiritTypeTwo,
		SpiritTypeThree,
		SpiritTypeFour,
		SpiritTypeFive:
		return true
	default:
		return false
	}
}

var (
	castingSpiritStringMap = map[SpiritType]string{
		SpiritTypeOne:   "铸灵A",
		SpiritTypeTwo:   "铸灵B",
		SpiritTypeThree: "铸灵C",
		SpiritTypeFour:  "铸灵D",
		SpiritTypeFive:  "铸灵E",
	}
)

func (t SpiritType) String() string {
	return castingSpiritStringMap[t]
}

type ForgeSoulType int32

const (
	ForgeSoulTypeDragon   ForgeSoulType = iota //青龙
	ForgeSoulTypeTiger                         //白虎
	ForgeSoulTypeSparrow                       //朱雀
	ForgeSoulTypeTortoise                      //玄武
)

const (
	MinForgeSoulType = ForgeSoulTypeDragon
	MaxForgeSoulType = ForgeSoulTypeTortoise
)

func (t ForgeSoulType) Valid() bool {
	switch t {
	case ForgeSoulTypeDragon,
		ForgeSoulTypeTiger,
		ForgeSoulTypeSparrow,
		ForgeSoulTypeTortoise:
		return true
	default:
		return false
	}
}

var (
	forgeSoulStringMap = map[ForgeSoulType]string{
		ForgeSoulTypeDragon:   "青龙",
		ForgeSoulTypeTiger:    "白虎",
		ForgeSoulTypeSparrow:  "朱雀",
		ForgeSoulTypeTortoise: "玄武",
	}
)

func (t ForgeSoulType) String() string {
	return forgeSoulStringMap[t]
}
