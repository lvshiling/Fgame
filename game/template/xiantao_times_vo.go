/*此类自动生成,请勿修改*/
package template

/*仙桃大会劫取配置*/
type XianTaoTimesTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int `json:"next_id"`

	//第N次
	Times int32 `json:"times"`

	//劫取成功后获得的仙桃比例（万分比）
	JieQuPercent int32 `json:"jiequ_percent"`

	//被劫取后损失的仙桃比例（万分比）
	SunShiPercent int32 `json:"sunshi_percent"`
}
