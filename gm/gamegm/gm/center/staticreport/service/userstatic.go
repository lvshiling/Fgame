package service

import (
	"context"
	"fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	"fgame/fgame/gm/gamegm/gm/center/staticreport/entity"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type IPlayerStatic interface {
	GetPlayerRegister(p_dblink gmdb.GameDbLink, serverId int32, p_start int64, p_end int64) (*entity.RegisterStaticPlayer, error)
	GetServerOnLine(platformId int, serverId int, p_start int64, p_end int64) ([]*entity.ServerOnLineStatic, error)

	GetTradeItemList(p_platformId int32, p_serverId int32, p_start int64, p_end int64, p_tradeId int64, p_playerId int64, p_level int32, p_state int32, p_page int32) ([]*entity.TradeItem, error)
	GetTradeItemCount(p_platformId int32, p_serverId int32, p_start int64, p_end int64, p_tradeId int64, p_playerId int64, p_level int32, p_state int32) (int32, error)
}

type playerStatic struct {
	db       gmdb.DBService
	centerDb gmdb.DBService
}

var (
	registerSql = `SELECT COUNT(1) AS totalPlayerCount
	,SUM(CASE WHEN createTime >= ? THEN 1 ELSE 0 END) AS todayRegisterCount
FROM t_player
WHERE
	createTime < ? and originServerId=?`
)

func (m *playerStatic) GetPlayerRegister(p_dblink gmdb.GameDbLink, serverId int32, p_start int64, p_end int64) (*entity.RegisterStaticPlayer, error) {
	rst := &entity.RegisterStaticPlayer{}
	db := m.getdb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	exdb := db.DB().Raw(registerSql, p_start, p_end, serverId).Scan(rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

var (
	onLineSql = `SELECT
	A.onLineDate
	,SUM(CASE WHEN B.onLineIndex = 0 THEN 1 ELSE 0 END) AS num0
	,SUM(CASE WHEN B.onLineIndex = 1 THEN 1 ELSE 0 END) AS num1
	,SUM(CASE WHEN B.onLineIndex = 2 THEN 1 ELSE 0 END) AS num2
	,SUM(CASE WHEN B.onLineIndex = 3 THEN 1 ELSE 0 END) AS num3
	,SUM(CASE WHEN B.onLineIndex = 4 THEN 1 ELSE 0 END) AS num4
	,SUM(CASE WHEN B.onLineIndex = 5 THEN 1 ELSE 0 END) AS num5
	,SUM(CASE WHEN B.onLineIndex = 6 THEN 1 ELSE 0 END) AS num6
	,SUM(CASE WHEN B.onLineIndex = 7 THEN 1 ELSE 0 END) AS num7
	,SUM(CASE WHEN B.onLineIndex = 14 THEN 1 ELSE 0 END) AS num14
	,SUM(CASE WHEN B.onLineIndex = 30 THEN 1 ELSE 0 END) AS num30
FROM t_server_online A
LEFT JOIN t_server_online B
ON A.playerId = B.playerId
WHERE
	A.onLineDate >= ? and A.onLineDate <= ? and A.platformId = ? AND A.serverId = ? AND A.onLineIndex = 0
GROUP BY A.onLineDate
ORDER BY A.onLineDate ASC`
)

func (m *playerStatic) GetServerOnLine(p_platformId int, p_serverId int, p_start int64, p_end int64) ([]*entity.ServerOnLineStatic, error) {
	rst := make([]*entity.ServerOnLineStatic, 0)
	exdb := m.db.DB().Raw(onLineSql, p_start, p_end, p_platformId, p_serverId).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return rst, exdb.Error
	}
	return rst, nil
}

func (m *playerStatic) GetTradeItemList(p_platformId int32, p_serverId int32, p_start int64, p_end int64, p_tradeId int64, p_playerId int64, p_level int32, p_state int32, p_index int32) ([]*entity.TradeItem, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize

	rst := make([]*entity.TradeItem, 0)
	paramList := make([]interface{}, 0)
	where := "platform = ?"
	paramList = append(paramList, p_platformId)
	if p_serverId > 0 {
		where += " and serverId=?"
		paramList = append(paramList, p_serverId)
	}
	if p_start > 0 {
		where += " and createTime >= ?"
		paramList = append(paramList, p_start)
	}
	if p_end > 0 {
		where += " and createTime < ?"
		paramList = append(paramList, p_end)
	}
	if p_tradeId > 0 {
		where += " and tradeId = ?"
		paramList = append(paramList, p_tradeId)
	}
	if p_playerId > 0 {
		where += " and playerId = ?"
		paramList = append(paramList, p_playerId)
	}
	if p_level > 0 {
		where += " and level = ?"
		paramList = append(paramList, p_level)
	}
	if p_state > -1 {
		where += " and status = ?"
		paramList = append(paramList, p_state)
	}

	exdb := m.centerDb.DB().Where(where, paramList...).Offset(offset).Limit(limit).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *playerStatic) GetTradeItemCount(p_platformId int32, p_serverId int32, p_start int64, p_end int64, p_tradeId int64, p_playerId int64, p_level int32, p_state int32) (int32, error) {
	rst := int32(0)
	paramList := make([]interface{}, 0)
	where := "platform = ?"
	paramList = append(paramList, p_platformId)
	if p_serverId > 0 {
		where += " and serverId=?"
		paramList = append(paramList, p_serverId)
	}
	if p_start > 0 {
		where += " and createTime >= ?"
		paramList = append(paramList, p_start)
	}
	if p_end > 0 {
		where += " and createTime < ?"
		paramList = append(paramList, p_end)
	}
	if p_tradeId > 0 {
		where += " and tradeId = ?"
		paramList = append(paramList, p_tradeId)
	}
	if p_playerId > 0 {
		where += " and playerId = ?"
		paramList = append(paramList, p_playerId)
	}
	if p_level > 0 {
		where += " and level = ?"
		paramList = append(paramList, p_level)
	}
	if p_state > -1 {
		where += " and status = ?"
		paramList = append(paramList, p_state)
	}

	exdb := m.centerDb.DB().Table("t_trade_item").Where(where, paramList...).Count(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return int32(0), exdb.Error
	}
	return rst, nil
}

func (m *playerStatic) getdb(p_dblink gmdb.GameDbLink) gmdb.DBService {
	return gmdb.GetDb(p_dblink)
}

func NewPlayerStatic(db gmdb.DBService, centerdb gmdb.DBService) IPlayerStatic {
	rst := &playerStatic{}
	rst.db = db
	rst.centerDb = centerdb
	return rst
}

const (
	playerStaticKey = contextKey("PlayerStatic")
)

func WithPlayerStatic(ctx context.Context, ls IPlayerStatic) context.Context {
	return context.WithValue(ctx, playerStaticKey, ls)
}

func PlayerStaticInContext(ctx context.Context) IPlayerStatic {
	us, ok := ctx.Value(playerStaticKey).(IPlayerStatic)
	if !ok {
		return nil
	}
	return us
}

func SetupPlayerStaticHandler(ls IPlayerStatic) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithPlayerStatic(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
