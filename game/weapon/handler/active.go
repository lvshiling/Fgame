package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
	weapontypes "fgame/fgame/game/weapon/types"
	"fgame/fgame/game/weapon/weapon"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEAPON_ACTIVE_TYPE), dispatch.HandlerFunc(handleWeaponActive))
}

//处理兵魂激活信息
func handleWeaponActive(s session.Session, msg interface{}) (err error) {
	log.Debug("weapon:处理兵魂激活信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWeaponActive := msg.(*uipb.CSWeaponActive)
	weaponId := csWeaponActive.GetWeaponId()

	err = weaponActive(tpl, weaponId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
				"error":    err,
			}).Error("weapon:处理兵魂激活信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Debug("weapon:处理兵魂激活信息完成")
	return nil
}

//处理兵魂激活信息逻辑
func weaponActive(pl player.Player, weaponId int32) (err error) {
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	flag := weaponManager.IsValid(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = weaponManager.IfWeaponExist(weaponId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:该兵魂已激活,无需激活")
		playerlogic.SendSystemMessage(pl, lang.WeaponRepeatActive)
		return
	}

	titleTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	if titleTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
			}).Warn("weapon:兵魂激活失败，该兵魂模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if titleTemplate.GetWeaponTag() != weapontypes.WeaponTagTypePermanent {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
			}).Warn("weapon:兵魂激活失败，该兵魂不能手动激活")
		playerlogic.SendSystemMessage(pl, lang.WeaponActivateFail)
		return
	}

	items := titleTemplate.GetNeedItemMap(pl.GetRole())
	if len(items) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//判断物品是否足够
		flag := inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
			}).Warn("weapon:当前道具不足，无法激活该兵魂")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonWeaponActive
		reasonText := fmt.Sprintf(inventoryReason.String(), weaponId)
		flag = inventoryManager.BatchRemove(items, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("weapon: weaponActive use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = weaponManager.WeaponActive(weaponId, true)
	if !flag {
		panic(fmt.Errorf("weapon: weaponActive should be ok"))
	}

	//同步属性
	weaponlogic.WeaponPropertyChanged(pl)

	scWeaponActive := pbutil.BuildSCWeaponActive(weaponId)
	pl.SendMsg(scWeaponActive)
	return
}
