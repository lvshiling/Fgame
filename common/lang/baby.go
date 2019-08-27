package lang

const (
	BabyBornTitle LangCode = BabyBase + iota
	BabyBornContent
	BabyMaxNum
	BabyCoupleNotOnline
	BabyFullTonic
	BabyFullChaoSheng
	BabyFailReturnMailContent
	BabyFailActivitySkill
	BabyToySlotHadEquip
	BabyToyFullLevel
	BabyLearnFullLevel
	BabyTalentLockFail
	BabyPregnantNotMarry
)

var (
	babyLangMap = map[LangCode]string{
		BabyBornTitle:             "宝宝出生啦",
		BabyBornContent:           "您的伴侣成功培养出了一个宝宝，双方都将拥有，请查收！",
		BabyMaxNum:                "当前宝宝个数已达上限，无法使用，超生可提高宝宝数量",
		BabyCoupleNotOnline:       "您的伴侣不在线，无法洞房",
		BabyFullTonic:             "当前补品食用已达上限，请等待宝宝出生吧",
		BabyFullChaoSheng:         "当前已达超生最大数量，无法再超生",
		BabyFailReturnMailContent: "洞房失败返还物品",
		BabyFailActivitySkill:     "没有可激活的天赋",
		BabyToySlotHadEquip:       "该位置已经装备玩具",
		BabyToyFullLevel:          "玩具已经满级",
		BabyLearnFullLevel:        "四书五经已达最高级",
		BabyTalentLockFail:        "天赋技能已解锁或已锁定",
		BabyPregnantNotMarry:      "没有伴侣",
	}
)

func init() {
	mergeLang(babyLangMap)
}
