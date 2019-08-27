/*此类自动生成,请勿修改*/
package template

/*装备模板*/
type EquipTemplateVO struct {

	//id
	Id int `json:"id"`

	//装备强化上限
	MaxStrengthen int `json:"max_strengthen"`

	//镶嵌宝石类型
	StoneType int32 `json:"stone_type"`

	//装备添加状态
	AddStatusid int32 `json:"add_statusid"`

	//套装
	Series int32 `json:"series"`

	//进阶需要物品1
	NeedItem1 int32 `json:"need_item1"`

	//进阶需要物品1数量
	NeedItemNum1 int32 `json:"need_item_num1"`

	//进阶需要物品2
	NeedItem2 int32 `json:"need_item2"`

	//进阶需要物品2数量
	NeedItemNum2 int32 `json:"need_item_num2"`

	//进阶需要物品3
	NeedItem3 int32 `json:"need_item3"`

	//进阶需要物品3数量
	NeedItemNum3 int32 `json:"need_item_num3"`

	//进阶成功几率
	SuccessRate int32 `json:"success_rate"`

	//下一阶物品
	Next int32 `json:"next"`

	//套装id
	TaozhuangId int32 `json:"taozhuang_id"`

	//血量
	Hp int64 `json:"hp"`

	//攻击
	Attack int64 `json:"attack"`

	//防御
	Defence int64 `json:"defence"`
}
