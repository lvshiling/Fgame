package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/emperor/emperor"
	emperorlogic "fgame/fgame/game/emperor/logic"
	"fgame/fgame/game/emperor/pbutil"
	emperortemplate "fgame/fgame/game/emperor/template"
	playermarry "fgame/fgame/game/marry/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMPEROR_ROB_TYPE), dispatch.HandlerFunc(handleEmperorRob))
}

//处理帝王争抢信息
func handleEmperorRob(s session.Session, msg interface{}) (err error) {
	log.Debug("emperor:处理帝王争抢信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = emperorRob(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("emperor:处理帝王争抢信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("emperor:处理帝王争抢信息完成")
	return nil
}

//处理帝王争抢界面信息逻辑
func emperorRob(pl player.Player) (err error) {
	//玩家信息
	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := marryManager.GetMarryInfo()
	spouseName := ""
	if marryInfo.SpouseId != 0 {
		spouseName = marryInfo.SpouseName
	}

	oldEmperorId, robNum := emperor.GetEmperorService().GetEmperorIdAndRobNum()
	if oldEmperorId == pl.GetId() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:已经是帝王")
		playerlogic.SendSystemMessage(pl, lang.EmperorWasMyself)
		return
	}
	nextRobNum := robNum + 1
	//元宝是否足够
	needGold := emperortemplate.GetEmperorTemplateService().GetEmperorRobNeedGold(nextRobNum)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(int64(needGold), false)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:当前元宝不足，请及时充值")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//抢帝王
	dropItemList, err := emperor.GetEmperorService().EmperorRobbed(pl, spouseName, nextRobNum)
	if err != nil {
		return
	}
	//消耗元宝
	goldReasonText := commonlog.GoldLogReasonRobEmperor.String()
	flag = propertyManager.CostGold(needGold, false, commonlog.GoldLogReasonRobEmperor, goldReasonText)
	if !flag {
		panic(fmt.Errorf("emperor: emperorRob cost gold should be ok"))
	}
	//同步元宝
	propertylogic.SnapChangedProperty(pl)
	//添加物品
	if len(dropItemList) != 0 {
		err = emperorlogic.OpenBoxReward(pl, dropItemList, false)
		if err != nil {
			return
		}
	}
	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	scEmperorRob := pbuitl.BuildSCEmperorRob(emperorObj, dropItemList, true)
	pl.SendMsg(scEmperorRob)
	oldpl := player.GetOnlinePlayerManager().GetPlayerById(oldEmperorId)
	if oldpl == nil {
		return
	}
	oldpl.SendMsg(scEmperorRob)
	return
}
