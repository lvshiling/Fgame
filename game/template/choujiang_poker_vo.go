/*此类自动生成,请勿修改*/
package template

/*摸金奖励配置*/
type ChouJiangPokerTemplateVO struct {

	//id
	Id int `json:"id"`

	//收集类型
	Type int32 `json:"type"`

	//活动id
	GroupId int32 `json:"group_id"`

	//奖励物品id
	RawItem string `json:"raw_item"`

	//奖励物品数量
	RawCount string `json:"raw_count"`

	//奖励银两
	RawYinliang int32 `json:"raw_yinliang"`

	//奖励绑元
	RawBindGold int32 `json:"raw_bind_gold"`

	//奖励元宝
	RawGold int32 `json:"raw_gold"`
}
