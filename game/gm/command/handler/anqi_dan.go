package handler

import (
	"fgame/fgame/common/lang"
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeAnqiDan, command.CommandHandlerFunc(handleAnqiDan))

}

func handleAnqiDan(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	anqiDanStr := c.Args[0]
	anqiDanLevel, err := strconv.ParseInt(anqiDanStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"anqiDanLevel": anqiDanLevel,
				"error":        err,
			}).Warn("gm:处理设置护暗器丹等级,anqiDanLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := anqitemplate.GetAnqiTemplateService().GetAnqiDan(int32(anqiDanLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":           pl.GetId(),
				"anqiDanLevel": anqiDanLevel,
				"error":        err,
			}).Warn("gm:处理设置护暗器丹等级,anqiDanLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	manager.GmSetAnqiAnqiDanLevel(int32(anqiDanLevel))

	//同步属性
	anqilogic.AnqiPropertyChanged(pl)

	scBodyShieldJJDan := pbutil.BuildSCAnqiEatDan(int32(anqiDanLevel), 0)
	pl.SendMsg(scBodyShieldJJDan)
	return
}
