package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	livenessentity "fgame/fgame/game/liveness/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "liveness"
)

type LivenessDao interface {
	//查询玩家活跃度信息
	GetLivenessEntity(playerId int64) (*livenessentity.PlayerLivenessEntity, error)
	//查询玩家活跃度任务
	GetLivenessQuestList(playerId int64) (questList []*livenessentity.PlayerLivenessQuestEntity, err error)
}

type livenessDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

//查询玩家活跃度信息
func (dao *livenessDao) GetLivenessEntity(playerId int64) (livenessEntity *livenessentity.PlayerLivenessEntity, err error) {
	livenessEntity = &livenessentity.PlayerLivenessEntity{}
	err = dao.ds.DB().First(livenessEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询玩家活跃度任务信息
func (dao *livenessDao) GetLivenessQuestList(playerId int64) (questList []*livenessentity.PlayerLivenessQuestEntity, err error) {
	err = dao.ds.DB().Find(&questList, "playerId=? ", playerId).Error
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
	dao  *livenessDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &livenessDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetLivenessDao() LivenessDao {
	return dao
}
