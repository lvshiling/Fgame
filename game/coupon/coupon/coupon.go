package coupon

import (
	"encoding/json"
	logintypes "fgame/fgame/account/login/types"
	couponeventtypes "fgame/fgame/game/coupon/event/types"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/httputils"
	"fmt"
	"runtime/debug"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type CouponConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

//兑换码活动
type CouponService interface {
	//兑换订单
	Exchange(pl player.Player, code string)
	GetHost() string
	GetPort() int32
}

type couponService struct {
	//读写锁
	rwm sync.RWMutex
	cfg *CouponConfig
}

//初始化
func (s *couponService) init(cfg *CouponConfig) (err error) {
	s.cfg = cfg
	return
}

const (
	exchangePath = "/api/coupon/exchange"
)

func (s *couponService) Exchange(pl player.Player, code string) {
	sdkType := pl.GetSDKType()
	deviceType := pl.GetDevicePlatformType()
	userId := pl.GetUserId()
	platformUserId := pl.GetPlatformUserId()
	playerId := pl.GetId()
	serverId := pl.GetServerId()
	playerName := pl.GetName()
	playerLevel := pl.GetLevel()
	playerVipLevel := pl.GetVip()
	platform := global.GetGame().GetPlatform()
	s.exchange(code, platform, sdkType, deviceType, serverId, platformUserId, userId, playerId, playerLevel, playerVipLevel, playerName)
}
func (s *couponService) GetHost() string {
	return s.cfg.Host
}

func (s *couponService) GetPort() int32 {
	return s.cfg.Port
}

func (s *couponService) exchangeFailed(playerId int64, code int32, msg string) {
	eventData := CreateExchangeFailedEventData(code, msg)
	gameevent.Emit(couponeventtypes.CouponEventTypeExchangeFailed, playerId, eventData)
}

func (s *couponService) exchangeFinish(playerId int64, title string, content string, attachment string) {
	eventData := CreateExchangeFinishEventData(title, content, attachment)
	gameevent.Emit(couponeventtypes.CouponEventTypeExchangeFinish, playerId, eventData)
}

func (s *couponService) exchange(code string, platform int32, sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, platformUserId string, userId int64, playerId int64, playerLevel int32, playerVipLevel int32, name string) {

	go func(host string, port int32) {
		var err error
		var title string
		var content string
		var attachment string
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				exceptionContent := string(debug.Stack())
				log.WithFields(
					log.Fields{
						"code":           code,
						"SDKType":        sdkType.String(),
						"DeviceType":     deviceType.String(),
						"ServerId":       serverId,
						"PlatformUserId": platformUserId,
						"platform":       platform,
						"UserId":         userId,
						"PlayerId":       playerId,
						"playerName":     name,
						"playerLevel":    playerLevel,
						"playerVipLevel": playerVipLevel,
						"error":          r,
						"stack":          exceptionContent,
					}).Error("coupon:兑换兑换码,错误")
				gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			}
		}()
		defer func() {
			if err != nil {
				log.WithFields(
					log.Fields{
						"code":           code,
						"SDKType":        sdkType.String(),
						"DeviceType":     deviceType.String(),
						"ServerId":       serverId,
						"PlatformUserId": platformUserId,
						"platform":       platform,
						"UserId":         userId,
						"PlayerId":       playerId,
						"playerName":     name,
						"playerLevel":    playerLevel,
						"playerVipLevel": playerVipLevel,
						"error":          err,
					}).Warn("coupon:兑换兑换码,失败")
				s.exchangeFailed(playerId, 0, "")
				return
			}
			s.exchangeFinish(playerId, title, content, attachment)
		}()

		type exchangeForm struct {
			Code           string `json:"code"`
			SDKType        int32  `json:"sdkType"`
			DeviceType     int32  `json:"deviceType"`
			ServerId       int32  `json:"serverId"`
			Platform       int32  `json:"platform"`
			PlatformUserId string `json:"PlatformUserId"`
			UserId         int64  `json:"userId"`
			PlayerId       int64  `json:"playerId"`
			Name           string `json:"name"`
			PlayerLevel    int32  `json:"playerLevel"`
			PlayerVipLevel int32  `json:"playerVipLevel"`
		}

		getExchangePath := fmt.Sprintf("http://%s:%d%s", host, port, exchangePath)
		form := &exchangeForm{
			Code:           code,
			SDKType:        int32(sdkType),
			DeviceType:     int32(deviceType),
			ServerId:       serverId,
			PlatformUserId: platformUserId,
			Platform:       platform,
			UserId:         userId,
			PlayerId:       playerId,
			PlayerLevel:    playerLevel,
			PlayerVipLevel: playerVipLevel,
			Name:           name,
		}

		result, err := httputils.PostJsonWithRawMessage(getExchangePath, nil, form)
		if err != nil {
			return
		}
		if result.ErrorCode != 0 {
			s.exchangeFailed(playerId, int32(result.ErrorCode), result.ErrorMsg)
			return
		}

		type getExchangeResponse struct {
			Title      string `json:"title"`
			Content    string `json:"content"`
			Attachment string `json:"attachment"`
		}
		res := &getExchangeResponse{}
		err = json.Unmarshal(result.Result, res)
		if err != nil {
			return
		}
		title = res.Title
		content = res.Content
		attachment = res.Attachment

	}(s.cfg.Host, s.cfg.Port)
}

var (
	once sync.Once
	s    *couponService
)

func Init(cfg *CouponConfig) (err error) {
	once.Do(func() {
		s = &couponService{}
		err = s.init(cfg)
	})
	return err
}

func GetCouponService() CouponService {
	return s
}
