package api

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/charge"
	"net/http"

	fgamehttpputils "fgame/fgame/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

// type getOrderRequest struct {
// 	SDKType  logintypes.SDKType `form:"sdkType" json:"sdkType"`
// 	ServerId int32              `form:"serverId" json:"serverId"`
// 	UserId   int64              `form:"userId" json:"userId"`
// 	PlayerId int64              `form:"playerId" json:"playerId"`
// 	ChargeId int32              `form:"chargeId" json:"chargeId"`
// 	Money    int32              `form:"money" json:"money"`
// }

// type getOrderResponse struct {
// 	OrderId string `json:"orderId"`
// }

// func handleGetOrder(rw http.ResponseWriter, req *http.Request) {
// 	log.WithFields(
// 		log.Fields{
// 			"ip": req.RemoteAddr,
// 		}).Info("charge:获取订单")
// 	form := &getOrderRequest{}
// 	err := httputils.Bind(req, form)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"ip":    req.RemoteAddr,
// 				"error": err,
// 			}).Error("charge:获取订单,解析请求错误")
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	sdkType := form.SDKType
// 	serverId := form.ServerId
// 	userId := form.UserId
// 	playerId := form.PlayerId
// 	chargeId := form.ChargeId
// 	money := form.Money
// 	ctx := req.Context()
// 	chargeService := charge.ChargeServiceInContext(ctx)
// 	obj, err := chargeService.GetOrder(sdkType, serverId, userId, playerId, chargeId, money)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"ip":       req.RemoteAddr,
// 				"sdkType":  sdkType,
// 				"serverId": serverId,
// 				"userId":   userId,
// 				"playerId": playerId,
// 				"chargeId": chargeId,
// 				"money":    money,
// 				"error":    err,
// 			}).Error("charge:获取订单,请求订单失败")
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	result := &getOrderResponse{
// 		OrderId: obj.GetOrderId(),
// 	}
// 	res := fgamehttpputils.NewSuccessResult(result)
// 	httputils.WriteJSON(rw, http.StatusOK, res)

// 	log.WithFields(
// 		log.Fields{
// 			"ip":       req.RemoteAddr,
// 			"sdkType":  sdkType,
// 			"serverId": serverId,
// 			"userId":   userId,
// 			"playerId": playerId,
// 			"chargeId": chargeId,
// 			"money":    money,
// 			"orderId":  obj.GetOrderId(),
// 		}).Info("charge:获取订单")
// }

type getOrderRequest struct {
	SDKType        logintypes.SDKType            `form:"sdkType" json:"sdkType"`
	DeviceType     logintypes.DevicePlatformType `form:"deviceType" json:"deviceType"`
	ServerId       int32                         `form:"serverId" json:"serverId"`
	PlatformUserId string                        `form:"platformUserId" json:"platformUserId"`
	UserId         int64                         `form:"userId" json:"userId"`
	PlayerId       int64                         `form:"playerId" json:"playerId"`
	ChargeId       int32                         `form:"chargeId" json:"chargeId"`
	Money          int32                         `form:"money" json:"money"`
	Name           string                        `form:"name" json:"name"`
	PlayerLevel    int32                         `form:"playerLevel" json:"playerLevel"`
	Gold           int32                         `form:"gold" json:"gold"`
}

type getOrderResponse struct {
	OrderId    string `json:"orderId"`
	SdkOrderId string `json:"sdkOrderId"`
	NotifyUrl  string `json:"notifyUrl"`
	Extension  string `json:"extension"`
}

func handleGetOrder(rw http.ResponseWriter, req *http.Request) {
	log.WithFields(
		log.Fields{
			"ip": req.RemoteAddr,
		}).Info("charge:获取订单")
	form := &getOrderRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Error("charge:获取订单,解析请求错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	sdkType := logintypes.SDKType(form.SDKType)
	deviceType := logintypes.DevicePlatformType(form.DeviceType)

	serverId := form.ServerId
	userId := form.UserId
	platformUser := form.PlatformUserId
	playerId := form.PlayerId
	playerLevel := form.PlayerLevel
	chargeId := form.ChargeId
	playerName := form.Name
	gold := form.Gold
	money := form.Money
	ctx := req.Context()
	chargeService := charge.ChargeServiceInContext(ctx)
	obj, err := chargeService.GetOrder(sdkType, deviceType, serverId, userId, playerId, playerLevel, playerName, chargeId, money, gold)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":          req.RemoteAddr,
				"sdkType":     sdkType,
				"serverId":    serverId,
				"userId":      userId,
				"playerId":    playerId,
				"playerLevel": playerLevel,
				"playerName":  playerName,
				"chargeId":    chargeId,
				"money":       money,
				"error":       err,
			}).Error("charge:获取订单,请求订单失败")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	//判断是否有第三方下单
	getOrderHandler := charge.GetGetSDKOrderHandler(sdkType)
	var thirdOrderId string
	var notifyUrl string
	var extension string
	if getOrderHandler != nil {
		orderId := obj.GetOrderId()
		tnotifyUrl, tthirdOrderId, textension, flag := getOrderHandler.GetSDKOrder(deviceType, platformUser, chargeId, money, playerId, playerName, serverId, orderId)
		if !flag {
			//TODO 修改下单失败
			log.WithFields(
				log.Fields{
					"ip":           req.RemoteAddr,
					"sdkType":      sdkType,
					"serverId":     serverId,
					"userId":       userId,
					"platformUser": platformUser,
					"playerId":     playerId,
					"chargeId":     chargeId,
					"money":        money,
					"error":        err,
				}).Error("charge:获取订单,请求订单失败")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		thirdOrderId = tthirdOrderId
		notifyUrl = tnotifyUrl
		extension = textension
		//TODO 是否需要记录第三方orderId
	}

	result := &getOrderResponse{
		OrderId:    obj.GetOrderId(),
		SdkOrderId: thirdOrderId,
		NotifyUrl:  notifyUrl,
		Extension:  extension,
	}
	res := fgamehttpputils.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, res)

	log.WithFields(
		log.Fields{
			"ip":         req.RemoteAddr,
			"sdkType":    sdkType,
			"serverId":   serverId,
			"userId":     userId,
			"playerId":   playerId,
			"chargeId":   chargeId,
			"money":      money,
			"orderId":    obj.GetOrderId(),
			"sdkOrderId": thirdOrderId,
			"notifyUrl":  notifyUrl,
			"extension":  extension,
		}).Info("charge:获取订单")
}
