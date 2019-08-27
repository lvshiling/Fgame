package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	goldequipentity "fgame/fgame/game/goldequip/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "gold_equip"
)

type GoldEquipDao interface {
	//加载槽位
	GetGoldEquipSlotList(playerId int64) ([]*goldequipentity.PlayerGoldEquipSlotEntity, error)
	//元神金装日志
	GetPlayerGoldEquipLogEntityList(playerId int64) (logEntityList []*goldequipentity.PlayerGoldEquipLogEntity, err error)
	// 元神金装设置
	GetPlayerGoldEquipSettingEntity(playerId int64) (entity *goldequipentity.PlayerGoldEquipSettingEntity, err error)
	// 获取玩家元神金装数据
	GetPlayerGoldEquipEntity(playerId int64) (entity *goldequipentity.PlayerGoldEquipEntity, err error)
}

type goldEquipDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *goldEquipDao) GetGoldEquipSlotList(playerId int64) (equipmentSlotList []*goldequipentity.PlayerGoldEquipSlotEntity, err error) {
	err = dao.ds.DB().Order("`slotId` ASC").Find(&equipmentSlotList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *goldEquipDao) GetPlayerGoldEquipLogEntityList(playerId int64) (logEntityList []*goldequipentity.PlayerGoldEquipLogEntity, err error) {
	err = dao.ds.DB().Order("updateTime ASC").Find(&logEntityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *goldEquipDao) GetPlayerGoldEquipSettingEntity(playerId int64) (entity *goldequipentity.PlayerGoldEquipSettingEntity, err error) {
	entity = &goldequipentity.PlayerGoldEquipSettingEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *goldEquipDao) GetPlayerGoldEquipEntity(playerId int64) (entity *goldequipentity.PlayerGoldEquipEntity, err error) {
	entity = &goldequipentity.PlayerGoldEquipEntity{}
	err = dao.ds.DB().First(entity, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *goldEquipDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &goldEquipDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetGoldEquipDao() GoldEquipDao {
	return dao
}
