package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/zhenfa/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "zhenfa"
)

type ZhenFaDao interface {
	//查询玩家阵法
	GetZhenFaList(playerId int64) ([]*entity.PlayerZhenFaEntity, error)
	//查询玩家阵旗
	GetZhenQiList(playerId int64) ([]*entity.PlayerZhenQiEntity, error)
	//查询玩家阵旗仙火
	GetZhenQiXianHuoList(playerId int64) ([]*entity.PlayerZhenQiXianHuoEntity, error)
	//查询玩家阵法战力
	GetZhenFaPowerEntity(playerId int64) (*entity.PlayerZhenFaPowerEntity, error)
}

type zhenFaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家阵法
func (dao *zhenFaDao) GetZhenFaList(playerId int64) (zhenFaList []*entity.PlayerZhenFaEntity, err error) {
	err = dao.ds.DB().Find(&zhenFaList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//查询玩家阵旗
func (dao *zhenFaDao) GetZhenQiList(playerId int64) (zhenQiList []*entity.PlayerZhenQiEntity, err error) {
	err = dao.ds.DB().Find(&zhenQiList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//查询玩家阵旗仙火
func (dao *zhenFaDao) GetZhenQiXianHuoList(playerId int64) (zhenQiXianHuoList []*entity.PlayerZhenQiXianHuoEntity, err error) {
	err = dao.ds.DB().Find(&zhenQiXianHuoList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *zhenFaDao) GetZhenFaPowerEntity(playerId int64) (powerEntity *entity.PlayerZhenFaPowerEntity, err error) {
	powerEntity = &entity.PlayerZhenFaPowerEntity{}
	err = dao.ds.DB().First(powerEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *zhenFaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &zhenFaDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetZhenFaDao() ZhenFaDao {
	return dao
}
