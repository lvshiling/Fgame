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
	majortypes "fgame/fgame/game/major/types"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MAJOR_INVITE_TYPE), dispatch.HandlerFunc(handleMajorInvite))
}

//处理双修邀请信息
func handleMajorInvite(s session.Session, msg interface{}) (err error) {
	log.Debug("major:处理双修邀请消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMajorInvite)
	majorInt := csMsg.GetMajorType()
	fubenId := csMsg.GetFubenId()

	majorType := majortypes.MajorType(majorInt)
	if !majorType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"majorType": majorType,
			}).Warn("major:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = majorInvite(tpl, majorType, fubenId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("major:处理双修邀请消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("major:处理双修邀请消息完成")
	return nil

}

//双修邀请信息的逻辑
func majorInvite(pl player.Player, majorType majortypes.MajorType, fubenId int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
	flag := manager.HasMajorNum(majorType)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:今日双修次数已用完")
		playerlogic.SendSystemMessage(pl, lang.MajorInviteNoTimes)
		return
	}

	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()
	if marryInfo.SpouseId == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:前没有配偶,结婚后可邀请配偶双修")
		playerlogic.SendSystemMessage(pl, lang.MajorInviteNoSpouse)
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	if spl == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:您的配偶当前不在线,无法双修")
		playerlogic.SendSystemMessage(pl, lang.MajorInviteSpouseNoOnline)
		return
	}

	if marryInfo.Status != marrytypes.MarryStatusTypeMarried {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:举办婚礼成为正式夫妻才可进入")
		playerlogic.SendSystemMessage(pl, lang.MajorNoHoldWed)
		return
	}

	ss := spl.GetScene()
	if ss == nil || ss.MapTemplate().IsFuBen() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("major:您的配偶当前在其他副本场景,无法邀请")
		playerlogic.SendSystemMessage(pl, lang.MajorInviteSpouseInOtherFuBen)
		return
	}

	flag = manager.InviteFrequent(majorType)
	inviteTime := int64(0)
	if flag {
		playerlogic.SendSystemMessage(pl, lang.CommonOperFrequent)
		return
	} else {
		inviteTime = manager.InviteTime(majorType)
	}

	major.GetMajorService().MajorInvite(pl.GetId(), pl.GetName(), spl.GetId(), spl.GetName(), majorType, fubenId)
	scMajorInvitePushSpouse := pbutil.BuildSCMajorInvitePushSpouse(pl.GetId(), int32(majorType), fubenId)
	spl.SendMsg(scMajorInvitePushSpouse)
	scMajorInvite := pbutil.BuildSCMajorInvite(inviteTime, int32(majorType), fubenId)
	pl.SendMsg(scMajorInvite)
	return
}
