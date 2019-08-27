package log

type FriendLogReason int32

const (
	FriendLogReasonGM FriendLogReason = iota + 1
	FriendLogReasonAgreeGive
)

func (zslr FriendLogReason) Reason() int32 {
	return int32(zslr)
}

var (
	friendLogReasonMap = map[FriendLogReason]string{
		FriendLogReasonGM:        "gm修改",
		FriendLogReasonAgreeGive: "同意赠送好友装备,好友id：%d，物品id：%d，数量：%d",
	}
)

func (ar FriendLogReason) String() string {
	return friendLogReasonMap[ar]
}
