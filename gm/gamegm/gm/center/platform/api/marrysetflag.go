package api

import (
	gmcenterPlatform "fgame/fgame/gm/gamegm/gm/center/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	centerservice "fgame/fgame/gm/gamegm/gm/center/server/service"

	"fgame/fgame/gm/gamegm/gm/center/platform/model"
	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type marrySetFlagRequest struct {
	CenterPlatformId int64 `form:"centerPlatformId" json:"centerPlatformId"`
	KindType         int32 `json:"kindType"`
}

func handleCenterPlatformMarrySetFlag(rw http.ResponseWriter, req *http.Request) {
	form := &marrySetFlagRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置状态，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterPlatform.CenterPlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置状态，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceContentInfo, err := service.GetPlatformMarrySet(form.CenterPlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置状态，获取数据为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceContentInfo.PlatformId = form.CenterPlatformId
	priceContentInfo.KindType = form.KindType
	err = service.SavePlatformMarrySet(priceContentInfo)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置状态，设置数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	remoteCs := remoteservice.CenterServiceInContext(req.Context())
	err = remoteCs.RefreshMarryPrice(int32(form.CenterPlatformId))
	centerSuccessFlag := true
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("中心平台列表结婚配置设置状态，设置中心服异常")
		centerSuccessFlag = false
		// rr := gmhttp.NewFailedResultWithMsg(100, "刷新中心服异常")
		// httputils.WriteJSON(rw, http.StatusOK, rr)
		// return
	}

	successFlag := true
	cs := centerservice.CenterServerServiceInContext(req.Context())
	serverList, err := cs.GetCenterServerListByPlatform(int(form.CenterPlatformId))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心平台列表结婚配置设置状态，获取服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rs := remoteservice.UserRemoteServiceInContext(req.Context())
	now := timeutils.TimeToMillisecond(time.Now())
	for _, value := range serverList {
		err = rs.SetMarryBanquetHouTaiType(int32(value.Id), form.KindType)
		logInfo := &model.PlatformMarrySendLog{
			CreateTime:  now,
			UpdateTime:  now,
			PlatformId:  form.CenterPlatformId,
			ServerId:    int32(value.ServerId),
			KindType:    form.KindType,
			SuccessFlag: 1,
		}
		if err != nil {
			logInfo.SuccessFlag = 0
			logInfo.FailMsg = err.Error()
			successFlag = false
		}
		err = service.AddPlatformMarryServerLog(logInfo)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("中心平台列表结婚配置设置状态,保存日志异常")

		}
	}
	errString := ""
	if !successFlag {
		errString += "部分服务器刷新异常!"
	}
	if !centerSuccessFlag {
		errString += "中心服刷新异常"
	}
	if !successFlag || !centerSuccessFlag {
		rr := gmhttp.NewFailedResultWithMsg(100, errString)
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
