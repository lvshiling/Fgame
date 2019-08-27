package api

import (
	"net/http"

	errhttp "fgame/fgame/gm/gamegm/error/utils"
	channelmodel "fgame/fgame/gm/gamegm/gm/channel/model"
	gmchannel "fgame/fgame/gm/gamegm/gm/channel/service"
	us "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleAllChannelList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("获取渠道列表")
	userId := us.GmUserIdInContext(req.Context())

	usservice := us.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取渠道列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmchannel.ChannelServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("获取渠道列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userInfo, err := usservice.GetUserInfo(userId) //得缓存
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取渠道列表，获取用户信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst := make([]*channelmodel.ChannelInfo, 0)

	if userInfo.ChannelID > 0 {
		channelInfo, err := service.GetChannelInfo(userInfo.ChannelID)
		if err != nil {
			log.WithFields(log.Fields{
				"error":     err,
				"channelid": userInfo.ChannelID,
			}).Error("获取渠道列表，获取异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rst = append(rst, channelInfo)
	} else {
		rst, err = service.GetAllChannelList()
		if err != nil {
			errhttp.ResponseWithError(rw, err)
			return
		}
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

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
