package entity

//交易物品数据
type TradeItemEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	PlayerId       int64  `gorm:"column:playerId"`
	OriginServerId int32  `gorm:"column:originServerId"`
	ServerId       int32  `gorm:"column:serverId"`
	ItemId         int32  `gorm:"column:itemId"`
	ItemNum        int32  `gorm:"column:itemNum"`
	Gold           int32  `gorm:"column:gold"`
	PropertyData   string `gorm:"column:porpertyData"`
	Level          int32  `gorm:"column:level"`
	Status         int32  `gorm:"column:status"`
	System         int32  `gorm:"column:system"`
	GlobalTradeId  int64  `gorm:"column:globalTradeId"`
	BuyPlatform    int32  `gorm:"column:buyPlatform"`
	BuyServerId    int32  `gorm:"column:buyServerId"`
	BuyPlayerId    int64  `gorm:"column:buyPlayerId"`
	BuyPlayerName  string `gorm:"column:buyPlayerName"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *TradeItemEntity) GetId() int64 {
	return e.Id
}

func (e *TradeItemEntity) TableName() string {
	return "t_trade_item"
}
