package types

import (
	chattypes "fgame/fgame/game/chat/types"
	itemtypes "fgame/fgame/game/item/types"
)

type HongBaoEventType string

const (
	//红包发送
	EventTypeHongBaoSend HongBaoEventType = "HongBaoSend"
	//红包感谢发言
	EventTypeHongBaoThanks HongBaoEventType = "HongBaoThanks"
)

//红包发送类型
type HongBaoSendEventData struct {
	hongBaoId          int64
	hongBaoType        itemtypes.ItemHongBaoSubType
	hongBaoChannelType chattypes.ChannelType
	cliArgs            string
}

func CreateHongBaoSendEventData(id int64, typ itemtypes.ItemHongBaoSubType, channelType chattypes.ChannelType, args string) *HongBaoSendEventData {
	d := &HongBaoSendEventData{
		hongBaoId:          id,
		hongBaoType:        typ,
		hongBaoChannelType: channelType,
		cliArgs:            args,
	}
	return d
}

func (d *HongBaoSendEventData) GetHongBaoId() int64 {
	return d.hongBaoId
}

func (d *HongBaoSendEventData) GetHongBaoType() itemtypes.ItemHongBaoSubType {
	return d.hongBaoType
}

func (d *HongBaoSendEventData) GetHongBaoChannelType() chattypes.ChannelType {
	return d.hongBaoChannelType
}

func (d *HongBaoSendEventData) GetCliArgs() string {
	return d.cliArgs
}
