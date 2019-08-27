package api

import (
	"net/http"

	"fgame/fgame/coupon_server/coupon"
	fgamehttpputils "fgame/fgame/pkg/httputils"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type exchangeCouponRequest struct {
	Code           string `form:"code" json:"code"`
	SDKType        int32  `form:"sdkType" json:"sdkType"`
	DeviceType     int32  `form:"deviceType" json:"deviceType"`
	ServerId       int32  `form:"serverId" json:"serverId"`
	Platform       int32  `form:"platform" json:"platform"`
	PlatformUserId string `form:"platformUserId" json:"platformUserId"`
	UserId         int64  `form:"userId" json:"userId"`
	PlayerId       int64  `form:"playerId" json:"playerId"`
	Name           string `form:"name" json:"name"`
	PlayerLevel    int32  `form:"playerLevel" json:"playerLevel"`
	PlayerVipLevel int32  `form:"playerVipLevel" json:"playerVipLevel"`
}

type exchangeCouponResponse struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

func handleExchangeCoupon(rw http.ResponseWriter, req *http.Request) {

	form := &exchangeCouponRequest{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(
			log.Fields{
				"ip":    req.RemoteAddr,
				"error": err,
			}).Error("coupon:兑换兑换码,解析请求错误")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	code := form.Code
	sdkType := form.SDKType
	deviceType := form.DeviceType

	serverId := form.ServerId
	userId := form.UserId
	platformUserId := form.PlatformUserId
	playerId := form.PlayerId
	playerLevel := form.PlayerLevel
	playerName := form.Name
	platform := form.Platform
	playerVipLevel := form.PlayerVipLevel
	log.WithFields(
		log.Fields{
			"ip":             req.RemoteAddr,
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
		}).Info("coupon:兑换兑换码")
	ctx := req.Context()
	couponService := coupon.CouponServiceInContext(ctx)
	title, content, attachment, err := couponService.Exchange(code, platform, sdkType, platformUserId, deviceType, serverId, userId, playerId, playerLevel, playerVipLevel, playerName)
	if err != nil {
		codeErr, ok := err.(coupon.CouponError)
		if !ok {
			log.WithFields(
				log.Fields{
					"ip":             req.RemoteAddr,
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
					"error":          err,
				}).Error("coupon:兑换兑换码,兑换失败")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := fgamehttpputils.NewFailedResultWithMsg(int(codeErr.Code()), codeErr.Code().String())
		httputils.WriteJSON(rw, http.StatusOK, res)
		return
	}

	result := &exchangeCouponResponse{
		Title:      title,
		Content:    content,
		Attachment: attachment,
	}
	res := fgamehttpputils.NewSuccessResult(result)
	httputils.WriteJSON(rw, http.StatusOK, res)

	log.WithFields(
		log.Fields{
			"ip":             req.RemoteAddr,
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
		}).Info("coupon:兑换兑换码")
}
