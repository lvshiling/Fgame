/*此类自动生成,请勿修改*/
package template

/*神魔排行榜配置*/
type ShenMoRankTemplateVO struct {

	//id
	Id int `json:"id"`

	//仙盟排名区间最小值
	RankMin int32 `json:"rank_min"`

	//仙盟排名区间最大值
	RankMax int32 `json:"rank_max"`

	//仙盟排名获得的银两
	GetSilver int32 `json:"get_silver"`

	//仙盟排名获得的绑元
	GetBindGold int32 `json:"get_bind_gold"`

	//获得的功勋
	GetGongXun int32 `json:"get_gongxun"`

	//仙盟排名获得的物品
	GetItemId string `json:"get_item_id"`

	//仙盟排名获得的物品数量
	GetItemCount string `json:"get_item_count"`
}
