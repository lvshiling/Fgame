package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/bagua/bagua"
	"fgame/fgame/game/bagua/pbutil"
	playerbagua "fgame/fgame/game/bagua/player"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BAGUA_PAIR_TYPE), dispatch.HandlerFunc(handleBaGuaPair))
}

//处理夫妻助战邀请信息
func handleBaGuaPair(s session.Session, msg interface{}) (err error) {
	log.Debug("bagua:处理夫妻助战邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = baGuaPair(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bagua:处理夫妻助战邀请消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bagua:处理夫妻助战邀请消息完成")
	return nil

}

//夫妻助战邀请信息的逻辑
func baGuaPair(pl player.Player) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	flag := manager.IfFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:八卦秘境已达最高层")
		playerlogic.SendSystemMessage(pl, lang.BaGuaReachLimit)
		return
	}

	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()
	if marryInfo.SpouseId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:当前没有配偶,结婚后可邀请配偶帮忙闯关八卦秘境")
		playerlogic.SendSystemMessage(pl, lang.BaGuaPairNoSpouse)
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	if spl == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:您的配偶当前不在线,无法邀请助战")
		playerlogic.SendSystemMessage(pl, lang.BaGuaPairSpouseNoOnline)
		return
	}

	if marryInfo.Status != marrytypes.MarryStatusTypeMarried {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:举办婚礼后可共同闯关")
		playerlogic.SendSystemMessage(pl, lang.BaGuaNoHoldWed)
		return
	}

	ss := spl.GetScene()
	if ss == nil || ss.MapTemplate().IsFuBen() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("bagua:您的配偶当前在其他副本场景,无法邀请")
		playerlogic.SendSystemMessage(pl, lang.BaGuaSpouseInOtherFuBen)
		return
	}

	flag = manager.InviteFrequent()
	inviteTime := int64(0)
	if flag {
		leftTimeMs := manager.GetInviteLeftTime()
		leftTime := timeutils.MillisecondToSecondCeil(leftTimeMs)
		leftTimeStr := fmt.Sprintf("%d", leftTime)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("bagua:邀请配偶,cd")
		playerlogic.SendSystemMessage(pl, lang.BaGuaInviteCd, leftTimeStr)
		return
	} else {
		inviteTime = manager.InviteTime()
	}

	curlevel := manager.GetLevel()
	nextLevel := curlevel + 1
	bagua.GetBaGuaService().PairInvite(pl.GetId(), pl.GetName(), spl.GetId(), spl.GetName(), nextLevel)
	scBaGuaPairPushSpouse := pbutil.BuildSCBaGuaPairPushSpouse(pl.GetId(), nextLevel)
	spl.SendMsg(scBaGuaPairPushSpouse)
	scRealmPair := pbutil.BuildSCBaGuaPair(inviteTime)
	pl.SendMsg(scRealmPair)
	return
}
