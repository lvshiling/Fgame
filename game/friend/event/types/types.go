package types

import (
	commonlog "fgame/fgame/common/log"
)

type FriendEventType string

const (
	EventTypeFriendGift            FriendEventType = "FriendGift"
	EventTypeFriendAddRefuse       FriendEventType = "FriendAddRefuse"
	EventTypeFriendAdd             FriendEventType = "FriendAdd"
	EventTypeFriendDelete          FriendEventType = "FriendDelete"
	EventTypeFriendBlack           FriendEventType = "FriendBlack"
	EventTypeFriendRemoveBlack     FriendEventType = "FriendRemoveBlack"
	EventTypeFriendPointChanged    FriendEventType = "FriendPointChanged"
	EventTypeFriendInvite          FriendEventType = "FriendInvite"
	EventTypeFriendBatchRefuse     FriendEventType = "FriendBatchRefuse"
	EventTypeFriendDummyNumChanged FriendEventType = "FriendDummyNumChanged"
	EventTypeFriendAddAll          FriendEventType = "FriendAddAll"
	EventTypeFriendGiveLog         FriendEventType = "FriendGiveLog"
)

type FriendGiftEventData struct {
	playerId  int64
	friendId  int64
	itemId    int32
	itemCount int32
	num       int32 //豪气值
	charmNum  int32 //魅力值
}

func (d *FriendGiftEventData) GetPlayerId() int64 {
	return d.playerId
}

func (d *FriendGiftEventData) GetFriendId() int64 {
	return d.friendId
}
func (d *FriendGiftEventData) GetNum() int32 {
	return d.num
}

func (d *FriendGiftEventData) GetCharmNum() int32 {
	return d.charmNum
}

func (d *FriendGiftEventData) GetItemId() int32 {
	return d.itemId
}

func (d *FriendGiftEventData) GetItemCount() int32 {
	return d.itemCount
}

func CreateFriendGiftEventData(playerId int64, friendId int64, itemId int32, itemCount int32, num int32, charmNum int32) *FriendGiftEventData {
	return &FriendGiftEventData{
		playerId:  playerId,
		friendId:  friendId,
		itemId:    itemId,
		itemCount: itemCount,
		num:       num,
		charmNum:  charmNum,
	}
}

//赠送装备日志
type FriendGiveEventData struct {
	reason     commonlog.FriendLogReason
	reasonText string
}

func (d *FriendGiveEventData) GetReason() commonlog.FriendLogReason {
	return d.reason
}

func (d *FriendGiveEventData) GetReasonText() string {
	return d.reasonText
}

func CreateFriendGiveEventData(reason commonlog.FriendLogReason, reasonText string) *FriendGiveEventData {
	return &FriendGiveEventData{
		reason:     reason,
		reasonText: reasonText,
	}
}
