package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type PlatformEntity struct {
	CenterPlatformId int32 `gorm:"column:centerPlatformId"`
	SdkType          int32 `gorm:"column:sdkType"`
}

func (e *PlatformEntity) TableName() string {
	return "t_platform"
}

type PlatformStore interface {
	//获取所有服务器
	GetAll() ([]*PlatformEntity, error)
}

type platformStore struct {
	db coredb.DBService
}

func (s *platformStore) GetAll() (eList []*PlatformEntity, err error) {
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("platform", err)
		}
		return
	}
	return
}

func NewPlatformStore(db coredb.DBService) PlatformStore {
	s := &platformStore{
		db: db,
	}
	return s
}
