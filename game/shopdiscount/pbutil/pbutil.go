package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playershopdiscount "fgame/fgame/game/shopdiscount/player"
)

func BuildSCShopDiscountGet(obj *playershopdiscount.PlayerShopDiscountObject) *uipb.SCShopDiscountGet {
	scShopDiscountGet := &uipb.SCShopDiscountGet{}
	scShopDiscountGet.Info = buildShopDiscountInfo(obj)
	return scShopDiscountGet
}

func BuildSCShopDiscountNotice(obj *playershopdiscount.PlayerShopDiscountObject) *uipb.SCShopDiscountNotice {
	scShopDiscountNotice := &uipb.SCShopDiscountNotice{}
	scShopDiscountNotice.Info = buildShopDiscountInfo(obj)
	return scShopDiscountNotice
}

func buildShopDiscountInfo(obj *playershopdiscount.PlayerShopDiscountObject) *uipb.ShopDiscountInfo {
	info := &uipb.ShopDiscountInfo{}
	typ := int32(obj.Typ)
	startTime := obj.StartTime
	endTime := obj.EndTime
	info.Typ = &typ
	info.StartTime = &startTime
	info.EndTime = &endTime
	return info
}
