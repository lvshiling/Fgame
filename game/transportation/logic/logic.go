package logic

import (
	"context"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	"fgame/fgame/game/transportation/pbutil"
	playertransportation "fgame/fgame/game/transportation/player"
	transportationtemplate "fgame/fgame/game/transportation/template"
	"fgame/fgame/game/transportation/transpotation"
	transportationtypes "fgame/fgame/game/transportation/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func CheckIfCanTransportaion(pl player.Player) (typ transportationtypes.TransportationType, flag bool) {
	//功能开启判断
	flag = pl.IsFuncOpen(funcopentypes.FuncOpenTypePersonalTransportation)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:押镖功能未开启")
		return
	}
	transManager := pl.GetPlayerDataManager(types.PlayerTransportationType).(*playertransportation.PlayerTransportationDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//是否足够次数
	if !transManager.IsEnoughTransportTimes() {
		flag = false
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:没有足够次数")
		return
	}
	//生成镖车
	biaoChe := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoChe != nil {
		flag = true
		return
	}

	typ = transportationtypes.TransportationTypeSilver
	silverBiaoChe := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(transportationtypes.TransportationTypeSilver)
	//是否足够银两
	flag = propertyManager.HasEnoughSilver(int64(silverBiaoChe.BiaocheSilver))
	if flag {
		return
	}
	typ = transportationtypes.TransportationTypeGold
	goldBiaoChe := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(transportationtypes.TransportationTypeGold)
	//是否足够元宝
	flag = propertyManager.HasEnoughGold(int64(goldBiaoChe.BiaocheGold), true)
	if flag {
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Warn("transport:银两或元宝不足")
	return
}

func HandlePersonalTransportation(pl player.Player, typ transportationtypes.TransportationType) (err error) {
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
	AddBiaoChe(pl, biaoChe, s, moveTemp.GetPosition())

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

func HandleReceiveTransportationRew(pl player.Player) (err error) {
	biaoChe := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoChe == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:玩家镖车不存在")
		return
	}

	obj := biaoChe.GetTransportationObject()
	tem := transportationtemplate.GetTransportationTemplateService().GetTransportationTemplate(obj.GetTransportType())
	if tem == nil {
		log.WithFields(
			log.Fields{
				"transportType": obj.GetTransportType(),
			}).Warn("transport:镖车模板不存在")
		return
	}

	//领取资源
	rewItemMap := make(map[int32]int32)
	addGoldNum := int64(0)
	addSilverNum := int64(0)
	var goldReason commonlog.GoldLogReason
	var silverReason commonlog.SilverLogReason
	var itemGetReason commonlog.InventoryLogReason
	switch obj.GetState() {
	case transportationtypes.TransportStateTypeFail:
		{
			addGoldNum = int64(tem.BiaocheLoseGold)
			addSilverNum = int64(tem.BiaocheLoseSilver)
			rewItemMap = tem.GetLastItemMap()
			goldReason = commonlog.GoldLogReasonTransportationFailRew
			silverReason = commonlog.SilverLogReasonTransportationFailRew
			itemGetReason = commonlog.InventoryLogReasonTransportFailRew
			break
		}
	case transportationtypes.TransportStateTypeFinish:
		{
			addGoldNum = int64(tem.BiaocheAwardGold)
			addSilverNum = int64(tem.BiaocheAwardSilver)
			rewItemMap = tem.GetFinishItemMap()
			goldReason = commonlog.GoldLogReasonTransportationFinishRew
			silverReason = commonlog.SilverLogReasonTransportationFinishRew
			itemGetReason = commonlog.InventoryLogReasonTransportRew
			break
		}
	default:
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("transport:玩家镖车运输中")
		return
	}

	//奖励倍数
	activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeBiaoChe)
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemp, _ := activityTemp.GetActivityTimeTemplate(now, openTime, mergeTime)
	if timeTemp != nil {
		ratio := timeTemp.BeiShu
		addGoldNum = addGoldNum * int64(ratio)
		addSilverNum = addSilverNum * int64(ratio)
		rewItemMap = coreutils.MultMap(rewItemMap, ratio)
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//需要判断是否大于0
	if addGoldNum > 0 {
		goldReasonText := fmt.Sprintf(goldReason.String(), obj.GetTransportType().String())
		propertyManager.AddGold(addGoldNum, true, goldReason, goldReasonText)
	}

	if addSilverNum > 0 {
		silverReasonText := fmt.Sprintf(silverReason.String(), obj.GetTransportType().String())
		propertyManager.AddSilver(addSilverNum, silverReason, silverReasonText)
	}

	//获得物品
	if len(rewItemMap) > 0 {
		//背包空间
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if !inventoryManager.HasEnoughSlots(rewItemMap) {
			title := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
			content := lang.GetLangService().ReadLang(lang.TransportationNotEnoughSlot)
			emaillogic.AddEmail(pl, title, content, rewItemMap)
		} else {
			reasonText := fmt.Sprintf(itemGetReason.String(), obj.GetTransportType(), obj.GetState())
			flag := inventoryManager.BatchAdd(rewItemMap, itemGetReason, reasonText)
			if !flag {
				panic("transport: add item should be ok")
			}
		}
	}

	transpotation.GetTransportService().RemoveTransportation(pl.GetId())
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scReceiveTransportRew := pbutil.BuildSCReceiveTransportRew(addGoldNum, addSilverNum, obj.GetTransportType())
	pl.SendMsg(scReceiveTransportRew)
	return
}

func AddBiaoChe(pl player.Player, biaoChe *biaochenpc.BiaocheNPC, s scene.Scene, pos coretypes.Position) (err error) {
	playerScene := pl.GetScene()
	if playerScene == s {
		biaoChe.SetPosition(pos)
		s.AddSceneObject(biaoChe)
	} else {
		biaoChe.SetPosition(pos)
		scenelogic.NPCEnterScene(biaoChe, s, pos)
	}
	return
}

// 邮件数据
type postEmailData struct {
	title       string
	content     string
	sendTime    int64
	attachement map[int32]int32
}

func newPostEmailData(title, content string, sendTime int64, attachement map[int32]int32) *postEmailData {
	d := &postEmailData{
		title:       title,
		content:     content,
		sendTime:    sendTime,
		attachement: attachement,
	}

	return d
}

//仙盟镖车邮件
func AllianceTransportRewBroadcast(allianceId int64, title, content string, rewMap map[int32]int32) {
	memberList := alliance.GetAllianceService().GetAlliance(allianceId).GetMemberList()
	now := global.GetGame().GetTimeService().Now()
	for _, memberObj := range memberList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(memberObj.GetMemberId())
		if pl == nil {
			emaillogic.AddOfflineEmail(memberObj.GetMemberId(), title, content, rewMap)
		} else {
			ctx := scene.WithPlayer(context.Background(), pl)
			data := newPostEmailData(title, content, now, rewMap)
			pl.Post(message.NewScheduleMessage(onSendRewEmail, ctx, data, nil))
		}
	}
}

func onSendRewEmail(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*postEmailData)

	emaillogic.AddEmailDefinTime(pl, data.title, data.content, data.sendTime, data.attachement)
	return nil
}
