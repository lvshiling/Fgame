/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerShopBuy)(nil))
}

/*商店购买*/
type PlayerShopBuy struct {
	PlayerLogMsg `bson:",inline"`

	//商铺id
	ShopId int32 `json:"shopId"`

	//商品名字
	ShopName string `json:"shopName"`

	//购买数量
	BuyNum int32 `json:"buyNum"`

	//购买花费
	CostMoney int32 `json:"costMoney"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerShopBuy) LogName() string {
	return "player_shop_buy"
}
