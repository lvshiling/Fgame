package api

import (
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type orderStaticRequest struct {
	SdkType int `form:"sdkType" json:"sdkType"`
}

type orderStaticRespon struct {
	ItemArray  []*orderStaticResponItem `json:"itemArray"`
	TotalCount int                      `json:"total"`
}

type orderStaticResponItem struct {
	SdkType                int `json:"sdkType"`
	TodayAmount            int `json:"todayAmount"`
	TodayPerson            int `json:"todayPerson"`
	TodayRegisterPerson    int `json:"todayRegisterPerson"`
	YestodayAmount         int `json:"yestodayAmount"`
	YestodayPerson         int `json:"yestodayPerson"`
	YestodayRegisterPerson int `json:"yestodayRegisterPerson"`
	TotalAmount            int `json:"totalAmount"`
	TotalPerson            int `json:"totalPerson"`
	TotalRegisterPerson    int `json:"totalRegisterPerson"`
	MonthAmount            int `json:"monthAmount"`
	MonthPerson            int `json:"monthPerson"`
	MonthActivityPerson    int `json:"monthActivityPerson"`
	TodayAvtivityPerson    int `json:"todayActivityPerson"`
	ThreeDayActivityPerson int `json:"threeDayActivityPerson"`
	WeekActivityPerson     int `json:"weekActivityPerson"`
	YestodayAvtivityPerson int `json:"yestodayAvtivityPerson"`
}

func handleOrderStatic(rw http.ResponseWriter, req *http.Request) {
	form := &orderStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单统计，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.WithFields(log.Fields{
		"sdkType": form.SdkType,
	}).Debug("请求获取中心订单统计")

	service := gmcenterorder.OrderServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单统计，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
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

	rpservice := stservice.StaticReportServiceInContext(req.Context())
	if rpservice == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，获取mongo统计服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	lastNowMonth := timeutils.TimeToMillisecond(time.Now().AddDate(0, -1, 0))
	log.Debug("lastNowMonth:", lastNowMonth, "Now:", now, "sdkList", userSdkTypeList, "formsdk", form.SdkType)
	sdkCountList, err := rpservice.GetOnLinePlayerNumGroupSdk(lastNowMonth, now, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计月在线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	beginYestoday, _ := timeutils.BeginOfYesterday()
	beginNow, _ := timeutils.BeginOfNow(now)
	sdkCountYestoday, err := rpservice.GetOnLinePlayerNumGroupSdk(beginYestoday, beginNow, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计昨日线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	sdkCountToday, err := rpservice.GetOnLinePlayerNumGroupSdk(beginNow, now, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计今日线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	threeDayAgo := now - 3*24*int64(time.Hour/time.Millisecond)
	weekDayAgo := now - 7*24*int64(time.Hour/time.Millisecond)
	sdkCountThreeDay, err := rpservice.GetOnLinePlayerNumGroupSdk(threeDayAgo, now, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计三日线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	sdkCountWeekDay, err := rpservice.GetOnLinePlayerNumGroupSdk(weekDayAgo, now, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计一周线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalYestodayPlayerCount := 0
	totalTodayPlayerCount := 0
	totalPlayerCount := 0
	for _, value := range sdkCountList {
		totalPlayerCount += value.Count
	}
	for _, value := range sdkCountYestoday {
		totalYestodayPlayerCount += value.Count
	}
	for _, value := range sdkCountToday {
		totalTodayPlayerCount += value.Count
	}

	rsp := &orderStaticRespon{}
	rsp.ItemArray = make([]*orderStaticResponItem, 0)

	rst, err := service.GetCenterOrderStatic(form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单统计异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &orderStaticResponItem{
			SdkType:                value.SdkType,
			TodayAmount:            value.TodayAmount,
			TodayPerson:            value.TodayPerson,
			TodayRegisterPerson:    value.TodayRegisterPerson,
			YestodayAmount:         value.YestodayAmount,
			YestodayPerson:         value.YestodayPerson,
			YestodayRegisterPerson: value.YestodayRegisterPerson,
			TotalAmount:            value.TotalAmount,
			TotalPerson:            value.TotalPerson,
			TotalRegisterPerson:    value.TotalRegisterPerson,
			MonthAmount:            value.MonthAmount,
			MonthPerson:            value.MonthPerson,
		}
		for _, sdkValue := range sdkCountList {
			if sdkValue.Id.SdkType == value.SdkType {
				item.MonthActivityPerson = sdkValue.Count
				break
			}
		}
		for _, countValue := range sdkCountToday {
			if countValue.Id.SdkType == value.SdkType {
				item.TodayAvtivityPerson = countValue.Count
				break
			}
		}
		for _, countValue := range sdkCountThreeDay {
			if countValue.Id.SdkType == value.SdkType {
				item.ThreeDayActivityPerson = countValue.Count
				break
			}
		}
		for _, countValue := range sdkCountWeekDay {
			if countValue.Id.SdkType == value.SdkType {
				item.WeekActivityPerson = countValue.Count
				break
			}
		}

		for _, countValue := range sdkCountYestoday {
			if countValue.Id.SdkType == value.SdkType {
				item.YestodayAvtivityPerson = countValue.Count
				break
			}
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	totalRst, err := service.GetCenterOrderTotalStatic(form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单统计异常")
	}
	if len(totalRst) > 0 {
		value := totalRst[0]
		item := &orderStaticResponItem{
			SdkType:                -1,
			TodayAmount:            value.TodayAmount,
			TodayPerson:            value.TodayPerson,
			TodayRegisterPerson:    value.TodayRegisterPerson,
			YestodayAmount:         value.YestodayAmount,
			YestodayPerson:         value.YestodayPerson,
			YestodayRegisterPerson: value.YestodayRegisterPerson,
			TotalAmount:            value.TotalAmount,
			TotalPerson:            value.TotalPerson,
			TotalRegisterPerson:    value.TotalRegisterPerson,
			MonthAmount:            value.MonthAmount,
			MonthPerson:            value.MonthPerson,
			MonthActivityPerson:    totalPlayerCount,
			TodayAvtivityPerson:    totalTodayPlayerCount,
			YestodayAvtivityPerson: totalYestodayPlayerCount,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
