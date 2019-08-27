package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	inventoryentity "fgame/fgame/game/inventory/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "inventory"
)

type InventoryDao interface {
	//加载背包
	GetInventoryEntity(playerId int64) (*inventoryentity.PlayerInventoryEntity, error)
	//加载物品
	GetItemList(playerId int64) ([]*inventoryentity.PlayerItemEntity, error)
	//加载槽位
	GetEquipmentSlotList(playerId int64) ([]*inventoryentity.PlayerEquipmentSlotEntity, error)
	//加载仓库物品
	GetDepotItemList(playerId int64) ([]*inventoryentity.PlayerItemEntity, error)
	//加载秘宝仓库物品
	GetMiBaoDepotItemList(playerId int64) ([]*inventoryentity.PlayerItemEntity, error)
	//加载材料仓库物品
	GetMaterialDepotItemList(playerId int64) ([]*inventoryentity.PlayerItemEntity, error)
	//加载使用记录
	GetItemUseList(playerId int64) ([]*inventoryentity.PlayerItemUseEntity, error)
}

type inventoryDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *inventoryDao) GetInventoryEntity(playerId int64) (inventoryEntity *inventoryentity.PlayerInventoryEntity, err error) {
	inventoryEntity = &inventoryentity.PlayerInventoryEntity{}
	err = dao.ds.DB().First(inventoryEntity, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *inventoryDao) GetItemList(playerId int64) (itemList []*inventoryentity.PlayerItemEntity, err error) {
	err = dao.ds.DB().Order("`index` ASC").Find(&itemList, "playerId=? AND isDepot=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *inventoryDao) GetEquipmentSlotList(playerId int64) (equipmentSlotList []*inventoryentity.PlayerEquipmentSlotEntity, err error) {
	err = dao.ds.DB().Order("`slotId` ASC").Find(&equipmentSlotList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *inventoryDao) GetDepotItemList(playerId int64) (itemList []*inventoryentity.PlayerItemEntity, err error) {
	err = dao.ds.DB().Order("`index` ASC").Find(&itemList, "playerId=? AND isDepot=1", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *inventoryDao) GetMiBaoDepotItemList(playerId int64) (itemList []*inventoryentity.PlayerItemEntity, err error) {
	err = dao.ds.DB().Order("`index` ASC").Find(&itemList, "playerId=? AND isDepot=2", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *inventoryDao) GetMaterialDepotItemList(playerId int64) (itemList []*inventoryentity.PlayerItemEntity, err error) {
	err = dao.ds.DB().Order("`index` ASC").Find(&itemList, "playerId=? AND isDepot=3", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *inventoryDao) GetItemUseList(playerId int64) (itemUseList []*inventoryentity.PlayerItemUseEntity, err error) {
	err = dao.ds.DB().Find(&itemUseList, "playerId=?", playerId).Error
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
	dao  *inventoryDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &inventoryDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetInventoryDao() InventoryDao {
	return dao
}
