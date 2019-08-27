package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"

	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateForbidRequest struct {
	UserId     int    `json:"userId"`
	Forbid     int    `json:"forbid"`
	ForbidTime int64  `json:"forbidTime"`
	ForbidText string `json:"reason"`
}

func handleUpdateForbid(rw http.ResponseWriter, req *http.Request) {
	form := &updateForbidRequest{}
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

	rds := service.CenterUserServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新禁止账户，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rds.UpdateForbid(int64(form.UserId), form.Forbid, now, forbidEndTime, userInfo.UserName, form.ForbidText)
	if err != nil {
		log.WithFields(log.Fields{
			"error":         err,
			"userid":        form.UserId,
			"Forbid":        form.Forbid,
			"ForbidTime":    now,
			"ForbidEndTime": forbidEndTime,
			"ForbidName":    userInfo.UserName,
			"ForbidText":    form.ForbidText,
		}).Error("更新禁止账户失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
