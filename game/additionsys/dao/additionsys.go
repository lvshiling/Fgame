package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	additionsysentity "fgame/fgame/game/additionsys/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "addition_sys"
)

type AdditionSysDao interface {
	//加载槽位
	GetAdditionSysSlotList(playerId int64) ([]*additionsysentity.PlayerAdditionSysSlotEntity, error)
	//加载升级数据
	GetAdditionSysLevelList(playerId int64) ([]*additionsysentity.PlayerAdditionSysLevelEntity, error)
	//加载觉醒数据
	GetAdditionSysAwakeList(playerId int64) ([]*additionsysentity.PlayerAdditionSysAwakeEntity, error)
	//加载升级数据
	GetAdditionSysTongLingList(playerId int64) ([]*additionsysentity.PlayerAdditionSysTongLingEntity, error)
	//加载玩家圣痕
	GetPlayerShengHenEntity(playerId int64) (*additionsysentity.PlayerShengHenEntity, error)
	//加载玩家灵珠
	GetAdditionSysLingZhuList(playerId int64) ([]*additionsysentity.PlayerAdditionSysLingZhuEntity, error)
}

type additionSysDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *additionSysDao) GetAdditionSysSlotList(playerId int64) (equipmentSlotList []*additionsysentity.PlayerAdditionSysSlotEntity, err error) {
	err = dao.ds.DB().Order("`slotId` ASC").Find(&equipmentSlotList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *additionSysDao) GetAdditionSysLevelList(playerId int64) (levelList []*additionsysentity.PlayerAdditionSysLevelEntity, err error) {
	err = dao.ds.DB().Find(&levelList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *additionSysDao) GetAdditionSysAwakeList(playerId int64) (awakeList []*additionsysentity.PlayerAdditionSysAwakeEntity, err error) {
	err = dao.ds.DB().Find(&awakeList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *additionSysDao) GetAdditionSysTongLingList(playerId int64) (tongLingList []*additionsysentity.PlayerAdditionSysTongLingEntity, err error) {
	err = dao.ds.DB().Find(&tongLingList, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *additionSysDao) GetPlayerShengHenEntity(playerId int64) (shengHenEntity *additionsysentity.PlayerShengHenEntity, err error) {
	shengHenEntity = &additionsysentity.PlayerShengHenEntity{}
	err = dao.ds.DB().First(&shengHenEntity, "playerId=? ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *additionSysDao) GetAdditionSysLingZhuList(playerId int64) (lingZhuList []*additionsysentity.PlayerAdditionSysLingZhuEntity, err error) {
	err = dao.ds.DB().Find(&lingZhuList, "playerId=? ", playerId).Error
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
	dao  *additionSysDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &additionSysDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetAdditionSysDao() AdditionSysDao {
	return dao
}
