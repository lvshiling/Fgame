package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	minggelogic "fgame/fgame/game/mingge/logic"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetemplate "fgame/fgame/game/mingge/template"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMingPanLevel, command.CommandHandlerFunc(handleMingPanLevel))

}

func handleMingPanLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	mingGeSubTypeStr := c.Args[0]
	numberStr := c.Args[1]
	starStr := c.Args[2]

	mingGeSub, err := strconv.ParseInt(mingGeSubTypeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":               pl.GetId(),
				"mingGeSubTypeStr": mingGeSubTypeStr,
				"numberStr":        numberStr,
				"starStr":          starStr,
				"error":            err,
			}).Warn("gm:处理设置命盘阶别,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	mingGeSubType := minggetypes.MingGeAllSubType(mingGeSub)
	if !mingGeSubType.Valid() {
		log.WithFields(
			log.Fields{
				"id":               pl.GetId(),
				"mingGeSubTypeStr": mingGeSubTypeStr,
				"numberStr":        numberStr,
				"starStr":          starStr,
				"error":            err,
			}).Warn("gm:处理设置命盘阶别,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	numberInt64, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":               pl.GetId(),
				"mingGeSubTypeStr": mingGeSubTypeStr,
				"numberStr":        numberStr,
				"starStr":          starStr,
				"error":            err,
			}).Warn("gm:处理设置命盘阶别,number不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	starInt64, err := strconv.ParseInt(starStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":               pl.GetId(),
				"mingGeSubTypeStr": mingGeSubTypeStr,
				"numberStr":        numberStr,
				"starStr":          starStr,
				"error":            err,
			}).Warn("gm:处理设置命盘阶别,star不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	number := int32(numberInt64)
	star := int32(starInt64)

	mingPanTemplate := minggetemplate.GetMingGeTemplateService().GetMingGeMingPanTemplate(mingGeSubType, number, star)
	if mingPanTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"number": number,
				"error":  err,
			}).Warn("gm:处理设置命盘阶别,number模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	mingGeAllSubTypeMap := make(map[minggetypes.MingGeAllSubType]bool)
	mingGeAllSubTypeMap[mingGeSubType] = true

	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	manager.GmSetMingPanLevel(mingGeSubType, number, star)

	//同步属性
	minggelogic.MingGePropertyChanged(pl)
	mingGePanRefinedMap := manager.GetMingGePanRefinedMap()
	scMingGeRefined := pbutil.BuildSCMingGeRefined(mingGePanRefinedMap, mingGeAllSubTypeMap)
	pl.SendMsg(scMingGeRefined)
	return
}
