/*此类自动生成,请勿修改*/
package template

/*商城配置*/
type ShopTemplateVO struct {

	//id
	Id int `json:"id"`

	//购买规则
	Order int32 `json:"order"`

	//名字
	Name string `json:"name"`

	//商店类型
	Type int32 `json:"type"`

	//同种商店不同类型道具标签页
	Table int32 `json:"table"`

	//物品模板id
	ItemId int32 `json:"item"`

	//一次性购买的数量
	BuyCount int32 `json:"buy_count"`

	//最大购买数量
	MaxCount int32 `json:"max_count"`

	//消耗资源类型
	ConsumeType int32 `json:"consume_type"`

	//消耗物品id
	ConsumeItemId int32 `json:"consume_item_id"`

	//消耗资源数量
	ConsumeData1 int32 `json:"consume_data1"`

	//是否绑定
	IsBind int32 `json:"is_bind"`

	//每日限购次数,第二天清空处理
	LimitCount int32 `json:"limit_count"`
}
