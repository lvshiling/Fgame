package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	soulruinsentity "fgame/fgame/game/soulruins/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "soulRuins"
)

type SoulRuinsDao interface {
	//查询玩家帝陵遗迹挑战次数
	GetSoulRuinsNumEntity(playerId int64) (*soulruinsentity.PlayerSoulRuinsNumEntity, error)
	//查询玩家帝陵遗迹列表
	GetSoulRuinsList(playerId int64) ([]*soulruinsentity.PlayerSoulRuinsEntity, error)
	//查询玩家帝陵遗迹章节奖励列表
	GetSoulRuinsRewChapterList(playerId int64) ([]*soulruinsentity.PlayerSoulRuinsRewChapterEntity, error)
}

type soulRuinsDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *soulRuinsDao) GetSoulRuinsNumEntity(playerId int64) (soulRuinsEntity *soulruinsentity.PlayerSoulRuinsNumEntity, err error) {
	soulRuinsEntity = &soulruinsentity.PlayerSoulRuinsNumEntity{}
	err = dao.ds.DB().First(soulRuinsEntity, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *soulRuinsDao) GetSoulRuinsList(playerId int64) (soulRuinsList []*soulruinsentity.PlayerSoulRuinsEntity, err error) {
	err = dao.ds.DB().Find(&soulRuinsList, "playerId=? and deleteTime =0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *soulRuinsDao) GetSoulRuinsRewChapterList(playerId int64) (soulRuinsRewChapterList []*soulruinsentity.PlayerSoulRuinsRewChapterEntity, err error) {
	err = dao.ds.DB().Find(&soulRuinsRewChapterList, "playerId=? ", playerId).Error
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
	dao  *soulRuinsDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &soulRuinsDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSoulRuinsDao() SoulRuinsDao {
	return dao
}
