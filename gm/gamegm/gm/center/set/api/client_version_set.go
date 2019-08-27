package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	setservice "fgame/fgame/gm/gamegm/gm/center/set/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type clientVersionSetRequest struct {
	AndroidVersion string `form:"androidVersion" json:"androidVersion"`
	IosVersion     string `form:"iosVersion" json:"iosVersion"`
}

func handleClientVersionSet(rw http.ResponseWriter, req *http.Request) {
	form := &clientVersionSetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("设置客户端版本号，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := setservice.CenterSetServiceInContext(req.Context())
	err = service.SaveClientVersion(form.IosVersion, form.AndroidVersion)
	if err != nil {
		log.WithFields(log.Fields{
			"error":          err,
			"androidVersion": form.AndroidVersion,
			"iosVersion":     form.IosVersion,
		}).Error("设置客户端版本号，保存异常")
		errhttp.ResponseWithError(rw, err)
		return
	}
	rs := remoteservice.CenterServiceInContext(req.Context())
	err = rs.RefreshClientVersion()
	if err != nil {
		log.WithFields(log.Fields{
			"error":          err,
			"androidVersion": form.AndroidVersion,
			"iosVersion":     form.IosVersion,
		}).Error("设置客户端版本号，刷新remote异常")
		failErr := gmhttp.NewFailedResultWithMsg(10001, err.Error())
		httputils.WriteJSON(rw, http.StatusOK, failErr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
