package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/center/user/model"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ICenterUserService interface {
	GetUserList(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_index int, p_platformArray []int) ([]*model.UserInfo, error)
	GetUserCount(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_platformArray []int) (int, error)
	GetNeiGuaUserList(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_index int, p_platformArray []int) ([]*model.UserInfo, error)
	GetNeiGuaUserCount(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_platformArray []int) (int, error)
	UpdateGm(p_userid int64, p_gmflag int, p_userName string, p_passWord string) error
	UpdateUserInfo(p_userid int64, p_userName string, p_passWord string) error
	UpdateForbid(p_userid int64, p_forbid int, p_forbidTime int64, p_forbidEndTime int64, p_forbidName string, p_forbidText string) error
	ExistsUserName(p_userid int64, p_userName string) (bool, error)
	GetUserInfo(p_userid int64) (*model.UserInfo, error)
	GetIpForbidInfo(p_ip string) (*model.IpForbidInfo, error)
	UpdateIpForbid(p_ip string, p_forbid int, p_forbidTime int64, p_forbidEndTime int64, p_forbidName string, p_forbidText string) error
	GetUserInfoByUserName(p_skdtype int, p_userId string) (*model.UserInfo, error)
}

type centerUserService struct {
	db gmdb.DBService
}

func (m *centerUserService) GetUserList(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_index int, p_platformArray []int) ([]*model.UserInfo, error) {
	rst := make([]*model.UserInfo, 0)
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize

	where := "deleteTime = 0"
	if p_platform > 0 {
		where += fmt.Sprintf(" AND platform = %d", p_platform)
	}
	if p_userid > 0 {
		where += fmt.Sprintf(" AND id = %d", p_userid)
	}
	if len(p_userName) > 0 {
		where += fmt.Sprintf(" AND name LIKE '%s'", "%"+p_userName+"%")
	}
	if len(p_platformUserId) > 0 {
		where += fmt.Sprintf(" AND platformUserId = '%s'", p_platformUserId)
	}
	if len(p_platformArray) > 0 {
		where += fmt.Sprintf(" AND platform IN (%s)", common.CombinIntArray(p_platformArray))
	}
	exdb := m.db.DB().Where(where).Offset(offset).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *centerUserService) GetUserCount(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_platformArray []int) (int, error) {
	rst := 0

	where := "deleteTime = 0"
	if p_platform > 0 {
		where += fmt.Sprintf(" AND platform = %d", p_platform)
	}
	if p_userid > 0 {
		where += fmt.Sprintf(" AND id = %d", p_userid)
	}
	if len(p_userName) > 0 {
		where += fmt.Sprintf(" AND name LIKE '%s'", "%"+p_userName+"%")
	}
	if len(p_platformUserId) > 0 {
		where += fmt.Sprintf(" AND platformUserId = '%s'", p_platformUserId)
	}
	if len(p_platformArray) > 0 {
		where += fmt.Sprintf(" AND platform IN (%s)", common.CombinIntArray(p_platformArray))
	}
	exdb := m.db.DB().Table("t_user").Where(where).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *centerUserService) GetNeiGuaUserList(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_index int, p_platformArray []int) ([]*model.UserInfo, error) {
	rst := make([]*model.UserInfo, 0)
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize

	where := "deleteTime = 0 AND name is not null and name != ''"
	if p_platform > 0 {
		where += fmt.Sprintf(" AND platform = %d", p_platform)
	}
	if p_userid > 0 {
		where += fmt.Sprintf(" AND id = %d", p_userid)
	}
	if len(p_userName) > 0 {
		where += fmt.Sprintf(" AND name LIKE '%s'", "%"+p_userName+"%")
	}
	if len(p_platformUserId) > 0 {
		where += fmt.Sprintf(" AND platformUserId = '%s'", p_platformUserId)
	}
	if len(p_platformArray) > 0 {
		where += fmt.Sprintf(" AND platform IN (%s)", common.CombinIntArray(p_platformArray))
	}
	exdb := m.db.DB().Where(where).Offset(offset).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *centerUserService) GetNeiGuaUserCount(p_platform int, p_userid int64, p_userName string, p_platformUserId string, p_platformArray []int) (int, error) {
	rst := 0

	where := "deleteTime = 0 AND name is not null and name != ''"
	if p_platform > 0 {
		where += fmt.Sprintf(" AND platform = %d", p_platform)
	}
	if p_userid > 0 {
		where += fmt.Sprintf(" AND id = %d", p_userid)
	}
	if len(p_userName) > 0 {
		where += fmt.Sprintf(" AND name LIKE '%s'", "%"+p_userName+"%")
	}
	if len(p_platformUserId) > 0 {
		where += fmt.Sprintf(" AND platformUserId = '%s'", p_platformUserId)
	}
	if len(p_platformArray) > 0 {
		where += fmt.Sprintf(" AND platform IN (%s)", common.CombinIntArray(p_platformArray))
	}
	exdb := m.db.DB().Table("t_user").Where(where).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *centerUserService) UpdateGm(p_userid int64, p_gmflag int, p_userName string, p_passWord string) error {
	exdb := m.db.DB().Table("t_user").Where("id = ?", p_userid).Updates(map[string]interface{}{"gm": p_gmflag, "name": p_userName, "password": p_passWord})
	if exdb.Error != nil {
		return exdb.Error
	}
	return nil
}

func (m *centerUserService) UpdateUserInfo(p_userid int64, p_userName string, p_passWord string) error {
	exdb := m.db.DB().Table("t_user").Where("id = ?", p_userid).Updates(map[string]interface{}{"name": p_userName, "password": p_passWord})
	if exdb.Error != nil {
		return exdb.Error
	}
	return nil
}

func (m *centerUserService) UpdateForbid(p_userid int64, p_forbid int, p_forbidTime int64, p_forbidEndTime int64, p_forbidName string, p_forbidText string) error {
	exdb := m.db.DB().Table("t_user").Where("id = ?", p_userid).Updates(map[string]interface{}{"forbid": p_forbid, "forbidTime": p_forbidTime, "forbidEndTime": p_forbidEndTime, "forbidName": p_forbidName, "forbidText": p_forbidText})
	if exdb.Error != nil {
		return exdb.Error
	}
	return nil
}

func (m *centerUserService) ExistsUserName(p_userid int64, p_userName string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_user").Where("id != ? and name=?", p_userid, p_userName).Count(&count)
	if exdb.Error != nil {
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerUserService) GetUserInfo(p_userid int64) (*model.UserInfo, error) {
	info := &model.UserInfo{}
	exdb := m.db.DB().Where("id = ?", p_userid).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	info.Id = int(p_userid)
	return info, nil
}

func (m *centerUserService) GetIpForbidInfo(p_ip string) (*model.IpForbidInfo, error) {
	info := &model.IpForbidInfo{}
	exdb := m.db.DB().Where("ip = ?", p_ip).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	info.Ip = p_ip
	return info, nil
}

func (m *centerUserService) UpdateIpForbid(p_ip string, p_forbid int, p_forbidTime int64, p_forbidEndTime int64, p_forbidName string, p_forbidText string) error {
	info, err := m.GetIpForbidInfo(p_ip)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	info.Forbid = p_forbid
	info.ForbidTime = p_forbidTime
	info.ForbidEndTime = p_forbidEndTime
	info.ForbidName = p_forbidName
	info.ForbidText = p_forbidText
	exdb := m.db.DB().Save(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *centerUserService) GetUserInfoByUserName(p_platform int, p_userId string) (*model.UserInfo, error) {
	info := &model.UserInfo{}
	exdb := m.db.DB().Where("platform = ? and platformUserId = ?", p_platform, p_userId).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func NewCenterUserService(p_db gmdb.DBService) ICenterUserService {
	rst := &centerUserService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	centerUserServiceKey = contextKey("CenterUserService")
)

func WithCenterUserService(ctx context.Context, ls ICenterUserService) context.Context {
	return context.WithValue(ctx, centerUserServiceKey, ls)
}

func CenterUserServiceInContext(ctx context.Context) ICenterUserService {
	us, ok := ctx.Value(centerUserServiceKey).(ICenterUserService)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterUserServiceHandler(ls ICenterUserService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterUserService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
