package check

import (
	playeranqi "fgame/fgame/game/anqi/player"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	playerdianxing "fgame/fgame/game/dianxing/player"
	playerfabao "fgame/fgame/game/fabao/player"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	playerlingyu "fgame/fgame/game/lingyu/player"
	playermassacre "fgame/fgame/game/massacre/player"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	gametemplate "fgame/fgame/game/template"
	playertianmoti "fgame/fgame/game/tianmo/player"
	playerwing "fgame/fgame/game/wing/player"
	playerxianti "fgame/fgame/game/xianti/player"
	playerxuedun "fgame/fgame/game/xuedun/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeSystemX, quest.CheckHandlerFunc(handleSystemReachX))
}

//check 指定系统达到x级
func handleSystemReachX(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理指系统达到x级")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, _ := range questDemandMap {
		reachXType := questtypes.SystemReachXType(demandId)
		if !reachXType.Valid() {
			return
		}
		order := int32(0)
		switch reachXType {
		case questtypes.SystemReachXTypeMount:
			{
				manager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
				mountInfo := manager.GetMountInfo()
				order = int32(mountInfo.AdvanceId)
				break
			}
		case questtypes.SystemReachXTypeAnQi:
			{
				manager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
				anQiInfo := manager.GetAnqiInfo()
				order = int32(anQiInfo.AdvanceId)
				break
			}
		case questtypes.SystemReachXTypeWing:
			{
				manager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
				wingInfo := manager.GetWingInfo()
				order = int32(wingInfo.AdvanceId)
				break
			}
		case questtypes.SystemReachXTypeBodyShield:
			{
				manager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
				bodyShiedInfo := manager.GetBodyShiedInfo()
				order = int32(bodyShiedInfo.AdvanceId)
				break
			}
		case questtypes.SystemReachXTypeLingYu:
			{
				manager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
				lingyuInfo := manager.GetLingyuInfo()
				order = int32(lingyuInfo.AdvanceId)
				break
			}
		case questtypes.SystemReachXTypeShenFa:
			{
				manager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
				shenfaInfo := manager.GetShenfaInfo()
				order = int32(shenfaInfo.AdvanceId)
				break
			}
		case questtypes.SystemReachXTypeFaBao:
			{
				manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
				fabaoInfo := manager.GetFaBaoInfo()
				order = fabaoInfo.GetAdvancedId()
				break
			}
		case questtypes.SystemReachXTypeXianTi:
			{
				manager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
				order = manager.GetXianTiAdvancedId()
				break
			}
		case questtypes.SystemReachXTypeLuXianRen:
			{
				manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
				massacreInfo := manager.GetMassacreInfo()
				order = int32(massacreInfo.CurrLevel)
				break
			}
		case questtypes.SystemReachXTypeXueDun:
			{
				manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
				xueDunInfo := manager.GetXueDunInfo()
				order = xueDunInfo.GetNumber()
				break
			}
		case questtypes.SystemReachXTypeDianXing:
			{
				manager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
				xueDunInfo := manager.GetDianXingObject()
				order = xueDunInfo.CurrType
				break
			}
		case questtypes.SystemReachXTypeShiHunFan:
			{
				manager := pl.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
				order = manager.GetShiHunFanAdvanced()
				break
			}
		case questtypes.SystemReachXTypeTianMoTi:
			{
				manager := pl.GetPlayerDataManager(types.PlayerTianMoDataManagerType).(*playertianmoti.PlayerTianMoDataManager)
				order = manager.GetTianMoAdvanced()
				break
			}
		case questtypes.SystemReachXTypeLingTongWeapon:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingBing)
				break
			}
		case questtypes.SystemReachXTypeLingTongMount:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingQi)
				break
			}
		case questtypes.SystemReachXTypeLingTongWing:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingYi)
				break
			}
		case questtypes.SystemReachXTypeLingTongShenFa:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingShen)
				break
			}
		case questtypes.SystemReachXTypeLingTongLingYu:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingYu)
				break
			}
		case questtypes.SystemReachXTypeLingTongFaBao:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingBao)
				break
			}
		case questtypes.SystemReachXTypeLingTongXianTi:
			{
				order = lingTongDevAdvaced(pl, lingtongdevtypes.LingTongDevSysTypeLingTi)
				break
			}

		}

		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, order)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理指系统达到x级,完成")
	return nil
}

func lingTongDevAdvaced(pl player.Player, classType lingtongdevtypes.LingTongDevSysType) (advacedId int32) {
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	if lingTongDevInfo == nil {
		return
	}
	return lingTongDevInfo.GetAdvancedId()
}
