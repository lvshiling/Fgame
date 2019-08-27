package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMarryTree, command.CommandHandlerFunc(handleMarryTree))

}

func handleMarryTree(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	treeLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"treeLevel": treeLevel,
				"error":     err,
			}).Warn("gm:处理设置爱情树等级,treeLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()

	treeTemplate := marrytemplate.GetMarryTemplateService().GetMarryLoveTreeTemplate(int32(treeLevel))
	if treeTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"treeLevel": treeLevel,
				"error":     err,
			}).Warn("gm:处理设置爱情树等级,treeLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager.GmMarryTreeLevel(int32(treeLevel))

	//同步属性
	marrylogic.MarryPropertyChanged(pl)
	spl := player.GetOnlinePlayerManager().GetPlayerById(marryInfo.SpouseId)
	//爱情树等级同步给配偶
	if spl != nil {
		scMarryTLevelChange := pbuitl.BuildSCMarryTLevelChange(pl.GetId(), marryInfo.TreeLevel)
		spl.SendMsg(scMarryTLevelChange)
	}

	scMarryTreeFeed := pbuitl.BuildSCMarryTreeFeed(marryInfo.TreeLevel, marryInfo.TreeExp)
	pl.SendMsg(scMarryTreeFeed)
	return
}
