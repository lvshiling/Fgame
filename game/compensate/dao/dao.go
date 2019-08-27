package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	compensateentity "fgame/fgame/game/compensate/entity"
	"fgame/fgame/game/global"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "compensate"
)

type CompensateDao interface {
	// 获取全服补偿数据
	GetCompensateEntityList() (compensateEntityList []*compensateentity.CompensateEntity, err error)
	// 获取玩家补偿数据
	GetPlayerCompensateEntityList(playerId int64) (entityList []*compensateentity.PlayerCompensateEntity, err error)
}

type compensateDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *compensateDao) GetCompensateEntityList() (compensateEntityList []*compensateentity.CompensateEntity, err error) {
	err = dao.ds.DB().Find(&compensateEntityList, "serverId=? AND deleteTime=0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *compensateDao) GetPlayerCompensateEntityList(playerId int64) (entityList []*compensateentity.PlayerCompensateEntity, err error) {
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
	dao  *compensateDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &compensateDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetCompensateDao() CompensateDao {
	return dao
}
