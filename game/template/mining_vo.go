/*此类自动生成,请勿修改*/
package template

/*挖矿配置*/
type MiningTemplateVO struct {

	//id
	Id int `json:"id"`

	//矿工人数(矿山等级)
	Level int32 `json:"level"`

	//间隔时间(秒)
	IntervalTime int32 `json:"interval_time"`

	//收益(间隔时间)
	Revenue int32 `json:"revenue"`

	//收益存储上限
	LimitMax int32 `json:"limit_max"`

	//升级所需银两
	NeedYinLiang int32 `json:"need_yinliang"`

	//升级所需元宝
	NeedGold int32 `json:"need_gold"`

	//升级所需物品ID1
	ItemId1 int32 `json:"item_id_1"`

	//升级所需物品数量1
	ItemCount1 int32 `json:"item_count_1"`
}
