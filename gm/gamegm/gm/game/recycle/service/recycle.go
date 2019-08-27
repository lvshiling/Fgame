package service

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	recyclemodel "fgame/fgame/gm/gamegm/gm/game/recycle/model"

	"github.com/codegangsta/negroni"
)

type RecycleService interface {
	GetServerRecycleModel(p_dblink gmdb.GameDbLink, p_serverId int) (*recyclemodel.TradeRecycleEntity, error)
}

type recycleService struct {
}

func (m *recycleService) GetServerRecycleModel(p_dblink gmdb.GameDbLink, p_serverId int) (*recyclemodel.TradeRecycleEntity, error) {
	info := &recyclemodel.TradeRecycleEntity{}
	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	exdb := db.DB().Where("serverId = ? and deleteTime=0", p_serverId).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *recycleService) getdb(p_dblink gmdb.GameDbLink) gmdb.DBService {
	return gmdb.GetDb(p_dblink)
}

func NewRecycleService() RecycleService {
	rst := &recycleService{}
	return rst
}

type contextKey string

const (
	recycleServiceKey = contextKey("recycleService")
)

func WithRecycleService(ctx context.Context, ls RecycleService) context.Context {
	return context.WithValue(ctx, recycleServiceKey, ls)
}

func RecycleServiceInContext(ctx context.Context) RecycleService {
	us, ok := ctx.Value(recycleServiceKey).(RecycleService)
	if !ok {
		return nil
	}
	return us
}

func SetupRecycleServiceHandler(ls RecycleService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithRecycleService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
