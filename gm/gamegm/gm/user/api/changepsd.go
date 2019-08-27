package api

import (
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type deleteChangePassWordRequest struct {
	UserId   int64  `form:"userid" json:"userId"`
	Password string `form:"password" json:"password"`
}

func handleChangePassWord(rw http.ResponseWriter, req *http.Request) {
	form := &deleteChangePassWordRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改密码，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmUserService.GetGmUserServiceInstance()
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("修改密码，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	myerr := service.ChangePassWord(form.UserId, form.Password)
	if myerr != nil {
		errhttp.ResponseWithError(rw, myerr)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
