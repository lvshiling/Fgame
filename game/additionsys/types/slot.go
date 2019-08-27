package types

type SlotPositionType int32

const (
	//位置1
	SlotPositionTypeOne SlotPositionType = iota
	//位置2
	SlotPositionTypeTwo
	//位置3
	SlotPositionTypeThree
	//位置4
	SlotPositionTypeFour
)

const (
	MinPosition = SlotPositionTypeOne
	MaxPosition = SlotPositionTypeFour
)

var (
	slotPositionTypeMap = map[SlotPositionType]string{
		SlotPositionTypeOne:   "位置1",
		SlotPositionTypeTwo:   "位置2",
		SlotPositionTypeThree: "位置3",
		SlotPositionTypeFour:  "位置4",
	}
)

func (spt SlotPositionType) Valid() bool {
	switch spt {
	case SlotPositionTypeOne,
		SlotPositionTypeTwo,
		SlotPositionTypeThree,
		SlotPositionTypeFour:
		return true
	}
	return false
}

func (spt SlotPositionType) String() string {
	return slotPositionTypeMap[spt]
}
