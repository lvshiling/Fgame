package api

import (
	gmdb "fgame/fgame/gm/gamegm/db"
	feeservice "fgame/fgame/gm/gamegm/gm/feedbackfee/service"
	"fgame/fgame/gm/gamegm/monitor"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"fgame/fgame/gm/gamegm/utils"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type gameFeedBackExchangeListRequest struct {
	PageIndex int32  `form:"pageIndex" json:"pageIndex"`
	ServerId  int    `form:"serverId" json:"serverId"`
	StartTime int64  `form:"startTime" json:"startTime"`
	EndTime   int64  `form:"endTime" json:"endTime"`
	PlayerId  string `form:"playerId" json:"playerId"`
	Code      string `json:"code"`
}

type gameFeedBackExchangeListRespon struct {
	ItemArray  []*gameFeedBackExchangeListResponItem `json:"itemArray"`
	TotalCount int32                                 `json:"total"`
}

type gameFeedBackExchangeListResponItem struct {
	Id          int64  `json:"id"`
	ServerId    int32  `json:"serverId"`
	PlayerId    string `json:"playerId"`
	ExchangeId  int64  `json:"exchangeId"`
	ExpiredTime int64  `json:"expiredTime"`
	Money       int32  `json:"money"`
	Code        string `json:"code"`
	Status      int32  `json:"status"`
	UpdateTime  int64  `json:"updateTime"`
	CreateTime  int64  `json:"createTime"`
	DeleteTime  int64  `json:"deleteTime"`
}

func handleGameFeeBackExchangeList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("handleGameFeeBackExchangeList:获取游戏服返还订单")
	form := &gameFeedBackExchangeListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("handleGameFeeBackExchangeList，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	playerId, _ := strconv.ParseInt(form.PlayerId, 10, 64)

	centerService := monitor.CenterServerServiceInContext(req.Context())

	acServerId, err := centerService.GetServerId(int64(form.ServerId))
	log.Debug("中心序号id:", acServerId)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("handleGameFeeBackExchangeList，获取服务id异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	feedBackService := feeservice.FeedBackFeeServiceInContext(req.Context())
	list, err := feedBackService.GetGameFeedBackFeeList(gmdb.GameDbLink(form.ServerId), acServerId, playerId, form.Code, form.StartTime, form.EndTime, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("handleGameFeeBackExchangeList，获取feedbacklist异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	count, err := feedBackService.GetGameFeedBackFeeCount(gmdb.GameDbLink(form.ServerId), acServerId, playerId, form.Code, form.StartTime, form.EndTime)
	if err != nil {
		log.WithFields(log.Fields{
			"dbid":  form.ServerId,
			"error": err,
		}).Error("handleGameFeeBackExchangeList，获取feedbackcount异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respon := &gameFeedBackExchangeListRespon{}
	respon.TotalCount = count
	for _, value := range list {
		item := &gameFeedBackExchangeListResponItem{
			Id:          value.Id,
			ServerId:    value.ServerId,
			PlayerId:    utils.ConverInt64ToString(value.PlayerId),
			ExchangeId:  value.ExchangeId,
			ExpiredTime: value.ExpiredTime,
			Money:       value.Money,
			Code:        value.Code,
			Status:      value.Status,
			UpdateTime:  value.UpdateTime,
			CreateTime:  value.CreateTime,
			DeleteTime:  value.DeleteTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
