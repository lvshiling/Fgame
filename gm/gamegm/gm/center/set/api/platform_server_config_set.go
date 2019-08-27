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

type platformServerConfigSetRequest struct {
	TradeServerIp string `form:"tradeServerIp" json:"tradeServerIp"`
}

func handlePlatformServerConfigSet(rw http.ResponseWriter, req *http.Request) {
	form := &platformServerConfigSetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("设置交易服务器，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := setservice.CenterSetServiceInContext(req.Context())
	err = service.SavePlatformServerConfig(form.TradeServerIp)
	if err != nil {
		log.WithFields(log.Fields{
			"error":         err,
			"tradeServerIp": form.TradeServerIp,
		}).Error("设置交易服务器，保存异常")
		errhttp.ResponseWithError(rw, err)
		return
	}
	rs := remoteservice.CenterServiceInContext(req.Context())
	err = rs.RefreshPlatformServerConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("设置交易服务器，刷新remote异常")
		failErr := gmhttp.NewFailedResultWithMsg(10001, err.Error())
		httputils.WriteJSON(rw, http.StatusOK, failErr)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
