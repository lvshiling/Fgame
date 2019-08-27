package api

import (
	"fgame/fgame/gm/gamegm/gm/center/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type updateGmRequest struct {
	UserId   int    `json:"id"`
	GmFlag   int    `json:"gm"`
	UserName string `json:"name"`
	PassWord string `json:"password"`
}

func handleUpdateGmList(rw http.ResponseWriter, req *http.Request) {
	form := &updateGmRequest{}
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
	if len(form.UserName) != 0 {
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
	}
	err = rds.UpdateGm(int64(form.UserId), form.GmFlag, form.UserName, form.PassWord)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userid": form.UserId,
			"gmflag": form.GmFlag,
		}).Error("更新gm失败")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
