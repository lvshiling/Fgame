package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/chess/pbutil"
	playerchess "fgame/fgame/game/chess/player"
	chesstypes "fgame/fgame/game/chess/types"
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHESS_CHANGED_TYPE), dispatch.HandlerFunc(handleChessChanged))

}

//处理苍龙棋局换一批
func handleChessChanged(s session.Session, msg interface{}) (err error) {
	log.Debug("chess:处理苍龙棋局换一批")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csChessChanged := msg.(*uipb.CSChessChanged)
	typ := chesstypes.ChessType(csChessChanged.GetTyp())
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("chess:破解棋局错误,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	//银两无换一批功能
	if typ == chesstypes.ChessTypeSilver {
		return
	}

	err = chessChanged(tpl, typ)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("chess:处理苍龙棋局换一批,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chess:处理苍龙棋局换一批完成")
	return nil

}

//苍龙棋局换一批逻辑
func chessChanged(pl player.Player, chessType chesstypes.ChessType) (err error) {

	chessManager := pl.GetPlayerDataManager(playertypes.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	needGold := int64(constant.GetConstantService().GetConstant(chessType.GetChangeConstantType()))

	isUseBindGold := false
	if chessType == chesstypes.ChessTypeBindGold {
		isUseBindGold = true
	}

	//是否足够元宝
	flag := propertyManager.HasEnoughGold(int64(needGold), isUseBindGold)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("chess:破解棋局错误，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//消耗元宝
	if needGold > 0 {
		goldUseReason := commonlog.GoldLogReasonChessUse
		goldUseReasonText := fmt.Sprintf(goldUseReason.String(), chesstypes.ChessTypeGold.String())
		flag := propertyManager.CostGold(needGold, isUseBindGold, goldUseReason, goldUseReasonText)
		if !flag {
			panic("chess:消耗元宝应该成功")
		}
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	//更新棋局
	newDropId := chessManager.ChangedChess(chessType)

	scChessChanged := pbutil.BuildSCChessChanged(newDropId)
	pl.SendMsg(scChessChanged)
	return
}
