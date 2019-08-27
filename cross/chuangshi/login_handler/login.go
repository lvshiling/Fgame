package login_handler

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/cross/chuangshi/chuangshi"
	"fgame/fgame/cross/login/login"
	"fgame/fgame/cross/player/player"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	crosstypes "fgame/fgame/game/cross/types"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	login.RegisterLoginHandler(crosstypes.CrossTypeChuangShi, login.LogincHandlerFunc(chuangShiEnterCityLogin))
}

// 进入城池
func chuangShiEnterCityLogin(pl *player.Player, ct crosstypes.CrossType, crossArgs ...string) bool {
	if len(crossArgs) <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"len":      len(crossArgs),
			}).Warn("login:玩家加载数据,参数不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return false
	}

	cityId, err := strconv.ParseInt(crossArgs[0], 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"arg":      crossArgs[0],
			}).Warn("login:玩家加载数据,参数不对")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return false
	}

	cityData := chuangshi.GetChuangShiService().GetChuangShiCityData(cityId)
	campType := cityData.GetCity().GetCampType()
	cityType := cityData.GetCity().GetType()
	index := cityData.GetCity().GetIndex()
	cityTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCityTemp(campType, cityType, index)
	if cityTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"cityId":   cityId,
			}).Warn("login:玩家加载数据,城池模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return false
	}

	var s scene.Scene
	switch cityType {
	case chuangshitypes.ChuangShiCityTypeMain:
		{
			s = chuangshi.GetChuangShiService().GetChuangShiMainScene(campType)
		}
	case chuangshitypes.ChuangShiCityTypeZhongli:
		{
			s = chuangshi.GetChuangShiService().GetChuangShiZhongLiScene()
		}
	case chuangshitypes.ChuangShiCityTypeFushu:
		{
			s = chuangshi.GetChuangShiService().GetChuangShiFuShuScene(cityId)
		}
	}

	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"campType": campType,
			}).Warnln("login:玩家进入城池错误，场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return false
	}

	bornPos := cityTemp.GetBornPos(pl.GetCamp())
	if !scenelogic.PlayerEnterScene(pl, s, bornPos) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"scene":    s,
			}).Warnln("login:玩家进入城池错误，进入场景失败")
	}
	return true
}
