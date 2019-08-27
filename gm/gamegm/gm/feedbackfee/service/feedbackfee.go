package service

import (
	"context"
	"fgame/fgame/gm/gamegm/constant"
	"fgame/fgame/gm/gamegm/db"
	gmdb "fgame/fgame/gm/gamegm/db"
	feemodel "fgame/fgame/gm/gamegm/gm/feedbackfee/model"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type FeedBackFeeService interface {
	GetGameFeedBackFeeList(p_dblink gmdb.GameDbLink, p_serverId int, p_playerId int64, p_code string, p_begin int64, p_end int64, p_index int32) ([]*feemodel.GamePlayerFeedBackFeeExchange, error)
	GetGameFeedBackFeeCount(p_dblink gmdb.GameDbLink, p_serverId int, p_playerId int64, p_code string, p_begin int64, p_end int64) (int32, error)

	GetCenterFeedBackFeeList(p_platformId int32, p_serverId int32, p_playerId int64, p_code string, p_begin int64, p_end int64, p_index int32) ([]*feemodel.CenterPlayerFeedBackFeeExchange, error)
	GetCenterFeedBackFeeCount(p_platformId int32, p_serverId int32, p_playerId int64, p_code string, p_begin int64, p_end int64) (int32, error)
}

type feedBackFeeService struct {
	centerDs db.DBService
}

func (m *feedBackFeeService) GetGameFeedBackFeeList(p_dblink gmdb.GameDbLink, p_serverId int, p_playerId int64, p_code string, p_begin int64, p_end int64, p_index int32) ([]*feemodel.GamePlayerFeedBackFeeExchange, error) {
	rst := make([]*feemodel.GamePlayerFeedBackFeeExchange, 0)
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	limit := constant.DefaultPageSize
	where := "deleteTime=0 and serverId=?"
	if p_playerId > 0 {
		where += fmt.Sprintf(" and playerId=%d", p_playerId)
	}
	if p_begin > 0 {
		where += fmt.Sprintf(" and createTime>=%d", p_begin)
	}
	if p_end > 0 {
		where += fmt.Sprintf(" and createTime < %d", p_end)
	}
	if len(p_code) > 0 {
		where += fmt.Sprintf(" and code='%s'", p_code)
	}
	exdb := db.DB().Where(where, p_serverId).Offset(offect).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *feedBackFeeService) GetGameFeedBackFeeCount(p_dblink gmdb.GameDbLink, p_serverId int, p_playerId int64, p_code string, p_begin int64, p_end int64) (int32, error) {
	rst := int32(0)
	db := m.getdb(p_dblink)
	if db == nil {
		return rst, fmt.Errorf("DB服务为空")
	}
	where := "deleteTime=0 and serverId=?"
	if p_playerId > 0 {
		where += fmt.Sprintf(" and playerId=%d", p_playerId)
	}
	if p_begin > 0 {
		where += fmt.Sprintf(" and createTime>=%d", p_begin)
	}
	if p_end > 0 {
		where += fmt.Sprintf(" and createTime < %d", p_end)
	}
	if len(p_code) > 0 {
		where += fmt.Sprintf(" and code='%s'", p_code)
	}
	exdb := db.DB().Table("t_feedback_exchange").Where(where, p_serverId).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *feedBackFeeService) GetCenterFeedBackFeeList(p_platformId int32, p_serverId int32, p_playerId int64, p_code string, p_begin int64, p_end int64, p_index int32) ([]*feemodel.CenterPlayerFeedBackFeeExchange, error) {
	rst := make([]*feemodel.CenterPlayerFeedBackFeeExchange, 0)
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	limit := constant.DefaultPageSize

	where := "deleteTime=0"
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId=%d", p_serverId)
	}
	if p_playerId > 0 {
		where += fmt.Sprintf(" and playerId=%d", p_playerId)
	}
	if p_platformId > 0 {
		where += fmt.Sprintf(" and platform=%d", p_platformId)
	}
	if p_begin > 0 {
		where += fmt.Sprintf(" and createTime>=%d", p_begin)
	}
	if p_end > 0 {
		where += fmt.Sprintf(" and createTime < %d", p_end)
	}
	if len(p_code) > 0 {
		where += fmt.Sprintf(" and code='%s'", p_code)
	}
	exdb := m.centerDs.DB().Where(where).Offset(offect).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *feedBackFeeService) GetCenterFeedBackFeeCount(p_platformId int32, p_serverId int32, p_playerId int64, p_code string, p_begin int64, p_end int64) (int32, error) {
	rst := int32(0)
	where := "deleteTime=0"
	if p_serverId > 0 {
		where += fmt.Sprintf(" and serverId=%d", p_serverId)
	}
	if p_playerId > 0 {
		where += fmt.Sprintf(" and playerId=%d", p_playerId)
	}
	if p_platformId > 0 {
		where += fmt.Sprintf(" and platform=%d", p_platformId)
	}
	if p_begin > 0 {
		where += fmt.Sprintf(" and createTime>=%d", p_begin)
	}
	if p_end > 0 {
		where += fmt.Sprintf(" and createTime < %d", p_end)
	}
	if len(p_code) > 0 {
		where += fmt.Sprintf(" and code='%s'", p_code)
	}
	exdb := m.centerDs.DB().Table("t_feedbackfee_exchange").Where(where).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *feedBackFeeService) getdb(p_dblink gmdb.GameDbLink) gmdb.DBService {
	return gmdb.GetDb(p_dblink)
}

func NewFeedBackFeeService(cs db.DBService) FeedBackFeeService {
	rst := &feedBackFeeService{
		centerDs: cs,
	}
	return rst
}

type contextKey string

const (
	feeBackFeeKey = contextKey("feeBackFee")
)

func WithFeedBackFeeService(ctx context.Context, ls FeedBackFeeService) context.Context {
	return context.WithValue(ctx, feeBackFeeKey, ls)
}

func FeedBackFeeServiceInContext(ctx context.Context) FeedBackFeeService {
	us, ok := ctx.Value(feeBackFeeKey).(FeedBackFeeService)
	if !ok {
		return nil
	}
	return us
}

func SetupFeedBackFeeServiceHandler(ls FeedBackFeeService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithFeedBackFeeService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
