package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	"fgame/fgame/gm/gamegm/gm/center/platform/model"
	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerPlatformSettingRequest struct {
	CenterPlatformId int64  `form:"centerPlatformId" json:"centerPlatformId"`
	Setting          string `json:"setting"`
}

func handleCenterPlatformSettingSave(rw http.ResponseWriter, req *http.Request) {
	form := &centerPlatformSettingRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("settingSave:设置中心平台配置，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithFields(log.Fields{
		"setting": form.Setting,
	}).Debug("配置设置")

	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("settingSave:设置中心平台配置，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	now := timeutils.TimeToMillisecond(time.Now())
	info := &model.CenterPlatformSetInfo{}
	info.PlatformId = form.CenterPlatformId
	info.SettingContent = form.Setting
	info.UpdateTime = now
	info.CreateTime = now
	err = service.SavePlatformSet(info)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("settingSave:设置中心平台配置，配置异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rs := remoteservice.CenterServiceInContext(req.Context())
	err = rs.RefreshPlatformConfig(int32(form.CenterPlatformId))
	if err != nil {
		rsErr := gmhttp.NewFailedResultWithMsg(111, err.Error())
		httputils.WriteJSON(rw, http.StatusOK, rsErr)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
