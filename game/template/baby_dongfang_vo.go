/*此类自动生成,请勿修改*/
package template

/*宝宝洞房模板配置*/
type BabyDongFangTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级
	NextId int `json:"next_id"`

	//宝宝数量
	BabyCount int32 `json:"baby_count"`

	//怀孕概率
	PregnantRate int32 `json:"pregnant_rate"`

	//洞房消耗物品
	PregnantItem int32 `json:"pregnant_item"`

	//洞房消耗物品数量
	PregnantCount int32 `json:"pregnant_count"`

	//失败返回物品
	FailReturnItem string `json:"fail_return_item"`

	//失败返回物品数量
	FailReturnItemCount string `json:"fail_return_item_count"`
}
