/*此类自动生成,请勿修改*/
package template

/*命格合成配置*/
type MingGeMingPanTemplateVO struct {

	//id
	Id int `json:"id"`

	//命格子类型
	SubType int32 `json:"sub_type"`

	//下一级id
	NextId int32 `json:"next_id"`

	//命盘阶数
	Number int32 `json:"number"`

	//命盘星级
	Star int32 `json:"star"`

	//权重
	Rate int32 `json:"rate"`

	//消耗的物品id
	UseItemId int32 `json:"use_item_id"`

	//消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//成功率
	UpdateWfb int32 `json:"update_wfb"`

	//最小升阶次数
	TimesMin int32 `json:"times_min"`

	//最大升阶次数
	TimesMax int32 `json:"times_max"`

	//增加的最小祝福值
	AddMin int32 `json:"add_min"`

	//增加的最大祝福值
	AddMax int32 `json:"add_max"`

	//祝福最大总值
	ZhuFuMax int32 `json:"zhufu_max"`

	//增加的生命
	Hp int64 `json:"hp"`

	//增加的攻击值
	Attack int64 `json:"attack"`

	//增加的防御值
	Defence int64 `json:"defence"`

	//给命盘增加的属性百分比（万分比）
	Percent int32 `json:"percent"`
}
