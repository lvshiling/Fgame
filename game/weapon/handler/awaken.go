package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
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
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WEAPON_AWAKEN_TYPE), dispatch.HandlerFunc(handleWeaponAwaken))
}

//处理兵魂觉醒信息
func handleWeaponAwaken(s session.Session, msg interface{}) (err error) {
	log.Debug("weapon:处理兵魂觉醒信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWeaponAwaken := msg.(*uipb.CSWeaponAwaken)
	weaponId := csWeaponAwaken.GetWeaponId()

	err = weaponAwanken(tpl, weaponId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("weapon:处理兵魂觉醒信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("weapon:处理兵魂觉醒信息完成")
	return nil

}

//兵魂觉醒逻辑
func weaponAwanken(pl player.Player, weaponId int32) (err error) {
	weaponManager := pl.GetPlayerDataManager(types.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	flag := weaponManager.IfIsAwaken(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = weaponManager.IfWeaponExist(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:未激活的兵魂,无法觉醒")
		playerlogic.SendSystemMessage(pl, lang.WeaponNotActiveNotAwaken)
		return
	}

	state := weaponManager.GetWeaponState(weaponId)
	if state == weapontypes.WeaponAwakenStatusTypeOk {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:兵魂重复觉醒")
		playerlogic.SendSystemMessage(pl, lang.WeaponRepeatAwaken)
		return
	}

	flag = weaponManager.IfAwakenStar(weaponId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"weaponId": weaponId,
		}).Warn("weapon:道具不足,无法觉醒")
		playerlogic.SendSystemMessage(pl, lang.WeaponAwakenNotStar)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	weaponTemplate := weapon.GetWeaponService().GetWeaponTemplate(int(weaponId))
	items := weaponTemplate.GetAwakenItemMap()
	//物品判断
	if len(items) != 0 {
		flag = inventoryManager.HasEnoughItems(items)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"weaponId": weaponId,
			}).Warn("weapon:道具不足,无法觉醒")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		reasonText := commonlog.InventoryLogReasonWeaponAwaken.String()
		flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonWeaponAwaken, reasonText)
		if !flag {
			panic(fmt.Errorf("weapon: weaponAwanken use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	sucess := mathutils.RandomHit(common.MAX_RATE, int(weaponTemplate.AwakenSuccessRate))
	if sucess {
		flag = weaponManager.Awaken(weaponId)
		if !flag {
			panic(fmt.Errorf("weapon: Awaken  should be ok"))
		}
		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
	}

	scWeaponAwaken := pbutil.BuildSCWeaponAwaken(weaponId, sucess)
	pl.SendMsg(scWeaponAwaken)
	return
}
