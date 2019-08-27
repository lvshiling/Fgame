package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerPlatformMarrySetListRespon struct {
	ItemArray  []*centerPlatformMarrySetListResponItem `json:"itemArray"`
	TotalCount int                                     `json:"total"`
}

type centerPlatformMarrySetListResponItem struct {
	CenterPlatformId   int64  `json:"centerPlatformId"`
	CenterPlatformName string `json:"centerPlatformName"`
	KindType           int32  `json:"kindType"`
	PriceSetFlag       int32  `json:"priceSetFlag"`
	// MarrySet           string `json:"marrySet"`
}

func handleCenterPlatformMarrySetList(rw http.ResponseWriter, req *http.Request) {
	form := &centerPlatformListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetCenterPlatformList(form.CenterPlatformName, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error":              err,
			"CenterPlatformName": form.CenterPlatformName,
			"index":              form.PageIndex,
		}).Error("获取中心平台列表结婚配置列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &centerPlatformMarrySetListRespon{}
	respon.ItemArray = make([]*centerPlatformMarrySetListResponItem, 0)
	for _, value := range rst {
		item := &centerPlatformMarrySetListResponItem{
			CenterPlatformId:   value.PlatformId,
			CenterPlatformName: value.Name,
		}
		marryInfo, err := service.GetPlatformMarrySet(value.PlatformId)
		if err != nil {
			log.WithFields(log.Fields{
				"error":      err,
				"PlatformId": value.PlatformId,
			}).Error("获取中心平台结婚配置，执行异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		if marryInfo == nil || marryInfo.Id == 0 {
			item.KindType = 1
		} else {
			item.KindType = int32(marryInfo.KindType)
			// item.MarrySet = marryInfo.PriceContent
			// if len(item.MarrySet) > 0 {
			// 	item.PriceSetFlag = 1
			// }
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
