package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmchannel "fgame/fgame/gm/gamegm/gm/channel/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type channelDeleteRequest struct {
	ChannelId int64 `form:"channelId" json:"channelId"`
}

func handleDeleteChannel(rw http.ResponseWriter, req *http.Request) {
	form := &channelDeleteRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除渠道，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchannel.ChannelServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("删除渠道,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.DeleteChannel(form.ChannelId)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
