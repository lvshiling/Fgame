/*此类自动生成,请勿修改*/
package template

/*折扣配置*/
type DiscountTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动id
	Group int32 `json:"group"`

	//更新天数
	DayGroup int32 `json:"day_group"`

	//位置
	Index int32 `json:"index"`

	//物品模板id
	ItemId int32 `json:"item_id"`

	//名称
	Name string `json:"name"`

	//一次性购买的数量
	BuyCount int32 `json:"buy_count"`

	//最大购买数量
	MaxCount int32 `json:"max_count"`

	//原价
	YuanGold int32 `json:"yuan_gold"`

	//实际价格
	UseGold int32 `json:"use_gold"`

	//折扣
	Dazhe int32 `json:"dazhe"`

	//是否绑定
	IsBind int32 `json:"is_bind"`

	//全服限购
	LimitQuanfu int32 `json:"limit_quanfu"`

	//每日限购次数,第二天清空处理
	LimitCount int32 `json:"limit_count"`
}
