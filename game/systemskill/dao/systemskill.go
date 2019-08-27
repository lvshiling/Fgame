package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	systemskillentity "fgame/fgame/game/systemskill/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "systemskill"
)

type SystemSkillDao interface {
	//查询玩家系统技能列表
	GetSystemSkillList(playerId int64) ([]*systemskillentity.PlayerSystemSkillEntity, error)
}

type xinFaDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *xinFaDao) GetSystemSkillList(playerId int64) (skillList []*systemskillentity.PlayerSystemSkillEntity, err error) {
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
	dao  *xinFaDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &xinFaDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSystemSkillDao() SystemSkillDao {
	return dao
}
