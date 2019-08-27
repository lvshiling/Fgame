package exchange

import (
	"context"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	couponservertypes "fgame/fgame/coupon_server/types"
	"fgame/fgame/pkg/timeutils"
	"net/http"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/codegangsta/negroni"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func SetupExchangeServiceHandler(s ExchangeService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithExchangeService(ctx, s)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

type ExchangeService interface {
	//生成兑换码
	GenerateCode(platform int32, serverId int32, playerId int64, exchangeId int64, money int32, expiredTime int64) (obj *FeedbackfeeExchangeObject, err error)
	//过期
	Expire(exchangeId int64) (obj *FeedbackfeeExchangeObject, err error)
	//兑换
	Exchange(code string, wxId string) (obj *FeedbackfeeExchangeObject, err error)
}

type contextKey string

const (
	exchangeServiceKey contextKey = "fgame.exchange"
)

func WithExchangeService(parent context.Context, s ExchangeService) context.Context {
	ctx := context.WithValue(parent, exchangeServiceKey, s)
	return ctx
}

func ExchangeServiceInContext(parent context.Context) ExchangeService {
	s := parent.Value(exchangeServiceKey)
	if s == nil {
		return nil
	}
	ts, ok := s.(ExchangeService)
	if !ok {
		return nil
	}
	return ts
}

var (
	cacheSize = 500
)

type exchangeService struct {
	m             sync.Mutex
	db            coredb.DBService
	rs            coreredis.RedisService
	initMap       map[int64]*FeedbackfeeExchangeObject
	codeMap       map[string]*FeedbackfeeExchangeObject
	codeCache     *lru.TwoQueueCache
	exchangeCache *lru.TwoQueueCache
}

func (s *exchangeService) init() (err error) {
	err = s.loadInitCode()
	if err != nil {
		return
	}
	return nil
}

func (s *exchangeService) loadInitCode() (err error) {
	s.codeCache, err = lru.New2Q(cacheSize)
	if err != nil {
		return
	}
	s.exchangeCache, err = lru.New2Q(cacheSize)
	if err != nil {
		return
	}

	s.initMap = make(map[int64]*FeedbackfeeExchangeObject)
	s.codeMap = make(map[string]*FeedbackfeeExchangeObject)

	feedbackfeeExchangeEntityList := make([]*FeedbackfeeExchangeEntity, 0, 8)
	err = s.db.DB().Find(&feedbackfeeExchangeEntityList, "status=?", int32(couponservertypes.ExchangeStatusInit)).Error
	if err != nil {
		return err
	}

	for _, e := range feedbackfeeExchangeEntityList {
		obj := NewFeedbackfeeExchangeObject()
		obj.FromEntity(e)
		s.addCode(obj)
	}

	return
}

func (s *exchangeService) getExchangeByExchangeId(exchangeId int64) (obj *FeedbackfeeExchangeObject, err error) {
	//查找缓存
	exchangeObjInter, ok := s.exchangeCache.Get(exchangeId)
	if ok {
		tobj, ok := exchangeObjInter.(*FeedbackfeeExchangeObject)
		if ok {
			obj = tobj
			return
		}
	}
	feedbackfeeExchangeEntity := &FeedbackfeeExchangeEntity{}
	err = s.db.DB().First(feedbackfeeExchangeEntity, "exchangeId=?", exchangeId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	obj = NewFeedbackfeeExchangeObject()
	obj.FromEntity(feedbackfeeExchangeEntity)
	s.exchangeCache.Add(exchangeId, obj)
	return
}

func (s *exchangeService) getExchangeByCode(code string) (obj *FeedbackfeeExchangeObject, err error) {
	//查找缓存
	exchangeObjInter, ok := s.codeCache.Get(code)
	if ok {
		tobj, ok := exchangeObjInter.(*FeedbackfeeExchangeObject)
		if ok {
			obj = tobj
			return
		}
	}
	feedbackfeeExchangeEntity := &FeedbackfeeExchangeEntity{}
	err = s.db.DB().First(feedbackfeeExchangeEntity, "code=?", code).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		return nil, nil
	}
	obj = NewFeedbackfeeExchangeObject()
	obj.FromEntity(feedbackfeeExchangeEntity)
	s.codeCache.Add(code, obj)
	return
}

const (
	codePrefix = "ws_"
)

//兑换
func (s *exchangeService) GenerateCode(platform int32, serverId int32, playerId int64, exchangeId int64, money int32, expiredTime int64) (obj *FeedbackfeeExchangeObject, err error) {
	//查找
	s.m.Lock()
	defer s.m.Unlock()
	obj, ok := s.initMap[exchangeId]
	if ok {
		return obj, nil
	}
	obj, err = s.getExchangeByExchangeId(exchangeId)
	if err != nil {
		return
	}
	if obj != nil {
		return
	}

	//查找兑换id
	feedbackfeeExchangeEntity := &FeedbackfeeExchangeEntity{}
	str := uuid.NewV4().String()
	generateCode := codePrefix + strings.Replace(str, "-", "", -1)

	now := timeutils.TimeToMillisecond(time.Now())
	feedbackfeeExchangeEntity = &FeedbackfeeExchangeEntity{}
	feedbackfeeExchangeEntity.Money = money
	feedbackfeeExchangeEntity.Platform = platform
	feedbackfeeExchangeEntity.ServerId = serverId
	feedbackfeeExchangeEntity.PlayerId = playerId

	feedbackfeeExchangeEntity.ExchangeId = exchangeId
	feedbackfeeExchangeEntity.ExpiredTime = expiredTime
	feedbackfeeExchangeEntity.CreateTime = now
	feedbackfeeExchangeEntity.Code = generateCode
	feedbackfeeExchangeEntity.Status = int32(couponservertypes.ExchangeStatusInit)
	err = s.db.DB().Save(feedbackfeeExchangeEntity).Error
	if err != nil {
		return
	}

	obj = NewFeedbackfeeExchangeObject()
	obj.FromEntity(feedbackfeeExchangeEntity)
	log.WithFields(
		log.Fields{
			"platform":    platform,
			"serverId":    serverId,
			"playerId":    playerId,
			"exchangeId":  exchangeId,
			"money":       money,
			"expiredTime": expiredTime,
			"code":        feedbackfeeExchangeEntity.Code,
		}).Info("exchange:生成兑换码")
	s.addCode(obj)

	return
}

func (s *exchangeService) addCode(obj *FeedbackfeeExchangeObject) {
	s.initMap[obj.exchangeId] = obj
	s.codeMap[obj.code] = obj
}

func (s *exchangeService) removeCode(obj *FeedbackfeeExchangeObject) {
	delete(s.initMap, obj.exchangeId)
	delete(s.codeMap, obj.code)
}

//兑换
func (s *exchangeService) Expire(exchangeId int64) (obj *FeedbackfeeExchangeObject, err error) {
	//查找
	s.m.Lock()
	defer s.m.Unlock()
	obj, ok := s.initMap[exchangeId]
	if ok {
		now := timeutils.TimeToMillisecond(time.Now())
		flag := obj.Expired(now)
		if !flag {
			return nil, nil
		}
		e, err := obj.ToEntity()
		if err != nil {
			return nil, err
		}
		err = s.db.DB().Save(e).Error
		if err != nil {
			return obj, nil
		}
		s.removeCode(obj)
		return obj, err
	}
	obj, err = s.getExchangeByExchangeId(exchangeId)
	if err != nil {
		return
	}
	if obj == nil {
		return
	}

	//幂等操作
	if obj.status == couponservertypes.ExchangeStatusExpired {
		return
	}

	return nil, nil
}

//兑换
func (s *exchangeService) Exchange(code string, wxId string) (obj *FeedbackfeeExchangeObject, err error) {
	//查找
	s.m.Lock()
	defer s.m.Unlock()
	obj, ok := s.codeMap[code]
	if ok {
		str := uuid.NewV4().String()
		orderId := strings.Replace(str, "-", "", -1)
		now := timeutils.TimeToMillisecond(time.Now())
		flag := obj.Finish(wxId, orderId, now)
		if !flag {
			return nil, nil
		}
		e, err := obj.ToEntity()
		if err != nil {
			return nil, err
		}
		err = s.db.DB().Save(e).Error
		if err != nil {
			return nil, err
		}
		s.removeCode(obj)
		return obj, err
	}

	obj, err = s.getExchangeByCode(code)
	if err != nil {
		return
	}

	//幂等操作
	if obj.status == couponservertypes.ExchangeStatusFinish || obj.status == couponservertypes.ExchangeStatusNotify {
		return
	}

	return nil, nil
}

func NewExchangeService(db coredb.DBService, rs coreredis.RedisService) (s ExchangeService, err error) {
	ts := &exchangeService{
		db: db,
		rs: rs,
	}
	err = ts.init()
	if err != nil {
		return nil, err
	}
	s = ts
	return s, nil
}
