/*此类自动生成,请勿修改*/
package template

/*神域配置*/
type ShenYuTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一层id
	NextId int `json:"next_id"`

	//神域类型
	RoundType int32 `json:"type"`

	//活动持续时间
	RoundTime int64 `json:"jiesuan_time"`

	//地图id
	MapId int32 `json:"map_id"`

	//每轮入围人数
	WinRank int32 `json:"win_rank"`

	//幸运奖时间
	LuckyRewTime int32 `json:"lucky_time"`

	//幸运奖人数
	LuckyPalyerCount int32 `json:"lucky_palyer_count"`

	//幸运奖励物品
	LuckyItemId string `json:"lucky_item_id"`

	//幸运奖励物品数量
	LuckyItemCount string `json:"lucky_item_count"`

	//钥匙重置标志
	ResetKeyFlag int32 `json:"reset_key_flag"`
}
