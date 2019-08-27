/*此类自动生成,请勿修改*/
package template

/*系统装备配置*/
type SystemEquipTemplateVO struct {

	//id
	Id int `json:"id"`

	//第几阶
	Series int32 `json:"series"`

	//套装id
	SuitGroup int32 `json:"suit_group"`

	//需要物品
	NeedItem string `json:"need_item"`

	//需要物品数量
	NeedItemNum string `json:"need_item_num"`

	//成功率
	SuccessRate int32 `json:"success_rate"`

	//下一阶物品
	Next int32 `json:"next"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//分解元神经验
	TushiExp int32 `json:"tushi_exp"`

	//分解返还物品id
	MeltingReturnId string `json:"melting_return_id"`

	//分解返还物品数量
	MeltingReturnCount string `json:"melting_return_count"`
}
