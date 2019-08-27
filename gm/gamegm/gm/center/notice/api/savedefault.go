package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmloginnotice "fgame/fgame/gm/gamegm/gm/center/notice/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type saveDefaultNoticeRequest struct {
	Content string `form:"content" json:"content"`
}

func handleUpdateDefaultNotice(rw http.ResponseWriter, req *http.Request) {
	form := &saveDefaultNoticeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("默认公告添加，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmloginnotice.LoginNoticeServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("默认公告添加,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.UpdateDefaultNotice(form.Content)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
