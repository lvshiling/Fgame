package pool

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	gmerror "fgame/fgame/gm/gamegm/error"
	gmerrutil "fgame/fgame/gm/gamegm/error/utils"
	poollogic "fgame/fgame/gm/gamegm/gm/manage/supportplayer/logic"
	paramapi "fgame/fgame/gm/gamegm/gm/openapi/param"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	plservice "fgame/fgame/gm/gamegm/gm/platform/service"

	gmutils "fgame/fgame/gm/gamegm/utils"

	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"

	playservice "fgame/fgame/gm/gamegm/gm/game/player/service"

	userservice "fgame/fgame/gm/gamegm/gm/center/user/service"

	futils "fgame/fgame/core/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleAddPoolPlayer(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleAddPoolPlayer...")
	postData := paramapi.ApiDataInContext(req.Context())
	if postData == nil {
		err := gmerror.GetError(gmerror.ErrorCodeOpenApiParam)
		gmerrutil.ResponseWithError(rw, err)
		return
	}
	platformInfo := plservice.ApiPlatformInfoInContext(req.Context())
	logicParam := &poollogic.LogicSupportPlayerParam{}
	logicParam.ChannelId = int(platformInfo.ChannelId)
	logicParam.CenterPlatformId = int(platformInfo.CenterPlatformID)
	logicParam.PlatformId = int(platformInfo.PlatformID)
	logicParam.UserName = platformInfo.PlatformName + "api"
	serverId, err := gmutils.ConverStringToInt64Error(postData["serverId"])
	if err != nil {
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiParam))
		return
	}
	centerServer := centerserver.CenterServerServiceInContext(req.Context())
	serverInfo, err := centerServer.GetCenterServerInfo(logicParam.CenterPlatformId, int(serverId))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("OpenApi handleAddPoolPlayer:获取服务器信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logicParam.ServerKeyId = int32(serverInfo.Id)
	logicParam.ServerName = serverInfo.ServerName
	gold, err := futils.SplitAsIntArray(postData["gold"])
	if err != nil {
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiParam))
		return
	}
	logicParam.Gold = gold
	userId := postData["userId"]

	usService := userservice.CenterUserServiceInContext(req.Context())
	userInfo, err := usService.GetUserInfoByUserName(platformInfo.SdkType, userId)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("OpenApi handleAddPoolPlayer:获取中心用户信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	playerService := playservice.PlayerServiceInContext(req.Context())
	playerInfo, err := playerService.GetOriginPlayerInfoByUserId(gmdb.GameDbLink(logicParam.ServerKeyId), int(serverId), int64(userInfo.Id))
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"playerId": userId,
		}).Error("OpenApi handleAddPoolPlayer:获取玩家信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if playerInfo == nil {
		log.Debug("OpenApi handleAddPoolPlayer:扶持玩家为空")
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiUserNotExists))
		return
	}

	if playerInfo.PrivilegeType == 0 {
		log.Debug("OpenApi handleAddPoolPlayer:不是扶持号")
		logicAddUser := &poollogic.LogicAddSupportPlayerParam{}
		logicAddUser.CenterPlatformId = logicParam.CenterPlatformId
		logicAddUser.ChannelId = logicParam.ChannelId
		logicAddUser.PlatformId = logicParam.PlatformId
		logicAddUser.PlayerId = playerInfo.Id
		logicAddUser.ServerKeyId = logicParam.ServerKeyId

		err = poollogic.HandleAddSupportPlayer(rw, req, logicAddUser)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("OpenApi handleAddSupportPlayer:设置扶持账号异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	logicParam.PlayerId = playerInfo.Id
	logicParam.Reason = postData["reason"]
	// = gmutils.ConverStringToInt32Error(postData["goldNum"])
	num, err := futils.SplitAsIntArray(postData["goldNum"])
	if err != nil {
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiParam))
		return
	}
	logicParam.Num = num
	if len(logicParam.Num) != len(logicParam.Gold) {
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiSupportPoolGoldNumNotEqual))
		return
	}

	logicParam.PlayerName = playerInfo.Name
	err = poollogic.HandlePrivilegeChargeSupportPlayer(rw, req, logicParam)
	if err != nil {
		gmerrutil.ResponseWithError(rw, err)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
