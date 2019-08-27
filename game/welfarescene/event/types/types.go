package types

type WelfareSceneEventType string

const (
	EventTypeWelfareSceneFinish  WelfareSceneEventType = "WelfareSceneFinish"  //运营活动副本结束
	EventTypeWelfareSceneRefresh                       = "WelfareSceneRefresh" //运营活动副本刷新
)
