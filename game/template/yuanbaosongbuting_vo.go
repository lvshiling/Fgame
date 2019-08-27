/*此类自动生成,请勿修改*/
package template

/*元宝送不停模板配置*/
type YuanBaoSongBuTingTemplateVO struct {

	//id
	Id int `json:"id"`

	//id
	NextId int `json:"next_id"`

	//运营活动名字
	Name string `json:"name"`

	//购买需要元宝
	NeedGold int32 `json:"need_gold"`

	//奖励银两
	RewSilver int32 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑定元宝
	RewBindGold int32 `json:"rew_bind_gold"`

	//每日奖励物品
	RewItem string `json:"rew_item"`

	//每日奖励物品数量
	RewItemCount string `json:"rew_item_count"`
}
