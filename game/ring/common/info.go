package common

import (
	ringtypes "fgame/fgame/game/ring/types"
)

// 特戒信息
type RingInfo struct {
	Typ          ringtypes.RingType          `json:"typ"`
	ItemId       int32                       `json:"itemId"`
	PropertyData *ringtypes.RingPropertyData `json:"property"`
}
