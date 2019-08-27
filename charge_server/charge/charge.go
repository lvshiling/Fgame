package charge

import (
	"context"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/charge_server/types"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"

	uuid "github.com/satori/go.uuid"
)

func SetupChargeServiceHandler(s ChargeService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithChargeService(ctx, s)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

type ChargeService interface {
	//获取订单
	GetOrder(typ logintypes.SDKType, devicePlatform logintypes.DevicePlatformType, serverId int32, userId int64, playerId int64, playerLevel int32, playerName string, chargeId int32, money int32, gold int32) (obj *OrderObject, err error)
	//充值成功
	OrderPay(orderId string, sdkOrderId string, typ logintypes.SDKType, money int32, platformUserId string, receiveTime int64) (obj *OrderObject, repeat bool, err error)
	//充值失败
	OrderFail(orderId string, typ logintypes.SDKType) (err error)
	//TODO 缓存防攻击
}

type contextKey string

const (
	chargeServiceKey contextKey = "fgame.charge"
)

func WithChargeService(parent context.Context, s ChargeService) context.Context {
	ctx := context.WithValue(parent, chargeServiceKey, s)
	return ctx
}

func ChargeServiceInContext(parent context.Context) ChargeService {
	s := parent.Value(chargeServiceKey)
	if s == nil {
		return nil
	}
	ts, ok := s.(ChargeService)
	if !ok {
		return nil
	}
	return ts
}

type chargeService struct {
	db coredb.DBService
	rs coreredis.RedisService
}

//获取订单
func (s *chargeService) GetOrder(typ logintypes.SDKType, devicePlatform logintypes.DevicePlatformType, serverId int32, userId int64, playerId int64, playerLevel int32, playerName string, chargeId int32, money int32, gold int32) (obj *OrderObject, err error) {
	obj = &OrderObject{}
	str := uuid.NewV4().String()
	orderId := strings.Replace(str, "-", "", -1)
	obj.orderId = orderId
	obj.userId = userId
	obj.serverId = serverId
	obj.sdkType = typ
	obj.devicePlatform = devicePlatform
	obj.playerId = playerId
	obj.playerLevel = playerLevel
	obj.playerName = playerName
	obj.chargeId = chargeId
	obj.money = money
	obj.gold = gold
	obj.status = types.OrderStatusInit
	now := timeutils.TimeToMillisecond(time.Now())
	obj.createTime = now
	e := obj.ToEntity()
	if err = s.db.DB().Save(e).Error; err != nil {
		return
	}
	return
}

func (s *chargeService) OrderPay(orderId string, sdkOrderId string, typ logintypes.SDKType, money int32, platformUserId string, receiveTime int64) (obj *OrderObject, repeat bool, err error) {
	orderEntity := &OrderEntity{}
	if err = s.db.DB().First(orderEntity, "orderId=?", orderId).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, false, nil
	}

	obj = NewOrderObject()
	obj.FromEntity(orderEntity)

	if obj.money != money {
		log.WithFields(
			log.Fields{
				"orderId":     orderId,
				"sdkOrderId":  sdkOrderId,
				"actualMoney": obj.money,
				"money":       money,
			}).Warn("charge:充值,金额错误")
		obj = nil
		return
	}

	//已经充值了
	if obj.status != types.OrderStatusInit {
		repeat = true
		return
	}

	obj.platformUserId = platformUserId
	obj.status = types.OrderStatusPay
	obj.sdkOrderId = sdkOrderId
	obj.receivePayTime = receiveTime
	now := timeutils.TimeToMillisecond(time.Now())
	obj.updateTime = now
	e := obj.ToEntity()
	if err = s.db.DB().Save(e).Error; err != nil {
		return
	}

	return
}

func (s *chargeService) OrderFail(orderId string, typ logintypes.SDKType) (err error) {
	orderEntity := &OrderEntity{}
	if err = s.db.DB().First(orderEntity, "orderId=?  and sdkType=?", orderId, int32(typ)).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil
	}
	//TODO 判断其它参数
	obj := NewOrderObject()
	obj.FromEntity(orderEntity)
	//已经充值了
	if obj.status != types.OrderStatusInit {
		return
	}

	obj.status = types.OrderStatusFail
	now := timeutils.TimeToMillisecond(time.Now())
	obj.updateTime = now
	e := obj.ToEntity()
	if err = s.db.DB().Save(e).Error; err != nil {
		return
	}
	return
}

func NewChargeService(db coredb.DBService, rs coreredis.RedisService) ChargeService {
	s := &chargeService{
		db: db,
		rs: rs,
	}
	return s
}
