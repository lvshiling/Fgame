package types

type ShopDiscountType int32

const (
	ShopDiscountTypeDefault ShopDiscountType = iota
	//618大促特权
	ShopDiscountType618
)

const (
	MinType = ShopDiscountType618
	MaxType = ShopDiscountType618
)

func (wt ShopDiscountType) Valid() bool {
	switch wt {
	case ShopDiscountType618:
		return true
	}
	return false
}

var (
	shopDiscountTypeMap = map[ShopDiscountType]string{
		ShopDiscountType618: "618大促特权",
	}
)

func (spt ShopDiscountType) String() string {
	return shopDiscountTypeMap[spt]
}
