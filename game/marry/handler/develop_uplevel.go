package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	marrylogic "fgame/fgame/game/marry/logic"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MARRY_DEVELOP_UPLEVEL_TYPE), dispatch.HandlerFunc(handleMarryDevelopUplevel))
}

//处理表白升级
func handleMarryDevelopUplevel(s session.Session, msg interface{}) (err error) {
	log.Debug("marry:处理表白升级处理消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = marryDevelopUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("marry:处理表白升级处理消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("marry:处理表白升级处理消息完成")
	return nil
}

//处理表白升级处理信息逻辑
func marryDevelopUp(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	nextLevel := manager.GetMarryDevelopLevel() + 1
	nextDevelopTemp := marrytemplate.GetMarryTemplateService().GetMarryDeveopTemplate(nextLevel)
	if nextDevelopTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"nextLevel": nextLevel,
			}).Warn("marry:您当前等级已达上限，无法继续升级")
		playerlogic.SendSystemMessage(pl, lang.MarryDevelopMaxLevel)
		return
	}

	if manager.GetMarryDevelopExp() < nextDevelopTemp.Experience {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"curExp":   manager.GetMarryDevelopExp(),
				"levelExp": nextDevelopTemp.Experience,
			}).Warn("marry:表白升级，表白经验不足")
		playerlogic.SendSystemMessage(pl, lang.MarryDevelopExpNotEnough)
		return
	}

	flag := manager.MarryDevelopUpLevel()
	if !flag {
		panic(fmt.Errorf("表白升级应该成功"))
	}
	marrylogic.MarryPropertyChanged(pl)

	//通知配偶
	spl := player.GetOnlinePlayerManager().GetPlayerById(manager.GetSpouseId())
	if spl != nil {
		scMarryTLevelChange := pbuitl.BuildSCMarryDevelopLevelChange(pl.GetId(), nextLevel)
		spl.SendMsg(scMarryTLevelChange)
	}

	curExp := manager.GetMarryDevelopExp()
	scMsg := pbuitl.BuildSCMarryDevelopUplevel(nextLevel, curExp)
	pl.SendMsg(scMsg)
	return
}
