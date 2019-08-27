package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	chargeentity "fgame/fgame/game/charge/entity"
	"fgame/fgame/game/charge/types"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "charge"
)

type ChargeDao interface {
	//获取玩家扶持充值记录
	GetPlayerPrivilegeChargeEntityList(playerId int64) (entityList []*chargeentity.PlayerPrivilegeChargeEntity, err error)
	//获取玩家充值记录
	GetPlayerChargeEntityList(playerId int64) (entityList []*chargeentity.PlayerChargeEntity, err error)
	//获取玩家档次首充记录
	GetPlayerFirstChargeRecordEntityList(playerId int64) (entityList []*chargeentity.PlayerFirstChargeRecordEntity, err error)
	//获取玩家新档次首冲记录
	GetPlayerNewFirstChargeRecordEntity(playerId int64) (entity *chargeentity.PlayerNewFirstChargeRecordEntity, err error)
	//获取玩家每日首充记录
	GetPlayerCycleChargeRecordEntity(playerId int64) (entity *chargeentity.PlayerCycleChargeRecordEntity, err error)
	//获取首冲时间
	GetFirstCharge(serverId int32) (firstChargeEntity *chargeentity.FirstChargeEntity, err error)
	//获取新首充时间
	GetNewFirstCharge(serverId int32) (firstChargeEntity *chargeentity.NewFirstChargeEntity, err error)
	//获取所有订单列表
	GetOrderList(serverId int32, orderStatus types.OrderStatus) (eList []*chargeentity.OrderEntity, err error)
	//获取所有后台充值列表
	GetPrivilegeChargeList(serverId int32, orderStatus types.OrderStatus) (eList []*chargeentity.PrivilegeChargeEntity, err error)
	//获取订单
	GetOrder(serverId int32, orderId string) (*chargeentity.OrderEntity, error)
}
type chargeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *chargeDao) GetPlayerPrivilegeChargeEntityList(playerId int64) (entityList []*chargeentity.PlayerPrivilegeChargeEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chargeDao) GetPlayerChargeEntityList(playerId int64) (entityList []*chargeentity.PlayerChargeEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chargeDao) GetPlayerFirstChargeRecordEntityList(playerId int64) (entityList []*chargeentity.PlayerFirstChargeRecordEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chargeDao) GetPlayerNewFirstChargeRecordEntity(playerId int64) (entity *chargeentity.PlayerNewFirstChargeRecordEntity, err error) {
	entity = &chargeentity.PlayerNewFirstChargeRecordEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chargeDao) GetPlayerCycleChargeRecordEntity(playerId int64) (entity *chargeentity.PlayerCycleChargeRecordEntity, err error) {
	entity = &chargeentity.PlayerCycleChargeRecordEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chargeDao) GetFirstCharge(serverId int32) (firstChargeEntity *chargeentity.FirstChargeEntity, err error) {
	firstChargeEntity = &chargeentity.FirstChargeEntity{}
	err = dao.ds.DB().First(firstChargeEntity, "serverId=? and deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *chargeDao) GetNewFirstCharge(serverId int32) (firstChargeEntity *chargeentity.NewFirstChargeEntity, err error) {
	firstChargeEntity = &chargeentity.NewFirstChargeEntity{}
	err = dao.ds.DB().First(firstChargeEntity, "serverId=? and deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取订单列表
func (dao *chargeDao) GetOrderList(serverId int32, orderStatus types.OrderStatus) (eList []*chargeentity.OrderEntity, err error) {
	err = dao.ds.DB().Find(&eList, "serverId=? and orderStatus=? and deleteTime=0", serverId, int32(orderStatus)).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取订单列表
func (dao *chargeDao) GetPrivilegeChargeList(serverId int32, orderStatus types.OrderStatus) (eList []*chargeentity.PrivilegeChargeEntity, err error) {
	err = dao.ds.DB().Find(&eList, "serverId=? and status=? and deleteTime=0", serverId, int32(orderStatus)).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取订单
func (dao *chargeDao) GetOrder(serverId int32, orderId string) (e *chargeentity.OrderEntity, err error) {
	e = &chargeentity.OrderEntity{}
	err = dao.ds.DB().First(e, "serverId=? and orderId=? and deleteTime=0", serverId, orderId).Error
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
	dao  *chargeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &chargeDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetChargeDao() ChargeDao {
	return dao
}
