package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	wingentity "fgame/fgame/game/wing/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "wing"
)

type WingDao interface {
	//查询玩家战翼试用阶数
	GetWingTrialEntity(playerId int64) (*wingentity.PlayerWingTrialEntity, error)
	//查询玩家战翼信息
	GetWingEntity(playerId int64) (*wingentity.PlayerWingEntity, error)
	//查询玩家非进阶战翼信息
	GetWingOtherList(playerId int64) ([]*wingentity.PlayerWingOtherEntity, error)
}

type wingDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *wingDao) GetWingTrialEntity(playerId int64) (wingTrialEntity *wingentity.PlayerWingTrialEntity, err error) {
	wingTrialEntity = &wingentity.PlayerWingTrialEntity{}
	err = dao.ds.DB().First(wingTrialEntity, "playerId=? and deleteTime =0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("wing_trial", err)
		}
		return nil, nil
	}
	return
}

func (dao *wingDao) GetWingEntity(playerId int64) (wingEntity *wingentity.PlayerWingEntity, err error) {
	wingEntity = &wingentity.PlayerWingEntity{}
	err = dao.ds.DB().First(wingEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询玩家非进阶战翼信息
func (dao *wingDao) GetWingOtherList(playerId int64) (wingOtherList []*wingentity.PlayerWingOtherEntity, err error) {
	err = dao.ds.DB().Find(&wingOtherList, "playerId=? ", playerId).Error
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

func GetWingDao() WingDao {
	return dao
}
