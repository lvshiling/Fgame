package entity

type TradeItem struct {
	Id                int64  `gorm:"primary_key;column:id"`
	Platform          int32  `gorm:"column:platform"`
	ServerId          int32  `gorm:"column:serverId"`
	TradeId           int64  `gorm:"column:tradeId"`
	PlayerId          int64  `gorm:"column:playerId"`
	PlayerName        string `gorm:"column:playerName"`
	ItemId            int32  `gorm:"column:itemId"`
	ItemNum           int32  `gorm:"column:itemNum"`
	Level             int32  `gorm:"column:level"`
	Gold              int64  `gorm:"column:gold"`
	PropertyData      string `gorm:"column:propertyData"`
	Status            int32  `gorm:"column:status"`
	BuyPlayerPlatform int32  `gorm:"column:buyPlayerPlatform"`
	BuyPlayerServerId int32  `gorm:"column:buyPlayerServerId"`
	BuyPlayerId       int64  `gorm:"column:buyPlayerId"`
	BuyPlayerName     string `gorm:"column:buyPlayerName"`
	UpdateTime        int64  `gorm:"column:updateTime"`
	CreateTime        int64  `gorm:"column:createTime"`
	DeleteTime        int64  `gorm:"column:deleteTime"`
}

func (m *TradeItem) TableName() string {
	return "t_trade_item"
}
