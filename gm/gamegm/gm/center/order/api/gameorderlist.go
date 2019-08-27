package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	monitor "fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/utils"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type gameOrderListRequest struct {
	PageIndex   int    `form:"pageIndex" json:"pageIndex"`
	OrderColumn int    `form:"ordercol" json:"ordercol"`
	OrderType   int    `form:"ordertype" json:"ordertype"`
	ServerId    int    `form:"serverId" json:"serverId"`
	StartTime   int64  `form:"startTime" json:"startTime"`
	EndTime     int64  `form:"endTime" json:"endTime"`
	MinAmount   int    `form:"minAmount" json:"minAmount"`
	MaxAmount   int    `form:"maxAmount" json:"maxAmount"`
	PlayerId    string `form:"playerId" json:"playerId"`
	UserId      string `form:"userId" json:"userId"`
	OrderId     string `form:"orderId" json:"orderId"`
	SdkOrderId  string `form:"sdkOrderId" json:"sdkOrderId"`
	PlayerName  string `form:"playerName" json:"playerName"`
	SdkType     int    `form:"sdkType" json:"sdkType"`
}

type gameOrderListRespon struct {
	ItemArray  []*gameOrderListResponItem `json:"itemArray"`
	TotalCount int                        `json:"total"`
}

type gameOrderListResponItem struct {
	Id          int64  `json:"id"`
	ServerId    int    `json:"serverId"`
	OrderId     string `json:"orderId"`
	OrderStatus int    `json:"orderStatus"`
	UserId      int64  `json:"userId"`
	PlayerId    string `json:"playerId"`
	ChargeId    int    `json:"chargeId"`
	Money       int    `json:"money"`
	UpdateTime  int64  `json:"updateTime"`
	CreateTime  int64  `json:"createTime"`
	DeleteTime  int64  `json:"deleteTime"`
	PlayerLevel int    `json:"playerLevel"`
	Gold        int    `json:"gold"`
	PlayerName  string `json:"playerName"`
}

func handleGameOrderList(rw http.ResponseWriter, req *http.Request) {
	form := &gameOrderListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单列表，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := gmcenterorder.OrderServiceInContext(req.Context())
	if service == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单列表，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("获取game订单列表，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp := &gameOrderListRespon{}
	rsp.ItemArray = make([]*gameOrderListResponItem, 0)
	playerId, _ := strconv.ParseInt(form.PlayerId, 10, 64)
	userid, _ := strconv.ParseInt(form.UserId, 10, 64)

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

	rst, err := service.GetGameOrderList(gmdb.GameDbLink(form.ServerId), acServerId, form.StartTime, form.EndTime, form.MinAmount, form.MaxAmount, playerId, userid, form.OrderId, form.SdkOrderId, form.PlayerName, form.PageIndex, form.OrderColumn, form.OrderType, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &gameOrderListResponItem{
			Id:          value.Id,
			OrderId:     value.OrderId,
			ServerId:    value.ServerId,
			OrderStatus: value.OrderStatus,
			UserId:      value.UserId,
			PlayerId:    utils.ConverInt64ToString(value.PlayerId),
			ChargeId:    value.ChargeId,
			Money:       value.Money,
			UpdateTime:  value.UpdateTime,
			CreateTime:  value.CreateTime,
			DeleteTime:  value.DeleteTime,
			PlayerLevel: value.PlayerLevel,
			Gold:        value.Gold,
			PlayerName:  value.PlayerName,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetGameOrderCount(gmdb.GameDbLink(form.ServerId), acServerId, form.StartTime, form.EndTime, form.MinAmount, form.MaxAmount, playerId, userid, form.OrderId, form.SdkOrderId, form.PlayerName, form.SdkType, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取game订单列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rsp.TotalCount = count

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
