package lang

const (
	TuLongEquipStrengthenNotAllow LangCode = TuLongEquipBase + iota
	TuLongEquipReachStrengthenMax
	TuLongEquipRongHeQualityNotEnough
	TuLongEquipRongHeItemNotEnough
	TuLongEquipRongHeLevelNotEqual
	TuLongEquipRongHeFailed
	TuLongEquipZhuanHuaItemNotEnough
	TuLongEquipZhuanHuaLevelNotEqual
	TuLongEquipZhuanHuaFailed
	TuLongEquipSkillFailed
	TuLongEquipZhuanHuaPosNotEqual
)

var (
	tulongEquipLangMap = map[LangCode]string{
		TuLongEquipStrengthenNotAllow:     "该屠龙无法被强化",
		TuLongEquipReachStrengthenMax:     "这件屠龙已经强化到极限，无法继续强化",
		TuLongEquipRongHeQualityNotEnough: "非橙色装备无法进行融合",
		TuLongEquipRongHeItemNotEnough:    "材料不足，无法进行融合",
		TuLongEquipRongHeLevelNotEqual:    "请选择相同的装备阶数进行融合",
		TuLongEquipRongHeFailed:           "该装备不可融合",

		TuLongEquipZhuanHuaItemNotEnough: "材料不足，无法进行转化",
		TuLongEquipZhuanHuaLevelNotEqual: "请选择相同的装备阶数进行转化",
		TuLongEquipZhuanHuaFailed:        "该装备不可转化",
		TuLongEquipZhuanHuaPosNotEqual:   "请选择部位一致的装备进行转化",

		TuLongEquipSkillFailed: "激活条件未满足，请继续努力",
	}
)

func init() {
	mergeLang(tulongEquipLangMap)
}
