package coupon

import (
	"context"
	coredb "fgame/fgame/core/db"
	coreredis "fgame/fgame/core/redis"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
)

const (
	ErrorCodeCouponInvalid ErrorCode = ErrorCodeCoupon + 1 + iota
	ErrorCodeCouponExpired
	ErrorCodeCouponTimesLimit
	ErrorCodeCouponLevelLimit
	ErrorCodeCouponVipLimit
)

var (
	errorCodeCouponMap = map[ErrorCode]string{
		ErrorCodeCouponInvalid:    "兑换码无效",
		ErrorCodeCouponExpired:    "兑换码已过期",
		ErrorCodeCouponTimesLimit: "兑换码次数已达上限",
		ErrorCodeCouponLevelLimit: "兑换码等级不够",
		ErrorCodeCouponVipLimit:   "兑换码vip等级不够",
	}
)

func init() {
	MergeErrorCodeMap(errorCodeCouponMap)
}

func SetupCouponServiceHandler(s CouponService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithCouponService(ctx, s)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}

type CouponService interface {
	Exchange(code string, platform int32, sdkType int32, platformUserId string, deviceType int32, serverId int32, userId int64, playerId int64, playerLevel int32, playerVipLevel int32, playerName string) (title, content, attachment string, err error)
}

type contextKey string

const (
	couponServiceKey contextKey = "fgame.coupon"
)

func WithCouponService(parent context.Context, s CouponService) context.Context {
	ctx := context.WithValue(parent, couponServiceKey, s)
	return ctx
}

func CouponServiceInContext(parent context.Context) CouponService {
	s := parent.Value(couponServiceKey)
	if s == nil {
		return nil
	}
	ts, ok := s.(CouponService)
	if !ok {
		return nil
	}
	return ts
}

type couponService struct {
	m  sync.Mutex
	db coredb.DBService
	rs coreredis.RedisService
}

var (
	findCodeSql = `
		select 
		rc.id as id,
		rc.redeemCode as redeemCode,
		rc.redeemId as redeemId,
		rc.useNum as useNum,
		rc.createTime as createTime,
		rc.deleteTime as deleteTime,
		rc.updateTime as updateTime,
		re.giftBagName as giftBagName,
		re.giftBagDesc as giftBagDesc,
		re.giftBagContent as giftBagContent,
		re.redeemPlayerUseNum as redeemPlayerUseNum,
		re.redeemServerUseNum as redeemServerUseNum,
		re.redeemUseNum as redeemUseNum,
		re.sdkTypes as sdkTypes,
		re.sendType as sendType,
		re.endTime as endTime,
		re.startTime as startTime,
		re.minPlayerLevel as minPlayerLevel,
		re.minVipLevel as minVipLevel
		from t_redeem_code rc left join
		t_redeem re on rc.redeemId=re.id 
		where rc.deleteTime=0 and rc.redeemCode=? limit 1
	`
)

//兑换
func (s *couponService) Exchange(code string, platform int32, sdkType int32, platformUserId string, deviceType int32, serverId int32, userId int64, playerId int64, playerLevel int32, playerVipLevel int32, playerName string) (title string, content string, attachment string, err error) {
	//查找
	s.m.Lock()
	defer s.m.Unlock()
	redeemCodeEntity := &RedeemCodeEntity{}
	err = s.db.DB().First(redeemCodeEntity, "redeemCode=?", code).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//兑换码不存在
			err = ErrorCodeCouponInvalid
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
				}).Warn("coupon:兑换码不存在")
			return
		}
		return
	}
	now := timeutils.TimeToMillisecond(time.Now())
	redeemUseNumEntity := &RedeemUseNumEntity{}
	err = s.db.DB().First(redeemUseNumEntity, "redeemId=?", redeemCodeEntity.RedeemId).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
		redeemUseNumEntity = &RedeemUseNumEntity{
			RedeemId:   redeemCodeEntity.RedeemId,
			CreateTime: now,
		}
	}

	redeemCompleteCodeEntity := &RedeemCompleteCodeEntity{}
	//查找兑换码
	err = s.db.DB().Raw(findCodeSql, code).Scan(redeemCompleteCodeEntity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//兑换码不存在
			err = ErrorCodeCouponInvalid
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
				}).Warn("coupon:兑换码不存在")
			return
		}
		return
	}
	sdkTypeStr := fmt.Sprintf("%d", sdkType)
	sdkTypeList := strings.Split(redeemCompleteCodeEntity.SdkTypes, ",")
	isContainSdk := false
	if len(sdkTypeList) != 0 {
		for _, tempSdkType := range sdkTypeList {
			if tempSdkType == sdkTypeStr {
				isContainSdk = true
				break
			}
		}
	} else {
		isContainSdk = true
	}
	//验证渠道
	if !isContainSdk {
		log.WithFields(
			log.Fields{
				"code":           code,
				"platform":       platform,
				"sdkType":        sdkType,
				"platformUserId": platformUserId,
				"deviceType":     deviceType,
				"serverId":       serverId,
				"userId":         userId,
				"playerId":       playerId,
				"playerLevel":    playerLevel,
				"playerName":     playerName,
				"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
			}).Warn("coupon:兑换码sdk不一样")
		err = ErrorCodeCouponInvalid
		return
	}

	//验证是否失效
	if now < redeemCompleteCodeEntity.StartTime {
		log.WithFields(
			log.Fields{
				"code":           code,
				"platform":       platform,
				"sdkType":        sdkType,
				"platformUserId": platformUserId,
				"deviceType":     deviceType,
				"serverId":       serverId,
				"userId":         userId,
				"playerId":       playerId,
				"playerLevel":    playerLevel,
				"playerName":     playerName,
				"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
				"startTime":      redeemCompleteCodeEntity.StartTime,
			}).Warn("coupon:兑换码,还没开始")
		err = ErrorCodeCouponInvalid
		return
	}

	if now > redeemCompleteCodeEntity.EndTime {
		log.WithFields(
			log.Fields{
				"code":           code,
				"platform":       platform,
				"sdkType":        sdkType,
				"platformUserId": platformUserId,
				"deviceType":     deviceType,
				"serverId":       serverId,
				"userId":         userId,
				"playerId":       playerId,
				"playerLevel":    playerLevel,
				"playerName":     playerName,
				"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
				"startTime":      redeemCompleteCodeEntity.StartTime,
				"endTime":        redeemCompleteCodeEntity.EndTime,
			}).Warn("coupon:兑换码,已过期")
		err = ErrorCodeCouponExpired
		return
	}

	if redeemCompleteCodeEntity.ReedmUseNum != 0 {
		redeemRecordList := make([]*RedeemRecordEntity, 0, 1)
		//获取兑换记录
		err = s.db.DB().Find(&redeemRecordList, "redeemCode=?", code).Error
		useTimes := int32(0)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return
			}
			err = nil
		}
		useTimes = int32(len(redeemRecordList))
		if useTimes >= redeemCompleteCodeEntity.ReedmUseNum {
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
					"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
					"startTime":      redeemCompleteCodeEntity.StartTime,
					"endTime":        redeemCompleteCodeEntity.EndTime,
					"useNum":         redeemCompleteCodeEntity.UseNum,
					"useTimes":       useTimes,
					"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				}).Warn("coupon:兑换码,兑换码次数已达上限")
			err = ErrorCodeCouponTimesLimit
			return
		}

	}

	//验证是否已达全服上限
	if redeemCompleteCodeEntity.RedeemServerUseNum != 0 {
		if redeemCompleteCodeEntity.UseNum >= redeemCompleteCodeEntity.RedeemServerUseNum {
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
					"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
					"startTime":      redeemCompleteCodeEntity.StartTime,
					"endTime":        redeemCompleteCodeEntity.EndTime,
					"useNum":         redeemCompleteCodeEntity.UseNum,
					"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				}).Warn("coupon:兑换码,全服次数已达上限")
			err = ErrorCodeCouponTimesLimit
			return
		}

		if redeemUseNumEntity.UseNum >= redeemCompleteCodeEntity.RedeemServerUseNum {
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
					"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
					"startTime":      redeemCompleteCodeEntity.StartTime,
					"endTime":        redeemCompleteCodeEntity.EndTime,
					"useNum":         redeemUseNumEntity.UseNum,
					"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				}).Warn("coupon:兑换礼包,全服次数已达上限")
			err = ErrorCodeCouponTimesLimit
			return
		}
	}
	//验证是否已达次数上限
	if redeemCompleteCodeEntity.RedeemPlayerUseNum != 0 {
		redeemRecordList := make([]*RedeemRecordEntity, 0, 1)
		//获取兑换记录
		err = s.db.DB().Find(&redeemRecordList, "redeemCode=? and playerId=?", code, playerId).Error
		useTimes := int32(0)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return
			}
			err = nil
		}
		useTimes = int32(len(redeemRecordList))

		if useTimes >= redeemCompleteCodeEntity.RedeemPlayerUseNum {
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
					"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
					"startTime":      redeemCompleteCodeEntity.StartTime,
					"endTime":        redeemCompleteCodeEntity.EndTime,
					"useNum":         redeemCompleteCodeEntity.UseNum,
					"useTimes":       useTimes,
					"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				}).Warn("coupon:兑换码,个人次数已达上限")
			err = ErrorCodeCouponTimesLimit
			return
		}

		redeemIdRecordList := make([]*RedeemRecordEntity, 0, 1)
		//获取兑换记录
		err = s.db.DB().Find(&redeemIdRecordList, "redeemId=? and playerId=?", redeemCodeEntity.RedeemId, playerId).Error
		reedIdUseTimes := int32(0)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return
			}
			err = nil
		}
		reedIdUseTimes = int32(len(redeemIdRecordList))

		if reedIdUseTimes >= redeemCompleteCodeEntity.RedeemPlayerUseNum {
			log.WithFields(
				log.Fields{
					"code":           code,
					"platform":       platform,
					"sdkType":        sdkType,
					"platformUserId": platformUserId,
					"deviceType":     deviceType,
					"serverId":       serverId,
					"userId":         userId,
					"playerId":       playerId,
					"playerLevel":    playerLevel,
					"playerName":     playerName,
					"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
					"startTime":      redeemCompleteCodeEntity.StartTime,
					"endTime":        redeemCompleteCodeEntity.EndTime,
					"useNum":         redeemCompleteCodeEntity.UseNum,
					"useTimes":       useTimes,
					"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				}).Warn("coupon:兑换礼包,个人次数已达上限")
			err = ErrorCodeCouponTimesLimit
			return
		}
	}
	if redeemCompleteCodeEntity.MinPlayerLevel > playerLevel {
		log.WithFields(
			log.Fields{
				"code":           code,
				"platform":       platform,
				"sdkType":        sdkType,
				"platformUserId": platformUserId,
				"deviceType":     deviceType,
				"serverId":       serverId,
				"userId":         userId,
				"playerId":       playerId,
				"playerLevel":    playerLevel,
				"playerVipLevel": playerVipLevel,
				"playerName":     playerName,
				"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
				"startTime":      redeemCompleteCodeEntity.StartTime,
				"endTime":        redeemCompleteCodeEntity.EndTime,
				"useNum":         redeemCompleteCodeEntity.UseNum,
				"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				"minVipLevel":    redeemCompleteCodeEntity.MinVipLevel,
				"minPlayerLevel": redeemCompleteCodeEntity.MinPlayerLevel,
			}).Warn("coupon:兑换码,vip等级不够")
		err = ErrorCodeCouponVipLimit
		return
	}
	if redeemCompleteCodeEntity.MinVipLevel > playerVipLevel {
		log.WithFields(
			log.Fields{
				"code":           code,
				"platform":       platform,
				"sdkType":        sdkType,
				"platformUserId": platformUserId,
				"deviceType":     deviceType,
				"serverId":       serverId,
				"userId":         userId,
				"playerId":       playerId,
				"playerLevel":    playerLevel,
				"playerVipLevel": playerVipLevel,
				"playerName":     playerName,
				"sdkTypes":       redeemCompleteCodeEntity.SdkTypes,
				"startTime":      redeemCompleteCodeEntity.StartTime,
				"endTime":        redeemCompleteCodeEntity.EndTime,
				"useNum":         redeemCompleteCodeEntity.UseNum,
				"serverUseNum":   redeemCompleteCodeEntity.RedeemServerUseNum,
				"minVipLevel":    redeemCompleteCodeEntity.MinVipLevel,
			}).Warn("coupon:兑换码,等级不够")
		err = ErrorCodeCouponLevelLimit
		return
	}
	trans := s.db.DB().Begin()
	defer func() {
		if err != nil {
			trans.Rollback()
		}
	}()
	//更新使用次数
	redeemCodeEntity.UseNum += 1
	redeemCodeEntity.UpdateTime = now
	err = trans.Save(redeemCodeEntity).Error
	if err != nil {
		return
	}
	redeemUseNumEntity.UseNum += 1
	redeemUseNumEntity.UpdateTime = now
	err = trans.Save(redeemUseNumEntity).Error
	if err != nil {
		return
	}
	//添加记录
	redeemRecordEntity := &RedeemRecordEntity{}
	redeemRecordEntity.RedeemCode = code
	redeemRecordEntity.RedeemId = redeemCompleteCodeEntity.RedeemId
	redeemRecordEntity.PlatformId = platform
	redeemRecordEntity.ServerId = serverId
	redeemRecordEntity.SdkType = sdkType
	redeemRecordEntity.PlatformUserId = platformUserId
	redeemRecordEntity.UserId = userId
	redeemRecordEntity.PlayerId = playerId
	redeemRecordEntity.PlayerLevel = playerLevel
	redeemRecordEntity.PlayerVipLevel = playerVipLevel
	redeemRecordEntity.PlayerName = playerName
	redeemRecordEntity.CreateTime = now

	err = trans.Save(redeemRecordEntity).Error
	if err != nil {
		return
	}
	title = redeemCompleteCodeEntity.GiftBagName
	content = redeemCompleteCodeEntity.GiftBagDesc
	attachment = redeemCompleteCodeEntity.GiftBagContent
	err = trans.Commit().Error
	if err != nil {
		return
	}
	return
}

func NewCouponService(db coredb.DBService, rs coreredis.RedisService) CouponService {
	s := &couponService{
		db: db,
		rs: rs,
	}
	return s
}
