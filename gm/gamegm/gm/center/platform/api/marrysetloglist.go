package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type marrySetLogListRequest struct {
	PageIndex  int `form:"pageIndex" json:"pageIndex"`
	PlatformId int `form:"platformId" json:"platformId"`
}

type marrySetLogListRespon struct {
	ItemArray  []*marrySetLogListResponItem `json:"itemArray"`
	TotalCount int                          `json:"total"`
}

type marrySetLogListResponItem struct {
	Id          int32  `json:"id"`
	PlatformId  int64  `json:"centerPlatformId"`
	ServerId    int32  `json:"centerServerId"`
	SuccessFlag int32  `json:"successFlag"`
	KindType    int32  `json:"kindType"`
	FailMsg     string `json:"failMsg"`
	UpdateTime  int64  `json:"updateTime"`
}

func handleMarrySetLogList(rw http.ResponseWriter, req *http.Request) {
	form := &marrySetLogListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚日志，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚日志，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	list, err := service.GetPlatformMarryServerLogList(int64(form.PlatformId), -1, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚日志，获取中心列表日志异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &marrySetLogListRespon{}
	respon.ItemArray = make([]*marrySetLogListResponItem, 0)
	for _, value := range list {
		item := &marrySetLogListResponItem{
			Id:          int32(value.Id),
			PlatformId:  value.PlatformId,
			ServerId:    value.ServerId,
			SuccessFlag: value.SuccessFlag,
			KindType:    value.KindType,
			FailMsg:     value.FailMsg,
			UpdateTime:  value.UpdateTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetPlatformMarryServerLogCount(int64(form.PlatformId), -1)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚日志，获取中心列表日志个数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
