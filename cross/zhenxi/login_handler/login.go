package login_handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/shareboss/shareboss"
	crosstypes "fgame/fgame/game/cross/types"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeZhenXi, login.LogincHandlerFunc(zhenXiBossLogin))
}

func zhenXiBossLogin(pl *player.Player, ct crosstypes.CrossType, crossArgs ...string) bool {
	if len(crossArgs) <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(crossArgs),
			}).Warn("login:玩家加载数据,参数不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return false
	}
	bossId, err := strconv.ParseInt(crossArgs[0], 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"arg":      crossArgs[0],
			}).Warn("login:玩家加载数据,参数不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return false
	}

	n := shareboss.GetShareBossService().GetShareBoss(worldbosstypes.BossTypeZhenXi, int32(bossId))
	if n == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bossId":   bossId,
			}).Warn("login:玩家加载数据,boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return false
	}
	s := n.GetScene()
	scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos())
	return true
}
