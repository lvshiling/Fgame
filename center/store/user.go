package store

import (
	coredb "fgame/fgame/core/db"
	timeutils "fgame/fgame/pkg/timeutils"
	"time"

	"github.com/jinzhu/gorm"
)

type UserEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	Platform       int32  `gorm:"column:platform"`
	PlatformUserId string `gorm:"column:platformUserId"`
	Name           string `gorm:"column:name"`
	Password       string `gorm:"column:password"`
	PhoneNum       string `gorm:"column:phoneNum"`
	IdCard         int32  `gorm:"column:idCard"`
	RealName       int32  `gorm:"column:realName"`
	RealNameState  int32  `gorm:"column:realNameState"`
	Gm             int32  `gorm:"column:gm"`
	Forbid         int32  `gorm:"column:forbid"`
	ForbidTime     int64  `gorm:"column:forbidTime"`
	ForbidEndTime  int64  `gorm:"column:forbidEndTime"`
	ForbidName     string `gorm:"column:forbidName"`
	ForbidText     string `gorm:"column:forbidText"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *UserEntity) GetId() int64 {
	return e.Id
}

func (e *UserEntity) TableName() string {
	return "t_user"
}

func NewUserEntity() *UserEntity {
	userEntity := &UserEntity{}
	return userEntity
}

type UserStore interface {
	//获取用户
	GetUserByPlatform(platform int32, platformUserId string) (*UserEntity, error)
	GetUserByNameAndPassword(name string, password string) (*UserEntity, error)
	RegisterPlatformUser(platform int32, platformUserId string, name string, passowrd string) (*UserEntity, error)
	GetUser(userId int64) (*UserEntity, error)
}

type userStore struct {
	db coredb.DBService
}

func (s *userStore) GetUserByNameAndPassword(name string, password string) (e *UserEntity, err error) {
	e = &UserEntity{}
	err = s.db.DB().First(e, "name=? and password=?", name, password).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("user", err)
		}
		return nil, nil
	}
	return
}

func (s *userStore) GetUserByPlatform(platform int32, platformUserId string) (e *UserEntity, err error) {
	e = &UserEntity{}
	err = s.db.DB().First(e, "platform=? and platformUserId=?", platform, platformUserId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("user", err)
		}
		return nil, nil
	}
	return
}

func (s *userStore) RegisterPlatformUser(platform int32, platformUserId string, name string, password string) (e *UserEntity, err error) {
	now := timeutils.TimeToMillisecond(time.Now())
	e = &UserEntity{}
	e.Platform = platform
	e.PlatformUserId = platformUserId
	e.Name = name
	e.Password = password
	e.CreateTime = now
	err = s.db.DB().Save(e).Error
	if err != nil {
		return
	}
	return
}

func (s *userStore) GetUser(userId int64) (e *UserEntity, err error) {
	e = &UserEntity{}
	err = s.db.DB().First(e, "id=?", userId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, coredb.NewDBError("user", err)
		}
		return nil, nil
	}
	return
}

func NewUserStore(db coredb.DBService) UserStore {
	s := &userStore{
		db: db,
	}
	return s
}
