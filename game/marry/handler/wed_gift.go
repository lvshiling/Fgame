package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	fireworkslogic "fgame/fgame/game/fireworks/logic"
	friendlogic "fgame/fgame/game/friend/logic"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marryscenetypes "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GIFT_TYPE), dispatch.HandlerFunc(handleMarryWedGift))
}

//处理赠送贺礼信息
func handleMarryWedGift(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理赠送贺礼处理消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMarryWedGift := msg.(*uipb.CSMarryWedGift)
	period := csMarryWedGift.GetPeriod()
	grade := csMarryWedGift.GetGrade()
	autoFlag := csMarryWedGift.GetAutoFlag()

	err = marryWedGift(tpl, period, grade, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"period":   period,
				"grade":    grade,
				"autuFlag": autoFlag,
				"error":    err,
			}).Error("marry:处理赠送贺礼处理消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理赠送贺礼处理消息完成")
	return nil
}

//处理赠送贺礼处理信息逻辑
func marryWedGift(pl player.Player, period int32, grade int32, autoFlag bool) (err error) {
	marryScene := marry.GetMarryService().GetScene()
	if marryScene != pl.GetScene() {
		return
	}

	sd := marry.GetMarryService().GetMarrySceneData()
	if sd.Status != marryscenetypes.MarrySceneStatusBanquet {
		return
	}
	banquetType := sd.Grade
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	wedNum := marrytemplate.GetMarryTemplateService().GetMarryConstWedNum()
	giftTemplate := marrytemplate.GetMarryTemplateService().GetMarryGiftTeamplate(grade)
	wedTotalTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()
	if giftTemplate == nil || period < 0 || period > wedNum {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"grade":    grade,
			"period":   period,
			"autoFlag": autoFlag,
		}).Warn("marry:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	now := global.GetGame().GetTimeService().Now()
	wedFirstTime, err := marrytemplate.GetMarryTemplateService().GetMarryFisrtWedTime(now)
	if err != nil {
		return err
	}
	wedStarTime := wedFirstTime + int64(period-1)*wedTotalTime
	if wedStarTime < wedFirstTime || wedStarTime > now {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"grade":    grade,
			"period":   period,
			"autoFlag": autoFlag,
		}).Warn("marry:无效婚期场次")
		playerlogic.SendSystemMessage(pl, lang.MarryWedGiftNoWedExist)
		return
	}
	flag := marry.GetMarryService().IfCanGiveWedGift(period)
	if !flag {
		playerlogic.SendSystemMessage(pl, lang.MarryWedGiftNoWedExist)
		return
	}

	//扣减贺礼
	isReturn := marrylogic.PlayerGiveWeddingGift(pl, giftTemplate, autoFlag)
	if isReturn {
		return
	}
	rosesNum := giftTemplate.UseItemAmount
	costSilver := int64(giftTemplate.UseSilver)
	useItemId := giftTemplate.UseItemId

	itemTemplate := item.GetItemService().GetItem(int(useItemId))
	if itemTemplate == nil {
		return
	}

	//豪气值
	heroism := giftTemplate.AddNum
	outOfTime := wedStarTime + wedTotalTime
	manager.AddHeroism(heroism, outOfTime)
	//更新场景数据
	marryScene.SceneDelegate().(marryscenetypes.MarrySceneData).GiveGift(pl, period, useItemId, rosesNum, costSilver, int64(heroism))

	buffNum := giftTemplate.BuffAmount
	banquetTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeWed, banquetType)
	buffId := banquetTemplate.AddBuffId
	//增加buff

	scenelogic.AddBuffs(pl, buffId, pl.GetId(), buffNum, common.MAX_RATE)

	name := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(1)))
	linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(useItemId)}
	itemNameLink := coreutils.FormatLink(itemName, linkArgs)
	giftType := giftTemplate.GetGiftType()
	var content string
	switch giftType {
	case marrytypes.MarryGiftTypeItem:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryWedGiveGiftFlower), name, itemNameLink)
			wedPlayerId := sd.PlayerId
			wedSpouseId := sd.SpouseId
			wedPlayerName := sd.PlayerName
			wedSpouseName := sd.SpouseName
			playerIdList := make([]int64, 0, 3)
			playerNameList := make([]string, 0, 3)
			playerIdList = append(playerIdList, wedPlayerId, wedSpouseId, pl.GetId())
			playerNameList = append(playerNameList, wedPlayerName, wedSpouseName, pl.GetName())
			friendlogic.BroadcastMsg(pl, useItemId, rosesNum, playerIdList, playerNameList)
			break
		}
	case marrytypes.MarryGiftTypeSilver:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryWedGiveGiftSilver), name, int32(costSilver))
			break
		}
	case marrytypes.MarryGiftTypeFireworks:
		{
			content = fmt.Sprintf(lang.GetLangService().ReadLang(lang.FireworksNotice), name, itemNameLink)
			fireworkslogic.BroadcastMsg(pl, useItemId, int32(giftTemplate.UseItemAmount))
			break
		}
	}

	if len(content) != 0 {
		//跑马灯
		noticelogic.NoticeNumBroadcastScene(pl.GetScene(), []byte(content), 0, int32(1))
		//系统频道
		chatlogic.BroadcastScene(pl.GetScene(), chattypes.MsgTypeText, []byte(content))
	}

	scMarryWedGift := pbuitl.BuildSCMarryWedGift(buffId, period, grade, autoFlag)
	pl.SendMsg(scMarryWedGift)
	return
}
