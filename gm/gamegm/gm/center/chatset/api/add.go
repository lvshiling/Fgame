package api

import (
	"net/http"

	gmchatSet "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type chatSetRequest struct {
	PlatformId       int   `form:"centerPlatformId" json:"centerPlatformId"`
	ServerId         []int `form:"centerServerId" json:"centerServerId"`
	WorldVip         int   `form:"worldVip" json:"worldVip"`
	WorldPlayerLevel int   `form:"worldPlayerLevel" json:"worldPlayerLevel"`
	PChatVip         int   `form:"pChatVip" json:"pChatVip"`
	PChatPlayerLevel int   `form:"pChatPlayerLevel" json:"pChatPlayerLevel"`
	GuildVip         int   `form:"guildVip" json:"guildVip"`
	GuildPlayerLevel int   `form:"guildPlayerLevel" json:"guildPlayerLevel"`
	TeamVip          int   `form:"teamVip" json:"teamVip"`
	TeamPlayerLevel  int   `form:"teamPlayerLevel" json:"teamPlayerLevel"`
	SdkType          int   `form:"sdkType" json:"sdkType"`
}

func handleAddChatSet(rw http.ResponseWriter, req *http.Request) {
	form := &chatSetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加聊天配置，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchatSet.ChatSetServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加聊天配置,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	centerService := monitor.CenterServerServiceInContext(req.Context())
	centerServerService := centerserver.CenterServerServiceInContext(req.Context())
	errMsg := ""
	for _, value := range form.ServerId {
		valueFormServerId := value
		serverid := centerService.GetCenterServerDBId(int32(form.PlatformId), int32(valueFormServerId))
		if serverid < 1 {
			log.WithFields(log.Fields{
				"PlatformId": form.PlatformId,
				"ServerId":   form.ServerId,
			}).Error("添加聊天配置，获得服务器id为空")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		remoteService := userremote.UserRemoteServiceInContext(req.Context())
		if remoteService == nil {
			log.Error("添加聊天配置，remote服务为空")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		serverInfo, err := centerServerService.GetCenterServer(serverid)
		if err != nil {
			log.WithFields(log.Fields{
				"serverId": serverid,
			}).Error("添加聊天配置，获取服务器信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = remoteService.ChatSet(int32(serverid), int32(form.WorldVip), int32(form.WorldPlayerLevel), int32(form.GuildVip), int32(form.GuildPlayerLevel), int32(form.PChatVip), int32(form.PChatPlayerLevel), int32(form.TeamVip), int32(form.TeamPlayerLevel))
		if err != nil {
			log.WithFields(log.Fields{
				"error":              err,
				"serveridkey":        serverid,
				"worldvip":           form.WorldVip,
				"WorldPlayerLevel":   form.WorldPlayerLevel,
				"allianceVip":        form.GuildVip,
				"alliancePlayLevel":  form.GuildPlayerLevel,
				"privateVip":         form.PChatVip,
				"privatePlayerLevel": form.PChatPlayerLevel,
			}).Error("发送服务器设置异常")
			// codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
			// errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
			errMsg += serverInfo.ServerName + ","
			continue
		}
		log.Debug("remote聊天设置成功")

		chatInfo, err := service.GetChatSet(form.PlatformId, valueFormServerId)
		if err != nil {
			log.WithFields(log.Fields{
				"error":             err,
				"PlatformId":        form.PlatformId,
				"valueFormServerId": valueFormServerId,
			}).Error("设置聊天配置,获取聊天配置失败")
			errMsg += serverInfo.ServerName + ","
			continue
		}
		if chatInfo == nil || chatInfo.Id == 0 {
			err = service.AddChatSet(form.PlatformId, valueFormServerId, form.WorldVip, form.WorldPlayerLevel, form.PChatVip, form.PChatPlayerLevel, form.GuildVip, form.GuildPlayerLevel, form.SdkType, form.TeamVip, form.TeamPlayerLevel)
		} else {
			err = service.UpdateChatSet(chatInfo.Id, form.PlatformId, valueFormServerId, form.WorldVip, form.WorldPlayerLevel, form.PChatVip, form.PChatPlayerLevel, form.GuildVip, form.GuildPlayerLevel, form.SdkType, form.TeamVip, form.TeamPlayerLevel)
		}
		if err != nil {
			// errhttp.ResponseWithError(rw, err)
			errMsg += serverInfo.ServerName + ","
			continue
		}
	}
	if len(errMsg) > 0 {
		rr := gmhttp.NewFailedResultWithMsg(100, errMsg+"设置失败")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
