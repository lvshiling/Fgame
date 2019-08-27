package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	quizeventtypes "fgame/fgame/game/quiz/event/types"
	"fgame/fgame/game/quiz/pbutil"
	quiz "fgame/fgame/game/quiz/quiz"
	quiztemplate "fgame/fgame/game/quiz/template"
	quiztypes "fgame/fgame/game/quiz/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_QUIZ_ANSWER_TYPE), dispatch.HandlerFunc(handleQuizAnswer))

}

//仙尊问答答题
func handleQuizAnswer(s session.Session, msg interface{}) (err error) {
	log.Debug("quiz:仙尊问答答题")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csQuizAnswer := msg.(*uipb.CSQuizAnswer)
	answerType := quiztypes.QuizAnswerType(csQuizAnswer.GetAnswerIdx())

	if !answerType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"AnswerIdx": int32(answerType),
			}).Warn("quiz:仙尊问答答题,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = quizAnswer(tpl, answerType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"AnswerIdx": int32(answerType),
				"error":     err,
			}).Error("quiz:仙尊问答答题,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quiz:仙尊问答答题完成")
	return nil

}

//仙尊问答答题逻辑
func quizAnswer(pl player.Player, answerType quiztypes.QuizAnswerType) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeQuiz) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"answerType": int32(answerType),
			}).Warn("quiz:仙尊问答答题，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	quizService := quiz.GetQuizService()
	quizObj := quizService.GetQuizObj()
	curTemplate := quiztemplate.GetQuizTemplateService().GetQuizByTemplateId(quizObj.GetLastQuizId())
	if curTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"answerType": int32(answerType),
			}).Warn("quiz:仙尊问答答题,题目模板错误")
		return
	}
	//是否答题时间内
	if !quizService.CheckQuizAnswerTime() {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"answerType": int32(answerType),
			}).Warn("quiz:仙尊问答答题,超过答题时间")

		playerlogic.SendSystemMessage(pl, lang.QuizAnswerTimeLimit)
		return
	}

	//是否已经答过了
	if quizService.IsAnswerPlayerById(pl.GetId()) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"answerType": int32(answerType),
			}).Warn("quiz:仙尊问答答题,已经答过题目了")

		playerlogic.SendSystemMessage(pl, lang.QuizAnswerAlreadyGet)
		return
	}

	//是否答题对
	result := int32(0)
	if int32(curTemplate.GetRightAnswer()) == int32(answerType) {
		result = 1
	}
	//加入答题列表
	quizService.AddAnswerPlayer(pl.GetId())
	//奖励
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	var isRewItem bool
	var rewItemMap map[int32]int32
	var rewData *propertytypes.RewData
	if result == 1 {
		rewData = curTemplate.GetRewData()
		//奖励物品
		isRewItem = mathutils.RandomHit(common.MAX_RATE, int(curTemplate.RewardItemRate))
		rewItemMap = curTemplate.GetRewItemMap()
	} else {
		rewData = curTemplate.GetRewErrorData()
		//奖励物品
		isRewItem = mathutils.RandomHit(common.MAX_RATE, int(curTemplate.ErrorRewardItemRate))
		rewItemMap = curTemplate.GetErrorRewItemMap()
	}

	if rewData != nil {
		reasonGold := commonlog.GoldLogReasonQuizAnswer
		reasonSilver := commonlog.SilverLogReasonQuizAnswer
		reasonLevel := commonlog.LevelLogReasonQuizAnswer
		reasonGoldText := fmt.Sprintf(reasonGold.String(), pl.GetLevel(), result)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), pl.GetLevel(), result)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), pl.GetLevel(), result)

		flagRewData := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flagRewData {
			panic(fmt.Errorf("quiz: GivequizReward AddRewData  should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	if isRewItem && len(rewItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			emailTitle := lang.GetLangService().ReadLang(lang.QuizAnswerRewTitle)
			emailContent := lang.GetLangService().ReadLang(lang.QuizAnswerRewContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			inventoryLogReason := commonlog.InventoryLogReasonQuizAnswer
			reasonText := fmt.Sprintf(inventoryLogReason.String(), pl.GetLevel(), result)
			flag = inventoryManager.BatchAdd(rewItemMap, inventoryLogReason, reasonText)
			if !flag {
				panic(fmt.Errorf("quiz: GivequizReward BatchAdd should be ok"))
			}
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	//答题消息事件
	data := quizeventtypes.CreateQuizAnswerChatEventData(pl.GetId(), answerType)
	gameevent.Emit(quizeventtypes.EventTypeQuizAnswerChat, nil, data)
	//返回前端
	scMsg := pbutil.BuildSCQuizAnswer(result, int32(curTemplate.GetRightAnswer()))
	pl.SendMsg(scMsg)
	return
}
