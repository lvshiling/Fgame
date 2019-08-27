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
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	shoptypes "fgame/fgame/game/shop/types"
	"fmt"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_GIVE_TYPE), dispatch.HandlerFunc(handleJieYiTokenGive))
}

func handleJieYiTokenGive(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理兄弟信物赠送请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiTokenGive)
	receiverId := csMsg.GetPlayerId()
	leaveWord := csMsg.GetLeaveWord()
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
			}).Warn("jieyi: 信物赠送方式不合法")
		return
	}

	err = jieYiTokenGive(tpl, receiverId, token, leaveWord, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理兄弟信物赠送请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理兄弟信物赠送请求消息,成功")

	return

}

func jieYiTokenGive(pl player.Player, receiverId int64, token jieyitypes.JieYiTokenType, leaveWord string, typ jieyitypes.JieYiTokenChangeMethod) (err error) {
	receiverPl := player.GetOnlinePlayerManager().GetPlayerById(receiverId)
	if receiverPl == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 对方不在线")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotOnline)
		return
	}

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if plObj == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}
	receiverObj := jieyi.GetJieYiService().GetJieYiMemberInfo(receiverId)
	if receiverObj == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 对方未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	if plObj.GetJieYiId() != receiverObj.GetJieYiId() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"receiverId": receiverId,
			}).Warn("jieyi: 不在同一结义阵营")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotIsSameJieYi)
		return
	}

	lastToken := receiverObj.GetTokenType()
	if int(token) <= int(lastToken) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"tokenType": int32(token),
			}).Warn("jieyi: 信物级别不够")
		playerlogic.SendSystemMessage(pl, lang.JieYiTokenNotChange)
		return
	}

	lastTokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenTemplate(lastToken)
	tokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenTemplate(token)
	if tokenTemp == nil {
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
	lastItemMap := make(map[int32]int32)
	if lastTokenTemp != nil {
		lastItemMap = lastTokenTemp.GetNeedItemMap()
	}
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

	// 判断是否补差价升级
	if typ == jieyitypes.JieYiTokenChangeMethodBuChaJia {
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
		flag := jieyi.GetJieYiService().TokenChangeSucess(receiverId, token)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"receiverId":    receiverId,
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 无法赠送")
			return
		}

		//更新自动购买每日限购次数
		if len(shopIdMap) != 0 {
			shoplogic.ShopDayCountChanged(pl, shopIdMap)
		}

		//消耗钱
		goldUseReason := commonlog.GoldLogReasonJieYiTokenChange
		goldReason := fmt.Sprintf(goldUseReason.String(), lastToken.String(), token.String(), jieyitypes.JieYiItemUseTypeGiveXiongDi.String())
		silverUseReason := commonlog.SilverLogReasonJieYiTokenChange
		silverReason := fmt.Sprintf(silverUseReason.String(), lastToken.String(), token.String(), jieyitypes.JieYiItemUseTypeGiveXiongDi.String())
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
		flag := jieyi.GetJieYiService().TokenChangeSucess(receiverId, token)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"receiverId":    receiverId,
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 无法赠送")
			return
		}
		reason := commonlog.InventoryLogReasonJieYiTokenTypeChangeUse
		reasonText := fmt.Sprintf(reason.String(), token.String())
		flag = inventoryManager.BatchRemove(itemMap, reason, reasonText)
		if !flag {
			panic("jieyi: 消耗物品应该成功")
		}
		// 物品改变推送
		inventorylogic.SnapInventoryChanged(pl)
	}

	jieYi := receiverObj.GetJieYi()
	scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(receiverObj)
	jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)
	jieyilogic.JieYiMemberChanged(jieYi)

	if receiverPl != nil {
		scJieYiTokenGiveNotice := pbutil.BuildSCJieYiTokenGiveNotice(pl.GetId(), pl.GetName(), leaveWord, int32(token), int32(lastToken))
		receiverPl.SendMsg(scJieYiTokenGiveNotice)
	}

	scJieYiTokenGive := pbutil.BuildSCJieYiTokenGive(int32(token), receiverId, receiverObj.GetTokenLev())
	pl.SendMsg(scJieYiTokenGive)
	return
}
