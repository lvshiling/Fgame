package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type PlatformChatSetEntity struct {
	Id               int64 `gorm:"primary_key;column:id"`
	PlatformId       int32 `gorm:"column:platformId"`
	MinVip           int32 `gorm:"column:minVip"`
	MinPlayerlevel   int32 `gorm:"column:minPlayerlevel"`
	WorldVip         int32 `gorm:"column:worldVip"`
	WorldPlayerLevel int32 `gorm:"column:worldPlayerLevel"`
	PChatVip         int32 `gorm:"column:pChatVip"`
	PChatPlayerLevel int32 `gorm:"column:pChatPlayerLevel"`
	GuildVip         int32 `gorm:"column:guildVip"`
	GuildPlayerLevel int32 `gorm:"column:guildPlayerLevel"`
	TeamVip          int32 `gorm:"column:teamVip"`
	TeamPlayerLevel  int32 `gorm:"column:teamPlayerLevel"`
	UpdateTime       int64 `gorm:"column:updateTime"`
	CreateTime       int64 `gorm:"column:createTime"`
	DeleteTime       int64 `gorm:"column:deleteTime"`
}

func (e *PlatformChatSetEntity) TableName() string {
	return "t_platform_chatset"
}

type PlatformChatsetStore interface {
	//获取平台配置
	GetPlatformChatset(platform int32) (*PlatformChatSetEntity, error)
	//获取所有平台配置
	GetAllPlatformChatset() ([]*PlatformChatSetEntity, error)
}

var (
	platformChatSetDbName = "platformChatSet"
)

type platformChatsetStore struct {
	db coredb.DBService
}

func (s *platformChatsetStore) GetAllPlatformChatset() (eList []*PlatformChatSetEntity, err error) {
	eList = make([]*PlatformChatSetEntity, 0, 8)
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(platformChatSetDbName, err)
		}
		return nil, nil
	}
	return
}

func (s *platformChatsetStore) GetPlatformChatset(platform int32) (e *PlatformChatSetEntity, err error) {
	e = &PlatformChatSetEntity{}
	err = s.db.DB().First(e, "deleteTime=0 and platformId=?", platform).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(platformChatSetDbName, err)
		}
		return nil, nil
	}
	return
}

func NewPlatformChatsetStore(db coredb.DBService) PlatformChatsetStore {
	s := &platformChatsetStore{
		db: db,
	}
	return s
}
