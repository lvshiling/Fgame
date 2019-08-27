package api

import (
	"net/http"

	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmchatSet "fgame/fgame/gm/gamegm/gm/center/chatset/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	remoteservice "fgame/fgame/gm/gamegm/remote/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type chatSetDeletePlatformRequest struct {
	ChatSetId int64 `form:"chatSetId" json:"chatSetId"`
}

func handleDeleteChatSetPlatform(rw http.ResponseWriter, req *http.Request) {
	form := &chatSetDeletePlatformRequest{}
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

	chatInfo, err := service.GetChatSetPlatformById(int(form.ChatSetId))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除聊天配置,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.DeleteChatSetPlatform(form.ChatSetId)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	remoteService := remoteservice.CenterServiceInContext(req.Context())
	if remoteService == nil {
		log.Error("添加聊天配置，remote服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = remoteService.RefreshPlatformConfig(int32(chatInfo.PlatformId))
	if err != nil {
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
