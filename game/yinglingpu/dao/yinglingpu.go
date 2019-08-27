package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	yinglingpuentity "fgame/fgame/game/yinglingpu/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "yinglingpu"
)

type YingLingPuDao interface {
	GetYingLingPu(p_playerId int64) ([]*yinglingpuentity.YingLingPuEntity, error)
	GetYingLingPuSuiPian(p_playerId int64) ([]*yinglingpuentity.YingLingPuSuiPianEntity, error)
}

type yingLingPuDao struct {
	ds coredb.DBService
}

func (t *yingLingPuDao) GetYingLingPu(p_playerId int64) ([]*yinglingpuentity.YingLingPuEntity, error) {
	rst := make([]*yinglingpuentity.YingLingPuEntity, 0, 8)
	exdb := t.ds.DB().Where("playerid = ? and deleteTime =0 ", p_playerId).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, coredb.NewDBError(dbName, exdb.Error)
	}
	return rst, nil
}

func (t *yingLingPuDao) GetYingLingPuSuiPian(p_playerId int64) ([]*yinglingpuentity.YingLingPuSuiPianEntity, error) {
	rst := make([]*yinglingpuentity.YingLingPuSuiPianEntity, 0, 8)
	exdb := t.ds.DB().Where("playerid = ? and deleteTime =0", p_playerId).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, coredb.NewDBError(dbName, exdb.Error)
	}
	return rst, nil
}

var (
	once sync.Once
	dao  *yingLingPuDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &yingLingPuDao{
			ds: ds,
		}
	})
	return nil
}

func GetYingLingPuDao() YingLingPuDao {
	return dao
}
