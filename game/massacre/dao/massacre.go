package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	massacreentity "fgame/fgame/game/massacre/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "massacre"
)

type MassacreDao interface {
	//查询玩家戮仙刃
	GetMassacreEntity(playerId int64) (*massacreentity.PlayerMassacreEntity, error)
}

type massacreDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *massacreDao) GetMassacreEntity(playerId int64) (massacreEntity *massacreentity.PlayerMassacreEntity, err error) {
	massacreEntity = &massacreentity.PlayerMassacreEntity{}
	err = dao.ds.DB().First(massacreEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *massacreDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &massacreDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMassacreDao() MassacreDao {
	return dao
}
