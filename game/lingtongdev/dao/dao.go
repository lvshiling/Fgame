package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/lingtongdev/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "lingtong"
)

type LingTongDevDao interface {
	//查询玩家灵童养成信息
	GetLingTongDevList(playerId int64) ([]*entity.PlayerLingTongDevEntity, error)
	//查询玩家非进阶灵童养成信息
	GetLingTongOtherList(playerId int64) ([]*entity.PlayerLingTongOtherEntity, error)
	//玩家灵童养成战力信息
	GetLingTongPowerList(playerId int64) ([]*entity.PlayerLingTongPowerEntity, error)
}

type lingTongDevDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家灵童养成信息
func (dao *lingTongDevDao) GetLingTongDevList(playerId int64) (lingTongDevList []*entity.PlayerLingTongDevEntity, err error) {
	err = dao.ds.DB().Find(&lingTongDevList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *lingTongDevDao) GetLingTongOtherList(playerId int64) (lingTongOtherList []*entity.PlayerLingTongOtherEntity, err error) {
	err = dao.ds.DB().Find(&lingTongOtherList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *lingTongDevDao) GetLingTongPowerList(playerId int64) (lingTongPowerList []*entity.PlayerLingTongPowerEntity, err error) {
	err = dao.ds.DB().Find(&lingTongPowerList, "playerId=? ", playerId).Error
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
	dao  *lingTongDevDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &lingTongDevDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetLingTongDevDao() LingTongDevDao {
	return dao
}
