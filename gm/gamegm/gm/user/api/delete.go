package api

import (
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type deleteUserRequest struct {
	UserId int64 `form:"userid" json:"userId"`
}

func handleDeleteUser(rw http.ResponseWriter, req *http.Request) {
	form := &deleteUserRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("用户登陆，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmUserService.GetGmUserServiceInstance()
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取用户列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	myerr := service.DeleteUser(form.UserId)
	if err != nil {
		errhttp.ResponseWithError(rw, myerr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
