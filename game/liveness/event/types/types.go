package types

type LivenessEventType string

const (
	EventTypeLivenessChanged   LivenessEventType = "LivenessChanged"   //活跃度改变
	EventTypeLivenessCrossFive LivenessEventType = "LivenessCrossFive" //活跃度过5点
)
