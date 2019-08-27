package api

import (
	gmcenter "fgame/fgame/gm/gamegm/gm/center/service"
	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"
	us "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerPlatformRespon struct {
	ItemArray []*centerPlatformResponItem `json:"itemArray"`
}

type centerPlatformResponItem struct {
	PlatformId int64  `json:"centerPlatformId"`
	Name       string `json:"centerPlatformName"`
}

func handleAllCenterPlatform(rw http.ResponseWriter, req *http.Request) {
	service := gmcenter.CenterPlatformServiceInContext(req.Context())
	rsp := &centerPlatformRespon{}
	rsp.ItemArray = make([]*centerPlatformResponItem, 0)

	rst, err := service.GetAllCenterPlatform()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	gmPlatformService := gmplatform.PlatformServiceInContext(req.Context())
	if gmPlatformService == nil {
		log.WithFields(log.Fields{}).Error("获取平台列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userId := us.GmUserIdInContext(req.Context())
	centerPlatformList, err := gmPlatformService.GetAllUserCenterPlatformList(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"UserId": userId,
			"error":  err,
		}).Error("获取平台列表，获取中心权限异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	centerMap := make(map[int64]int64)
	for _, value := range centerPlatformList {
		centerMap[value] = value
	}

	for _, value := range rst {
		if len(centerMap) > 0 { //为0则表示所有的都有
			_, exists := centerMap[value.PlatformId]
			if !exists {
				continue
			}
		}
		item := &centerPlatformResponItem{
			PlatformId: value.PlatformId,
			Name:       value.Name,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
