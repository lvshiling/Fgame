package api

import (
	"net/http"

	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerPlatformListRequest struct {
	PageIndex          int    `form:"pageIndex" json:"pageIndex"`
	CenterPlatformName string `form:"centerPlatformName" json:"centerPlatformName"`
}

type centerPlatformListRespon struct {
	ItemArray  []*centerPlatformListResponItem `json:"itemArray"`
	TotalCount int                             `json:"total"`
}

type centerPlatformListResponItem struct {
	CenterPlatformId   int64  `json:"centerPlatformId"`
	CenterPlatformName string `json:"centerPlatformName"`
	SkdType            int    `json:"sdkType"`
}

func handleCenterPlatformList(rw http.ResponseWriter, req *http.Request) {
	form := &centerPlatformListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetCenterPlatformList(form.CenterPlatformName, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error":              err,
			"CenterPlatformName": form.CenterPlatformName,
			"index":              form.PageIndex,
		}).Error("获取中心平台列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &centerPlatformListRespon{}
	respon.ItemArray = make([]*centerPlatformListResponItem, 0)

	for _, value := range rst {
		item := &centerPlatformListResponItem{
			CenterPlatformId:   value.PlatformId,
			CenterPlatformName: value.Name,
			SkdType:            value.SkdType,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetCenterPlatformCount(form.CenterPlatformName)
	if err != nil {
		log.WithFields(log.Fields{
			"error":              err,
			"CenterPlatformName": form.CenterPlatformName,
			"index":              form.PageIndex,
		}).Error("获取中心平台列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
