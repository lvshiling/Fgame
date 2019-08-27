package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	transportationlogic "fgame/fgame/game/transportation/logic"
	"fgame/fgame/game/transportation/pbutil"
	playertransportation "fgame/fgame/game/transportation/player"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_TRANSPORTATION_TYPE), dispatch.HandlerFunc(handlerAllianceTransportation))
}

//仙盟押镖
func handlerAllianceTransportation(s session.Session, msg interface{}) (err error) {
	log.Debug("found:处理仙盟押镖请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)

	err = allianceTransportation(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("found:处理仙盟押镖请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("found：处理仙盟押镖请求完成")

	return
}

func allianceTransportation(pl player.Player) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}
	moveTemp := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplateFirst()
	s := scene.GetSceneService().GetWorldSceneByMapId(moveTemp.MapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:押镖路线场景不存在")

		return
	}

	//功能开启判断
	flag := pl.IsFuncOpen(funcopentypes.FuncOpenTypeAllianceTransportation)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:领取仙盟押镖任务错误，功能未开启，无法押镖")

		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	//生成镖车
	member, biaoChe, err := alliance.GetAllianceService().AddTransportation(pl.GetId())
	if err != nil {
		return
	}
	transportationlogic.AddBiaoChe(pl, biaoChe, s, moveTemp.GetPosition())

	//推送系统公告
	format := lang.GetLangService().ReadLang(lang.TransportationAllianceAcceptSystemBroadcast)
	allianceName := coreutils.FormatColor(chattypes.ColorTypePlayerName, member.GetAlliance().GetAllianceName())
	content := fmt.Sprintf(format, allianceName)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

	//镖车信息
	transportationObj := biaoChe.GetTransportationObject()
	scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(transportationObj, biaoChe)
	pl.SendMsg(scTransportBriefInfoNotice)

	//次数信息
	transManager := pl.GetPlayerDataManager(types.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	personalTimes := transManager.GetTranspotTimes()
	allianceTimes := alliance.GetAllianceService().GetAllianceTransportTimes(pl.GetId())
	scPlayerTransportationBriefInfo := pbutil.BuildSCPlayerTransportationBriefInfo(personalTimes, allianceTimes)
	pl.SendMsg(scPlayerTransportationBriefInfo)

	scAllianceTransportation := pbutil.BuildSCAllianceTransportation(biaoChe.GetTransportationObject().GetCreateTime())
	pl.SendMsg(scAllianceTransportation)

	return
}
