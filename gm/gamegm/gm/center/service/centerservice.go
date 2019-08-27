package service

import (
	"context"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmcentermodel "fgame/fgame/gm/gamegm/gm/center/model"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ICenterService interface {
	GetAllCenterPlatform() ([]*gmcentermodel.CenterPlatformInfo, error)
	GetGroupByPlatForm(p_platid int64) ([]*gmcentermodel.CenterGroupInfo, error)
	GetServerByPlatForm(p_platid int64) ([]*gmcentermodel.CenterServer, error)
	GetAllServerByPlatForm(p_platid int64) ([]*gmcentermodel.CenterServer, error)
}

type centerService struct {
	db gmdb.DBService
}

var (
	getAllCenterPlatformSql = `SELECT * FROM t_platform WHERE deleteTime = 0`
)

func (m *centerService) GetAllCenterPlatform() ([]*gmcentermodel.CenterPlatformInfo, error) {
	rst := make([]*gmcentermodel.CenterPlatformInfo, 0)

	dberr := m.db.DB().Raw(getAllCenterPlatformSql).Scan(&rst)
	if dberr.Error != nil && dberr.Error != gorm.ErrRecordNotFound {
		return nil, dberr.Error
	}
	return rst, nil
}

var (
	getGroupByPlatFormSql = `SELECT
	A.id
	,A.serverType
	,A.serverId
	,A.platform
	,CONCAT(
		CONCAT(CASE WHEN B.id > 0 THEN CONCAT('---[(',C.serverId,CONCAT(')'),C.name,']') ELSE '' END, CONCAT('(',A.serverId,')',A.name))
		, '[开',to_days(current_timestamp())-to_days(from_unixtime(A.startTime/1000))+1,'天]'
		,CASE WHEN B.id > 0 THEN CONCAT('[合',to_days(current_timestamp())-to_days(from_unixtime(B.mergeTime/1000))+1,'天]') ELSE '' END) 
		AS name
	,A.startTime
	,A.updateTime
	,A.createTime
	,A.deleteTime
	,A.serverIp
	,A.serverPort
	,A.serverRemoteIp
	,A.serverRemotePort
	,A.serverDBIp
	,A.serverDBPort
	,A.serverDBName
	,A.serverDBUser
	,A.serverDBPassword
	,A.serverTag
	,A.serverStatus
FROM
	t_server A
	LEFT JOIN t_merge_record B
	ON A.platform = B.platform and A.serverId = B.fromServerId AND A.serverType=0
	LEFT JOIN t_server C
	ON B.platform=C.platform AND B.finalServerId = C.serverId  AND C.serverType=0
WHERE
	A.deleteTime = 0 AND A.serverType=0 AND A.platform=?
ORDER BY A.platform,IFNULL(C.serverId,A.serverId),A.serverId ASC`
)

func (m *centerService) GetGroupByPlatForm(p_platid int64) ([]*gmcentermodel.CenterGroupInfo, error) {
	rst := make([]*gmcentermodel.CenterGroupInfo, 0)

	dberr := m.db.DB().Raw(getGroupByPlatFormSql, p_platid).Scan(&rst)
	if dberr.Error != nil && dberr.Error != gorm.ErrRecordNotFound {
		return nil, dberr.Error
	}
	return rst, nil
}

var (
	getServerByPlatFormSql = `SELECT
	A.id
	,A.serverType
	,A.serverId
	,A.platform
	,CONCAT(
		CONCAT(CASE WHEN B.id > 0 THEN CONCAT('---[(',C.serverId,CONCAT(')'),C.name,']') ELSE '' END, CONCAT('(',A.serverId,')',A.name))
		, '[开',to_days(current_timestamp())-to_days(from_unixtime(A.startTime/1000))+1,'天]'
		,CASE WHEN B.id > 0 THEN CONCAT('[合',to_days(current_timestamp())-to_days(from_unixtime(B.mergeTime/1000))+1,'天]') ELSE '' END) 
		AS name
	,A.startTime
	,A.updateTime
	,A.createTime
	,A.deleteTime
	,A.serverIp
	,A.serverPort
	,A.serverRemoteIp
	,A.serverRemotePort
	,A.serverDBIp
	,A.serverDBPort
	,A.serverDBName
	,A.serverDBUser
	,A.serverDBPassword
	,A.serverTag
	,A.serverStatus
FROM
	t_server A
	LEFT JOIN t_merge_record B
	ON A.platform = B.platform and A.serverId = B.fromServerId AND A.serverType=0
	LEFT JOIN t_server C
	ON B.platform=C.platform AND B.finalServerId = C.serverId  AND C.serverType=0
WHERE
	A.deleteTime = 0 AND A.serverType=0 and A.platform=?
ORDER BY A.platform,IFNULL(C.serverId,A.serverId),A.serverId ASC`
)

func (m *centerService) GetServerByPlatForm(p_platid int64) ([]*gmcentermodel.CenterServer, error) {
	rst := make([]*gmcentermodel.CenterServer, 0)

	dberr := m.db.DB().Raw(getServerByPlatFormSql, p_platid).Scan(&rst)
	if dberr.Error != nil && dberr.Error != gorm.ErrRecordNotFound {
		return nil, dberr.Error
	}
	return rst, nil
}

var (
	getAllServerByPlatFormSql = `SELECT
	A.id
	,A.serverType
	,A.serverId
	,A.platform
	,CONCAT(
		CONCAT(CASE WHEN B.id > 0 THEN CONCAT('---[(',C.serverId,CONCAT(')'),C.name,']') ELSE '' END, CONCAT('(',A.serverId,')',A.name))
		, '[开',to_days(current_timestamp())-to_days(from_unixtime(A.startTime/1000))+1,'天]'
		,CASE WHEN B.id > 0 THEN CONCAT('[合',to_days(current_timestamp())-to_days(from_unixtime(B.mergeTime/1000))+1,'天]') ELSE '' END) 
		AS name
	,A.startTime
	,A.updateTime
	,A.createTime
	,A.deleteTime
	,A.serverIp
	,A.serverPort
	,A.serverRemoteIp
	,A.serverRemotePort
	,A.serverDBIp
	,A.serverDBPort
	,A.serverDBName
	,A.serverDBUser
	,A.serverDBPassword
	,A.serverTag
	,A.serverStatus
FROM
	t_server A
	LEFT JOIN t_merge_record B
	ON A.platform = B.platform and A.serverId = B.fromServerId AND A.serverType=0
	LEFT JOIN t_server C
	ON B.platform=C.platform AND B.finalServerId = C.serverId  AND C.serverType=0
WHERE
	A.deleteTime = 0  AND A.serverType=0 and A.platform=?
ORDER BY A.platform,IFNULL(C.serverId,A.serverId),A.serverId ASC`
)

func (m *centerService) GetAllServerByPlatForm(p_platid int64) ([]*gmcentermodel.CenterServer, error) {
	rst := make([]*gmcentermodel.CenterServer, 0)

	dberr := m.db.DB().Raw(getAllServerByPlatFormSql, p_platid).Scan(&rst)
	if dberr.Error != nil && dberr.Error != gorm.ErrRecordNotFound {
		return nil, dberr.Error
	}
	return rst, nil
}

func NewCenterPlatformService(p_db gmdb.DBService) ICenterService {
	rst := &centerService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	platformServiceKey = contextKey("CenterPlatformService")
)

func WithCenterPlatformService(ctx context.Context, ls ICenterService) context.Context {
	return context.WithValue(ctx, platformServiceKey, ls)
}

func CenterPlatformServiceInContext(ctx context.Context) ICenterService {
	us, ok := ctx.Value(platformServiceKey).(ICenterService)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterPlatformServiceHandler(ls ICenterService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterPlatformService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
