package types

type LingZhuType int32

const (
	//灵珠1
	LingZhuTypeOne LingZhuType = iota
	//灵珠2
	LingZhuTypeTwo
	//灵珠3
	LingZhuTypeThree
	//灵珠4
	LingZhuTypeFour
	//灵珠5
	LingZhuTypeFive
)

const (
	MinLingZhuType = LingZhuTypeOne
	MaxLingZhuType = LingZhuTypeFive
)

var (
	lingZhuTypeMap = map[LingZhuType]string{
		LingZhuTypeOne:   "灵珠1",
		LingZhuTypeTwo:   "灵珠2",
		LingZhuTypeThree: "灵珠3",
		LingZhuTypeFour:  "灵珠4",
		LingZhuTypeFive:  "灵珠5",
	}
)

func (t LingZhuType) Valid() bool {
	switch t {
	case LingZhuTypeOne,
		LingZhuTypeTwo,
		LingZhuTypeThree,
		LingZhuTypeFour,
		LingZhuTypeFive:
		return true
	}
	return false
}

func (t LingZhuType) String() string {
	return lingZhuTypeMap[t]
}
