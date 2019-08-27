package lang

const (
	InventorySlotNoEnough LangCode = InventoryBase + iota
	InventoryDepotSlotNoEnough
	InventoryCanNotSell
	InventoryCanNotAddSlot
	InventoryCanNotAddDepotSlot
	InventoryEquipCanNotTakeOff
	InventoryItemNoExist
	InventoryItemCanNotEquip
	InventoryEquipmentSlotNoEquip
	InventoryEquipmentSlotHadEquip
	InventoryEquipmentSlotStarMax
	InventoryEquipmentSlotLevelMax
	InventoryEquipmentUpgradeMax
	InventoryEquipmentSlotStrengthNoSlot
	InventoryItemNoEnough
	InventoryItemCanNotUse
	InventoryEquipGemCanNotTakeOff
	InventoryEquipGemAlreadyTakeOff
	InventoryEquipGemCanNotPutOn
	InventoryEquipGemSlotNoExist
	InventoryDepotNotAllowStore
	InventoryTodayUseTimesNotEnough
	InventoryTotalUseTimesNotEnough
	InventoryUseItemInCd
	InventoryItemExpire
	InventoryEquipmentSlotLevelExceedLevel
	InventoryQianKunDaiUseOut
	InventoryChessItemNotice
	InventoryBombOreItemNotice
	InventoryLuckyTrayItemNotice
	InventoryLuckyChargeTrayItemNotice
	InventorySecretCardItemNotice
	InventoryInCross
	InventoryTeShuItemNotEnough
	InventoryItemNotReachEffectiveTime
	InventoryClearEquipmentGemReurnMailTitle
	InventoryClearEquipmentGemReurnMailContent
	InventoryEquipBaoKuItemNotice
	InventoryEquipBaoKuLuckyBoxItemNotice
	InventorySlotNoEnoughSlot
	InventoryMiBaoDepotSlotNoEnough
	InventoryItemCanNotTrade
	InventoryItemCanNotChaiJie
	InventoryUseExpendSlotCardSucceed
	InventoryMaterialBaoKuItemNotice
	InventoryMaterialBaoKuLuckyBoxItemNotice
)

var (
	inventoryLangMap = map[LangCode]string{
		InventorySlotNoEnough:          "背包空间不足,请整理后再来",
		InventoryDepotSlotNoEnough:     "仓库已满，无法放入",
		InventoryCanNotSell:            "不能出售",
		InventoryCanNotAddSlot:         "不能添加槽位",
		InventoryCanNotAddDepotSlot:    "不能添加仓库槽位",
		InventoryEquipCanNotTakeOff:    "装备不能被卸下",
		InventoryItemNoExist:           "物品不存在",
		InventoryItemCanNotEquip:       "物品不能装备",
		InventoryEquipmentSlotNoEquip:  "装备槽没有装上装备",
		InventoryEquipmentSlotHadEquip: "无法使用该物品，请先卸下装备",
		InventoryEquipmentSlotStarMax:  "装备槽已经升到最高星",

		InventoryEquipmentSlotLevelMax:           "装备槽已经强化到最高级",
		InventoryEquipmentUpgradeMax:             "装备槽已经升到最高阶",
		InventoryEquipmentSlotStrengthNoSlot:     "装备槽没有可以强化升级的东西",
		InventoryItemNoEnough:                    "物品不足",
		InventoryItemCanNotUse:                   "物品不能使用",
		InventoryEquipGemCanNotTakeOff:           "宝石不能被卸下",
		InventoryEquipGemAlreadyTakeOff:          "宝石已经被卸下",
		InventoryEquipGemCanNotPutOn:             "宝石不能镶嵌",
		InventoryEquipGemSlotNoExist:             "宝石槽位不存在",
		InventoryDepotNotAllowStore:              "不允许放入仓库",
		InventoryTodayUseTimesNotEnough:          "今日使用次数已达上限，请明日再来",
		InventoryTotalUseTimesNotEnough:          "物品可使用次数已达上限，无法使用该物品",
		InventoryUseItemInCd:                     "物品使用CD中",
		InventoryItemExpire:                      "物品已过期",
		InventoryEquipmentSlotLevelExceedLevel:   "装备槽已经强化超过人物等级上限",
		InventoryQianKunDaiUseOut:                "乾坤袋开启完毕，无法再次使用",
		InventoryInCross:                         "跨服中无法使用",
		InventoryClearEquipmentGemReurnMailTitle: "宝石退还",
		InventoryItemNotReachEffectiveTime:       "该物品要在%s之后才可以使用哦 ~",

		InventoryChessItemNotice:                   "%s破局成功，在苍龙棋局中获得%s",
		InventoryBombOreItemNotice:                 "%s往矿洞丢了一捆炸药，轰隆一声，矿中爆出%s",
		InventoryLuckyTrayItemNotice:               "寻龙点穴，%s在天玄罗盘指引下获得%s",
		InventoryLuckyChargeTrayItemNotice:         "寻龙点穴，%s在幸运转盘指引下获得%s",
		InventorySecretCardItemNotice:              "天眷之人，%s完成天机牌，获得%s",
		InventoryTeShuItemNotEnough:                "缺少黄金杀猪刀，黄金杀猪刀可在【充值返利-1000档】中获得",
		InventoryClearEquipmentGemReurnMailContent: "亲爱的玩家，由于系统更新导致装备宝石问题，系统自动将宝石退回给您，请您重新进行镶嵌，给您带来的不变，敬请谅解",
		InventoryEquipBaoKuItemNotice:              "%s鸿运当头，在装备宝库抽取中了%s，一统仙界指日可待！",
		InventoryEquipBaoKuLuckyBoxItemNotice:      "%s鸿运当头，在装备宝库的幸运值抽取中兑换到了%s，一统仙界指日可待！",
		InventorySlotNoEnoughSlot:                  "背包空间不足，请至少保留%s个背包空位",
		InventoryMiBaoDepotSlotNoEnough:            "秘宝仓库空间不足，无法放入",
		InventoryItemCanNotTrade:                   "不能交易",
		InventoryItemCanNotChaiJie:                 "不能拆解",
		InventoryUseExpendSlotCardSucceed:          "成功使用%s个%s级背包扩充符，背包格子数+%s",
		InventoryMaterialBaoKuItemNotice:           "%s鸿运当头，在材料宝库抽取中了%s，一统仙界指日可待！",
		InventoryMaterialBaoKuLuckyBoxItemNotice:   "%s鸿运当头，在材料宝库的幸运值抽取中兑换到了%s，一统仙界指日可待！",
	}
)

func init() {
	mergeLang(inventoryLangMap)
}
