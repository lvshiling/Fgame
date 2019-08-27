package api

import (
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	"fgame/fgame/gm/gamegm/gm/types"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type orderStaticTotalRequest struct {
	SdkType int `form:"sdkType" json:"sdkType"`
}

type orderStaticTotalRespon struct {
	SdkType        int `json:"sdkType"`
	TodayAmount    int `json:"todayAmount"`
	TodayPerson    int `json:"todayPerson"`
	YestodayAmount int `json:"yestodayAmount"`
	YestodayPerson int `json:"yestodayPerson"`
	TotalAmount    int `json:"totalAmount"`
	TotalPerson    int `json:"totalPerson"`
}

func handleorderStaticTotal(rw http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{}).Debug("请求获取中心订单统计")

	service := gmcenterorder.OrderServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{}).Error("获取订单统计汇总，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp := &orderStaticTotalRespon{}
	userPrivilege := types.PrivilegeLevel(gmUserService.PrivilegeInContext(req.Context()))
	if !userPrivilege.HasCanShouYeOrder() {
		rr := gmhttp.NewSuccessResult(rsp)
		httputils.WriteJSON(rw, http.StatusOK, rr)
	}

	gmuserid := gmUserService.GmUserIdInContext(req.Context())

	usservice := gmUserService.GmUserServiceInContext(req.Context())
	if usservice == nil {
		log.WithFields(log.Fields{}).Error("邮件列表，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userSdkTypeList, err := usservice.GetUserSdkTypeList(gmuserid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取game订单列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalRst, err := service.GetCenterOrderTotalStatic(0, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单统计汇总异常")
	}
	if len(totalRst) > 0 {
		value := totalRst[0]
		rsp = &orderStaticTotalRespon{
			SdkType:        -1,
			TodayAmount:    value.TodayAmount,
			TodayPerson:    value.TodayPerson,
			YestodayAmount: value.YestodayAmount,
			YestodayPerson: value.YestodayPerson,
			TotalAmount:    value.TotalAmount,
			TotalPerson:    value.TotalPerson,
		}
	}

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
