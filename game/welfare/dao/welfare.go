package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/global"
	welfareentity "fgame/fgame/game/welfare/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "welfare"
)

type WelfareDao interface {
	//开服活动
	GetOpenActivityList(playerId int64) (openActivityEntityList []*welfareentity.PlayerOpenActivityEntity, err error)
	//活动充值信息
	GetOpenActivityChargeList(playerId int64) (chrgeEntityList []*welfareentity.PlayerOpenActivityChargeEntity, err error)
	//活动消费信息
	GetOpenActivityCostList(playerId int64) (costEntityList []*welfareentity.PlayerOpenActivityCostEntity, err error)
	//首充信息
	GetPlayerFirstCharge(playerId int64) (firstChargeEntity *welfareentity.PlayerFirstChargeEntity, err error)
	//排行榜邮件奖励记录
	GetOpenActivityEmailRecordList() (recordEntityList []*welfareentity.OpenActivityEmailRecordEntity, err error)
	//活动抽奖信息
	GetActivityNumRecordList(playerId int64) (timesEntityList []*welfareentity.PlayerActivityNumRecordEntity, err error)
	//活动增长数据
	GetActivityAddNumList(playerId int64) (addNumEntityList []*welfareentity.PlayerActivityAddNumEntity, err error)
	//全服奖励次数限制
	GetOpenActivityRewardsLimitList() (entityList []*welfareentity.OpenActivityRewardsLimitEntity, err error)
	//全服折扣次数限制
	GetOpenActivityDiscountLimitList() (entityList []*welfareentity.OpenActivityDiscountLimitEntity, err error)
	//玩家活动开启邮件数据
	GetPlayerActivityOpenMailRecordList(playerId int64) (entityList []*welfareentity.PlayerActivityOpenMailEntity, err error)
	//活动开启邮件通知记录
	GetOpenActivityStartEmailList() (recordEntityList []*welfareentity.OpenActivityStartEmailEntity, err error)
	//BOSS首杀记录
	GetOpenActivityBossKilledIdList(serverId int32) (entityList []*welfareentity.OpenActivityBossKillEntity, err error)
	//城战助威
	GetOpenActivityAllianceCheerList(serverId int32) (entityList []*welfareentity.OpenActivityAllianceCheerEntity, err error)
	//循环活动
	GetOpenActivityXunHuan(serverId int32) (entity *welfareentity.OpenActivityXunHuanEntity, err error)
}

type welfareDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *welfareDao) GetOpenActivityList(playerId int64) (openActivityEntityList []*welfareentity.PlayerOpenActivityEntity, err error) {
	err = dao.ds.DB().Find(&openActivityEntityList, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}
func (dao *welfareDao) GetOpenActivityChargeList(playerId int64) (chrgeEntityList []*welfareentity.PlayerOpenActivityChargeEntity, err error) {
	err = dao.ds.DB().Find(&chrgeEntityList, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityCostList(playerId int64) (costEntityList []*welfareentity.PlayerOpenActivityCostEntity, err error) {
	err = dao.ds.DB().Find(&costEntityList, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetActivityNumRecordList(playerId int64) (timesEntityList []*welfareentity.PlayerActivityNumRecordEntity, err error) {
	err = dao.ds.DB().Find(&timesEntityList, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetActivityAddNumList(playerId int64) (addNumEntityList []*welfareentity.PlayerActivityAddNumEntity, err error) {
	err = dao.ds.DB().Find(&addNumEntityList, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetPlayerActivityOpenMailRecordList(playerId int64) (entityList []*welfareentity.PlayerActivityOpenMailEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetPlayerFirstCharge(playerId int64) (firstChargeEntity *welfareentity.PlayerFirstChargeEntity, err error) {
	firstChargeEntity = &welfareentity.PlayerFirstChargeEntity{}
	err = dao.ds.DB().First(firstChargeEntity, "playerId=? and deleteTime=0 ", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityEmailRecordList() (recordEntityList []*welfareentity.OpenActivityEmailRecordEntity, err error) {
	err = dao.ds.DB().Find(&recordEntityList, "deleteTime = 0 AND serverId=?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityStartEmailList() (recordEntityList []*welfareentity.OpenActivityStartEmailEntity, err error) {
	err = dao.ds.DB().Find(&recordEntityList, "deleteTime = 0 AND serverId=?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityRewardsLimitList() (entityList []*welfareentity.OpenActivityRewardsLimitEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime = 0 AND serverId=?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityDiscountLimitList() (entityList []*welfareentity.OpenActivityDiscountLimitEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime = 0 AND serverId=?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityBossKilledIdList(serverId int32) (entityList []*welfareentity.OpenActivityBossKillEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityAllianceCheerList(serverId int32) (entityList []*welfareentity.OpenActivityAllianceCheerEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "serverId=? AND deleteTime=0", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *welfareDao) GetOpenActivityXunHuan(serverId int32) (entity *welfareentity.OpenActivityXunHuanEntity, err error) {
	entity = &welfareentity.OpenActivityXunHuanEntity{}
	err = dao.ds.DB().First(&entity, "serverId=? AND deleteTime=0", serverId).Error
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
	dao  *welfareDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &welfareDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetWelfareDao() WelfareDao {
	return dao
}
