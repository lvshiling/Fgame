package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fmt"
	"net/http"
	"strings"

	gmerr "fgame/fgame/gm/gamegm/error"
	gmerrutil "fgame/fgame/gm/gamegm/error/utils"
	centerserver "fgame/fgame/gm/gamegm/gm/center/server/service"
	poolservice "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type privilegeChargeMulityRequest struct {
	ChannelId  int    `json:"channelId"`
	PlatformId int    `json:"platformId"`
	ServerId   int32  `json:"serverId"`
	PlayerId   string `json:"playerId"`
	PlayerName string `json:"playerName"`
	Gold       int64  `json:"gold"`
	Num        int32  `json:"num"`
	Reason     string `json:"reason"`
	AllServer  bool   `json:"allServer"`
}

func handlePrivilegeChargeMulity(rw http.ResponseWriter, req *http.Request) {
	log.Debug("扶持玩家")
	form := &privilegeChargeMulityRequest{}
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

	service := supportplayer.SupportPlayerServiceInContext(req.Context())
	// centerService := monitor.CenterServerServiceInContext(req.Context())
	// , err := centerService.GetServerId(int64(form.ServerId))
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
	acServerId := serverInfo.ServerId
	serverName := serverInfo.ServerName
	centerPlatformId := serverInfo.Platform

	playArray := make([]int64, 0)
	if form.AllServer { //全服更新
		allGm, err := service.GetAllGmPlayerList(gmdb.GameDbLink(form.ServerId), acServerId)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("扶持玩家设置，获取所有扶持玩家异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, value := range allGm {
			playArray = append(playArray, value.Id)
		}
	}
	if !form.AllServer {
		if len(form.PlayerId) == 0 && len(form.PlayerName) == 0 {
			rr := gmhttp.NewFailedResultWithMsg(112, "玩家信息不能为空")
			httputils.WriteJSON(rw, http.StatusOK, rr)
			return
		}

		formPlayArray := strings.Split(form.PlayerId, ",")
		playNameArray := strings.Split(form.PlayerName, ",")
		if len(formPlayArray) > 0 && len(form.PlayerId) > 0 {
			for _, value := range formPlayArray {
				formPlayerId := changeStringToInt64(value)
				dbPlayerid, dberr := service.CheckPlayerId(gmdb.GameDbLink(form.ServerId), acServerId, formPlayerId)
				if dberr != nil {
					log.WithFields(log.Fields{
						"playerid": value,
						"error":    dberr,
					}).Error("扶持玩家设置，获取玩家id异常")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				if dbPlayerid <= 0 {
					rr := gmhttp.NewFailedResultWithMsg(113, fmt.Sprintf("玩家信息异常，玩家ID:%d不存在", formPlayerId))
					httputils.WriteJSON(rw, http.StatusOK, rr)
					return
				}
				fuchFlag, dberr := service.CheckPlayerIdFuchi(gmdb.GameDbLink(form.ServerId), dbPlayerid)
				if dberr != nil {
					log.WithFields(log.Fields{
						"playerid": value,
						"error":    dberr,
					}).Error("扶持玩家设置，获取玩家id异常")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				if !fuchFlag {
					rr := gmhttp.NewFailedResultWithMsg(113, fmt.Sprintf("玩家信息异常，玩家ID:%d不是扶持账号", formPlayerId))
					httputils.WriteJSON(rw, http.StatusOK, rr)
					return
				}
				if dbPlayerid > 0 {
					playArray = append(playArray, dbPlayerid)
				}
			}
		}
		if len(playArray) == 0 {
			for _, value := range playNameArray {
				dbPlayerid, dberr := service.GetPlayerId(gmdb.GameDbLink(form.ServerId), acServerId, value)
				if dberr != nil {
					log.WithFields(log.Fields{
						"PlayerName": value,
						"error":      dberr,
					}).Error("扶持玩家设置，获取玩家id异常")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				if dbPlayerid <= 0 {
					rr := gmhttp.NewFailedResultWithMsg(113, fmt.Sprintf("玩家信息异常，玩家名:%s不存在", value))
					httputils.WriteJSON(rw, http.StatusOK, rr)
					return
				}
				fuchFlag, dberr := service.CheckPlayerIdFuchi(gmdb.GameDbLink(form.ServerId), dbPlayerid)
				if dberr != nil {
					log.WithFields(log.Fields{
						"playerid": value,
						"error":    dberr,
					}).Error("扶持玩家设置，获取玩家id异常")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				if !fuchFlag {
					rr := gmhttp.NewFailedResultWithMsg(113, fmt.Sprintf("玩家信息异常，玩家名:%s不是扶持账号", value))
					httputils.WriteJSON(rw, http.StatusOK, rr)
					return
				}
				if dbPlayerid > 0 {
					playArray = append(playArray, dbPlayerid)
				}
			}
		}
	}

	if len(playArray) == 0 {
		rr := gmhttp.NewFailedResultWithMsg(112, "玩家信息异常，不存在相关的玩家Id和玩家名字")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}
	num := int64(len(playArray))
	totalCost := form.Gold * num

	//池不够
	if int64(poolInfo.CurGold) < totalCost {
		err = gmerr.GetError(gmerr.ErrorCodeServerSupportPoolNotEnought)
		gmerrutil.ResponseWithError(rw, err)
		return
	}

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

	//开始发
	// playerId := changeStringToInt64(form.PlayerId)
	errPlayerId := make([]int64, 0)
	for _, value := range playArray {
		err = remoteService.PrivilegeCharge(form.ServerId, value, form.Gold, form.Num)
		if err != nil {
			log.WithFields(log.Fields{
				"serverId":      form.ServerId,
				"playerid":      value,
				"form.PlayerId": form.PlayerId,
				"gold":          form.Gold,
				"error":         err,
			}).Error("Remote发送池异常")
			errPlayerId = append(errPlayerId, value)
			continue
		}
		gold := form.Gold * int64(form.Num)
		playerName, _ := service.GetPlayerName(gmdb.GameDbLink(form.ServerId), value)
		err = service.AddChargeLog(form.ChannelId, form.PlatformId, int(centerPlatformId), int(form.ServerId), value, serverName, int(gold), userName, form.Reason, playerName)
		if err != nil {
			log.WithFields(log.Fields{
				"serverId": form.ServerId,
				"error":    err,
			}).Error("添加日志失败")
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
	errorCount := int64(len(errPlayerId))
	gold := form.Gold * int64(form.Num)
	totalCost = gold * int64(num-errorCount)
	log.Debug("remote发送完成...")
	err = poolService.ReduceServerGold(form.ServerId, int(totalCost))
	if err != nil {
		log.WithFields(log.Fields{
			"serverId": form.ServerId,
			"gold":     form.Gold,
			"error":    err,
		}).Error("减少本地池失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(errPlayerId) > 0 {
		errmsg := fmt.Sprintf("玩家Id列表：%d发送失败", errPlayerId)
		rr := gmhttp.NewFailedResultWithMsg(112, errmsg)
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
