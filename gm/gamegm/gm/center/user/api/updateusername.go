package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateUserNameRequest struct {
	UserId   int    `json:"id"`
	UserName string `json:"name"`
	PassWord string `json:"password"`
}

func handleUpdateUserNameList(rw http.ResponseWriter, req *http.Request) {
	form := &updateUserNameRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新gm，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rds := service.CenterUserServiceInContext(req.Context())
	if rds == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("更新gm，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	exflag, err := rds.ExistsUserName(int64(form.UserId), form.UserName)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"userid":   form.UserId,
			"UserName": form.UserName,
		}).Error("中心用户名判断存在失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exflag {
		rr := gmhttp.NewFailedResultWithMsg(1000, "用户名已经存在")
		httputils.WriteJSON(rw, http.StatusOK, rr)
		return
	}
	err = rds.UpdateUserInfo(int64(form.UserId), form.UserName, form.PassWord)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userid": form.UserId,
		}).Error("更新用户名失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
