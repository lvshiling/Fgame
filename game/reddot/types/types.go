package types

type RedDotInfo struct {
	GroupId  int32
	IsReddot bool
}

func NewRedDotInfo(groupId int32, isRed bool) *RedDotInfo {
	d := &RedDotInfo{}
	d.GroupId = groupId
	d.IsReddot = isRed
	return d
}
