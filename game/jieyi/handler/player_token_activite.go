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
	"fmt"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_ACTIVITE_TYPE), dispatch.HandlerFunc(handlePlayerTokenActivite))
}

func handlePlayerTokenActivite(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家激活信物请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiTokenActivite)
	token := jieyitypes.JieYiTokenType(csMsg.GetToken())
	if !token.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"token":    int32(token),
			}).Warn("jieyi: 信物类型不合法")
		return
	}

	err = playerTokenActivite(tpl, token)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家激活信物请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理玩家激活信物请求消息,成功")

	return
}

func playerTokenActivite(pl player.Player, token jieyitypes.JieYiTokenType) (err error) {

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if plObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	//判断信物
	if plObj.GetTokenType() > jieyitypes.JieYiTokenTypeInvalid {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"tokenType": int32(token),
			}).Warn("jieyi: 信物已经激活")
		playerlogic.SendSystemMessage(pl, lang.JieYiTokenAlreadyActivite)
		return
	}

	tokenTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiTokenTemplate(token)
	if tokenTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"tokenType": int32(token),
			}).Warn("jieyi: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		return
	}
	itemMap := tokenTemp.GetNeedItemMap()

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
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"tokenType": int32(token),
			}).Warn("jieyi: 信物激活失败")
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

	jieYi := plObj.GetJieYi()

	scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(plObj)
	jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)

	jieyilogic.JieYiMemberChanged(jieYi)

	scMsg := pbutil.BuildSCJieYiTokenActivite(int32(token))
	pl.SendMsg(scMsg)

	return
}
