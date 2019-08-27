package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/songbuting/player"
)

func BuildSCSongBuTingChanged(songBuTingObj *player.PlayerSongBuTingObject) *uipb.SCSongBuTingChanged {
	scSongBuTingChanged := &uipb.SCSongBuTingChanged{}
	isReceive := songBuTingObj.GetIsReceive()
	times := songBuTingObj.GetTimes()
	scSongBuTingChanged.IsReceive = &isReceive
	scSongBuTingChanged.Times = &times
	return scSongBuTingChanged
}
