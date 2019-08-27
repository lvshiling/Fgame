package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancetypes "fgame/fgame/game/alliance/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_CALL_TYPE), dispatch.HandlerFunc(handleAllianceMemberCall))
}

//处理仙盟召集令
func handleAllianceMemberCall(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟召集")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSAllianceMemberCall)
	typeInt := csMsg.GetCallType()

	callType := alliancetypes.AllianceCallType(typeInt)
	if !callType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"callType": callType,
			}).Warn("alliance:仙盟召集，召集类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = allianceMemberCall(tpl, callType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟召集,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟召集,完成")
	return nil

}

//仙盟召集
func allianceMemberCall(pl player.Player, callType alliancetypes.AllianceCallType) (err error) {
	s := pl.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟召集，不在场景中")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}
	al := alliance.GetAllianceService().GetAlliance(pl.GetAllianceId())
	if al == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟召集，没有仙盟")
		playerlogic.SendSystemMessage(pl, lang.PlayerNoInScene)
		return
	}

	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	if allianceManager.IsMemberCallCD(callType) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟召集，召集CD中")
		playerlogic.SendSystemMessage(pl, lang.AllianceMemberCallNoticeCD)
		return
	}

	// 盟主/副盟主额外处理
	member := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	if member.IsMengZhu() || member.IsFuMengZhu() {
		scMsg := pbutil.BuildSCAllianceMemberCallBroadcast(pl.GetName(), s.MapId(), pl.GetPos(), int32(callType))
		for _, mem := range al.GetMemberList() {
			if pl.GetId() == mem.GetMemberId() {
				continue
			}

			memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
			if memPl == nil {
				continue
			}
			memPl.SendMsg(scMsg)
		}
	}

	//广播帮派
	pos := pl.GetPos()
	memberName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, pl.GetName())
	posStr := coreutils.FormatColor(chattypes.ColorTypeBossMap, coreutils.FormatStrPosiotn(fmt.Sprintf("%.0f", pl.GetPos().X), fmt.Sprintf("%.0f", pl.GetPos().Z)))
	mapName := coreutils.FormatColor(chattypes.ColorTypeBossMap, s.MapTemplate().Name)
	args := []int64{int64(chattypes.ChatAllianceRescue), int64(funcopentypes.FuncOpenTypeAlliance), int64(pl.GetScene().MapId()), int64(pos.X), int64(pos.Y), int64(pos.Z)}
	link := coreutils.FormatLink(chattypes.ButtonTypeToRescue, args)

	format := lang.GetLangService().ReadLang(lang.AllianceMemberCallNotice)
	content := fmt.Sprintf(format, memberName, mapName, posStr, link)
	chatlogic.SystemBroadcastAlliance(al, chattypes.MsgTypeText, []byte(content))

	lastCallTime, flag := allianceManager.UpdateLastMemberCall(callType)
	if !flag {
		panic("alliance:更新召集时间应该成功")
	}

	scMsg := pbutil.BuildSCAllianceMemberCall(lastCallTime, int32(callType))
	pl.SendMsg(scMsg)
	return
}
