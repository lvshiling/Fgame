package api

import (
	gmcenterServer "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerServerZhanQuRequest struct {
	PlatformId int `form:"centerPlatformId" json:"centerPlatformId"`
}

type centerServerZhanQuRespon struct {
	ItemArray []*centerServerZhanQuResponItem `json:"itemArray"`
}

type centerServerZhanQuResponItem struct {
	ZhanQuServerId   int32  `json:"zhanquServerId"`
	ZhanQuServerName string `json:"zhanquServerName"`
	ZhanQuChild      string `json:"zhanQuChild"`
}

func handleCenterServerZhanQuList(rw http.ResponseWriter, req *http.Request) {

	form := &centerServerZhanQuRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterServerZhanQuList，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterServer.CenterServerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterServerZhanQuList，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	zhanQuList, err := service.GetCenterServerZhanQu(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"platformId": form.PlatformId,
			"error":      err,
		}).Error("handleCenterServerZhanQuList，获取战区服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	zhanQuServerId := make([]int, 0)
	for _, value := range zhanQuList {
		zhanQuServerId = append(zhanQuServerId, value.ServerId)
	}

	childServer, err := service.GetZhanQuServer(form.PlatformId, zhanQuServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"platformId": form.PlatformId,
			"zhanQuList": zhanQuServerId,
			"error":      err,
		}).Error("handleCenterServerZhanQuList，子服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &centerServerZhanQuRespon{}
	respon.ItemArray = make([]*centerServerZhanQuResponItem, 0)
	for _, value := range zhanQuList {
		item := &centerServerZhanQuResponItem{
			ZhanQuServerId:   int32(value.ServerId),
			ZhanQuServerName: value.ServerName,
		}
		for _, childValue := range childServer {
			if childValue.ParentServerId != value.ServerId {
				continue
			}
			if len(item.ZhanQuChild) > 0 {
				item.ZhanQuChild += ","
			}
			item.ZhanQuChild += strconv.Itoa(childValue.ServerId)
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
