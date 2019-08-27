package api

import (
	"net/http"

	gmdb "fgame/fgame/gm/gamegm/db"
	singleserverservice "fgame/fgame/gm/gamegm/gm/game/singleserver/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type doubleChargeRequest struct {
	ServerId int `json:"serverId"`
}

type doubleChargeRespon struct {
	StartTime int64 `json:"startTime"`
}

func handleDoubleCharge(rw http.ResponseWriter, req *http.Request) {
	form := &doubleChargeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("服务器启用状态，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := singleserverservice.SingleServerServiceInContext(req.Context())
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetFirstCharge(gmdb.GameDbLink(form.ServerId), acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &doubleChargeRespon{}
	respon.StartTime = rst
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
