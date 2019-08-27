/*此类自动生成,请勿修改*/
package template

/*竞技场pvp淘汰配置*/
type ArenapvpTaoTaiTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型
	Type int32 `json:"changci"`

	//银两
	WinSilver int32 `json:"win_silver"`

	//元宝
	WinGold int32 `json:"win_gold"`

	//绑元
	WinBindGold int32 `json:"win_bind_gold"`

	//物品id
	WinItemId string `json:"win_item_id"`

	//物品数量
	WinItemCount string `json:"win_item_count"`

	//积分
	WinJifen int32 `json:"win_jifen"`

	//银两
	LoseSilver int32 `json:"lose_silver"`

	//元宝
	LoseGold int32 `json:"lose_gold"`

	//绑元
	LoseBindGold int32 `json:"lose_bind_gold"`

	//物品id
	LoseItemId string `json:"lose_item_id"`

	//物品数量
	LoseItemCount string `json:"lose_item_count"`

	//积分
	LoseJifen int32 `json:"rew_jifen"`
}
