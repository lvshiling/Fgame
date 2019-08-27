package api

import (
	"net/http"

	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type platformListRequest struct {
	PageIndex    int    `form:"pageIndex" json:"pageIndex"`
	ChannelId    int64  `form:"channelId" json:"channelId"`
	PlatformName string `form:"platformName" json:"platformName"`
}

type platformListRespon struct {
	ItemArray  []*platformListResponItem `json:"itemArray"`
	TotalCount int                       `json:"total"`
}

type platformListResponItem struct {
	PlatformId       int64  `json:"platformId"`
	PlatformName     string `json:"platformName"`
	ChannelId        int64  `json:"channelId"`
	CenterPlatformId int64  `json:"centerPlatformId"`
	SdkType          int    `json:"sdkType"`
	SignKey          string `json:"signKey"`
}

func handlePlatformList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("获取平台列表:", req.URL.RawQuery)
	form := &platformListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取平台列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmplatform.PlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取平台列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetPlatformList(form.PlatformName, form.ChannelId, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error":        err,
			"PlatformName": form.PlatformName,
			"index":        form.PageIndex,
		}).Error("获取平台列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &platformListRespon{}
	respon.ItemArray = make([]*platformListResponItem, 0)

	for _, value := range rst {
		item := &platformListResponItem{
			PlatformId:       value.PlatformID,
			PlatformName:     value.PlatformName,
			ChannelId:        value.ChannelId,
			CenterPlatformId: value.CenterPlatformID,
			SdkType:          value.SdkType,
			SignKey:          value.SignKey,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetPlatformCount(form.PlatformName, form.ChannelId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":        err,
			"PlatformName": form.PlatformName,
			"index":        form.PageIndex,
		}).Error("获取平台列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
