package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/arena/arena"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_WIN_TYPE), dispatch.HandlerFunc(handleArenaWin))
}

//处理获胜
func handleArenaWin(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3匹配获胜")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isArenaWin := msg.(*crosspb.ISArenaWin)
	level := isArenaWin.GetLevel()
	extra := isArenaWin.GetExtra()
	err = arenaWin(tpl, level, extra)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3匹配获胜,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3匹配获胜,完成")
	return nil

}

//3v3匹配
func arenaWin(pl player.Player, level int32, win bool) (err error) {

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

	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaObj := arenaManager.GetPlayerArenaObjectByRefresh()
	if win && arenaObj.IfExtralWinRew() {
		nextDayWinCount := arenaObj.GetDayWinCount() + 1
		arenaTemplate := arenatemplate.GetArenaTemplateService().GetArenaTemplate(arenatypes.ArenaTypeArena, nextDayWinCount)
		if arenaTemplate != nil {
			newExtraItemMap := make(map[int32]int32)
			extraItemMap := arenaTemplate.GetExtraItemMap()
			newExtraItemMap = utils.MergeMap(newExtraItemMap, extraItemMap)
			newExtraItemMap = utils.MultMap(extraItemMap, ratio)
			if len(newExtraItemMap) > 0 {
				inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
				if inventoryManager.HasEnoughSlots(newExtraItemMap) {
					reason := commonlog.InventoryLogReasonArenaExtra
					reasonText := fmt.Sprintf(reason.String(), level)
					if !inventoryManager.BatchAdd(newExtraItemMap, reason, reasonText) {
						panic(fmt.Errorf("arena:处理跨服3v3匹配获胜,应该成功"))
					}
					inventorylogic.SnapInventoryChanged(pl)
				} else {
					title := lang.GetLangService().ReadLang(lang.ArenaMailTitle)
					content := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
					emaillogic.AddEmail(pl, title, content, newExtraItemMap)
				}
			}
		}
	}

	//竞技场获胜
	maxJifenPercent, addJifenPercent := xianzuncardlogic.ArenaExtralPercent(pl)
	arenaManager.ArenaFinish(win, ratio, maxJifenPercent, addJifenPercent)
	arena.GetArenaService().UpdateWinCount(win, pl.GetId(), pl.GetName())
	gameevent.Emit(arenaeventtypes.EventTypeArenaLianSheng, pl, arenaObj.GetWinCount())

	scMsg := pbutil.BuildSCPlayerArenaInfo(arenaObj)
	pl.SendMsg(scMsg)

	siArenaWin := pbutil.BuildSIArenaWin()
	pl.SendMsg(siArenaWin)
	return
}
