package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmchannel "fgame/fgame/gm/gamegm/gm/channel/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type channelRequest struct {
	ChannelName string `form:"channelName" json:"channelName"`
}

func handleAddChannel(rw http.ResponseWriter, req *http.Request) {
	form := &channelRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加渠道，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchannel.ChannelServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("添加渠道,服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.AddChannel(form.ChannelName)
	if err != nil {
		errhttp.ResponseWithError(rw, err)
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
