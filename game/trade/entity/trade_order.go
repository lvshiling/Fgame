package entity

//交易订单数据
type TradeOrderEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	ServerId       int32  `gorm:"column:serverId"`
	PlayerId       int64  `gorm:"column:playerId"`
	PlayerName     string `gorm:"column:playerName"`
	BuyServerId    int32  `gorm:"column:buyServerId"`
	TradeId        int64  `gorm:"column:tradeId"`
	ItemId         int32  `gorm:"column:itemId"`
	ItemNum        int32  `gorm:"column:itemNum"`
	Gold           int32  `gorm:"column:gold"`
	PropertyData   string `gorm:"column:porpertyData"`
	Level          int32  `gorm:"column:level"`
	Status         int32  `gorm:"column:status"`
	SellPlatform   int32  `gorm:"column:sellPlatform"`
	SellServerId   int32  `gorm:"column:sellServerId"`
	SellPlayerId   int64  `gorm:"column:sellPlayerId"`
	SellPlayerName string `gorm:"column:sellPlayerName"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *TradeOrderEntity) GetId() int64 {
	return e.Id
}

func (e *TradeOrderEntity) TableName() string {
	return "t_trade_order"
}
