/*此类自动生成,请勿修改*/
package template

/*合成配置*/
type SynthesisTemplateVO struct {

	//id
	Id int `json:"id"`

	//左侧标签名称
	Type1 string `json:"type1"`

	//子标签ID
	FirstTabId int32 `json:"first_tab_id"`

	//合成配方名称
	Name1 string `json:"name1"`

	//性别需求
	NeedGender int32 `json:"need_gender"`

	//职业需求
	NeedProfession int32 `json:"need_profession"`

	//合成所需物品ID1
	NeedItemId1 int32 `json:"need_item_id_1"`

	//合成所需物品的数量1
	NeedItemCount1 int32 `json:"need_item_count_1"`

	//合成所需物品ID2
	NeedItemId2 int32 `json:"need_item_id_2"`

	//合成所需物品的数量2
	NeedItemCount2 int32 `json:"need_item_count_2"`

	//合成所需物品ID3
	NeedItemId3 int32 `json:"need_item_id_3"`

	//合成所需物品的数量3
	NeedItemCount3 int32 `json:"need_item_count_3"`

	//合成所需物品ID4
	NeedItemId4 int32 `json:"need_item_id_4"`

	//合成所需物品的数量4
	NeedItemCount4 int32 `json:"need_item_count_4"`

	//合成所需物品ID5
	NeedItemId5 int32 `json:"need_item_id_5"`

	//合成所需物品的数量5
	NeedItemCount5 int32 `json:"need_item_count_5"`

	//合成后的物品ID
	ItemId int32 `json:"item_id"`

	//合成后的物品数量
	ItemCount int32 `json:"item_count"`

	//合成所需银两
	NeedSilver int32 `json:"need_silver"`

	//合成所需元宝
	NeedGold int32 `json:"need_gold"`

	//合成所需绑定元宝
	NeedBindGold int32 `json:"need_bind_gold"`

	//合成成功率万分比
	SuccessRate int32 `json:"success_rate"`

	//单次最大合成数量
	MaxCount int32 `json:"max_count"`

	//功能开启模块ID
	HintContent int32 `json:"hint_content"`

	//防爆符物品id
	ExplosionId int32 `json:"explosion_id"`

	//防爆符物品数量
	ExplosionCount int32 `json:"explosion_count"`
}
