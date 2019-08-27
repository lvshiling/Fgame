/*此类自动生成,请勿修改*/
package template

/*升阶系统补偿配置*/
type ReturnXiTongTemplateVO struct {

	//id
	Id int `json:"id"`

	//阶数
	Number int32 `json:"number"`

	//系统类型
	Type int32 `json:"type"`

	//补偿的物品id
	ReturnItemId string `json:"return_item_id"`

	//补偿的物品数量
	ReturnItemCount string `json:"return_item_count"`
}
