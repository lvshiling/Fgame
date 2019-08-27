package types

const (
	defaultKillName = "未知生物"
)

const (
	FeiShengKillerId int64 = -1000 + iota
)

var (
	killNameMap = map[int64]string{
		FeiShengKillerId: "飞升雷劫",
	}
)

func GetKillName(killId int64) string {
	name, ok := killNameMap[killId]
	if !ok {
		return defaultKillName
	}
	return name
}
