package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	emperorentity "fgame/fgame/game/emperor/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "emperor"
)

type EmperorDao interface {
	//查询玩家膜拜次数
	GetEmperorWorshipEntity(playerId int64) (*emperorentity.PlayerEmperorWorshipEntity, error)
	//查询当前帝王
	GetEmperorEntity(serverId int32) (*emperorentity.EmperorEntity, error)
	//查询龙椅抢夺记录
	GetEmperorRecordsList(serverId int32) ([]*emperorentity.EmperorRecordsEntity, error)
	//合服使用
	GetEmperorList(serverId int32) ([]*emperorentity.EmperorEntity, error)
}

type emperorDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家膜拜次数
func (dao *emperorDao) GetEmperorWorshipEntity(playerId int64) (Emperor *emperorentity.PlayerEmperorWorshipEntity, err error) {
	Emperor = &emperorentity.PlayerEmperorWorshipEntity{}
	err = dao.ds.DB().First(Emperor, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询当前帝王
func (dao *emperorDao) GetEmperorEntity(serverId int32) (emperor *emperorentity.EmperorEntity, err error) {
	emperor = &emperorentity.EmperorEntity{}
	err = dao.ds.DB().First(emperor, "serverId=? and deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//合服使用
func (dao *emperorDao) GetEmperorList(serverId int32) (emperorList []*emperorentity.EmperorEntity, err error) {
	err = dao.ds.DB().Find(&emperorList, "serverId=? and deleteTime =0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询龙椅抢夺记录
func (dao *emperorDao) GetEmperorRecordsList(serverId int32) (emperorRecords []*emperorentity.EmperorRecordsEntity, err error) {
	err = dao.ds.DB().Order("`robTime` DESC").Find(&emperorRecords, "serverId=? and deleteTime =0", serverId).Error
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
	dao  *emperorDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &emperorDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetEmperorDao() EmperorDao {
	return dao
}
