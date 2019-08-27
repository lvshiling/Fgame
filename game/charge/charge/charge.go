package charge

import (
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/charge/dao"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	chargetemplate "fgame/fgame/game/charge/template"
	"fgame/fgame/game/charge/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/httputils"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"runtime/debug"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type ChargeConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

//充值活动
type ChargeService interface {
	//获取首冲时间
	GetChargeTime() *FirstChargeObject
	// 重置首冲时间
	GMResetChargeTime(chargeTime int64)
	//充值
	Charge(orderId string, playerId int64, chargeId int32, money int32) (flag bool, err error)
	//获取订单
	GetOrder(pl player.Player, chargeId int32) bool
	FinishOrder(pl player.Player, orderId string)
	//获取未完成订单
	GetUnfinishOrderList(pl player.Player) []*OrderObject
	// 获取后台充值未完成列表
	GetUnfinishPrivilegeChargeList(pl player.Player) []*PrivilegeChargeObject
	// 获取新首充活动开始时间和持续时间
	GetNewFirstChargeTime() (startTime int64, duration int64)
	// 添加后台充值
	PrivilegeCharge(playerId, goldNum int64)
	// 完成后台充值
	FinishPriviCharge(pl player.Player, privilegeId int64)
	// 是否在新首充时间之内
	IsNewFirstChargeDuration() bool

	// Gm 设置新首充开始时间
	GmSetNewFirstChargeTime(startTime int64) bool
	SetNewFirstChargeTime(startTime int64) bool
	ResetFirstCharge() bool
}

type chargeService struct {
	//读写锁
	rwm sync.RWMutex
	cfg *ChargeConfig
	//废弃:zrc
	//首冲时间
	firstChargeObj *FirstChargeObject
	//新首充对象
	newFirstChargeObj          *NewFirstChargeObject
	unfinishOrderListMap       map[int64][]*OrderObject
	unfinishPrivilegeChargeMap map[int64][]*PrivilegeChargeObject
}

//初始化
func (s *chargeService) init(cfg *ChargeConfig) (err error) {
	s.cfg = cfg
	if err = s.initFirstCharge(); err != nil {
		return
	}
	if err = s.initUnfinishedOrders(); err != nil {
		return
	}

	if err = s.initUnfinishedPrivilegeCharges(); err != nil {
		return
	}
	if err = s.initNewFirstCharge(); err != nil {
		return
	}
	return
}

//初始化首冲
func (s *chargeService) initFirstCharge() (err error) {
	serverId := global.GetGame().GetServerIndex()
	firstChargeEntity, err := dao.GetChargeDao().GetFirstCharge(serverId)
	if err != nil {
		return
	}
	if firstChargeEntity == nil {
		s.initFirstChargeObj()
	} else {
		obj := newFirstChargeObject()
		obj.FromEntity(firstChargeEntity)
		s.firstChargeObj = obj
	}

	//合服重置首冲
	if merge.GetMergeService().IsMerge() {
		now := global.GetGame().GetTimeService().Now()
		s.firstChargeObj.ChargeTime = now
		s.firstChargeObj.SetModified()
	}
	return
}

func (s *chargeService) initNewFirstCharge() (err error) {
	serverId := global.GetGame().GetServerIndex()
	newFirstChargeEntity, err := dao.GetChargeDao().GetNewFirstCharge(serverId)
	if err != nil {
		return
	}
	if newFirstChargeEntity == nil {
		s.initNewFirstChargeObj()
	} else {
		obj := newNewFirstChargeObject()
		obj.FromEntity(newFirstChargeEntity)
		s.newFirstChargeObj = obj
	}
	s.resetFirstCharge()
	//合服重置首冲
	// if merge.GetMergeService().IsMerge() {
	// 	now := global.GetGame().GetTimeService().Now()
	// 	s.newFirstChargeObj.startTime = now
	// 	s.newFirstChargeObj.updateTime = now
	// 	s.newFirstChargeObj.SetModified()
	// }
	return
}

func (s *chargeService) ResetFirstCharge() bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.resetFirstCharge()
}

func (s *chargeService) resetFirstCharge() bool {
	now := global.GetGame().GetTimeService().Now()
	mergeTime := merge.GetMergeService().GetMergeTime()
	if mergeTime != 0 {
		beginTime, _ := timeutils.BeginOfNow(mergeTime)
		if s.newFirstChargeObj.startTime < beginTime {
			s.newFirstChargeObj.startTime = beginTime
			s.newFirstChargeObj.updateTime = now
			s.newFirstChargeObj.SetModified()
			return true
		}
	} else {
		startTime := center.GetCenterService().GetStartTime()
		beginTime, _ := timeutils.BeginOfNow(startTime)
		if s.newFirstChargeObj.startTime < beginTime {
			s.newFirstChargeObj.startTime = beginTime
			s.newFirstChargeObj.updateTime = now
			s.newFirstChargeObj.SetModified()
			return true
		}
	}
	return false
}

func (s *chargeService) initUnfinishedOrders() (err error) {
	s.unfinishOrderListMap = make(map[int64][]*OrderObject)

	serverId := global.GetGame().GetServerIndex()
	orderEntityList, err := dao.GetChargeDao().GetOrderList(serverId, types.OrderStatusInit)
	if err != nil {
		return
	}

	//初始化未完成订单列表
	for _, orderEntity := range orderEntityList {
		orderObj := newOrderObject()
		err = orderObj.FromEntity(orderEntity)
		if err != nil {
			return err
		}
		s.unfinishOrderListMap[orderObj.GetPlayerId()] = append(s.unfinishOrderListMap[orderObj.GetPlayerId()], orderObj)
	}
	return
}

func (s *chargeService) initUnfinishedPrivilegeCharges() (err error) {
	s.unfinishPrivilegeChargeMap = make(map[int64][]*PrivilegeChargeObject)

	serverId := global.GetGame().GetServerIndex()
	privilegeEntityList, err := dao.GetChargeDao().GetPrivilegeChargeList(serverId, types.OrderStatusInit)
	if err != nil {
		return
	}

	//初始化未完成后台充值列表
	for _, entity := range privilegeEntityList {
		obj := newPrivilegeChargeObject()
		obj.FromEntity(entity)
		s.unfinishPrivilegeChargeMap[obj.GetPlayerId()] = append(s.unfinishPrivilegeChargeMap[obj.GetPlayerId()], obj)
	}
	return
}

func (s *chargeService) initFirstChargeObj() {
	fco := newFirstChargeObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	fco.Id = id
	fco.CreateTime = now
	fco.ChargeTime = 0
	fco.SetModified()
	s.firstChargeObj = fco
}

func (s *chargeService) initNewFirstChargeObj() {
	obj := newNewFirstChargeObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	serverId := global.GetGame().GetServerIndex()
	obj.id = id
	obj.serverId = serverId
	obj.createTime = now
	obj.SetModified()
	s.newFirstChargeObj = obj
}

func (s *chargeService) GetChargeTime() *FirstChargeObject {
	return s.firstChargeObj
}

func (s *chargeService) GMResetChargeTime(chargeTime int64) {
	now := global.GetGame().GetTimeService().Now()
	s.firstChargeObj.ChargeTime = chargeTime
	s.firstChargeObj.UpdateTime = now
	s.firstChargeObj.SetModified()
}

func (s *chargeService) Charge(orderId string, playerId int64, chargeId int32, money int32) (flag bool, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTemplate == nil {
		return false, nil
	}
	serverId := global.GetGame().GetServerIndex()
	//TODO 优化让查询更快
	orderEntity, err := dao.GetChargeDao().GetOrder(serverId, orderId)
	if err != nil {
		return false, err
	}
	orderObj := newOrderObject()
	if orderEntity == nil {
		id, _ := idutil.GetId()
		orderObj.id = id
		orderObj.serverId = serverId
		orderObj.orderStatus = types.OrderStatusInit
		orderObj.orderId = orderId
		orderObj.playerId = playerId
		orderObj.chargeId = chargeId
		orderObj.money = money
		now := global.GetGame().GetTimeService().Now()
		orderObj.createTime = now
		orderObj.SetModified()
	} else {
		log.WithFields(
			log.Fields{
				"orderId":  orderId,
				"PlayerId": playerId,
				"ChargeId": chargeId,
				"Money":    money,
			}).Warn("charge:重复发送订单")
		return true, nil
	}

	s.unfinishOrderListMap[playerId] = append(s.unfinishOrderListMap[playerId], orderObj)

	//发送事件完成充值
	eventData := CreateOrderChargeEventData(chargeId, orderId)
	gameevent.Emit(chargeeventtypes.ChargeEventTypeOrderCharge, playerId, eventData)
	return true, nil
}

func (s *chargeService) PrivilegeCharge(playerId, goldNum int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := newPrivilegeChargeObject()
	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	obj.id = id
	obj.serverId = global.GetGame().GetServerIndex()
	obj.playerId = playerId
	obj.status = types.OrderStatusInit
	obj.goldNum = goldNum
	obj.createTime = now
	obj.SetModified()

	s.unfinishPrivilegeChargeMap[playerId] = append(s.unfinishPrivilegeChargeMap[playerId], obj)

	// 后台充值
	gameevent.Emit(chargeeventtypes.ChargeEventTypePrivilegeCharge, playerId, nil)
	return
}

const (
	orderPath = "/api/charge/get"
)

func (s *chargeService) GetOrder(pl player.Player, chargeId int32) bool {
	chargeTemplate := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTemplate == nil {
		return false
	}
	sdkType := pl.GetSDKType()
	deviceType := pl.GetDevicePlatformType()
	userId := pl.GetUserId()
	platformUserId := pl.GetPlatformUserId()
	playerId := pl.GetId()
	money := chargeTemplate.Rmb
	gold := chargeTemplate.Gold
	serverId := pl.GetServerId()
	playerName := pl.GetName()
	playerLevel := pl.GetLevel()
	s.getOrder(sdkType, deviceType, serverId, platformUserId, userId, playerId, playerLevel, chargeId, money, gold, playerName)
	return true
}

func (s *chargeService) getOrderFailed(playerId int64, chargeId int32) {
	gameevent.Emit(chargeeventtypes.ChargeEventTypeGetOrderFailed, playerId, chargeId)
}

func (s *chargeService) getOrderFinish(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, notifyUrl string, serverId int32, platformUserId string, userId int64, playerId int64, chargeId int32, money int32, name string, orderId string, sdkOrderId string, extension string) {
	eventData := CreateGetOrderFinishEventData(sdkType, deviceType, notifyUrl, serverId, platformUserId, userId, playerId, chargeId, money, name, orderId, sdkOrderId, extension)
	gameevent.Emit(chargeeventtypes.ChargeEventTypeGetOrderFinish, playerId, eventData)
}

func (s *chargeService) getOrder(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, platformUserId string, userId int64, playerId int64, playerLevel int32, chargeId int32, money int32, gold int32, name string) {

	go func(host string, port int32) {
		var err error
		var orderId string
		var sdkOrderId string
		var extension string
		var notifyUrl string
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"SDKType":        sdkType.String(),
						"DeviceType":     deviceType.String(),
						"ServerId":       serverId,
						"PlatformUserId": platformUserId,
						"UserId":         userId,
						"PlayerId":       playerId,
						"ChargeId":       chargeId,
						"Money":          money,
						"playerName":     name,
						"playerLevel":    playerLevel,
						"notifyUrl":      notifyUrl,
						"orderId":        orderId,
						"sdkOrderId":     sdkOrderId,
						"extension":      extension,
						"error":          r,
						"stack":          exceptionContent,
					}).Error("charge:获取订单号,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()
		defer func() {
			if err != nil {
				log.WithFields(
					log.Fields{
						"SDKType":        sdkType.String(),
						"DeviceType":     deviceType.String(),
						"ServerId":       serverId,
						"UserId":         userId,
						"PlatformUserId": platformUserId,
						"PlayerId":       playerId,
						"ChargeId":       chargeId,
						"Money":          money,
						"playerName":     name,
						"playerLevel":    playerLevel,
						"notifyUrl":      notifyUrl,
						"orderId":        orderId,
						"sdkOrderId":     sdkOrderId,
						"extension":      extension,
						"error":          err,
					}).Warn("charge:获取订单号,失败")
				s.getOrderFailed(playerId, chargeId)
				return
			}
			s.getOrderFinish(sdkType, deviceType, notifyUrl, serverId, platformUserId, userId, playerId, chargeId, money, name, orderId, sdkOrderId, extension)
		}()

		type getOrderForm struct {
			SDKType        logintypes.SDKType            `json:"sdkType"`
			DeviceType     logintypes.DevicePlatformType `json:"deviceType"`
			ServerId       int32                         `json:"serverId"`
			PlatformUserId string                        `json:"platformUserId"`
			UserId         int64                         `json:"userId"`
			PlayerId       int64                         `json:"playerId"`
			PlayerLevel    int32                         `json:"playerLevel"`
			ChargeId       int32                         `json:"chargeId"`
			Money          int32                         `json:"money"`
			Gold           int32                         `json:"gold"`
			Name           string                        `json:"name"`
		}

		getOrderPath := fmt.Sprintf("http://%s:%d%s", host, port, orderPath)
		form := &getOrderForm{
			SDKType:        sdkType,
			DeviceType:     deviceType,
			ServerId:       serverId,
			PlatformUserId: platformUserId,
			UserId:         userId,
			PlayerId:       playerId,
			PlayerLevel:    playerLevel,
			ChargeId:       chargeId,
			Gold:           gold,
			Money:          money,
			Name:           name,
		}

		result, err := httputils.PostJsonWithRawMessage(getOrderPath, nil, form)
		if err != nil {
			return
		}
		if result.ErrorCode != 0 {
			err = fmt.Errorf("error_code %d", result.ErrorCode)
			return
		}
		type getOrderResponse struct {
			OrderId    string `json:"orderId"`
			SdkOrderId string `json:"sdkOrderId"`
			NotifyUrl  string `json:"notifyUrl"`
			Extension  string `json:"extension"`
		}
		res := &getOrderResponse{}
		err = json.Unmarshal(result.Result, res)
		if err != nil {
			return
		}
		orderId = res.OrderId
		sdkOrderId = res.SdkOrderId
		notifyUrl = res.NotifyUrl
		extension = res.Extension
	}(s.cfg.Host, s.cfg.Port)
}

func (s *chargeService) FinishOrder(pl player.Player, orderId string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	unfinishOrderList := s.unfinishOrderListMap[pl.GetId()]

	findIndex := -1
	for index, unfinishOrder := range unfinishOrderList {
		if unfinishOrder.GetOrderId() == orderId {
			findIndex = index
			break
		}
	}
	if findIndex < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	orderObj := unfinishOrderList[findIndex]
	orderObj.orderStatus = types.OrderStatusFinish
	orderObj.updateTime = now
	orderObj.SetModified()
	remainOrderList := make([]*OrderObject, 0, 1)
	if findIndex >= 0 {
		remainOrderList = append(remainOrderList, unfinishOrderList[:findIndex]...)
		remainOrderList = append(remainOrderList, unfinishOrderList[findIndex+1:]...)
		s.unfinishOrderListMap[pl.GetId()] = remainOrderList
	}
}

func (s *chargeService) FinishPriviCharge(pl player.Player, privilegeId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	playerId := pl.GetId()
	findIndex := -1
	unfinishPrivilegeList := s.unfinishPrivilegeChargeMap[playerId]
	for index, obj := range unfinishPrivilegeList {
		if obj.id != privilegeId {
			continue
		}

		findIndex = index
		break
	}

	if findIndex < 0 {
		return
	}

	obj := unfinishPrivilegeList[findIndex]
	now := global.GetGame().GetTimeService().Now()
	obj.status = types.OrderStatusFinish
	obj.updateTime = now
	obj.SetModified()

	remainList := make([]*PrivilegeChargeObject, 0, 1)
	remainList = append(remainList, unfinishPrivilegeList[:findIndex]...)
	remainList = append(remainList, unfinishPrivilegeList[findIndex+1:]...)
	s.unfinishPrivilegeChargeMap[playerId] = remainList
}

func (s *chargeService) GetUnfinishOrderList(pl player.Player) []*OrderObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.unfinishOrderListMap[pl.GetId()]
}

func (s *chargeService) GetUnfinishPrivilegeChargeList(pl player.Player) []*PrivilegeChargeObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.unfinishPrivilegeChargeMap[pl.GetId()]
}

func (s *chargeService) GetNewFirstChargeTime() (startTime int64, dur int64) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	startTime = s.newFirstChargeObj.startTime
	dur = int64(duration) * int64(common.DAY)
	return
}

func (s *chargeService) IsNewFirstChargeDuration() bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	now := global.GetGame().GetTimeService().Now()
	startTime := s.newFirstChargeObj.startTime
	endTime := s.newFirstChargeObj.startTime + int64(duration)*int64(common.DAY)
	if now < startTime || now > endTime {
		return false
	}
	return true
}

func (s *chargeService) GmSetNewFirstChargeTime(now int64) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if !s.ifCanReset(now) {
		return false
	}
	return s.setNewFirstChargeTime(now)
	// now := global.GetGame().GetTimeService().Now()
	// s.newFirstChargeObj.startTime = startTime
	// s.newFirstChargeObj.updateTime = now
	// s.newFirstChargeObj.SetModified()

	// gameevent.Emit(chargeeventtypes.ChargeEventTypeNewFirstChargeTimeChangeLog, s.newFirstChargeObj, duration)
}

func (s *chargeService) SetNewFirstChargeTime(now int64) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if !s.ifCanReset(now) {
		return false
	}
	return s.setNewFirstChargeTime(now)
	// now := global.GetGame().GetTimeService().Now()
	// s.newFirstChargeObj.startTime = startTime
	// s.newFirstChargeObj.updateTime = now
	// s.newFirstChargeObj.SetModified()

	// gameevent.Emit(chargeeventtypes.ChargeEventTypeNewFirstChargeTimeChangeLog, s.newFirstChargeObj, duration)
}

func (s *chargeService) ifCanReset(now int64) bool {
	diff, _ := timeutils.DiffDay(now, s.newFirstChargeObj.startTime)
	if diff < duration {
		return false
	}

	return true
}

func (s *chargeService) setNewFirstChargeTime(now int64) bool {
	flag := s.ifCanReset(now)
	if !flag {
		return false
	}

	beginTime, _ := timeutils.BeginOfNow(now)
	s.newFirstChargeObj.startTime = beginTime
	s.newFirstChargeObj.updateTime = now
	s.newFirstChargeObj.SetModified()

	//添加日志
	logObj := newNewFirstChargeLogObject()
	logObj.id, _ = idutil.GetId()
	logObj.createTime = now
	logObj.serverId = global.GetGame().GetServerIndex()
	logObj.SetModified()
	return true
}

const (
	duration = 7
)

var (
	once sync.Once
	s    *chargeService
)

func Init(cfg *ChargeConfig) (err error) {
	once.Do(func() {
		s = &chargeService{}
		err = s.init(cfg)
	})
	return err
}

func GetChargeService() ChargeService {
	return s
}
