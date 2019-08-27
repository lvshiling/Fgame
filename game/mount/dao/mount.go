package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	mountentity "fgame/fgame/game/mount/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "mount"
)

type MountDao interface {
	//查询玩家坐骑信息
	GetMountEntity(playerId int64) (*mountentity.PlayerMountEntity, error)
	//查询玩家非进阶坐骑信息
	GetMountOtherList(playerId int64) ([]*mountentity.PlayerMountOtherEntity, error)
}

type mountDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *mountDao) GetMountEntity(playerId int64) (mountEntity *mountentity.PlayerMountEntity, err error) {
	mountEntity = &mountentity.PlayerMountEntity{}
	err = dao.ds.DB().First(mountEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询玩家非进阶坐骑信息
func (dao *mountDao) GetMountOtherList(playerId int64) (mountOtherList []*mountentity.PlayerMountOtherEntity, err error) {
	err = dao.ds.DB().Find(&mountOtherList, "playerId=? ", playerId).Error
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
	dao  *mountDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &mountDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMountDao() MountDao {
	return dao
}
