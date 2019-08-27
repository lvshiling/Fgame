package login

import (
	"context"
	fgamedb "fgame/fgame/core/db"

	fgameredis "fgame/fgame/core/redis"
	"fgame/fgame/login/model"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id         int64  `json:"id"`
	DeviceMac  string `json:"deviceMac"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	UpdateTime int64  `json:"updateTime"`
	CreateTime int64  `json:"createTime"`
	DeleteTime int64  `json:"deleteTime"`
}

func (u *User) ConvertToModel() (um interface{}, err error) {
	um = convertToModel(u)
	return
}

func convertFromModel(um *model.User) (u *User) {
	u = &User{
		Id:         um.Id,
		DeviceMac:  um.DeviceMac,
		Name:       um.Name,
		UpdateTime: um.UpdateTime,
		CreateTime: um.CreateTime,
		DeleteTime: um.DeleteTime,
	}
	return
}

func convertToModel(um *User) (u *model.User) {
	u = &model.User{
		Id:         um.Id,
		DeviceMac:  um.DeviceMac,
		Name:       um.Name,
		UpdateTime: um.UpdateTime,
		CreateTime: um.CreateTime,
		DeleteTime: um.DeleteTime,
	}
	return
}

type UserService interface {
	Register(deviceMac string, name string) (user *User, err error)
	RegisterUser(name string, password string) (user *User, err error)
	GetUserById(id int64) (u *model.User, err error)
	GetUserByName(name string) (u *User, err error)
	GetUserByDeviceMac(deviceMac string) (u *User, err error)
}

type userService struct {
	rs fgameredis.RedisService
	db fgamedb.DBService
}

func (us *userService) Register(deviceMac string, name string) (u *User, err error) {
	user := &model.User{}
	user.DeviceMac = deviceMac
	user.Name = name
	user.CreateTime = time.Now().UnixNano() / int64(time.Millisecond)
	user.UpdateTime = user.CreateTime
	err = us.db.DB().Save(user).Error
	if err != nil {
		return
	}
	u = convertFromModel(user)
	return
}

func (us *userService) RegisterUser(name string, password string) (u *User, err error) {
	user := &model.User{}
	user.Name = name
	user.Password = password
	user.CreateTime = time.Now().UnixNano() / int64(time.Millisecond)
	user.UpdateTime = user.CreateTime
	err = us.db.DB().Save(user).Error
	if err != nil {
		return
	}
	u = convertFromModel(user)
	return
}

func (us *userService) GetUserById(id int64) (u *model.User, err error) {
	u = &model.User{}
	tdb := us.db.DB().First(u, "id=?", id)
	err = tdb.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (us *userService) GetUserByDeviceMac(deviceMac string) (user *User, err error) {
	u := &model.User{}
	tdb := us.db.DB().First(u, "deviceMac=?", deviceMac)
	err = tdb.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	user = convertFromModel(u)
	return
}

func (us *userService) GetUserByName(name string) (user *User, err error) {
	u := &model.User{}
	tdb := us.db.DB().First(u, "name=?", name)
	err = tdb.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	user = convertFromModel(u)
	return
}

func NewUserService(db fgamedb.DBService, rs fgameredis.RedisService) (us UserService) {
	us = &userService{
		rs: rs,
		db: db,
	}
	return
}

const (
	userServiceKey = "UserService"
)

func WithUserService(ctx context.Context, us UserService) context.Context {
	return context.WithValue(ctx, userServiceKey, us)
}

func UserServiceInContext(ctx context.Context) UserService {
	us, ok := ctx.Value(userServiceKey).(UserService)
	if !ok {
		return nil
	}
	return us
}
