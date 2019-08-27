package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	serversupp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type deleteServerSupportPoolRequest struct {
	ID   int `json:"id"`
	Gold int `json:"gold"`
}

func handleDeleteSupportPool(rw http.ResponseWriter, req *http.Request) {
	form := &deleteServerSupportPoolRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除扶植池，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	poolService := serversupp.ServerSupportPoolInContext(req.Context())
	if poolService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除扶植池，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// userid := gmUserService.GmUserIdInContext(req.Context())
	err = poolService.DeleteServerSupportPool(form.ID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除扶植池异常")
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
