/*此类自动生成,请勿修改*/
package template

/*元神金装强化配置*/
type GoldEquipUpstarTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//成功率
	SuccessRate int32 `json:"success_rate"`

	//强化等级
	Level int32 `json:"level"`

	//使用物品id
	NeedItem string `json:"need_item"`

	//使用物品数量
	NeedCount string `json:"need_count"`

	//失败率
	FailReturnRate int32 `json:"fail_return_rate"`

	//失败回退等级
	FailReturnLevel int32 `json:"fail_return_level"`

	//返还物品Id
	MeltingReturnId string `json:"melting_return_id"`

	//返还物品数量
	MeltingReturnCount string `json:"melting_return_count"`

	//防爆物品id
	ProtectItemId int32 `json:"protect_item_id"`

	//防爆物品数量
	ProtectItemCount int32 `json:"protect_item_count"`

	//生命
	AddHp int32 `json:"add_hp"`

	//攻击
	AddAttack int32 `json:"add_attack"`

	//防御
	AddDef int32 `json:"add_def"`
}
