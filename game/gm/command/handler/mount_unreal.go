package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	mountlogic "fgame/fgame/game/mount/logic"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMountUnreal, command.CommandHandlerFunc(handleMountUnreal))

}

func handleMountUnreal(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	mountIdStr := c.Args[0]
	mountId, err := strconv.ParseInt(mountIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"mountId": mountId,
				"error":   err,
			}).Warn("gm:处理设置坐骑幻化,mountId不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := mount.GetMountService().GetMount(int(mountId))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":      pl.GetId(),
				"mountId": mountId,
				"error":   err,
			}).Warn("gm:处理设置坐骑幻化,mountId模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	manager.GmSetMountUnreal(int(mountId))

	//同步属性
	mountlogic.MountPropertyChanged(pl)

	scMountUnreal := pbutil.BuildSCMountUnreal(int32(mountId))
	pl.SendMsg(scMountUnreal)
	return
}
