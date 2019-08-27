package api

import (
	"net/http"

	gmplatform "fgame/fgame/gm/gamegm/gm/platform/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"

	platformmodel "fgame/fgame/gm/gamegm/gm/platform/model"
	us "fgame/fgame/gm/gamegm/gm/user/service"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

func handleAllPlatformList(rw http.ResponseWriter, req *http.Request) {

	userId := us.GmUserIdInContext(req.Context())

	usservice := us.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取平台列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := gmplatform.PlatformServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("获取平台列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userInfo, err := usservice.GetUserInfo(userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取平台列表，获取用户信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst := make([]*platformmodel.PlatformInfo, 0)

	if userInfo.PlatformId > 0 {
		platInfo, err := service.GetPlatformInfo(userInfo.PlatformId)
		if err != nil {
			log.WithFields(log.Fields{
				"error":      err,
				"PlatformId": userInfo.PlatformId,
			}).Error("获取平台列表，执行异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rst = append(rst, platInfo)
	} else {
		if userInfo.ChannelID > 0 {
			rst, err = service.GetPlatformByChannel(userInfo.ChannelID)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("获取平台列表，执行异常")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			rst, err = service.GetAllPlatformList()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("获取平台列表，执行异常")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	respon := &platformListRespon{}
	respon.ItemArray = make([]*platformListResponItem, 0)

	for _, value := range rst {
		item := &platformListResponItem{
			PlatformId:       value.PlatformID,
			PlatformName:     value.PlatformName,
			ChannelId:        value.ChannelId,
			CenterPlatformId: value.CenterPlatformID,
			SdkType:          value.SdkType,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}
	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
