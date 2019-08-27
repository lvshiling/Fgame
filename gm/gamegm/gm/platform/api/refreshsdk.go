package api

import (
	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	centerremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleRefreshSdk(rw http.ResponseWriter, req *http.Request) {
	service := centerremote.CenterServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("刷新skd,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := service.RefreshSDK()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("刷新skd,刷新失败")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteCenter)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
