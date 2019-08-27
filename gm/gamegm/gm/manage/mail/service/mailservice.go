package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	model "fgame/fgame/gm/gamegm/gm/manage/mail/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
)

type IMailService interface {
	//提交
	AddMailInfo(p_mailType int, p_serverId []int, p_title string, p_content string, p_playlist string, p_proplist string, p_freezTme int, p_effectDays int, p_startRoleTime int64, p_endRoleTime int64, p_startLevel int, p_endLevel int, p_user int64, p_sdkType int, p_centerPlatformId []int64, p_bindFlag int, p_remark string, p_mailState int32, p_sendState []int32) error
	UpdateMailInfo(p_id int64, p_mailType int, p_serverId int, p_title string, p_content string, p_playlist string, p_proplist string, p_freezTme int, p_effectDays int, p_startRoleTime int64, p_endRoleTime int64, p_startLevel int, p_endLevel int, p_user int64, p_sdkType int, p_centerPlatformId int64, p_bindFlag int, p_remark string) error
	DeleteInfo(p_id int64) error
	GetApplyList(p_user int64, p_pageindex int, p_title string, p_state int, p_platformList []int64, p_delFlag bool, p_playerId string) ([]*model.MailApply, error)
	GetApplyCount(p_user int64, p_title string, p_state int, p_platformList []int64, p_delFlag bool, p_playerId string) (int, error)

	//审核
	ApproveMail(p_id int64, p_userid int64, p_state int, p_reason string) error
	GetApproveList(p_pageindex int, p_title string, p_state int, p_platformList []int64, p_playerId string) ([]*model.MailApply, error)
	GetApproveCount(p_title string, p_state int, p_platformList []int64, p_playerId string) (int, error)

	//获取
	GetMailInfo(p_id int64) (*model.MailApply, error)

	UpdateSendFlag(p_id int64, p_flag int) error
}

type mailService struct {
	db gmdb.DBService
}

func (m *mailService) AddMailInfo(p_mailType int, p_serverId []int, p_title string, p_content string, p_playlist string, p_proplist string, p_freezTme int, p_effectDays int, p_startRoleTime int64, p_endRoleTime int64, p_startLevel int, p_endLevel int, p_user int64, p_sdkType int, p_centerPlatformId []int64, p_bindFlag int, p_remark string, p_mailState int32, p_sendState []int32) error {
	if len(p_serverId) != len(p_centerPlatformId) || len(p_serverId) != len(p_sendState) {
		return fmt.Errorf("服务器参数与中心id数量不等")
	}
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.db.DB().Begin()
	for index, value := range p_serverId {
		centerid := p_centerPlatformId[index]
		info := &model.MailApply{
			MailType:         p_mailType,
			ServerId:         value,
			Title:            p_title,
			Content:          p_content,
			Playerlist:       p_playlist,
			Proplist:         p_proplist,
			FreezTime:        p_freezTme,
			EffectDays:       p_effectDays,
			RoleStartTime:    p_startRoleTime,
			RoleEndTime:      p_endRoleTime,
			MinLevel:         p_startLevel,
			MaxLevel:         p_endLevel,
			CreateTime:       now,
			MailTime:         now,
			MailState:        int(p_mailState),
			MailUser:         p_user,
			SdkType:          p_sdkType,
			CenterPlatformId: centerid,
			BindFlag:         p_bindFlag,
			Remark:           p_remark,
		}
		info.SendFlag = int(p_sendState[index])
		errdb := exdb.Save(info)
		if errdb.Error != nil {
			exdb.Rollback()
			return errdb.Error
		}
	}
	exdb.Commit()
	return nil
}

func (m *mailService) UpdateMailInfo(p_id int64, p_mailType int, p_serverId int, p_title string, p_content string, p_playlist string, p_proplist string, p_freezTme int, p_effectDays int, p_startRoleTime int64, p_endRoleTime int64, p_startLevel int, p_endLevel int, p_user int64, p_sdkType int, p_centerPlatformId int64, p_bindFlag int, p_remark string) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := &model.MailApply{}
	errdb := m.db.DB().Where("id = ?", p_id).First(info)
	if errdb.Error != nil {
		return errdb.Error
	}
	info.MailType = p_mailType
	info.ServerId = p_serverId
	info.Title = p_title
	info.Content = p_content
	info.Playerlist = p_playlist
	info.Proplist = p_proplist
	info.FreezTime = p_freezTme
	info.EffectDays = p_effectDays
	info.RoleStartTime = p_startRoleTime
	info.RoleEndTime = p_endRoleTime
	info.MinLevel = p_startLevel
	info.MaxLevel = p_endLevel
	info.UpdateTime = now
	info.MailTime = now
	info.MailUser = p_user
	info.SdkType = p_sdkType
	info.CenterPlatformId = p_centerPlatformId
	info.BindFlag = p_bindFlag
	info.Remark = p_remark
	errdb = m.db.DB().Save(info)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *mailService) DeleteInfo(p_id int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_mail_apply").Where("id = ?", p_id).Update("deleteTime", now)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *mailService) GetApplyList(p_user int64, p_pageindex int, p_title string, p_state int, p_platformList []int64, p_delFlag bool, p_playerId string) ([]*model.MailApply, error) {
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*model.MailApply, 0)
	whereStr := "mailUser=? "
	if !p_delFlag {
		whereStr += " and deleteTime =0"
	}
	if len(p_title) > 0 {
		whereStr += fmt.Sprintf(" and title LIKE '%s'", "%"+p_title+"%")
	}
	if p_state > 0 {
		whereStr += fmt.Sprintf(" and mailState=%d", p_state)
	}
	if len(p_platformList) > 0 {
		whereStr += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if len(p_playerId) > 0 {
		whereStr += fmt.Sprintf(" and playerlist LIKE '%s'", "%"+p_playerId+"%")
	}
	errdb := m.db.DB().Where(whereStr, p_user).Order("Id desc").Offset(offset).Limit(limit).Find(&rst)
	if errdb.Error != nil {
		return nil, errdb.Error
	}
	return rst, nil
}

func (m *mailService) GetApplyCount(p_user int64, p_title string, p_state int, p_platformList []int64, p_delFlag bool, p_playerId string) (int, error) {
	rst := 0
	whereStr := "mailUser=? "
	if !p_delFlag {
		whereStr += " and deleteTime =0"
	}
	if len(p_title) > 0 {
		whereStr += fmt.Sprintf(" and title LIKE '%s'", "%"+p_title+"%")
	}
	if p_state > 0 {
		whereStr += fmt.Sprintf(" and mailState=%d", p_state)
	}
	if len(p_platformList) > 0 {
		whereStr += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if len(p_playerId) > 0 {
		whereStr += fmt.Sprintf(" and playerlist LIKE '%s'", "%"+p_playerId+"%")
	}
	errdb := m.db.DB().Table("t_mail_apply").Where(whereStr, p_user).Count(&rst)
	if errdb.Error != nil {
		return 0, errdb.Error
	}
	return rst, nil
}

func (m *mailService) ApproveMail(p_id int64, p_userid int64, p_state int, p_reason string) error {
	now := timeutils.TimeToMillisecond(time.Now())
	info := &model.MailApply{}
	errdb := m.db.DB().Where("id = ?", p_id).First(info)
	if errdb.Error != nil {
		return errdb.Error
	}
	info.ApproveReason = p_reason
	info.ApproveTime = now
	info.MailState = p_state
	info.UpdateTime = now
	errdb = m.db.DB().Save(info)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *mailService) GetApproveList(p_pageindex int, p_title string, p_state int, p_platformList []int64, p_playerId string) ([]*model.MailApply, error) {
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*model.MailApply, 0)
	whereStr := "deleteTime =0 "
	if len(p_title) > 0 {
		whereStr += fmt.Sprintf(" and title LIKE '%s'", "%"+p_title+"%")
	}
	if p_state > 0 {
		whereStr += fmt.Sprintf(" and mailState=%d", p_state)
	}
	if len(p_platformList) > 0 {
		whereStr += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if len(p_playerId) > 0 {
		whereStr += fmt.Sprintf(" and playerlist LIKE '%s'", "%"+p_playerId+"%")
	}
	errdb := m.db.DB().Where(whereStr).Order("Id desc").Offset(offset).Limit(limit).Find(&rst)
	if errdb.Error != nil {
		return nil, errdb.Error
	}
	return rst, nil
}

func (m *mailService) GetApproveCount(p_title string, p_state int, p_platformList []int64, p_playerId string) (int, error) {
	rst := 0
	whereStr := "deleteTime =0 "
	if len(p_title) > 0 {
		whereStr += fmt.Sprintf(" and title LIKE '%s'", "%"+p_title+"%")
	}
	if p_state > 0 {
		whereStr += fmt.Sprintf(" and mailState=%d", p_state)
	}
	if len(p_platformList) > 0 {
		whereStr += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if len(p_playerId) > 0 {
		whereStr += fmt.Sprintf(" and playerlist LIKE '%s'", "%"+p_playerId+"%")
	}
	errdb := m.db.DB().Table("t_mail_apply").Where(whereStr).Count(&rst)
	if errdb.Error != nil {
		return 0, errdb.Error
	}
	return rst, nil
}

func (m *mailService) GetMailInfo(p_id int64) (*model.MailApply, error) {
	info := &model.MailApply{}
	errdb := m.db.DB().Where("id=?", p_id).First(info)
	if errdb.Error != nil {
		return nil, errdb.Error
	}
	return info, nil
}

func (m *mailService) UpdateSendFlag(p_id int64, p_flag int) error {
	errdb := m.db.DB().Table("t_mail_apply").Where("id = ?", p_id).Update("sendFlag", p_flag)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func NewMailService(p_db gmdb.DBService) IMailService {
	rst := &mailService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	mailServiceKey = contextKey("MailService")
)

func WithMailService(ctx context.Context, ls IMailService) context.Context {
	return context.WithValue(ctx, mailServiceKey, ls)
}

func MailServiceInContext(ctx context.Context) IMailService {
	us, ok := ctx.Value(mailServiceKey).(IMailService)
	if !ok {
		return nil
	}
	return us
}

func SetupMailServiceHandler(ls IMailService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithMailService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
