package api

import (
	"net/http"

	mailservice "fgame/fgame/gm/gamegm/gm/manage/mail/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"

	platform "fgame/fgame/gm/gamegm/gm/platform/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type addMailRequest struct {
	MailType         int    `json:"mailType"`
	ServerId         int    `json:"serverId"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Playerlist       string `json:"playerlist"`
	Proplist         string `json:"proplist"`
	FreezTime        int    `json:"freezTime"`
	EffectDays       int    `json:"effectDays"`
	RoleStartTime    int64  `json:"roleStartTime"`
	RoleEndTime      int64  `json:"roleEndTime"`
	MinLevel         int    `json:"minLevel"`
	MaxLevel         int    `json:"maxLevel"`
	SdkType          int    `json:"sdkType"`
	CenterPlatformId int64  `json:"centerPlatformId"`
	PlatformId       int    `json:"platformId"`
	ChannelId        int    `json:"channelId"`
	BindFlag         int    `json:"bindFlag"`
	Remark           string `json:"remark"`
}

func handleAddMail(rw http.ResponseWriter, req *http.Request) {
	form := &addMailRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加邮件，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	mailService := mailservice.MailServiceInContext(req.Context())
	if mailService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加邮件，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	platformService := platform.PlatformServiceInContext(req.Context())
	if platformService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加邮件，平台服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	serverService := centerserver.CenterServerServiceInContext(req.Context())
	if serverService == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加邮件，中心服务器服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userid := gmUserService.GmUserIdInContext(req.Context())
	serverId := make([]int, 0)
	centerPlatformList := make([]int64, 0)
	if form.ServerId > 0 {
		serverId = append(serverId, form.ServerId)
		centerPlatformList = append(centerPlatformList, form.CenterPlatformId)
	}

	if form.ServerId == 0 && form.PlatformId > 0 {
		serverList, err := serverService.GetCenterServerListByPlatformEnable(int(form.CenterPlatformId))
		if err != nil {
			log.WithFields(log.Fields{
				"centerPlatformId": form.CenterPlatformId,
				"error":            err,
			}).Error("添加邮件，获取服务器列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, value := range serverList {
			serverId = append(serverId, int(value.Id))
			centerPlatformList = append(centerPlatformList, value.Platform)
		}
	}

	if form.ServerId == 0 && form.PlatformId == 0 && form.ChannelId > 0 {
		platformList, err := platformService.GetPlatformByChannel(int64(form.ChannelId))
		if err != nil {
			log.WithFields(log.Fields{
				"channelid": form.ChannelId,
				"error":     err,
			}).Error("添加邮件，获取平台列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		queryCenterPlatList := make([]int, 0)
		for _, value := range platformList {
			queryCenterPlatList = append(queryCenterPlatList, int(value.CenterPlatformID))
		}
		serverList, err := serverService.GetCenterServerListByPlatformArrayEnable(queryCenterPlatList)
		if err != nil {
			log.WithFields(log.Fields{
				"platformArray": queryCenterPlatList,
				"error":         err,
			}).Error("添加邮件，获取服务器列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, serValue := range serverList {
			serverId = append(serverId, int(serValue.Id))
			centerPlatformList = append(centerPlatformList, serValue.Platform)
		}
	}
	mailState := int32(1)
	if len(form.Proplist) == 0 {
		mailState = int32(2)
	}
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	sendStateArray := make([]int32, 0)
	remoteErrFlag := false
	for _, value := range serverId {
		if len(form.Proplist) != 0 || form.MailType != 1 {
			sendStateArray = append(sendStateArray, int32(0))
			continue
		}

		if form.MailType == 1 {
			err = remoteService.SendPlayerCompensate(int32(value), form.Playerlist, form.Title, form.Content, form.Proplist, form.BindFlag)
			if err != nil {
				log.WithFields(log.Fields{
					"ServerId":   value,
					"Playerlist": form.Playerlist,
					"Title":      form.Title,
					"Content":    form.Content,
					"Proplist":   form.Proplist,
					"error":      err,
				}).Error("审核邮件异常,remote发送失败")
				sendStateArray = append(sendStateArray, int32(0))
				remoteErrFlag = true
			} else {
				sendStateArray = append(sendStateArray, int32(1))
			}
		} else {
			err = remoteService.SendServerCompensate(int32(value), form.Title, form.Content, form.Proplist, int32(form.MinLevel), form.RoleStartTime, form.BindFlag)
			if err != nil {
				log.WithFields(log.Fields{
					"ServerId":   value,
					"Playerlist": form.Playerlist,
					"Title":      form.Title,
					"Content":    form.Content,
					"Proplist":   form.Proplist,
					"error":      err,
				}).Error("审核邮件异常,remote发送失败")
				sendStateArray = append(sendStateArray, int32(0))
				remoteErrFlag = true
			} else {
				sendStateArray = append(sendStateArray, int32(1))
			}
		}
	}

	err = mailService.AddMailInfo(form.MailType, serverId, form.Title, form.Content, form.Playerlist, form.Proplist, form.FreezTime, form.EffectDays, form.RoleStartTime, form.RoleEndTime, form.MinLevel, form.MaxLevel, userid, form.SdkType, centerPlatformList, form.BindFlag, form.Remark, mailState, sendStateArray)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加邮件异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if remoteErrFlag {
		rr := gmhttp.NewFailedResultWithMsg(100, "添加成功，但是有部分远程发送失败")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)

}
