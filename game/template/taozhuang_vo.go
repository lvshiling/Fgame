/*此类自动生成,请勿修改*/
package template

/*套装配置*/
type TaozhuangTemplateVO struct {

	//id
	Id int `json:"id"`

	//名字
	Name string `json:"name"`

	//类型
	Type int32 `json:"type"`

	//属性
	AttrId int32 `json:"attr_id"`

	//数量
	Number int32 `json:"number"`

	//血量
	Hp int64 `json:"hp"`

	//攻击
	Attack int64 `json:"attack"`

	//防御
	Defence int64 `json:"defence"`

	//血量
	HpPercent int64 `json:"hp_percent"`

	//攻击
	AttPercent int64 `json:"att_percent"`

	//防御
	DefPercent int64 `json:"def_percent"`
}
