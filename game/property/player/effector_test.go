package player_test

import (
	"fgame/fgame/game/property/player/types"
	"testing"
)

var (
	mask = uint64(9223372036854775807)
)

func TestMask(t *testing.T) {
	for i := types.PlayerPropertyEffectorTypeInit; i <= types.PlayerPropertyEffectorTypeBaby; i++ {
		if i.Mask()&mask != 0 {
			t.Error(i.String())
		}
	}
}
