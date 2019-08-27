/*此类自动生成,请勿修改*/
package template

/*创世之战官职模板配置*/
type ChuangShiGuanZhiTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int `json:"next_id"`

	//官职等级
	Level int32 `json:"level"`

	//官职名字
	Name string `json:"name"`

	//升级成功率
	UpLevPercent int32 `json:"update_percent"`

	//需要的威望
	UseWeiWang int32 `json:"use_weiwang"`

	//需要的银两
	UseMoney int32 `json:"use_money"`

	//需要的物品id
	ItemId int32 `json:"use_item"`

	//需要的物品数量
	ItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//获取的物品id
	GetItemId int32 `json:"get_item_id"`

	//获取的物品数量
	GetItemCount int32 `json:"get_item_count"`

	//时装id
	FashionId int32 `json:"fashion_id"`
}
