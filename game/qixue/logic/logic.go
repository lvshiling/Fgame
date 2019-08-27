package logic

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/common/common"
	consttypes "fgame/fgame/game/constant/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/qixue/pbutil"
	playerqixue "fgame/fgame/game/qixue/player"
	qixuetemplate "fgame/fgame/game/qixue/template"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//变更泣血枪属性
func QiXuePropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeQiXue.Mask())
	return
}

//泣血枪进阶
func HandleQiXueAdvanced(pl player.Player) (err error) {

	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeQiXueQiang) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("qixue:功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	qixueManager := pl.GetPlayerDataManager(playertypes.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	qiXueObj := qixueManager.GetQiXueInfo()
	curLev := qiXueObj.GetLevel()
	curStar := qiXueObj.GetStar()

	qiXueTemplate := qixuetemplate.GetQiXueTemplateService().GetQiXueTemplateByLevel(curLev, curStar)
	if qiXueTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"curLev":   curLev,
				"curStar":  curStar,
			}).Warn("qixue:泣血枪模板不存")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	nextTemp := qiXueTemplate.GetNextTemp()
	if nextTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("qixue:泣血枪已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.QiXueAdanvacedReachedLimit)
		return
	}

	if !qiXueObj.IfEnoughShaLuNum() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("qixue:泣血枪注入失败，没有杀戮心")
		playerlogic.SendSystemMessage(pl, lang.QiXueAdanvacedShaLuNotEnough)
		return
	}

	//注入杀戮心
	flag := qixueManager.UseShaLuNum()
	if !flag {
		panic(fmt.Errorf("qixue: qixueAdvanced use shaqinum should be ok"))
	}

	//同步属性
	QiXuePropertyChanged(pl)

	scMsg := pbutil.BuildSCQiXueAdavanced(qiXueObj)
	pl.SendMsg(scMsg)
	return
}

//玩家爆杀戮心
func QiXueProcessDrop(pl player.Player, attackId int64, attackName string) (itemId int32, dropNum int64) {
	manager := pl.GetPlayerDataManager(types.PlayerQiXueDataManagerType).(*playerqixue.PlayerQiXueDataManager)
	if !manager.IfDropCD() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("qixue:处理获取泣血枪掉落,掉落冷却中")
		return
	}

	//掉落储存的杀戮心
	itemId = int32(consttypes.ShaLuXin)
	//获取掉落的
	flag, dropNum, costStar := manager.QiXueDrop()
	if !flag {
		return
	}

	//掉落数量
	if dropNum > 0 {
		systemRecycleRatio := qixuetemplate.GetQiXueTemplateService().GetQiXueConstantTemplate().DropSystemReturn
		backBuyXiShu := float64(1 - float64(systemRecycleRatio)/float64(common.MAX_RATE))
		dropNum = int64(math.Ceil(float64(dropNum) * backBuyXiShu))
	}

	//同步属性
	if costStar > 0 {
		QiXuePropertyChanged(pl)
	}

	if dropNum > 0 || costStar > 0 {
		//告诉前端掉落杀气了
		scMsg := pbutil.BuildSCQiXueShaQiDrop(manager.GetQiXueInfo(), costStar, int32(dropNum), attackName)
		pl.SendMsg(scMsg)
	}

	return
}
