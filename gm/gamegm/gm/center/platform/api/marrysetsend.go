package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	remoteservice "fgame/fgame/gm/gamegm/remote/service"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type marrySetLogSendRequest struct {
	PlatformId int   `form:"centerPlatformId" json:"centerPlatformId"`
	ServerId   int32 `json:"centerServerId"`
	Id         int32 `json:"id"`
	KindType   int32 `json:"kindType"`
}

func handleMarrySetLogSend(rw http.ResponseWriter, req *http.Request) {
	form := &marrySetLogSendRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心平台列表结婚日志发送，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerService := monitor.CenterServerServiceInContext(req.Context())

	serverid := centerService.GetCenterServerDBId(int32(form.PlatformId), int32(form.ServerId))
	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("中心平台列表结婚日志发送，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rs := remoteservice.UserRemoteServiceInContext(req.Context())
	err = rs.SetMarryBanquetHouTaiType(int32(serverid), form.KindType)
	if err != nil {
		log.WithFields(log.Fields{
			"serverid": serverid,
			"err":      err,
		}).Error("中心平台列表结婚日志发送，发送失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚日志，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.UpdatePlatformMarryServerLogState(form.Id, 1)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚日志，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
