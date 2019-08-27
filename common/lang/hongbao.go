package lang

const (
	HongBaoSnatchCountReachedLimit = HongBaoBase + iota
	HongBaoExpire
	HongBaoSnatchNeedConditionLimit
	HongBaoSnatchThanksBoss
)

var (
	hongbaoLangMap = map[LangCode]string{
		HongBaoSnatchCountReachedLimit:  "每日抢红包次数已达上限",
		HongBaoExpire:                   "红包已过期",
		HongBaoSnatchNeedConditionLimit: "不满足抢红包条件",
		HongBaoSnatchThanksBoss:         "谢谢老板",
	}
)

func init() {
	mergeLang(hongbaoLangMap)
}
