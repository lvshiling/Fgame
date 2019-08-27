/*此类自动生成,请勿修改*/
package template

/*定情信物套装增加*/
type MarryXinWuSuitTemplateVO struct {

	//id
	Id int `json:"id"`

	//1技能，2属性
	Type int32 `json:"type"`

	//Type:1技能Id，Type2属性
	Value int32 `json:"value"`

	//套装物品数量
	Num int32 `json:"num"`

	//增加生命
	Hp int32 `json:"hp"`

	//增加攻击
	Attack int32 `json:"attack"`

	//增加防御
	Defence int32 `json:"defence"`

	//套装描述
	Describe string `json:"describe"`
}
