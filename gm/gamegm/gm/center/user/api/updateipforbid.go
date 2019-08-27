package api

import (
	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	"fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	remoteservice "fgame/fgame/gm/gamegm/remote/service"
	userremote "fgame/fgame/gm/gamegm/remote/service"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateIpForbidRequest struct {
	Ip         string `json:"ip"`
	Forbid     int    `json:"forbid"`
	ForbidTime int64  `json:"forbidTime"`
	ForbidText string `json:"reason"`
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	PlayerId   string `json:"playerId"`
}

func handleIpUpdateForbid(rw http.ResponseWriter, req *http.Request) {
	form := &updateIpForbidRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新禁止账户，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithFields(log.Fields{
		"forbidTime": form.ForbidTime,
		"forbid":     form.Forbid,
	}).Debug("解析参数")

	now := timeutils.TimeToMillisecond(time.Now())
	forbidEndTime := now + form.ForbidTime*int64(time.Second/time.Millisecond)
	if form.ForbidTime == 0 {
		forbidEndTime = 0
	}

	us := gmUserService.GmUserServiceInContext(req.Context())
	userid := gmUserService.GmUserIdInContext(req.Context())
	if us == nil {
		log.Error("更新禁止账户，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userInfo, err := us.GetUserInfo(userid)
	if err != nil || userInfo == nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userid": userid,
		}).Error("更新禁止账户，获取用户信息失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rs := remoteservice.CenterServiceInContext(req.Context())
	err = rs.ForbidIp(form.Ip, true)
	if err != nil {
		rr := gmhttp.NewFailedResultWithMsg(1000, err.Error())
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	//以下更新数据库其实可以不要
	rds := service.CenterUserServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新禁止账户，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rds.UpdateIpForbid(form.Ip, form.Forbid, now, forbidEndTime, userInfo.UserName, form.ForbidText)
	if err != nil {
		log.WithFields(log.Fields{
			"error":         err,
			"Ip":            form.Ip,
			"Forbid":        form.Forbid,
			"ForbidTime":    now,
			"ForbidEndTime": forbidEndTime,
			"ForbidName":    userInfo.UserName,
			"ForbidText":    form.ForbidText,
		}).Error("更新禁止IP失败")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
	}

	//踢掉这个人
	usRemote := userremote.UserRemoteServiceInContext(req.Context())
	playerId, _ := strconv.ParseInt(form.PlayerId, 10, 64)
	centerService := monitor.CenterServerServiceInContext(req.Context())
	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)
	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("封禁ip，踢出玩家，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = usRemote.KickOut(int32(serverid), playerId, form.ForbidText)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"serverid": serverid,
		}).Error("封禁ip，踢出玩家，踢出玩家异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}
	log.Debug("封禁IP，踢出成功")

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
