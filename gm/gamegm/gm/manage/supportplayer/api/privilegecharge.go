package api

import (
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	gmError "fgame/fgame/gm/gamegm/error"
	gmerr "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmerrutil "fgame/fgame/gm/gamegm/error/utils"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	poolservice "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type privilegeChargeRequest struct {
	ChannelId  int    `json:"channelId"`
	PlatformId int    `json:"platformId"`
	ServerId   int32  `json:"serverId"`
	PlayerId   string `json:"playerId"`
	PlayerName string `json:"playerName"`
	Gold       int64  `json:"gold"`
	Num        int32  `json:"num"`
	Reason     string `json:"reason"`
}

func handlePrivilegeCharge(rw http.ResponseWriter, req *http.Request) {
	log.Debug("扶持玩家")
	form := &privilegeChargeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("扶持玩家，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	poolService := poolservice.ServerSupportPoolInContext(req.Context())
	if poolService == nil {
		log.Error("扶持玩家，pool服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("扶持玩家，Remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	poolInfo, err := poolService.GetServerSupportPoolInfo(form.ServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": form.ServerId,
			"error":    err,
		}).Error("获取池信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if poolInfo == nil {
		err = gmerr.GetError(gmerr.ErrorCodeServerSupportPoolNotExists)
		gmerrutil.ResponseWithError(rw, err)
		return
	}

	gold := form.Gold * int64(form.Num)
	//池不够
	if int64(poolInfo.CurGold) < gold {
		err = gmerr.GetError(gmerr.ErrorCodeServerSupportPoolNotEnought)
		gmerrutil.ResponseWithError(rw, err)
		return
	}
	//开始发
	playerId := changeStringToInt64(form.PlayerId)
	err = remoteService.PrivilegeCharge(form.ServerId, playerId, form.Gold, form.Num)
	if err != nil {
		log.WithFields(log.Fields{
			"serverId":      form.ServerId,
			"playerid":      playerId,
			"form.PlayerId": form.PlayerId,
			"gold":          form.Gold,
			"error":         err,
		}).Error("Remote发送池异常")
		// rr := gmhttp.NewFailedResultWithMsg(1000, err.Error())
		// httputils.WriteJSON(rw, http.StatusOK, rr)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}
	log.Debug("remote发送完成...")
	err = poolService.ReduceServerGold(form.ServerId, int(gold))
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": form.ServerId,
			"gold":     form.Gold,
			"error":    err,
		}).Error("减少本地池失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := supportplayer.SupportPlayerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("SupportPlayerService服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerservice := centerserver.CenterServerServiceInContext(req.Context())
	serverInfo, err := centerservice.GetCenterServer(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": form.ServerId,
			"error":    err,
		}).Error("获取中心服务器信息失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	serverName := serverInfo.ServerName
	centerPlatformId := serverInfo.Platform
	userid := gmUserService.GmUserIdInContext(req.Context())
	us := gmUserService.GmUserServiceInContext(req.Context())
	userInfo, err := us.GetUserInfo(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"userid": userid,
			"error":  err,
		}).Error("获取用户信息失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userName := userInfo.UserName

	err = service.AddChargeLog(form.ChannelId, form.PlatformId, int(centerPlatformId), int(form.ServerId), playerId, serverName, int(gold), userName, form.Reason, form.PlayerName)
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": form.ServerId,
			"error":    err,
		}).Error("添加日志失败")
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
