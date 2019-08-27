package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	fabaoentity "fgame/fgame/game/fabao/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "fabao"
)

type FaBaoDao interface {
	//查询玩家法宝信息
	GetFaBaoEntity(playerId int64) (*fabaoentity.PlayerFaBaoEntity, error)
	//查询玩家非进阶法宝信息
	GetFaBaoOtherList(playerId int64) ([]*fabaoentity.PlayerFaBaoOtherEntity, error)
}

type wingDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *wingDao) GetFaBaoEntity(playerId int64) (faBaoEntity *fabaoentity.PlayerFaBaoEntity, err error) {
	faBaoEntity = &fabaoentity.PlayerFaBaoEntity{}
	err = dao.ds.DB().First(faBaoEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询玩家非进阶法宝信息
func (dao *wingDao) GetFaBaoOtherList(playerId int64) (faBaoOtherList []*fabaoentity.PlayerFaBaoOtherEntity, err error) {
	err = dao.ds.DB().Find(&faBaoOtherList, "playerId=? ", playerId).Error
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
	dao  *wingDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &wingDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFaBaoDao() FaBaoDao {
	return dao
}
