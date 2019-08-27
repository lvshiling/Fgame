package api

import (
	"net/http"

	"fgame/fgame/coupon_server/exchange"
	fgamehttpputils "fgame/fgame/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type exchangeCodeRequest struct {
	ServerId int32 `form:"serverId" json:"serverId"`
	Platform int32 `form:"platform" json:"platform"`

	PlayerId int64 `form:"playerId" json:"playerId"`

	ExchangeId  int64 `form:"money" json:"exchangeId"`
	Money       int32 `form:"money" json:"money"`
	ExpiredTime int64 `form:"money" json:"expiredTime"`
}

type exchangeCodeResponse struct {
	Code string `form:"code" json:"code"`
}

func handleExchangeCode(rw http.ResponseWriter, req *http.Request) {

	form := &exchangeCodeRequest{}
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

	serverId := form.ServerId

	playerId := form.PlayerId

	platform := form.Platform

	exchangeId := form.ExchangeId
	money := form.Money
	expiredTime := form.ExpiredTime
	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"platform": platform,

			"serverId": serverId,

			"playerId": playerId,

			"exchangeId":  exchangeId,
			"money":       money,
			"expiredTime": expiredTime,
		}).Info("exchange:产生兑换码")
	ctx := req.Context()
	exchangeService := exchange.ExchangeServiceInContext(ctx)
	obj, err := exchangeService.GenerateCode(platform, serverId, playerId, exchangeId, money, expiredTime)
	if err != nil {
		codeErr, ok := err.(exchange.ExchangeError)
		if !ok {
			log.WithFields(
				log.Fields{
					"ip":       req.RemoteAddr,
					"platform": platform,

					"serverId": serverId,

					"playerId": playerId,

					"exchangeId":  exchangeId,
					"money":       money,
					"expiredTime": expiredTime,
					"error":       err,
				}).Error("exchange:产生兑换码,兑换失败")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := fgamehttpputils.NewFailedResultWithMsg(int(codeErr.Code()), codeErr.Code().String())
		httputils.WriteJSON(rw, http.StatusOK, res)
		return
	}

	result := &exchangeCodeResponse{
		Code: obj.GetCode(),
	}
	res := fgamehttpputils.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, res)

	log.WithFields(
		log.Fields{
			"ip":       req.RemoteAddr,
			"platform": platform,

			"serverId": serverId,

			"playerId": playerId,

			"exchangeId":  exchangeId,
			"money":       money,
			"expiredTime": expiredTime,
		}).Info("coupon:产生兑换码")
}
