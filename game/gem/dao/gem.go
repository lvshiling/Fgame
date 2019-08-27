package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	gementity "fgame/fgame/game/gem/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "Gem"
)

type GemDao interface {
	//查询玩家矿山信息
	GetMineEntity(playerId int64) (*gementity.PlayerMiningEntity, error)
	//查询玩家赌石列表
	GetGambleList(playerId int64) ([]*gementity.PlayerGambleEntity, error)
}

type gemDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *gemDao) GetMineEntity(playerId int64) (miningEntity *gementity.PlayerMiningEntity, err error) {
	miningEntity = &gementity.PlayerMiningEntity{}
	err = dao.ds.DB().First(miningEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *gemDao) GetGambleList(playerId int64) (gambleList []*gementity.PlayerGambleEntity, err error) {
	err = dao.ds.DB().Find(&gambleList, "playerId=? ", playerId).Error
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
	dao  *gemDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &gemDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetGemDao() GemDao {
	return dao
}
