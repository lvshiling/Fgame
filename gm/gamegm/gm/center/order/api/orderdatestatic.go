package api

import (
	pubmodel "fgame/fgame/gm/gamegm/gm/center/order/pubmodel"
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	organizeservice "fgame/fgame/gm/gamegm/gm/organize/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type orderDateStaticRequest struct {
	StartTime    int64   `json:"startTime"`
	EndTime      int64   `json:"endTime"`
	ChannelList  []int64 `json:"channelId"`
	PlatformList []int   `json:"platformId"`
	ServerList   []int   `json:"serverId"`
}

type orderDateStaticRespon struct {
	ItemArray []*orderDateStaticResponItem `json:"itemArray"`
}

type orderDateStaticResponItem struct {
	OrderDate             string `json:"orderDate"`
	OrderPlayerNum        int    `json:"orderPlayerNum"`
	OrderNum              int    `json:"orderNum"`
	OrderMoney            int    `json:"orderMoney"`
	OrderGold             int    `json:"orderGold"`
	TotalPlayerCount      int    `json:"totalPlayerCount"`
	ServerCount           int    `json:"serverCount"`
	DateNewPlayerCount    int    `json:"dateNewPlayerCount"`
	OrderNewPlayerCount   int    `json:"orderNewPlayerCount"`
	OrderNewMoney         int    `json:"orderNewMoney"`
	OrderFirstPlayerCount int    `json:"orderFirstPlayerCount"`
	OrderFirstMoneyCount  int    `json:"orderFirstMoneyCount"`
}

func handleOrderDateStaticList(rw http.ResponseWriter, req *http.Request) {
	form := &orderDateStaticRequest{}
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

	userSdkTypeList, err := usservice.GetUserSdkTypeList(gmuserid)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	orgs := organizeservice.OrganizeServiceInContext(req.Context())
	if orgs == nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，获取组织服务异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serverMap, err := orgs.GetSdkServer(form.ChannelList, form.PlatformList, form.ServerList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，组织查询列表条件异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rst, err := service.GetCenterOrderStaticMultiple(serverMap, form.StartTime, form.EndTime, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，查询订单数据异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	userCountList, err := service.GetPlayerCountDate(serverMap, form.StartTime, form.EndTime, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，查询用户列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug("获取玩家数据", userCountList)
	serverCount, err := orgs.GetSdkServerCount(form.ChannelList, form.PlatformList, form.ServerList)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"userId": gmuserid,
		}).Error("获取订单列表日期，获得服务器个数异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &orderDateStaticRespon{}
	respon.ItemArray = make([]*orderDateStaticResponItem, 0)
	for _, value := range rst {
		item := &orderDateStaticResponItem{
			OrderDate:             value.OrderDate,
			OrderPlayerNum:        value.OrderPlayerNum,
			OrderGold:             value.OrderGold,
			OrderMoney:            value.OrderMoney,
			OrderNum:              value.OrderNum,
			ServerCount:           serverCount,
			OrderNewPlayerCount:   value.OrderNewPlayer,
			OrderNewMoney:         value.OrderNewMoney,
			OrderFirstPlayerCount: value.OrderFirstPlayer,
			OrderFirstMoneyCount:  value.OrderFirstMoney,
		}
		userCount := findTimeUserCount(value.OrderDate, userCountList)
		item.TotalPlayerCount = userCount.LeiJiCount
		item.DateNewPlayerCount = userCount.DateCount
		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}

func findTimeUserCount(p_date string, userMap map[int64]*pubmodel.UserDateCount) *pubmodel.UserDateCount {
	dateStr := strings.Replace(p_date, "-", "", -1)
	time, err := timeutils.ParseYYYYMMDD(dateStr)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"timestr": dateStr,
		}).Error("转换异常")
		return nil
	}
	log.Debug("查找人数:", time)
	if value, ok := userMap[time]; ok {
		log.Debug("查找到了", time, "人数：", value)
		return value
	}
	log.Debug("没查找到", time)
	return nil
}
