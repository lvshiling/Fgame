package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/transportation/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "biaoche"
)

type TransportationDao interface {
	//查找玩家镖车信息
	GetPlayerTransportInfo(playerId int64) (*entity.PlayerTransportationEntity, error)
	//查找镖车列表
	GetTransportList(serverId int32) (transportList []*entity.TransportationEntity, err error)
}

type transportationDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *transportationDao) GetPlayerTransportInfo(playerId int64) (transprot *entity.PlayerTransportationEntity, err error) {
	transprot = &entity.PlayerTransportationEntity{}
	err = dao.ds.DB().First(&transprot, "playerId = ? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *transportationDao) GetTransportList(serverId int32) (transportList []*entity.TransportationEntity, err error) {
	err = dao.ds.DB().Find(&transportList, "deleteTime = 0 AND serverId=?", serverId).Error
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
	tdao *transportationDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		tdao = &transportationDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetTransportationDao() TransportationDao {
	return tdao
}
