package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"

	sceneentity "fgame/fgame/game/scene/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "scene"
)

type SceneDao interface {
	//查询玩家场景
	GetSceneEntity(playerId int64) (*sceneentity.PlayerSceneEntity, error)
	//查询玩家点星
	// GetDingShiBossEntityList(serverId int32) (dingShiBossList []*sceneentity.DingShiBossEntity, err error)
}

type sceneDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *sceneDao) GetSceneEntity(playerId int64) (sceneEntity *sceneentity.PlayerSceneEntity, err error) {
	sceneEntity = &sceneentity.PlayerSceneEntity{}
	err = dao.ds.DB().First(sceneEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

// func (dao *sceneDao) GetDingShiBossEntityList(serverId int32) (dingShiBossList []*sceneentity.DingShiBossEntity, err error) {
// 	dingShiBossList = make([]*sceneentity.DingShiBossEntity, 0, 8)
// 	err = dao.ds.DB().Find(&dingShiBossList, "serverId=? AND deleteTime=0", serverId).Error
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return nil, coredb.NewDBError(dbName, err)
// 		}
// 		return nil, nil
// 	}
// 	return

// }

var (
	once sync.Once
	dao  *sceneDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &sceneDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSceneDao() SceneDao {
	return dao
}
