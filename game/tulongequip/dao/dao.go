package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	tulongequipentity "fgame/fgame/game/tulongequip/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "tulong_equip"
)

type TuLongEquipDao interface {
	//加载槽位列表
	GetTuLongEquipSlotList(playerId int64) ([]*tulongequipentity.PlayerTuLongEquipSlotEntity, error)
	//获取套装技能列表
	GetTuLongSuitSkillList(playerId int64) ([]*tulongequipentity.PlayerTuLongSuitSkillEntity, error)
	//获取屠龙装数据
	GetTuLongEquipEntity(playerId int64) (*tulongequipentity.PlayerTuLongEquipEntity, error)
}

type tuLongEquipDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *tuLongEquipDao) GetTuLongEquipSlotList(playerId int64) (entityList []*tulongequipentity.PlayerTuLongEquipSlotEntity, err error) {
	err = dao.ds.DB().Order("`slotId` ASC").Find(&entityList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *tuLongEquipDao) GetTuLongSuitSkillList(playerId int64) (entityList []*tulongequipentity.PlayerTuLongSuitSkillEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *tuLongEquipDao) GetTuLongEquipEntity(playerId int64) (entity *tulongequipentity.PlayerTuLongEquipEntity, err error) {
	entity = &tulongequipentity.PlayerTuLongEquipEntity{}
	err = dao.ds.DB().First(&entity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *tuLongEquipDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &tuLongEquipDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetTuLongEquipDao() TuLongEquipDao {
	return dao
}
