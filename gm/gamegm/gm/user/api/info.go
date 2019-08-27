package api

import (
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleInfo(rw http.ResponseWriter, req *http.Request) {
	log.Debug("获得用户信息")
	service := gmUserService.LoginServiceInContext(req.Context())
	if service == nil {
		log.Error("用户登陆，获取登陆服务异常")
	}
	gmUserId := gmUserService.GmUserIdInContext(req.Context())

	userInfo, err := service.GetUserInfo(gmUserId)
	if err != nil {
		log.WithFields(log.Fields{
			"gmuserid": gmUserId,
			"error":    err,
		}).Error("用户登陆异常")
		rw.WriteHeader(http.StatusInternalServerError)
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
