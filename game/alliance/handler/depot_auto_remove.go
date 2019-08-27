package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/center/center"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DEPOT_AUTO_REMOVE_TYPE), dispatch.HandlerFunc(handleAllianceDepotAutoRemove))
}

//处理仙盟仓库自动销毁
func handleAllianceDepotAutoRemove(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟仓库自动销毁")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSAllianceDepotAutoRemove)
	isAuto := csMsg.GetIsAuto()
	zhuansheng := csMsg.GetMaxZhuanSheng()
	qualityInt := csMsg.GetMaxQuality()

	quality := itemtypes.ItemQualityType(qualityInt)
	if !quality.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isAuto":   isAuto,
			}).Warn("alliance:处理仙盟仓库自动销毁,品质参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = allianceDepotAutoRemove(tpl, isAuto, zhuansheng, quality)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isAuto":   isAuto,
				"error":    err,
			}).Error("alliance:处理仙盟仓库自动销毁,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"isAuto":   isAuto,
		}).Debug("alliance:处理仙盟仓库自动销毁,完成")
	return nil
}

//仙盟仓库自动销毁
func allianceDepotAutoRemove(pl player.Player, isAuto, zhuansheng int32, quality itemtypes.ItemQualityType) (err error) {
	if !center.GetCenterService().IsAllianceOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟仓库关闭中")
		playerlogic.SendSystemMessage(pl, lang.AllianceDepotClose)
		return
	}
	al, err := alliance.GetAllianceService().AutoRemoveDepot(pl.GetId(), isAuto, zhuansheng, quality)
	if err != nil {
		return
	}

	//设置推送
	for _, member := range al.GetMemberList() {
		if member.GetMemberId() == pl.GetId() {
			continue
		}

		memberPlayer := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPlayer == nil {
			continue
		}

		scMsg := pbutil.BuildSCAllianceDepotSettingNotice(isAuto, zhuansheng, int32(quality))
		memberPlayer.SendMsg(scMsg)
	}

	scMsg := pbutil.BuildSCAllianceDepotAutoRemove(isAuto, zhuansheng, int32(quality))
	pl.SendMsg(scMsg)
	return
}
