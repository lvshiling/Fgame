package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	feishengtemplate "fgame/fgame/game/feisheng/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_DU_JIE_TYPE), dispatch.HandlerFunc(handleFeiShengDuJie))
}

//处理飞升渡劫
func handleFeiShengDuJie(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理飞升渡劫消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = feiShengDuJie(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("feisheng:处理飞升渡劫消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("feisheng:处理飞升渡劫消息完成")
	return nil

}

//飞升渡劫界面逻辑
func feiShengDuJie(pl player.Player) (err error) {

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengInfo := feiManager.GetFeiShengInfo()
	nextFeiTemplate := feishengtemplate.GetFeiShengTemplateService().GetFeiShengTemplate(feiShengInfo.GetFeiLevel() + 1)
	if nextFeiTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feisheng:飞升渡劫失败，模板不存在，已达满级")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//功德数
	if feiShengInfo.GetGongDeNum() < int64(nextFeiTemplate.GongDe) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("feisheng:飞升渡劫失败，功德不足")
		playerlogic.SendSystemMessage(pl, lang.FeiShengGongDeNotEnough)
		return
	}

	curRate := feiShengInfo.GetAddRate() + nextFeiTemplate.Rate
	isSuccess := mathutils.RandomHit(common.MAX_RATE, int(curRate))

	//上buff
	buffId := int32(0)
	if isSuccess {
		buffId = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFeiShengSuccessBuffId)
	} else {
		buffId = constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFeiShengFaildBuffId)
	}
	scenelogic.AddBuff(pl, buffId, scenetypes.FeiShengKillerId, common.MAX_RATE)
	feiManager.FeiShengDuJie(isSuccess)

	scMsg := pbutil.BuildSCFeiShengDuJie()
	pl.SendMsg(scMsg)
	return
}
