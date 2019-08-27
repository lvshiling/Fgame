package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	arenapvpentity "fgame/fgame/game/arenapvp/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "arenapvp"
)

type ArenapvpDao interface {
	//获取竞技场
	GetPlayerArenapvp(playerId int64) (*arenapvpentity.PlayerArenapvpEntity, error)
	//获取玩家竞技场竞猜日志列表
	GetPlayerArenapvpGuessLogList(playerId int64) ([]*arenapvpentity.PlayerArenapvpGuessLogEntity, error)
	//获取竞技场竞猜日志列表
	GetArenapvpGuessRecordList(serverId int32) ([]*arenapvpentity.ArenapvpGuessRecordEntity, error)
}

type arenapvpDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *arenapvpDao) GetPlayerArenapvp(playerId int64) (arenapvpEntity *arenapvpentity.PlayerArenapvpEntity, err error) {
	arenapvpEntity = &arenapvpentity.PlayerArenapvpEntity{}
	err = dao.ds.DB().First(arenapvpEntity, "playerId=? AND deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}

		return nil, nil
	}
	return
}

//获取玩家竞技场竞猜日志列表
func (dao *arenapvpDao) GetPlayerArenapvpGuessLogList(playerId int64) (entityList []*arenapvpentity.PlayerArenapvpGuessLogEntity, err error) {
	err = dao.ds.DB().Order("createTime ASC").Find(&entityList, "playerId=? AND deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}

		return nil, nil
	}
	return
}

//获取竞技场竞猜日志列表
func (dao *arenapvpDao) GetArenapvpGuessRecordList(serverId int32) (entityList []*arenapvpentity.ArenapvpGuessRecordEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "serverId=? AND deleteTime=0 ", serverId).Error
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
	dao  *arenapvpDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &arenapvpDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetArenapvpDao() ArenapvpDao {
	return dao
}
