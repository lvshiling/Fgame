package service

import (
	"context"
	"fgame/fgame/core/db"
	constant "fgame/fgame/gm/gamegm/constant"
	"fgame/fgame/gm/gamegm/gm/mglog/model"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type PlayerStatsService interface {
	GetPlayerStats(p_startTime int64, p_endTime int64, p_index int) ([]*model.PlayerStatEntity, error)
	GetPlayerStatsCount(p_startTime int64, p_endTime int64) (int, error)
}

type playerStatsService struct {
	ds db.DBService
}

func (m *playerStatsService) GetPlayerStats(p_startTime int64, p_endTime int64, p_index int) ([]*model.PlayerStatEntity, error) {
	rst := make([]*model.PlayerStatEntity, 0)
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize

	exdb := m.ds.DB().Where("beginTime >=? and beginTime <=?", p_startTime, p_endTime).Offset(offset).Limit(limit).Order("beginTime asc,statCount desc").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}

	return rst, nil
}

func (m *playerStatsService) GetPlayerStatsCount(p_startTime int64, p_endTime int64) (int, error) {
	rst := 0
	exdb := m.ds.DB().Table("t_player_stats").Where("beginTime >=? and beginTime <=?", p_startTime, p_endTime).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return 0, exdb.Error
	}
	return rst, nil
}

func NewplayerStatsService(p_ds db.DBService) PlayerStatsService {
	rst := &playerStatsService{
		ds: p_ds,
	}
	return rst
}

type contextKey string

const (
	playerStatsServiceKey = contextKey("PlayerStatsService")
)

func WithPlayerStatsService(ctx context.Context, ls PlayerStatsService) context.Context {
	return context.WithValue(ctx, playerStatsServiceKey, ls)
}

func PlayerStatsServiceInContext(ctx context.Context) PlayerStatsService {
	us, ok := ctx.Value(playerStatsServiceKey).(PlayerStatsService)
	if !ok {
		return nil
	}
	return us
}

func SetupPlayerStatsServiceHandler(ls PlayerStatsService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithPlayerStatsService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
