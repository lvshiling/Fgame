package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/game/jieyi/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "jieyi"
)

type JieYiDao interface {
	// 查询玩家结义信息
	GetPlayerJieYiEntity(playerId int64) (*entity.PlayerJieYiEntity, error)
	// 查询结义信息
	GetJieYiListEntity(serverId int32) ([]*entity.JieYiEntity, error)
	// 查询结义成员信息
	GetJieYiMemberListEntity(serverId int32) ([]*entity.JieYiMemberEntity, error)
	// 查询结义留言信息
	GetJieYiLeaveWordListEntity(serverId int32) ([]*entity.JieYiLeaveWordEntity, error)
	// 查询结义邀请信息
	GetJieYiInviteListEntity(serverId int32) ([]*entity.JieYiInviteEntity, error)
}

type jieYiDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *jieYiDao) GetPlayerJieYiEntity(playerId int64) (jieyiEntity *entity.PlayerJieYiEntity, err error) {
	jieyiEntity = &entity.PlayerJieYiEntity{}
	err = dao.ds.DB().First(jieyiEntity, "playerId=? AND deleteTime=0", playerId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return
}

func (dao *jieYiDao) GetJieYiListEntity(serverId int32) (jieyiListEntity []*entity.JieYiEntity, err error) {
	err = dao.ds.DB().Find(&jieyiListEntity, "deleteTime=0 AND serverId = ?", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *jieYiDao) GetJieYiMemberListEntity(serverId int32) (jieyiListEntity []*entity.JieYiMemberEntity, err error) {
	err = dao.ds.DB().Find(&jieyiListEntity, "deleteTime=0 AND serverId = ?", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *jieYiDao) GetJieYiLeaveWordListEntity(serverId int32) (jieyiListEntity []*entity.JieYiLeaveWordEntity, err error) {
	err = dao.ds.DB().Find(&jieyiListEntity, "deleteTime=0 AND serverId = ?", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	return
}

func (dao *jieYiDao) GetJieYiInviteListEntity(serverId int32) (jieyiListEntity []*entity.JieYiInviteEntity, err error) {
	err = dao.ds.DB().Find(&jieyiListEntity, "deleteTime=0 AND serverId = ?", serverId).Error
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
	dao  *jieYiDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &jieYiDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetJieYiDao() JieYiDao {
	return dao
}
