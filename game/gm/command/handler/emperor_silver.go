package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	playeremperor "fgame/fgame/game/emperor/player"
	emperortemplate "fgame/fgame/game/emperor/template"
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
	command.Register(gmcommandtypes.CommandTypeEmperorSilver, command.CommandHandlerFunc(handleEmperorSilver))
}

//抢龙椅银两库存
func handleEmperorSilver(p scene.Player, c *command.Command) (err error) {
	log.Debug("gm:silver")
	pl := p.(player.Player)

	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	silverStr := c.Args[0]
	silver, err := strconv.ParseInt(silverStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"silver": silver,
				"error":  err,
			}).Warn("gm:silver,silver不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	maxSilver := emperortemplate.GetEmperorTemplateService().GetEmperorChestMax()

	if silver < 0 || silver > int64(maxSilver) {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"silver": silver,
				"error":  err,
			}).Warn("gm:silver,silver不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	err = emperorSilver(pl, silver)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:处理抢龙椅银两库存,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:处理抢龙椅银两库存完成")
	return
}

func emperorSilver(pl player.Player, silver int64) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerEmperorDataManagerType).(*playeremperor.PlayerEmperorDataManager)
	emperor.GetEmperorService().GmSetEmperorSilver(silver)
	emperorObj := emperor.GetEmperorService().GetEmperorInfo()
	worshipNum := manager.GetWorshipNum()
	scEmperorGet := pbuitl.BuildSCEmperorGet(emperorObj, worshipNum)
	pl.SendMsg(scEmperorGet)
	return
}
