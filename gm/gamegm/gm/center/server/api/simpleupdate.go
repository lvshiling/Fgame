package api

import (
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type simpleUpdateRequest struct {
	CenterServerId int    `form:"id" json:"id"`
	ServerName     string `form:"serverName" json:"serverName"`
}

func handleUpdateCenterServerName(rw http.ResponseWriter, req *http.Request) {
	form := &simpleUpdateRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新中心服务器名称，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新中心服务器名称,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	remoteService := userremote.CenterServiceInContext(req.Context())
	if remoteService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("刷新中心服务器，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serverInfo, err := service.GetCenterServer(int64(form.CenterServerId))
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	err = service.UpdateCenterServerName(int64(form.CenterServerId), form.ServerName, serverInfo.Platform)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	centerPlatformId := serverInfo.Platform
	err = remoteService.Refresh(int32(centerPlatformId))
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
