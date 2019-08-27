package api

import (
	"net/http"

	jyzqservice "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type jiaoYiZhanQuDeleteRequest struct {
	Id int32 `json:"id"`
}

func handleJiaoYiZhanQuDelete(rw http.ResponseWriter, req *http.Request) {
	form := &jiaoYiZhanQuDeleteRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区更新，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := jyzqservice.JiaoYiZhanQuServiceInContext(req.Context())

	err = service.DeleteZhanQu(form.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区更新，更新战区信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
