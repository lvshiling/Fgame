/*此类自动生成,请勿修改*/
package template

/*四神遗迹宝箱配置*/
type FourGodBoxTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//消耗物品id
	UseItemId int32 `json:"use_item_id"`

	//消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//宝箱类型
	Type int32 `json:"type"`

	//采集物id
	BiologyId int32 `json:"biology_id"`

	//消耗物品的最小数量
	KeyMin int32 `json:"key_min"`

	//消耗物品的最大数量
	KeyMax int32 `json:"key_max"`

	//奖励的银两
	AwardSilver int32 `json:"award_silver"`

	//奖励的元宝
	AwardGold int32 `json:"award_gold"`

	//奖励的绑元
	AwardBindGold int32 `json:"award_bindgold"`

	//奖励的经验
	AwardExp int32 `json:"award_exp"`

	//奖励的经验点
	AwardExpPoint int32 `json:"award_exp_point"`

	//奖励的经验点
	award_item_id int32 `json:"award_item_id"`

	//掉落id
	DropId string `json:"drop_id"`
}
