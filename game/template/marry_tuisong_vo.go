/*此类自动生成,请勿修改*/
package template

/*婚车游车时赠礼配置*/
type MarryTuiSongTemplateVO struct {

	//id
	Id int `json:"id"`

	//赠送类型，1鲜花，2赠礼
	Type int32 `json:"type"`

	//消耗元宝
	NeedGold int32 `json:"need_gold"`

	//获得经验
	RewardExp int32 `json:"reward_exp"`

	//获得经验
	RewardExpPoint int32 `json:"reward_exp_point"`

	//回馈物品Id
	ZhuHeItem string `json:"zhuhe_item"`

	//回馈物品数量
	ZhuheItemCount string `json:"zhuhe_item_count"`
}
