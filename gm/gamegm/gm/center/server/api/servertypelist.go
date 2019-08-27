package api

import (
	"net/http"

	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerTypeServerListRequest struct {
	ServerType       int `form:"serverType" json:"serverType"`
	CenterPlatformId int `form:"platformId" json:"platformId"`
}

type centerTypeServerListRespon struct {
	ItemArray  []*centerTypeServerListResponItem `json:"itemArray"`
	TotalCount int                               `json:"total"`
}

type centerTypeServerListResponItem struct {
	CenterServerId       int    `json:"serverId"`
	ServerName           string `json:"serverName"`
	FinnalServerId       int32  `json:"finnalServerId"`
	Id                   int64  `json:"id"`
	ParentServerId       int    `json:"parentServerId"`
	JiaoYiZhanQuServerId int32  `json:"jiaoYiZhanQuServerId"`
	PingTaiFuServerId    int32  `json:"pingTaiFuServerId"`
	PlatformId           int32  `json:"platformId"`
	ChengZhanServerId    int32  `json:"chengZhanServerId"`
}

func handleCenterTypeServerList(rw http.ResponseWriter, req *http.Request) {
	form := &centerTypeServerListRequest{}
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

	rst, err := service.QueryAllCenterServerListByType(form.CenterPlatformId, form.ServerType)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"ServerType": form.ServerType,
		}).Error("获取中心服务器列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &centerTypeServerListRespon{}
	respon.ItemArray = make([]*centerTypeServerListResponItem, 0)

	for _, value := range rst {
		item := &centerTypeServerListResponItem{
			CenterServerId:       value.ServerId,
			ServerName:           value.ServerName,
			Id:                   value.Id,
			ParentServerId:       value.ParentServerId,
			JiaoYiZhanQuServerId: value.JiaoYiZhanQuServerId,
			PingTaiFuServerId:    value.PingTaiFuServerId,
			PlatformId:           int32(value.Platform),
			FinnalServerId:       value.FinnalServerId,
			ChengZhanServerId:    value.ChengZhanServerId,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
