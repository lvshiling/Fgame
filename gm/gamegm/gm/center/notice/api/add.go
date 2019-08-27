package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmLoginNotice "fgame/fgame/gm/gamegm/gm/center/notice/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type addLoginNoticeRequest struct {
	PlatformId int64  `form:"platformId" json:"platformId"`
	Content    string `form:"content" json:"content"`
}

func handleAddLoginNotice(rw http.ResponseWriter, req *http.Request) {
	form := &addLoginNoticeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加中心登陆公告，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmLoginNotice.LoginNoticeServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加中心登陆公告,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = service.AddLoginNotice(form.PlatformId, form.Content)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
