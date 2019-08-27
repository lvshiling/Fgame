package api

import (
	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	userremotemodel "fgame/fgame/gm/gamegm/remote/model"
	userremote "fgame/fgame/gm/gamegm/remote/service"

	monitor "fgame/fgame/gm/gamegm/monitor"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type ignoreRequest struct {
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	Reason     string `json:"reason"`
	PlayerId   string `json:"playerId"`
	ForbidTime int64  `json:"forbidTime"`
}

func handleIgnoreChat(rw http.ResponseWriter, req *http.Request) {
	form := &ignoreRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("玩家禁默，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := userremote.UserRemoteServiceInContext(req.Context())
	if service == nil {
		log.Error("玩家禁默，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	//获得用户信息：这边先从库拿之后再优化
	us := gmUserService.GmUserServiceInContext(req.Context())
	userid := gmUserService.GmUserIdInContext(req.Context())
	if us == nil {
		log.Error("玩家禁默，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userInfo, err := us.GetUserInfo(userid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userid": userid,
		}).Error("玩家禁默，获取用户信息失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	info := &userremotemodel.IgnorePlayerChat{
		PlayerId:     changeStringToInt64(form.PlayerId),
		ForbidReason: form.Reason,
		ForbidName:   userInfo.UserName,
		ForbidTime:   form.ForbidTime * 1000,
	}

	centerService := monitor.CenterServerServiceInContext(req.Context())

	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)
	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("玩家禁默，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithFields(log.Fields{
		"PlayerId":     info.PlayerId,
		"ForbidReason": info.ForbidReason,
		"ForbidName":   info.ForbidName,
		"ForbidTime":   info.ForbidTime,
	}).Debug("玩家禁默认")

	err = service.IgnorePlayerChat(int32(serverid), info)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"serverid": serverid,
		}).Error("玩家禁默，玩家禁默异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
