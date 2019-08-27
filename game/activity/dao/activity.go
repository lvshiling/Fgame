package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	activityentity "fgame/fgame/game/activity/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "activity"
)

type ActivityDao interface {
	//玩家活动大厅列表
	GetActivitys(playerId int64) ([]*activityentity.PlayerActivityEntity, error)
	GetActivityPkList(playerId int64) ([]*activityentity.PlayerActivityPkEntity, error)
	GetActivityRankList(playerId int64) ([]*activityentity.PlayerActivityRankEntity, error)
	GetActivityCollectList(playerId int64) ([]*activityentity.PlayerActivityCollectEntity, error)
	GetActivityEndRecordList(serverId int32) ([]*activityentity.ActivityEndRecordEntity, error)
}

type activityDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *activityDao) GetActivitys(playerId int64) (activitys []*activityentity.PlayerActivityEntity, err error) {
	err = dao.ds.DB().Find(&activitys, "playerId = ? AND deleteTime = 0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}

		return nil, nil
	}
	return
}

func (dao *activityDao) GetActivityPkList(playerId int64) (activitys []*activityentity.PlayerActivityPkEntity, err error) {
	err = dao.ds.DB().Find(&activitys, "playerId = ? AND deleteTime = 0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}

		return nil, nil
	}
	return
}

func (dao *activityDao) GetActivityRankList(playerId int64) (activitys []*activityentity.PlayerActivityRankEntity, err error) {
	err = dao.ds.DB().Find(&activitys, "playerId = ? AND deleteTime = 0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}

		return nil, nil
	}
	return
}

func (dao *activityDao) GetActivityCollectList(playerId int64) (entityList []*activityentity.PlayerActivityCollectEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *activityDao) GetActivityEndRecordList(serverId int32) (entityList []*activityentity.ActivityEndRecordEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "serverId=? AND deleteTime=0", serverId).Error
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
	dao  *activityDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &activityDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetActivityDao() ActivityDao {
	return dao
}
