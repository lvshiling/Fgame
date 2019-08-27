/*此类自动生成,请勿修改*/
package template

/*命格合成配置*/
type MingGeSynthesisTemplateVO struct {

	//id
	Id int `json:"id"`

	//合成所需的物品id
	NeedItemId string `json:"need_item_id"`

	//合成所需的物品数量
	NeedItemCount string `json:"need_item_count"`

	//合成后的物品ID
	ItemId int32 `json:"item_id"`

	//合成后的物品数量
	ItemCount int32 `json:"item_count"`

	//合成所需银两
	NeedSilver int32 `json:"need_silver"`

	//合成失败后获得的物品关联到掉落表
	ShiBaiDrop int32 `json:"shibai_drop"`

	//合成所需元宝
	NeedGold int32 `json:"need_gold"`

	//合成所需绑定元宝
	NeedBindGold int32 `json:"need_bind_gold"`

	//合成成功率万分比
	SuccessRate int32 `json:"success_rate"`
}
