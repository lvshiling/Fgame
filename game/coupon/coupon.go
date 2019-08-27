package coupon

import (
	"fgame/fgame/game/coupon/coupon"

	//注册管理器
	_ "fgame/fgame/game/coupon/event/listener"
)

func Init(cfg *coupon.CouponConfig) (err error) {

	err = coupon.Init(cfg)
	if err != nil {
		return
	}
	return
}
