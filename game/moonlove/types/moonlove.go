package types

type MoonloveRankType int32

const (
	MoonloveRankTypeGenerous      MoonloveRankType = iota //豪气
	MoonloveRankTypeGenerousCharm                         //魅力
)

func (t MoonloveRankType) Valid() bool {
	switch t {
	case MoonloveRankTypeGenerous,
		MoonloveRankTypeGenerousCharm:
		return true
	default:
		return false
	}
}

const (
	ColorTypeEmailRanking = "#4377ef" //邮件奖励排名颜色
)
