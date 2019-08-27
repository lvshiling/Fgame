package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/jieyi/jieyi"
	jieyilogic "fgame/fgame/game/jieyi/logic"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JIEYI_NAME_UP_LEV_TYPE), dispatch.HandlerFunc(handleJieYiNameUpLev))
}

func handleJieYiNameUpLev(s session.Session, msg interface{}) (err error) {
	log.Debug("jieyi: 开始处理结义威名升级请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = jieYiNameUpLev(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("jieyi: 处理结义威名升级请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("jieyi: 处理结义威名升级请求消息,成功")

	return
}

func jieYiNameUpLev(pl player.Player) (err error) {
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !jieYiManager.IsJieYi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jieyi: 玩家未结义")
		playerlogic.SendSystemMessage(pl, lang.JieYiNotJieYi)
		return
	}

	playerJieYiObj := jieYiManager.GetPlayerJieYiObj()

	level := playerJieYiObj.GetNameLev() + 1
	nameTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiNameTemplate(level)
	if nameTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("jieyi: 威名等级达到上限,无法升级威名")
		playerlogic.SendSystemMessage(pl, lang.JieYiNameAlreadyTopLevel)
		return
	}

	needShengWei := nameTemp.UseShengWei
	if playerJieYiObj.GetShengWeiZhi() < needShengWei {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    level,
			}).Warn("jieyi: 声威值不足")
		playerlogic.SendSystemMessage(pl, lang.JieYiShengWeiZhiNotEnough)
		return
	}

	success := jieyilogic.JieYiNameUpLev(playerJieYiObj.GetNameNum(), nameTemp)

	// 同步数据
	jieYiManager.NameUpLevel(success, needShengWei)

	// 同步属性
	jieyilogic.JieYiPropertyChange(pl)
	obj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	if obj != nil {
		scMsg := pbutil.BuildSCJieBrotherInfoOnChange(obj)
		pl.SendMsg(scMsg)
	}

	if !success {
		level--
	}

	scMsg := pbutil.BuildSCJieYiNameUpLev(success, level, 0, 0)
	pl.SendMsg(scMsg)

	return
}
