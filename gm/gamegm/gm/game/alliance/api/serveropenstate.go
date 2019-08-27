package api

import (
	"net/http"

	gmdb "fgame/fgame/gm/gamegm/db"
	alliservice "fgame/fgame/gm/gamegm/gm/game/alliance/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type serverOpenRequest struct {
	ServerId int `json:"serverId"`
}

type serverOpenRespon struct {
	Open int `json:"open"`
}

func handleServerOpenList(rw http.ResponseWriter, req *http.Request) {
	form := &serverOpenRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("服务器启用状态，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := alliservice.AllianceServiceInContext(req.Context())
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

	rst, err := service.GetServerRegisterFlag(gmdb.GameDbLink(form.ServerId), acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("服务器启用状态，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &serverOpenRespon{}
	respon.Open = rst
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
