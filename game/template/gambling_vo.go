/*此类自动生成,请勿修改*/
package template

/*赌石配置*/
type GamblingTemplateVO struct {

	//id
	Id int `json:"id"`

	//赌石类型
	Type int32 `json:"type"`

	//奖励预览
	RewardPreview string `json:"reward_preview"`

	//收益(间隔时间)
	Revenue int32 `json:"revenue"`

	//单次所需银两
	NeedYinLiang int32 `json:"need_yinliang"`

	//单次所需元宝
	NeedGold int32 `json:"need_gold"`

	//单次所需原石数量
	NeedYuanShi int32 `json:"need_yuanshi"`

	//所需物品ID
	NeedItem int32 `json:"need_item"`

	//所需物品数量
	NeedItemNum int32 `json:"need_item_num"`

	//间隔几次走1的掉落包
	IntervalNum1 int32 `json:"interval_num_1"`

	//掉落包1
	DropId1 int32 `json:"drop_id_1"`

	//间隔几次走2的掉落包
	IntervalNum2 int32 `json:"interval_num_2"`

	//掉落包2
	DropId2 int32 `json:"drop_id_2"`

	//间隔几次走3的掉落包
	IntervalNum3 int32 `json:"interval_num_3"`

	//掉落包3
	DropId3 int32 `json:"drop_id_3"`

	//首次掉落包
	FirstDrop int32 `json:"first_drop"`
}
