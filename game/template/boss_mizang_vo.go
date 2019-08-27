/*此类自动生成,请勿修改*/
package template

/*boss密藏配置*/
type BossMizangTemplateVO struct {

	//id
	Id int `json:"id"`

	//物品id
	SilverItemId int32 `json:"silver_item_id"`

	//物品数量
	SilverItemCount int32 `json:"silver_item_count"`

	//掉落
	SilverDrop string `json:"silver__drop"`

	//比例
	SilverRateAdd int32 `json:"silver_rate_add"`

	//物品id
	GoldItemId int32 `json:"gold_item_id"`

	//物品数量
	GoldItemCount int32 `json:"gold_item_count"`

	//掉落
	GoldDrop string `json:"gold_drop"`

	//比例
	GoldRateAdd int32 `json:"gold_rate_add"`

	//采集物
	CaijiBiologyId int32 `json:"caiji_biology_id"`
}
