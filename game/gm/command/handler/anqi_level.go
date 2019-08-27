package handler

import (
	"fgame/fgame/common/lang"
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	commontypes "fgame/fgame/game/common/types"
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
	command.Register(gmcommandtypes.CommandTypeAnqiLevel, command.CommandHandlerFunc(handleAnqiLevel))

}

func handleAnqiLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	anqiLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"anqiLevel": anqiLevel,
				"error":     err,
			}).Warn("gm:处理设置暗器阶别,anqiLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := anqitemplate.GetAnqiTemplateService().GetAnqi(int(anqiLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"anqiLevel": anqiLevel,
				"error":     err,
			}).Warn("gm:处理设置暗器阶别,暗器模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	manager.GmSetAnqiAdvanced(int(anqiLevel))

	//同步属性
	anqilogic.AnqiPropertyChanged(pl)

	scAnqiAdvanced := pbutil.BuildSCAnqiAdavancedFinshed(int32(anqiLevel), commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scAnqiAdvanced)
	return
}
