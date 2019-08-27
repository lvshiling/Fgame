package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerServerUpdateParentRequest struct {
	CenterServerId int `form:"id" json:"id"`
	ParentServerId int `form:"parentServerId" json:"parentServerId"`
}

func handleUpdateCenterServerParent(rw http.ResponseWriter, req *http.Request) {
	form := &centerServerUpdateParentRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新中心服务器，解析异常")
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

	err = service.UpdateParentServerId(form.CenterServerId, form.ParentServerId)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
