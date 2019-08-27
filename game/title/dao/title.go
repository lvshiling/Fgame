package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	titleentity "fgame/fgame/game/title/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "title"
)

type TitleDao interface {
	//查询玩家称号穿戴
	GetTitleWearEntity(playerId int64) (*titleentity.PlayerWearTitleEntity, error)
	//查询玩家称号列表
	GetTitleList(playerId int64) ([]*titleentity.PlayerTitleEntity, error)
}

type titleDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *titleDao) GetTitleWearEntity(playerId int64) (titleWear *titleentity.PlayerWearTitleEntity, err error) {
	titleWear = &titleentity.PlayerWearTitleEntity{}
	err = dao.ds.DB().First(titleWear, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *titleDao) GetTitleList(playerId int64) (titleList []*titleentity.PlayerTitleEntity, err error) {
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
	dao  *titleDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &titleDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTitleDao() TitleDao {
	return dao
}
