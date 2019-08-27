/*此类自动生成,请勿修改*/
package template

/*竞技场配置*/
type ArenaTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型
	Type int32 `json:"type"`

	//连续次数
	LianXuCount int32 `json:"lianxu_count"`

	//连续获得积分
	LianXuGetJifen int32 `json:"lianxu_get_jifen"`

	//竞技场奖励物品
	LianxuGetItemId string `json:"lianxu_get_item_id"`

	//竞技场奖励数量
	LianxuGetItemCount string `json:"lianxu_get_item_count"`

	//竞技场等级
	ArenaNum int32 `json:"arena_num"`

	//竞技场银两奖励
	ArenaSilver int32 `json:"arena_silver"`

	//竞技场奖励物品
	ArenaItemId string `json:"arena_item_id"`

	//竞技场奖励数量
	ArenaItemAmount string `json:"arena_item_amount"`
}
