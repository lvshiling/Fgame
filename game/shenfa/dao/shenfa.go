package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/shenfa/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shenfa"
)

type ShenfaDao interface {
	//查询玩家身法信息
	GetShenfaEntity(playerId int64) (shenfaEntity *entity.PlayerShenfaEntity, err error)
	//查询玩家非进阶战翼信息
	GetShenfaOtherList(playerId int64) ([]*entity.PlayerShenfaOtherEntity, error)
}

type shenfaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shenfaDao) GetShenfaEntity(playerId int64) (shenfaEntity *entity.PlayerShenfaEntity, err error) {
	shenfaEntity = &entity.PlayerShenfaEntity{}
	err = dao.ds.DB().First(shenfaEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenfaDao) GetShenfaOtherList(playerId int64) (shenfaOtherList []*entity.PlayerShenfaOtherEntity, err error) {
	err = dao.ds.DB().Find(&shenfaOtherList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *shenfaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shenfaDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShenfaDao() ShenfaDao {
	return dao
}
