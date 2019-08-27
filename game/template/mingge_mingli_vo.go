/*此类自动生成,请勿修改*/
package template

/*命格命理配置*/
type MingGeMingLiTemplateVO struct {

	//id
	Id int `json:"id"`

	//命理类型
	Type int32 `json:"type"`

	//命理部位
	SubType int32 `json:"sub_type"`

	//该部位能够随机的类型(二进制)
	AttrPool int32 `json:"attr_pool"`

	//该部位指定的最大收益属性
	ZhiDingAttr int32 `json:"zhiding_attr"`

	//基础属性倍数
	AttrOne int32 `json:"attr_one"`

	//当3条命宫属性中不包含指定属性时洗练单条属性洗练到指定属性的概率(万分比)
	ZhiDingRate1 int32 `json:"zhiding_rate1"`

	//当3条命宫属性中包含1条指定属性时洗练单条属性洗练到指定属性的概率(万分比)
	ZhiDingRate2 int32 `json:"zhiding_rate2"`

	//当3条命宫属性中包含2条指定属性时洗练单条属性洗练到指定属性的概率(万分比)
	ZhiDingRate3 int32 `json:"zhiding_rate3"`

	//洗练消耗的物品id
	UseItemId int32 `json:"use_item_id"`

	//洗练消耗的物品基础倍数基础数量为1*该字段
	UseItemOne int32 `json:"use_item_one"`

	//洗练消耗的物品系数向上取整
	CoefficientUse1 float64 `json:"coefficient_use1"`

	//消耗物品数量的固定值
	CoefficientUse2 int32 `json:"coefficient_use2"`

	//洗练获得属性系数向上取整
	CoefficientAttr1 float64 `json:"coefficient_attr1"`

	//洗练属性的固定值
	CoefficientAttr2 float64 `json:"coefficient_attr2"`

	//公式里洗练次数的最大值
	XilianLimitCount int32 `json:"xilian_limit_count"`

	//收益万分比
	ShouYiPercent int32 `json:"shouyi_percent"`

	//指定属性的收益系数
	CoefficientZhiDing int32 `json:"coefficient_zhiding"`
}
