package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_BASIC_PROPERTY_CHANGED_TYPE), dispatch.HandlerFunc(handlePlayerBasicPropertyChanged))
}

//处理基础属性变更
func handlePlayerBasicPropertyChanged(s session.Session, msg interface{}) error {
	log.Debug("login:处理基础属性变更")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siPlayerBasicPropertyChanged := msg.(*crosspb.SIPlayerBasicPropertyChanged)
	basicPropertyData := siPlayerBasicPropertyChanged.GetBasicPropertyData()
	baseProperties := pbutil.ConvertFromBaseProperty(basicPropertyData)
	err := playerBasePropertyChanged(pl, baseProperties)

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家跨服登陆,失败")
		return err
	}

	log.Debug("login:处理跨服登陆消息完成")
	return nil
}

//玩家基础属性变化
func playerBasePropertyChanged(pl *player.Player, baseProperties map[int32]int64) (err error) {
	return
}
