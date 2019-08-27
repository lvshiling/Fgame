package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerServerUpdateJiaoYiZhanQuArrayRequest struct {
	CenterServerId       []int `form:"id" json:"id"`
	JiaoYiZhanQuServerId int   `form:"jiaoYiZhanQuServerId" json:"jiaoYiZhanQuServerId"`
}

func handleUpdateCenterServerJiaoYiZhanQuArray(rw http.ResponseWriter, req *http.Request) {
	form := &centerServerUpdateJiaoYiZhanQuArrayRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新中心服务器交易战区服，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新中心服务器,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.UpdateJiaoYiZhanQuFuArray(form.CenterServerId, form.JiaoYiZhanQuServerId)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
