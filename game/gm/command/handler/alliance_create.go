package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {

	command.Register(gmcommandtypes.CommandTypeAllianceCreate, command.CommandHandlerFunc(handleAllianceCreate))
}

func handleAllianceCreate(pl scene.Player, c *command.Command) (err error) {
	log.Debug("gm:创建仙盟")
	if len(c.Args) <= 2 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	name := c.Args[0]
	versionStr := c.Args[1]
	typStr := c.Args[2]

	versionInt, err := strconv.ParseInt(versionStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:创建仙盟,参数错误")
		return
	}
	version := alliancetypes.AllianceVersionType(versionInt)
	if !version.Valid() {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:创建仙盟,参数错误")
		return
	}

	typInt, err := strconv.ParseInt(typStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:创建仙盟,参数错误")
		return
	}
	typ := alliancetypes.AllianceNewType(typInt)
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Warn("gm:创建仙盟,参数错误")
		return
	}

	err = allianceCreate(pl, name, version, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":    pl.GetId(),
				"error": err,
			}).Error("gm:创建仙盟,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"id": pl.GetId(),
		}).Debug("gm:创建仙盟,完成")
	return
}

func allianceCreate(p scene.Player, name string, version alliancetypes.AllianceVersionType, typ alliancetypes.AllianceNewType) (err error) {
	pl := p.(player.Player)
	lingYuId := pl.GetLingyuInfo().AdvanceId
	// al, err := alliance.GetAllianceService().CreateAlliance(pl.GetCamp(), pl.GetId(), name, pl.GetRole(), pl.GetSex(), pl.GetName(), pl.GetVip(), version, typ)
	al, err := alliance.GetAllianceService().CreateAlliance(pl.GetId(), name, pl.GetRole(), pl.GetSex(), pl.GetName(), pl.GetVip(), lingYuId, pl.GetLevel(), version, typ)
	if err != nil {
		return
	}
	if al == nil {
		panic("alliance:创建仙盟应该成功")
	}
	alliance.GetAllianceService().SyncMemberInfo(pl.GetId(), pl.GetName(), pl.GetSex(), pl.GetLevel(), pl.GetForce(), pl.GetZhuanSheng(), pl.GetLingyuInfo().AdvanceId, pl.GetVip())

	scAllianceCreate := pbutil.BuildSCAllianceCreate(al)
	pl.SendMsg(scAllianceCreate)
	return
}
