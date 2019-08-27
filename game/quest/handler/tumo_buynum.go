package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_BUYNUM_TYPE), dispatch.HandlerFunc(handleQuestTuMoBuyNum))
}

//处理购买屠魔次数
func handleQuestTuMoBuyNum(s session.Session, msg interface{}) (err error) {
	log.Debug("quest:处理购买屠魔次数")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csQuestTuMoBuyNum := msg.(*uipb.CSQuestTuMoBuyNum)
	num := csQuestTuMoBuyNum.GetNum()

	err = questTuMoBuyNum(tpl, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
				"error":    err,
			}).Error("quest:处理购买屠魔次数,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理购买屠魔次数,完成")
	return nil
}

//购买屠魔次数
func questTuMoBuyNum(pl player.Player, num int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	flag, err := manager.IfReachBuyLimit()
	if err != nil {
		return err
	}
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest:购买屠魔任务次数达上限")
		playerlogic.SendSystemMessage(pl, lang.QuestTuMoBuyNumReachLimit)
		return
	}

	costGold := questtemplate.GetQuestTemplateService().GetQuestTuMoBuyNumCostGold()
	//元宝是否足够
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag = propertyManager.HasEnoughGold(int64(costGold), false)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest:当前元宝不足，请及时充值")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	goldReasonText := commonlog.GoldLogReasonTuMoBuyNum.String()
	flag = propertyManager.CostGold(costGold, false, commonlog.GoldLogReasonTuMoBuyNum, goldReasonText)
	if !flag {
		panic(fmt.Errorf("quest: tuMoBuyNum cost gold should be ok"))
	}
	//同步元宝
	propertylogic.SnapChangedProperty(pl)
	//增加购买次数
	err = manager.AddNumByBuy()
	if err != nil {
		return
	}
	buyNum := manager.GetBuyNum()
	extraNum := manager.GetExtraNum()
	scQuestTuMoBuyNum := pbutil.BuildSCQuestTuMoBuyNum(buyNum, extraNum)
	pl.SendMsg(scQuestTuMoBuyNum)
	return

}
