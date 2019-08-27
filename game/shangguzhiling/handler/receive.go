package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shangguzhiling/pbutil"
	playershangguzhiling "fgame/fgame/game/shangguzhiling/player"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_RECEIVE_TYPE), dispatch.HandlerFunc(handleShangguzhilingReceive))
}

func handleShangguzhilingReceive(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSShangguzhilingReceive)
	lingshouType := shangguzhilingtypes.LingshouType(csMsg.GetType())

	if !lingshouType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:上古之灵领取奖励请求,灵兽类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = shangguzhilingReceive(tpl, lingshouType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shangguzhiling:上古之灵领取奖励请求,错误")

		return err
	}
	return nil
}

func shangguzhilingReceive(pl player.Player, lingshouType shangguzhilingtypes.LingshouType) (err error) {
	lingShouManager := pl.GetPlayerDataManager(playertypes.PlayerShangguzhilingDataManagerType).(*playershangguzhiling.PlayerShangguzhilingDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//是否解锁
	if !lingShouManager.IsLingShouUnlock(lingshouType) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:上古之灵领取奖励请求,灵兽未解锁")
		playerlogic.SendSystemMessage(pl, lang.ShangguzhilingLingShouUnLock)
		return
	}
	obj := lingShouManager.GetLingShouObj(lingshouType)

	//相关模板
	curLevel := obj.GetLevel()
	baseTemp := shangguzhilingtemplate.GetShangguzhilingTemplateService().GetLingShouTemplate(lingshouType)
	if baseTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
			}).Warn("shangguzhiling:上古之灵领取奖励请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	levelTemp := baseTemp.GetLevelTemp(curLevel)
	if levelTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
				"curLevel":     curLevel,
			}).Warn("shangguzhiling:上古之灵领取奖励请求,等级模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//CD
	cdTime := int64(levelTemp.BaoxiangCd)
	now := global.GetGame().GetTimeService().Now()
	lastReceiveTime := obj.GetLastReceiveTime()
	if lastReceiveTime+cdTime > now {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"error":           err,
				"lingshouType":    lingshouType,
				"curLevel":        curLevel,
				"lastReceiveTime": lastReceiveTime,
				"cdTime":          cdTime,
			}).Warn("shangguzhiling:上古之灵领取奖励请求,冷却时间不足")
		playerlogic.SendSystemMessage(pl, lang.ShangguzhilingReceiveCDNotEnough)
		return
	}

	//掉落
	dropId := levelTemp.BaoxiangDrop
	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
	if dropData == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"error":        err,
				"lingshouType": lingshouType,
				"dropId":       dropId,
			}).Warn("shangguzhiling:上古之灵领取奖励请求,掉落模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//更新领取时间
	lingShouManager.UpdateLastReceiveTime(lingshouType)

	//获得物品
	itemGetReason := commonlog.InventoryLogReasonLingShouReceive
	itemGetReasonText := fmt.Sprintf(itemGetReason.String(), lingshouType.String())
	flag := inventoryManager.AddItemLevel(dropData, itemGetReason, itemGetReasonText)
	if !flag {
		panic("shangguzhiling:增加物品应该成功")
	}

	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCShangguzhilingReceive(lingshouType, obj.GetLastReceiveTime(), dropData)
	pl.SendMsg(scMsg)
	return
}
