package api

import (
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	organizeservice "fgame/fgame/gm/gamegm/gm/organize/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type orderDatePlatformStaticRequest struct {
	StartTime    int64   `json:"startTime"`
	EndTime      int64   `json:"endTime"`
	ChannelList  []int64 `json:"channelId"`
	PlatformList []int64 `json:"platformId"`
}

type orderDatePlatformStaticRespon struct {
	ItemArray []*orderDatePlatformStaticResponItem `json:"itemArray"`
}

type orderDatePlatformStaticResponItem struct {
	SdkType        int `json:"sdkType"`
	OrderMoney     int `json:"orderMoney"`
	OrderPlayerNum int `json:"orderPlayerNum"`
	OrderNum       int `json:"orderNum"`
}

func handleOrderDatePlatformStaticList(rw http.ResponseWriter, req *http.Request) {
	form := &orderDatePlatformStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单按天列表解析，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterorder.OrderServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	gmuserid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("获取订单列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// userSdkTypeList, err := usservice.GetUserSdkTypeList(gmuserid)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"error":  err,
	// 		"userId": gmuserid,
	// 	}).Error("获取订单列表日期，获取权限中心平台列表异常")
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	orgs := organizeservice.OrganizeServiceInContext(req.Context())
	if orgs == nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，获取组织服务异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	var userSdkTypeList []int
	if len(form.PlatformList) == 0 {
		sdkList, err := orgs.GetSdkList(form.ChannelList)
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"userId": gmuserid,
			}).Error("获取订单列表日期，获取权限中心平台列表异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		userSdkTypeList = append(userSdkTypeList, sdkList...)
	} else {
		// sdkList, err := orgs.GetSdkListByPlatform(form.PlatformList)
		// if err != nil {
		// 	log.WithFields(log.Fields{
		// 		"error":  err,
		// 		"userId": gmuserid,
		// 	}).Error("获取订单列表日期，获取权限中心平台列表异常")
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		for _, platform := range form.PlatformList {
			userSdkTypeList = append(userSdkTypeList, int(platform))
		}
	}

	rst, err := service.GetCenterOrderStaticPlatform(form.StartTime, form.EndTime, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，查询订单数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &orderDatePlatformStaticRespon{}
	respon.ItemArray = make([]*orderDatePlatformStaticResponItem, 0)
	for _, value := range rst {
		item := &orderDatePlatformStaticResponItem{
			SdkType:        value.SdkType,
			OrderPlayerNum: value.OrderPlayerNum,
			OrderNum:       value.OrderNum,
			OrderMoney:     value.OrderMoney,
		}

		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
