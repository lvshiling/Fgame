package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_PEIYANG_TYPE), dispatch.HandlerFunc(handleLingTongPeiYang))

}

//处理灵童培养信息
func handleLingTongPeiYang(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理灵童培养信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongPeiYang := msg.(*uipb.CSLingTongPeiYang)
	lingTongId := csLingTongPeiYang.GetLingTongId()
	num := csLingTongPeiYang.GetNum()

	err = lingTongPeiYang(tpl, lingTongId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
				"num":        num,
				"error":      err,
			}).Error("lingtong:处理灵童培养信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
			"num":        num,
		}).Debug("lingtong:处理灵童培养信息完成")
	return nil

}

//灵童培养逻辑
func lingTongPeiYang(pl player.Player, lingTongId int32, num int32) (err error) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
			"num":        num,
		}).Warn("lingtong:模板为空")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
			"num":        num,
		}).Warn("lingtong:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongInfo, flag := manager.GetLingTongInfo(lingTongId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
			"num":        num,
		}).Warn("lingtong:未激活该灵童")
		playerlogic.SendSystemMessage(pl, lang.LingTongNoActive)
		return
	}
	curLevel := lingTongInfo.GetPeiYangLevel()
	nextLevel := curLevel + 1
	nextLingTongPeiYangTemplate := lingTongTemplate.GetLingTongPeiYangByLevel(nextLevel)
	if nextLingTongPeiYangTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
			"num":        num,
		}).Warn("lingtong:培养已达满级")
		playerlogic.SendSystemMessage(pl, lang.LingTongPeiYangReachFull)
		return
	}

	reachCaoLiaoTemplate, flag := lingtongtemplate.GetLingTongTemplateService().GetLingTongPeiYangUpgrade(lingTongId, curLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
			"num":        num,
		}).Warn("lingtong:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	useItem := nextLingTongPeiYangTemplate.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
				"num":        num,
			}).Warn("lingtong:当前培养丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongEatClu.String(), lingTongId)
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonLingTongDevEatClu, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = manager.EatCulDan(lingTongId, reachCaoLiaoTemplate.Level)
	if !flag {
		panic(fmt.Errorf("lingtong:EatCulDan should be ok"))
	}
	lingtonglogic.LingTongPropertyChanged(pl)

	scLingTongPeiYang := pbutil.BuildSCLingTongPeiYang(lingTongId, lingTongInfo.GetPeiYangLevel(), lingTongInfo.GetPeiYangPro())
	pl.SendMsg(scLingTongPeiYang)
	return
}
