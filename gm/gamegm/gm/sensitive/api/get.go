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

type sensitiveRespon struct {
	Content string `form:"content" json:"content"`
}

func handleGetSensitive(rw http.ResponseWriter, req *http.Request) {
	service := gmsensitive.SensitiveServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("添加敏感词,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userid := gmuser.GmUserIdInContext(req.Context())

	rst, err := service.GetSensitive(userid)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}
	respon := &sensitiveRespon{
		Content: rst.Content,
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
