package api

import (
	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type recycleGoldRequest struct {
	ServerId int   `json:"serverId"`
	Gold     int64 `json:"gold"`
}

func handlerecycleGold(rw http.ResponseWriter, req *http.Request) {
	log.Debug("recycleGold:增加自定义回收")
	form := &recycleGoldRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("recycleGold，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := userremote.UserRemoteServiceInContext(req.Context())
	if service == nil {
		log.Error("recycleGold，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.CustomRecycleGold(int32(form.ServerId), form.Gold)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"serverid": form.ServerId,
		}).Error("recycleGold，增加自定义回收异常")
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
