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
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_HANDLE_TOKEN_SUO_YAO_TYPE), dispatch.HandlerFunc(handleJieYiTokenSuoYaoResp))
}

// 同意给索要人信物
func handleJieYiTokenSuoYaoResp(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理处理信物索要请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiHandleTokenSuoYao)
	suoYaoRenId := csMsg.GetPlayerId()
	token := jieyitypes.JieYiTokenType(csMsg.GetToken())
	if !token.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"token":    int32(token),
			}).Warn("jieyi: 信物类型不合法")
		return
	}

	err = jieYiTokenSuoYaoResp(tpl, suoYaoRenId, token)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理处理信物索要请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理处理信物索要请求消息,成功")
	return
}

func jieYiTokenSuoYaoResp(pl player.Player, suoYaoRenId int64, token jieyitypes.JieYiTokenType) (err error) {
	suoYaoPl := player.GetOnlinePlayerManager().GetPlayerById(suoYaoRenId)
	if suoYaoPl == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"suoYaoRenId": suoYaoRenId,
			}).Warn("jieyi: 对方不在线")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotOnline)
		return
	}

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	suoYaoRenObj := jieyi.GetJieYiService().GetJieYiMemberInfo(suoYaoRenId)
	if plObj == nil || suoYaoRenObj == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"suoYaoRenId": suoYaoRenId,
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	if plObj.GetJieYiId() != suoYaoRenObj.GetJieYiId() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"suoYaoRenId": suoYaoRenId,
			}).Warn("jieyi: 不在同一结义阵营")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotIsSameJieYi)
		return
	}

	lastToken := suoYaoRenObj.GetTokenType()
	if int(token) <= int(lastToken) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"tokenType": int32(token),
			}).Warn("jieyi: 信物级别不够")
		playerlogic.SendSystemMessage(pl, lang.JieYiTokenNotChange)
		return
	}

	// 扣除信物
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

	// 判断背包里信物是否足够
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	enoughFlag := true
	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			enoughFlag = false
		}
	}

	if enoughFlag {
		flag := jieyi.GetJieYiService().TokenChangeSucess(suoYaoRenId, token)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"suoYaoRenId":   suoYaoRenId,
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 替换失败")
			return
		}
		// 足够则消耗物品
		reason := commonlog.InventoryLogReasonJieYiTokenTypeChangeUse
		reasonText := fmt.Sprintf(reason.String(), token.String())
		flag = inventoryManager.BatchRemove(itemMap, reason, reasonText)
		if !flag {
			panic("jieyi: 消耗物品应该成功")
		}
		// 物品改变推送
		inventorylogic.SnapInventoryChanged(pl)
	} else {
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

		flag := jieyi.GetJieYiService().TokenChangeSucess(suoYaoRenId, token)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"suoYaoRenId":   suoYaoRenId,
				"tokenType":     int32(token),
				"lastTokenType": int32(lastToken),
			}).Warn("jieyi: 替换失败")
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
	}

	jieYi := suoYaoRenObj.GetJieYi()

	scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(suoYaoRenObj)
	jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)

	jieyilogic.JieYiMemberChanged(jieYi)

	scMsg := pbutil.BuildSCJieYiHandleTokenSuoYao(suoYaoRenId, int32(token))
	pl.SendMsg(scMsg)

	return
}
