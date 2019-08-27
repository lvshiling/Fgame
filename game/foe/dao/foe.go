package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	foeentity "fgame/fgame/game/foe/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "foe"
)

type FoeDao interface {
	//查询仇人列表
	GetFoeList(playerId int64) ([]*foeentity.PlayerFoeEntity, error)
	GetFoeProtect(playerId int64) (*foeentity.PlayerFoeProtectEntity, error)
	GetFoeFeedbackList(playerId int64) ([]*foeentity.PlayerFoeFeedbackEntity, error)
}

type foeDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *foeDao) GetFoeList(playerId int64) (foeEntityList []*foeentity.PlayerFoeEntity, err error) {
	err = dao.ds.DB().Order("`killTime` DESC").Find(&foeEntityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *foeDao) GetFoeProtect(playerId int64) (entity *foeentity.PlayerFoeProtectEntity, err error) {
	entity = &foeentity.PlayerFoeProtectEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *foeDao) GetFoeFeedbackList(playerId int64) (entityList []*foeentity.PlayerFoeFeedbackEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *foeDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &foeDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFoeDao() FoeDao {
	return dao
}
