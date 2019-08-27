package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateIpUnForbidRequest struct {
	Ip string `json:"ip"`
}

func handleIpUpdateUnForbid(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleIpUpdateUnForbid:解禁ip")
	form := &updateIpUnForbidRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新解封IP，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithFields(log.Fields{
		"ip": form.Ip,
	}).Debug("解析参数")

	rs := remoteservice.CenterServiceInContext(req.Context())
	err = rs.ForbidIp(form.Ip, false)
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
		}).Error("更新解封IP，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
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
	now := timeutils.TimeToMillisecond(time.Now())
	err = rds.UpdateIpForbid(form.Ip, 0, now, 0, userInfo.UserName, "")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"Ip":    form.Ip,
		}).Error("更新禁止IP失败")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
