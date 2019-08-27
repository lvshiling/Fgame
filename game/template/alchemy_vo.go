/*此类自动生成,请勿修改*/
package template

/*炼丹配置*/
type AlchemyTemplateVO struct {

	//id
	Id int `json:"id"`

	//合成丹药id
	SynthetiseId int32 `json:"synthetise_id"`

	//合成丹药数量
	SynthetiseNum int32 `json:"synthetise_num"`

	//需要物品id1
	NeedItemId1 int32 `json:"need_item_id1"`

	//需要物品数量1
	NeedItemNum1 int32 `json:"need_item_num1"`

	//需要物品id2
	NeedItemId2 int32 `json:"need_item_id2"`

	//需要物品数量2
	NeedItemNum2 int32 `json:"need_item_num2"`

	//需要物品id3
	NeedItemId3 int32 `json:"need_item_id3"`

	//需要物品数量3
	NeedItemNum3 int32 `json:"need_item_num3"`

	//单个丹药练成所需时间(毫秒)
	Time int32 `json:"time"`

	//单个丹药加速消耗元宝
	AccelerateMoney int32 `json:"accelerate_money"`
}
