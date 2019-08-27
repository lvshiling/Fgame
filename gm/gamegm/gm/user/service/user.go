package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmusermodel "fgame/fgame/gm/gamegm/gm/user/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	gmredis "fgame/fgame/gm/gamegm/redis"
	"fmt"
	"net/http"
	"strconv"
	"time"

	gmError "fgame/fgame/gm/gamegm/error"
	types "fgame/fgame/gm/gamegm/gm/types"

	platmode "fgame/fgame/gm/gamegm/gm/platform/model"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IGmUserService interface {
	GetUserList(p_userName string, p_privilege int, p_index int, p_channel int, p_platform int, p_child []int) ([]*gmusermodel.DBGmUserInfo, error)
	GetUserCount(p_userName string, p_privilege int, p_channel int, p_platform int, p_child []int) (int, error)

	AddUser(p_user *gmusermodel.DBGmUserInfo) error
	UpdateUser(p_user *gmusermodel.DBGmUserInfo) error
	DeleteUser(p_userid int64) error
	ChangePassWord(p_userid int64, p_password string) error
	GetUserInfo(p_userId int64) (*gmusermodel.DBGmUserInfo, error)

	GetUserCenterPlatList(p_userId int64) ([]int64, error)
	GetUserSdkTypeList(p_userId int64) ([]int, error)
}

type gmUserService struct {
	_db gmdb.DBService
}

func (m *gmUserService) GetUserList(p_userName string, p_privilege int, p_index int, p_channel int, p_platform int, p_child []int) ([]*gmusermodel.DBGmUserInfo, error) {
	rst := make([]*gmusermodel.DBGmUserInfo, 0)

	if len(p_child) == 0 {
		return nil, nil
	}
	offest := (p_index - 1) * constant.DefaultPageSize
	if offest < 0 {
		offest = 0
	}
	limit := constant.DefaultPageSize
	whereName := "%" + p_userName + "%"
	wherePrivile := "and deleteTime =0 "
	if p_privilege >= 0 {
		wherePrivile += " and privilege_level=" + strconv.Itoa(p_privilege)
	}
	if p_channel > 0 {
		wherePrivile += fmt.Sprintf(" and channelId=%d ", p_channel)
	}
	if p_platform > 0 {
		wherePrivile += fmt.Sprintf(" and platformId=%d ", p_platform)
	}
	if len(p_child) > 0 {
		childStr := ""
		for index, value := range p_child {
			if index != 0 {
				childStr += ","
			}
			childStr += fmt.Sprintf("%d", value)
		}
		wherePrivile += fmt.Sprintf(" and privilege_level IN (%s)", childStr)
	}
	selDB := m._db.DB().Limit(limit).Offset(offest).Order("id asc").Where("userName LIKE ? "+wherePrivile, whereName)

	exDB := selDB.Find(&rst)
	if exDB.Error != nil {
		return nil, exDB.Error
	}
	return rst, nil
}

func (m *gmUserService) GetUserCount(p_userName string, p_privilege int, p_channel int, p_platform int, p_child []int) (int, error) {

	if len(p_child) == 0 {
		return 0, nil
	}
	whereName := "%" + p_userName + "%"
	wherePrivile := "and deleteTime =0 "
	if p_privilege >= 0 {
		wherePrivile += " and privilege_level=" + strconv.Itoa(p_privilege)
	}
	if p_channel > 0 {
		wherePrivile += fmt.Sprintf(" and channelId=%d ", p_channel)
	}
	if p_platform > 0 {
		wherePrivile += fmt.Sprintf(" and platformId=%d ", p_platform)
	}
	if len(p_child) > 0 {
		childStr := ""
		for index, value := range p_child {
			if index != 0 {
				childStr += ","
			}
			childStr += fmt.Sprintf("%d", value)
		}
		wherePrivile += fmt.Sprintf(" and privilege_level IN (%s)", childStr)
	}
	rst := 0
	selDB := m._db.DB().Table("t_gmuser").Order("id asc").Where("deleteTime =0 AND userName LIKE ? "+wherePrivile, whereName).Count(&rst)
	if selDB.Error != nil {
		return 0, selDB.Error
	}
	return rst, nil
}

func (m *gmUserService) AddUser(p_user *gmusermodel.DBGmUserInfo) error {
	if len(p_user.UserName) == 0 || len(p_user.Psd) == 0 || p_user.PrivilegeLevel < 1 {
		return gmError.GetError(gmError.ErrorCodeUserParamEmpty)
	}
	userPrivilege := types.PrivilegeLevel(p_user.PrivilegeLevel)
	if userPrivilege == types.PrivilegeLevelChannel && p_user.ChannelID < 1 {
		return gmError.GetError(gmError.ErrorCodeMissChannel)
	}
	if userPrivilege == types.PrivilegeLevelPlatform && p_user.PlatformId < 1 {
		return gmError.GetError(gmError.ErrorCodeMissPlatform)
	}
	exflag, exErr := m.existUserName(p_user.UserName)
	if exErr != nil {
		return exErr
	}
	if exflag { //已经存在用户名
		return gmError.GetError(gmError.ErrorCodeUserExist)
	}
	p_user.CreateTime = timeutils.TimeToMillisecond(time.Now())
	if !userPrivilege.HasChannel() {
		p_user.ChannelID = 0
	}

	if !userPrivilege.HasPlatform() {
		p_user.PlatformId = 0
	}
	err := m.save(p_user)
	if err != nil {
		return exErr
	}
	return nil
}

func (m *gmUserService) UpdateUser(p_user *gmusermodel.DBGmUserInfo) error {
	if len(p_user.UserName) == 0 || p_user.PrivilegeLevel < 1 || p_user.UserId < 1 {
		return gmError.GetError(gmError.ErrorCodeUserParamEmpty)
	}

	userPrivilege := types.PrivilegeLevel(p_user.PrivilegeLevel)
	if userPrivilege == types.PrivilegeLevelChannel && p_user.ChannelID < 1 {
		return gmError.GetError(gmError.ErrorCodeMissChannel)
	}
	if (userPrivilege == types.PrivilegeLevelPlatform || userPrivilege == types.PrivilegeLevelKeFu) && p_user.PlatformId < 1 {
		return gmError.GetError(gmError.ErrorCodeMissPlatform)
	}

	if !userPrivilege.HasChannel() {
		p_user.ChannelID = 0
	}

	if !userPrivilege.HasPlatform() {
		p_user.PlatformId = 0
	}

	exflag, exErr := m.existUserNameWithId(p_user.UserId, p_user.UserName)
	if exErr != nil {
		return exErr
	}
	if exflag { //已经存在用户名
		return gmError.GetError(gmError.ErrorCodeUserExist)
	}
	myuser, err := m.getUser(p_user.UserId)
	if err != nil {
		return err
	}
	myuser.UserName = p_user.UserName
	myuser.PrivilegeLevel = p_user.PrivilegeLevel
	myuser.ChannelID = p_user.ChannelID
	myuser.PlatformId = p_user.PlatformId
	myuser.UpdateTime = timeutils.TimeToMillisecond(time.Now())

	err = m.save(myuser)
	if err != nil {
		return err
	}
	return nil
}

func (m *gmUserService) DeleteUser(p_userid int64) error {
	dberr := m._db.DB().Table("t_gmuser").Where("id = ?", p_userid).Update("deleteTime", timeutils.TimeToMillisecond(time.Now()))
	if dberr.Error != nil {
		return dberr.Error
	}
	return nil
}

func (m *gmUserService) ChangePassWord(p_userid int64, p_password string) error {
	dberr := m._db.DB().Table("t_gmuser").Where("id = ?", p_userid).Update("psd", p_password)
	if dberr.Error != nil {
		return dberr.Error
	}
	return nil
}

func (m *gmUserService) GetUserInfo(p_userId int64) (*gmusermodel.DBGmUserInfo, error) {
	rst := &gmusermodel.DBGmUserInfo{}
	dberr := m._db.DB().Where("id=?", p_userId).First(rst)
	if dberr.Error != nil && dberr.Error != gorm.ErrRecordNotFound {
		return nil, dberr.Error
	}
	return rst, nil

}

func (m *gmUserService) getUser(p_id int64) (*gmusermodel.DBGmUserInfo, error) {
	rst := &gmusermodel.DBGmUserInfo{}
	getErr := m._db.DB().Where("id = ?", p_id).First(rst)
	if getErr.Error != nil && getErr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"userId": p_id,
			"error":  getErr.Error,
		}).Error("获取用户信息失败getUser")
		return nil, getErr.Error
	}
	return rst, nil
}

func (m *gmUserService) save(p_user *gmusermodel.DBGmUserInfo) error {
	savedb := m._db.DB().Save(p_user)
	if savedb.Error != nil {
		log.WithFields(log.Fields{
			"userName":       p_user.UserName,
			"userId":         p_user.UserId,
			"Avator":         p_user.Avator,
			"PrivilegeLevel": p_user.PrivilegeLevel,
			"error":          savedb.Error,
		}).Error("保存用户信息失败userSave")
		return savedb.Error
	}
	return nil
}

func (m *gmUserService) existUserName(p_username string) (bool, error) {
	count := 0
	seldb := m._db.DB().Table("t_gmuser").Where("deleteTime=0 and userName=?", p_username).Count(&count)
	if seldb.Error != nil && seldb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"userName": p_username,
			"error":    seldb.Error,
		}).Error("检查用户失败existUserName")
		return false, seldb.Error
	}
	rst := false
	if count > 0 {
		rst = true
	}
	return rst, nil
}

func (m *gmUserService) existUserNameWithId(p_player int64, p_username string) (bool, error) {
	count := 0
	seldb := m._db.DB().Table("t_gmuser").Where("deleteTime=0 and userName=? AND id != ?", p_username, p_player).Count(&count)
	if seldb.Error != nil && seldb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"userName": p_username,
			"error":    seldb.Error,
		}).Error("检查用户失败existUserNameWithId")
		return false, seldb.Error
	}
	rst := false
	if count > 0 {
		rst = true
	}
	return rst, nil
}

func (m *gmUserService) GetUserCenterPlatList(p_userId int64) ([]int64, error) {
	rst := make([]int64, 0)
	userInfo, err := m.GetUserInfo(p_userId)
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return rst, fmt.Errorf("nil userInfo")
	}
	if userInfo.PlatformId > 0 {
		platInfo := &platmode.PlatformInfo{}
		exdb := m._db.DB().Where("platformId = ?", userInfo.PlatformId).First(platInfo)
		if exdb.Error != nil {
			return rst, exdb.Error
		}
		rst = append(rst, platInfo.CenterPlatformID)
		return rst, nil
	}
	if userInfo.ChannelID > 0 {
		platList := make([]*platmode.PlatformInfo, 0)
		exdb := m._db.DB().Where("channelId = ?", userInfo.ChannelID).Find(&platList)
		if exdb.Error != nil {
			return rst, exdb.Error
		}
		for _, value := range platList {
			rst = append(rst, value.CenterPlatformID)
		}
		return rst, nil
	}

	return rst, nil
}

func (m *gmUserService) GetUserSdkTypeList(p_userId int64) ([]int, error) {
	rst := make([]int, 0)
	userInfo, err := m.GetUserInfo(p_userId)
	if err != nil {
		return nil, err
	}
	if userInfo == nil {
		return rst, fmt.Errorf("nil userInfo")
	}
	if userInfo.PlatformId > 0 {
		platInfo := &platmode.PlatformInfo{}
		exdb := m._db.DB().Where("platformId = ?", userInfo.PlatformId).First(platInfo)
		if exdb.Error != nil {
			return rst, exdb.Error
		}
		rst = append(rst, platInfo.SdkType)
		return rst, nil
	}
	if userInfo.ChannelID > 0 {
		platList := make([]*platmode.PlatformInfo, 0)
		exdb := m._db.DB().Where("channelId = ?", userInfo.ChannelID).Find(&platList)
		if exdb.Error != nil {
			return rst, exdb.Error
		}
		for _, value := range platList {
			rst = append(rst, value.SdkType)
		}
		return rst, nil
	}

	return rst, nil
}

func NewGmUserService(p_db gmdb.DBService) IGmUserService {
	rst := &gmUserService{
		_db: p_db,
	}
	return rst
}

var (
	_singleGMUserService IGmUserService
)

func GetGmUserServiceInstance() IGmUserService {
	return _singleGMUserService
}

func InitGmUserService() {
	_singleGMUserService = NewGmUserService(gmdb.GetInstanceDB())
}

type contextKey string

const (
	userServiceKey = contextKey("GMUserService")
)

func WithGmUserService(ctx context.Context, us IGmUserService) context.Context {
	return context.WithValue(ctx, userServiceKey, us)
}

func GmUserServiceInContext(ctx context.Context) IGmUserService {
	us, ok := ctx.Value(userServiceKey).(IGmUserService)
	if !ok {
		return nil
	}
	return us
}

func SetupUserServiceHandler(ls IGmUserService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithGmUserService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

const (
	gmUserIdKey = contextKey("GmUserId")
	gmUserName  = contextKey("GmUserName")
)

func WithGmUserId(ctx context.Context, gmUserId int64) context.Context {
	return context.WithValue(ctx, gmUserIdKey, gmUserId)
}

func GmUserIdInContext(ctx context.Context) int64 {
	tDealerId, ok := ctx.Value(gmUserIdKey).(int64)
	if !ok {
		return 0
	}
	return int64(tDealerId)
}

func WithGmUserName(ctx context.Context, gmUserName string) context.Context {
	return context.WithValue(ctx, gmUserName, gmUserName)
}

func GmUserNameInContext(ctx context.Context) string {
	tDealerId, ok := ctx.Value(gmUserName).(string)
	if !ok {
		return ""
	}
	return tDealerId
}

const (
	privilegeKey = contextKey("Privilege")
)

func WithPrivilege(ctx context.Context, privilege int32) context.Context {
	return context.WithValue(ctx, privilegeKey, privilege)
}

func PrivilegeInContext(ctx context.Context) int32 {
	privilege, ok := ctx.Value(privilegeKey).(int32)
	if !ok {
		return 0
	}
	return privilege
}

const (
	userRedisKey        = "game.gm.user"
	userSessionRedisKey = "game.gm.user.session"
	privilegeRedisKey   = "game.gm.user.privilege"
)

func GetGmUserKey(dealerId int64) string {
	return gmredis.Combine(userRedisKey, fmt.Sprintf("%d", dealerId))
}

func getGmUserSessionKey(dealerId int64) string {
	return gmredis.Combine(userSessionRedisKey, fmt.Sprintf("%d", dealerId))
}

func getGmUserPrivilegeKey(dealerId int64) string {
	return gmredis.Combine(privilegeRedisKey, fmt.Sprintf("%d", dealerId))
}
