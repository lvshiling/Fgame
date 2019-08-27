package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	jxentity "fgame/fgame/game/juexue/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "jueXue"
)

type JueXueDao interface {
	//查询玩家绝学使用
	GetJueXueUseEntity(playerId int64) (*jxentity.PlayerJueXueUseEntity, error)
	//查询玩家绝学列表
	GetJueXueList(playerId int64) ([]*jxentity.PlayerJueXueEntity, error)
}

type jueXueDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *jueXueDao) GetJueXueUseEntity(playerId int64) (jueXueUseEntity *jxentity.PlayerJueXueUseEntity, err error) {
	jueXueUseEntity = &jxentity.PlayerJueXueUseEntity{}
	err = dao.ds.DB().First(jueXueUseEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *jueXueDao) GetJueXueList(playerId int64) (jueXueList []*jxentity.PlayerJueXueEntity, err error) {
	err = dao.ds.DB().Find(&jueXueList, "playerId=? ", playerId).Error
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
	dao  *jueXueDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &jueXueDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetJueXueDao() JueXueDao {
	return dao
}
