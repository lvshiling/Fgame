package service

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/center/set/model"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type CenterSetService interface {
	SaveClientVersion(iosVersion string, androidVersion string) error
	GetClientVersion() (rst *model.ClientVersionEntity, err error)

	SavePlatformServerConfig(tradeServerIp string) error
	GetPlatformServerConfig() (rst *model.PlatformServerConfigEntity, err error)
}

type centerSetService struct {
	ds gmdb.DBService
}

func (m *centerSetService) SaveClientVersion(iosVersion string, androidVersion string) error {
	rstInfo := &model.ClientVersionEntity{}
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.ds.DB().First(rstInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if rstInfo.Id == 0 {
		rstInfo.CreateTime = now
	}
	rstInfo.UpdateTime = now
	rstInfo.AndroidVersion = androidVersion
	rstInfo.IosVersion = iosVersion
	exdb = m.ds.DB().Save(rstInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *centerSetService) GetClientVersion() (*model.ClientVersionEntity, error) {
	rstInfo := &model.ClientVersionEntity{}
	exdb := m.ds.DB().First(rstInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rstInfo, nil
}

func (m *centerSetService) SavePlatformServerConfig(tradeServerIp string) error {
	rstInfo := &model.PlatformServerConfigEntity{}
	now := timeutils.TimeToMillisecond(time.Now())
	exdb := m.ds.DB().First(rstInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	if rstInfo.Id == 0 {
		rstInfo.CreateTime = now
	}
	rstInfo.UpdateTime = now
	rstInfo.TradeServerIp = tradeServerIp
	exdb = m.ds.DB().Save(rstInfo)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return exdb.Error
	}
	return nil
}

func (m *centerSetService) GetPlatformServerConfig() (rst *model.PlatformServerConfigEntity, err error) {
	rst = &model.PlatformServerConfigEntity{}
	exdb := m.ds.DB().First(rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		err = exdb.Error
		return
	}
	return
}

func NewCenterSetService(p_db gmdb.DBService) CenterSetService {
	rst := &centerSetService{
		ds: p_db,
	}
	return rst
}

type contextKey string

const (
	centerSetServiceKey = contextKey("CenterSetService")
)

func WithCenterSetService(ctx context.Context, ls CenterSetService) context.Context {
	return context.WithValue(ctx, centerSetServiceKey, ls)
}

func CenterSetServiceInContext(ctx context.Context) CenterSetService {
	us, ok := ctx.Value(centerSetServiceKey).(CenterSetService)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterSetServiceHandler(ls CenterSetService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterSetService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
