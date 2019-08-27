package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	yuxientity "fgame/fgame/game/yuxi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "yuxi"
)

type YuXiDao interface {
	//获取玩家仙盟数据
	GetPlayerAllianceYuXi(playerId int64) (*yuxientity.PlayerAlliancYuXiEntity, error)
}

type yuxiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *yuxiDao) GetPlayerAllianceYuXi(playerId int64) (entity *yuxientity.PlayerAlliancYuXiEntity, err error) {
	entity = &yuxientity.PlayerAlliancYuXiEntity{}
	err = dao.ds.DB().First(entity, "deleteTime=0 and playerId=?", playerId).Error
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
	dao  *yuxiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &yuxiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetYuXiDao() YuXiDao {
	return dao
}
