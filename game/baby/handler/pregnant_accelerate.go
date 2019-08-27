package handler
/*
import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_BORN_ACCELERATE_TYPE), dispatch.HandlerFunc(handleBabyBornAccelerate))
}

//处理加速出生
func handleBabyBornAccelerate(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理加速出生消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = babyBornAccelerate(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理加速出生消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理加速出生消息完成")
	return nil

}

//加速出生界面逻辑
func babyBornAccelerate(pl player.Player) (err error) {

	// 是否怀孕状态
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	if !babyManager.IsPregnant() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("baby:处理加速出生消息,不是怀孕状态")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	pregnantInfo := babyManager.GetPregnantInfo()
	now := global.GetGame().GetTimeService().Now()
	needGold := babyConstantTemplate.GetAccelerateNeedGold(now, pregnantInfo.GetPregnantTime())
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("baby:处理加速出生消息, 元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	goldCostReason := commonlog.GoldLogReasonBabyAccelerateUse
	flag := propertyManager.CostGold(needGold, false, goldCostReason, goldCostReason.String())
	if !flag {
		panic(fmt.Errorf("baby: 消耗元宝应该成功"))
	}

	//加速出生
	baby := babyManager.AccelerateBorn()
	if baby == nil {
		panic(fmt.Errorf("baby: 宝宝出生应该成功"))
	}
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCBabyBornAccelerate()
	pl.SendMsg(scMsg)
	return
}
*/