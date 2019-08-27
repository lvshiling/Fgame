package api

import (
	"net/http"

	gmerr "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	serversupp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type addPlatformSupportPoolSetRequest struct {
	Gold             int32 `json:"gold"`
	Percent          int32 `json:"percent"`
	CenterPlatformId int64 `json:"centerPlatformId"`
}

func handleAddPlatformSupportPoolSet(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleAddPlatformSupportPoolSet:添加平添配置")
	form := &addPlatformSupportPoolSetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加平台扶植池，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	poolService := serversupp.ServerSupportPoolInContext(req.Context())
	if poolService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加平台扶植池，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	// userid := gmUserService.GmUserIdInContext(req.Context())
	poolSetInfo, err := poolService.GetPlatformSupportPoolSet(form.CenterPlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加平台扶植池，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if poolSetInfo.Id > 0 {
		errhttp.ResponseWithError(rw, gmerr.GetError(gmerr.ErrorCodePlatformSupportPoolSetExists))
		return
	}
	err = poolService.AddPlatformSupportPoolSet(form.CenterPlatformId, form.Gold, form.Percent)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加平台扶植池异常")
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
