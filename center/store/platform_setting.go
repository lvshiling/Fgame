package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type PlatformSettingEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	PlatformId     int32  `gorm:"column:platformId"`
	SettingContent string `gorm:"column:settingContent"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *PlatformSettingEntity) TableName() string {
	return "t_platform_setting"
}

type PlatformSettingStore interface {
	//获取平台配置
	GetPlatformSetting(platform int32) (*PlatformSettingEntity, error)
	//获取所有平台配置
	GetAllPlatformSetting() ([]*PlatformSettingEntity, error)
}

var (
	platformSettingDbName = "platform_setting"
)

type platformSettingStore struct {
	db coredb.DBService
}

func (s *platformSettingStore) GetPlatformSetting(platform int32) (e *PlatformSettingEntity, err error) {
	e = &PlatformSettingEntity{}
	err = s.db.DB().First(e, "deleteTime=0 and platformId=?", platform).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(platformSettingDbName, err)
		}
		return nil, nil
	}
	return
}

func (s *platformSettingStore) GetAllPlatformSetting() (eList []*PlatformSettingEntity, err error) {
	eList = make([]*PlatformSettingEntity, 0, 8)
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(platformSettingDbName, err)
		}
		return nil, nil
	}
	return
}

func NewPlatformSettingStore(db coredb.DBService) PlatformSettingStore {
	s := &platformSettingStore{
		db: db,
	}
	return s
}
