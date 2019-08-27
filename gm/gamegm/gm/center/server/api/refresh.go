package api

import (
	"net/http"

	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerRefreshRequest struct {
	ServerId int32 `form:"serverId" json:"serverId"`
}

func handleCenterRefresh(rw http.ResponseWriter, req *http.Request) {
	form := &centerRefreshRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("刷新中心服务器，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := userremote.CenterServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("刷新中心服务器，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.Refresh(form.ServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("刷新中心服务器，异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteCenter)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
