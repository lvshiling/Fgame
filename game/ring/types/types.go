package types

// 特戒类型
type RingType int32

const (
	RingTypeTypeMaBi     RingType = iota // 麻痹戒指
	RingTypeTypeShenNong                 // 神农戒指
	RingTypeTypeShiHun                   // 噬魂戒指
	RingTypeTypeHaoTian                  // 昊天戒指
)

func (t RingType) Valid() bool {
	switch t {
	case RingTypeTypeMaBi,
		RingTypeTypeShenNong,
		RingTypeTypeShiHun,
		RingTypeTypeHaoTian:
		return true
	default:
		return false
	}
}

var (
	ringMap = map[RingType]string{
		RingTypeTypeMaBi:     "麻痹戒指",
		RingTypeTypeShenNong: "神农戒指",
		RingTypeTypeShiHun:   "噬魂戒指",
		RingTypeTypeHaoTian:  "昊天戒指",
	}
)

func (t RingType) String() string {
	return ringMap[t]
}

// 积分宝库类型
type BaoKuType int32

const (
	BaoKuTypeRing BaoKuType = iota // 特戒宝库
)

func (t BaoKuType) Valid() bool {
	switch t {
	case BaoKuTypeRing:
		return true
	default:
		return false
	}
}

var (
	baoKuMap = map[BaoKuType]string{
		BaoKuTypeRing: "特戒宝库",
	}
)

func (t BaoKuType) String() string {
	return baoKuMap[t]
}

// 积分宝库寻宝类型
type BaoKuAttendType int32

const (
	BaoKuAttendTypeOne BaoKuAttendType = iota // 单次
	BaoKuAttendTypeTen                        // 十连
)

func (t BaoKuAttendType) Valid() bool {
	switch t {
	case BaoKuAttendTypeOne,
		BaoKuAttendTypeTen:
		return true
	default:
		return false
	}
}

var (
	BaoKuAttendTypeMap = map[BaoKuAttendType]int32{
		BaoKuAttendTypeOne: 1,
		BaoKuAttendTypeTen: 10,
	}
)

func (t BaoKuAttendType) Int() int32 {
	return BaoKuAttendTypeMap[t]
}
