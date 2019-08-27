package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmError "fgame/fgame/gm/gamegm/error"
	platformmodel "fgame/fgame/gm/gamegm/gm/platform/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	usermodel "fgame/fgame/gm/gamegm/gm/user/model"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IPlatformService interface {
	AddPlatform(p_name string, p_channelid int64, p_center int64, p_sdkType int, p_signKey string) error
	UpdatePlatform(p_platformid int64, p_name string, p_channelid int64, p_center int64, p_sdkType int, p_signKey string) error
	DeletePlatform(p_platformid int64) error

	GetPlatformList(p_name string, p_channelid int64, p_index int) ([]*platformmodel.PlatformInfo, error)
	GetAllPlatformList() ([]*platformmodel.PlatformInfo, error)
	GetPlatformCount(p_name string, p_channelid int64) (int, error)

	GetPlatformInfo(p_platformid int64) (*platformmodel.PlatformInfo, error)
	GetPlatformInfoArray(p_platformid []int64) ([]*platformmodel.PlatformInfo, error)
	GetPlatformByChannel(p_channelid int64) ([]*platformmodel.PlatformInfo, error)
	GetPlatformByChannelArray(p_channelid []int64) ([]*platformmodel.PlatformInfo, error)

	//获取用户所有的中心平台
	GetAllUserCenterPlatformList(p_userId int64) ([]int64, error)
}

type platformService struct {
	db gmdb.DBService
}

func (m *platformService) AddPlatform(p_name string, p_channelid int64, p_center int64, p_sdkType int, p_signKey string) error {
	if len(p_name) == 0 || p_channelid < 1 {
		return gmError.GetError(gmError.ErrorCodePlatformEmpty)
	}

	exflag, err := m.existsPlatform(p_name)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodePlatformExist)
	}

	exflag, err = m.existsPlatformSdk(p_sdkType)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodePlatformExistSdk)
	}

	modelInfo := &platformmodel.PlatformInfo{
		PlatformName:     p_name,
		ChannelId:        p_channelid,
		CenterPlatformID: p_center,
		SdkType:          p_sdkType,
		SignKey:          p_signKey,
	}
	err = m.savePlatform(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *platformService) UpdatePlatform(p_platformid int64, p_name string, p_channelid int64, p_center int64, p_sdkType int, p_signKey string) error {
	if len(p_name) == 0 || p_channelid < 1 {
		return gmError.GetError(gmError.ErrorCodePlatformEmpty)
	}

	exflag, err := m.existsPlatformWithId(p_platformid, p_name)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodePlatformExist)
	}

	exflag, err = m.existsPlatformWithIdSdk(p_platformid, p_sdkType)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodePlatformExistSdk)
	}

	modelInfo := &platformmodel.PlatformInfo{
		PlatformID:       p_platformid,
		PlatformName:     p_name,
		ChannelId:        p_channelid,
		CenterPlatformID: p_center,
		SdkType:          p_sdkType,
		SignKey:          p_signKey,
	}
	err = m.savePlatform(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *platformService) DeletePlatform(p_platformid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_platform").Where("platformId = ?", p_platformid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"platformid": p_platformid,
			"error":      errdb.Error,
		}).Error("删除平台失败")
		return errdb.Error
	}
	return nil
}

func (m *platformService) GetPlatformList(p_name string, p_channelid int64, p_index int) ([]*platformmodel.PlatformInfo, error) {
	return m.getPlatformList(p_name, p_channelid, p_index)
}

func (m *platformService) GetPlatformCount(p_name string, p_channelid int64) (int, error) {
	return m.getPlatformCount(p_name, p_channelid)
}

func (m *platformService) GetPlatformInfo(p_platformid int64) (*platformmodel.PlatformInfo, error) {
	info := &platformmodel.PlatformInfo{}

	dber := m.db.DB().Where("platformId = ?", p_platformid).First(info)
	if dber.Error != nil && dber.Error != gorm.ErrRecordNotFound {
		return info, nil
	}

	return info, nil
}

func (m *platformService) GetPlatformInfoArray(p_platformid []int64) ([]*platformmodel.PlatformInfo, error) {
	rst := make([]*platformmodel.PlatformInfo, 0)

	dber := m.db.DB().Where("platformId IN (?)", p_platformid).First(&rst)
	if dber.Error != nil && dber.Error != gorm.ErrRecordNotFound {
		return rst, nil
	}

	return rst, nil
}

func (m *platformService) GetPlatformByChannel(p_channelid int64) ([]*platformmodel.PlatformInfo, error) {
	rst := make([]*platformmodel.PlatformInfo, 0)
	dber := m.db.DB().Where("channelId = ? and deleteTime = 0", p_channelid).Find(&rst)
	if dber.Error != nil && dber.Error != gorm.ErrRecordNotFound {
		return rst, nil
	}
	return rst, nil
}

func (m *platformService) GetPlatformByChannelArray(p_channelid []int64) ([]*platformmodel.PlatformInfo, error) {
	rst := make([]*platformmodel.PlatformInfo, 0)
	dber := m.db.DB().Where("channelId IN (?) and deleteTime = 0", p_channelid).Find(&rst)
	if dber.Error != nil && dber.Error != gorm.ErrRecordNotFound {
		return rst, nil
	}
	return rst, nil
}

func (m *platformService) existsPlatform(p_name string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("platformName = ? and deleteTime = 0", p_name).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformName": p_name,
			"error":        exdb.Error,
		}).Error("查询平台信息失败existsPlatform")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *platformService) existsPlatformWithId(p_platformid int64, p_name string) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("platformName = ? and platformId != ? and deleteTime=0", p_name, p_platformid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformName": p_name,
			"platformid":   p_platformid,
			"error":        exdb.Error,
		}).Error("查询平台信息失败existsPlatformWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *platformService) existsPlatformSdk(p_sdkType int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("sdkType = ? and deleteTime = 0", p_sdkType).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"p_sdkType": p_sdkType,
			"error":     exdb.Error,
		}).Error("查询平台信息失败existsPlatformSdk")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *platformService) existsPlatformWithIdSdk(p_platformid int64, p_sdkType int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform").Where("sdkType = ? and platformId != ? and deleteTime=0", p_sdkType, p_platformid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"p_sdkType":  p_sdkType,
			"platformid": p_platformid,
			"error":      exdb.Error,
		}).Error("查询平台信息失败existsPlatformWithIdSdk")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *platformService) savePlatform(p_info *platformmodel.PlatformInfo) error {
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.PlatformID > 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}

	prePlatFormInfo := &platformmodel.PlatformInfo{}
	if p_info.PlatformID > 0 {
		exdbdd := m.db.DB().Where("platformId = ?", p_info.PlatformID).First(&prePlatFormInfo)
		if exdbdd.Error != nil {
			log.WithFields(log.Fields{
				"platformName": p_info.PlatformName,
				"platformid":   p_info.PlatformID,
				"error":        exdbdd.Error,
			}).Error("保存平台失败,获取平台失败")
			return exdbdd.Error
		}
	}

	trans := m.db.DB().Begin()

	exdb := trans.Save(p_info)
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"platformName": p_info.PlatformID,
			"platformid":   p_info.PlatformName,
			"error":        exdb.Error,
		}).Error("保存平台失败")
		trans.Rollback()
		return exdb.Error
	}
	if p_info.PlatformID > 0 && prePlatFormInfo.ChannelId != p_info.ChannelId { //更改的时候也要将用户里的给更新
		exdb = trans.Table("t_gmuser").Where("platformId = ? and deleteTime=0", p_info.PlatformID).Update("channelId", p_info.ChannelId)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			log.WithFields(log.Fields{
				"platformName": p_info.PlatformID,
				"platformid":   p_info.PlatformName,
				"error":        exdb.Error,
			}).Error("更新用户那边的失败")
			trans.Rollback()
			return exdb.Error
		}
	}
	trans.Commit()
	return nil
}

func (m *platformService) getPlatformList(p_name string, p_channelid int64, p_index int) ([]*platformmodel.PlatformInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*platformmodel.PlatformInfo, 0)
	whereStr := "deleteTime =0"
	if p_channelid > 0 {
		whereStr += fmt.Sprintf(" and channelId=%d", p_channelid)
	}
	exerr := m.db.DB().Where(whereStr+" and platformName like ?", "%"+p_name+"%").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformName": p_name,
			"error":        exerr.Error,
		}).Error("获取平台列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *platformService) getPlatformCount(p_name string, p_channelid int64) (int, error) {
	rst := 0
	whereStr := "deleteTime =0"
	if p_channelid > 0 {
		whereStr += fmt.Sprintf(" and channelId=%d", p_channelid)
	}
	exerr := m.db.DB().Table("t_platform").Where(whereStr+" and platformName like ?", "%"+p_name+"%").Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformName": p_name,
			"error":        exerr.Error,
		}).Error("获取平台列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

func (m *platformService) GetAllPlatformList() ([]*platformmodel.PlatformInfo, error) {
	rst := make([]*platformmodel.PlatformInfo, 0)
	exerr := m.db.DB().Where("deleteTime =0").Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取All平台列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *platformService) GetAllUserCenterPlatformList(p_userId int64) ([]int64, error) {
	userInfo := &usermodel.DBGmUserInfo{}
	exdb := m.db.DB().Where("id = ?", p_userId).First(userInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	rst := make([]int64, 0)
	if userInfo.PlatformId > 0 { //拥有平台
		platformList := &platformmodel.PlatformInfo{}
		exdb = m.db.DB().Where("platformId = ? and deleteTime = 0", userInfo.PlatformId).First(platformList)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			return nil, exdb.Error
		}
		rst = append(rst, platformList.CenterPlatformID)
	} else {
		if userInfo.ChannelID > 0 { //拥有渠道
			platformList := make([]*platformmodel.PlatformInfo, 0)
			exdb = m.db.DB().Where("channelId = ? and deleteTime = 0", userInfo.ChannelID).First(&platformList)
			if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
				return nil, exdb.Error
			}
			for _, value := range platformList {
				rst = append(rst, value.CenterPlatformID)
			}
		}
	}
	return rst, nil
}

func NewPlatformService(p_db gmdb.DBService) IPlatformService {
	rst := &platformService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	platformServiceKey = contextKey("PlatformService")
	apiPlatformId      = contextKey("ApiPlatformId")
	apiPlatformInfo    = contextKey("ApiPlatformInfo")
)

func WithPlatformService(ctx context.Context, ls IPlatformService) context.Context {
	return context.WithValue(ctx, platformServiceKey, ls)
}

func PlatformServiceInContext(ctx context.Context) IPlatformService {
	us, ok := ctx.Value(platformServiceKey).(IPlatformService)
	if !ok {
		return nil
	}
	return us
}

func SetupPlatformServiceHandler(ls IPlatformService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithPlatformService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

func WithApiPlatformId(ctx context.Context, ls int64) context.Context {
	return context.WithValue(ctx, apiPlatformId, ls)
}

func ApiPlatformIdInContext(ctx context.Context) int64 {
	us, ok := ctx.Value(apiPlatformId).(int64)
	if !ok {
		return 0
	}
	return us
}

func WithApiPlatformInfo(ctx context.Context, ls *platformmodel.PlatformInfo) context.Context {
	return context.WithValue(ctx, apiPlatformInfo, ls)
}

func ApiPlatformInfoInContext(ctx context.Context) *platformmodel.PlatformInfo {
	us, ok := ctx.Value(apiPlatformInfo).(*platformmodel.PlatformInfo)
	if !ok {
		return nil
	}
	return us
}
