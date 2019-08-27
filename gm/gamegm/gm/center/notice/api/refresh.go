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

func handleCenterNoticeRefresh(rw http.ResponseWriter, req *http.Request) {

	service := userremote.NoticeInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("刷新中心服务器通知，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := service.RefreshNotice()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("刷新中心服务器通知，异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteNotice)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
