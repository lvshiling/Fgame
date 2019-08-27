package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	playerservice "fgame/fgame/gm/gamegm/gm/game/player/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	csservice "fgame/fgame/gm/gamegm/gm/center/user/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type playerInfoRequest struct {
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	PlayerId   string `json:"playerId"`
	Ip         string `json:"ip"`
}

type playerInfoRespon struct {
	Forbid       int `json:"forbid"`
	ForbidChat   int `json:"forbidChat"`
	IgnoreChat   int `json:"ignoreChat"`
	CenterForbid int `json:"centerForbid"`
	IpForbid     int `json:"ipForbid"`
}

func handlePlayerInfo(rw http.ResponseWriter, req *http.Request) {
	form := &playerInfoRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取游戏玩家信息，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := playerservice.PlayerServiceInContext(req.Context())
	rsp := &playerInfoRespon{}

	centerService := monitor.CenterServerServiceInContext(req.Context())
	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)

	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("获取游戏玩家信息，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	playerId := changeStringToInt64(form.PlayerId)
	rst, err := service.GetPlayerInfo(gmdb.GameDbLink(serverid), playerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取玩家信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp.Forbid = rst.Forbid
	rsp.ForbidChat = rst.ForbidChat
	rsp.IgnoreChat = rst.IgnoreChat
	//中心用户获取
	centerUserService := csservice.CenterUserServiceInContext(req.Context())
	if centerUserService != nil {
		cusInfo, err := centerUserService.GetUserInfo(rst.UserId)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("获取中心玩家信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rsp.CenterForbid = cusInfo.Forbid
		ipInfo, err := centerUserService.GetIpForbidInfo(form.Ip)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("获取中心玩家信息异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rsp.IpForbid = ipInfo.Forbid
	}
	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
