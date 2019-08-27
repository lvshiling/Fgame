package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
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
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	transportationlogic "fgame/fgame/game/transportation/logic"
	"fgame/fgame/game/transportation/pbutil"
	playertransportation "fgame/fgame/game/transportation/player"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
	transportationtypes "fgame/fgame/game/transportation/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_PERSONAL_TRANSPORTATION_TYPE), dispatch.HandlerFunc(handlerPersonalTransportation))
}

//押镖
func handlerPersonalTransportation(s session.Session, msg interface{}) (err error) {
	log.Debug("transport:处理押镖请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csPersonalTransportation := msg.(*uipb.CSPersonalTransportation)
	typ := csPersonalTransportation.GetTyp()

	transportType := transportationtypes.TransportationType(typ)
	if !transportType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      transportType,
			}).Warn("transport:参数错误")

		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = personalTransportation(tpl, transportType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"typ":      typ,
				"err":      err,
			}).Error("transport:处理押镖请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"typ":      typ,
		}).Debug("transport:处理押镖请求完成")

	return
}

func personalTransportation(pl player.Player, typ transportationtypes.TransportationType) (err error) {
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
	tem := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(typ)
	if tem == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("transport:参数错误,模板不存在")

		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//功能开启判断
	flag := pl.IsFuncOpen(funcopentypes.FuncOpenTypePersonalTransportation)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:领取押镖任务错误，功能未开启，无法押镖")

		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	transManager := pl.GetPlayerDataManager(types.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//是否足够次数
	if !transManager.IsEnoughTransportTimes() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
			}).Warn("transport:领取押镖任务错误，次数不足")

		playerlogic.SendSystemMessage(pl, lang.TransportationAcceptNumNoEnough)
		return
	}

	if typ == transportationtypes.TransportationTypeSilver {
		//是否足够银两
		flag := propertyManager.HasEnoughSilver(int64(tem.BiaocheSilver))
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
				}).Warn("transport:领取押镖任务错误，银两不足")

			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	if typ == transportationtypes.TransportationTypeGold {
		//是否足够元宝
		flag := propertyManager.HasEnoughGold(int64(tem.BiaocheGold), true)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ,
				}).Warn("transport:领取押镖任务错误，元宝不足")

			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

	}

	//生成镖车
	biaoChe, err := transpotation.GetTransportService().AddPersonalTransportation(pl.GetId(), pl.GetName(), pl.GetAllianceId(), typ)
	if err != nil {
		return
	}

	//消耗银两
	if typ == transportationtypes.TransportationTypeSilver {
		silverReason := commonlog.SilverLogReasonTransportationCost
		silverReasonText := fmt.Sprintf(silverReason.String(), typ.String())
		flag := propertyManager.CostSilver(int64(tem.BiaocheSilver), silverReason, silverReasonText)
		if !flag {
			panic("transportation：消耗银两应该成功")
		}
	}

	//消耗元宝
	if typ == transportationtypes.TransportationTypeGold {
		goldReason := commonlog.GoldLogReasonTransportationCost
		goldReasonText := fmt.Sprintf(goldReason.String(), typ.String())
		flag := propertyManager.CostGold(int64(tem.BiaocheGold), true, goldReason, goldReasonText)
		if !flag {
			panic("transpotation:消耗元宝应该成功")
		}
	}

	transportationObj := biaoChe.GetTransportationObject()
	transportationlogic.AddBiaoChe(pl, biaoChe, s, moveTemp.GetPosition())

	//消耗次数
	transManager.AcceptTransportation(typ)

	//推送系统公告
	format := lang.GetLangService().ReadLang(lang.TransportationAcceptSystemBroadcast)
	args := []int64{int64(chattypes.ChatLinkTypeNpc), int64(tem.BiologyId)}
	joinLink := coreutils.FormatLink("【我要押镖】", args)
	content := fmt.Sprintf(format, pl.GetName(), typ.String(), joinLink)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	//镖车信息
	scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(transportationObj, biaoChe)
	pl.SendMsg(scTransportBriefInfoNotice)

	//次数信息
	personalTimes := transManager.GetTranspotTimes()
	allianceTimes := alliance.GetAllianceService().GetAllianceTransportTimes(pl.GetId())
	scPlayerTransportationBriefInfo := pbutil.BuildSCPlayerTransportationBriefInfo(personalTimes, allianceTimes)
	pl.SendMsg(scPlayerTransportationBriefInfo)

	scPersonalTransportation := pbutil.BuildSCPersonalTransportation(biaoChe.GetTransportationObject().GetCreateTime())
	pl.SendMsg(scPersonalTransportation)

	return
}
