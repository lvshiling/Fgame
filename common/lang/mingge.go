package lang

const (
	MingGeMosaicItemNoCorrect LangCode = MingGeBase + iota
	MingGeMosaicItemSlotNoCorrect
	MingGeMosaicItemQualityNoHigher
	MingGeSlotUnloadNoItemId
	MingGeRefinedAllFull
	MingGeMingLiBaptize
	MingGeMingLiSetNum
	MingGeMingLiSetProperty
	MingGeRefundMailTitle
	MingGeRefundMailContent
)

var (
	mingGeLangMap = map[LangCode]string{
		MingGeMosaicItemNoCorrect:       "命盘镶嵌,命格物品不对",
		MingGeMosaicItemSlotNoCorrect:   "命盘镶嵌,槽位不存在",
		MingGeMosaicItemQualityNoHigher: "命盘镶嵌,替换的命格品质不高于当前命格",
		MingGeSlotUnloadNoItemId:        "该槽位没有镶嵌命格,无法卸下",
		MingGeRefinedAllFull:            "命盘祭炼所有命盘都满级了",
		MingGeMingLiBaptize:             "未激活的命宫,无法洗练",
		MingGeMingLiSetNum:              "请先洗练一次,获取对象,才能设置次数",
		MingGeMingLiSetProperty:         "请先洗练一次,获取对象,才能设置属性",
		MingGeRefundMailTitle:           "系统补偿",
		MingGeRefundMailContent:         "亲爱的玩家，由于命宫系统战力有部分计算延迟（仅存在于洗练出三条同属性的玩家），当前版本已经修复了该问题，但可能导致一定幅度的战力波动，故奉上大量补偿，给您带来的不便深表歉意",
	}
)

func init() {
	mergeLang(mingGeLangMap)
}
