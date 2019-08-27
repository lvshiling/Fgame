/*此类自动生成,请勿修改*/
package template

/*房子配置*/
type HouseTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级di
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//房子类型
	Type int32 `json:"type"`

	//房子序号
	HouseIndex int32 `json:"number"`

	//升级所需物品ID
	UseItemId string `json:"use_item_id"`

	//升级所需物品数量
	UseItemCount string `json:"use_item_count"`

	//损坏几率
	BrokenPercent int32 `json:"broken_percent"`

	//维修消耗物品ID
	FixItemId string `json:"fix_item_id"`

	//维修消耗物品数量
	FixItemCount string `json:"fix_item_count"`

	//租金：根据房子类型获得对应的银两/绑元
	Rent int32 `json:"rent"`

	//房价：根据房子类型出售房子可获得的银两/绑元数量
	HousePrice int32 `json:"house_price"`

	//提前出售获得的房价万分比
	AdvanceSalePercent int32 `json:"advance_sale_percent"`

	//每次装修可获得物品ID
	UplevGetItem string `json:"uplev_get_item"`

	//每次装修可获得物品数量
	UplevGetItemCount string `json:"uplev_get_item_count"`
}
