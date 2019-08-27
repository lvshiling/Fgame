package lang

const (
	GuaJiNoGuaJiPlayer LangCode = GuaJiBase + iota
	GuaJiProcessing
	GuaJiArgsInvalid
	GuaJiAutoBuyBlood
	GuaJiGoldEquipmentWear
	GuaJiSkillUpgrade
	GuaJiAllianceBatchJoin
	GuaJiBiaoCheRewardGet
	GuaJiSoulActive
	GuaJiMountAdvanced
	GuaJiWingAdvanced
	GuaJiBodyshieldAdvanced
	GuaJiAnqiAdvanced
	GuaJiFabaoAdvanced
	GuaJiShenfaAdvanced
	GuaJiXiantiAdvanced
	GuaJiLingyuAdvanced
	GuaJiShihunfanAdvanced
	GuaJiTianmotiAdvanced
	GuaJiFeatherAdvanced
	GuaJiShieldAdvanced
	GuaJiMassacreAdvanced
	GuaJiLingTongWeaponAdvanced
	GuaJiLingTongMountAdvanced
	GuaJiLingTongWingAdvanced
	GuaJiLingTongShenFaAdvanced
	GuaJiLingTongLingYuAdvanced
	GuaJiLingTongFaBaoAdvanced
	GuaJiLingTongXianTiAdvanced
)

var (
	guaJiLangMap = map[LangCode]string{
		GuaJiNoGuaJiPlayer:          "不是挂机玩家",
		GuaJiProcessing:             "正在挂机中",
		GuaJiArgsInvalid:            "挂机参数错误",
		GuaJiAutoBuyBlood:           "正在购买血药",
		GuaJiGoldEquipmentWear:      "正在替换元神金装",
		GuaJiSkillUpgrade:           "正在升级技能",
		GuaJiAllianceBatchJoin:      "正在仙盟一键加入",
		GuaJiBiaoCheRewardGet:       "正在获取镖车奖励",
		GuaJiSoulActive:             "正在激活帝魂",
		GuaJiMountAdvanced:          "正在坐骑自动进阶",
		GuaJiWingAdvanced:           "正在战翼自动进阶",
		GuaJiBodyshieldAdvanced:     "正在护体盾自动进阶",
		GuaJiAnqiAdvanced:           "正在暗器自动进阶",
		GuaJiFabaoAdvanced:          "正在法宝自动进阶",
		GuaJiShenfaAdvanced:         "正在身法自动进阶",
		GuaJiXiantiAdvanced:         "正在仙体自动进阶",
		GuaJiLingyuAdvanced:         "正在领域自动进阶",
		GuaJiShihunfanAdvanced:      "正在噬魂番自动进阶",
		GuaJiTianmotiAdvanced:       "正在天魔体自动进阶",
		GuaJiFeatherAdvanced:        "正在护体仙羽自动进阶",
		GuaJiShieldAdvanced:         "正在盾刺自动进阶",
		GuaJiMassacreAdvanced:       "正在戮仙刃自动进阶",
		GuaJiLingTongWeaponAdvanced: "正在灵兵自动进阶",
		GuaJiLingTongMountAdvanced:  "正在灵骑自动进阶",
		GuaJiLingTongWingAdvanced:   "正在灵翼自动进阶",
		GuaJiLingTongShenFaAdvanced: "正在灵身自动进阶",
		GuaJiLingTongLingYuAdvanced: "正在灵域自动进阶",
		GuaJiLingTongFaBaoAdvanced:  "正在灵宝自动进阶",
		GuaJiLingTongXianTiAdvanced: "正在灵体自动进阶",
	}
)

func init() {
	mergeLang(guaJiLangMap)
}
