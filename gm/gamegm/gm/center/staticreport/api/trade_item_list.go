package api

import (
	stservice "fgame/fgame/gm/gamegm/gm/center/staticreport/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	"fgame/fgame/gm/gamegm/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type tradeItemListRequest struct {
	PlatformId int32  `json:"platformId"`
	ServerId   int32  `json:"serverId"`
	StartTime  int64  `json:"startTime"`
	EndTime    int64  `json:"endTime"`
	TradeId    string `json:"tradeId"`
	PlayerId   string `json:"playerId"`
	Level      int32  `json:"level"`
	State      int32  `json:"state"`
	PageIndex  int32  `json:"pageIndex"`
}

type tradeItemListRespon struct {
	ItemArray  []*tradeItemListResponItem `json:"itemArray"`
	TotalCount int32                      `json:"total"`
}

type tradeItemListResponItem struct {
	Id                int64  `json:"id"`
	Platform          int32  `json:"platform"`
	ServerId          int32  `json:"serverId"`
	TradeId           int64  `json:"tradeId"`
	PlayerId          int64  `json:"playerId"`
	PlayerName        string `json:"playerName"`
	ItemId            int32  `json:"itemId"`
	ItemNum           int32  `json:"itemNum"`
	Level             int32  `json:"level"`
	Gold              int64  `json:"gold"`
	PropertyData      string `json:"propertyData"`
	Status            int32  `json:"status"`
	BuyPlayerPlatform int32  `json:"buyPlayerPlatform"`
	BuyPlayerServerId int32  `json:"buyPlayerServerId"`
	BuyPlayerId       int64  `json:"buyPlayerId"`
	BuyPlayerName     string `json:"buyPlayerName"`
	UpdateTime        int64  `json:"updateTime"`
	CreateTime        int64  `json:"createTime"`
	DeleteTime        int64  `json:"deleteTime"`
}

//交易行
func handleTradeItemList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("交易行查询")
	form := &tradeItemListRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("交易行查询，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := stservice.PlayerStaticInContext(req.Context())
	tradeId := utils.ConverStringToInt64(form.TradeId)
	playerId := utils.ConverStringToInt64(form.PlayerId)
	centerPlatformId := form.PlatformId
	centerServerId := form.ServerId
	rst, err := service.GetTradeItemList(centerPlatformId, centerServerId, form.StartTime, form.EndTime, tradeId, playerId, form.Level, form.State, form.PageIndex)
	if err != nil {
		log.WithFields(log.Fields{
			"plarformId": form.PlatformId,
			"serverId":   form.ServerId,
			"StartTime":  form.StartTime,
			"EndTime":    form.EndTime,
			"tradeId":    form.TradeId,
			"playerId":   form.PlayerId,
			"Level":      form.Level,
			"State":      form.State,
			"PageIndex":  form.PageIndex,
			"error":      err,
		}).Error("交易行查询，查询异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	count, err := service.GetTradeItemCount(centerPlatformId, centerServerId, form.StartTime, form.EndTime, tradeId, playerId, form.Level, form.State)
	if err != nil {
		log.WithFields(log.Fields{
			"plarformId": form.PlatformId,
			"serverId":   form.ServerId,
			"StartTime":  form.StartTime,
			"EndTime":    form.EndTime,
			"tradeId":    form.TradeId,
			"playerId":   form.PlayerId,
			"Level":      form.Level,
			"State":      form.State,
			"PageIndex":  form.PageIndex,
			"error":      err,
		}).Error("交易行查询，查询数量异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	respon := &tradeItemListRespon{}
	respon.ItemArray = make([]*tradeItemListResponItem, 0)
	for _, value := range rst {
		item := &tradeItemListResponItem{
			Id:                value.Id,
			Platform:          value.Platform,
			ServerId:          value.ServerId,
			TradeId:           value.TradeId,
			PlayerId:          value.PlayerId,
			PlayerName:        value.PlayerName,
			ItemId:            value.ItemId,
			ItemNum:           value.ItemNum,
			Level:             value.Level,
			Gold:              value.Gold,
			PropertyData:      value.PropertyData,
			Status:            value.Status,
			BuyPlayerPlatform: value.BuyPlayerPlatform,
			BuyPlayerServerId: value.BuyPlayerServerId,
			BuyPlayerId:       value.BuyPlayerId,
			BuyPlayerName:     value.BuyPlayerName,
			UpdateTime:        value.UpdateTime,
			CreateTime:        value.CreateTime,
			DeleteTime:        value.DeleteTime,
		}
		respon.ItemArray = append(respon.ItemArray, item)
	}
	respon.TotalCount = count

	rr := gmhttp.NewSuccessResult(respon)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
