package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/arena/arena"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_GIVE_UP_TYPE), dispatch.HandlerFunc(handleArenaGiveUp))
}

//处理放弃/退出
func handleArenaGiveUp(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3匹配放弃/退出")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaGiveUp(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3匹配放弃/退出,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3匹配放弃/退出,完成")
	return nil

}

//3v3匹配
func arenaGiveUp(pl player.Player) (err error) {

	//活动倍数
	activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeArena)
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemp, _ := activityTemp.GetActivityTimeTemplate(now, openTime, mergeTime)
	ratio := int32(1)
	if timeTemp != nil {
		ratio = int32(timeTemp.BeiShu)
	}
	maxJifenPercent, addJifenPercent := xianzuncardlogic.ArenaExtralPercent(pl)

	//竞技场放弃/退出
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaManager.ArenaGiveUp(ratio, maxJifenPercent, addJifenPercent)
	arena.GetArenaService().UpdateWinCount(false, pl.GetId(), pl.GetName())

	//失败邮件
	title := lang.GetLangService().ReadLang(lang.ArenaMailTitleFail)
	content := lang.GetLangService().ReadLang(lang.ArenaMailContentFail)
	emaillogic.AddEmail(pl, title, content, nil)

	arenaObj := arenaManager.GetPlayerArenaObjectByRefresh()
	scMsg := pbutil.BuildSCPlayerArenaInfo(arenaObj)
	pl.SendMsg(scMsg)

	siMsg := pbutil.BuildSIArenaGiveUp()
	pl.SendMsg(siMsg)
	return
}
