package handler

import (
	"context"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playerinventory "fgame/fgame/game/inventory/player"
	marry "fgame/fgame/game/marry/marry"
	pbutil "fgame/fgame/game/marry/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	commonlog "fgame/fgame/common/log"
	inventorylogic "fgame/fgame/game/inventory/logic"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrytemplate "fgame/fgame/game/marry/template"
	mtypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_PRE_GIFT), dispatch.HandlerFunc(handleMarryWedPreGift))
}

//处理赠送贺礼信息
func handleMarryWedPreGift(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:游车玩家赠送")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMarryPreGift := msg.(*uipb.CSMarryPreGift)
	giftType := csMarryPreGift.GetGiftType()
	period := csMarryPreGift.GetPeriod()
	err = webPreGift(tpl, period, mtypes.MarryPreGiftType(giftType))
	if err != nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
			"err":      err,
		}).Error("marry:游车玩家赠送,失败")
		return
	}

	log.WithFields(log.Fields{
		"playerId": pl.GetId(),
		"giftType": giftType,
		"period":   period,
	}).Debug("marry:游车玩家赠送,成功")
	return
}

func webPreGift(pl player.Player, period int32, giftType mtypes.MarryPreGiftType) (err error) {
	flag := checkPeriod(period) //校验场次
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
		}).Warn("marry:游车玩家赠送,无效婚姻场次")
		playerlogic.SendSystemMessage(pl, lang.MarryWedGiftNoWedExist)
		return
	}
	giftTemplate := marrytemplate.GetMarryTemplateService().GetMarryPreGiftTemplate(giftType)
	if giftTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
		}).Warn("marry:游车玩家赠送,无效礼物类型")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	marryInfo := marry.GetMarryService().GetMarrySceneData()
	if marryInfo == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
		}).Warn("marry:游车玩家赠送,婚礼场景不存在")
		playerlogic.SendSystemMessage(pl, lang.MarryWedGiftNoWedExist)
		return
	}

	playerId := marryInfo.PlayerId
	spouseId := marryInfo.SpouseId
	playerSn := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if playerSn == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
		}).Warn("marry:游车玩家赠送,玩家不在线")
		// playerlogic.SendSystemMessage(pl, lang.MarryWedGiftNoWedExist)
		// return
	}
	spouseSn := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if spouseSn == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
		}).Warn("marry:游车玩家赠送,配偶不在线")
		// playerlogic.SendSystemMessage(pl, lang.MarryWedGiftNoWedExist)
		// return
	}
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	goldNum := int64(giftTemplate.NeedGold)
	if goldNum > 0 { //使用元宝大于0要判断是否需要消耗元宝
		if !propertyManager.HasEnoughGold(goldNum, false) {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"giftType": giftType,
				"period":   period,
			}).Warn("marry:游车玩家赠送,元宝不足")

			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
		reasonText := commonlog.GoldLogReasonMarryPreGift.String()
		flag := propertyManager.CostGold(goldNum, flag, commonlog.GoldLogReasonMarryPreGift, reasonText)
		if !flag {
			panic(fmt.Errorf("marry:赠送礼物花费元宝应该成功"))
		}
	}

	//赠送回馈
	itemMap := giftTemplate.GetReturnItemMap()
	if !inventoryManager.HasEnoughSlots(itemMap) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"giftType": giftType,
			"period":   period,
		}).Debug("marry:游车玩家赠送,背包不足不足")
		title := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
		mailContent := lang.GetLangService().ReadLang(lang.MarryPreGiftContent)
		emaillogic.AddEmail(pl, title, mailContent, itemMap)
	} else {

		returnText := commonlog.InventoryLogReasonMarryPreGift.String()
		flag = inventoryManager.BatchAdd(itemMap, commonlog.InventoryLogReasonMarryPreGift, returnText)
		if !flag {
			panic(fmt.Errorf("marry:赠送礼物,回赠应该成功"))
		}
	}
	//发送结婚增加经验
	exp := giftTemplate.GetRewardExp()
	expPoint := giftTemplate.GetRewardExpPoint()
	giftPlayerId := pl.GetId()
	giftPlayerName := pl.GetName()
	addExd(playerSn, int32(giftType), exp, expPoint, giftPlayerId, giftPlayerName)
	addExd(spouseSn, int32(giftType), exp, expPoint, giftPlayerId, giftPlayerName)

	//同步属性
	propertylogic.SnapChangedProperty(pl)
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)
	rtnMsg := pbutil.BuildSCMarryPreGift(itemMap)
	pl.SendMsg(rtnMsg)

	return
}

func addExd(pl player.Player, giftType int32, exp int64, expPoint int64, giftPlayerId int64, giftPlayerName string) error {
	if pl == nil {
		return nil
	}
	splCtx := scene.WithPlayer(context.Background(), pl)
	rst := marryeventtypes.CreateMarryPreWedGiftEventData(giftType, exp, expPoint, giftPlayerId, giftPlayerName)
	msg := message.NewScheduleMessage(onAddExp, splCtx, rst, nil)
	pl.Post(msg)

	return nil
}

func onAddExp(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	info := result.(*marryeventtypes.MarryPreWedGiftEventData)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	preGiftText := commonlog.LevelLogReasonMarryPreGift.String()
	exp := info.GetExp()
	exppoint := info.GetExpPoint()
	pointExp := propertyManager.GetExpPoint(exppoint)
	propertyManager.AddExpPoint(exppoint, commonlog.LevelLogReasonMarryPreGift, preGiftText)
	propertyManager.AddExp(exp, commonlog.LevelLogReasonMarryPreGift, preGiftText)

	totalExp := exp + pointExp
	giftPlayerId := info.GetGiftPlayerId()
	giftPlayerName := info.GetGiftPlayerName()
	reply := pbutil.BuildSCMarryPreGiftMsg(giftPlayerId, giftPlayerName, info.GetGiftType(), totalExp)
	pl.SendMsg(reply)
	propertylogic.SnapChangedProperty(pl)
	return nil
}

func checkPeriod(period int32) bool {
	wedNum := marrytemplate.GetMarryTemplateService().GetMarryConstWedNum()
	wedTotalTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()

	now := global.GetGame().GetTimeService().Now()
	wedFirstTime, err := marrytemplate.GetMarryTemplateService().GetMarryFisrtWedTime(now)
	if err != nil {
		return false
	}
	if period < 0 || period > wedNum {
		return false
	}
	wedStarTime := wedFirstTime + int64(period-1)*wedTotalTime
	if wedStarTime < wedFirstTime || wedStarTime > now {
		return false
	}

	return true
}
