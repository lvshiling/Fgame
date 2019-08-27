/*此类自动生成,请勿修改*/
package template

/*宝宝读书模板配置*/
type BabyLearnTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级id
	NextId int `json:"next_id"`

	//读书等级
	Level int32 `json:"level"`

	//经验
	Experience int32 `json:"experience"`

	//属性二进制
	AttrType int32 `json:"attr_type"`

	//属性倍数
	BeiShu int32 `json:"beishu"`

	//转世返还物品
	ZsReturnItem string `json:"zs_return_item"`

	//转世返还数量
	ZsReturnCount string `json:"zs_return_count"`
}
