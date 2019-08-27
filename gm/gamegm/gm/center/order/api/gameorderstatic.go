package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type gameOrderStaticRequest struct {
	ServerId int `form:"serverId" json:"serverId"`
	SdkType  int `form:"sdkType" json:"sdkType"`
}

type gameOrderStaticRespon struct {
	ItemArray  []*gameOrderStaticResponItem `json:"itemArray"`
	TotalCount int                          `json:"total"`
}

type gameOrderStaticResponItem struct {
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

func handleGameOrderStatic(rw http.ResponseWriter, req *http.Request) {
	form := &gameOrderStaticRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterorder.OrderServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	log.Debug("中心序号id:", acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取game订单统计，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	serverInfo, err := centerService.GetCenterServerDbInfo(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取game订单统计，获取服务信息异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rpservice := stservice.StaticReportServiceInContext(req.Context())
	if rpservice == nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取game订单统计，获取mongo统计服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := timeutils.TimeToMillisecond(time.Now())
	lastNowMonth := timeutils.TimeToMillisecond(time.Now().AddDate(0, -1, 0))
	count, err := rpservice.GetOnLinePlayerNum(lastNowMonth, now, serverInfo.Platform, acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计月在线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	beginYestoday, _ := timeutils.BeginOfYesterday()
	beginNow, _ := timeutils.BeginOfNow(now)
	countYestoday, err := rpservice.GetOnLinePlayerNum(beginYestoday, beginNow, serverInfo.Platform, acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计昨日线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	countToday, err := rpservice.GetOnLinePlayerNum(beginNow, now, serverInfo.Platform, acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计今日线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	threeDayAgo := now - 3*24*int64(time.Hour/time.Millisecond)
	weekDayAgo := now - 7*24*int64(time.Hour/time.Millisecond)
	threeActivityCount, err := rpservice.GetOnLinePlayerNum(threeDayAgo, now, serverInfo.Platform, acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计三日线人数出错")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	weekActivityCount, err := rpservice.GetOnLinePlayerNum(weekDayAgo, now, serverInfo.Platform, acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计，统计周线人数出错")
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

	rsp := &gameOrderStaticRespon{}
	rsp.ItemArray = make([]*gameOrderStaticResponItem, 0)

	rst, err := service.GetGameOrderStatic(gmdb.GameDbLink(form.ServerId), acServerId, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单统计异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &gameOrderStaticResponItem{
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
			MonthActivityPerson:    count,
			TodayAvtivityPerson:    countToday,
			YestodayAvtivityPerson: countYestoday,
			ThreeDayActivityPerson: threeActivityCount,
			WeekActivityPerson:     weekActivityCount,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	// count, err := service.GetGameOrderCount(gmdb.GameDbLink(form.ServerId), acServerId, form.StartTime, form.EndTime, form.MinAmount, form.MaxAmount, playerId, userid, form.OrderId, form.SdkOrderId, form.PlayerName)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"error": err,
	// 	}).Error("获取game订单统计异常")
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// rsp.TotalCount = count

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
