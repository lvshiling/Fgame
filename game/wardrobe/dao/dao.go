package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	wardrobeentity "fgame/fgame/game/wardrobe/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "wardrobe"
)

type WardrobeDao interface {
	//查询玩家衣橱列表
	GetWardrobeList(playerId int64) ([]*wardrobeentity.PlayerWardrobeEntity, error)
	//查询玩家套装培养资质丹
	GetWardrobePeiYangList(playerId int64) ([]*wardrobeentity.PlayerWardrobePeiYangEntity, error)
}

type wardrobeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *wardrobeDao) GetWardrobeList(playerId int64) (wardrobeList []*wardrobeentity.PlayerWardrobeEntity, err error) {
	err = dao.ds.DB().Find(&wardrobeList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *wardrobeDao) GetWardrobePeiYangList(playerId int64) (wardrobeList []*wardrobeentity.PlayerWardrobePeiYangEntity, err error) {
	err = dao.ds.DB().Find(&wardrobeList, "playerId=? ", playerId).Error
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
	dao  *wardrobeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &wardrobeDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetWardrobeDao() WardrobeDao {
	return dao
}
