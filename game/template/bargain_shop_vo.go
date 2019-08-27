/*此类自动生成,请勿修改*/
package template

/*打折礼包配置*/
type BargainShopTemplateVO struct {

	//id
	Id int `json:"id"`

	//活动id
	Group int32 `json:"group"`

	//类型
	Type int32 `json:"type"`

	//物品id
	ItemId string `json:"item_id"`

	//数量
	ItemCount string `json:"item_count"`

	//职业
	Profession int32 `json:"profession"`

	//性别
	Gender int32 `json:"gender"`

	//原价
	YuanGold int32 `json:"yuan_gold"`

	//折扣起始id
	BargainBegain int32 `json:"bargain_begain"`
}
