package api

import (
	feeservice "fgame/fgame/gm/gamegm/gm/feedbackfee/service"
	"fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/utils"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type centerFeedBackExchangeListRequest struct {
	PageIndex  int32  `form:"pageIndex" json:"pageIndex"`
	PlatformId int32  `form:"platformId" json:"platformId"`
	ServerId   int    `form:"serverId" json:"serverId"`
	StartTime  int64  `form:"startTime" json:"startTime"`
	EndTime    int64  `form:"endTime" json:"endTime"`
	PlayerId   string `form:"playerId" json:"playerId"`
	Code       string `json:"code"`
}

type centerFeedBackExchangeListRespon struct {
	ItemArray  []*centerFeedBackExchangeListResponItem `json:"itemArray"`
	TotalCount int32                                   `json:"total"`
}

type centerFeedBackExchangeListResponItem struct {
	Id          int64  `json:"id"`
	Platform    int32  `json:"platform"`
	ServerId    int32  `json:"serverId"`
	PlayerId    string `json:"playerId"`
	ExchangeId  int64  `json:"exchangeId"`
	ExpiredTime int64  `json:"expiredTime"`
	Money       int32  `json:"money"`
	Code        string `json:"code"`
	Status      int32  `json:"status"`
	WxId        string `json:"wxId"`
	OrderId     string `json:"orderId"`
	UpdateTime  int64  `json:"updateTime"`
	CreateTime  int64  `json:"createTime"`
	DeleteTime  int64  `json:"deleteTime"`
}

func handleCenterFeeBackExchangeList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleCenterFeeBackExchangeList:获取游戏服返还订单")
	form := &centerFeedBackExchangeListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleCenterFeeBackExchangeList，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	playerId, _ := strconv.ParseInt(form.PlayerId, 10, 64)

	acServerId := int32(0)
	if form.ServerId > 0 {
		centerService := monitor.CenterServerServiceInContext(req.Context())

		acServerId, err := centerService.GetServerId(int64(form.ServerId))
		log.Debug("中心序号id:", acServerId)
		if err != nil {
			log.WithFields(log.Fields{
				"dbid":  form.ServerId,
				"error": err,
			}).Error("handleCenterFeeBackExchangeList，获取服务id异常")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	feedBackService := feeservice.FeedBackFeeServiceInContext(req.Context())
	list, err := feedBackService.GetCenterFeedBackFeeList(form.PlatformId, acServerId, playerId, form.Code, form.StartTime, form.EndTime, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("handleCenterFeeBackExchangeList，获取feedbacklist异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	count, err := feedBackService.GetCenterFeedBackFeeCount(form.PlatformId, acServerId, playerId, form.Code, form.StartTime, form.EndTime)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("handleCenterFeeBackExchangeList，获取feedbackcount异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &centerFeedBackExchangeListRespon{}
	respon.TotalCount = count
	for _, value := range list {
		item := &centerFeedBackExchangeListResponItem{
			Id:          value.Id,
			Platform:    value.Platform,
			ServerId:    value.ServerId,
			PlayerId:    utils.ConverInt64ToString(value.PlayerId),
			ExchangeId:  value.ExchangeId,
			ExpiredTime: value.ExpiredTime,
			Money:       value.Money,
			Code:        value.Code,
			Status:      value.Status,
			WxId:        value.WxId,
			OrderId:     value.OrderId,
			UpdateTime:  value.UpdateTime,
			CreateTime:  value.CreateTime,
			DeleteTime:  value.DeleteTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
