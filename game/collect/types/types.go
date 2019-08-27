package types

type CollectType int32

const (
	CollectTypeServer CollectType = iota
	CollectTypeClient
)

func (t CollectType) Valid() bool {
	switch t {
	case CollectTypeClient,
		CollectTypeServer:
		return true
	}
	return false
}

type CollectPointFinishType int32

const (
	CollectPointFinishTypeAway CollectPointFinishType = iota //移除
	CollectPointFinishTypeLive                               //保存
)

func (t CollectPointFinishType) Valid() bool {
	switch t {
	case CollectPointFinishTypeAway,
		CollectPointFinishTypeLive:
		return true
	}
	return false
}

type CollectChooseFinishType int32

const (
	CollectChooseFinishTypeLow CollectChooseFinishType = iota
	CollectChooseFinishTypeExpert
)

func (t CollectChooseFinishType) Valid() bool {
	switch t {
	case CollectChooseFinishTypeLow,
		CollectChooseFinishTypeExpert:
		return true
	}
	return false
}

const (
	MinChooseType = CollectChooseFinishTypeLow
	MaxChooseType = CollectChooseFinishTypeExpert
)

//密藏类型
type MiZangOpenType int32

const (
	MiZangOpenTypeSilver MiZangOpenType = iota //银铲子
	MiZangOpenTypeGold                         //金铲子
)

func (t MiZangOpenType) Valid() bool {
	switch t {
	case MiZangOpenTypeSilver,
		MiZangOpenTypeGold:
		return true
	}
	return false
}

var (
	miZangOpenTypeMap = map[MiZangOpenType]string{
		MiZangOpenTypeSilver: "银铲子",
		MiZangOpenTypeGold:   "金铲子",
	}
)

func (t MiZangOpenType) String() string {
	return miZangOpenTypeMap[t]
}
