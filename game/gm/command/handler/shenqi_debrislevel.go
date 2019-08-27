package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
	shenqilogic "fgame/fgame/game/shenqi/logic"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitemplate "fgame/fgame/game/shenqi/template"
	shenqitypes "fgame/fgame/game/shenqi/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeShenQiDebrisLevel, command.CommandHandlerFunc(handleSetShenQiDebrisLevel))
}

func handleSetShenQiDebrisLevel(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) < 3 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	typStr := c.Args[0]
	posStr := c.Args[1]
	levStr := c.Args[2]
	typInt, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"typInt": typInt,
				"error":  err,
			}).Warn("gm:处理神器碎片等级任务,类型typInt不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	posInt, err := strconv.ParseInt(posStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":     pl.GetId(),
				"posInt": posInt,
				"error":  err,
			}).Warn("gm:处理神器碎片等级任务,类型posInt不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	lev, err := strconv.ParseInt(levStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"lev":   lev,
				"error": err,
			}).Warn("gm:处理神器碎片等级任务,类型lev不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	//参数验证
	typ := shenqitypes.ShenQiType(int32(typInt))
	if !typ.Valid() {
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("gm:神器类型,错误")
		return
	}
	pos := shenqitypes.DebrisType(int32(posInt))
	if !pos.Valid() {
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("gm:神器类型,错误")
		return
	}

	tempTemplatet := shenqitemplate.GetShenQiTemplateService().GetShenQiDebrisUpByArg(typ, pos, int32(lev))
	if tempTemplatet == nil {
		log.WithFields(
			log.Fields{
				"id":  pl.GetId(),
				"typ": typ.String(),
				"pos": pos.String(),
				"lev": lev,
			}).Warn("gm:处理神器碎片等级任务,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//修改等级
	manager := pl.GetPlayerDataManager(types.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	slotObj := manager.GmSetShenQiDebrisLevel(typ, pos, tempTemplatet.Level)

	//同步属性
	shenqilogic.ShenQiPropertyChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCShenQiDebrisUp(slotObj, false)
	pl.SendMsg(scMsg)
	return
}
