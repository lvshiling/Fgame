/*此类自动生成,请勿修改*/
package template

/*天书配置*/
type TianShuTemplateVO struct {

	//id
	Id int `json:"id"`

	//天书类型
	Type int32 `json:"type"`

	//天书等级
	Level int32 `json:"level"`

	//升级成功率
	SuccessRate int32 `json:"success_rate"`

	//升级消耗物品id
	LevelItem string `json:"level_item"`

	//升级消耗物品数量
	LevelItemCount string `json:"level_item_count"`

	//下一级id
	NextId int32 `json:"next_id"`

	//激活需要充值的数量
	NeedGold int32 `json:"need_gold"`

	//激活属性
	Hp int32 `json:"hp"`

	//激活属性
	Attack int32 `json:"attack"`

	//激活属性
	Defence int32 `json:"defence"`

	//激活奖励id
	FreeGiftId string `json:"free_gift_id"`

	//激活奖励数量
	FreeGiftCount string `json:"free_gift_count"`

	//激活奖励银两
	FreeGiftSilver int32 `json:"free_gift_silver"`

	//激活奖励绑元
	FreeGiftBindgold int32 `json:"free_gift_bindgold"`

	//激活奖励元宝
	FreeGiftGold int32 `json:"free_gift_gold"`

	//特权比率
	Tequan int32 `json:"tequan"`
}
