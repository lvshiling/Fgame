/*此类自动生成,请勿修改*/
package template

/*屠龙装备强化配置*/
type TuLongEquipStrengthenTemplateVO struct {

	//id
	Id int `json:"id"`

	//套装类型
	Type int32 `json:"type"`

	//部位
	SubType int32 `json:"sub_type"`

	//强化等级
	Level int32 `json:"level"`

	//下级id
	NextId int32 `json:"next_id"`

	//强化概率
	Rate int32 `json:"rate"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//强化的物品id
	NeedItem int32 `json:"need_item"`

	//强化的物品数量
	ItemCount int32 `json:"item_count"`
}
