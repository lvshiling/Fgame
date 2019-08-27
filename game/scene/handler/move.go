package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_CS_OBJECT_MOVE_TYPE), dispatch.HandlerFunc(handleMove))
}

//处理移动包
func handleMove(s session.Session, msg interface{}) error {
	log.Debug("scene:处理对象移动消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(scene.Player)
	csObjectMove := msg.(*scenepb.CSObjectMove)

	pos := csObjectMove.GetPos()
	moveSpeed := csObjectMove.GetMoveSpeed()
	angle := csObjectMove.GetAngle()
	moveTypeInt := csObjectMove.GetMoveType()
	aPos := coretypes.Position{
		X: float64(pos.GetPosX()),
		Y: float64(pos.GetPosY()),
		Z: float64(pos.GetPosZ()),
	}
	moveType := types.MoveType(moveTypeInt)
	uId := csObjectMove.GetUid()
	moveFlag := csObjectMove.GetFlag()
	aCurPos := aPos
	curPos := csObjectMove.GetCurPos()
	if curPos != nil {
		aCurPos = coretypes.Position{
			X: float64(curPos.GetPosX()),
			Y: float64(curPos.GetPosY()),
			Z: float64(curPos.GetPosZ()),
		}
	}

	//用户移动
	scenelogic.HandleObjectMove(tpl, uId, aPos, aCurPos, float64(moveSpeed), float64(angle), moveType, moveFlag)
	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"pos":       aPos,
			"angle":     angle,
			"moveType":  moveType,
			"moveSpeed": moveSpeed,
		}).Debug("scene:处理对象移动消息,完成")
	return nil
}
