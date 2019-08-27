/*此类自动生成,请勿修改*/
package template

/*神域之战排行榜配置*/
type ShenYuRankTemplateVO struct {

	//id
	Id int `json:"id"`

	//next_id
	NextId int `json:"next_id"`

	//神域类型
	RoundType int32 `json:"type"`

	//排名区间最小值
	RankMin int32 `json:"rank_min"`

	//排名区间最大值
	RankMax int32 `json:"rank_max"`

	//排名获得的经验值
	RewExp int32 `json:"rew_exp"`

	//排名获得的经验值点
	RewExpPoint int32 `json:"rew_exp_point"`

	//排名获得的银两
	RewSilver int32 `json:"rew_silver"`

	//排名获得的绑元
	RewBindGold int32 `json:"rew_bind_gold"`

	//排名获得的元宝
	RewGold int32 `json:"rew_gold"`

	//排名获得的物品
	RewItemId string `json:"rew_item_id"`

	//排名获得的物品数量
	RewItemCount string `json:"rew_item_count"`
}
