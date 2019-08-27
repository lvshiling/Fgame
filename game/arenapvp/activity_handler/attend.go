package activity_handler

import (
	"fgame/fgame/game/activity/activity"
	activitytypes "fgame/fgame/game/activity/types"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	activity.RegisterActivityHandler(activitytypes.ActivityTypeArenapvp, activity.ActivityAttendHandlerFunc(playerEnterActivity))
}

func playerEnterActivity(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {

	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpInfo := arenapvpManager.GetPlayerArenapvpObj()
	if !arenapvpInfo.IfBuyTicket() {
		constantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
		needBindGold := int64(constantTemp.RuchangUseBindgold)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if !propertyManager.HasEnoughGold(needBindGold, true) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"needBindGold": needBindGold,
				}).Warn("arenapvp:比武大会购买门票,元宝不足")
			return
		}
	}

	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeArenapvp)
	flag = true
	return
}
