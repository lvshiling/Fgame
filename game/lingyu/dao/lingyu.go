package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/lingyu/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "lingyu"
)

type LingyuDao interface {
	//查询玩家领域信息
	GetLingyuEntity(playerId int64) (lingyuEntity *entity.PlayerLingyuEntity, err error)
	//查询玩家非进阶领域信息
	GetLingyuOtherList(playerId int64) ([]*entity.PlayerLingyuOtherEntity, error)
}

type lingyuDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *lingyuDao) GetLingyuEntity(playerId int64) (lingyuEntity *entity.PlayerLingyuEntity, err error) {
	lingyuEntity = &entity.PlayerLingyuEntity{}
	err = dao.ds.DB().First(lingyuEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *lingyuDao) GetLingyuOtherList(playerId int64) (lingyuOtherList []*entity.PlayerLingyuOtherEntity, err error) {
	err = dao.ds.DB().Find(&lingyuOtherList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *lingyuDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &lingyuDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetLingyuDao() LingyuDao {
	return dao
}
