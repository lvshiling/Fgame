package types

type ShengTanEventType string

const (
	EventTypeShengTanSceneEnd             ShengTanEventType = "ShengTanSceneEnd"
	EventTypeShengTanSceneJiuNiangChanged                   = "ShengTanSceneJiuLiangChanged"
	EventTypeShengTanSceneTickReward                        = "ShengTanSceneTickReward"
)
