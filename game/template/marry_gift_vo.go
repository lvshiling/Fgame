/*此类自动生成,请勿修改*/
package template

/*婚宴贺礼配置*/
type MarryGiftTemplateVO struct {

	//id
	Id int `json:"id"`

	//贺礼类型
	Type int32 `json:"type"`

	//消耗银两
	UseSilver int32 `json:"use_silver"`

	//消耗物品id
	UseItemId int32 `json:"use_item_id"`

	//消耗物品id
	UseItemAmount int32 `json:"use_item_amount"`

	//增加豪气值
	AddNum int32 `json:"add_num"`

	//获得buff数量
	BuffAmount int32 `json:"buff_amount"`
}
