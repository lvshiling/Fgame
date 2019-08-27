package lang

const (
	HouseActivateFali = HouseBase + iota
	HouseFullLevel
	HouseHasActivate
	HouseNotEnoughTimes
	HouseHadBroken
	HouseHadRent
	HouseRentEmailTitle
	HouseRentEmailContent
)

var (
	houseLangMap = map[LangCode]string{
		HouseActivateFali:     "激活本房子需要上一个房子装修等级达到最高",
		HouseFullLevel:        "当前房子已装修到最高档，无法继续装修",
		HouseHasActivate:      "房子已经激活",
		HouseNotEnoughTimes:   "今天工人已经很累了，请明天再来装修",
		HouseHadBroken:        "房子已经损坏，无法收取租金",
		HouseHadRent:          "房子已经领取租金",
		HouseRentEmailTitle:   "房租收益",
		HouseRentEmailContent: "您的第%d座房子昨日租金：%d，您未进行领取，系统自动为您进行补发，敬请查收",
	}
)

func init() {
	mergeLang(houseLangMap)
}
