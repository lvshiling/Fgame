package log

type ShopLogReason int32

const (
	ShopLogReasonGM ShopLogReason = iota + 1
	ShopBuyLogReason
)

func (r ShopLogReason) Reason() int32 {
	return int32(r)
}

func (r ShopLogReason) String() string {
	return shopLogReasonGMMap[r]
}

var (
	shopLogReasonGMMap = map[ShopLogReason]string{
		ShopLogReasonGM:  "gm修改",
		ShopBuyLogReason: "商店购买shopId:%d,数量:%d,消耗钱:%d,购买方式:%s",
	}
)
