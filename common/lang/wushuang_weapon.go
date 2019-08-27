package lang

const (
	WushuangWeaponExperienceFull LangCode = WushuangBase + iota
	WushuangWeaponStrengthenItemTypeWrong
	WushuangWeaponExperienceNotFull
	WushuangWeaponLevelFull
	WushuangWeaponBreakthroughItemTypeWrong
	WushuangWeaponBreakthroughItemZhuanshuWrong
	WushuangWeaponBreakthroughTopLevel
	WushuangWeaponStrengthenItemNotEnough
	WushuangWeaponBodyPosEquipmentNotExist
	WushuangWeaponPowerUp
	WushuangWeaponItemNotExist
	WushuangWeaponBuchangEmailContent
	WushuangWeaponBuchangEmailTitle
)

var (
	wushuangWeaponLangMap = map[LangCode]string{
		WushuangWeaponExperienceFull:                "吞噬经验已满",
		WushuangWeaponStrengthenItemTypeWrong:       "吞噬物品类型错误",
		WushuangWeaponExperienceNotFull:             "突破经验值未满",
		WushuangWeaponLevelFull:                     "无双神器满级了",
		WushuangWeaponBreakthroughItemTypeWrong:     "突破物品类型错误",
		WushuangWeaponBreakthroughItemZhuanshuWrong: "突破物品转数不足够",
		WushuangWeaponBreakthroughTopLevel:          "恭喜%s将%s强化至顶级，成功激活独有外观，%s!",
		WushuangWeaponStrengthenItemNotEnough:       "吞噬物品数量不足",
		WushuangWeaponBodyPosEquipmentNotExist:      "部位装备不存在",
		WushuangWeaponPowerUp:                       "战力飙升%d",
		WushuangWeaponItemNotExist:                  "物品ID不存在",
		WushuangWeaponBuchangEmailTitle:             "神甲活动补偿",
		WushuangWeaponBuchangEmailContent:           "由于部分老服玩家无法获得无双神甲，因此本次系统额外开启一场神甲认主活动，由于您当前已有此神甲，系统给予您以下补偿，敬请查收",
	}
)

func init() {
	mergeLang(wushuangWeaponLangMap)
}
