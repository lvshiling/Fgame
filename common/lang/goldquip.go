package lang

const (
	GoldEquipGreathenStrengthenItemUseMax LangCode = GoldEquipBase + iota
	GoldEquipReachStrengthenMax
	GoldEquipStrengthenItemNotSame
	GoldEquipStrengthenNotAllow
	GoldEquipStrengthenQualityNotEnough
	GoldEquipChongzhuItemNotEnough
	GoldEquipChongzhuQualityNotEnough
	GoldEquipNotAllowBeItem
	GoldEquipZhuanShengReachLimit
	GoldEquipZhuanShengNeedZhuanShuNoEnough
	GoldEquipZhuanShengLevelNoEnough
	GoldEquipZhuanShengGoldEquipNoEnough
	GoldEquipStrengthenNotice
	GoldEquipCanNotEatBind
	GoldEquipLevelToHigh
	GoldEquipSlotHadEquip
	GoldEquipOpenLightNotAllow
	GoldEquipOpenLightFullLevel
	GoldEquipUpstarNotAllow
	GoldEquipReachUpstarFullLevel
	GoldEquipEquipGemCanNotTakeOff
	GoldEquipEquipGemAlreadyTakeOff
	GoldEquipEquipGemCanNotPutOn
	GoldEquipEquipGemSlotNoExist
	GoldEquipEquipmentSlotNoEquip
	GoldEquipEquipmentNotItemEat
	GoldEquipEquipmentNotGoldEquip
	GoldEquipCantNotUseExtend
	GoldEquipAlreadyUnlock
	GoldEquipNoUnlock
	GoldEquipAutoFenJieMailTitle
	GoldEquipAutoFenJieSlotNotEnoughMailContent
	GoldEquipUseItemStrengthenBuWei
	GoldEquipGodCastingLevelNotEnough
	GoldEquipGodCastingCastingSpiritLevelFull
	GoldEquipGodCastingForgeSoulLevelFull
	GoldEquipGodCastingLevelFull
	GoldEquipGodCastingInheritLevelTooLow
	GoldEquipGodCastingInheritMailTitle
	GoldEquipGodCastingInheritMailContent
)

var (
	goldEquipLangMap = map[LangCode]string{
		GoldEquipGreathenStrengthenItemUseMax:       "超出升星允许材料数量上限",
		GoldEquipReachStrengthenMax:                 "这件装备已经升星到极限，无法继续升星",
		GoldEquipStrengthenItemNotSame:              "请选择相同的装备进行升星",
		GoldEquipStrengthenNotAllow:                 "该金装无法被升星",
		GoldEquipStrengthenQualityNotEnough:         "仅橙色装备才可进行升星",
		GoldEquipChongzhuItemNotEnough:              "材料不足，无法进行重铸",
		GoldEquipChongzhuQualityNotEnough:           "非橙色装备无法进行重铸",
		GoldEquipNotAllowBeItem:                     "不允许作为材料",
		GoldEquipZhuanShengReachLimit:               "您当前已达最高转数,无法再转生",
		GoldEquipZhuanShengNeedZhuanShuNoEnough:     "转生需要的玩家转数不足",
		GoldEquipZhuanShengLevelNoEnough:            "您当前等级不足,无法转生",
		GoldEquipZhuanShengGoldEquipNoEnough:        "元神装备不足,无法转生",
		GoldEquipStrengthenNotice:                   "上天眷顾，%s成功将%s升星到%s",
		GoldEquipCanNotEatBind:                      "绑定装备无法进行吞噬",
		GoldEquipLevelToHigh:                        "元神金装升星等级过高",
		GoldEquipSlotHadEquip:                       "无法使用该物品，请先卸下元神金装",
		GoldEquipOpenLightNotAllow:                  "仅橙色品质元神装备可以进行开光",
		GoldEquipOpenLightFullLevel:                 "装备开光等级已达上限，无法继续开光",
		GoldEquipUpstarNotAllow:                     "仅蓝色以及之上品质元神装备可以进行强化",
		GoldEquipReachUpstarFullLevel:               "当前装备强化已达最高，无法继续强化",
		GoldEquipEquipGemCanNotTakeOff:              "宝石不能被卸下",
		GoldEquipEquipGemAlreadyTakeOff:             "宝石已经被卸下",
		GoldEquipEquipGemCanNotPutOn:                "宝石不能镶嵌",
		GoldEquipEquipGemSlotNoExist:                "宝石槽位不存在",
		GoldEquipEquipmentSlotNoEquip:               "装备槽没有装上装备",
		GoldEquipEquipmentNotItemEat:                "请选择装备分解",
		GoldEquipEquipmentNotGoldEquip:              "目标不是元神金装",
		GoldEquipCantNotUseExtend:                   "该装备升星等级低于当前装备，无法进行继承",
		GoldEquipAlreadyUnlock:                      "已经解锁过",
		GoldEquipNoUnlock:                           "未解锁",
		GoldEquipAutoFenJieMailTitle:                "自动分解",
		GoldEquipAutoFenJieSlotNotEnoughMailContent: "由于您设置了自动分解，但是背包空间不足，分解获得的东西获得的东西现以邮件返还给您，请查收！",
		GoldEquipUseItemStrengthenBuWei:             "成功将%s部位强化至%s级",
		GoldEquipGodCastingLevelNotEnough:           "神铸等级不足",
		GoldEquipGodCastingCastingSpiritLevelFull:   "神铸铸灵满级",
		GoldEquipGodCastingForgeSoulLevelFull:       "神铸锻魂满级",
		GoldEquipGodCastingLevelFull:                "神铸装备满级",
		GoldEquipGodCastingInheritLevelTooLow:       "继承装备神铸等级小于等于身上装备神铸等级",
		GoldEquipGodCastingInheritMailTitle:         "神铸继承",
		GoldEquipGodCastingInheritMailContent:       "恭喜您继承了神铸次数，原材料装备销毁，原装备升星等级：%s，开光次数：%s，现将升星和开光的材料消耗返给您，请查收！",
	}
)

func init() {
	mergeLang(goldEquipLangMap)
}
