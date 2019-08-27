/*此类自动生成,请勿修改*/
package template

/*3v3排行榜配置*/
type ArenaRankTemplateVO struct {

	//id
	Id int `json:"id"`

	//排名区间最小值
	RankMin int32 `json:"rank_min"`

	//排名区间最大值
	RankMax int32 `json:"rank_max"`

	//排名获得的银两
	GetSilver int32 `json:"rew_silver"`

	//排名获得的绑元
	GetBindGold int32 `json:"rew_bind_gold"`

	//排名获得的元宝
	GetGold int32 `json:"rew_gold"`

	//排名获得的物品
	GetItemId string `json:"rew_item_id"`

	//排名获得的物品数量
	GetItemCount string `json:"rew_item_count"`
}
