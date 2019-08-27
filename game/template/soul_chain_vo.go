/*此类自动生成,请勿修改*/
package template

/*帝魂锁链配置*/
type SoulChainTemplateVO struct {

	//id
	Id int `json:"id"`

	//所需觉醒的帝魂数量
	NeedCount int32 `json:"need_count"`

	//效果类型
	Type int32 `json:"type"`

	//对应的效果加成
	Value int32 `json:"value"`

	//该条属性的描述
	Describe string `json:"describe"`
}
