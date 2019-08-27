package types

// 仙尊特权卡类型
type XianZunCardType int32

const (
	XianZunCardTypeSliver XianZunCardType = iota // 白银仙尊
	JieYiDaoJuTypeGold                           // 黄金仙尊
	JieYiDaoJuTypeDiamond                        // 钻石仙尊
)

func (t XianZunCardType) Valid() bool {
	switch t {
	case XianZunCardTypeSliver,
		JieYiDaoJuTypeGold,
		JieYiDaoJuTypeDiamond:
		return true
	default:
		return false

	}
}

var (
	xianZunCardMap = map[XianZunCardType]string{
		XianZunCardTypeSliver: "白银仙尊",
		JieYiDaoJuTypeGold:    "黄金仙尊",
		JieYiDaoJuTypeDiamond: "钻石仙尊",
	}
)

func (t XianZunCardType) String() string {
	return xianZunCardMap[t]
}
