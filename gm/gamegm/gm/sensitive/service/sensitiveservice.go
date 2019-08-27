package service

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	sensitivemodel "fgame/fgame/gm/gamegm/gm/sensitive/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ISensitiveService interface {
	Save(p_userid int64, p_sensitive string) error
	GetSensitive(p_userid int64) (*sensitivemodel.UserSensitive, error)
}

type sensitiveService struct {
	db gmdb.DBService
}

func (m *sensitiveService) Save(p_userid int64, p_sensitive string) error {
	info := &sensitivemodel.UserSensitive{}
	exdb := m.db.DB().Where("userId = ? and deleteTime=0", p_userid).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	info.UserId = p_userid
	info.Content = p_sensitive
	info.UpdateTime = timeutils.TimeToMillisecond(time.Now())
	exdb = m.db.DB().Save(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *sensitiveService) GetSensitive(p_userid int64) (*sensitivemodel.UserSensitive, error) {
	info := &sensitivemodel.UserSensitive{}
	exdb := m.db.DB().Where("userId = ? and deleteTime=0", p_userid).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func NewSensitiveService(p_db gmdb.DBService) ISensitiveService {
	rst := &sensitiveService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	sensitiveServiceKey = contextKey("SensitiveService")
)

func WithSensitiveService(ctx context.Context, ls ISensitiveService) context.Context {
	return context.WithValue(ctx, sensitiveServiceKey, ls)
}

func SensitiveServiceInContext(ctx context.Context) ISensitiveService {
	us, ok := ctx.Value(sensitiveServiceKey).(ISensitiveService)
	if !ok {
		return nil
	}
	return us
}

func SetupSensitiveServiceHandler(ls ISensitiveService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithSensitiveService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
