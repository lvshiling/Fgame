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

type allianceGongGaoForm struct {
	ServerId   int    `json:"serverId"`
	GongGao    string `json:"gongGao"`
	AllianceId string `json:"allianceId"`
}

func handleAllianceGongGao(rw http.ResponseWriter, req *http.Request) {
	form := &allianceGongGaoForm{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("alliance:仙盟公告修改，解析异常")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("alliance:仙盟公告修改，remote服务为空")
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
	err = remoteService.ModifyAllianceGongGao(int32(form.ServerId), allianceId, form.GongGao)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("alliance:仙盟公告修改，获取失败")
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}
	result := &struct{}{}
	rr := gmhttp.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
