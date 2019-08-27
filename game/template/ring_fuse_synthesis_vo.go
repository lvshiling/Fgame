/*此类自动生成,请勿修改*/
package template

/*特戒融合合成配置*/
type RingFuseSynthesisTemplateVO struct {

	//id
	Id int `json:"id"`

	//合成所需物品1id
	NeedItemId1 int32 `json:"need_item_id_1"`

	//合成所需物品1数量
	NeedItemCount1 int32 `json:"need_item_count_1"`

	//合成所需物品2id
	NeedItemId2 int32 `json:"need_item_id_2"`

	//合成所需物品2数量
	NeedItemCount2 int32 `json:"need_item_count_2"`

	//合成后物品id
	ItemId int32 `json:"item_id"`

	//合成后物品数量
	ItemCount int32 `json:"item_count"`

	//合成成功率万分比
	SuccessRate int32 `json:"success_rate"`

	//合成所需银两
	NeedSilver int32 `json:"need_silver"`

	//合成所需元宝
	NeedGold int32 `json:"need_gold"`

	//合成所需绑元
	NeedBindGold int32 `json:"need_bind_gold"`
}
