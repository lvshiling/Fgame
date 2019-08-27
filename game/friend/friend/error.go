package friend

import (
	"fgame/fgame/common/lang"
	gamecommon "fgame/fgame/game/common/common"
)

var (
	ErrorFriendIsNotFriend    = gamecommon.CodeError(lang.FriendIsNotFriend)
	ErrorFriendPeerAlreadFull = gamecommon.CodeError(lang.FriendPeerAlreadFull)
)
