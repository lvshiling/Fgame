package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	wushuangweaponentity "fgame/fgame/game/wushuangweapon/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbSlotName     = "wushuangweapon_slot"
	dbSettingsName = "wushuang_settings"
	dbBuchangName  = "wushuang_buchang"
)

type WushuangWeaponDao interface {
	GetAllWushuangWeaponSlotEntity(playerId int64) (wushuangEntityList []*wushuangweaponentity.PlayerWushuangWeaponSlotEntity, err error)
	GetAllWushuangSettings(playerId int64) (wushuangSettingsEntityList []*wushuangweaponentity.PlayerWushuangSettingsEntity, err error)
	GetAllWushuangBuchang(playerId int64) (wushuangSettingsEntityList *wushuangweaponentity.PlayerWushuangBuchangEntity, err error)
}

type wushuangWeaponDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *wushuangWeaponDao) GetAllWushuangWeaponSlotEntity(playerId int64) (wushuangEntityList []*wushuangweaponentity.PlayerWushuangWeaponSlotEntity, err error) {
	err = dao.ds.DB().Order("`slotId` ASC").Find(&wushuangEntityList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbSlotName, err)
		}
		return nil, nil
	}
	return
}

func (dao *wushuangWeaponDao) GetAllWushuangBuchang(playerId int64) (wushuangBuchangEntity *wushuangweaponentity.PlayerWushuangBuchangEntity, err error) {
	wushuangBuchangEntity = &wushuangweaponentity.PlayerWushuangBuchangEntity{}
	err = dao.ds.DB().First(wushuangBuchangEntity, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbBuchangName, err)
		}
		return nil, nil
	}
	return
}

func (dao *wushuangWeaponDao) GetAllWushuangSettings(playerId int64) (wushuangSettingsEntityList []*wushuangweaponentity.PlayerWushuangSettingsEntity, err error) {
	err = dao.ds.DB().Order("`itemId` ASC").Find(&wushuangSettingsEntityList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbSettingsName, err)
		}
		return nil, nil
	}
	return
}

var (
	once sync.Once
	dao  *wushuangWeaponDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &wushuangWeaponDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetWushuangWeaponDao() WushuangWeaponDao {
	return dao
}
