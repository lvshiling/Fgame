package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/supremetitle/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "supreme_title"
)

type SupremeTitleDao interface {
	//查询玩家至尊称号穿戴
	GetSupremeTitleWearEntity(playerId int64) (*entity.PlayerWearSupremeTitleEntity, error)
	//查询玩家至尊称号列表
	GetSupremeTitleList(playerId int64) ([]*entity.PlayerSupremeTitleEntity, error)
}

type supremeTitleDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *supremeTitleDao) GetSupremeTitleWearEntity(playerId int64) (titleWear *entity.PlayerWearSupremeTitleEntity, err error) {
	titleWear = &entity.PlayerWearSupremeTitleEntity{}
	err = dao.ds.DB().First(titleWear, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *supremeTitleDao) GetSupremeTitleList(playerId int64) (titleList []*entity.PlayerSupremeTitleEntity, err error) {
	err = dao.ds.DB().Find(&titleList, "playerId=? ", playerId).Error
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
	dao  *supremeTitleDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &supremeTitleDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSupremeTitleDao() SupremeTitleDao {
	return dao
}
