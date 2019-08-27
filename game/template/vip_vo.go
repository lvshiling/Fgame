/*此类自动生成,请勿修改*/
package template

/*vip模板配置*/
type VipTemplateVO struct {

	//id
	Id int `json:"id"`

	//vip等级
	Level int32 `json:"level"`

	//星级
	Star int32 `json:"star"`

	//下一级
	NextId int32 `json:"next_id"`

	//需要消费的元宝
	NeedValue int32 `json:"need_value"`

	//血量加成
	Hp int32 `json:"hp"`

	//攻击加成
	Attack int32 `json:"attack"`

	//防御加成
	Defence int32 `json:"defence"`

	//礼包
	GiftId string `json:"gift_id"`

	//数量
	GiftCount string `json:"gift_count"`

	//礼物银两
	GiftSilver int64 `json:"gift_silver"`

	//免费礼包
	FreeGiftId string `json:"free_gift_id"`

	//免费礼包数量
	FreeGiftCount string `json:"free_gift_count"`

	//免费礼物银两
	FreeGiftSilver int64 `json:"free_gift_silver"`

	//现价
	Price int32 `json:"price"`

	//原价
	CostPrice int32 `json:"cost_price"`

	//购买疲劳值上限
	BuyPilaoCount int32 `json:"buy_pilao_count"`

	//vip额外领取红包数量
	HbSnatchCount int32 `json:"buy_hb_count"`

	//VIP超生数量
	BabyChaoSheng int32 `json:"baby_chaosheng"`

	//手续费
	Shouxu int32 `json:"shouxu"`
}
