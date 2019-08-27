package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_UNLOAD_TYPE), dispatch.HandlerFunc(handleLingTongFashionUnload))

}

//处理灵童时装卸下信息
func handleLingTongFashionUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理灵童时装卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = lingTongFashionUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("lingtong:处理灵童时装卸下信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingtong:处理灵童时装卸下信息完成")
	return nil

}

//灵童时装卸下逻辑
func lingTongFashionUnload(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingtong:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	lingTongId := lingTongObj.GetLingTongId()
	lingTongFashionObj := manager.GetLingTongFashionById(lingTongId)
	if lingTongFashionObj == nil {
		panic(fmt.Errorf("lingtong: 灵童时装应该是存在的"))
	}
	fashionId := lingTongFashionObj.GetFashionId()
	flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if flag {
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionBornNoUnload)
		return
	}

	fashionWear, flag := manager.Unload()
	if !flag {
		panic(fmt.Errorf("lingtong: 灵童时装卸下应该是ok的"))
	}
	scLingTongFashionUnload := pbutil.BuildSCLingTongFashionUnload(fashionWear)
	pl.SendMsg(scLingTongFashionUnload)
	return
}
