package store

import (
	coredb "fgame/fgame/core/db"

	"github.com/jinzhu/gorm"
)

type NoticeEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlatformId int32  `gorm:"column:platformId"`
	Content    string `gorm:"column:content"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *NoticeEntity) TableName() string {
	return "t_notice_login"
}

type NoticeStore interface {
	//获取所有服务器
	GetAll() ([]*NoticeEntity, error)
}

var (
	noticeDbName = "notice"
)

type noticeStore struct {
	db coredb.DBService
}

func (s *noticeStore) GetAll() (eList []*NoticeEntity, err error) {
	err = s.db.DB().Find(&eList, "deleteTime=0").Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError(dbName, err)
		}
		return
	}
	return
}

func NewNoticeStore(db coredb.DBService) NoticeStore {
	s := &noticeStore{
		db: db,
	}
	return s
}
