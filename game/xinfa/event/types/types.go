package types

type XinFaEventType string

const (
	//心法激活
	EventTypeXinFaActive XinFaEventType = "xinfaActive"
	//心法升级
	EventTypeXinFaUpgrade XinFaEventType = "xinfaUpgrade"
)
