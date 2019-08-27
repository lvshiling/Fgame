package api

import (
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	jyzqservice "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	jyzqmodel "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/model"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type jiaoYiZhanQuUpdateRequest struct {
	Id         int32  `json:"id"`
	ServerId   int32  `json:"serverId"`
	ZhanQuName string `json:"zhanquName"`
	PlatformId int32  `form:"platformId" json:"platformId"`
}

func handleJiaoYiZhanQuUpdate(rw http.ResponseWriter, req *http.Request) {
	form := &jiaoYiZhanQuUpdateRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区更新，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := jyzqservice.JiaoYiZhanQuServiceInContext(req.Context())
	info := &jyzqmodel.JiaoYiZhanQuEntity{}
	now := timeutils.TimeToMillisecond(time.Now())
	info.Id = form.Id
	info.ServerId = form.ServerId
	info.PlatformId = form.PlatformId
	info.JiaoYiName = form.ZhanQuName
	info.CreateTime = now
	info.UpdateTime = now
	err = service.UpdateJiaoYiZhanQu(info)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区更新，更新战区信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
