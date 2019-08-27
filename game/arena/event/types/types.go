package types

type ArenaEventType string

const (
	//竞技场获胜次数变化
	EventTypeArenaWinChanged ArenaEventType = "ArenaWinChanged"
	//复活次数变化
	EventTypeArenaReliveChanged = "ArenaReliveChanged"
	//参加3v3活动
	EventTypeArenaJoin = "ArenaJoin"
	//3v3层数
	EventTypeArenaLianSheng = "ArenaCengShu"
	//3v3积分变化
	EventTypeArenaJiFenChanged = "ArenaJiFenChanged"
)
