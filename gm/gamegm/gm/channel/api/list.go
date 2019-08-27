package api

import (
	"net/http"

	gmchannel "fgame/fgame/gm/gamegm/gm/channel/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type channelListRequest struct {
	PageIndex   int    `form:"pageIndex" json:"pageIndex"`
	ChannelName string `form:"channelName" json:"channelName"`
}

type channelListRespon struct {
	ItemArray  []*channelListResponItem `json:"itemArray"`
	TotalCount int                      `json:"total"`
}

type channelListResponItem struct {
	ChannelId   int64  `json:"channelId"`
	ChannelName string `json:"channelName"`
}

func handleChannelList(rw http.ResponseWriter, req *http.Request) {
	form := &channelListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取渠道列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmchannel.ChannelServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取渠道列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetChannelList(form.ChannelName, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"error":       err,
			"ChannelName": form.ChannelName,
			"index":       form.PageIndex,
		}).Error("获取渠道列表，执行异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &channelListRespon{}
	respon.ItemArray = make([]*channelListResponItem, 0)

	for _, value := range rst {
		item := &channelListResponItem{
			ChannelId:   value.ChannelID,
			ChannelName: value.ChannelName,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	count, err := service.GetChannelCount(form.ChannelName)
	if err != nil {
		log.WithFields(log.Fields{
			"error":       err,
			"ChannelName": form.ChannelName,
			"index":       form.PageIndex,
		}).Error("获取渠道列表，执行异常")
	}
	respon.TotalCount = count
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
