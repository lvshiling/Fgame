/*此类自动生成,请勿修改*/
package template

/*系统装备强化配置*/
type SystemStrengthenTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//强化等级
	Level int32 `json:"level"`

	//系统类型
	Type int32 `json:"type"`

	//部位类型
	SubType int32 `json:"position"`

	//强化成功概率
	SuccessRate int32 `json:"success_rate"`

	//物品
	CostItemId string `json:"need_item"`

	//物品数量
	CostItemCount string `json:"need_count"`

	//失败回退概率
	FailBackRate int32 `json:"fail_return_rate"`

	//失败回退等级
	FailBacklevel int32 `json:"fail_return_level"`

	//生命
	Hp int32 `json:"add_hp"`

	//攻击
	Attack int32 `json:"add_attack"`

	//防御
	Defence int32 `json:"add_def"`

	//保级物品id
	ProtectItemId int32 `json:"protect_item_id"`

	//保级物品数量
	ProtectItemCount int32 `json:"protect_item_count"`
}
