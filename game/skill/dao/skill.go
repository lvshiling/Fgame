package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	skillentity "fgame/fgame/game/skill/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "skill"
)

type SkillDao interface {
	//查询玩家职业技能列表
	GetRoleSkillList(playerId int64) ([]*skillentity.PlayerSkillEntity, error)
	//查询玩家技能cd时间
	GetSkillCdList(playerId int64) ([]*skillentity.PlayerSkillCdEntity, error)
}

type skillDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *skillDao) GetRoleSkillList(playerId int64) (roleSkillList []*skillentity.PlayerSkillEntity, err error) {
	err = dao.ds.DB().Find(&roleSkillList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *skillDao) GetSkillCdList(playerId int64) (skillCdList []*skillentity.PlayerSkillCdEntity, err error) {
	err = dao.ds.DB().Find(&skillCdList, "playerId=? ", playerId).Error
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
	dao  *skillDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &skillDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetSkillDao() SkillDao {
	return dao
}
