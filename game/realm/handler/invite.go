package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/realm/pbutil"
	playerrealm "fgame/fgame/game/realm/player"
	"fgame/fgame/game/realm/realm"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_REALM_PAIR_TYPE), dispatch.HandlerFunc(handleRealmPair))
}

//处理夫妻助战邀请信息
func handleRealmPair(s session.Session, msg interface{}) (err error) {
	log.Debug("realm:处理夫妻助战邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = realmPair(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("realm:处理夫妻助战邀请消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("realm:处理夫妻助战邀请消息完成")
	return nil

}

//夫妻助战邀请信息的逻辑
func realmPair(pl player.Player) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	flag := manager.IfFullLevel()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:境界已达最高层")
		playerlogic.SendSystemMessage(pl, lang.RealmReachLimit)
		return
	}

	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()
	if marryInfo.SpouseId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:当前没有配偶,结婚后可邀请配偶帮忙闯关天劫塔")
		playerlogic.SendSystemMessage(pl, lang.RealmPairNoSpouse)
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	if spl == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:您的配偶当前不在线,无法邀请助战")
		playerlogic.SendSystemMessage(pl, lang.RealmPairSpouseNoOnline)
		return
	}

	if marryInfo.Status != marrytypes.MarryStatusTypeMarried {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:举办婚礼后可共同闯关")
		playerlogic.SendSystemMessage(pl, lang.RealmNoHoldWed)
		return
	}

	ss := spl.GetScene()
	if ss == nil || ss.MapTemplate().IsFuBen() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("realm:您的配偶当前在其他副本场景,无法邀请")
		playerlogic.SendSystemMessage(pl, lang.RealmSpouseInOtherFuBen)
		return
	}

	flag = manager.InviteFrequent()
	inviteTime := int64(0)
	if flag {
		playerlogic.SendSystemMessage(pl, lang.CommonOperFrequent)
		return
	} else {
		inviteTime = manager.InviteTime()
	}

	curlevel := manager.GetTianJieTaLevel()
	nextLevel := curlevel + 1
	realm.GetRealmRankService().PairInvite(pl.GetId(), pl.GetName(), spl.GetId(), spl.GetName(), nextLevel)
	scRealmPairPushSpouse := pbutil.BuildSCRealmPairPushSpouse(pl.GetId(), nextLevel)
	spl.SendMsg(scRealmPairPushSpouse)
	scRealmPair := pbutil.BuildSCRealmPair(inviteTime)
	pl.SendMsg(scRealmPair)
	return
}
