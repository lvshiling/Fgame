package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/manage/notice/model"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

type INoticeService interface {
	AddNoticeInfo(p_info *model.NoticeInfo) error
	GetNoticeList(p_successFlag int, p_beginTime int64, p_endTime int64, p_pageindex int, p_platformList []int64) ([]*model.NoticeInfo, error)
	GetNoticeCount(p_successFlag int, p_beginTime int64, p_endTime int64, p_platformList []int64) (int, error)
}

type noticeService struct {
	db gmdb.DBService
}

func (m *noticeService) AddNoticeInfo(p_info *model.NoticeInfo) error {
	exdb := m.db.DB().Save(p_info)
	if exdb.Error != nil {
		return exdb.Error
	}
	return nil
}

func (m *noticeService) GetNoticeList(p_successFlag int, p_beginTime int64, p_endTime int64, p_pageindex int, p_platformList []int64) ([]*model.NoticeInfo, error) {
	rst := make([]*model.NoticeInfo, 0)
	offset := (p_pageindex - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	where := "deleteTime=0"
	if p_beginTime > 0 {
		where += fmt.Sprintf(" and createTime >= %d", p_beginTime)
	}
	if p_endTime > 0 {
		where += fmt.Sprintf(" and createTime < %d", p_endTime)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if p_successFlag > -1 {
		where += fmt.Sprintf(" and successFlag = %d", p_successFlag)
	}
	dberr := m.db.DB().Where(where).Offset(offset).Limit(limit).Find(&rst)
	if dberr.Error != nil {
		return nil, dberr.Error
	}
	return rst, nil
}

func (m *noticeService) GetNoticeCount(p_successFlag int, p_beginTime int64, p_endTime int64, p_platformList []int64) (int, error) {
	rst := 0

	where := "deleteTime=0"
	if p_beginTime > 0 {
		where += fmt.Sprintf(" and createTime >= %d", p_beginTime)
	}
	if p_endTime > 0 {
		where += fmt.Sprintf(" and createTime < %d", p_endTime)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and centerPlatformId IN (%s)", common.CombinInt64Array(p_platformList))
	}
	if p_successFlag > -1 {
		where += fmt.Sprintf(" and successFlag = %d", p_successFlag)
	}
	dberr := m.db.DB().Table("t_notice").Where(where).Count(&rst)
	if dberr.Error != nil {
		return 0, dberr.Error
	}
	return rst, nil
}

func NewNoticeInfo(p_db gmdb.DBService) INoticeService {
	rst := &noticeService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	noticeServiceKey = contextKey("noticeService")
)

func WithNoticeService(ctx context.Context, ls INoticeService) context.Context {
	return context.WithValue(ctx, noticeServiceKey, ls)
}

func NoticeServiceInContext(ctx context.Context) INoticeService {
	us, ok := ctx.Value(noticeServiceKey).(INoticeService)
	if !ok {
		return nil
	}
	return us
}

func SetupNoticeServiceHandler(ls INoticeService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithNoticeService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
