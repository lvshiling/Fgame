package api

import (
	"net/http"
	"strings"

	"fgame/fgame/login/login"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type VisitorLoginForm struct {
	DeviceMac string `form:"deviceMac" json:"deviceMac"`
}

func handleVisitorLogin(rw http.ResponseWriter, req *http.Request) {
	log.Debug("游客登录")
	loginForm := &VisitorLoginForm{}
	if err := httputils.Bind(req, loginForm); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.WithFields(log.Fields{
			"error": err,
		}).Error("游客登录,解析失败")
		return
	}
	deviceMac := loginForm.DeviceMac
	deviceMac = strings.TrimSpace(deviceMac)
	//TODO 验证参数
	if len(deviceMac) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Error("游客登录,失败,设备码是空")
		return
	}

	ls := login.LoginServiceInContext(req.Context())

	t, expiredTime, err := ls.VisitLogin(deviceMac)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"deviceMac": deviceMac,
			"error":     err,
		}).Error("游客登录,失败")
		return
	}

	lr := &LoginResponse{}
	lr.Token = t
	lr.ExpireTime = expiredTime

	rr := RestResult{}
	rr.ErrorCode = 0
	rr.Result = lr
	httputils.WriteJSON(rw, http.StatusOK, rr)
	log.WithFields(
		log.Fields{
			"token":       t,
			"expiredTime": expiredTime,
		},
	).Debug("游客登录成功")
}
