package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	arenaentity "fgame/fgame/game/arena/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "arena"
)

type ArenaDao interface {
	//获取竞技场
	GetArena(playerId int64) (*arenaentity.PlayerArenaEntity, error)
	//查询3v3排行榜
	GetArenaRankList(serverId int32) ([]*arenaentity.ArenaRankEntity, error)
	//排行榜刷新时间戳
	GetArenaRankTimeEntity(serverId int32) (*arenaentity.ArenaRankTimeEntity, error)
}

type arenaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *arenaDao) GetArena(playerId int64) (arenaEntity *arenaentity.PlayerArenaEntity, err error) {
	arenaEntity = &arenaentity.PlayerArenaEntity{}
	err = dao.ds.DB().First(arenaEntity, "playerId = ? AND deleteTime = 0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}

		return nil, nil
	}
	return
}

func (dao *arenaDao) GetArenaRankList(serverId int32) (arenaRankList []*arenaentity.ArenaRankEntity, err error) {
	err = dao.ds.DB().Find(&arenaRankList, "serverId=? and  deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *arenaDao) GetArenaRankTimeEntity(serverId int32) (arenaRankTimeEntity *arenaentity.ArenaRankTimeEntity, err error) {
	arenaRankTimeEntity = &arenaentity.ArenaRankTimeEntity{}
	err = dao.ds.DB().First(arenaRankTimeEntity, "serverId=? and  deleteTime=0", serverId).Error
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
	dao  *arenaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) error {
	once.Do(func() {
		dao = &arenaDao{
			ds: ds,
			rs: rs,
		}
	})

	return nil
}

func GetArenaDao() ArenaDao {
	return dao
}
