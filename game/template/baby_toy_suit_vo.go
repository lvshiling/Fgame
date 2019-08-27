/*此类自动生成,请勿修改*/
package template

/*宝宝玩具套装属性配置*/
type BabyToySuitTemplateVO struct {

	//id
	Id int `json:"id"`

	//属性类型
	Type int32 `json:"type"`

	//当type=1时,此字段为攻击时附加buff的id
	Value1 int32 `json:"value1"`

	//激活属性所需数量
	Num int32 `json:"num"`
}
