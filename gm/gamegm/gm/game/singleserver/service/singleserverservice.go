package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/game/singleserver/model"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

type SingleServerService interface {
	GetFirstCharge(p_dblink gmdb.GameDbLink, p_serverId int) (int64, error)
	GetSingleServerDoubleChargeRegisterLogList(p_dblink gmdb.GameDbLink, p_serverId int, p_pageIndex int) ([]*model.NewFirstChargeLog, error)
	GetSingleServerDoubleChargeRegisterLogCount(p_dblink gmdb.GameDbLink, p_serverId int) (int32, error)
}

type singleServerService struct {
}

func (m *singleServerService) GetFirstCharge(p_dblink gmdb.GameDbLink, p_serverId int) (int64, error) {
	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := &model.NewFirstCharge{}
	exdb := db.DB().Where("serverId = ? and deleteTime=0", p_serverId).First(rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst.StartTime, nil
}

func (m *singleServerService) GetSingleServerDoubleChargeRegisterLogList(p_dblink gmdb.GameDbLink, p_serverId int, p_index int) ([]*model.NewFirstChargeLog, error) {
	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	rst := make([]*model.NewFirstChargeLog, 0)
	errdb := db.DB().Where("serverId = ? and deleteTime=0", p_serverId).Order("createTime desc").Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if errdb.Error != nil {
		return nil, errdb.Error
	}
	return rst, nil
}

func (m *singleServerService) GetSingleServerDoubleChargeRegisterLogCount(p_dblink gmdb.GameDbLink, p_serverId int) (int32, error) {
	db := gmdb.GetDb(p_dblink)
	rst := 0
	errdb := db.DB().Table("t_new_first_charge_log").Where("serverId = ? and deleteTime=0", p_serverId).Count(&rst)
	if errdb.Error != nil {
		return 0, errdb.Error
	}
	return int32(rst), nil
}

func NewSingleServerService() SingleServerService {
	rst := &singleServerService{}
	return rst
}

type contextKey string

const (
	singleServerServiceKey = contextKey("singleServerService")
)

func WithSingleServerService(ctx context.Context, ls SingleServerService) context.Context {
	return context.WithValue(ctx, singleServerServiceKey, ls)
}

func SingleServerServiceInContext(ctx context.Context) SingleServerService {
	us, ok := ctx.Value(singleServerServiceKey).(SingleServerService)
	if !ok {
		return nil
	}
	return us
}

func SetupSingleServerServiceHandler(ls SingleServerService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithSingleServerService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
