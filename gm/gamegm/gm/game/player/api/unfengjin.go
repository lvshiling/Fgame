package api

import (
	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	userremote "fgame/fgame/gm/gamegm/remote/service"

	monitor "fgame/fgame/gm/gamegm/monitor"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type unForbidRequest struct {
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	PlayerId   string `json:"playerId"`
}

func handleUnForbid(rw http.ResponseWriter, req *http.Request) {
	form := &unForbidRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("玩家封禁，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := userremote.UserRemoteServiceInContext(req.Context())
	if service == nil {
		log.Error("玩家封禁，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	centerService := monitor.CenterServerServiceInContext(req.Context())

	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)
	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("玩家封禁，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.UnForbidPlayer(int32(serverid), changeStringToInt64(form.PlayerId))
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"serverid": serverid,
		}).Error("玩家封禁，玩家封禁异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
