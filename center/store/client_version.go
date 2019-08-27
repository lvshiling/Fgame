package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type ClientVersionEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	IosVersion     string `gorm:"column:iosVersion"`
	AndroidVersion string `gorm:"column:androidVersion"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *ClientVersionEntity) TableName() string {
	return "t_client_version"
}

type ClientVersionStore interface {
	//获取客户端版本
	GetClientVersion() (*ClientVersionEntity, error)
}

var (
	clientVersionDbName = "client_version"
)

type clientVersionStore struct {
	db coredb.DBService
}

func (s *clientVersionStore) GetClientVersion() (e *ClientVersionEntity, err error) {
	e = &ClientVersionEntity{}
	err = s.db.DB().First(e, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(clientVersionDbName, err)
		}
		return nil, nil
	}
	return
}

func NewClientVersionStore(db coredb.DBService) ClientVersionStore {
	s := &clientVersionStore{
		db: db,
	}
	return s
}
