package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/core/template"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	feishengtemplate "fgame/fgame/game/feisheng/template"
	feishengtypes "fgame/fgame/game/feisheng/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	gametemplate "fgame/fgame/game/template"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_SAN_GONG_TYPE), dispatch.HandlerFunc(handleFeiShengSanGong))
}

//处理飞升散功
func handleFeiShengSanGong(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理飞升散功消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFeiShengSanGong)
	typ := csMsg.GetType()

	sanGongType := feishengtypes.SanGongType(typ)
	if !sanGongType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"sanGongType": sanGongType,
			}).Warn("feisheng:飞升散功失败，类型不存在")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = feiShengSanGong(tpl, sanGongType.GetReduceLevelNum())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("feisheng:处理飞升散功消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("feisheng:处理飞升散功消息完成")
	return nil

}

const (
	sangongLimit = 200
)

//飞升散功界面逻辑
func feiShengSanGong(pl player.Player, reduceNum int32) (err error) {

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengInfo := feiManager.GetFeiShengInfo()
	feiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(feiShengInfo.GetFeiLevel())
	if feiTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feisheng:飞升散功失败，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if pl.GetLevel() < sangongLimit {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feisheng:飞升散功失败，等级不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	costExp := int64(0)
	curLevel := pl.GetLevel()
	for reduceNum > 0 {
		levelTemplate := template.GetTemplateService().Get(int(curLevel), (*gametemplate.CharacterLevelTemplate)(nil)).(*gametemplate.CharacterLevelTemplate)
		costExp += int64(levelTemplate.Experience)

		nextTempObj := template.GetTemplateService().Get(int(curLevel+1), (*gametemplate.CharacterLevelTemplate)(nil))
		if nextTempObj == nil {
			nextTempObj = levelTemplate
		}
		nextLevelTemplate := nextTempObj.(*gametemplate.CharacterLevelTemplate)
		curExp := propertyManager.GetExp()
		curPercent := float64(curExp) / float64(nextLevelTemplate.Experience)
		costExp += curExp - int64(math.Ceil(curPercent*float64(levelTemplate.Experience)))

		curLevel -= 1
		reduceNum -= 1
	}

	costLevelReason := commonlog.LevelLogReasonFeiShengSanGong
	propertyManager.CostExp(costExp, costLevelReason, costLevelReason.String())
	propertylogic.SnapChangedProperty(pl)

	//散功
	addGongDe := feiManager.FeiShengSanGong(costExp)

	scMsg := pbutil.BuildSCFeiShengSanGong(addGongDe)
	pl.SendMsg(scMsg)
	return
}
