/*此类自动生成,请勿修改*/
package template

/*系统装备套装配置*/
type SystemTaozhuangTemplateVO struct {

	//id
	Id int `json:"id"`

	//套装名称
	Name string `json:"name"`

	//套装类型
	Type int32 `json:"type"`

	//激活所需数量
	Number int32 `json:"number"`

	//套装描述
	Describe string `json:"describe"`

	//套装品质
	Pos1Quality int32 `json:"pos1_quality"`

	//增加对系统基础属性万分比
	AttrPercent int32 `json:"attr_percent"`
}
