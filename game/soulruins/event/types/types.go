package types

type SoulruinsChallengeEventType string

const (
	//帝魂遗迹挑战事件
	EventTypeSoulruinsChallenge SoulruinsChallengeEventType = "SoulruinsChallenge"
	//帝陵遗迹完成(含扫荡)
	EventTypeSoulruinsFinish SoulruinsChallengeEventType = "SoulruinsFinish"
	//帝陵遗迹扫荡
	EventTypeSoulruinsSweep SoulruinsChallengeEventType = "SoulruinsSweep"
)

type SoulRuinsFinishEventData struct {
	soulRuinsId int32
	num         int32
}

func (d *SoulRuinsFinishEventData) GetSoulRuinsId() int32 {
	return d.soulRuinsId
}

func (d *SoulRuinsFinishEventData) GetNum() int32 {
	return d.num
}

func CreateSoulRuinsFinishEventData(soulRuinsId int32, num int32) *SoulRuinsFinishEventData {
	d := &SoulRuinsFinishEventData{
		soulRuinsId: soulRuinsId,
		num:         num,
	}
	return d
}
