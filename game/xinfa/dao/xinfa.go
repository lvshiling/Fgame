package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	xinfaentity "fgame/fgame/game/xinfa/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "xinfa"
)

type XinFaDao interface {
	//查询玩家心法列表
	GetXinFaList(playerId int64) ([]*xinfaentity.PlayerXinFaEntity, error)
}

type xinFaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xinFaDao) GetXinFaList(playerId int64) (xinFaList []*xinfaentity.PlayerXinFaEntity, err error) {
	err = dao.ds.DB().Find(&xinFaList, "playerId=? ", playerId).Error
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
	dao  *xinFaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xinFaDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetXinFaDao() XinFaDao {
	return dao
}
