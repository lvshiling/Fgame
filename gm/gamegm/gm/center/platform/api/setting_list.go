package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerPlatformSettingListRequest struct {
	PageIndex          int    `form:"pageIndex" json:"pageIndex"`
	CenterPlatformName string `form:"centerPlatformName" json:"centerPlatformName"`
}

type centerPlatformSettingListRespon struct {
	ItemArray  []*centerPlatformSettingListResponItem `json:"itemArray"`
	TotalCount int                                    `json:"total"`
}

type centerPlatformSettingListResponItem struct {
	CenterPlatformId   int64  `json:"centerPlatformId"`
	CenterPlatformName string `json:"centerPlatformName"`
	Setting            string `json:"setting"`
}

func handleCenterPlatformSettingList(rw http.ResponseWriter, req *http.Request) {
	form := &centerPlatformSettingListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterPlatformSettingList，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterPlatformSettingList，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rst, err := service.GetCenterPlatformSettingList(form.CenterPlatformName, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterPlatformSettingList，获取列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	total, err := service.GetCenterPlatformSettingCount(form.CenterPlatformName)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterPlatformSettingList，获取个数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &centerPlatformSettingListRespon{}
	for _, value := range rst {
		item := &centerPlatformSettingListResponItem{
			CenterPlatformId:   value.PlatformId,
			Setting:            value.SettingContent,
			CenterPlatformName: value.PlatformName,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}
	respon.TotalCount = total

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
