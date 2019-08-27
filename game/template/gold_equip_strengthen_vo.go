/*此类自动生成,请勿修改*/
package template

/*元神金装强化配置*/
type GoldEquipStrengthenTemplateVO struct {

	//id
	Id int `json:"id"`

	//套装id
	Level int32 `json:"level"`

	//强化id
	NextId int32 `json:"next_id"`

	//可用于强化的物品id
	NeedItemId string `json:"need_item_id"`

	//0强化到1的概念
	GiveRate1 int32 `json:"give_rate1"`

	//1强化到2的概念
	GiveRate2 int32 `json:"give_rate2"`

	//2强化到3的概念
	GiveRate3 int32 `json:"give_rate3"`

	//3强化到4的概念
	GiveRate4 int32 `json:"give_rate4"`

	//4强化到5的概念
	GiveRate5 int32 `json:"give_rate5"`

	//5强化到6的概念
	GiveRate6 int32 `json:"give_rate6"`

	//6强化到7的概念
	GiveRate7 int32 `json:"give_rate7"`

	//7强化到8的概念
	GiveRate8 int32 `json:"give_rate8"`

	//8强化到9的概念
	GiveRate9 int32 `json:"give_rate9"`

	//9强化到10的概念
	GiveRate10 int32 `json:"give_rate10"`

	//生命
	Hp int32 `json:"hp"`

	//攻击
	Attack int32 `json:"attack"`

	//防御
	Defence int32 `json:"defence"`

	//生命加成万分比
	HpPercent int32 `json:"hp_percent"`

	//攻击加成万分比
	AttPercent int32 `json:"att_percent"`

	//防御加成万分比
	DefPercent int32 `json:"def_percent"`

	//返还物品id
	MeltingReturnId string `json:"melting_return_id"`

	//返还物品数量
	MeltingReturnCount string `json:"melting_return_count"`
}
