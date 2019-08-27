package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type platformDeleteRequest struct {
	PlatformId int64 `form:"platformId" json:"platformId"`
}

func handleDeletePlatform(rw http.ResponseWriter, req *http.Request) {
	form := &platformDeleteRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除渠道，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmplatform.PlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除渠道,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.DeletePlatform(form.PlatformId)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
