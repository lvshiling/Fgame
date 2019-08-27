package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shenqientity "fgame/fgame/game/shenqi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shenqi"
)

type ShenQiDao interface {
	//加载器灵槽位
	GetShenQiQiLingList(playerId int64) ([]*shenqientity.PlayerShenQiQiLingEntity, error)
	//加载淬炼槽位
	GetShenQiSmeltList(playerId int64) ([]*shenqientity.PlayerShenQiSmeltEntity, error)
	//加载碎片槽位
	GetShenQiDebrisList(playerId int64) ([]*shenqientity.PlayerShenQiDebrisEntity, error)
	//加载神器数据
	GetShenQiEntity(playerId int64) (*shenqientity.PlayerShenQiEntity, error)
}

type shenQiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shenQiDao) GetShenQiQiLingList(playerId int64) (slotList []*shenqientity.PlayerShenQiQiLingEntity, err error) {
	err = dao.ds.DB().Order("`shenQiType` ASC").Find(&slotList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenQiDao) GetShenQiSmeltList(playerId int64) (slotList []*shenqientity.PlayerShenQiSmeltEntity, err error) {
	err = dao.ds.DB().Order("`shenQiType` ASC").Find(&slotList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenQiDao) GetShenQiDebrisList(playerId int64) (slotList []*shenqientity.PlayerShenQiDebrisEntity, err error) {
	err = dao.ds.DB().Order("`shenQiType` ASC").Find(&slotList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenQiDao) GetShenQiEntity(playerId int64) (shenQiEntity *shenqientity.PlayerShenQiEntity, err error) {
	shenQiEntity = &shenqientity.PlayerShenQiEntity{}
	err = dao.ds.DB().First(shenQiEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *shenQiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shenQiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShenQiDao() ShenQiDao {
	return dao
}
