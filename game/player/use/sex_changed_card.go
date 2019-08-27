package use

import (
	"fgame/fgame/common/lang"
	playergoldequip "fgame/fgame/game/goldequip/player"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/marry/marry"
	moonlovelogic "fgame/fgame/game/moonlove/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/pbutil"
	playertypes "fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeSexCard, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handlerGenderChanged))
}

// 性别变更卡
func handlerGenderChanged(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if pl.IsCross() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:处理变更性别,玩家处于跨服")
		playerlogic.SendSystemMessage(pl, lang.InventoryInCross)
		return
	}

	// 结婚状态
	marryObj := marry.GetMarryService().GetMarry(pl.GetId())
	if marryObj != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:处理变更性别,您当前处于结婚状态，无法使用该物品")
		playerlogic.SendSystemMessage(pl, lang.MarryDealIsMarried)
		return
	}

	// 月下情缘状态
	s := pl.GetScene()

	if s != nil && s.MapTemplate().GetMapType() == scenetypes.SceneTypeYueXiaQingYuan {
		sd := s.SceneDelegate().(moonlovelogic.MoonloveSceneData)
		if sd.IsCouple(pl.GetId()) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:处理变更性别,您当前处于双人赏月状态，无法使用该物品")
			playerlogic.SendSystemMessage(pl, lang.MoonloveIsCoupleUseItem)
			return
		}
	}

	// 装备
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipBag := inventoryManager.GetEquipmentBag()
	if equipBag.IsHadEquipment() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:处理变更性别,请先卸下装备")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotHadEquip)
		return
	}

	// 元神金装
	goldEquipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldEquipBag := goldEquipManager.GetGoldEquipBag()
	if goldEquipBag.IsHadEquipment() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:处理变更名称,请先卸下元神金装")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipSlotHadEquip)
		return
	}

	newSex := pl.ChangeSex()
	scMsg := pbutil.BuildSCPlayerSexChanged(newSex)
	pl.SendMsg(scMsg)

	flag = true
	return
}
