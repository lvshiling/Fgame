package service

import (
	"context"
	"fgame/fgame/gm/gamegm/common"
	constant "fgame/fgame/gm/gamegm/constant"
	gmdb "fgame/fgame/gm/gamegm/db"
	gmError "fgame/fgame/gm/gamegm/error"
	centerServermodel "fgame/fgame/gm/gamegm/gm/center/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	userremote "fgame/fgame/gm/gamegm/remote/service"
	"fmt"
	"net/http"
	"time"

	servermodel "fgame/fgame/gm/gamegm/gm/center/server/model"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

type ICenterServerService interface {
	AddCenterServer(p_serverType int, p_serverId int, p_platformId int64, p_name string, p_startTime int64, p_serverIp string, p_serverPort string, p_serverRemoteIp string, p_serverRemotePort string, p_serverDbIp string, p_serverDbPort string, p_serverDbName string, p_serverDBUser string, p_serverDBPassword string, p_serverTag int, p_serverStatus int, p_parentServerId int, p_preShow int) (int64, error)
	UpdateCenterServer(p_centerServerid int64, p_serverType int, p_serverId int, p_platformId int64, p_name string, p_startTime int64, p_serverIp string, p_serverPort string, p_serverRemoteIp string, p_serverRemotePort string, p_serverDbIp string, p_serverDbPort string, p_serverDbName string, p_serverDBUser string, p_serverDBPassword string, p_serverTag int, p_serverStatus int, p_parentServerId int, p_preShow int) error
	DeleteCenterServer(p_centerServerid int64) error
	UpdateParentServerId(p_id int, p_parentId int) error
	UpdateParentServerIdArray(p_id []int, p_parentId int) error
	UpdateJiaoYiZhanQuFu(p_id int, p_jiaoyizhanquFuId int) error
	UpdateJiaoYiZhanQuFuArray(p_id []int, p_jiaoyizhanquFuId int) error
	UpdatePingTaiFu(p_id int, p_pingTaiFuId int) error
	UpdatePingTaiFuArray(p_id []int, p_pingTaiFuId int) error
	UpdateChengZhanFu(p_id int, p_chengZhanFu int) error
	UpdateChengZhanFuArray(p_id []int, p_chengZhanFu int) error
	UpdateCenterServerName(p_centerServerId int64, p_name string, p_platformId int64) error

	GetCenterServerList(p_name string, p_platformId int, p_serverType int, p_index int) ([]*centerServermodel.CenterServer, error)
	GetCenterServerCount(p_name string, p_platformId int, p_serverType int) (int, error)
	GetAllCenterServerList() ([]*centerServermodel.CenterServer, error)
	GetAllCenterServerListArray(p_platformId []int64) ([]*centerServermodel.CenterServer, error)
	GetAllCenterServerListByType(p_centerplatformId int, p_serverType int) ([]*centerServermodel.CenterServer, error)
	QueryAllCenterServerListByType(p_centerplatformId int, p_serverType int) ([]*servermodel.CenterQueryServer, error)

	GetCenterServerListByPlatform(p_platformId int) ([]*centerServermodel.CenterServer, error)
	GetCenterServerListByPlatformArray(p_platformId []int) ([]*centerServermodel.CenterServer, error)
	GetCenterMainServerListByPlatform(p_platformId int) ([]*centerServermodel.CenterServer, error)

	GetCenterServerListByPlatformEnable(p_platformId int) ([]*centerServermodel.CenterServer, error)
	GetCenterServerListByPlatformArrayEnable(p_platformId []int) ([]*centerServermodel.CenterServer, error)
	GetAllCenterServerListEnable() ([]*centerServermodel.CenterServer, error)

	GetCenterServer(p_id int64) (*centerServermodel.CenterServer, error)
	GetCenterServerName(p_platformId int, p_serverId int) (string, error)
	GetCenterServerInfo(p_platformId int, p_serverId int) (*centerServermodel.CenterServer, error)

	GetCenterMergeRecord(p_platformId int, p_serverId int) ([]*servermodel.ServerMergeRecord, error)

	GetCenterServerZhanQu(p_platform int) ([]*centerServermodel.CenterServer, error)
	GetZhanQuServer(p_platform int, p_zhanQuServer []int) ([]*centerServermodel.CenterServer, error)

	RegisterDB(p_serverid gmdb.GameDbLink, p_host string, p_port string, p_dbName string, p_userName string, p_password string) error
	RegisterGrpc(p_serverId int32, p_host string, p_port string) error
}

type centerServerService struct {
	db gmdb.DBService
}

func (m *centerServerService) AddCenterServer(p_serverType int, p_serverId int, p_platformId int64, p_name string, p_startTime int64, p_serverIp string, p_serverPort string, p_serverRemoteIp string, p_serverRemotePort string, p_serverDbIp string, p_serverDbPort string, p_serverDbName string, p_serverDBUser string, p_serverDBPassword string, p_serverTag int, p_serverStatus int, p_parentServerId int, p_preShow int) (int64, error) {
	if len(p_name) == 0 || (p_serverType != 4 && p_serverType != 5 && p_platformId < 1) {
		return 0, gmError.GetError(gmError.ErrorCodeCenterServerEmpty)
	}

	exflag, err := m.existsCenterServer(p_name, p_platformId)
	if err != nil {
		return 0, err
	}
	if exflag {
		return 0, gmError.GetError(gmError.ErrorCodeCenterServerExist)
	}

	exflag, err = m.existsCenterServerId(p_serverId, p_platformId, p_serverType)
	if err != nil {
		return 0, err
	}
	if exflag {
		return 0, gmError.GetError(gmError.ErrorCodeCenterServerIdExist)
	}

	modelInfo := &centerServermodel.CenterServer{
		ServerName:       p_name,
		ServerType:       p_serverType,
		ServerId:         p_serverId,
		Platform:         p_platformId,
		StartTime:        p_startTime,
		ServerIp:         p_serverIp,
		ServerPort:       p_serverPort,
		ServerRemoteIp:   p_serverRemoteIp,
		ServerRemotePort: p_serverRemotePort,
		ServerDbIp:       p_serverDbIp,
		ServerDbPort:     p_serverDbPort,
		ServerDBName:     p_serverDbName,
		ServerDBUser:     p_serverDBUser,
		ServerDBPassword: p_serverDBPassword,
		ServerTag:        p_serverTag,
		ServerStatus:     p_serverStatus,
		ParentServerId:   p_parentServerId,
		PreShow:          p_preShow,
	}
	err = m.saveCenterServer(modelInfo)
	if err != nil {
		return 0, err
	}
	return modelInfo.Id, nil
}

func (m *centerServerService) UpdateCenterServer(p_centerServerid int64, p_serverType int, p_serverId int, p_platformId int64, p_name string, p_startTime int64, p_serverIp string, p_serverPort string, p_serverRemoteIp string, p_serverRemotePort string, p_serverDbIp string, p_serverDbPort string, p_serverDbName string, p_serverDBUser string, p_serverDBPassword string, p_serverTag int, p_serverStatus int, p_parentServerId int, p_preShow int) error {
	if len(p_name) == 0 {
		return gmError.GetError(gmError.ErrorCodeCenterServerEmpty)
	}

	exflag, err := m.existsCenterServerWithId(p_centerServerid, p_name, p_platformId)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeCenterServerExist)
	}

	exflag, err = m.existsCenterServerIdWithId(p_centerServerid, p_serverId, p_platformId, p_serverType)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeCenterServerIdExist)
	}
	modelInfo, err := m.GetCenterServer(p_centerServerid)
	if err != nil {
		return err
	}
	modelInfo.ServerType = p_serverType
	modelInfo.ServerId = p_serverId
	modelInfo.Platform = p_platformId
	modelInfo.ServerName = p_name
	modelInfo.StartTime = p_startTime
	modelInfo.ServerIp = p_serverIp
	modelInfo.ServerPort = p_serverPort
	modelInfo.ServerRemoteIp = p_serverRemoteIp
	modelInfo.ServerRemotePort = p_serverRemotePort
	modelInfo.ServerDbIp = p_serverDbIp
	modelInfo.ServerDbPort = p_serverDbPort
	modelInfo.ServerDBName = p_serverDbName
	modelInfo.ServerDBUser = p_serverDBUser
	modelInfo.ServerDBPassword = p_serverDBPassword
	modelInfo.ServerTag = p_serverTag
	modelInfo.ServerStatus = p_serverStatus
	modelInfo.ParentServerId = p_parentServerId
	modelInfo.PreShow = p_preShow

	// modelInfo = &centerServermodel.CenterServer{
	// 	Id:               p_centerServerid,
	// 	ServerType:       p_serverType,
	// 	ServerId:         p_serverId,
	// 	Platform:         p_platformId,
	// 	ServerName:       p_name,
	// 	StartTime:        p_startTime,
	// 	ServerIp:         p_serverIp,
	// 	ServerPort:       p_serverPort,
	// 	ServerRemoteIp:   p_serverRemoteIp,
	// 	ServerRemotePort: p_serverRemotePort,
	// 	ServerDbIp:       p_serverDbIp,
	// 	ServerDbPort:     p_serverDbPort,
	// 	ServerDBName:     p_serverDbName,
	// 	ServerDBUser:     p_serverDBUser,
	// 	ServerDBPassword: p_serverDBPassword,
	// 	ServerTag:        p_serverTag,
	// 	ServerStatus:     p_serverStatus,
	// 	ParentServerId:   p_parentServerId,
	// 	PreShow:          p_preShow,
	// }
	err = m.saveCenterServer(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *centerServerService) DeleteCenterServer(p_centerServerid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_server").Where("id = ?", p_centerServerid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"centerServerid": p_centerServerid,
			"error":          errdb.Error,
		}).Error("删除中心服务器失败")
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateParentServerId(p_id int, p_parentId int) error {
	errdb := m.db.DB().Table("t_server").Where("id = ?", p_id).Update("parentServerId", p_parentId)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateParentServerIdArray(p_id []int, p_parentId int) error {
	errdb := m.db.DB().Table("t_server").Where("id IN (?)", p_id).Update("parentServerId", p_parentId)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateJiaoYiZhanQuFu(p_id int, p_jiaoyizhanquFuId int) error {
	errdb := m.db.DB().Table("t_server").Where("id = ?", p_id).Update("jiaoYiZhanQuServerId", p_jiaoyizhanquFuId)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateJiaoYiZhanQuFuArray(p_id []int, p_jiaoyizhanquFuId int) error {
	errdb := m.db.DB().Table("t_server").Where("id IN (?)", p_id).Update("jiaoYiZhanQuServerId", p_jiaoyizhanquFuId)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdatePingTaiFu(p_id int, p_pingTaiFuId int) error {
	errdb := m.db.DB().Table("t_server").Where("id = ?", p_id).Update("pingTaiFuServerId", p_pingTaiFuId)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdatePingTaiFuArray(p_id []int, p_pingTaiFuId int) error {
	errdb := m.db.DB().Table("t_server").Where("id IN (?)", p_id).Update("pingTaiFuServerId", p_pingTaiFuId)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateChengZhanFu(p_id int, p_chengZhanFu int) error {
	errdb := m.db.DB().Table("t_server").Where("id = ?", p_id).Update("chengZhanServerId", p_chengZhanFu)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateChengZhanFuArray(p_id []int, p_chengZhanFu int) error {
	errdb := m.db.DB().Table("t_server").Where("id IN (?)", p_id).Update("chengZhanServerId", p_chengZhanFu)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) UpdateCenterServerName(p_id int64, p_name string, p_platformId int64) error {
	if len(p_name) == 0 {
		return gmError.GetError(gmError.ErrorCodeCenterServerEmpty)
	}
	exflag, err := m.existsCenterServerWithId(p_id, p_name, p_platformId)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeCenterServerExist)
	}

	errdb := m.db.DB().Table("t_server").Where("id = ?", p_id).Update("name", p_name)
	if errdb.Error != nil {
		return errdb.Error
	}
	return nil
}

func (m *centerServerService) GetCenterServerList(p_name string, p_platformId int, p_serverType int, p_index int) ([]*centerServermodel.CenterServer, error) {
	return m.getCenterServerList(p_name, p_platformId, p_serverType, p_index)
}

func (m *centerServerService) GetAllCenterServerList() ([]*centerServermodel.CenterServer, error) {
	return m.getAllCenterServerList()
}

func (m *centerServerService) GetAllCenterServerListArray(p_platformId []int64) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	where := ""
	if len(p_platformId) > 0 {
		where += fmt.Sprintf(" and platform IN (%s)", common.CombinInt64Array(p_platformId))
	}
	exerr := m.db.DB().Where("deleteTime=0 and serverType=0 " + where).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL中心服务器列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerServerService) GetAllCenterServerListByType(p_centerplatformId int, p_serverType int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	where := ""
	if p_centerplatformId > 0 {
		where += fmt.Sprintf(" and platform = %d", p_centerplatformId)
	}
	exerr := m.db.DB().Where("deleteTime=0 and serverType=?"+where, p_serverType).Order("platform asc,serverType asc,serverId asc").Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL中心服务器列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

// var (
// 	getAllCenterServerListByTypeQuerySql = `SELECT
// 	A.id
// 	,A.serverType
// 	,A.serverId
// 	,A.platform
// 	,A.name
// 	,A.parentServerId
// 	,A.preShow
// 	,A.jiaoYiZhanQuServerId
// 	,A.pingTaiFuServerId
// 	,A.chengZhanServerId
// 	,IFNULL(B.finalServerId,A.serverId) AS finnalServerId
// FROM
// 	t_server A
// 	LEFT JOIN t_merge_record B
// 	ON A.platform = B.platform AND A.serverId = B.fromServerId AND B.deleteTime=0
// WHERE
// 	A.deleteTime=0 AND A.serverType=?
// ORDER BY A.platform,A.serverId`
// )

var (
	getAllCenterServerListByTypeQuerySql = `SELECT
	A.id
	,A.serverType
	,A.serverId
	,A.platform
	,A.name
	,A.parentServerId
	,A.preShow
	,A.jiaoYiZhanQuServerId
	,A.pingTaiFuServerId
	,A.chengZhanServerId
	,IFNULL(B.finalServerId,A.serverId) AS finnalServerId
FROM
	t_server A
	LEFT JOIN t_merge_record B
	ON A.platform = B.platform AND A.serverId = B.fromServerId AND B.deleteTime=0
WHERE
	A.deleteTime=0 AND A.serverType=?`
)

func (m *centerServerService) QueryAllCenterServerListByType(p_centerplatformId int, p_serverType int) ([]*servermodel.CenterQueryServer, error) {
	rst := make([]*servermodel.CenterQueryServer, 0)
	where := ""
	orderBy := " ORDER BY A.platform,A.serverId"
	if p_centerplatformId > 0 {
		where += fmt.Sprintf(" and A.platform = %d", p_centerplatformId)
	}
	exdb := m.db.DB().Raw(getAllCenterServerListByTypeQuerySql+where+orderBy, p_serverType).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServerCount(p_name string, p_platformId int, p_serverType int) (int, error) {
	return m.getCenterServerCount(p_name, p_platformId, p_serverType)
}

func (m *centerServerService) GetCenterServerListByPlatform(p_platformId int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	exdb := m.db.DB().Where("deleteTime = 0 and serverType=0 and platform=?", p_platformId).Order("serverId asc").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServerListByPlatformArray(p_platformId []int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	exdb := m.db.DB().Where("deleteTime = 0 and serverType=0 and platform IN (?)", p_platformId).Order("serverId asc").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

var (
	getCenterMainServerListByPlatformSql = `SELECT *
	FROM
	(
		SELECT A.*,CASE WHEN B.id IS NULL THEN 1 WHEN A.serverId=B.finalServerId THEN 1 ELSE 0 END AS isMain
		FROM
		t_server A
		LEFT JOIN t_merge_record B
		ON A.platform = B.platform and A.serverId = B.fromServerId AND A.serverType=0
		WHERE A.serverType = 0 AND A.platform = ?
	) FA
	WHERE FA.isMain = 1
	`
)

func (m *centerServerService) GetCenterMainServerListByPlatform(p_platformId int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	exdb := m.db.DB().Raw(getCenterMainServerListByPlatformSql, p_platformId).Scan(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServerListByPlatformEnable(p_platformId int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	whereCondition := " and not exists(select 1 from t_merge_record where t_merge_record.platform=t_server.platform and t_merge_record.fromServerId=t_server.serverid)"
	exdb := m.db.DB().Where("deleteTime = 0 and serverType=0 and platform=? "+whereCondition, p_platformId).Order("serverId asc").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServerListByPlatformArrayEnable(p_platformId []int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	whereCondition := " and not exists(select 1 from t_merge_record where t_merge_record.platform=t_server.platform and t_merge_record.fromServerId=t_server.serverid)"
	exdb := m.db.DB().Where("deleteTime = 0 and serverType=0 and platform IN (?)"+whereCondition, p_platformId).Order("serverId asc").Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetAllCenterServerListEnable() ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	whereCondition := " and not exists(select 1 from t_merge_record where t_merge_record.platform=t_server.platform and t_merge_record.fromServerId=t_server.serverid)"
	exerr := m.db.DB().Where("deleteTime=0" + whereCondition).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL中心服务器列表Enable失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServer(p_id int64) (*centerServermodel.CenterServer, error) {
	info := &centerServermodel.CenterServer{}
	exdb := m.db.DB().Where("id = ?", p_id).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return info, nil
}

func (m *centerServerService) existsCenterServer(p_name string, p_platformId int64) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_server").Where("name = ? and platform=? and deleteTime = 0", p_name, p_platformId).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerServerName": p_name,
			"error":            exdb.Error,
		}).Error("查询中心服务器信息失败existsCenterServer")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerServerService) existsCenterServerWithId(p_centerServerid int64, p_name string, p_platformId int64) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_server").Where("name = ? and id != ? and platform=? and deleteTime=0", p_name, p_centerServerid, p_platformId).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerServerName": p_name,
			"centerServerid":   p_centerServerid,
			"error":            exdb.Error,
		}).Error("查询中心服务器信息失败existsCenterServerWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerServerService) existsCenterServerId(p_serverId int, p_platformId int64, p_serverType int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_server").Where("serverId = ? and platform=? and serverType=? and deleteTime = 0", p_serverId, p_platformId, p_serverType).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"p_serverId": p_serverId,
			"error":      exdb.Error,
		}).Error("查询中心服务器信息失败existsCenterServer")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerServerService) existsCenterServerIdWithId(p_centerServerid int64, p_serverId int, p_platformId int64, p_serverType int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_server").Where("serverId = ? and id != ? and platform=? and serverType=? and deleteTime=0", p_serverId, p_centerServerid, p_platformId, p_serverType).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"p_serverId":     p_serverId,
			"centerServerid": p_centerServerid,
			"error":          exdb.Error,
		}).Error("查询中心服务器信息失败existsCenterServerWithId")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *centerServerService) saveCenterServer(p_info *centerServermodel.CenterServer) error {
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.Id > 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}
	exdb := m.db.DB().Save(p_info)
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"centerServerName": p_info.Id,
			"centerServerid":   p_info.ServerName,
			"error":            exdb.Error,
		}).Error("保存中心服务器失败")
		return exdb.Error
	}
	return nil
}

func (m *centerServerService) getCenterServerList(p_name string, p_platformId int, p_serverType int, p_index int) ([]*centerServermodel.CenterServer, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*centerServermodel.CenterServer, 0)
	where := ""
	if p_platformId > 0 {
		where += fmt.Sprintf(" and platform=%d", p_platformId)
	}
	if p_serverType > -1 {
		where += fmt.Sprintf(" and serverType=%d", p_serverType)
	}

	exerr := m.db.DB().Where("deleteTime =0 and name like ?"+where, "%"+p_name+"%").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerServerName": p_name,
			"error":            exerr.Error,
		}).Error("获取中心服务器列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerServerService) getCenterServerCount(p_name string, p_platformId int, p_serverType int) (int, error) {
	rst := 0
	where := ""
	if p_platformId > 0 {
		where += fmt.Sprintf(" and platform=%d", p_platformId)
	}
	if p_serverType > -1 {
		where += fmt.Sprintf(" and serverType=%d", p_serverType)
	}
	exerr := m.db.DB().Table("t_server").Where("deleteTime =0 and name like ? "+where, "%"+p_name+"%").Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"centerServerName": p_name,
			"error":            exerr.Error,
		}).Error("获取中心服务器列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}

func (m *centerServerService) getAllCenterServerList() ([]*centerServermodel.CenterServer, error) {

	rst := make([]*centerServermodel.CenterServer, 0)

	exerr := m.db.DB().Where("deleteTime=0").Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取ALL中心服务器列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServerName(p_platformId int, p_serverId int) (string, error) {
	info := &centerServermodel.CenterServer{}
	exerr := m.db.DB().Where("platform=? and serverId=? and serverType=0 and deleteTime=0", p_platformId, p_serverId).First(info)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		return "", exerr.Error
	}
	return info.ServerName, nil
}

func (m *centerServerService) GetCenterServerInfo(p_platformId int, p_serverId int) (*centerServermodel.CenterServer, error) {
	info := &centerServermodel.CenterServer{}
	exerr := m.db.DB().Where("platform=? and serverId=? and serverType=0 and deleteTime=0", p_platformId, p_serverId).First(info)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		return nil, exerr.Error
	}
	return info, nil
}

func (m *centerServerService) GetCenterMergeRecord(p_platformId int, p_serverId int) ([]*servermodel.ServerMergeRecord, error) {
	rst := make([]*servermodel.ServerMergeRecord, 0)
	exdb := m.db.DB().Where("platform=? and finalServerId=? and deleteTime =0", p_platformId, p_serverId).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetCenterServerZhanQu(p_platform int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	exdb := m.db.DB().Where("platform=? and serverType=2 and deleteTime =0", p_platform).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) GetZhanQuServer(p_platform int, p_zhanQuServer []int) ([]*centerServermodel.CenterServer, error) {
	rst := make([]*centerServermodel.CenterServer, 0)
	exdb := m.db.DB().Where("platform=? and parentServerId IN (?) and serverType=0 and deleteTime =0", p_platform, p_zhanQuServer).Find(&rst)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		return nil, exdb.Error
	}
	return rst, nil
}

func (m *centerServerService) RegisterDB(p_serverid gmdb.GameDbLink, p_host string, p_port string, p_dbName string, p_userName string, p_password string) error {
	err := gmdb.RegisterDB(p_serverid, p_host, p_port, p_dbName, p_userName, p_password)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"host":     p_host,
			"port":     p_port,
			"dbName":   p_dbName,
			"userName": p_userName,
			"password": p_password,
		}).Error("注册数据库失败")
		return gmError.GetError(gmError.ErrorCodeCenterServerDBNotConnect)
	}
	return nil
}

func (m *centerServerService) RegisterGrpc(p_serverId int32, p_host string, p_port string) error {
	grpcHost := fmt.Sprintf("%s:%s", p_host, p_port)
	_, grpcerr := userremote.RegisterGrpc(p_serverId, grpcHost)
	if grpcerr != nil {
		log.WithFields(log.Fields{
			"error": grpcerr,
			"host":  p_host,
			"port":  p_port,
		}).Error("注册Remote失败")
		return gmError.GetError(gmError.ErrorCodeCenterServerRemoteNotConnect)
	}
	return nil
}

func NewCenterServerService(p_db gmdb.DBService) ICenterServerService {
	rst := &centerServerService{
		db: p_db,
	}
	return rst
}

type contextKey string

const (
	centerServerServiceKey = contextKey("CenterServerService")
)

func WithCenterServerService(ctx context.Context, ls ICenterServerService) context.Context {
	return context.WithValue(ctx, centerServerServiceKey, ls)
}

func CenterServerServiceInContext(ctx context.Context) ICenterServerService {
	us, ok := ctx.Value(centerServerServiceKey).(ICenterServerService)
	if !ok {
		return nil
	}
	return us
}

func SetupCenterServerServiceHandler(ls ICenterServerService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCenterServerService(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
