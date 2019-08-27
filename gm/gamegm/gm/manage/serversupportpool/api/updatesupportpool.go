package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	serversupp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateServerSupportPoolRequest struct {
	ID               int   `json:"id"`
	Gold             int   `json:"gold"`
	Percent          int32 `json:"percent"`
	SdkType          int   `json:"sdkType"`
	CenterPlatformId int64 `json:"centerPlatformId"`
}

func handleUpdateSupportPool(rw http.ResponseWriter, req *http.Request) {
	form := &updateServerSupportPoolRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新扶植池，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	poolService := serversupp.ServerSupportPoolInContext(req.Context())
	if poolService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新扶植池，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// userid := gmUserService.GmUserIdInContext(req.Context())
	err = poolService.UpdateServerSupportPool(form.ID, form.Gold, form.Percent)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新扶植池异常")
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
