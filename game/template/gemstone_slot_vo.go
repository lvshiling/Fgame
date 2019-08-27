/*此类自动生成,请勿修改*/
package template

/*宝石格子配置*/
type GemstoneSlotTemplateVO struct {

	//id
	Id int `json:"id"`

	//位置
	Position int32 `json:"position"`

	//槽位顺序
	Order int32 `json:"order"`

	//宝石类型
	GemstoneType int32 `json:"gemstone_type"`

	//需求等级
	NeedLevel int32 `json:"need_;evel"`

	//需求层数
	NeedLayer int32 `json:"need_layer"`

	//消耗的物品id
	NeedItemId string `json:"need_item_id"`

	//消耗的物品数量
	NeedItemCount string `json:"need_item_count"`
}
