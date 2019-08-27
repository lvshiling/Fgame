package service

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/manage/serverdaily/model"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ServerDailyStatService interface {
	SaveStatArray(array []*model.ServerDailyStats) error
	GetStatListByDate(begin int64) ([]*model.ServerDailyStats, error)
	GetStatCountByDate(begin int64) (int32, error)
	GetMinDate() (int64, error)
	GetPlatformStatListByDate(platformId int, begin int64, end int64) ([]*model.ServerDailyStats, error)
}

type serverDailyStatService struct {
	gmds gmdb.DBService
}

type tempServerDailyMinDate struct {
	MinDate int64 `gorm:"column:minDate"`
}

func (m *serverDailyStatService) SaveStatArray(array []*model.ServerDailyStats) error {
	execDs := m.gmds.DB().Begin()
	for _, value := range array {
		exdb := execDs.Save(value)
		if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
			execDs.Rollback()
			return exdb.Error
		}
	}
	execDs.Commit()
	return nil
}

func (m *serverDailyStatService) GetStatListByDate(begin int64) ([]*model.ServerDailyStats, error) {
	rst := make([]*model.ServerDailyStats, 0)
	exdb := m.gmds.DB().Where("curDate = ?", begin).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *serverDailyStatService) GetStatCountByDate(begin int64) (int32, error) {
	totalCount := int32(0)
	exdb := m.gmds.DB().Table("t_server_daily_stats").Where("curDate = ?", begin).Count(&totalCount)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return totalCount, nil
}

func (m *serverDailyStatService) GetPlatformStatListByDate(platformId int, begin int64, end int64) ([]*model.ServerDailyStats, error) {
	rst := make([]*model.ServerDailyStats, 0)
	exdb := m.gmds.DB().Where("platformId = ? and curDate >= ? and curDate <= ?", platformId, begin, end).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

var (
	minDateSql = `SELECT MIN(curDate) AS minDate FROM t_server_daily_stats`
)

func (m *serverDailyStatService) GetMinDate() (int64, error) {
	rst := &tempServerDailyMinDate{}
	exdb := m.gmds.DB().Raw(minDateSql).Scan(rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return rst.MinDate, nil
}

func NewServerDailyStatService(ds gmdb.DBService) ServerDailyStatService {
	rst := &serverDailyStatService{
		gmds: ds,
	}
	return rst
}

type contextKey string

const (
	serverDailyStatKey = contextKey("serverDailyStat")
)

func WithServerDailyStat(ctx context.Context, ls ServerDailyStatService) context.Context {
	return context.WithValue(ctx, serverDailyStatKey, ls)
}

func ServerDailyStatInContext(ctx context.Context) ServerDailyStatService {
	us, ok := ctx.Value(serverDailyStatKey).(ServerDailyStatService)
	if !ok {
		return nil
	}
	return us
}

func SetupserverDailyStatHandler(ls ServerDailyStatService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithServerDailyStat(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
