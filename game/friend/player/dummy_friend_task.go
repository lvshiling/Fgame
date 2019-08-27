package player

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"time"
)

const (
	taskAddTime = time.Second * 5
)

//虚拟好友数量
type AddDummyFriendTask struct {
	pl player.Player
}

func (t *AddDummyFriendTask) Run() {
	friendManager := t.pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*PlayerFriendDataManager)
	friendManager.addDummyFriend()
}

//间隔时间
func (t *AddDummyFriendTask) ElapseTime() time.Duration {
	return taskAddTime
}

func CreateAddDummyFriendTask(pl player.Player) *AddDummyFriendTask {
	t := &AddDummyFriendTask{
		pl: pl,
	}
	return t
}
