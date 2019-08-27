package api

import (
	gmcenterorder "fgame/fgame/gm/gamegm/gm/center/order/service"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/utils"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type orderListRequest struct {
	OrderId    string `form:"orderId" json:"orderId"`
	SdkOrderId string `form:"sdkOrderId" json:"sdkOrderId"`
	PageIndex  int    `form:"pageIndex" json:"pageIndex"`
	SdkType    int    `form:"sdkType" json:"sdkType"`
	StartTime  int64  `form:"startTime" json:"startTime"`
	EndTime    int64  `form:"endTime" json:"endTime"`
}

type orderListRespon struct {
	ItemArray  []*orderListResponItem `json:"itemArray"`
	TotalCount int                    `json:"total"`
}

type orderListResponItem struct {
	Id             int64  `json:"id"`
	OrderId        string `json:"orderId"`
	SdkOrderId     string `json:"sdkOrderId"`
	Status         int    `json:"status"`
	SdkType        int    `json:"sdkType"`
	ServerId       int    `json:"serverId"`
	UserId         int64  `json:"userId"`
	PlayerId       string `json:"playerId"`
	ChargeId       int    `json:"chargeId"`
	Money          int    `json:"money"`
	ReceivePayTime int64  `json:"receivePayTime"`
	UpdateTime     int64  `json:"updateTime"`
	CreateTime     int64  `json:"createTime"`
	DeleteTime     int64  `json:"deleteTime"`
	PlayerName     string `json:"playerName"`
	Gold           int    `json:"gold"`
	PlayerLevel    int    `json:"playerLevel"`
}

func handleOrderList(rw http.ResponseWriter, req *http.Request) {
	form := &orderListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单列表，解析异常")
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
		}).Error("获取game订单列表，获取权限中心平台列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp := &orderListRespon{}
	rsp.ItemArray = make([]*orderListResponItem, 0)

	rst, err := service.GetOrderList(form.OrderId, form.SdkOrderId, form.SdkType, form.StartTime, form.EndTime, form.PageIndex, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, value := range rst {
		item := &orderListResponItem{
			Id:             value.Id,
			OrderId:        value.OrderId,
			SdkOrderId:     value.SdkOrderId,
			Status:         value.Status,
			SdkType:        value.SdkType,
			ServerId:       value.ServerId,
			UserId:         value.UserId,
			PlayerId:       utils.ConverInt64ToString(value.PlayerId),
			ChargeId:       value.ChargeId,
			Money:          value.Money,
			ReceivePayTime: value.ReceivePayTime,
			UpdateTime:     value.UpdateTime,
			CreateTime:     value.CreateTime,
			DeleteTime:     value.DeleteTime,
			PlayerName:     value.PlayerName,
			Gold:           value.Gold,
			PlayerLevel:    value.PlayerLevel,
		}
		rsp.ItemArray = append(rsp.ItemArray, item)
	}

	count, err := service.GetOrderCount(form.OrderId, form.SdkOrderId, form.SdkType, form.StartTime, form.EndTime, userSdkTypeList)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("获取订单列表异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rsp.TotalCount = count

	rr := gmhttp.NewSuccessResult(rsp)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
