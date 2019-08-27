package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmsensitive "fgame/fgame/gm/gamegm/gm/sensitive/service"
	gmuser "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type sensitiveRequest struct {
	Content string `form:"content" json:"content"`
}

func handleAddSensitive(rw http.ResponseWriter, req *http.Request) {
	form := &sensitiveRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加敏感词，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmsensitive.SensitiveServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加敏感词,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid := gmuser.GmUserIdInContext(req.Context())

	err = service.Save(userid, form.Content)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
