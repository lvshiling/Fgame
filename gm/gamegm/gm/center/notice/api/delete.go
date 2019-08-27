package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmloginNotice "fgame/fgame/gm/gamegm/gm/center/notice/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type loginNoticeDeleteRequest struct {
	ID int64 `form:"id" json:"id"`
}

func handleDeleteLoginNotice(rw http.ResponseWriter, req *http.Request) {
	form := &loginNoticeDeleteRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除中心平台公告，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmloginNotice.LoginNoticeServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除中心平台公告,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.DeleteLoginNotice(form.ID)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
