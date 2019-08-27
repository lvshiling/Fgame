package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmError "fgame/fgame/gm/gamegm/error"
	loginnoticemodel "fgame/fgame/gm/gamegm/gm/center/notice/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ILoginNoticeService interface {
	AddLoginNotice(p_platformId int64, p_content string) (int64, error)
	UpdateLoginNotice(p_loginNoticeid int64, p_content string) error
	DeleteLoginNotice(p_loginNoticeid int64) error

	GetLoginNoticeList(p_index int) ([]*loginnoticemodel.LoginNotice, error)
	GetLoginNoticeCount() (int, error)

	GetLoginNotice(p_id int64) (*loginnoticemodel.LoginNotice, error)

	GetDefaultNotice() (*loginnoticemodel.LoginNotice, error)
	UpdateDefaultNotice(p_content string) error
}

type loginNoticeService struct {
	db gmdb.DBService
}

func (m *loginNoticeService) AddLoginNotice(p_platformId int64, p_content string) (int64, error) {

	exflag, err := m.existsLoginNotice(p_platformId)
	if err != nil {
		return 0, err
	}
	if exflag {
		return 0, gmError.GetError(gmError.ErrorCodeCenterLoginNoticeExist)
	}

	modelInfo := &loginnoticemodel.LoginNotice{
		PlatformId: int(p_platformId),
		Content:    p_content,
		CreateTime: timeutils.TimeToMillisecond(time.Now()),
	}
	err = m.saveLoginNotice(modelInfo)
	if err != nil {
		return 0, err
	}
	return modelInfo.Id, nil
}

func (m *loginNoticeService) UpdateLoginNotice(p_loginNoticeid int64, p_content string) error {
	if len(p_content) == 0 {
		return gmError.GetError(gmError.ErrorCodeCenterLoginNoticeEmpty)
	}

	modelInfo, err := m.GetLoginNotice(p_loginNoticeid)
	if err != nil {
		return err
	}
	modelInfo.UpdateTime = timeutils.TimeToMillisecond(time.Now())
	modelInfo.Content = p_content

	err = m.saveLoginNotice(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *loginNoticeService) DeleteLoginNotice(p_loginNoticeid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_notice_login").Where("id = ?", p_loginNoticeid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"loginNoticeid": p_loginNoticeid,
			"error":         errdb.Error,
		}).Error("删除中心服务器失败")
		return errdb.Error
	}
	return nil
}

func (m *loginNoticeService) GetLoginNoticeList(p_index int) ([]*loginnoticemodel.LoginNotice, error) {
	return m.getLoginNoticeList(p_index)
}

func (m *loginNoticeService) GetLoginNoticeCount() (int, error) {
	return m.getLoginNoticeCount()
}

func (m *loginNoticeService) GetLoginNotice(p_id int64) (*loginnoticemodel.LoginNotice, error) {
	info := &loginnoticemodel.LoginNotice{}
	exdb := m.db.DB().Where("id = ?", p_id).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *loginNoticeService) existsLoginNotice(p_platformId int64) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_notice_login").Where("platformId=? and deleteTime = 0", p_platformId).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"p_platformId": p_platformId,
			"error":        exdb.Error,
		}).Error("查询中心服务器信息失败existsLoginNotice")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *loginNoticeService) existsLoginNoticeWithId(p_loginNoticeid int64, p_name string, p_platformId int64) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_notice_login").Where("name = ? and id != ? and platformId=? and deleteTime=0", p_name, p_loginNoticeid, p_platformId).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"loginNoticeName": p_name,
			"loginNoticeid":   p_loginNoticeid,
			"error":           exdb.Error,
		}).Error("查询中心服务器信息失败existsLoginNoticeWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *loginNoticeService) saveLoginNotice(p_info *loginnoticemodel.LoginNotice) error {
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.Id > 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}
	exdb := m.db.DB().Save(p_info)
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"loginNoticeName": p_info.Id,
			"error":           exdb.Error,
		}).Error("保存中心服务器失败")
		return exdb.Error
	}
	return nil
}

func (m *loginNoticeService) getLoginNoticeList(p_index int) ([]*loginnoticemodel.LoginNotice, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*loginnoticemodel.LoginNotice, 0)

	exerr := m.db.DB().Where("deleteTime =0 and platformId != 0").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取中心服务器列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *loginNoticeService) getLoginNoticeCount() (int, error) {
	rst := 0

	exerr := m.db.DB().Table("t_notice_login").Where("deleteTime =0 and platformId != 0").Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取中心服务器列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

func (m *loginNoticeService) GetDefaultNotice() (*loginnoticemodel.LoginNotice, error) {
	info := &loginnoticemodel.LoginNotice{}
	exdb := m.db.DB().Where("platformId = 0 and deleteTime=0").First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return info, exdb.Error
	}
	return info, nil
}

func (m *loginNoticeService) UpdateDefaultNotice(p_content string) error {
	info, err := m.GetDefaultNotice()
	if err != nil {
		return err
	}
	now := timeutils.TimeToMillisecond(time.Now())
	info.Content = p_content
	info.UpdateTime = now
	if info.Id == 0 {
		info.CreateTime = now
	}
	exdb := m.db.DB().Save(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func NewLoginNoticeService(p_db gmdb.DBService) ILoginNoticeService {
	rst := &loginNoticeService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	loginNoticeServiceKey = contextKey("LoginNoticeService")
)

func WithLoginNoticeService(ctx context.Context, ls ILoginNoticeService) context.Context {
	return context.WithValue(ctx, loginNoticeServiceKey, ls)
}

func LoginNoticeServiceInContext(ctx context.Context) ILoginNoticeService {
	us, ok := ctx.Value(loginNoticeServiceKey).(ILoginNoticeService)
	if !ok {
		return nil
	}
	return us
}

func SetupLoginNoticeServiceHandler(ls ILoginNoticeService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithLoginNoticeService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
