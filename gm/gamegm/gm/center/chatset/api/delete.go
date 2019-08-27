package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmchatSet "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type chatSetDeleteRequest struct {
	ChatSetId int64 `form:"chatSetId" json:"chatSetId"`
}

func handleDeleteChatSet(rw http.ResponseWriter, req *http.Request) {
	form := &chatSetDeleteRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除聊天配置，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchatSet.ChatSetServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除聊天配置,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.DeleteChatSet(form.ChatSetId)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
