package api

import (
	jyzqservice "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type jiaoYiZhanQuListRequest struct {
	PlatformId int32 `form:"platformId" json:"platformId"`
	PageIndex  int32 `form:"pageIndex" json:"pageIndex"`
}

type jiaoYiZhanQuListRespon struct {
	ItemArray  []*jiaoYiZhanQuListResponItem `json:"itemArray"`
	TotalCount int32                         `json:"totalCount"`
}

type jiaoYiZhanQuListResponItem struct {
	Id         int32  `json:"id"`
	ServerId   int32  `json:"serverId"`
	ZhanQuName string `json:"zhanquName"`
	PlatformId int32  `json:"platformId"`
	CreateTime int64  `json:"createTime"`
}

func handleJiaoYiZhanQuList(rw http.ResponseWriter, req *http.Request) {
	form := &jiaoYiZhanQuListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := jyzqservice.JiaoYiZhanQuServiceInContext(req.Context())

	respon := &jiaoYiZhanQuListRespon{}
	respon.ItemArray = make([]*jiaoYiZhanQuListResponItem, 0)
	totalCount, err := service.GetJiaoYiZhanQuCount(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区列表，获取个数失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	list, err := service.GetJiaoYiZhanQuList(form.PlatformId, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区列表，获取列表失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, value := range list {
		item := &jiaoYiZhanQuListResponItem{
			Id:         value.Id,
			PlatformId: value.PlatformId,
			ServerId:   value.ServerId,
			ZhanQuName: value.JiaoYiName,
			CreateTime: value.CreateTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	respon.TotalCount = totalCount

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
