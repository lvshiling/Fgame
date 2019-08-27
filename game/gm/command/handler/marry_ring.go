package handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/gm/command"
	gmcommandtypes "fgame/fgame/game/gm/command/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	command.Register(gmcommandtypes.CommandTypeMarryRing, command.CommandHandlerFunc(handleMarryRing))

}

func handleMarryRing(p scene.Player, c *command.Command) (err error) {
	pl := p.(player.Player)
	if len(c.Args) <= 0 {
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}
	levelStr := c.Args[0]
	ringLevel, err := strconv.ParseInt(levelStr, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"ringLevel": ringLevel,
				"error":     err,
			}).Warn("gm:处理设置婚戒等级,ringLevel不是数字")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		err = nil
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	marryInfo := manager.GetMarryInfo()
	spouseId := marryInfo.SpouseId

	//未婚
	if marryInfo.Status == marrytypes.MarryStatusTypeUnmarried {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("marry:未结婚,无法培养")
		playerlogic.SendSystemMessage(pl, lang.MarryTreeFeedNoMarried)
		return
	}

	ringTemplate := marrytemplate.GetMarryTemplateService().GetMarryRingTemplate(marryInfo.Ring, int32(ringLevel))
	if ringTemplate == nil {
		log.WithFields(
			log.Fields{
				"id":        pl.GetId(),
				"ringLevel": ringLevel,
				"error":     err,
			}).Warn("gm:处理设置婚戒等级,ringLevel模板不存在")
		playerlogic.SendSystemMessage(pl, lang.GMFormatWrong)
		return
	}

	//获取配偶婚戒等级
	sringLevel := int32(0)
	marryObj := marry.GetMarryService().GetMarry(pl.GetId())
	if marryObj != nil {
		if marryObj.PlayerId == pl.GetId() {
			sringLevel = marryObj.PlayerRingLevel
		} else {
			sringLevel = marryObj.SpouseRingLevel
		}
		if spouseId != 0 {
			//等级差
			levelDiff := marrytemplate.GetMarryTemplateService().GetMarryConstRingLevelDiff()
			diffLevel := int32(ringLevel) - sringLevel
			if diffLevel >= levelDiff || diffLevel <= -levelDiff {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
				}).Warn("marry:夫妻戒指等级差不满足,无法提升")
				playerlogic.SendSystemMessage(pl, lang.MarryRingFeedLevelNoEnough)
				return
			}
		}
	}

	manager.GmMarryRingLevel(int32(ringLevel))

	//同步属性
	marrylogic.MarryPropertyChanged(pl)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	//婚戒等级同步给配偶
	if spl != nil {
		scMarryRLevelChange := pbuitl.BuildSCMarryRLevelChange(pl.GetId(), marryInfo.RingLevel)
		spl.SendMsg(scMarryRLevelChange)
	}

	scMarryRingFeed := pbuitl.BuildSCMarryRingFeed(marryInfo.RingLevel, marryInfo.RingExp)
	pl.SendMsg(scMarryRingFeed)
	return
}
