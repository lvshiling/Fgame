package api

import (
	gmcenter "fgame/fgame/gm/gamegm/gm/center/server/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	userservice "fgame/fgame/gm/gamegm/gm/user/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleAllCenterServer(rw http.ResponseWriter, req *http.Request) {
	service := gmcenter.CenterServerServiceInContext(req.Context())
	rsp := &serverRespon{}
	rsp.ItemArray = make([]*serverResponItem, 0)

	userCenterPlatformList, err := userservice.GetUserCenterPlatList(req.Context())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心所有服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetAllCenterServerListArray(userCenterPlatformList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &serverResponItem{
			ID:         value.Id,
			ServerId:   value.ServerId,
			ServerType: value.ServerType,
			Name:       value.ServerName,
			PlatformId: value.Platform,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
