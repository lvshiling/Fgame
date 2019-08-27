/*此类自动生成,请勿修改*/
package template

/*装备强化模板*/
type EquipStrengthenTemplateVO struct {

	//id
	Id int `json:"id"`

	//装备强化类型
	Type int32 `json:"type"`

	//下一阶id
	NextId int32 `json:"next_id"`

	//位置
	Position int32 `json:"position"`

	//等级
	Level int32 `json:"level"`

	//成功率
	SuccessRate int32 `json:"success_rate"`

	//回退概率
	ReturnRate int32 `json:"return_rate"`

	//回退id
	FailReturnStrengthenId int32 `json:"fail_return_strengthen_id"`

	//银两数
	SilverNum int32 `json:"silver_num"`

	//需要物品
	NeedItem int32 `json:"need_item"`

	//需要物品1数量
	NeedItemNum int32 `json:"need_item_num"`

	//血量
	Hp int64 `json:"hp"`

	//攻击
	Attack int64 `json:"attack"`

	//防御
	Defence int64 `json:"defence"`
}
