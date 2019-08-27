/*此类自动生成,请勿修改*/
package template

/*命格配置*/
type MingGeTemplateVO struct {

	//id
	Id int `json:"id"`

	//命格类型
	Type int32 `json:"type"`

	//命格子类型
	SubType int32 `json:"sub_type"`

	//命格增加的生命
	Hp int64 `json:"hp"`

	//命格增加的攻击
	Attack int64 `json:"attack"`

	//命格增加的防御
	Defence int64 `json:"defence"`
}
