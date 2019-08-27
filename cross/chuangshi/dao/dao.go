package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	chuangshientity "fgame/fgame/cross/chuangshi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "chuangshi"
)

type ChuangShiDao interface {
	// 所有阵营
	GetCampList(platform int32, serverId int32) (eList []*chuangshientity.ChuangShiCampEntity, err error)
	//所有城市
	GetCityList(platform int32, serverId int32, camp int32) (eList []*chuangshientity.ChuangShiCityEntity, err error)
	//所有城市建设
	GetCityJianSheList(platform int32, serverId int32, cityId int64) (eList []*chuangshientity.ChuangShiCityJianSheEntity, err error)
	//神王报名列表
	GetShenWangSignUpList(platform int32, serverId int32, camp int32) (eList []*chuangshientity.ChuangShiShenWangSignUpEntity, err error)
	//神王投票列表
	GetShenWangVoteList(platform int32, serverId int32, camp int32) (eList []*chuangshientity.ChuangShiShenWangVoteEntity, err error)
	//神王投票记录列表
	GetShenWangVoteRecordList(platform int32, serverId int32) (eList []*chuangshientity.ChuangShiShenWangVoteRecordEntity, err error)
	//阵营成员列表
	GetChuangshiMemberList(platform int32, serverId int32, camp int32) (eList []*chuangshientity.ChuangShiMemberEntity, err error)
}

type chuangShiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (d *chuangShiDao) GetCampList(platform int32, serverId int32) (eList []*chuangshientity.ChuangShiCampEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiCampEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and  deleteTime=0", platform, serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (d *chuangShiDao) GetCityList(platform int32, serverId int32, campType int32) (eList []*chuangshientity.ChuangShiCityEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiCityEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and campType=? and  deleteTime=0", platform, serverId, campType).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (d *chuangShiDao) GetCityJianSheList(platform int32, serverId int32, cityId int64) (eList []*chuangshientity.ChuangShiCityJianSheEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiCityJianSheEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and cityId=? and  deleteTime=0", platform, serverId, cityId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (d *chuangShiDao) GetShenWangSignUpList(platform int32, serverId int32, campType int32) (eList []*chuangshientity.ChuangShiShenWangSignUpEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiShenWangSignUpEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and campType=? and  deleteTime=0", platform, serverId, campType).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (d *chuangShiDao) GetShenWangVoteList(platform int32, serverId int32, campType int32) (eList []*chuangshientity.ChuangShiShenWangVoteEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiShenWangVoteEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and campType=? and  deleteTime=0", platform, serverId, campType).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (d *chuangShiDao) GetShenWangVoteRecordList(platform int32, serverId int32) (eList []*chuangshientity.ChuangShiShenWangVoteRecordEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiShenWangVoteRecordEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and  deleteTime=0", platform, serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (d *chuangShiDao) GetChuangshiMemberList(platform int32, serverId int32, campType int32) (eList []*chuangshientity.ChuangShiMemberEntity, err error) {
	eList = make([]*chuangshientity.ChuangShiMemberEntity, 0, 8)
	err = dao.ds.DB().Find(&eList, "platform=? and serverId=? and campType=? and  deleteTime=0", platform, serverId, campType).Error
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
	dao  *chuangShiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &chuangShiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetChuangShiDao() ChuangShiDao {
	return dao
}
