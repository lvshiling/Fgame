package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	treasureboxentity "fgame/fgame/cross/treasurebox/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "treasurebox_log"
)

type TreasureBoxDao interface {
	//查询玩家
	GetTreasureBoxLogList(platform int32, num int32) ([]*treasureboxentity.TreasureBoxLogEntity, error)
}

type treasureBoxDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *treasureBoxDao) GetTreasureBoxLogList(platform int32, num int32) (boxLogList []*treasureboxentity.TreasureBoxLogEntity, err error) {
	err = dao.ds.DB().Limit(num).Order("`lastTime` DESC").Find(&boxLogList, "areaId=?", platform).Error
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
	dao  *treasureBoxDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &treasureBoxDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTreasureBoXDao() TreasureBoxDao {
	return dao
}
