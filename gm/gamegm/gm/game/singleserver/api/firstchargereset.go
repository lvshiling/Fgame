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

type doubleChargeResetRequest struct {
	ServerId int32 `json:"serverId"`
}

func handleDoubleChargeReset(rw http.ResponseWriter, req *http.Request) {
	form := &doubleChargeResetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("服务器启用状态设置，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("添加聊天配置，remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = remoteService.FirstChargeReset(form.ServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态设置，获取失败")
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
