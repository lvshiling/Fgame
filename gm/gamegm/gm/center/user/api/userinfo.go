package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type userInfoRequest struct {
	UserId int64 `json:"id"`
}

type userInfoRespon struct {
	UserId int64  `json:"id"`
	Name   string `json:"name"`
}

func handleUserInfo(rw http.ResponseWriter, req *http.Request) {
	form := &userInfoRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心用户信息，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rds := service.CenterUserServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取中心用户信息，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userInfo, err := rds.GetUserInfo(form.UserId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": form.UserId,
		}).Error("获取中心用户信息，获取失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := userInfoRespon{
		UserId: int64(userInfo.Id),
		Name:   userInfo.Name,
	}
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
