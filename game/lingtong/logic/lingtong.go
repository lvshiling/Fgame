package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coretypes "fgame/fgame/core/types"
	commomlogic "fgame/fgame/game/common/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/lingtong/lingtong"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/idutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//检查是否可以激活
func CheckIfLingTongActivate(pl player.Player, lingTongId int32) bool {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		return false
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	_, tflag := manager.GetLingTongInfo(lingTongId)
	if tflag {
		return false
	}

	useItem := lingTongTemplate.UseItemId
	num := lingTongTemplate.UseItemCount
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			return false
		}
	}
	return true
}

//获取灵童激活界面信息逻辑
func HandleLingTongActivate(pl player.Player, lingTongId int32) (err error) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:模板为空")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	_, flag := manager.GetLingTongInfo(lingTongId)
	if flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:重复激活")
		return
	}

	useItem := lingTongTemplate.UseItemId
	num := lingTongTemplate.UseItemCount
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
			}).Warn("lingtong:物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongActivate.String(), lingTongId)
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonLingTongActivate, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	lingTongInfo := manager.LingTongActivate(lingTongId)
	if lingTongInfo == nil {
		panic(fmt.Errorf("lingtong:LingTongActivate should be ok"))
	}
	lingTongFashionInfo := manager.GetLingTongFashionById(lingTongId)
	if lingTongFashionInfo == nil {
		panic(fmt.Errorf("lingtong:GetLingTongFashionById should be ok"))
	}
	LingTongPropertyChanged(pl)

	fashionId := lingTongFashionInfo.GetFashionId()
	scLingTongActivate := pbutil.BuildSCLingTongActivate(fashionId, lingTongInfo)
	pl.SendMsg(scLingTongActivate)
	return
}

//变更灵童属性
func LingTongPropertyChanged(pl player.Player) (err error) {
	//同步属性
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLingTong.Mask())
	LingTongSelfAllPropertyChanged(pl)
	return
}

func GetLingTongPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	basePower := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTong)
	power = basePower
	return
}

func GetLingTongFashionPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	basePower := propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongFashion)
	power = basePower
	return
}

func GetLingTongMountPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongMount)
	return
}

func GetLingTongWingPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongWing)
	return
}

func GetLingTongWeaponPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongWeapon)
	return
}

func GetLingTongShenFaPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongShenFa)
	return
}

func GetLingTongLingYuPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongLingYu)
	return
}

func GetLingTongXianTiPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongXianTi)
	return
}

func GetLingTongFaBaoPower(pl player.Player) (power int64) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	power = propertyManager.GetModuleForce(playerpropertytypes.PlayerPropertyEffectorTypeLingTongFaBao)
	return
}

//灵童养成类皮肤升星判断
func LingTongShengJi(curTimesNum int32, curBless int32, shengJiTemplate *gametemplate.LingTongShengJiTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := shengJiTemplate.TimesMin
	timesMax := shengJiTemplate.TimesMax
	updateRate := shengJiTemplate.UpdateWfb
	blessMax := shengJiTemplate.ZhufuMax
	addMin := shengJiTemplate.AddMin
	addMax := shengJiTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//灵童升星判断
func LingTongUpstarJudge(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.LingTongUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpdateWfb
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//初始化灵童
func InitLingTong(pl player.Player) (flag bool) {
	//更新灵童属性
	LingTongSelfAllPropertyChanged(pl)

	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongDevManager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongObj := lingTongManager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		return
	}
	id, _ := idutil.GetId()
	lingTongId := lingTongObj.GetLingTongId()
	lingTongInfoData, _ := lingTongManager.GetLingTongInfo(lingTongId)
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	fashionId := lingTongManager.GetLingTongFashionById(lingTongId).GetFashionId()
	weaponId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingBing)
	wingId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingYi)
	mountId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingQi)
	shenFaId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingShen)
	lingYuId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingYu)
	faBaoId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingBao)
	xianTiId := lingTongDevManager.GetLingTongDevSeqId(lingtongdevtypes.LingTongDevSysTypeLingTi)

	battleProperties := lingTongManager.GetAllSystemBattleProperties()
	lingTongShowObj := lingtong.CreateLingTongShowObject(fashionId, weaponId, 0, 0, wingId, mountId, false, shenFaId, lingYuId, faBaoId, xianTiId)
	pos := coretypes.Position{}
	angle := float64(0)
	if pl.GetScene() != nil {
		pos = pl.GetPosition()
		angle = pl.GetAngle()
	}
	name := lingTongInfoData.GetLingTongName()
	lingTong := lingtong.CreateLingTong(pl, id, name, pos, angle, lingTongTemplate, lingTongShowObj, battleProperties)

	//设置灵童
	pl.UpdateLingTong(lingTong)
	return true
}

//更新灵童
func UpdateLingTong(pl player.Player) (flag bool) {
	lingTong := pl.GetLingTong()
	if lingTong == nil {
		return false
	}
	LingTongSelfAllPropertyChanged(pl)

	lingTongManager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongId := lingTongManager.GetLingTong().GetLingTongId()
	lingTongInfo, _ := lingTongManager.GetLingTongInfo(lingTongId)
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)

	lingTongFashion := lingTongManager.GetLingTongFashionById(lingTongId)
	lingTong.SetLingTongFashionId(lingTongFashion.GetFashionId())
	battleProperties := lingTongManager.GetChangedBattlePropertiesAndReset()
	lingTong.UpdateSystemBattleProperty(battleProperties)

	lingTong.UpdateLingTongTemplate(lingTongTemplate, lingTongInfo.GetLingTongName())
	return true
}
