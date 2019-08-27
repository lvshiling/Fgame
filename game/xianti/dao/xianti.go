package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	xiantientity "fgame/fgame/game/xianti/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "mount"
)

type XianTiDao interface {
	//查询玩家仙体信息
	GetXianTiEntity(playerId int64) (*xiantientity.PlayerXianTiEntity, error)
	//查询玩家非进阶仙体信息
	GetXianTiOtherList(playerId int64) ([]*xiantientity.PlayerXianTiOtherEntity, error)
}

type xianTiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家仙体信息
func (dao *xianTiDao) GetXianTiEntity(playerId int64) (xiantiEntity *xiantientity.PlayerXianTiEntity, err error) {
	xiantiEntity = &xiantientity.PlayerXianTiEntity{}
	err = dao.ds.DB().First(xiantiEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询玩家非进阶仙体信息
func (dao *xianTiDao) GetXianTiOtherList(playerId int64) (xiantiOtherList []*xiantientity.PlayerXianTiOtherEntity, err error) {
	err = dao.ds.DB().Find(&xiantiOtherList, "playerId=? ", playerId).Error
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
	dao  *xianTiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xianTiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetXianTiDao() XianTiDao {
	return dao
}
