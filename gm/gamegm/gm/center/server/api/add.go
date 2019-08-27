package api

import (
	"net/http"

	gmdb "fgame/fgame/gm/gamegm/db"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerServerRequest struct {
	ServerType       int    `form:"serverType" json:"serverType"`
	ServerId         int    `form:"serverId" json:"serverId"`
	PlatformId       int64  `form:"platformId" json:"platformId"`
	ServerName       string `form:"serverName" json:"serverName"`
	StartTime        int64  `form:"startTime" json:"startTime"`
	ServerIp         string `form:"serverIp" json:"serverIp"`
	ServerPort       string `form:"serverPort" json:"serverPort"`
	ServerRemoteIp   string `form:"serverRemoteIp" json:"serverRemoteIp"`
	ServerRemotePort string `form:"serverRemotePort" json:"serverRemotePort"`
	ServerDbIp       string `form:"serverDBIp" json:"serverDBIp"`
	ServerDbPort     string `form:"serverDBPort" json:"serverDBPort"`
	ServerDBName     string `form:"serverDBName" json:"serverDBName"`
	ServerDBUser     string `form:"serverDBUser" json:"serverDBUser"`
	ServerDBPassword string `form:"serverDBPassword" json:"serverDBPassword"`
	ServerTag        int    `form:"serverTag" json:"serverTag"`
	ServerStatus     int    `form:"serverStatus" json:"serverStatus"`
	ParentServerId   int    `form:"parentServerId" json:"parentServerId"`
	PreShow          bool   `form:"preShow" json:"preShow"`
}

func handleAddCenterServer(rw http.ResponseWriter, req *http.Request) {
	form := &centerServerRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加中心平台，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加中心平台,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	preShow := 0
	if form.PreShow {
		preShow = 1
	}

	id, err := service.AddCenterServer(form.ServerType, form.ServerId, form.PlatformId, form.ServerName, form.StartTime, form.ServerIp, form.ServerPort, form.ServerRemoteIp, form.ServerRemotePort, form.ServerDbIp, form.ServerDbPort, form.ServerDBName, form.ServerDBUser, form.ServerDBPassword, form.ServerTag, form.ServerStatus, form.ParentServerId, preShow)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	if id > 0 && form.ServerType == 0 {
		err = service.RegisterDB(gmdb.GameDbLink(int32(id)), form.ServerDbIp, form.ServerDbPort, form.ServerDBName, form.ServerDBUser, form.ServerDBPassword)
		if err != nil {
			errhttp.ResponseWithError(rw, err)
			return
		}
		err = service.RegisterGrpc(int32(id), form.ServerRemoteIp, form.ServerRemotePort)
		if err != nil {
			errhttp.ResponseWithError(rw, err)
			return
		}
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
