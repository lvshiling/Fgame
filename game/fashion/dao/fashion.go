package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	fashionentity "fgame/fgame/game/fashion/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "Fashion"
)

type FashionDao interface {
	//查询玩家时装穿戴
	GetFashionWearEntity(playerId int64) (*fashionentity.PlayerWearFashionEntity, error)
	//查询玩家时装列表
	GetFashionList(playerId int64) ([]*fashionentity.PlayerFashionEntity, error)
	//查询玩家时装试用列表
	GetFashionTrialList(playerId int64) ([]*fashionentity.PlayerFashionTrialEntity, error)
}

type fashionDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *fashionDao) GetFashionWearEntity(playerId int64) (fashionWear *fashionentity.PlayerWearFashionEntity, err error) {
	fashionWear = &fashionentity.PlayerWearFashionEntity{}
	err = dao.ds.DB().First(fashionWear, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *fashionDao) GetFashionList(playerId int64) (fashionList []*fashionentity.PlayerFashionEntity, err error) {
	err = dao.ds.DB().Find(&fashionList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *fashionDao) GetFashionTrialList(playerId int64) (fashionTrialList []*fashionentity.PlayerFashionTrialEntity, err error) {
	err = dao.ds.DB().Find(&fashionTrialList, "playerId=? AND deleteTime=0 ", playerId).Error
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
	dao  *fashionDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &fashionDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFashionDao() FashionDao {
	return dao
}
