/*此类自动生成,请勿修改*/
package template

/*打折配置*/
type BargainTemplateVo struct {

	//id
	Id int `json:"id"`

	//下一级id
	NextId int `json:"next_id"`

	//打折次数
	BargainTimes int32 `json:"bargain_times"`

	//打折下限
	DazheMin int32 `json:"dazhe_min"`

	//打折上限
	DazheMax int32 `json:"dazhe_max"`
}
