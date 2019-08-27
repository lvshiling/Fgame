/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerTradeBuy)(nil))
}

/*交易所购买*/
type PlayerTradeBuy struct {
	PlayerTradeLogMsg `bson:",inline"`

	//物品id
	ItemId int32 `json:"itemId"`

	//物品数量
	ItemNum int32 `json:"itemNum"`

	//价格
	Gold int32 `json:"gold"`

	//原因编号
	Reason int32 `json:"reason"`

	//原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerTradeBuy) LogName() string {
	return "player_trade_buy"
}
