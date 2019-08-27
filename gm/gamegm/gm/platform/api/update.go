package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type platformUpdateRequest struct {
	PlatformId       int64  `form:"platformId" json:"platformId"`
	CenterPlatformId int64  `form:"centerPlatformId" json:"centerPlatformId"`
	PlatformName     string `form:"platformName" json:"platformName"`
	ChannelId        int64  `form:"channelId" json:"channelId"`
	SdkType          int    `form:"sdkType" json:"sdkType"`
	SignKey          string `form:"signKey" json:"signKey"`
}

func handleUpdatePlatform(rw http.ResponseWriter, req *http.Request) {
	form := &platformUpdateRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新渠道，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmplatform.PlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新渠道,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.UpdatePlatform(form.PlatformId, form.PlatformName, form.ChannelId, form.CenterPlatformId, form.SdkType, form.SignKey)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
