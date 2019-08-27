package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/major/major"
	"fgame/fgame/game/major/pbutil"
	playermajor "fgame/fgame/game/major/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MAJOR_INVITE_DEAL_TYPE), dispatch.HandlerFunc(handleMajorInviteDeal))
}

//处理双修决策信息
func handleMajorInviteDeal(s session.Session, msg interface{}) (err error) {
	log.Debug("major:处理双修决策消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMajorInviteDeal)
	result := csMsg.GetResult()

	err = majorInviteDeal(tpl, result)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"result":   result,
				"error":    err,
			}).Error("major:处理双修决策消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("major:处理双修决策消息完成")
	return nil
}

//处理双修决策信息逻辑
func majorInviteDeal(pl player.Player, result bool) (err error) {
	majorInvite := major.GetMajorService().GetMajorInvite(pl.GetId())
	if majorInvite == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("major:配偶邀请已经取消")
		playerlogic.SendSystemMessage(pl, lang.MajorHadCancle)
		return
	}

	if result {
		//判断次数是否足够
		manager := pl.GetPlayerDataManager(types.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
		flag := manager.HasMajorNum(majorInvite.FuBenType)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("major:今日双修次数已用完")
			playerlogic.SendSystemMessage(pl, lang.MajorInviteNoTimes)
			return
		}
		if !playerlogic.CheckCanEnterScene(pl) {
			return
		}

		peerPlayerId := majorInvite.PlayerId
		teamDataObj := team.GetTeamService().GetTeamByPlayerId(peerPlayerId)
		if teamDataObj != nil && teamDataObj.IsMatch() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("major:对方当前正在3v3匹配,邀请取消")
			playerlogic.SendSystemMessage(pl, lang.MajorPeerIn3v3Match)
			return
		}
	}

	_, codeResult := major.GetMajorService().MajorInviteDeal(pl, result)
	scMajorInviteDeal := pbutil.BuildSCMajorInviteDeal(int32(codeResult))
	pl.SendMsg(scMajorInviteDeal)
	return
}
