package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/global"
	onearenaentity "fgame/fgame/game/onearena/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "onearena"
)

type OneArenaDao interface {
	//灵池信息
	GetOneArenaList() ([]*onearenaentity.OneArenaEntity, error)
	//查询玩家灵池争夺信息
	GetPlayerOneArenaEntity(playerId int64) (*onearenaentity.PlayerOneArenaEntity, error)
	//查询玩家灵池争夺记录信息
	GetPlayerOneArenaRecordList(playerId int64) ([]*onearenaentity.PlayerOneArenaRecordEntity, error)
	//查询玩家灵池被抢记录信息
	GetPlayerOneArenaRobbedList(playerId int64) ([]*onearenaentity.PlayerOneArenaRobbedEntity, error)
	//获取灵池产出的鲲
	GetPlayerOneArenaKunList(playerId int64) ([]*onearenaentity.PlayerOneArenaKunEntity, error)
}

type oneArenaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *oneArenaDao) GetOneArenaList() (onearenaList []*onearenaentity.OneArenaEntity, err error) {
	err = dao.ds.DB().Find(&onearenaList, "serverId=? and deleteTime=0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *oneArenaDao) GetPlayerOneArenaEntity(playerId int64) (onearenaEntity *onearenaentity.PlayerOneArenaEntity, err error) {
	onearenaEntity = &onearenaentity.PlayerOneArenaEntity{}
	err = dao.ds.DB().First(onearenaEntity, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *oneArenaDao) GetPlayerOneArenaRecordList(playerId int64) (recordList []*onearenaentity.PlayerOneArenaRecordEntity, err error) {
	err = dao.ds.DB().Find(&recordList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *oneArenaDao) GetPlayerOneArenaRobbedList(playerId int64) (robbedList []*onearenaentity.PlayerOneArenaRobbedEntity, err error) {
	err = dao.ds.DB().Order("`robTime` DESC").Find(&robbedList, "playerId=? and deleteTime = 0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *oneArenaDao) GetPlayerOneArenaKunList(playerId int64) (kunList []*onearenaentity.PlayerOneArenaKunEntity, err error) {
	err = dao.ds.DB().Find(&kunList, "playerId=? and deleteTime = 0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *oneArenaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &oneArenaDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetOneArenaDao() OneArenaDao {
	return dao
}
