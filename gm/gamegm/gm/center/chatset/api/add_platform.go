package api

import (
	"net/http"

	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmchatSet "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type chatSetPlatformRequest struct {
	PlatformId       int `form:"centerPlatformId" json:"centerPlatformId"`
	WorldVip         int `form:"worldVip" json:"worldVip"`
	WorldPlayerLevel int `form:"worldPlayerLevel" json:"worldPlayerLevel"`
	PChatVip         int `form:"pChatVip" json:"pChatVip"`
	PChatPlayerLevel int `form:"pChatPlayerLevel" json:"pChatPlayerLevel"`
	GuildVip         int `form:"guildVip" json:"guildVip"`
	GuildPlayerLevel int `form:"guildPlayerLevel" json:"guildPlayerLevel"`
	TeamVip          int `form:"teamVip" json:"teamVip"`
	TeamPlayerLevel  int `form:"teamPlayerLevel" json:"teamPlayerLevel"`
}

func handleAddChatSetPlatform(rw http.ResponseWriter, req *http.Request) {
	form := &chatSetPlatformRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加平台聊天配置，解析异常")
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

	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"error":              err,
	// 		"serveridkey":        serverid,
	// 		"worldvip":           form.WorldVip,
	// 		"WorldPlayerLevel":   form.WorldPlayerLevel,
	// 		"allianceVip":        form.GuildVip,
	// 		"alliancePlayLevel":  form.GuildPlayerLevel,
	// 		"privateVip":         form.PChatVip,
	// 		"privatePlayerLevel": form.PChatPlayerLevel,
	// 	}).Error("发送服务器设置异常")
	// 	// codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
	// 	// errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
	// 	errMsg += serverInfo.ServerName + ","
	// 	continue
	// }
	log.Debug("remote聊天设置成功")

	chatInfo, err := service.GetChatSetPlatform(form.PlatformId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"PlatformId": form.PlatformId,
		}).Error("设置聊天配置,获取聊天配置失败")
	}
	if chatInfo == nil || chatInfo.Id == 0 {
		err = service.AddChatSetPlatform(form.PlatformId, form.WorldVip, form.WorldPlayerLevel, form.PChatVip, form.PChatPlayerLevel, form.GuildVip, form.GuildPlayerLevel, form.TeamVip, form.TeamPlayerLevel)
	} else {
		err = service.UpdateChatSetPlatform(chatInfo.Id, form.PlatformId, form.WorldVip, form.WorldPlayerLevel, form.PChatVip, form.PChatPlayerLevel, form.GuildVip, form.GuildPlayerLevel, form.TeamVip, form.TeamPlayerLevel)
	}
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	// centerService := monitor.CenterServerServiceInContext(req.Context())
	// centerServerService := centerserver.CenterServerServiceInContext(req.Context())
	errMsg := ""

	remoteService := remoteservice.CenterServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("添加聊天配置，remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = remoteService.RefreshPlatformConfig(int32(form.PlatformId))
	if err != nil {
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	if len(errMsg) > 0 {
		rr := gmhttp.NewFailedResultWithMsg(100, errMsg+"设置失败")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
