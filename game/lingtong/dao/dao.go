package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/lingtong/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName        = "lingtong"
	dbNamefashion = "lingtong_fashion_trial"
)

type LingTongDao interface {
	//玩家灵童信息
	GetLingTongEntity(playerId int64) (*entity.PlayerLingTongEntity, error)
	//玩家激活灵童信息
	GetLingTongInfoList(playerId int64) ([]*entity.PlayerLingTongInfoEntity, error)
	//玩家灵童时装
	GetLingTongFashionList(playerId int64) ([]*entity.PlayerLingTongFashionEntity, error)
	//玩家灵童时装信息
	GetLingTongFashionInfoList(playerId int64) ([]*entity.PlayerLingTongFashionInfoEntity, error)
	//试用时装
	GetLingTongFashionTrialEntity(playerId int64) (*entity.PlayerLingTongFashionTrialEntity, error)
}

type lingTongDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *lingTongDao) GetLingTongEntity(playerId int64) (lingTongEntity *entity.PlayerLingTongEntity, err error) {
	lingTongEntity = &entity.PlayerLingTongEntity{}
	err = dao.ds.DB().First(lingTongEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *lingTongDao) GetLingTongInfoList(playerId int64) (lingTongInfoList []*entity.PlayerLingTongInfoEntity, err error) {
	err = dao.ds.DB().Find(&lingTongInfoList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *lingTongDao) GetLingTongFashionList(playerId int64) (lingTongFashionList []*entity.PlayerLingTongFashionEntity, err error) {
	err = dao.ds.DB().Find(&lingTongFashionList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *lingTongDao) GetLingTongFashionInfoList(playerId int64) (lingTongFashionInfoList []*entity.PlayerLingTongFashionInfoEntity, err error) {
	err = dao.ds.DB().Find(&lingTongFashionInfoList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *lingTongDao) GetLingTongFashionTrialEntity(playerId int64) (lingTongFashionTrialEntity *entity.PlayerLingTongFashionTrialEntity, err error) {
	lingTongFashionTrialEntity = &entity.PlayerLingTongFashionTrialEntity{}
	err = dao.ds.DB().First(lingTongFashionTrialEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbNamefashion, err)
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *lingTongDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &lingTongDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetLingTongDao() LingTongDao {
	return dao
}
