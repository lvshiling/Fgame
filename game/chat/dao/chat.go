package dao

import (
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	chatentity "fgame/fgame/game/chat/entity"
	"sync"

	"github.com/jinzhu/gorm"
)

const (
	dbName = "chat"
)

type ChatDao interface {
	//获取聊天设置
	GetChatSetting(serverId int32) (*chatentity.ChatSettingEntity, error)
	//获取玩家聊天
	GetPlayerChatEntity(playerId int64) (*chatentity.PlayerChatEntity, error)
}

type chatDao struct {
	ds coredb.DBService
	rs coreredis.RedisService
}

func (dao *chatDao) GetChatSetting(serverId int32) (chatSettingEntity *chatentity.ChatSettingEntity, err error) {
	chatSettingEntity = &chatentity.ChatSettingEntity{}
	err = dao.ds.DB().First(chatSettingEntity, "serverId=?", serverId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return nil, nil
	}
	return

}

func (dao *chatDao) GetPlayerChatEntity(playerId int64) (entity *chatentity.PlayerChatEntity, err error) {
	entity = &chatentity.PlayerChatEntity{}
	err = dao.ds.DB().First(entity, "deleteTime=0 AND playerId=?", playerId).Error
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
	dao  *chatDao
)

func Init(ds coredb.DBService, rs coreredis.RedisService) (err error) {
	once.Do(func() {
		dao = &chatDao{
			ds: ds,
			rs: rs,
		}
	})
	return nil
}

func GetChatDao() ChatDao {
	return dao
}
