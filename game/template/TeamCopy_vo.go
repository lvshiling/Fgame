/*此类自动生成,请勿修改*/
package template

/*组队副本配置*/
type TeamCopyTemplateVO struct {

	//id
	Id int `json:"id"`

	//组队类型
	Type int32 `json:"type"`

	//名字
	Name string `json:"name"`

	//地图类型
	MapId int32 `json:"map_id"`

	//生物id
	BiologyId string `json:"biology_id"`

	//每日奖励次数
	RewardNumber int32 `json:"reward_number"`

	//复活次数
	ResurrectionNumber int32 `json:"resurrection_number"`

	//奖励物品
	ItemId string `json:"item_id"`

	//对应数量
	ItemCount string `json:"item_count"`

	//奖励元宝
	RewardGold int32 `json:"reward_gold"`

	//奖励绑元
	RewardBindgold int32 `json:"reward_bindgold"`

	//奖励银两
	RewardSilver int32 `json:"reward_silver"`

	//奖励银两
	RewardPreviewId string `json:"reward_preview_id"`

	//推荐战力
	RecForce int32 `json:"rec_force"`
}
