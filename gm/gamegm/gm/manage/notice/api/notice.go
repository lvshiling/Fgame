package api

import (
	"net/http"
	"time"

	centerServermodel "fgame/fgame/gm/gamegm/gm/center/model"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	ntmodel "fgame/fgame/gm/gamegm/gm/manage/notice/model"
	ntservice "fgame/fgame/gm/gamegm/gm/manage/notice/service"
	plservice "fgame/fgame/gm/gamegm/gm/platform/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type noticeRequest struct {
	ChannelId    int    `json:"channelId"`
	PlatformId   int    `json:"platformId"`
	ServerId     int    `json:"serverId"`
	Content      string `json:"content"`
	IntervalTime int64  `json:"intervalTime"`
	BeginTime    int64  `json:"beginTime"`
	EndTime      int64  `json:"endTime"`
}

func handleNotice(rw http.ResponseWriter, req *http.Request) {
	form := &noticeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("发送公告，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid := gmUserService.GmUserIdInContext(req.Context())
	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("发送公告，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userInfo, err := usservice.GetUserInfo(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("发送公告，获取用户信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userInfo.PlatformId > 0 {
		form.PlatformId = int(userInfo.PlatformId)
	}
	if userInfo.ChannelID > 0 {
		form.ChannelId = int(userInfo.ChannelID)
	}

	service := ntservice.NoticeServiceInContext(req.Context())
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("发送公告，Remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerservice := centerserver.CenterServerServiceInContext(req.Context())
	allServer, err := centerservice.GetAllCenterServerListEnable()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("发送公告，获取所有的服务器异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serverMap := make(map[int64]*centerServermodel.CenterServer)
	for _, value := range allServer {
		serverMap[value.Id] = value
	}

	ps := plservice.PlatformServiceInContext(req.Context())
	if ps == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("发送公告，平台服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	realServer := make([]int, 0)
	if form.ServerId > 0 {
		realServer = append(realServer, form.ServerId)
	}
	if form.PlatformId > 0 && form.ServerId == 0 {
		plinfo, err := ps.GetPlatformInfo(int64(form.PlatformId))
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"platid": form.PlatformId,
			}).Error("发送公告，获取平台信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		centerPlatId := plinfo.CenterPlatformID
		for _, value := range allServer {
			if value.Platform == centerPlatId && value.ServerType == 0 {
				realServer = append(realServer, int(value.Id))
			}
		}
	}
	if form.ChannelId > 0 && form.PlatformId == 0 && form.ServerId == 0 {
		platArray, err := ps.GetPlatformByChannel(int64(form.ChannelId))
		if err != nil {
			log.WithFields(log.Fields{
				"error":     err,
				"ChannelId": form.ChannelId,
			}).Error("发送公告，通过渠道获取平台信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		centerMap := make(map[int64]int64)
		for _, platValue := range platArray {
			centerMap[platValue.CenterPlatformID] = platValue.PlatformID
		}
		for _, value := range allServer {
			if value.ServerType != 0 {
				continue
			}

			if _, ok := centerMap[value.Platform]; ok {
				realServer = append(realServer, int(value.Id))
			}
		}
	}

	userName := userInfo.UserName

	for _, value := range realServer {
		log.Debug("发送服务器：", value)
		serverName := ""
		centerPlatformId := int64(0)
		if myserver, ok := serverMap[int64(value)]; ok {
			serverName = myserver.ServerName
			centerPlatformId = myserver.Platform
		}
		info := &ntmodel.NoticeInfo{
			ChannelId:        form.ChannelId,
			PlatformId:       form.PlatformId,
			ServerId:         form.ServerId,
			BeginTime:        form.BeginTime,
			EndTime:          form.EndTime,
			IntervalTime:     form.IntervalTime,
			Content:          form.Content,
			CreateTime:       timeutils.TimeToMillisecond(time.Now()),
			SuccessFlag:      1,
			ServerName:       serverName,
			CenterPlatformId: centerPlatformId,
			UserName:         userName,
		}
		err = remoteService.BroadcastNotice(int32(value), form.Content, form.BeginTime, form.EndTime, form.IntervalTime*60*1000)
		info.UpdateTime = timeutils.TimeToMillisecond(time.Now())
		if err != nil {
			info.SuccessFlag = 0
			info.ErrorMsg = err.Error()
		}
		adderr := service.AddNoticeInfo(info)
		if adderr != nil {
			log.WithFields(log.Fields{
				"error": adderr,
			}).Error("发送公告，保存信息日志")
		}
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
