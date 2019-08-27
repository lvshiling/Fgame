package types

type DebrisType int32

const (
	//碎片位1
	DebrisTypeOne DebrisType = iota
	//碎片位2
	DebrisTypeTwo
	//碎片位3
	DebrisTypeThree
	//碎片位4
	DebrisTypeFour
	//碎片位5
	DebrisTypeFive
	//碎片位6
	DebrisTypeSix
	//碎片位7
	DebrisTypeSeven
	//碎片位8
	DebrisTypeEight
)

const (
	MinDebrisType = DebrisTypeOne
	MaxDebrisType = DebrisTypeEight
)

var (
	debrisTypeMap = map[DebrisType]string{
		DebrisTypeOne:   "碎片位1",
		DebrisTypeTwo:   "碎片位2",
		DebrisTypeThree: "碎片位3",
		DebrisTypeFour:  "碎片位4",
		DebrisTypeFive:  "碎片位5",
		DebrisTypeSix:   "碎片位6",
		DebrisTypeSeven: "碎片位7",
		DebrisTypeEight: "碎片位8",
	}
)

func (spt DebrisType) Valid() bool {
	switch spt {
	case DebrisTypeOne,
		DebrisTypeTwo,
		DebrisTypeThree,
		DebrisTypeFour,
		DebrisTypeFive,
		DebrisTypeSix,
		DebrisTypeSeven,
		DebrisTypeEight:
		return true
	}
	return false
}

func (spt DebrisType) String() string {
	return debrisTypeMap[spt]
}
