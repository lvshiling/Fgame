package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/global"
	marryentity "fgame/fgame/game/marry/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "marry"
)

type MarryDao interface {
	//查看婚烟信息
	GetMarryList() ([]*marryentity.MarryEntity, error)
	//婚期安排列表
	GetMarryWedList() ([]*marryentity.MarryWedEntity, error)
	//获取求婚婚戒信息
	GetMarryRingList() ([]*marryentity.MarryRingEntity, error)
	//喜帖列表
	GetMarryWedCardList(now int64) ([]*marryentity.MarryWedCardEntity, error)
	//查询玩家婚姻信息
	GetMarryPlayerEntity(playerId int64) (*marryentity.PlayerMarryEntity, error)

	//获取玩家豪气值
	GetHeroismEntity(playerId int64) (*marryentity.PlayerMarryHeroismEntity, error)
	//协议离婚成功请求离婚者已下线
	GetConsentDivorce(playerId int64) (*marryentity.MarryDivorceConsentEntity, error)
	//查询玩家查看过喜帖
	GetViewWedCardList(playerId int64, filterTime int64) ([]*marryentity.PlayerViewWedCardEntity, error)
	//查询推送玩家婚礼按钮记录
	GetPlayerPushWedRecord(playerId int64) (*marryentity.PlayerPushWedRecordEntity, error)
	//获取预定婚期信息
	GetMarryPreWedList() ([]*marryentity.MarryPreWedEntity, error)
	//获取玩家纪念
	GetPlayerMarryJiNianList(playerId int64) ([]*marryentity.PlayerMarryJiNianEntity, error)
	//获取玩家定情
	GetPlayerMarryDingQingList(playerId int64) (*marryentity.PlayerMarryDingQingEntity, error)
	//获取纪念时装套装获取信息
	GetPlayerJiNianSjInfo(playerId int64) (*marryentity.PlayerMarryJiNianSjEntity, error)
}

type marryDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *marryDao) GetMarryPreWedList() (preWedList []*marryentity.MarryPreWedEntity, err error) {
	err = dao.ds.DB().Find(&preWedList, "serverId= ? and deleteTime =0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *marryDao) GetMarryRingList() (ringList []*marryentity.MarryRingEntity, err error) {
	err = dao.ds.DB().Find(&ringList, "serverId= ? and deleteTime =0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *marryDao) GetConsentDivorce(playerId int64) (divorceEntity *marryentity.MarryDivorceConsentEntity, err error) {
	divorceEntity = &marryentity.MarryDivorceConsentEntity{}
	err = dao.ds.DB().First(divorceEntity, "serverId = ? and playerId=? and deleteTime = 0", global.GetGame().GetServerIndex(), playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *marryDao) GetMarryPlayerEntity(playerId int64) (marryEntity *marryentity.PlayerMarryEntity, err error) {
	marryEntity = &marryentity.PlayerMarryEntity{}
	err = dao.ds.DB().First(marryEntity, "playerId=?", playerId).Error //获取结婚的信息
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *marryDao) GetPlayerPushWedRecord(playerId int64) (pushWedRecord *marryentity.PlayerPushWedRecordEntity, err error) {
	pushWedRecord = &marryentity.PlayerPushWedRecordEntity{}
	err = dao.ds.DB().First(pushWedRecord, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//查看婚烟信息
func (dao *marryDao) GetMarryList() (marryList []*marryentity.MarryEntity, err error) {
	// err = dao.ds.DB().Find(&marryList, "serverId=? and deleteTime =0", global.GetGame().GetServerIndex()).Error
	err = dao.ds.DB().Order("id ASC").Find(&marryList, "serverId=?", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//查询玩家查看过喜帖
func (dao *marryDao) GetViewWedCardList(playerId int64, filterTime int64) (viewWedCardList []*marryentity.PlayerViewWedCardEntity, err error) {
	err = dao.ds.DB().Find(&viewWedCardList, "playerId =? and viewTime >=?", playerId, filterTime).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//婚期安排列表
func (dao *marryDao) GetMarryWedList() (marryWedList []*marryentity.MarryWedEntity, err error) {
	err = dao.ds.DB().Order("`period` ASC").Find(&marryWedList, "serverId=? and deleteTime=0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//喜帖列表
func (dao *marryDao) GetMarryWedCardList(now int64) (marryWedCardList []*marryentity.MarryWedCardEntity, err error) {
	err = dao.ds.DB().Order("`createTime` DESC").Find(&marryWedCardList, "serverId=? and outOfTime >? and deleteTime=0", global.GetGame().GetServerIndex(), now).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

//获取玩家豪气值
func (dao *marryDao) GetHeroismEntity(playerId int64) (heroismEntity *marryentity.PlayerMarryHeroismEntity, err error) {
	heroismEntity = &marryentity.PlayerMarryHeroismEntity{}
	err = dao.ds.DB().Find(heroismEntity, "playerId =?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *marryDao) GetPlayerMarryJiNianList(playerId int64) ([]*marryentity.PlayerMarryJiNianEntity, error) {
	jinianList := make([]*marryentity.PlayerMarryJiNianEntity, 0)
	err := dao.ds.DB().Find(&jinianList, "playerId = ?", playerId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return jinianList, nil
}

func (dao *marryDao) GetPlayerMarryDingQingList(playerId int64) (*marryentity.PlayerMarryDingQingEntity, error) {
	dingQingList := &marryentity.PlayerMarryDingQingEntity{}
	err := dao.ds.DB().First(dingQingList, "playerId = ?", playerId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dingQingList, nil
}

func (dao *marryDao) GetPlayerJiNianSjInfo(playerId int64) (*marryentity.PlayerMarryJiNianSjEntity, error) {
	info := &marryentity.PlayerMarryJiNianSjEntity{}
	err := dao.ds.DB().First(info, "playerId = ? and deleteTime=0", playerId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return info, nil
}

var (
	once sync.Once
	dao  *marryDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &marryDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetMarryDao() MarryDao {
	return dao
}
