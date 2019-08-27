package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/foe/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	friendtemplate "fgame/fgame/game/friend/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOE_FEEDBACK_BUY_PROTECT_TYPE), dispatch.HandlerFunc(handleFoeFeedbackProtect))
}

//处理反馈保护购买
func handleFoeFeedbackProtect(s session.Session, msg interface{}) error {
	log.Debug("foe:处理仇人反馈保护购买")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err := foeBuyProtect(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("foe:处理仇人反馈保护购买,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("foe:处理仇人反馈保护购买,完成")
	return nil

}

//处理仇人反馈保护购买
func foeBuyProtect(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	if manager.IsOnProtected() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("foe:购买仇人反馈保护失败，保护未过期")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
	needGold := int64(noticeConstantTemp.BaoHuFei)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	if !propertyManager.HasEnoughGold(needGold, true) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("foe:购买仇人反馈保护失败，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	useReason := commonlog.GoldLogReasonFoeFeedbackBuyProtect
	flag := propertyManager.CostGold(needGold, true, useReason, useReason.String())
	if !flag {
		panic(fmt.Errorf("foe:购买仇人反馈保护消耗元宝应该成功"))
	}

	manager.BuyFoeFeedbackProtect()
	propertylogic.SnapChangedProperty(pl)

	expireTime := manager.GetFoeFeedbackProtectExpireTime()
	scMsg := pbutil.BuildSCFoeFeedbackBuyProtect(expireTime)
	pl.SendMsg(scMsg)
	return
}
