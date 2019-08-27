package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/shenmo/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "shenmo"
)

type ShenMoDao interface {
	//查询玩家神魔信息
	GetShenMoEntity(playerId int64) (shenmoEntity *entity.PlayerShenMoEntity, err error)
	//查询神魔排行榜
	GetShenMoRankList(serverId int32) ([]*entity.ShenMoRankEntity, error)
	//排行榜刷新时间戳
	GetShenMoRankTimeEntity(serverId int32) (*entity.ShenMoRankTimeEntity, error)
}

type shenMoDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shenMoDao) GetShenMoEntity(playerId int64) (shenmoEntity *entity.PlayerShenMoEntity, err error) {
	shenmoEntity = &entity.PlayerShenMoEntity{}
	err = dao.ds.DB().First(shenmoEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenMoDao) GetShenMoRankList(serverId int32) (shenMoRankList []*entity.ShenMoRankEntity, err error) {
	err = dao.ds.DB().Find(&shenMoRankList, "serverId=? and  deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenMoDao) GetShenMoRankTimeEntity(serverId int32) (shenMoRankTimeEntity *entity.ShenMoRankTimeEntity, err error) {
	shenMoRankTimeEntity = &entity.ShenMoRankTimeEntity{}
	err = dao.ds.DB().First(shenMoRankTimeEntity, "serverId=? and  deleteTime=0", serverId).Error
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
	dao  *shenMoDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &shenMoDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetShenMoDao() ShenMoDao {
	return dao
}
