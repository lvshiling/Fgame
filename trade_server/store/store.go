package store

import (
	coredb "fgame/fgame/core/db"
	tradeservertypes "fgame/fgame/trade_server/types"

	"github.com/jinzhu/gorm"
)

type TradeItemEntity struct {
	Id                int64  `gorm:"primary_key;column:id"`
	Platform          int32  `gorm:"column:platform"`
	ServerId          int32  `gorm:"column:serverId"`
	TradeId           int64  `gorm:"column:tradeId"`
	PlayerId          int64  `gorm:"column:playerId"`
	PlayerName        string `gorm:"column:playerName"`
	ItemId            int32  `gorm:"column:itemId"`
	ItemNum           int32  `gorm:"column:itemNum"`
	Level             int32  `gorm:"column:level"`
	Gold              int32  `gorm:"column:gold"`
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

func (e *TradeItemEntity) TableName() string {
	return "t_trade_item"
}

type TradeStore interface {
	//获取所有服务器
	GetAll(status tradeservertypes.TradeStatus) ([]*TradeItemEntity, error)
	//获取交易数据
	GetTradeItemByTradeId(tradeId int64) (*TradeItemEntity, error)
	//获取交易数据
	GetTradeItemByLocalTradeId(tradeId int64) (*TradeItemEntity, error)
}

var (
	dbName = "trade_item"
)

type tradeStore struct {
	db coredb.DBService
}

func (s *tradeStore) GetAll(status tradeservertypes.TradeStatus) (eList []*TradeItemEntity, err error) {
	err = s.db.DB().Find(&eList, "deleteTime=0 and status=?", status).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		err = nil
		return
	}
	return
}

func (s *tradeStore) GetTradeItemByTradeId(tradeId int64) (e *TradeItemEntity, err error) {
	e = &TradeItemEntity{}
	err = s.db.DB().First(e, "deleteTime=0 and id=?", tradeId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (s *tradeStore) GetTradeItemByLocalTradeId(tradeId int64) (e *TradeItemEntity, err error) {
	e = &TradeItemEntity{}
	err = s.db.DB().First(e, "deleteTime=0 and tradeId=?", tradeId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func NewTradeStore(db coredb.DBService) TradeStore {
	s := &tradeStore{
		db: db,
	}
	return s
}
