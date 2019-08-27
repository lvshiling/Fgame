package entity

//交易物品数据
type PlayerTradeLogEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	PlayerId       int64  `gorm:"column:playerId"`
	LogType        int32  `gorm:"column:logType"`
	TradeId        int64  `gorm:"column:tradeId"`
	SellServerId   int32  `gorm:"column:sellServerId"`
	SellPlayerId   int64  `gorm:"column:sellPlayerId"`
	SellPlayerName string `gorm:"column:sellPlayerName"`
	BuyServerId    int32  `gorm:"column:buyServerId"`
	BuyPlayerId    int64  `gorm:"column:buyPlayerId"`
	BuyPlayerName  string `gorm:"column:buyPlayerName"`
	GetGold        int32  `gorm:"column:getGold"`
	Gold           int32  `gorm:"column:gold"`
	Fee            int32  `gorm:"column:fee"`
	FeeRate        int32  `gorm:"column:feeRate"`
	ItemId         int32  `gorm:"column:itemId"`
	ItemNum        int32  `gorm:"column:itemNum"`
	PropertyData   string `gorm:"column:propertyData"`
	Level          int32  `gorm:"column:level"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *PlayerTradeLogEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTradeLogEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTradeLogEntity) TableName() string {
	return "t_player_trade_log"
}
