package lang

const (
	DanStillNotGet = DanBase + iota
	DanHasedFinshed
	DanNotFinished
	DanUseReachedLimit
	DanNotCanUse
	DanUseIsNotEnough
	DanLevelReachedLimit
	DanLevelReachedLimitNotEat
)

var (
	danLangMap = map[LangCode]string{
		DanStillNotGet:             "当前正在炼丹或有丹药未领取,无法请求炼丹",
		DanHasedFinshed:            "丹药已练成,无需加速",
		DanNotFinished:             "炼丹未完成,无法领取丹药",
		DanUseReachedLimit:         "当前食用已达上限,请升级后再食用",
		DanNotCanUse:               "当前等级无可食用丹药,请先炼丹",
		DanUseIsNotEnough:          "当前食用丹药进度未满,无法升级",
		DanLevelReachedLimit:       "当前食丹等级已达最高级,无法再升级",
		DanLevelReachedLimitNotEat: "当前食丹等级已达最高级,无法再食用",
	}
)

func init() {
	mergeLang(danLangMap)
}
