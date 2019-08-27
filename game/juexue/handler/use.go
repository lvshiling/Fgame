package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/juexue/juexue"
	"fgame/fgame/game/juexue/pbutil"
	playerjuexue "fgame/fgame/game/juexue/player"
	jxtypes "fgame/fgame/game/juexue/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_JUEXUE_USE_TYPE), dispatch.HandlerFunc(handleJueXueUse))
}

//处理绝学使用信息
func handleJueXueUse(s session.Session, msg interface{}) (err error) {
	log.Debug("juexue:处理绝学使用信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csJueXueUse := msg.(*uipb.CSJueXueUse)
	typ := csJueXueUse.GetTyp()

	err = juexueUse(tpl, jxtypes.JueXueType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("juexue:处理绝学使用信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Debug("juexue:处理绝学使用信息完成")
	return nil
}

//处理绝学使用信息逻辑
func juexueUse(pl player.Player, typ jxtypes.JueXueType) (err error) {
	jxManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjuexue.PlayerJueXueDataManager)
	if !typ.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := jxManager.IfJueXueExist(typ)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:该绝学还未激活,无法使用")
		playerlogic.SendSystemMessage(pl, lang.JueXueUseNotActive)
		return
	}

	flag = jxManager.IfUseExist(typ)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("juexue:绝学正在使用")
		playerlogic.SendSystemMessage(pl, lang.JueXueUseRepeat)
		return
	}

	flag = jxManager.JueXueUse(typ)
	if !flag {
		panic(fmt.Errorf("juexue: juexueUse should be ok"))
	}
	insight, level := jxManager.GetJueXueLevelByTyp(typ)
	jueXueTemplate := juexue.GetJueXueService().GetJueXueByTypeAndLevel(typ, insight, level)
	juexueId := int32(jueXueTemplate.TemplateId())
	scJueXueUse := pbutil.BuildSCJueXueUse(juexueId)
	pl.SendMsg(scJueXueUse)
	return
}
