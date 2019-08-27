package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENA_RELIVE_TYPE), dispatch.HandlerFunc(handleArenaRelive))
}

//处理复活
func handleArenaRelive(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理跨服3v3匹配复活")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = arenaRelive(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理跨服3v3匹配复活,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理跨服3v3匹配复活,完成")
	return nil

}

//3v3复活
func arenaRelive(pl player.Player) (err error) {
	culTime := pl.GetArenaReliveTime() + 1
	mi := float64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RevivePar) / float64(common.MAX_RATE)
	first := float64(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().ReviveValue)
	reliveItemId := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().ReviveItem
	//TODO: 防止越界
	//判断消耗
	itemNum := int32(math.Ceil(first * math.Pow(float64(culTime), mi)))
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughItem(reliveItemId, itemNum) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("relive:原地复活,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		siArenaRelive := pbutil.BuildSIArenaRelive(false)
		pl.SendCrossMsg(siArenaRelive)
		return
	}

	reasonText := fmt.Sprintf(commonlog.InventoryLogReasonArenaRelive.String(), culTime)
	flag := inventoryManager.UseItem(reliveItemId, itemNum, commonlog.InventoryLogReasonArenaRelive, reasonText)
	if !flag {
		panic(fmt.Errorf("relive:使用复活道具应该成功"))
	}

	inventorylogic.SnapInventoryChanged(pl)

	//竞技场获胜
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaManager.Relive()
	siArenaRelive := pbutil.BuildSIArenaRelive(true)
	pl.SendCrossMsg(siArenaRelive)
	return
}
