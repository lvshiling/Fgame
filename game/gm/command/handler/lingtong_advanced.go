package handler

import (
	"fgame/fgame/common/lang"
	commontypes "fgame/fgame/game/common/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtempalte "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeLingTongAdvanced, command.CommandHandlerFunc(handleLingTongAdvanced))

}

func handleLingTongAdvanced(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 1 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	classTypeStr := c.Args[0]
	advancedIdStr := c.Args[1]
	classTypeInt64, err := strconv.ParseInt(classTypeStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":            pl.GetId(),
				"classTypeStr":  classTypeStr,
				"advancedIdStr": advancedIdStr,
				"error":         err,
			}).Warn("gm:处理设置灵童养成类阶别,classTypeInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	advancedInt64, err := strconv.ParseInt(advancedIdStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":            pl.GetId(),
				"classTypeStr":  classTypeStr,
				"advancedIdStr": advancedIdStr,
				"error":         err,
			}).Warn("gm:处理设置灵童养成类阶别,advancedInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}
	classType := lingtongdevtypes.LingTongDevSysType(classTypeInt64)
	if !classType.Vaild() {
		log.WithFields(
			log.Fields{
				"id":            pl.GetId(),
				"classTypeStr":  classTypeStr,
				"advancedIdStr": advancedIdStr,
				"error":         err,
			}).Warn("gm:处理设置灵童养成类阶别,advancedInt64不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	return lingTongDevAdvanced(pl, advancedInt64, classType)
}

func lingTongDevAdvanced(pl player.Player, advancedInt64 int64, classType lingtongdevtypes.LingTongDevSysType) (err error) {
	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := lingTongManager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId":      pl.GetId(),
			"classType":     classType,
			"advancedInt64": advancedInt64,
		}).Warn("gm:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	lingTongDevTemplate := lingtongdevtempalte.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, int32(advancedInt64))
	if lingTongDevTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":            pl.GetId(),
				"classType":     classType,
				"advancedInt64": advancedInt64,
				"error":         err,
			}).Warn("gm:处理设置灵童养成类阶别,阶别模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	// if lingTongDevInfo == nil {
	// 	lingTongDevInfo = manager.AdvancedInit(classType)
	// }
	manager.GmSetLingTongDevAdvanced(classType, int(advancedInt64))

	//同步属性
	lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)
	advancedId := lingTongDevInfo.GetAdvancedId()
	seqId := lingTongDevInfo.GetSeqId()
	scLingTongDevAdavancedFinshed := pbutil.BuildSCLingTongDevAdavancedFinshed(int32(classType), advancedId, seqId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scLingTongDevAdavancedFinshed)
	return
}
