package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	shoptypes "fgame/fgame/game/shop/types"
	"fmt"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_DAOJU_CHANGE_TYPE), dispatch.HandlerFunc(handlePlayerDaoJuChange))
}

func handlePlayerDaoJuChange(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理玩家补充差价替换道具请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSJieYiDaoJuChange)
	daoJu := jieyitypes.JieYiDaoJuType(csMsg.GetDaoJu())
	if !daoJu.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"daoJu":    int32(daoJu),
			}).Warn("jieyi: 道具类型不合法")
		return
	}

	err = playerDaoJuChange(tpl, daoJu)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理玩家补充差价替换道具请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理玩家补充差价替换道具请求消息,成功")

	return
}

func playerDaoJuChange(pl player.Player, daoJu jieyitypes.JieYiDaoJuType) (err error) {

	plObj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if plObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	lastDaoJu := plObj.GetDaoJuType()
	if int(daoJu) <= int(lastDaoJu) {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"daojuType":     int32(daoJu),
				"lastDaojuType": int32(lastDaoJu),
			}).Warn("jieyi: 道具级别不够")
		playerlogic.SendSystemMessage(pl, lang.JieYiDaoJuNotChange)
		return
	}

	lastDaoJuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(lastDaoJu)
	daoJuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(daoJu)
	if daoJuTemp == nil || lastDaoJuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"daojuType":     int32(daoJu),
				"lastDaojuType": int32(lastDaoJu),
			}).Warn("jieyi: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.JieYiTemplateNotExist)
		return
	}
	itemMap := daoJuTemp.GetNeedItemMap()
	lastItemMap := lastDaoJuTemp.GetNeedItemMap()
	lastTokenSliver := int64(0)
	lastTokenBind := int64(0)
	lastTokenGold := int64(0)
	for id, num := range lastItemMap {
		shopTempMap := shop.GetShopService().GetShopItemMap(id)
		goldTempList, ok := shopTempMap[shoptypes.ShopConsumeTypeGold]
		if !ok {
			continue
		}
		lastTokenGold += int64(goldTempList[0].ConsumeData1) * int64(num)
	}
	// 补差价升级
	isEnoughBuyTimes, shopIdMap := shoplogic.MaxBuyTimesForPlayerMapComplementGold(pl, lastTokenGold, itemMap)
	if !isEnoughBuyTimes {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("jieyi:购买物品失败,补差价")
		playerlogic.SendSystemMessage(pl, lang.JieYiBuChaJiaFail)
		return
	}

	tokenBind, tokenGold, tokenSliver := shoplogic.ShopCostData(pl, shopIdMap)

	needSliver := (tokenSliver - lastTokenSliver)
	needBind := (tokenBind - lastTokenBind)
	needGold := (tokenGold - lastTokenGold)

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if needSliver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(needSliver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"daojuType":     int32(daoJu),
				"lastDaojuType": int32(lastDaoJu),
			}).Warn("jieyi: 银两不足,无法替换")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	if needGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"daojuType":     int32(daoJu),
				"lastDaojuType": int32(lastDaoJu),
			}).Warn("jieyi: 元宝不足,无法替换")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	costBind := needBind + needGold
	if costBind != 0 {
		flag := propertyManager.HasEnoughGold(int64(costBind), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"daojuType":     int32(daoJu),
				"lastDaojuType": int32(lastDaoJu),
			}).Warn("jieyi: 元宝不足,无法替换")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	// 同步数据
	flag := jieyi.GetJieYiService().DaoJuChangeSucess(pl.GetId(), daoJu)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"daoJu":     int(daoJu),
			"lastDaoJu": int(lastDaoJu),
		}).Warn("jieyi: 改变失败")
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonJieYiDaoJuChange
	goldReason := fmt.Sprintf(goldUseReason.String(), lastDaoJu.String(), daoJu.String(), jieyitypes.JieYiItemUseTypeTiHuan.String())
	silverUseReason := commonlog.SilverLogReasonJieYiDaoJuChange
	silverReason := fmt.Sprintf(silverUseReason.String(), lastDaoJu.String(), daoJu.String(), jieyitypes.JieYiItemUseTypeTiHuan.String())
	flag = propertyManager.Cost(int64(needBind), int64(needGold), goldUseReason, goldReason, int64(needSliver), silverUseReason, silverReason)
	if !flag {
		panic(fmt.Errorf("jieyi: 替换结义道具消耗钱应该成功"))
	}
	//同步钱
	if needBind != 0 || needGold != 0 || needSliver != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	jieYi := plObj.GetJieYi()
	scJieBrotherInfoOnChange := pbutil.BuildSCJieBrotherInfoOnChange(plObj)
	jieyilogic.BroadcastJieYi(jieYi, scJieBrotherInfoOnChange)

	jieyilogic.JieYiMemberChanged(jieYi)

	scMsg := pbutil.BuildSCJieYiDaoJuChange(int32(daoJu))
	pl.SendMsg(scMsg)

	return
}
