package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	shenmoentity "fgame/fgame/cross/shenmo/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName     = "shenmo_rank"
	dbNametime = "shenmo_rank_time"
)

type ShemMoDao interface {
	//查询神魔排行榜
	GetShenMoRankList(platform int32) ([]*shenmoentity.ShenMoRankEntity, error)
	//排行榜刷新时间戳
	GetShenMoRankTimeEntity(platform int32) (*shenmoentity.ShenMoRankTimeEntity, error)
}

type shenMoDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *shenMoDao) GetShenMoRankList(platform int32) (shenMoRankList []*shenmoentity.ShenMoRankEntity, err error) {
	err = dao.ds.DB().Find(&shenMoRankList, "platform=? and  deleteTime=0", platform).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *shenMoDao) GetShenMoRankTimeEntity(platform int32) (shenMoRankTimeEntity *shenmoentity.ShenMoRankTimeEntity, err error) {
	shenMoRankTimeEntity = &shenmoentity.ShenMoRankTimeEntity{}
	err = dao.ds.DB().First(shenMoRankTimeEntity, "platform=? and  deleteTime=0", platform).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbNametime, err)
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

func GetShenMoDao() ShemMoDao {
	return dao
}
