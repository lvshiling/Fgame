package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	outlandbossentity "fgame/fgame/game/outlandboss/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "outlandboss"
)

type OutlandBossDao interface {
	//获取玩家外域boss
	GetOutlandBossEntity(playerId int64) (*outlandbossentity.PlayerOutlandBossEntity, error)
	//查询外域boss掉落记录
	GetOutlandBossDropRecordsList(serverId int32) ([]*outlandbossentity.OutlandBossDropRecordsEntity, error)
}

type outlandbossDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *outlandbossDao) GetOutlandBossEntity(playerId int64) (outlandbossEntity *outlandbossentity.PlayerOutlandBossEntity, err error) {
	outlandbossEntity = &outlandbossentity.PlayerOutlandBossEntity{}
	err = dao.ds.DB().First(outlandbossEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查询外域boss掉落记录
func (dao *outlandbossDao) GetOutlandBossDropRecordsList(serverId int32) (records []*outlandbossentity.OutlandBossDropRecordsEntity, err error) {
	err = dao.ds.DB().Order("`dropTime` ASC").Find(&records, "serverId=? AND deleteTime=0", serverId).Error
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
	dao  *outlandbossDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &outlandbossDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetOutlandBossDao() OutlandBossDao {
	return dao
}
