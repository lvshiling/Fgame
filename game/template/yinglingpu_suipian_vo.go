/*此类自动生成,请勿修改*/
package template

/*英灵谱配置碎片*/
type YinglingpuSuiPianTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一部位id
	NextId int32 `json:"next_id"`

	//碎片部位编号,部位标识从1~4
	SuipianId int32 `json:"suipian_id"`

	//激活该部位消耗的物品id
	UseItemId int32 `json:"use_item_id"`

	//激活该部位消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//激活该部位后增加的生命
	Hp int32 `json:"hp"`

	//激活该部位后增加的攻击
	Attack int32 `json:"attack"`

	//激活该部位后增加的防御
	Defence int32 `json:"defence"`

	//该部位的图片资源
	Resource string `json:"resource"`
}
