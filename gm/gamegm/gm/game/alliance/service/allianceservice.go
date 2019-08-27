package service

import (
	"context"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	almodel "fgame/fgame/gm/gamegm/gm/game/alliance/model"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
)

type IAllianceService interface {
	GetAllianceList(p_dblink gmdb.GameDbLink, p_serverId int, p_allianceName string, p_index int, p_col int, p_asc int) (rst []*almodel.AllianceQuery, err error)
	GetAllianceCount(p_dblink gmdb.GameDbLink, p_serverId int, p_allianceName string) (int, error)

	GetServerRegisterFlag(p_dblink gmdb.GameDbLink, p_serverId int) (int, error)
	GetServerRegisterLogList(p_dblink gmdb.GameDbLink, p_serverId int, p_pageIndex int) ([]*almodel.ServerRegisterSettingLog, error)
	GetServerRegisterLogCount(p_dblink gmdb.GameDbLink, p_serverId int) (int, error)
}

var (
	allianceColumnOrderMap = map[int]string{
		1: "id",
		2: "allianceName",
		3: "allianceLevel",
		4: "totalForce",
		5: "createTime",
		6: "playerCount",
		7: "notice",
	}
)

func getColName(p_id int) string {
	if value, ok := allianceColumnOrderMap[p_id]; ok {
		return value
	}
	return "id"
}

func getOrderBy(p_ordercol int, p_ordertype int) string {
	colName := getColName(p_ordercol)
	asc := "asc"
	if p_ordertype > 0 {
		asc = "desc"
	}
	return " " + colName + " " + asc
}

type allianceService struct {
}

func (m *allianceService) GetAllianceList(p_dblink gmdb.GameDbLink, p_serverId int, p_allianceName string, p_index int, p_col int, p_asc int) (rst []*almodel.AllianceQuery, err error) {

	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	orderStr := getOrderBy(p_col, p_asc)

	sql := `select A.id,A.name AS allianceName
	,A.level AS allianceLevel
	,A.totalForce
	,A.createTime
	,(SELECT COUNT(1) FROM t_player_alliance B WHERE A.id=B.allianceId AND B.deleteTime=0) AS playerCount
	,A.notice
from t_alliance A
WHERE A.serverId = ? and A.name like ? and A.deleteTime=0
order by ` + orderStr + fmt.Sprintf(" limit %d,%d", offect, constant.DefaultPageSize)
	exdb := db.DB().Raw(sql, p_serverId, "%"+p_allianceName+"%").Scan(&rst)
	if exdb.Error != nil {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *allianceService) GetAllianceCount(p_dblink gmdb.GameDbLink, p_serverId int, p_allianceName string) (int, error) {
	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := 0
	exdb := db.DB().Table("t_alliance").Where("deleteTime=0 and name like ? and serverId=?", "%"+p_allianceName+"%", p_serverId).Count(&rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst, nil
}

func (m *allianceService) GetServerRegisterFlag(p_dblink gmdb.GameDbLink, p_serverId int) (int, error) {
	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return 0, fmt.Errorf("DB服务为空")
	}
	rst := &almodel.ServerRegisterSetting{}
	exdb := db.DB().Where("serverId = ? and deleteTime=0", p_serverId).First(rst)
	if exdb.Error != nil {
		return 0, exdb.Error
	}
	return rst.Open, nil
}

func (m *allianceService) GetServerRegisterLogList(p_dblink gmdb.GameDbLink, p_serverId int, p_index int) ([]*almodel.ServerRegisterSettingLog, error) {
	db := gmdb.GetDb(p_dblink)
	if db == nil {
		return nil, fmt.Errorf("DB服务为空")
	}
	offect := (p_index - 1) * constant.DefaultPageSize
	if offect < 0 {
		offect = 0
	}
	rst := make([]*almodel.ServerRegisterSettingLog, 0)
	errdb := db.DB().Where("serverId = ? and deleteTime=0", p_serverId).Order("createTime desc").Offset(offect).Limit(constant.DefaultPageSize).Find(&rst)
	if errdb.Error != nil {
		return nil, errdb.Error
	}
	return rst, nil
}

func (m *allianceService) GetServerRegisterLogCount(p_dblink gmdb.GameDbLink, p_serverId int) (int, error) {
	db := gmdb.GetDb(p_dblink)
	rst := 0
	errdb := db.DB().Table("t_register_setting_log").Where("serverId = ? and deleteTime=0", p_serverId).Count(&rst)
	if errdb.Error != nil {
		return 0, errdb.Error
	}
	return rst, nil
}

func NewAllianceService() IAllianceService {
	rst := &allianceService{}
	return rst
}

type contextKey string

const (
	allianceServiceKey = contextKey("allianceService")
)

func WithAllianceService(ctx context.Context, ls IAllianceService) context.Context {
	return context.WithValue(ctx, allianceServiceKey, ls)
}

func AllianceServiceInContext(ctx context.Context) IAllianceService {
	us, ok := ctx.Value(allianceServiceKey).(IAllianceService)
	if !ok {
		return nil
	}
	return us
}

func SetupAllianceServiceHandler(ls IAllianceService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithAllianceService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
