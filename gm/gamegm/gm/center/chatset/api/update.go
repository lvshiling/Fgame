package api

import (
	"net/http"

	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmchatSet "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type chatSetUpdateRequest struct {
	ChatSetId        int   `form:"chatSetId" json:"chatSetId"`
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

func handleUpdateChatSet(rw http.ResponseWriter, req *http.Request) {
	form := &chatSetUpdateRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新聊天配置，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchatSet.ChatSetServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新聊天配置,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(form.ServerId) == 0 {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	centerService := monitor.CenterServerServiceInContext(req.Context())

	updateServerId := form.ServerId[0]
	serverid := centerService.GetCenterServerDBId(int32(form.PlatformId), int32(updateServerId))
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
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}
	log.Debug("remote聊天设置成功")

	err = service.UpdateChatSet(form.ChatSetId, form.PlatformId, updateServerId, form.WorldVip, form.WorldPlayerLevel, form.PChatVip, form.PChatPlayerLevel, form.GuildVip, form.GuildPlayerLevel, form.SdkType, form.TeamVip, form.TeamPlayerLevel)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
