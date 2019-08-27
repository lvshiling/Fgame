/*此类自动生成,请勿修改*/
package template

/*龙宫探宝排行榜配置*/
type LongGongRankTemplateVO struct {

	//id
	Id int `json:"id"`

	//next_id
	NextId int `json:"next_id"`

	//玩家伤害排名区间最小值
	RankMin int32 `json:"rank_min"`

	//玩家伤害排名区间最大值
	RankMax int32 `json:"rank_max"`

	//玩家伤害排名获得的经验值
	RewExp int32 `json:"rew_exp"`

	//玩家伤害排名获得的银两
	RewSilver int32 `json:"rew_silver"`

	//玩家伤害排名获得的绑元
	RewBindGold int32 `json:"rew_bind_gold"`

	//玩家伤害排名获得的元宝
	RewGold int32 `json:"rew_gold"`

	//玩家伤害排名获得的物品
	RewItemId string `json:"rew_item_id"`

	//玩家伤害排名获得的物品数量
	RewItemCount string `json:"rew_item_count"`
}
