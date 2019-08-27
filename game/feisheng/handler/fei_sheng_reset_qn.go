package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	feishenglogic "fgame/fgame/game/feisheng/logic"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	feishengtemplate "fgame/fgame/game/feisheng/template"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_RESET_QN_TYPE), dispatch.HandlerFunc(handleFeiShengResteQn))
}

//处理飞升重置潜能
func handleFeiShengResteQn(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理飞升重置潜能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = feiShengResteQn(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("feisheng:处理飞升重置潜能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("feisheng:处理飞升重置潜能消息完成")
	return nil
}

//飞升重置潜能界面逻辑
func feiShengResteQn(pl player.Player) (err error) {

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengInfo := feiManager.GetFeiShengInfo()
	feiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(feiShengInfo.GetFeiLevel())
	if feiTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feisheng:飞升重置潜能失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//元宝
	needGold := int64(feiTemplate.XidianGold)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, true) {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("feisheng:飞升重置潜能失败，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	if needGold > 0 {
		useGoldReason := commonlog.GoldLogReasonFeiShengResetQn
		useGoldReasonText := fmt.Sprintf(useGoldReason.String(), feiShengInfo.GetFeiLevel())
		flag := propertyManager.CostGold(needGold, true, useGoldReason, useGoldReasonText)
		if !flag {
			panic(fmt.Errorf("feisheng: feisheng eat dan cost gold should be ok"))
		}
	}

	feiManager.ResetQn()
	propertylogic.SnapChangedProperty(pl)
	feishenglogic.FeiShengPropertyChanged(pl)

	scMsg := pbutil.BuildSCFeiShengResteQn(feiShengInfo)
	pl.SendMsg(scMsg)
	return
}
