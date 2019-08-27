package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/global"
	hongbaoentity "fgame/fgame/game/hongbao/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "hongbao"
)

type HongBaoDao interface {
	//获取所有红包信息
	GetAllHongBaoEntity(existTime int64) ([]*hongbaoentity.HongBaoEntity, error)
	//查询玩家红包信息
	GetPlayerHongBaoEntity(playerId int64) (*hongbaoentity.PlayerHongBaoEntity, error)
}

type hongbaoDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *hongbaoDao) GetAllHongBaoEntity(existTime int64) (entityList []*hongbaoentity.HongBaoEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime=0 AND serverId = ? AND createTime > ?", global.GetGame().GetServerIndex(), existTime).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *hongbaoDao) GetPlayerHongBaoEntity(playerId int64) (hongbaoEntity *hongbaoentity.PlayerHongBaoEntity, err error) {
	hongbaoEntity = &hongbaoentity.PlayerHongBaoEntity{}
	err = dao.ds.DB().First(hongbaoEntity, "deleteTime=0 AND playerId=?", playerId).Error
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
	dao  *hongbaoDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &hongbaoDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetHongBaoDao() HongBaoDao {
	return dao
}
