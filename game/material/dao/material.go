package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/material/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "material"
)

type MaterialDao interface {
	//查找玩家材料副本信息
	GetMaterialInfo(playerId int64) ([]*entity.PlayerMaterialEntity, error)
}

type materialDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *materialDao) GetMaterialInfo(playerId int64) (entityList []*entity.PlayerMaterialEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId = ? AND deleteTime=0", playerId).Error
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
	dao  *materialDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &materialDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetMaterialDao() MaterialDao {
	return dao
}
