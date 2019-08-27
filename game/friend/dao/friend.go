package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	friendentity "fgame/fgame/game/friend/entity"
	"fgame/fgame/game/global"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "Friend"
)

type FriendDao interface {
	//获取玩家好友日志
	GetFriendLogList(playerId int64) ([]*friendentity.PlayerFriendLogEntity, error)
	//获取好友系统数据
	GetFriendAllList() ([]*friendentity.FriendEntity, error)
	//获取好友黑名单列表
	GetFriendBlackList(playerId int64) ([]*friendentity.PlayerFriendBlackEntity, error)
	//获取邀请列表
	GetFriendInviteList(playerId int64) ([]*friendentity.PlayerFriendInviteEntity, error)
	//获取赞赏列表
	GetFriendFeedbackList(playerId int64) ([]*friendentity.PlayerFriendFeedbackEntity, error)
	//获取领奖记录
	GetFriendAddRew(playerId int64) (*friendentity.PlayerFriendAddRewEntity, error)
	//获取赞赏记录列表
	GetFriendAdmireList(playerId int64) ([]*friendentity.PlayerFriendAdmireEntity, error)
	//获取全局表白日志
	GetMarryDevelopLogEntityList() (logEntityList []*friendentity.FriendMarryDevelopLogEntity, err error)
	//获取玩家表白日志
	GetMarryDevelopSendLogList(playerId int64) (logEntityList []*friendentity.PlayerFriendMarryDevelopSendLogEntity, err error)
	//获取对玩家表白日志
	GetMarryDevelopRecvLogList(playerId int64) (logEntityList []*friendentity.PlayerFriendMarryDevelopRecvLogEntity, err error)
}

type friendDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *friendDao) GetFriendLogList(playerId int64) (friendLogEntityList []*friendentity.PlayerFriendLogEntity, err error) {
	err = dao.ds.DB().Order("`createTime` ASC").Find(&friendLogEntityList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取好友系统数据
func (dao *friendDao) GetFriendAllList() (allFriendList []*friendentity.FriendEntity, err error) {
	err = dao.ds.DB().Find(&allFriendList, "serverId=?  and deleteTime=0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取好友黑名单列表
func (dao *friendDao) GetFriendBlackList(playerId int64) (blackList []*friendentity.PlayerFriendBlackEntity, err error) {
	err = dao.ds.DB().Find(&blackList, "playerId=?", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取邀请列表
func (dao *friendDao) GetFriendInviteList(playerId int64) (inviteList []*friendentity.PlayerFriendInviteEntity, err error) {
	err = dao.ds.DB().Order("`UpdateTime` ASC").Find(&inviteList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取赞赏列表
func (dao *friendDao) GetFriendFeedbackList(playerId int64) (entityList []*friendentity.PlayerFriendFeedbackEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取赞赏列表
func (dao *friendDao) GetFriendAddRew(playerId int64) (entity *friendentity.PlayerFriendAddRewEntity, err error) {
	entity = &friendentity.PlayerFriendAddRewEntity{}
	err = dao.ds.DB().First(entity, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取赞赏列表
func (dao *friendDao) GetFriendAdmireList(playerId int64) (entityList []*friendentity.PlayerFriendAdmireEntity, err error) {
	err = dao.ds.DB().Find(&entityList, "playerId=? and deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取全局表白日志
func (dao *friendDao) GetMarryDevelopLogEntityList() (logEntityList []*friendentity.FriendMarryDevelopLogEntity, err error) {
	err = dao.ds.DB().Order("updateTime ASC").Find(&logEntityList, "serverId = ? AND deleteTime=0", global.GetGame().GetServerIndex()).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取玩家表白日志
func (dao *friendDao) GetMarryDevelopSendLogList(playerId int64) (logEntityList []*friendentity.PlayerFriendMarryDevelopSendLogEntity, err error) {
	err = dao.ds.DB().Order("updateTime ASC").Find(&logEntityList, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

//获取对玩家表白日志
func (dao *friendDao) GetMarryDevelopRecvLogList(playerId int64) (logEntityList []*friendentity.PlayerFriendMarryDevelopRecvLogEntity, err error) {
	err = dao.ds.DB().Order("updateTime ASC").Find(&logEntityList, "playerId=? AND deleteTime=0", playerId).Error
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
	dao  *friendDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &friendDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetFriendDao() FriendDao {
	return dao
}
