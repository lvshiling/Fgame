package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type MarrypriceEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlatformId int32 `gorm:"column:platformId"`
	KindType   int32 `gorm:"column:kindType"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *MarrypriceEntity) TableName() string {
	return "t_platform_marryprice"
}

type MarryPriceStore interface {
	//获取所有服务器
	GetAll() ([]*MarrypriceEntity, error)
	Get(platform int32) (*MarrypriceEntity, error)
}

var (
	marryPriceDbName = "marryprice"
)

type marryPriceStore struct {
	db coredb.DBService
}

func (s *marryPriceStore) GetAll() (eList []*MarrypriceEntity, err error) {
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(marryPriceDbName, err)
		}
		return nil, nil
	}
	return
}

func (s *marryPriceStore) Get(platform int32) (e *MarrypriceEntity, err error) {
	e = &MarrypriceEntity{}
	err = s.db.DB().First(e, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(marryPriceDbName, err)
		}
		return nil, nil
	}
	return
}

func NewMarryPriceStore(db coredb.DBService) MarryPriceStore {
	s := &marryPriceStore{
		db: db,
	}
	return s
}
