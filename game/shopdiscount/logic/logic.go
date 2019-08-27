package logic

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shoptype "fgame/fgame/game/shop/types"
	playershopdiscount "fgame/fgame/game/shopdiscount/player"
	shopdiscounttemplate "fgame/fgame/game/shopdiscount/template"
)

const (
	defaultRatio = float64(common.MAX_RATE)
)

// 获取商店折扣
func GetShopDiscount(pl player.Player, shopType shoptype.ShopConsumeType) float64 {

	discountManager := pl.GetPlayerDataManager(playertypes.PlayerShopDiscountDataManagerType).(*playershopdiscount.PlayerShopDiscountDataManager)
	discountType := discountManager.GetCurShopDiscountType()
	discountTemp := shopdiscounttemplate.GetShopDiscountTemplateService().GetShopDiscountTemplateByType(discountType, shopType)
	if discountTemp == nil {
		return defaultRatio
	}

	return float64(discountTemp.Discount)
}
