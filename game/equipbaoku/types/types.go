package types

type BaoKuType int32

const (
	BaoKuTypeEquip     BaoKuType = iota // 装备宝库
	BaoKuTypeMaterials                  // 材料宝库
)

func (t BaoKuType) Valid() bool {
	switch t {
	case BaoKuTypeEquip,
		BaoKuTypeMaterials:
		return true
	}
	return false
}

var BaoKuNameMap = map[BaoKuType]string{
	BaoKuTypeEquip:     "装备宝库",
	BaoKuTypeMaterials: "材料宝库",
}

func (t BaoKuType) GetBaoKuName() string {
	return BaoKuNameMap[t]
}
