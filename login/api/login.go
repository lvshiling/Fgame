package api

import (
	"net/http"

	"fgame/fgame/login/login"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type LoginForm struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
}

func handleLogin(rw http.ResponseWriter, req *http.Request) {
	log.Debug("登陆")
	loginForm := &LoginForm{}
	if err := httputils.Bind(req, loginForm); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.WithFields(log.Fields{
			"error": err,
		}).Error("登陆,解析失败")
		return
	}
	userName := loginForm.UserName
	password := loginForm.Password
	//TODO 验证参数
	if len(userName) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Error("登陆,失败,用户名是空")
		return
	}

	//TODO 验证参数
	if len(password) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		log.Error("登陆,失败,密码是空")
		return
	}

	ls := login.LoginServiceInContext(req.Context())

	t, expiredTime, err := ls.Login(userName, password)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"userName": userName,
			"password": password,
			"error":    err,
		}).Error("登陆,失败")
		return
	}
	if t == "" {
		rw.WriteHeader(http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"userName": userName,
			"password": password,
		}).Error("登陆,密码错误或用户不存在")
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
	).Debug("登陆成功")
}
