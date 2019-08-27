package types

// 花色
type PokerType int32

const (
	PokerNumberJoker PokerType = iota //大小王
	PokerNumberSpade                  //黑桃
	PokerNumberHeart                  //红心
	PokerNumberClub                   //梅花
	PokerNumberBlock                  //方块
)

func (t PokerType) Valid() bool {
	switch t {
	case PokerNumberHeart,
		PokerNumberBlock,
		PokerNumberSpade,
		PokerNumberClub,
		PokerNumberJoker:
		return true
	default:
		return false
	}
}

// 数字
type PokerNumberType int32

const (
	PokerNumberOne PokerNumberType = iota + 1
	PokerNumberTwo
	PokerNumberThree
	PokerNumberFour
	PokerNumberFive
	PokerNumberSix
	PokerNumberSeven
	PokerNumberEight
	PokerNumberNine
	PokerNumberTen
	PokerNumberEleven
	PokerNumberTwelve
	PokerNumberThirteen
	PokerNumberSmallJoker
	PokerNumberBigJoker
)

func GetInitPokerList() (initList []int32) {
	for pokerType, numberList := range collectMap {
		for _, pokerNumber := range numberList {
			newNum := MaskEncode(pokerType, pokerNumber)
			initList = append(initList, newNum)
		}
	}

	return
}

func MaskEncode(pokerType PokerType, pokerNumber PokerNumberType) int32 {
	return int32(pokerType) + int32(pokerNumber)<<3
}

func MaskDecode(number int32) (pokerType PokerType, pokerNumber int32) {
	pokerNumber = number >> 3
	pokerTypeInt := number - pokerNumber<<3
	pokerType = PokerType(pokerTypeInt)
	return
}

func IsFinishCollect(pokerType PokerType, collectNum int32) bool {
	numberList, ok := collectMap[pokerType]
	if !ok {
		return false
	}

	if len(numberList) != int(collectNum) {
		return false
	}

	return true
}

var (
	collectMap = map[PokerType][]PokerNumberType{
		PokerNumberHeart: []PokerNumberType{
			PokerNumberOne,
			PokerNumberTwo,
			PokerNumberThree,
			PokerNumberFour,
			PokerNumberFive,
			PokerNumberSix,
			PokerNumberSeven,
			PokerNumberEight,
			PokerNumberNine,
			PokerNumberTen,
			PokerNumberEleven,
			PokerNumberTwelve,
			PokerNumberThirteen,
		},
		PokerNumberBlock: []PokerNumberType{
			PokerNumberOne,
			PokerNumberTwo,
			PokerNumberThree,
			PokerNumberFour,
			PokerNumberFive,
			PokerNumberSix,
			PokerNumberSeven,
			PokerNumberEight,
			PokerNumberNine,
			PokerNumberTen,
			PokerNumberEleven,
			PokerNumberTwelve,
			PokerNumberThirteen,
		},
		PokerNumberSpade: []PokerNumberType{
			PokerNumberOne,
			PokerNumberTwo,
			PokerNumberThree,
			PokerNumberFour,
			PokerNumberFive,
			PokerNumberSix,
			PokerNumberSeven,
			PokerNumberEight,
			PokerNumberNine,
			PokerNumberTen,
			PokerNumberEleven,
			PokerNumberTwelve,
			PokerNumberThirteen,
		},
		PokerNumberClub: []PokerNumberType{
			PokerNumberOne,
			PokerNumberTwo,
			PokerNumberThree,
			PokerNumberFour,
			PokerNumberFive,
			PokerNumberSix,
			PokerNumberSeven,
			PokerNumberEight,
			PokerNumberNine,
			PokerNumberTen,
			PokerNumberEleven,
			PokerNumberTwelve,
			PokerNumberThirteen,
		},
		PokerNumberJoker: []PokerNumberType{
			PokerNumberSmallJoker,
			PokerNumberBigJoker,
		},
	}
)
