package api

import (
	"fmt"
	"net/http"

	"fgame/fgame/coupon_server/exchange"
	"fgame/fgame/coupon_server/remote"
	fgamehttpputils "fgame/fgame/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type exchangeRequest struct {
	Code string `form:"code" json:"code"`
	WxId string `form:"wxId" json:"wxId"`
}

type exchangeResponse struct {
	OrderId string `form:"orderId" json:"orderId"`
	Money   int32  `form:"money" json:"money"`
	Code    string `form:"code" json:"code"`
	WxId    string `form:"wxId" json:"wxId"`
}

func handleExchange(rw http.ResponseWriter, req *http.Request) {

	form := &exchangeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Error("exchange:兑换,解析请求错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	code := form.Code
	wxId := form.WxId

	log.WithFields(
		log.Fields{
			"ip":   req.RemoteAddr,
			"code": code,
			"wxId": wxId,
		}).Info("exchange:兑换")
	ctx := req.Context()
	exchangeService := exchange.ExchangeServiceInContext(ctx)
	obj, err := exchangeService.Exchange(code, wxId)
	if err != nil {
		codeErr, ok := err.(exchange.ExchangeError)
		if !ok {
			log.WithFields(
				log.Fields{
					"ip":   req.RemoteAddr,
					"wxId": wxId,
					"code": code,

					"error": err,
				}).Error("exchange:兑换,兑换失败")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := fgamehttpputils.NewFailedResultWithMsg(int(codeErr.Code()), codeErr.Code().String())
		httputils.WriteJSON(rw, http.StatusOK, res)
		return
	}
	var result *exchangeResponse
	if obj == nil {
		result = &exchangeResponse{
			OrderId: "",
		}
	} else {
		result = &exchangeResponse{
			OrderId: obj.GetOrderId(),
			Money:   obj.GetMoney(),
			WxId:    obj.GetWxId(),
			Code:    obj.GetCode(),
		}

		//放入回调队列中
		remoteService := remote.RemoteServiceInContext(ctx)
		flag := remoteService.Exchange(obj)
		if !flag {
			panic(fmt.Errorf("exchange:添加到回调队列应该成功"))
		}

	}

	res := fgamehttpputils.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, res)

	log.WithFields(
		log.Fields{
			"ip":   req.RemoteAddr,
			"wxId": wxId,
			"code": code,
		}).Info("exchange:兑换")
}
