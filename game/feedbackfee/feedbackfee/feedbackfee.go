package feedbackfee

import (
	"encoding/json"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/coupon/coupon"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	feedbackfeedao "fgame/fgame/game/feedbackfee/dao"
	feebackfeeeventypes "fgame/fgame/game/feedbackfee/event/types"
	feebackfeetypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/httputils"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"runtime/debug"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type FeedbackFeeService interface {
	Start()
	Stop()
	Heartbeat()
	Exchange(playerId int64, exchangeId int64, money int32, expiredTime int64)
	CodeGenerate()
	GetCodeGenerateList(playerId int64) []*FeedbackExchangeObject
	FillCode(id int64)
	CodeExpireCheck()
	CodeExchange(id int64, playerId int64, code string, money int32)
	GetEndList(playerId int64) []*FeedbackExchangeObject
	ExchangeNotify(id int64)
	CodeExchangeByCode(code string)
}

type feedbackFeeService struct {
	rwm               sync.RWMutex
	playerExchangeMap map[int64]*FeedbackExchangeObject
	statusExchangeMap map[feebackfeetypes.FeedbackExchangeStatus]map[int64]*FeedbackExchangeObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

//初始化
func (s *feedbackFeeService) init() (err error) {
	s.playerExchangeMap = make(map[int64]*FeedbackExchangeObject)
	s.statusExchangeMap = make(map[feebackfeetypes.FeedbackExchangeStatus]map[int64]*FeedbackExchangeObject)
	serverId := global.GetGame().GetServerIndex()
	unfinishFeedbackExchangeEntityList, err := feedbackfeedao.GetFeedbackFeeDao().GetUnfinishFeedbackExchangeList(serverId)
	if err != nil {
		return
	}

	for _, e := range unfinishFeedbackExchangeEntityList {
		obj := newFeedbackExchangeObject()
		err = obj.FromEntity(e)
		if err != nil {
			return
		}
		s.addFeedbackExchangeObj(obj)
	}
	s.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	s.heartbeatRunner.AddTask(createCodeExpireTask(s))
	s.heartbeatRunner.AddTask(createCodeGenerateTask(s))

	return
}

//初始化
func (s *feedbackFeeService) Start() {

	return
}

//初始化
func (s *feedbackFeeService) Stop() {

	return
}

func (s *feedbackFeeService) addFeedbackExchangeObj(obj *FeedbackExchangeObject) {
	exchangeMap, ok := s.statusExchangeMap[obj.status]
	if !ok {
		exchangeMap = make(map[int64]*FeedbackExchangeObject)
		s.statusExchangeMap[obj.status] = exchangeMap
	}
	exchangeMap[obj.id] = obj
	s.playerExchangeMap[obj.exchangeId] = obj
}

func (s *feedbackFeeService) removeFeedbackExchangeObj(obj *FeedbackExchangeObject) {
	exchangeMap, ok := s.statusExchangeMap[obj.status]
	if !ok {
		return
	}
	delete(exchangeMap, obj.id)
	delete(s.playerExchangeMap, obj.exchangeId)
}

func (s *feedbackFeeService) getFeedbackExchangeObj(status feebackfeetypes.FeedbackExchangeStatus, id int64) *FeedbackExchangeObject {
	exchangeMap, ok := s.statusExchangeMap[status]
	if !ok {
		return nil
	}
	obj, ok := exchangeMap[id]
	if !ok {
		return nil
	}
	return obj
}

func (s *feedbackFeeService) getFeedbackExchangeObjByCode(code string) *FeedbackExchangeObject {
	if len(code) == 0 {
		return nil
	}
	for _, exchangeMap := range s.statusExchangeMap {
		for _, obj := range exchangeMap {
			if obj.code == code {
				return obj
			}
		}
	}

	return nil
}

func (s *feedbackFeeService) Exchange(playerId int64, exchangeId int64, money int32, expiredTime int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	//已经上传过
	feedbackExchange, ok := s.playerExchangeMap[exchangeId]
	if ok {
		return
	}
	//检查有没有上传过
	now := global.GetGame().GetTimeService().Now()
	//生成新的代码
	feedbackExchange = newFeedbackExchangeObject()
	feedbackExchange.id, _ = idutil.GetId()
	feedbackExchange.exchangeId = exchangeId
	feedbackExchange.playerId = playerId
	feedbackExchange.serverId = global.GetGame().GetServerIndex()
	feedbackExchange.money = money
	feedbackExchange.status = feebackfeetypes.FeedbackExchangeStatusInit
	feedbackExchange.expiredTime = expiredTime
	feedbackExchange.createTime = now
	feedbackExchange.SetModified()
	s.addFeedbackExchangeObj(feedbackExchange)
	s.asyncCodeGenerate(feedbackExchange)
	return
}

func (s *feedbackFeeService) CodeGenerate() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	initExchangeMap, ok := s.statusExchangeMap[feebackfeetypes.FeedbackExchangeStatusInit]
	if !ok {
		return
	}
	for _, obj := range initExchangeMap {
		s.asyncCodeGenerate(obj)
	}
	return
}

func (s *feedbackFeeService) GetCodeGenerateList(playerId int64) (codeGenerateList []*FeedbackExchangeObject) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	codeExchangeMap, ok := s.statusExchangeMap[feebackfeetypes.FeedbackExchangeStatusGenerateCode]
	if !ok {
		return
	}
	for _, obj := range codeExchangeMap {
		if obj.playerId == playerId {
			codeGenerateList = append(codeGenerateList, obj)
		}
	}
	return codeGenerateList
}

const (
	codePath   = "/api/exchange/code"
	expirePath = "/api/exchange/expire"
)

func (s *feedbackFeeService) asyncCodeGenerate(obj *FeedbackExchangeObject) {
	host := coupon.GetCouponService().GetHost()
	port := coupon.GetCouponService().GetPort()

	go func(host string, port int32, obj *FeedbackExchangeObject) {
		var err error
		platform := global.GetGame().GetPlatform()
		serverId := obj.serverId
		playerId := obj.playerId
		exchangeId := obj.id
		money := obj.money
		expiredTime := obj.expiredTime
		log.WithFields(
			log.Fields{
				"platform":    platform,
				"serverId":    serverId,
				"playerId":    playerId,
				"exchangeId":  exchangeId,
				"money":       money,
				"expiredTime": expiredTime,
			}).Info("exchange:生成兑换码中")
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"platform":    platform,
						"serverId":    serverId,
						"playerId":    playerId,
						"exchangeId":  exchangeId,
						"money":       money,
						"expiredTime": expiredTime,
						"error":       r,
						"stack":       exceptionContent,
					}).Error("exchange:兑换兑换码,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()

		type exchangeForm struct {
			Platform    int32 `json:"platform"`
			ServerId    int32 `json:"serverId"`
			PlayerId    int64 `json:"playerId"`
			ExchangeId  int64 `json:"exchangeId"`
			Money       int32 `json:"money"`
			ExpiredTime int64 `json:"expiredTime"`
		}

		getExchangePath := fmt.Sprintf("http://%s:%d%s", host, port, codePath)
		form := &exchangeForm{
			Platform:    platform,
			ServerId:    serverId,
			ExchangeId:  exchangeId,
			PlayerId:    playerId,
			Money:       money,
			ExpiredTime: expiredTime,
		}

		result, err := httputils.PostJsonWithRawMessage(getExchangePath, nil, form)
		if err != nil {
			log.WithFields(
				log.Fields{
					"platform":    platform,
					"serverId":    serverId,
					"playerId":    playerId,
					"exchangeId":  exchangeId,
					"money":       money,
					"expiredTime": expiredTime,
					"err":         err.Error(),
				}).Error("exchange:兑换兑换码,生成码错误")
			return
		}
		if result.ErrorCode != 0 {
			log.WithFields(
				log.Fields{
					"platform":    platform,
					"serverId":    serverId,
					"playerId":    playerId,
					"exchangeId":  exchangeId,
					"money":       money,
					"expiredTime": expiredTime,
					"errorCode":   result.ErrorCode,
					"errorMsg":    result.ErrorMsg,
				}).Warn("feedbackfee:兑换,错误")
			// s.exchangeFailed(playerId, int32(result.ErrorCode), result.ErrorMsg)
			return
		}

		type getExchangeResponse struct {
			Code string `json:"code"`
		}
		res := &getExchangeResponse{}
		err = json.Unmarshal(result.Result, res)
		if err != nil {
			return
		}
		s.codeGenerateFinish(exchangeId, res.Code)

	}(host, port, obj)

}

//码产生
func (s *feedbackFeeService) codeGenerateFinish(id int64, code string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getFeedbackExchangeObj(feebackfeetypes.FeedbackExchangeStatusInit, id)
	if obj == nil {
		return
	}
	s.removeFeedbackExchangeObj(obj)
	flag := obj.CodeGenerate(code)
	if !flag {
		log.WithFields(
			log.Fields{
				"id":         id,
				"playerId":   obj.playerId,
				"serverId":   obj.serverId,
				"exchangeId": obj.exchangeId,
				"code":       obj.code,
				"money":      obj.money,
				"status":     obj.status,
			}).Warn("feedbackfee:代码产生失败")
		return
	}
	s.addFeedbackExchangeObj(obj)
	//发送事件
	gameevent.Emit(feebackfeeeventypes.EventTypeCodeGenerate, obj, nil)
}

func (s *feedbackFeeService) FillCode(id int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getFeedbackExchangeObj(feebackfeetypes.FeedbackExchangeStatusGenerateCode, id)
	if obj == nil {
		return
	}
	s.removeFeedbackExchangeObj(obj)
	flag := obj.FillCode()
	if !flag {
		log.WithFields(
			log.Fields{
				"id":         id,
				"playerId":   obj.playerId,
				"serverId":   obj.serverId,
				"exchangeId": obj.exchangeId,
				"code":       obj.code,
				"money":      obj.money,
				"status":     obj.status,
			}).Warn("feedbackfee:填充代码失败")
		return
	}

	s.addFeedbackExchangeObj(obj)
	return
}

func (s *feedbackFeeService) CodeExpireCheck() {
	processExchangeMap, ok := s.statusExchangeMap[feebackfeetypes.FeedbackExchangeStatusProcess]
	if !ok {
		return
	}
	for _, obj := range processExchangeMap {
		s.asyncCodeExpire(obj)
	}
	return
}

func (s *feedbackFeeService) asyncCodeExpire(obj *FeedbackExchangeObject) {
	now := global.GetGame().GetTimeService().Now()
	if now < obj.expiredTime {
		return
	}
	host := coupon.GetCouponService().GetHost()
	port := coupon.GetCouponService().GetPort()

	go func(host string, port int32, obj *FeedbackExchangeObject) {
		var err error
		platform := global.GetGame().GetPlatform()
		serverId := obj.serverId
		playerId := obj.playerId
		exchangeId := obj.id
		money := obj.money
		expiredTime := obj.expiredTime
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"platform":    platform,
						"serverId":    serverId,
						"playerId":    playerId,
						"exchangeId":  exchangeId,
						"money":       money,
						"expiredTime": expiredTime,
						"error":       r,
						"stack":       exceptionContent,
					}).Error("feedbackfee:兑换过期,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()

		type exchangeForm struct {
			ExchangeId int64 `json:"exchangeId"`
		}

		getExchangePath := fmt.Sprintf("http://%s:%d%s", host, port, expirePath)
		form := &exchangeForm{
			ExchangeId: exchangeId,
		}

		result, err := httputils.PostJsonWithRawMessage(getExchangePath, nil, form)
		if err != nil {
			return
		}
		if result.ErrorCode != 0 {
			log.WithFields(
				log.Fields{
					"platform":    platform,
					"serverId":    serverId,
					"playerId":    playerId,
					"exchangeId":  exchangeId,
					"money":       money,
					"expiredTime": expiredTime,
					"errorCode":   result.ErrorCode,
					"errorMsg":    result.ErrorMsg,
				}).Warn("feedbackfee:过期,错误")
			// s.exchangeFailed(playerId, int32(result.ErrorCode), result.ErrorMsg)
			return
		}

		type getExchangeResponse struct {
			ExchangeId int64 `json:"exchangeId"`
		}
		res := &getExchangeResponse{}
		err = json.Unmarshal(result.Result, res)
		if err != nil {
			return
		}
		s.ExchangeExpired(obj.id)

	}(host, port, obj)

}

//兑换过期
func (s *feedbackFeeService) ExchangeExpired(id int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getFeedbackExchangeObj(feebackfeetypes.FeedbackExchangeStatusProcess, id)
	if obj == nil {
		return
	}
	s.removeFeedbackExchangeObj(obj)
	flag := obj.Expired()
	if !flag {
		log.WithFields(
			log.Fields{
				"id":         id,
				"playerId":   obj.playerId,
				"serverId":   obj.serverId,
				"exchangeId": obj.exchangeId,
				"code":       obj.code,
				"money":      obj.money,
				"status":     obj.status,
			}).Warn("feedbackfee:兑换过期失败")
		return
	}

	s.addFeedbackExchangeObj(obj)
	//发送事件
	gameevent.Emit(feebackfeeeventypes.EventTypeCodeExpire, obj, nil)
	return
}

//兑换成功
func (s *feedbackFeeService) CodeExchange(id int64, playerId int64, code string, money int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getFeedbackExchangeObj(feebackfeetypes.FeedbackExchangeStatusProcess, id)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"id": id,
			}).Warn("feedbackfee:已经完成")
		return
	}
	s.removeFeedbackExchangeObj(obj)
	flag := obj.Finish()
	if !flag {
		log.WithFields(
			log.Fields{
				"id":         id,
				"playerId":   obj.playerId,
				"serverId":   obj.serverId,
				"exchangeId": obj.exchangeId,
				"code":       obj.code,
				"money":      obj.money,
				"status":     obj.status,
			}).Warn("feedbackfee:兑换失败")
		return
	}

	s.addFeedbackExchangeObj(obj)
	//发送事件
	gameevent.Emit(feebackfeeeventypes.EventTypeCodeExchange, obj, nil)
	return
}

func (s *feedbackFeeService) GetEndList(playerId int64) (endList []*FeedbackExchangeObject) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	failExchangeMap, ok := s.statusExchangeMap[feebackfeetypes.FeedbackExchangeStatusFailed]
	if ok {
		for _, obj := range failExchangeMap {
			if obj.playerId == playerId {
				endList = append(endList, obj)
			}
		}
	}
	finishExchangeMap, ok := s.statusExchangeMap[feebackfeetypes.FeedbackExchangeStatusFinish]
	if ok {
		for _, obj := range finishExchangeMap {
			if obj.playerId == playerId {
				endList = append(endList, obj)
			}
		}
	}
	return
}

//通知
func (s *feedbackFeeService) ExchangeNotify(id int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getFeedbackExchangeObj(feebackfeetypes.FeedbackExchangeStatusFailed, id)
	if obj == nil {
		obj = s.getFeedbackExchangeObj(feebackfeetypes.FeedbackExchangeStatusFinish, id)
		if obj == nil {
			return
		}
	}
	s.removeFeedbackExchangeObj(obj)
	flag := obj.Notify()
	if !flag {
		log.WithFields(
			log.Fields{
				"id":         id,
				"playerId":   obj.playerId,
				"serverId":   obj.serverId,
				"exchangeId": obj.exchangeId,
				"code":       obj.code,
				"money":      obj.money,
				"status":     obj.status,
			}).Warn("feedbackfee:通知失败")
		return
	}
	return
}

//通知
func (s *feedbackFeeService) CodeExchangeByCode(code string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	obj := s.getFeedbackExchangeObjByCode(code)
	if obj == nil {
		return
	}
	if obj.status != feebackfeetypes.FeedbackExchangeStatusProcess {
		return
	}
	s.removeFeedbackExchangeObj(obj)
	flag := obj.Finish()
	if !flag {
		log.WithFields(
			log.Fields{
				"id":         obj.id,
				"playerId":   obj.playerId,
				"serverId":   obj.serverId,
				"exchangeId": obj.exchangeId,
				"code":       obj.code,
				"money":      obj.money,
				"status":     obj.status,
			}).Warn("feedbackfee:兑换失败")
		return
	}
	s.addFeedbackExchangeObj(obj)
	gameevent.Emit(feebackfeeeventypes.EventTypeCodeExchange, obj, nil)
	return
}

func (s *feedbackFeeService) Heartbeat() {
	s.heartbeatRunner.Heartbeat()
}

var (
	once sync.Once
	cs   *feedbackFeeService
)

func Init() (err error) {
	once.Do(func() {
		cs = &feedbackFeeService{}
		err = cs.init()
	})
	return err
}

func GetFeedbackFeeService() FeedbackFeeService {
	return cs
}
