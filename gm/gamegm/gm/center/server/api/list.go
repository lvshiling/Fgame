package api

import (
	"net/http"

	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerServerListRequest struct {
	PageIndex        int    `form:"pageIndex" json:"pageIndex"`
	CenterServerName string `form:"centerServerName" json:"centerServerName"`
	PlatformId       int    `form:"platformId" json:"platformId"`
	ServerType       int    `form:"serverType" json:"serverType"`
}

type centerServerListRespon struct {
	ItemArray  []*centerServerListResponItem `json:"itemArray"`
	TotalCount int                           `json:"total"`
}

type centerServerListResponItem struct {
	CenterServerId   int64  `json:"id"`
	ServerType       int    `json:"serverType"`
	ServerId         int    `json:"serverId"`
	PlatformId       int64  `json:"platformId"`
	ServerName       string `json:"serverName"`
	StartTime        int64  `json:"startTime"`
	ServerIp         string `json:"serverIp"`
	ServerPort       string `json:"serverPort"`
	ServerRemoteIp   string `json:"serverRemoteIp"`
	ServerRemotePort string `json:"serverRemotePort"`
	ServerDbIp       string `json:"serverDBIp"`
	ServerDbPort     string `json:"serverDBPort"`
	ServerDBName     string `json:"serverDBName"`
	ServerDBUser     string `json:"serverDBUser"`
	ServerDBPassword string `json:"serverDBPassword"`
	ServerTag        int    `json:"serverTag"`
	ServerStatus     int    `json:"serverStatus"`
	ParentServerId   int    `json:"parentServerId"`
	PreShow          bool   `json:"preShow"`
}

func handleCenterServerList(rw http.ResponseWriter, req *http.Request) {
	form := &centerServerListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心服务器列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心服务器列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetCenterServerList(form.CenterServerName, form.PlatformId, form.ServerType, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error":            err,
			"CenterServerName": form.CenterServerName,
			"index":            form.PageIndex,
		}).Error("获取中心服务器列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &centerServerListRespon{}
	respon.ItemArray = make([]*centerServerListResponItem, 0)

	for _, value := range rst {
		item := &centerServerListResponItem{
			CenterServerId:   value.Id,
			ServerType:       value.ServerType,
			ServerId:         value.ServerId,
			PlatformId:       value.Platform,
			ServerName:       value.ServerName,
			StartTime:        value.StartTime,
			ServerIp:         value.ServerIp,
			ServerPort:       value.ServerPort,
			ServerRemoteIp:   value.ServerRemoteIp,
			ServerRemotePort: value.ServerRemotePort,
			ServerDbIp:       value.ServerDbIp,
			ServerDbPort:     value.ServerDbPort,
			ServerDBName:     value.ServerDBName,
			ServerDBUser:     value.ServerDBUser,
			ServerDBPassword: value.ServerDBPassword,
			ServerTag:        value.ServerTag,
			ServerStatus:     value.ServerStatus,
			ParentServerId:   value.ParentServerId,
		}
		if value.PreShow > 0 {
			item.PreShow = true
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetCenterServerCount(form.CenterServerName, form.PlatformId, form.ServerType)
	if err != nil {
		log.WithFields(log.Fields{
			"error":            err,
			"CenterServerName": form.CenterServerName,
			"index":            form.PageIndex,
		}).Error("获取中心服务器列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
