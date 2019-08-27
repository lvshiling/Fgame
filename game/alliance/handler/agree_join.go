package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	gamecommon "fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AGREE_JOIN_APPLY_TYPE), dispatch.HandlerFunc(handleAllianceAgreeJoinApply))
}

//处理仙盟加入申请
func handleAllianceAgreeJoinApply(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟加入申请")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAllianceAgreeJoinApply := msg.(*uipb.CSAllianceAgreeJoinApply)
	agree := csAllianceAgreeJoinApply.GetAgree()
	joinId := csAllianceAgreeJoinApply.GetJoinId()
	err = allianceAgreeJoin(tpl, joinId, agree)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"joinId":   joinId,
				"agree":    agree,
				"error":    err,
			}).Error("alliance:处理仙盟加入申请,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"joinId":   joinId,
			"agree":    agree,
		}).Debug("alliance:处理仙盟加入申请,完成")
	return nil

}

//仙盟加入
func allianceAgreeJoin(pl player.Player, joinId int64, agree bool) (err error) {
	err = agreeJoin(pl, joinId, agree)
	if err != nil {
		return
	}

	scMsg := pbutil.BuildSCAllianceAgreeJoinApply(joinId, agree)
	pl.SendMsg(scMsg)
	return
}

func agreeJoin(pl player.Player, joinId int64, agree bool) (err error) {
	//处理申请加入仙盟
	al, _, err := alliance.GetAllianceService().AgreeAllianceJoinApply(pl.GetId(), joinId, agree)
	if err != nil {
		terr, ok := err.(gamecommon.Error)
		if !ok {
			return
		}
		// 特殊处理
		if terr.Code() == lang.AllianceAlreadyFullApply {
			playerlogic.SendSystemMessage(pl, lang.AllianceAlreadyFullApply)
			agree = false
			err = nil
		} else {
			return
		}
	}

	//通知用户
	joinPlayer := player.GetOnlinePlayerManager().GetPlayerById(joinId)
	if joinPlayer != nil {
		if agree {
			//同步用户数据
			joinName := joinPlayer.GetName()
			joinSex := joinPlayer.GetSex()
			joinLevel := joinPlayer.GetLevel()
			joinForce := joinPlayer.GetForce()
			joinZhuanSheng := joinPlayer.GetZhuanSheng()
			//TODO xzk:优化不要用GetLingyuInfo
			joinLingyu := joinPlayer.GetLingyuInfo().AdvanceId
			joinVip := joinPlayer.GetVip()
			alliance.GetAllianceService().SyncMemberInfo(joinId, joinName, joinSex, joinLevel, joinForce, joinZhuanSheng, joinLingyu, joinVip)

			mem := alliance.GetAllianceService().GetAllianceMember(joinId)
			if mem == nil {
				panic("alliance agree join: 成员应该存在")
			}
			scAllianceInfo := pbutil.BuildSCAllianceInfo(al, mem)
			joinPlayer.SendMsg(scAllianceInfo)
		}

		allianceId := al.GetAllianceId()
		allianceName := al.GetAllianceObject().GetName()
		scAllianceAgreeJoinApplyToApply := pbutil.BuildSCAllianceAgreeJoinApplyToApply(allianceId, allianceName, agree)
		joinPlayer.SendMsg(scAllianceAgreeJoinApplyToApply)
	}

	//广播管理成员
	scAllianceAgreeJoinApplyToManager := pbutil.BuildSCAllianceAgreeJoinApplyToManager(joinId, agree)
	for _, manager := range al.GetAllManagers() {
		if manager.GetMemberId() == pl.GetId() {
			continue
		}
		managerPlayer := player.GetOnlinePlayerManager().GetPlayerById(manager.GetMemberId())
		if managerPlayer == nil {
			continue
		}
		managerPlayer.SendMsg(scAllianceAgreeJoinApplyToManager)
	}

	return
}
