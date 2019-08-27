package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/trade/entity"
	tradetypes "fgame/fgame/game/trade/types"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "trade"
)

type TradeDao interface {
	//获取交易物品
	GetTradeItemList(serverId int32, status tradetypes.TradeStatus) ([]*entity.TradeItemEntity, error)
	//获取订单
	GetTradeOrderList(serverId int32, status tradetypes.TradeOrderStatus) ([]*entity.TradeOrderEntity, error)
	//获取交易日志
	GetTradeLogList(playerId int64) ([]*entity.PlayerTradeLogEntity, error)
	//获取回收数据
	GetTradeRecycle(serverId int32) (*entity.TradeRecycleEntity, error)
	//获取个人回收数据
	GetPlayerTradeRecycle(playerId int64) (*entity.PlayerTradeRecycleEntity, error)
}

type tradeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *tradeDao) GetTradeItemList(serverId int32, status tradetypes.TradeStatus) (tradeItemList []*entity.TradeItemEntity, err error) {
	err = dao.ds.DB().Order("createTime desc").Find(&tradeItemList, "deleteTime = 0 AND serverId=? and status=?", serverId, int32(status)).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *tradeDao) GetTradeOrderList(serverId int32, status tradetypes.TradeOrderStatus) (tradeOrderList []*entity.TradeOrderEntity, err error) {
	err = dao.ds.DB().Order("createTime desc").Find(&tradeOrderList, "deleteTime = 0 AND serverId=? and status=?", serverId, int32(status)).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *tradeDao) GetTradeLogList(playerId int64) (tradeLogList []*entity.PlayerTradeLogEntity, err error) {
	err = dao.ds.DB().Order("createTime desc").Find(&tradeLogList, "deleteTime = 0 AND playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *tradeDao) GetTradeRecycle(serverId int32) (e *entity.TradeRecycleEntity, err error) {
	e = &entity.TradeRecycleEntity{}
	err = dao.ds.DB().First(e, "deleteTime = 0 AND serverId=?", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *tradeDao) GetPlayerTradeRecycle(playerId int64) (e *entity.PlayerTradeRecycleEntity, err error) {
	e = &entity.PlayerTradeRecycleEntity{}
	err = dao.ds.DB().First(e, "deleteTime = 0 AND playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	tdao *tradeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		tdao = &tradeDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetTradeDao() TradeDao {
	return tdao
}
