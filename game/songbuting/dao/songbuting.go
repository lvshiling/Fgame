package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/songbuting/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "yuanbao_songbuting"
)

type SongBuTingDao interface {
	//查询玩家身法信息
	GetSongBuTingEntity(playerId int64) (songBuTingEntity *entity.PlayerSongBuTingEntity, err error)
}

type songBuTingDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *songBuTingDao) GetSongBuTingEntity(playerId int64) (songBuTingEntity *entity.PlayerSongBuTingEntity, err error) {
	songBuTingEntity = &entity.PlayerSongBuTingEntity{}
	err = dao.ds.DB().First(songBuTingEntity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *songBuTingDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &songBuTingDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSongBuTingDao() SongBuTingDao {
	return dao
}
