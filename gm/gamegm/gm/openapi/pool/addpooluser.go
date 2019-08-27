package pool

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	gmerror "fgame/fgame/gm/gamegm/error"
	gmerrutil "fgame/fgame/gm/gamegm/error/utils"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	userservice "fgame/fgame/gm/gamegm/gm/center/user/service"
	playservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	poollogic "fgame/fgame/gm/gamegm/gm/manage/supportplayer/logic"
	paramapi "fgame/fgame/gm/gamegm/gm/openapi/param"
	plservice "fgame/fgame/gm/gamegm/gm/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	gmutils "fgame/fgame/gm/gamegm/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type openApiAddSupportPlayerRespon struct {
	UserId string `json:"userId"`
}

func handleAddSupportPlayer(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleAddSupportPlayer...")
	postData := paramapi.ApiDataInContext(req.Context())
	if postData == nil {
		err := gmerror.GetError(gmerror.ErrorCodeOpenApiParam)
		gmerrutil.ResponseWithError(rw, err)
		return
	}

	logicParam := &poollogic.LogicAddSupportPlayerParam{}
	platformInfo := plservice.ApiPlatformInfoInContext(req.Context())
	logicParam.ChannelId = int(platformInfo.ChannelId)
	logicParam.CenterPlatformId = int(platformInfo.CenterPlatformID)
	logicParam.PlatformId = int(platformInfo.PlatformID)
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
		}).Error("OpenApi handleAddSupportPlayer:获取服务器信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logicParam.ServerKeyId = int32(serverInfo.Id)
	userId := postData["userId"]

	usService := userservice.CenterUserServiceInContext(req.Context())
	userInfo, err := usService.GetUserInfoByUserName(platformInfo.SdkType, userId)
	if err != nil || userInfo == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("OpenApi handleAddSupportPlayer:获取中心用户信息异常")
		gmerrutil.ResponseWithError(rw, gmerror.GetError(gmerror.ErrorCodeOpenApiSupportUserNotExists))
		return
	}
	pls := playservice.PlayerServiceInContext(req.Context())
	gmInfo, err := pls.GetOriginPlayerInfoByUserId(gmdb.GameDbLink(logicParam.ServerKeyId), int(serverId), int64(userInfo.Id))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("OpenApi handleAddSupportPlayer:获取服务器用户信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logicParam.PlayerId = gmInfo.Id
	err = poollogic.HandleAddSupportPlayer(rw, req, logicParam)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("OpenApi handleAddSupportPlayer:设置扶持账号异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiRepson := &openApiAddSupportPlayerRespon{
		UserId: userId,
	}
	rr := gmhttp.NewSuccessResult(apiRepson)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
