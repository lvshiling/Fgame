/*此类自动生成,请勿修改*/
package template

/*转生礼包折扣配置*/
type CircleBargainTemplateVO struct {

	//id
	Id int `json:"id"`

	//折扣组
	Type int32 `json:"type"`

	//数量最小值
	ItemCountMin int32 `json:"item_count_min"`

	//数量最大值
	ItemCountMax int32 `json:"item_count_max"`

	//折扣
	Discount int32 `json:"discount"`
}
