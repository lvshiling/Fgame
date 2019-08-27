package api

import (
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	gmerror "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	jyzqservice "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type jiaoYiZhanQuAddRequest struct {
	ServerId   int32  `json:"serverId"`
	ZhanQuName string `json:"zhanquName"`
	PlatformId int32  `form:"platformId" json:"platformId"`
}

func handleJiaoYiZhanQuAdd(rw http.ResponseWriter, req *http.Request) {
	form := &jiaoYiZhanQuAddRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区添加，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := jyzqservice.JiaoYiZhanQuServiceInContext(req.Context())
	info, err := service.GetJiaoYiZhanQu(form.PlatformId, form.ServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区添加，获取战区信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if info.Id > 0 {
		errhttp.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeJiaoYiZhanQuExists))
		return
	}
	now := timeutils.TimeToMillisecond(time.Now())
	info.ServerId = form.ServerId
	info.PlatformId = form.PlatformId
	info.JiaoYiName = form.ZhanQuName
	info.CreateTime = now
	err = service.AddJiaoYiZhanQu(info)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易战区添加，添加战区信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
