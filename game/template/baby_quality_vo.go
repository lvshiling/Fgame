/*此类自动生成,请勿修改*/
package template

/*宝宝品质模板配置*/
type BabyQualityTemplateVO struct {

	//id
	Id int `json:"id"`

	//补品区间
	QuJian string `json:"qujian"`

	//下一级id
	NextId int `json:"next_id"`

	//品质1
	Type1 int32 `json:"type_1"`

	//品质2
	Type2 int32 `json:"type_2"`

	//品质3
	Type3 int32 `json:"type_3"`

	//品质4
	Type4 int32 `json:"type_4"`

	//权重
	Rate1 int32 `json:"rate_1"`

	//权重2
	Rate2 int32 `json:"rate_2"`

	//权重3
	Rate3 int32 `json:"rate_3"`

	//权重4
	Rate4 int32 `json:"rate_4"`
}
