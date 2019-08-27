package types

type YingLingPuSuiPianPositionType int

func (t YingLingPuSuiPianPositionType) Valid() bool {
	switch t {
	case YingLingPuSuiPianPositionTypeOne,
		YingLingPuSuiPianPositionTypeTwo,
		YingLingPuSuiPianPositionTypeThree,
		YingLingPuSuiPianPositionTypeFour:
		return true
	}
	return false
}

const (
	YingLingPuSuiPianPositionTypeOne YingLingPuSuiPianPositionType = iota
	YingLingPuSuiPianPositionTypeTwo
	YingLingPuSuiPianPositionTypeThree
	YingLingPuSuiPianPositionTypeFour
)

// 图鉴类型
type YingLingPuTuJianType int

const (
	YingLingPuTuJianTypeCommon YingLingPuTuJianType = iota
	YingLingPuTuJianTypeGread
	YingLingPuTuJianTypeHeroic
)

func (t YingLingPuTuJianType) Valid() bool {
	switch t {
	case YingLingPuTuJianTypeCommon,
		YingLingPuTuJianTypeGread,
		YingLingPuTuJianTypeHeroic:
		return true
	}
	return false
}

func GetMaxYingLingPuType() YingLingPuTuJianType {
	return YingLingPuTuJianTypeHeroic
}

var (
	yingLingPuTuJianMap = map[YingLingPuTuJianType]string{
		YingLingPuTuJianTypeCommon: "普通",
		YingLingPuTuJianTypeGread:  "超级",
		YingLingPuTuJianTypeHeroic: "史诗",
	}
)

func (t YingLingPuTuJianType) String() string {
	return yingLingPuTuJianMap[t]
}
