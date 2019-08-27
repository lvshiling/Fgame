package api

import (
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type ipGetStateRequest struct {
	Ip string `json:"ip"`
}

type ipGetStateRespon struct {
	Forbid bool `json:"forbid"`
}

func handleIpGetState(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleIpGetState:解禁ip")
	form := &ipGetStateRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleIpGetState，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithFields(log.Fields{
		"ip": form.Ip,
	}).Debug("handleIpGetState：解析参数")

	rs := remoteservice.CenterServiceInContext(req.Context())
	rmrst, err := rs.ForbidIpSearch(form.Ip)
	if err != nil {
		rr := gmhttp.NewFailedResultWithMsg(1000, err.Error())
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}
	respon := &ipGetStateRespon{
		Forbid: rmrst.Result,
	}
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
