package effect

import (
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	babytypes "fgame/fgame/game/baby/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeBaby, babyPropertyEffect)
}

//宝宝属性
func babyPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	babyList := babyManager.GetBabyInfoList()
	allToySlotMap := babyManager.GetAllToySlotMap()

	//宝宝属性
	for _, baby := range babyList {
		attrMap := countBabyProperty(baby.GetQuality(), baby.GetLearnLevel(), baby.GetDanBei(), baby.GetSkillList())
		for typ, val := range attrMap {
			val += prop.GetBase(typ)
			prop.SetBase(typ, val)
		}
	}

	//宝宝玩具
	toyAttrMap := countBabyToyProperty(allToySlotMap)
	for typ, val := range toyAttrMap {
		val += prop.GetBase(typ)
		prop.SetBase(typ, val)
	}

	//配偶宝宝贡献10%属性
	coupleBabyList := babyManager.GetAllCoupleBabyList()
	for _, babyData := range coupleBabyList {
		pregnantTemp := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(babyData.Quality)
		if pregnantTemp == nil {
			continue
		}

		attrMap := countBabyProperty(babyData.Quality, babyData.LearnLevel, babyData.Danbei, babyData.TalentList)
		for typ, val := range attrMap {
			val = int64(math.Ceil(float64(val) * float64(pregnantTemp.AttrShareRate) / float64(common.MAX_RATE)))
			val += prop.GetBase(typ)
			prop.SetBase(typ, val)
		}
	}

	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeBabyTalent, prop)
	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeBabyToySuit, prop)
}

func countBabyToyProperty(allToySlotMap map[babytypes.ToySuitType][]*playerbaby.PlayerBabyToySlotObject) map[propertytypes.BattlePropertyType]int64 {
	attrMap := make(map[propertytypes.BattlePropertyType]int64)
	for suitType, slotList := range allToySlotMap {
		for _, slot := range slotList {
			if slot.IsEmpty() {
				continue
			}
			itemId := int(slot.GetItemId())
			level := slot.GetLevel()
			pos := slot.GetSlotId()
			toyTemp := item.GetItemService().GetItem(itemId).GetBabyToyTemplate()
			if toyTemp == nil {
				continue
			}

			//装备属性
			for typ, val := range toyTemp.GetBattlePropertyMap() {
				attrMap[typ] += val
			}

			//强化属性
			toyUplevelTemp := babytemplate.GetBabyTemplateService().GetBabyToyUplevelTemplate(suitType, pos, level)
			if toyUplevelTemp != nil {
				for typ, val := range toyUplevelTemp.GetBattlePropertyMap() {
					attrMap[typ] += val
				}
			}
		}
	}
	return attrMap
}

func countBabyProperty(quality, learnLevel, danbei int32, talentList []*babytypes.TalentInfo) map[propertytypes.BattlePropertyType]int64 {
	proMap := make(map[propertytypes.BattlePropertyType]int64)
	pregnantTemp := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(quality)
	if pregnantTemp == nil {
		return nil
	}

	//宝宝属性
	babyDanBeiTempMap := babytemplate.GetBabyTemplateService().GetBabyDanBeiTemplateMap()
	for _, danbeiTemp := range babyDanBeiTempMap {
		if pregnantTemp.AttrType&danbeiTemp.GetDanBeiType().Mask() == 0 {
			continue
		}

		typ := danbeiTemp.GetDanBeiType().GetPropertyType()
		val := int64(math.Ceil(float64(danbeiTemp.Attr) * float64(danbei) / float64(common.MAX_RATE)))
		proMap[typ] += val
	}

	//宝宝读书
	babyLearnTemplate := babytemplate.GetBabyTemplateService().GetBabyLearnTemplate(learnLevel)
	if babyLearnTemplate != nil {
		ratio := float64(pregnantTemp.GrowthNum) / float64(common.MAX_RATE)

		for _, danbeiTemp := range babyDanBeiTempMap {
			if babyLearnTemplate.AttrType&danbeiTemp.GetDanBeiType().Mask() == 0 {
				continue
			}

			typ := danbeiTemp.GetDanBeiType().GetPropertyType()
			val := int64(math.Ceil(float64(danbeiTemp.Attr) * float64(babyLearnTemplate.BeiShu) * ratio))
			proMap[typ] += val
		}
	}

	//宝宝天赋
	for _, talent := range talentList {
		if talent.Type != babytypes.SkillTypeAttr {
			continue
		}

		skillTemp := skilltemplate.GetSkillTemplateService().GetSkillTemplate(talent.SkillId)
		if skillTemp == nil {
			continue
		}
		attrMap := skillTemp.GetAttrTemplate().GetAllBattleProperty()
		for typ, val := range attrMap {
			proMap[typ] += val
		}
	}

	return proMap
}
