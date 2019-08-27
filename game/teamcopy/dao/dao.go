package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	teamcopyentity "fgame/fgame/game/teamcopy/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "team_copy"
)

type TeamCopyDao interface {
	//获取玩家组队副本
	GetTeamCopyList(playerId int64) (teamCopyList []*teamcopyentity.PlayerTeamCopyEntity, err error)
}

type teamCopyDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *teamCopyDao) GetTeamCopyList(playerId int64) (teamCopyList []*teamcopyentity.PlayerTeamCopyEntity, err error) {
	err = dao.ds.DB().Find(&teamCopyList, "playerId=? ", playerId).Error
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
	dao  *teamCopyDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &teamCopyDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTeamCopyDao() TeamCopyDao {
	return dao
}
