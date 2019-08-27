package logic

import (
	"fmt"
	"net/http"

	gmerr "fgame/fgame/gm/gamegm/error"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	poolservice "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/service"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
)

type LogicSupportPlayerParam struct {
	ChannelId        int     //gm渠道Id
	PlatformId       int     //gm平台Id
	CenterPlatformId int     //中心服平台Id
	ServerKeyId      int32   //中心服主键ID
	PlayerId         int64   //玩家ID
	ServerName       string  //服务器名称
	Gold             []int32 //扶持金额
	UserName         string  //扶持操作人
	Reason           string  //扶持原因
	PlayerName       string  //游戏玩家名称
	Num              []int32 //数量
}

func HandlePrivilegeChargeSupportPlayer(rw http.ResponseWriter, req *http.Request, param *LogicSupportPlayerParam) error {
	poolService := poolservice.ServerSupportPoolInContext(req.Context())
	if poolService == nil {
		log.Error("扶持玩家，pool服务为空")
		return fmt.Errorf("扶持玩家，pool服务为空")
	}
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("扶持玩家，Remote服务为空")
		// rw.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("扶持玩家，Remote服务为空")
	}

	poolInfo, err := poolService.GetServerSupportPoolInfo(param.ServerKeyId)
	if err != nil {
		log.WithFields(log.Fields{
			"ServerKeyId": param.ServerKeyId,
			"error":       err,
		}).Error("获取池信息异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("获取池信息异常")
	}

	if poolInfo == nil {
		err = gmerr.GetError(gmerr.ErrorCodeServerSupportPoolNotExists)
		// gmerrutil.ResponseWithError(rw, err)
		return err
	}
	totalGold := 0
	for ix := 0; ix < len(param.Gold); ix++ {
		totalGold += int(param.Gold[ix] * param.Num[ix])
	}

	//池不够
	if poolInfo.CurGold < totalGold {
		err = gmerr.GetError(gmerr.ErrorCodeServerSupportPoolNotEnought)
		// gmerrutil.ResponseWithError(rw, err)
		return err
	}

	//开始发
	playerId := param.PlayerId
	for ix := 0; ix < len(param.Gold); ix++ {
		err = remoteService.PrivilegeCharge(param.ServerKeyId, playerId, int64(param.Gold[ix]), param.Num[ix])
		if err != nil {
			log.WithFields(log.Fields{
				"serverId":       param.ServerKeyId,
				"playerid":       playerId,
				"param.PlayerId": param.PlayerId,
				"gold":           param.Gold,
				"error":          err,
			}).Error("Remote发送池异常")
			// rr := gmhttp.NewFailedResultWithMsg(1000, err.Error())
			// httputils.WriteJSON(rw, http.StatusOK, rr)
			// return
			codeErr := gmerr.GetError(gmerr.ErrorCodeDefaultRemoteUser)
			// errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
			return codeErr
		}
	}
	log.Debug("remote发送完成...")
	err = poolService.ReduceServerGold(param.ServerKeyId, int(totalGold))
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": param.ServerKeyId,
			"gold":     param.Gold,
			"error":    err,
		}).Error("减少本地池失败")
		// rw.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("减少本地池失败")
	}

	service := supportplayer.SupportPlayerServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("logic:SupportPlayerService服务为空")
		// rw.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("logic:SupportPlayerService服务为空")
	}
	centerservice := centerserver.CenterServerServiceInContext(req.Context())
	serverInfo, err := centerservice.GetCenterServer(int64(param.ServerKeyId))
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": param.ServerKeyId,
			"error":    err,
		}).Error("获取中心服务器信息失败")
		// rw.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("获取中心服务器信息失败")
	}
	serverName := serverInfo.ServerName
	centerPlatformId := serverInfo.Platform

	userName := param.UserName

	err = service.AddChargeLog(param.ChannelId, param.PlatformId, int(centerPlatformId), int(param.ServerKeyId), playerId, serverName, int(totalGold), userName, param.Reason, param.PlayerName)
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": param.ServerKeyId,
			"error":    err,
		}).Error("添加日志失败")
		// rw.WriteHeader(http.StatusInternalServerError)
	}

	// rr := gmhttp.NewSuccessResult(nil)
	// httputils.WriteJSON(rw, http.StatusOK, rr)
	return nil
}
