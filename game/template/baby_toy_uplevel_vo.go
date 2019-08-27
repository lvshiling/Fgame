/*此类自动生成,请勿修改*/
package template

/*宝宝玩具升级配置*/
type BabyToyUplevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//套装类型
	SuitType int32 `json:"type"`

	//部位
	Position int32 `json:"position"`

	//玩具等级
	Level int32 `json:"level"`

	//下级id
	NextId int32 `json:"next_id"`

	//强化概率
	Rate int32 `json:"success_rate"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//回退概率
	ReturnRate int32 `json:"return_rate"`

	//回退等级
	FailReturnStrengthenId int32 `json:"fail_return_strengthen_id"`

	//强化消耗
	silver_num int32 `json:"silver_num"`

	//强化消耗的物品id
	NeedItem int32 `json:"need_item"`

	//强化消耗的物品数量
	ItemCount int32 `json:"need_item_num"`
}
