package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	allianceentity "fgame/fgame/game/alliance/entity"
	"fgame/fgame/game/global"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "alliance"
)

type AllianceDao interface {
	//查询所有仙盟
	GetAllAllianceList() ([]*allianceentity.AllianceEntity, error)
	//获取仙盟成员
	GetAllianceMemberList(allianceId int64) ([]*allianceentity.AllianceMemberEntity, error)
	//获取所有申请
	GetAllianceJoinApplyList(allianceId int64) ([]*allianceentity.AllianceJoinApplyEntity, error)
	//获取霸主
	GetAllianceHegemon() (*allianceentity.AllianceHegemonEntity, error)
	//获取玩家仙盟数据
	GetPlayerAlliance(playerId int64) (*allianceentity.PlayerAllianceEntity, error)
	//获取玩家仙盟仙术数据
	GetPlayerAllianceSkillList(playerId int64) ([]*allianceentity.PlayerAllianceSkillEntity, error)
	//获取玩家仙盟日志数据
	GetPlayerAllianceLogList(allianceId int64, limit int32) ([]*allianceentity.AllianceLogEntity, error)
	//获取所有邀请
	GetAllianceInvitationList(allianceId int64) ([]*allianceentity.AllianceInvitationEntity, error)
	//仓库列表
	GetAllianceDepotItemList(allianceId int64) ([]*allianceentity.AllianceDepotEntity, error)
	//仙盟boss
	GetAllianceBoss(allianceId int64) (*allianceentity.AllianceBossEntity, error)
}

type allianceDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *allianceDao) GetAllAllianceList() (entityList []*allianceentity.AllianceEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime=0 AND serverId = ?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *allianceDao) GetAllianceMemberList(allianceId int64) (entityList []*allianceentity.AllianceMemberEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime=0 and allianceId=?", allianceId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *allianceDao) GetAllianceJoinApplyList(allianceId int64) (entityList []*allianceentity.AllianceJoinApplyEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime=0 and allianceId=?", allianceId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *allianceDao) GetAllianceInvitationList(allianceId int64) (entityList []*allianceentity.AllianceInvitationEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "deleteTime=0 and allianceId=?", allianceId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *allianceDao) GetAllianceDepotItemList(allianceId int64) (entityList []*allianceentity.AllianceDepotEntity, err error) {
	err = dao.ds.DB().Order("`index` ASC").Find(&entityList, "deleteTime=0 and allianceId=?", allianceId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *allianceDao) GetAllianceHegemon() (entity *allianceentity.AllianceHegemonEntity, err error) {
	entity = &allianceentity.AllianceHegemonEntity{}
	err = dao.ds.DB().First(entity, "deleteTime=0 AND serverId=?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *allianceDao) GetPlayerAlliance(playerId int64) (entity *allianceentity.PlayerAllianceEntity, err error) {
	entity = &allianceentity.PlayerAllianceEntity{}
	err = dao.ds.DB().First(entity, "deleteTime=0 and playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *allianceDao) GetPlayerAllianceSkillList(playerId int64) (entityArr []*allianceentity.PlayerAllianceSkillEntity, err error) {
	err = dao.ds.DB().Find(&entityArr, "deleteTime=0 and playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *allianceDao) GetPlayerAllianceLogList(allianceId int64, limit int32) (entityArr []*allianceentity.AllianceLogEntity, err error) {
	err = dao.ds.DB().Limit(limit).Find(&entityArr, "deleteTime=0 and allianceId=?", allianceId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *allianceDao) GetAllianceBoss(allianceId int64) (bossEntity *allianceentity.AllianceBossEntity, err error) {
	bossEntity = &allianceentity.AllianceBossEntity{}
	err = dao.ds.DB().First(bossEntity, "deleteTime=0 and allianceId=?", allianceId).Error
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
	dao  *allianceDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &allianceDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetAllianceDao() AllianceDao {
	return dao
}
