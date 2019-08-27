package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	setservice "fgame/fgame/gm/gamegm/gm/center/set/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type clientVersionRespon struct {
	AndroidVersion string `json:"androidVersion"`
	IosVersion     string `json:"iosVersion"`
}

func handleClientVersionGet(rw http.ResponseWriter, req *http.Request) {
	service := setservice.CenterSetServiceInContext(req.Context())
	rst, err := service.GetClientVersion()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("设置客户端版本号，获取异常")
		errhttp.ResponseWithError(rw, err)
		return
	}
	respon := &clientVersionRespon{
		AndroidVersion: rst.AndroidVersion,
		IosVersion:     rst.IosVersion,
	}
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
