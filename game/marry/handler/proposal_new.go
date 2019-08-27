package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	playerfriend "fgame/fgame/game/friend/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_PROPOSAL_TYPE), dispatch.HandlerFunc(handleMarryProposal_New))
}

//处理求婚信息
func handleMarryProposal_New(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理求婚消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMarryProposal := msg.(*uipb.CSMarryProposal)
	ring := csMarryProposal.GetRingType()
	proposalId := csMarryProposal.GetPlayerId()
	proposalName := csMarryProposal.GetPlayerName()

	err = marryProposal_new(tpl, marrytypes.MarryRingType(ring), proposalId, proposalName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"ring":       ring,
				"proposalId": proposalId,
				"error":      err,
			}).Error("marry:处理求婚消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理求婚消息完成")
	return nil
}

//处理求婚信息逻辑
func marryProposal_new(pl player.Player, ringType marrytypes.MarryRingType, proposalId int64, proposalName string) (err error) {
	//参数校验
	if !ringType.Valid() {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"ringType":   ringType,
			"proposalId": proposalId,
		}).Warn("marry:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if ringType.HoutaiType() != marrytemplate.GetMarryTemplateService().GetHouTaiType() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"ringType":      ringType,
			"ringHouTai":    ringType.HoutaiType(),
			"proposalId":    proposalId,
			"currentHoutai": marrytemplate.GetMarryTemplateService().GetHouTaiType(),
		}).Warn("marry:求婚不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	friendManager := pl.GetPlayerDataManager(types.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	flag := friendManager.IsFriend(proposalId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"ringType":   ringType,
			"proposalId": proposalId,
		}).Warn("marry:未添加对方为好友,无法同意对方的求婚请求")
		playerlogic.SendSystemMessage(pl, lang.MarryProposalNoFriend, proposalName)
		return
	}

	spl := player.GetOnlinePlayerManager().GetPlayerById(proposalId)
	if spl == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"ringType":   ringType,
			"proposalId": proposalId,
		}).Warn("marry:玩家不在线,无法求婚")
		playerlogic.SendSystemMessage(pl, lang.MarrySpouseNoOnline)
		return
	}

	mySex := pl.GetSex()
	spouseSex := spl.GetSex()
	if mySex == spouseSex {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"ringType":   ringType,
			"proposalId": proposalId,
		}).Warn("marry:只能向异性求婚")
		playerlogic.SendSystemMessage(pl, lang.MarrySpouseSameSex)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	marryObj := marry.GetMarryService().GetMarry(proposalId)
	if marryInfo.SpouseId != 0 || marryObj != nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"ringType":   ringType,
			"proposalId": proposalId,
		}).Warn("marry:您或对方已拥有伴侣,无法进行求婚")
		playerlogic.SendSystemMessage(pl, lang.MarryHasedSpouse)
		return
	}

	friendObj := friendManager.GetFriend(proposalId)
	intimacy := friendObj.Point
	needIntimacy, flag := marrytemplate.GetMarryTemplateService().GetMarryConstIntimacy()
	if !flag {
		return
	}
	if intimacy < needIntimacy {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"ringType":   ringType,
			"proposalId": proposalId,
		}).Warn("marry:双方亲密度过低")
		intimacyStr := fmt.Sprintf("%d", needIntimacy)
		playerlogic.SendSystemMessage(pl, lang.MarryIntimacyNoEnough, intimacyStr)
		return
	}

	//是否已向其它人求过婚
	marryRingObj := marry.GetMarryService().GetMarryProposalRing(pl.GetId())
	if marryRingObj != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"ringType":   ringType,
				"proposalId": proposalId,
			}).Warn("marry:同时只能向一个人求婚,您已经向其它求婚")
		playerlogic.SendSystemMessage(pl, lang.MarryProposalIsExist, marryRingObj.PeerName)
		return
	}

	//获取婚烟对戒
	// itemTempalte := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeWedRing, ringType.ItemWedRingSubType())

	// ringTypeInt := int32(ringType)
	// merrySubRing := marrytypes.CreateMarryBanquetRingSubType(ringTypeInt) //转换成婚戒子类型配置
	ringItem := marrytemplate.GetMarryTemplateService().GetMarryBanquetTemplateByHouTai(ringType.HoutaiType(), marrytypes.MarryBanquetTypeRing, ringType.BanquetSubTypeRing())
	if ringItem == nil {
		log.WithFields(log.Fields{
			"playerId":             pl.GetId(),
			"ringType":             ringType,
			"proposalId":           proposalId,
			"MarryBanquetTypeRing": marrytypes.MarryBanquetTypeRing,
			"merrySubRing":         ringType.BanquetSubTypeRing(),
		}).Warn("marry:结婚类型错误")
		playerlogic.SendSystemMessage(pl, lang.MarryQiuHunNotExists)
		return
	}
	// if itemTempalte == nil {
	// 	log.WithFields(log.Fields{
	// 		"playerId":   pl.GetId(),
	// 		"ringType":   ringType,
	// 		"proposalId": proposalId,
	// 	}).Warn("marry:结婚类型错误")
	// 	playerlogic.SendSystemMessage(pl, lang.MarryQiuHunNotExists)
	// 	return
	// }

	// gold := marry.GetMarrySetService().GetCostGold(marrytypes.MarryBanquetTypeRing, ringType.BanquetSubTypeRing())
	buyGold := int64(ringItem.UseGold)
	bindGold := int64(ringItem.UseBinggold)
	silver := int64(ringItem.UseSilver)
	// buyGold := int64(gold) //修改了从多版本的获取
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//判断是否足够金钱
	// if !propertyManager.HasEnoughGold(buyGold, false) {
	// 	log.WithFields(log.Fields{}).Debug("marryProposal_new: not enought gold")
	// 	playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
	// 	return
	// }

	if !propertyManager.HasEnoughCost(bindGold, buyGold, silver) {
		log.WithFields(log.Fields{}).Debug("marryProposal_new: not enought gold")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//TODO:zrc 有可能已经求过了
	//记录求婚婚戒
	marry.GetMarryService().MarryProposalRing(pl.GetId(), proposalId, spl.GetName(), ringType)

	//扣除元宝
	buyText := commonlog.GoldLogReasonMarryProposal.String()

	flag = propertyManager.Cost(bindGold, buyGold, commonlog.GoldLogReasonMarryProposal, buyText, silver, commonlog.SilverLogReasonMarryProposal, buyText)
	if !flag {
		panic(fmt.Errorf("marryProposal_new:cost gold should be success"))
	}
	//同步物品
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	//发送事件
	eventData := marryeventtypes.CreateMarryProposalEventData(pl.GetId(), proposalId, ringType)
	gameevent.Emit(marryeventtypes.EventTypeMarryProposal, nil, eventData)

	scMarryProposal := pbuitl.BuildSCMarryProposal(0)
	pl.SendMsg(scMarryProposal)
	return
}
