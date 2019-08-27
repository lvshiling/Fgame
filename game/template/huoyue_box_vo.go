/*此类自动生成,请勿修改*/
package template

/*活跃度宝箱配置*/
type HuoYueBoxTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一星级id
	NextId int32 `json:"next_id"`

	//领取奖励所需要的星级
	NeedStar int32 `json:"need_star"`

	//获得的银两
	AwardSilver int32 `json:"award_silver"`

	//获取的元宝
	AwardGold int32 `json:"award_gold"`

	//获取的绑元
	AwardBindGold int32 `json:"award_bindgold"`

	//获得的物品
	AwardItemId string `json:"award_item_id"`

	//获得的物品数量
	AwardItemIdCount string `json:"award_item_id_count"`
}
