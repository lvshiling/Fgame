package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type ServerConfigEntity struct {
	Id            int64  `gorm:"primary_key;column:id"`
	TradeServerIp string `gorm:"column:tradeServerIp"`

	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *ServerConfigEntity) TableName() string {
	return "t_platform_server_config"
}

type ServerConfigStore interface {
	//获取客户端版本
	GetServerConfig() (*ServerConfigEntity, error)
}

var (
	serverConfigDbName = "server_config"
)

type serverConfigStore struct {
	db coredb.DBService
}

func (s *serverConfigStore) GetServerConfig() (e *ServerConfigEntity, err error) {
	e = &ServerConfigEntity{}
	err = s.db.DB().First(e, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(serverConfigDbName, err)
		}
		return nil, nil
	}
	return
}

func NewServerConfigStore(db coredb.DBService) ServerConfigStore {
	s := &serverConfigStore{
		db: db,
	}
	return s
}
