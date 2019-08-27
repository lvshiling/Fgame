package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	playerentity "fgame/fgame/game/player/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "player"
)

type PlayerDao interface {
	//查询玩家根据名字
	QueryByName(serverId int32, name string) (playerEntity *playerentity.PlayerEntity, err error)
	//查询根据玩家id
	QueryById(id int64) (playerEntity *playerentity.PlayerEntity, err error)
	//查询根据用户id
	QueryByUserId(userId int64, serverId int32) (playerEntity *playerentity.PlayerEntity, err error)
}

type playerDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *playerDao) QueryByName(serverId int32, name string) (playerEntity *playerentity.PlayerEntity, err error) {
	playerEntity = &playerentity.PlayerEntity{}
	err = dao.ds.DB().First(playerEntity, "originServerId=? and name=? and deleteTime=0", serverId, name).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return playerEntity, nil
}

func (dao *playerDao) QueryById(id int64) (playerEntity *playerentity.PlayerEntity, err error) {
	playerEntity = &playerentity.PlayerEntity{}
	err = dao.ds.DB().First(playerEntity, "id=? and deleteTime=0", id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return playerEntity, nil
}

func (dao *playerDao) QueryByUserId(userId int64, serverId int32) (playerEntity *playerentity.PlayerEntity, err error) {
	playerEntity = &playerentity.PlayerEntity{}
	err = dao.ds.DB().First(playerEntity, "userId=? and originServerId=? and deleteTime=0", userId, serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return playerEntity, nil
}


var (
	once sync.Once
	dao  *playerDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &playerDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetPlayerDao() PlayerDao {
	return dao
}
