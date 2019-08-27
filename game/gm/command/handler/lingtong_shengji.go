package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLingTongShengJi, command.CommandHandlerFunc(handleLingTongShengJi))

}

func handleLingTongShengJi(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	lingTongIdStr := c.Args[0]
	levelStr := c.Args[1]
	lingTongId, err := strconv.ParseInt(lingTongIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"lingTongId": lingTongId,
				"levelStr":   levelStr,
				"error":      err,
			}).Warn("gm:处理设置灵童升级等级,lingTongId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	level, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"lingTongId": lingTongId,
				"levelStr":   levelStr,
				"error":      err,
			}).Warn("gm:处理设置灵童升级等级,level不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	tempTemplateObject := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(int32(lingTongId))
	//升级等级灵童
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"lingTongId": lingTongId,
				"error":      err,
			}).Warn("gm:处理设置灵童升级等级,lingTongId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	lingTongShengJiTemplate := tempTemplateObject.GetLingTongShengJiByLevel(int32(level))
	if lingTongShengJiTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"lingTongId": lingTongId,
				"error":      err,
			}).Warn("gm:处理设置灵童升级等级,lingTongId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongInfo, flag := manager.GetLingTongInfo(int32(lingTongId))
	if !flag {
		flag = manager.GmLingTongActivate(int32(lingTongId))
		if !flag {
			return
		}

		lingTongFashionInfo := manager.GetLingTongFashionById(int32(lingTongId))
		if lingTongFashionInfo == nil {
			panic(fmt.Errorf("lingtong:GetLingTongFashionById should be ok"))
		}
		lingtonglogic.LingTongPropertyChanged(pl)

		lingTongInfo, flag := manager.GetLingTongInfo(int32(lingTongId))
		if !flag {
			return
		}
		fashionId := lingTongFashionInfo.GetFashionId()
		scLingTongActivate := pbutil.BuildSCLingTongActivate(fashionId, lingTongInfo)
		pl.SendMsg(scLingTongActivate)
	}
	flag = manager.GmLingTongShengJi(int32(lingTongId), int32(level))
	if !flag {
		return
	}
	lingtonglogic.LingTongPropertyChanged(pl)

	lingTongInfo, _ = manager.GetLingTongInfo(int32(lingTongId))
	scLingTongUpgrade := pbutil.BuildSCLingTongUpgrade(int32(lingTongId), lingTongInfo.GetLevel(), lingTongInfo.GetPro(), true)
	pl.SendMsg(scLingTongUpgrade)
	return
}
