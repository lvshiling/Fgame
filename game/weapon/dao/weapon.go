package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	weaponentity "fgame/fgame/game/weapon/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "weapon"
)

type WeaponDao interface {
	//查询玩家兵魂信息
	GetWeaponInfoEntity(playerId int64) (*weaponentity.PlayerWeaponInfoEntity, error)
	//查询玩家兵魂列表
	GetWeaponList(playerId int64) ([]*weaponentity.PlayerWeaponEntity, error)
}

type weaponDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *weaponDao) GetWeaponInfoEntity(playerId int64) (weaponInfo *weaponentity.PlayerWeaponInfoEntity, err error) {
	weaponInfo = &weaponentity.PlayerWeaponInfoEntity{}
	err = dao.ds.DB().First(weaponInfo, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *weaponDao) GetWeaponList(playerId int64) (weaponList []*weaponentity.PlayerWeaponEntity, err error) {
	err = dao.ds.DB().Find(&weaponList, "playerId=? ", playerId).Error
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
	dao  *weaponDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &weaponDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetWeaponDao() WeaponDao {
	return dao
}
