package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	setservice "fgame/fgame/gm/gamegm/gm/center/set/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type platformServerConfigGetRespon struct {
	TradeServerIp string `json:"tradeServerIp"`
}

func handleplatformServerConfigGet(rw http.ResponseWriter, req *http.Request) {
	service := setservice.CenterSetServiceInContext(req.Context())
	rst, err := service.GetPlatformServerConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取交易ip设置，获取异常")
		errhttp.ResponseWithError(rw, err)
		return
	}
	respon := &platformServerConfigGetRespon{
		TradeServerIp: rst.TradeServerIp,
	}
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
