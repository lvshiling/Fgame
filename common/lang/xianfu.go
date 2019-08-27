package lang

const (
	XianfuArgumentInvalid LangCode = XianfuBase + iota
	XianfuReachMaxGrade
	XianfuNotEnoughGold
	XinafuNotEnoughChallengeTimes
	XinafuNotEnoughSaodangItem
	XinafuNotEnoughChallengeItem
	XinafuUpgradeNotice
)

var (
	xianfuLangMap = map[LangCode]string{
		XianfuArgumentInvalid:         "参数无效",
		XianfuReachMaxGrade:           "当前建筑已达最高级，无法升级",
		XianfuNotEnoughGold:           "当前元宝不足，无法升级",
		XinafuNotEnoughChallengeTimes: "副本次数不足，无法扫荡",
		XinafuNotEnoughSaodangItem:    "当前扫荡券不足，无法扫荡",
		XinafuNotEnoughChallengeItem:  "副本挑战令不足，无法扫荡",
		XinafuUpgradeNotice:           "恭喜%s成功将%s提升到%s级，%s产出大幅度提升",
	}
)

func init() {
	mergeLang(xianfuLangMap)
}
