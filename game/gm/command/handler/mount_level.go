package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
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
	command.Register(gmcommandtypes.CommandTypeMountLevel, command.CommandHandlerFunc(handleMountLevel))

}

func handleMountLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	mountLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"mountLevel": mountLevel,
				"error":      err,
			}).Warn("gm:处理设置坐骑阶别,mountLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	tempTemplateObject := mount.GetMountService().GetMountNumber(int32(mountLevel))

	//修改等级
	if tempTemplateObject == nil {
		log.WithFields(
			log.Fields{
				"id":         pl.GetId(),
				"mountLevel": mountLevel,
				"error":      err,
			}).Warn("gm:处理设置坐骑阶别,mountLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	manager.GmSetMountAdvanced(int(mountLevel))

	//同步属性
	mountlogic.MountPropertyChanged(pl)

	mountId := manager.GetMountInfo().MountId
	scMountAdvanced := pbutil.BuildSCMountAdavancedFinshed(int32(mountLevel), mountId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scMountAdvanced)
	return
}
