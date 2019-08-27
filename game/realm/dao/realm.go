package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/global"
	realmentity "fgame/fgame/game/realm/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "realm"
)

const (
	realmRankSql = `SELECT 
				   		A.playerId,
				   		A.playerName,
				   		A.level,
				   		A.usedTime
				   FROM  
						t_player_tianjieta A
				   INNER JOIN
				   		t_player B ON serverId=? AND B.id=A.playerId  
				   Where 
				   		A.level > 0  
				   ORDER BY 
				   		level DESC, 
				   		usedTime ASC
				   LIMIT ?`
)

type RealmDao interface {
	//查询玩家时装穿戴
	GetTianJieTaEntity(playerId int64) (*realmentity.PlayerTianJieTaEntity, error)
	//玩家天劫塔排行
	GetRankRealmList(num int32) (realmList []*realmentity.PlayerTianJieTaEntity, err error)
}

type realmDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *realmDao) GetTianJieTaEntity(playerId int64) (tianJieTaEntity *realmentity.PlayerTianJieTaEntity, err error) {
	tianJieTaEntity = &realmentity.PlayerTianJieTaEntity{}
	err = dao.ds.DB().First(tianJieTaEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//玩家天劫塔排行
func (dao *realmDao) GetRankRealmList(num int32) (realmList []*realmentity.PlayerTianJieTaEntity, err error) {
	err = dao.ds.DB().Raw(realmRankSql, global.GetGame().GetServerIndex(), num).Scan(&realmList).Error
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
	dao  *realmDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &realmDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetRealmDao() RealmDao {
	return dao
}
