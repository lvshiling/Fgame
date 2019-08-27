package api

import (
	"net/http"
	"strconv"

	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type allianceDismissForm struct {
	ServerId   int    `json:"serverId"`
	AllianceId string `json:"allianceId"`
}

func handleAllianceDismiss(rw http.ResponseWriter, req *http.Request) {
	form := &allianceGongGaoForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("alliance:仙盟解散，解析异常")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("alliance:仙盟解散，remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	allianceId, err := strconv.ParseInt(form.AllianceId, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("alliance:仙盟解散，解析异常")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = remoteService.DismissAlliance(int32(form.ServerId), allianceId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("alliance:仙盟解散，获取失败")
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}
	result := &struct{}{}
	rr := gmhttp.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
