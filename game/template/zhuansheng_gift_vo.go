/*此类自动生成,请勿修改*/
package template

/*转生大礼包配置*/
type ZhuanShengGiftTemplateVO struct {

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

	//购买数量
	BuyCount int32 `json:"buy_count"`

	//最大数量
	MaxCount int32 `json:"max_count"`

	//原价
	YuanGold int32 `json:"yuan_gold"`

	//现价
	UseGold int32 `json:"use_gold"`

	//折数
	DaZhe int32 `json:"dazhe"`

	//是否绑定
	IsBind int32 `json:"is_bind"`

	//最大购买
	BuyMax int32 `json:"buy_max"`

	//消耗id
	UseItemId string `json:"use_item_id"`

	//消耗数量
	UseItemCount string `json:"use_item_count"`

	//充值后可购买
	NeedChongZhi int32 `json:"need_chongzhi"`

	//批量购买折扣
	Bargain int32 `json:"bargain"`

	//赠送id
	GiveItemId string `json:"give_item_id"`

	//赠送数量
	GiveItemCount string `json:"give_item_count"`

	//现积分价
	UsePoint int32 `json:"use_point"`
}
