package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	soulentity "fgame/fgame/game/soul/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "soul"
)

type SoulDao interface {
	//查询玩家帝魂镶嵌
	GetSoulEmbedEntity(playerId int64) (*soulentity.PlayerSoulEmbedEntity, error)
	//查询玩家帝魂列表
	GetSoulList(playerId int64) ([]*soulentity.PlayerSoulEntity, error)
}

type soulDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *soulDao) GetSoulEmbedEntity(playerId int64) (soulEntity *soulentity.PlayerSoulEmbedEntity, err error) {
	soulEntity = &soulentity.PlayerSoulEmbedEntity{}
	err = dao.ds.DB().First(soulEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *soulDao) GetSoulList(playerId int64) (soulList []*soulentity.PlayerSoulEntity, err error) {
	err = dao.ds.DB().Find(&soulList, "playerId=? ", playerId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

var (
	once sync.Once
	dao  *soulDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &soulDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSoulDao() SoulDao {
	return dao
}
