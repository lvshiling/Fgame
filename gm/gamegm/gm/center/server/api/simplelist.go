package api

import (
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerSimpleServerListRequest struct {
	PageIndex        int    `form:"pageIndex" json:"pageIndex"`
	CenterServerName string `form:"centerServerName" json:"centerServerName"`
	PlatformId       int    `form:"centerPlatformId" json:"centerPlatformId"`
	ServerType       int    `form:"serverType" json:"serverType"`
}

type centerSimpleServerListRespon struct {
	ItemArray  []*centerSimpleServerListResponItem `json:"itemArray"`
	TotalCount int                                 `json:"total"`
}

type centerSimpleServerListResponItem struct {
	CenterServerId int64  `json:"id"`
	ServerId       int    `json:"serverId"`
	PlatformId     int64  `json:"platformId"`
	ServerName     string `json:"serverName"`
	StartTime      int64  `json:"startTime"`
}

func handleSimpleCenterServerList(rw http.ResponseWriter, req *http.Request) {
	form := &centerSimpleServerListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心服务器列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心服务器列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetCenterServerList(form.CenterServerName, form.PlatformId, form.ServerType, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error":            err,
			"CenterServerName": form.CenterServerName,
			"index":            form.PageIndex,
		}).Error("获取中心服务器列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &centerSimpleServerListRespon{}
	respon.ItemArray = make([]*centerSimpleServerListResponItem, 0)

	for _, value := range rst {
		item := &centerSimpleServerListResponItem{
			CenterServerId: value.Id,
			ServerId:       value.ServerId,
			PlatformId:     value.Platform,
			ServerName:     value.ServerName,
			StartTime:      value.StartTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetCenterServerCount(form.CenterServerName, form.PlatformId, form.ServerType)
	if err != nil {
		log.WithFields(log.Fields{
			"error":            err,
			"CenterServerName": form.CenterServerName,
			"index":            form.PageIndex,
		}).Error("获取中心服务器列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
