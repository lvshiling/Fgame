package api

import (
	"net/http"

	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"

	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type loginRequest struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
}

type loginRespon struct {
	UserId           int64    `json:"user_id"`
	UserName         string   `json:"name"`
	Access           []string `json:"roles"`
	Token            string   `json:"token"`
	Avator           string   `json:"avatar"`
	ExpiredTime      int64    `json:"expiredTime"`
	ChannelId        int64    `json:"channelId"`
	PlatformId       int64    `json:"platformId"`
	CenterPlatformId int64    `json:"centerPlatformId"`
}

func handleLogin(rw http.ResponseWriter, req *http.Request) {
	form := &loginRequest{}
	// postbytte, posterr := ioutil.ReadAll(req.Body)
	// log.WithFields(log.Fields{
	// 	"postdata": string(postbytte),
	// 	"error":    posterr,
	// }).Debug("post数据")

	// rw.WriteHeader(http.StatusOK)
	// return
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("用户登陆，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"userName": form.UserName,
		"password": form.Password,
	}).Debug("用户登陆")

	service := gmUserService.LoginServiceInContext(req.Context())
	if service == nil {
		log.Error("用户登陆，获取登陆服务异常")
	}

	userInfo, err := service.Login(form.UserName, form.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"userName": form.UserName,
			"password": form.Password,
			"error":    err,
		}).Error("用户登陆异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userInfo == nil || userInfo.UserId < 1 {
		rr := gmhttp.NewFailedResultWithMsg(1, "用户名密码错误")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}

	respon := &loginRespon{
		UserId:      userInfo.UserId,
		UserName:    userInfo.UserName,
		Token:       userInfo.Token,
		ExpiredTime: userInfo.ExpiredTime,
		Access:      userInfo.Access,
		Avator:      userInfo.Avator,
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
