package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fmt"
	"net/http"
	"strings"

	poolservice "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/service"
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/service"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type privilegeSetRequest struct {
	ServerId   int32  `json:"serverId"`
	PlayerId   string `json:"playerId"`
	Privilege  int32  `json:"privilege"`
	PlayerName string `json:"playerName"`
}

func handlePrivilegeSet(rw http.ResponseWriter, req *http.Request) {
	log.Debug("扶持玩家设置")
	form := &privilegeSetRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("扶持玩家设置，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	poolService := poolservice.ServerSupportPoolInContext(req.Context())
	if poolService == nil {
		log.Error("扶持玩家设置，pool服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	remoteService := userremote.UserRemoteServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("扶持玩家设置，Remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := supportplayer.SupportPlayerServiceInContext(req.Context())
	centerService := monitor.CenterServerServiceInContext(req.Context())
	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("扶持玩家设置，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	formPlayArray := strings.Split(form.PlayerId, ",")
	playArray := make([]int64, 0)

	// formPlayerId := changeStringToInt64(form.PlayerId)
	// playerId := int64(0)
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
				rr := gmhttp.NewFailedResultWithMsg(113, fmt.Sprintf("玩家信息异常，玩家ID:%d不存在", value))
				httputils.WriteJSON(rw, http.StatusOK, rr)
				return
			}
			if dbPlayerid > 0 {
				playArray = append(playArray, dbPlayerid)
			}
		}
	}
	if len(playArray) == 0 {
		playNameArray := strings.Split(form.PlayerName, ",")
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
			if dbPlayerid > 0 {
				playArray = append(playArray, dbPlayerid)
			}
		}
	}
	if len(playArray) == 0 {
		rr := gmhttp.NewFailedResultWithMsg(112, "玩家信息异常，不存在相关的玩家Id和玩家名字")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	for _, value := range playArray {
		err = remoteService.PrivilegeChargeSet(form.ServerId, value, form.Privilege)
		if err != nil {
			log.WithFields(log.Fields{
				"playerid": form.PlayerId,
				"error":    err,
			}).Error("扶持玩家设置，remote设置异常")
			rr := gmhttp.NewFailedResultWithMsg(111, err.Error())
			httputils.WriteJSON(rw, http.StatusOK, rr)
			return
		}
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
