package api

// type testChargeRequest struct {
// 	SDKType        logintypes.SDKType `form:"sdkType" json:"sdkType"`
// 	PlatformUserId string             `form:"platformUserId" json:"platformUserId"`
// 	ServerId       int32              `form:"serverId" json:"serverId"`
// 	UserId         int64              `form:"userId" json:"userId"`
// 	PlayerId       int64              `form:"playerId" json:"playerId"`
// 	ChargeId       int32              `form:"chargeId" json:"chargeId"`
// 	Money          int32              `form:"money" json:"money"`
// }

// type testChargeResponse struct {
// 	OrderId string `json:"orderId"`
// }

// func handleTestCharge(rw http.ResponseWriter, req *http.Request) {
// 	log.WithFields(
// 		log.Fields{
// 			"ip": req.RemoteAddr,
// 		}).Info("charge:测试充值")
// 	form := &testChargeRequest{}
// 	err := httputils.Bind(req, form)
// 	if err != nil {
// 		log.WithFields(
// 			log.Fields{
// 				"ip":    req.RemoteAddr,
// 				"error": err,
// 			}).Error("charge:测试充值,解析请求错误")
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	sdkType := form.SDKType
// 	serverId := form.ServerId
// 	userId := form.UserId
// 	platformUserId := form.PlatformUserId
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
// 			}).Error("charge:测试充值,请求订单失败")
// 		rw.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	now := timeutils.TimeToMillisecond(time.Now())
// 	obj, err = chargeService.OrderPay(obj.GetOrderId(), "test", obj.GetSdkType(), serverId, platformUserId, now)
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
// 			}).Error("charge:测试充值,付款错误")
// 		rw.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	//放入回调队列中
// 	remoteService := remote.RemoteServiceInContext(ctx)
// 	flag := remoteService.Charge(obj)
// 	if !flag {
// 		panic(fmt.Errorf("charge:添加到回调队列应该成功"))
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
// 		}).Info("charge:测试充值")
// }
