package api

import (
	"net/http"

	"fgame/fgame/coupon_server/exchange"
	fgamehttpputils "fgame/fgame/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type expireExchangeRequest struct {
	ExchangeId int64 `form:"exchangeId" json:"exchangeId"`
}

type expireExchangeResponse struct {
	ExchangeId int64 `form:"exchangeId" json:"exchangeId"`
}

func handleExchangeExpire(rw http.ResponseWriter, req *http.Request) {

	form := &expireExchangeRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Error("exchange:产生兑换码,解析请求错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	exchangeId := form.ExchangeId

	log.WithFields(
		log.Fields{
			"ip": req.RemoteAddr,

			"exchangeId": exchangeId,
		}).Info("exchange:产生兑换码")
	ctx := req.Context()
	exchangeService := exchange.ExchangeServiceInContext(ctx)
	obj, err := exchangeService.Expire(exchangeId)
	if err != nil {
		codeErr, ok := err.(exchange.ExchangeError)
		if !ok {
			log.WithFields(
				log.Fields{
					"ip": req.RemoteAddr,

					"exchangeId": exchangeId,
					"error":      err,
				}).Error("exchange:过期,过期失败")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := fgamehttpputils.NewFailedResultWithMsg(int(codeErr.Code()), codeErr.Code().String())
		httputils.WriteJSON(rw, http.StatusOK, res)
		return
	}
	result := &expireExchangeResponse{}
	if obj == nil {
		result = &expireExchangeResponse{
			ExchangeId: 0,
		}
	} else {
		result = &expireExchangeResponse{
			ExchangeId: exchangeId,
		}
	}

	res := fgamehttpputils.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, res)

	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"exchangeId": exchangeId,
		}).Info("exchange:过期")
}
