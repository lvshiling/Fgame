package api

import (
	"net/http"

	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerServerPingRequest struct {
	CenterServerId int32 `form:"id" json:"id"`
	ServerId       int32 `form:"serverId" json:"serverId"`
	PlatformId     int32 `form:"platformId" json:"platformId"`
}

func handlePingCenterServer(rw http.ResponseWriter, req *http.Request) {
	form := &centerServerPingRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ping游戏服务器，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := remoteservice.UserRemoteServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ping游戏服务器,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.Ping(form.CenterServerId, form.ServerId, form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ping游戏服务器,ping异常")

		rr := gmhttp.NewFailedResultWithMsg(100, err.Error())
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
