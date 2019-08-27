package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerPlatformRequest struct {
	CenterPlatformName string `form:"centerPlatformName" json:"centerPlatformName"`
	SdkType            int    `form:"sdkType" json:"sdkType"`
}

func handleAddCenterPlatform(rw http.ResponseWriter, req *http.Request) {
	form := &centerPlatformRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加中心平台，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加中心平台,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.AddCenterPlatform(form.CenterPlatformName)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
