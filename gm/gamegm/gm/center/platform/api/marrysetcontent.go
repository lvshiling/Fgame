package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type marrySetContentRequest struct {
	CenterPlatformId int64                         `json:"centerPlatformId"`
	MarrySet         []*marrySetContentRequestItem `json:"marrySet"`
}

type marrySetContentRequestItem struct {
	KindType     int32 `json:"kindType"`     //1当前版本，2廉价版本
	MarryType    int32 `json:"marryType"`    //对应结婚的type：1婚宴2喜糖3婚车3婚戒
	MarrySubType int32 `json:"marrySubType"` //对应结婚模板的SubType,每种类型的三种档次,从1开始
	UseGold      int32 `json:"useGold"`
}

func handleCenterPlatformMarrySetContent(rw http.ResponseWriter, req *http.Request) {
	form := &marrySetContentRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置内容，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置内容，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceContentInfo, err := service.GetPlatformMarrySet(form.CenterPlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置内容，获取数据为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceContentInfo.PlatformId = form.CenterPlatformId
	// contentString, err := json.Marshal(form.MarrySet)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"error": err,
	// 	}).Error("获取中心平台列表结婚配置设置内容，序列化内容异常")
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// priceContentInfo.PriceContent = string(contentString)
	err = service.SavePlatformMarrySet(priceContentInfo)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置内容，设置数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
