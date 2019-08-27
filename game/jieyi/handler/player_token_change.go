package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	shoptypes "fgame/fgame/game/shop/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_CHANGE_TYPE), dispatch.HandlerFunc(handlePlayerTokenChange))
}

func handlePlayerTokenChange(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家替换信物请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiTokenChange)
	token := jieyitypes.JieYiTokenType(csMsg.GetToken())
	if !token.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"token":    int32(token),
			}).Warn("jieyi: 信物类型不合法")
		return
	}
	typ := jieyitypes.JieYiTokenChangeMethod(csMsg.GetType())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"typ":      int32(typ),
			}).Warn("jieyi: 玩家信物替换方式不合法")
		return
	}

	err = playerTokenChange(tpl, token, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家替换信物请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理玩家替换信物请求消息,成功")

	return
}

func playerTokenChange(pl player.Player, token jieyitypes.JieYiTokenType, typ jieyitypes.JieYiTokenChangeMethod) (err error) {
	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if plObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}
	lastToken := plObj.GetTokenType()
	if lastToken == jieyitypes.JieYiTokenTypeInvalid {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 信物未激活")
		playerlogic.SendSystemMessage(pl, lang.JieYiTokenNotActivite)
		return
	}
	//判断信物
	if plObj.GetTokenType() >= token {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"originToken": plObj.GetTokenType().String(),
				"tokenType":   token.String(),
			}).Warn("jieyi: 信物级别不够")
		playerlogic.SendSystemMessage(pl, lang.JieYiTokenAlreadyActivite)
		return
	}

	lastTokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenTemplate(lastToken)
	tokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenTemplate(token)
	if tokenTemp == nil || lastTokenTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		return
	}
	itemMap := tokenTemp.GetNeedItemMap()
	lastItemMap := lastTokenTemp.GetNeedItemMap()

	// 判断是否补差价升级
	if typ == jieyitypes.JieYiTokenChangeMethodBuChaJia {
		lastTokenSliver := int64(0)
		lastTokenBind := int64(0)
		lastTokenGold := int64(0)
		for id, num := range lastItemMap {
			shopTempMap := shop.GetShopService().GetShopItemMap(id)
			goldTempList, ok := shopTempMap[shoptypes.ShopConsumeTypeGold]
			if !ok {
				continue
			}
			lastTokenGold += int64(goldTempList[0].ConsumeData1) * int64(num)
		}
		isEnoughBuyTimes, shopIdMap := shoplogic.MaxBuyTimesForPlayerMapComplementGold(pl, lastTokenGold, itemMap)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
			}).Warn("jieyi:购买物品失败,补差价")
			playerlogic.SendSystemMessage(pl, lang.JieYiBuChaJiaFail)
			return
		}

		tokenBind, tokenGold, tokenSliver := shoplogic.ShopCostData(pl, shopIdMap)

		needSliver := (tokenSliver - lastTokenSliver)
		needBind := (tokenBind - lastTokenBind)
		needGold := (tokenGold - lastTokenGold)

		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if needSliver != 0 {
			flag := propertyManager.HasEnoughSilver(int64(needSliver))
			if !flag {
				log.WithFields(log.Fields{
					"playerId":      pl.GetId(),
					"tokenType":     int32(token),
					"lastTokenType": int32(lastToken),
				}).Warn("jieyi: 银两不足,无法替换")
				playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
				return
			}
		}

		if needGold != 0 {
			flag := propertyManager.HasEnoughGold(int64(needGold), false)
			if !flag {
				log.WithFields(log.Fields{
					"playerId":      pl.GetId(),
					"tokenType":     int32(token),
					"lastTokenType": int32(lastToken),
				}).Warn("jieyi: 元宝不足,无法替换")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}
		}

		costBind := needBind + needGold
		if costBind != 0 {
			flag := propertyManager.HasEnoughGold(int64(costBind), true)
			if !flag {
				log.WithFields(log.Fields{
					"playerId":      pl.GetId(),
					"tokenType":     int32(token),
					"lastTokenType": int32(lastToken),
				}).Warn("jieyi: 元宝不足,无法替换")
				playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
				return
			}
		}
		flag := jieyi.GetJieYiService().TokenChangeSucess(pl.GetId(), token)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 无法失败")
			return
		}

		//更新自动购买每日限购次数
		if len(shopIdMap) != 0 {
			shoplogic.ShopDayCountChanged(pl, shopIdMap)
		}

		//消耗钱
		goldUseReason := commonlog.GoldLogReasonJieYiTokenChange
		goldReason := fmt.Sprintf(goldUseReason.String(), int(lastToken), int(token))
		silverUseReason := commonlog.SilverLogReasonJieYiTokenChange
		silverReason := fmt.Sprintf(silverUseReason.String(), int(lastToken), int(token))
		flag = propertyManager.Cost(int64(needBind), int64(needGold), goldUseReason, goldReason, int64(needSliver), silverUseReason, silverReason)
		if !flag {
			panic(fmt.Errorf("jieyi: 替换结义信物消耗钱应该成功"))
		}
		//同步钱
		if needBind != 0 || needGold != 0 || needSliver != 0 {
			propertylogic.SnapChangedProperty(pl)
		}
	} else {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if len(itemMap) != 0 {
			if !inventoryManager.HasEnoughItems(itemMap) {
				log.WithFields(
					log.Fields{
						"playerId":  pl.GetId(),
						"tokenType": int32(token),
					}).Warn("jieyi: 所需物品不足")
				playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
				return
			}
		}
		flag := jieyi.GetJieYiService().TokenChangeSucess(pl.GetId(), token)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 无法失败")
			return
		}
		reason := commonlog.InventoryLogReasonJieYiTokenTypeChangeUse
		reasonText := fmt.Sprintf(reason.String(), token.String(), jieyitypes.JieYiItemUseTypeTiHuan.String())
		flag = inventoryManager.BatchRemove(itemMap, reason, reasonText)
		if !flag {
			panic("jieyi: 消耗物品应该成功")
		}
		// 物品改变推送
		inventorylogic.SnapInventoryChanged(pl)
	}

	jieYi := plObj.GetJieYi()

	scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(plObj)
	jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)

	jieyilogic.JieYiMemberChanged(jieYi)

	scMsg := pbutil.BuildSCJieYiTokenChange(int32(token), plObj.GetTokenLev())
	pl.SendMsg(scMsg)

	return
}
