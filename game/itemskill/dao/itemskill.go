package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	itemskillentity "fgame/fgame/game/itemskill/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "itemskill"
)

type ItemSkillDao interface {
	//查询玩家物品技能列表
	GetItemSkillList(playerId int64) ([]*itemskillentity.PlayerItemSkillEntity, error)
}

type itemSkillDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *itemSkillDao) GetItemSkillList(playerId int64) (skillList []*itemskillentity.PlayerItemSkillEntity, err error) {
	err = dao.ds.DB().Find(&skillList, "playerId=? and deleteTime=0", playerId).Error
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
	dao  *itemSkillDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &itemSkillDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetItemSkillDao() ItemSkillDao {
	return dao
}
