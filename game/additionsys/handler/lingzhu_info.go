package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_LINGTONG_LINGZHU_INFO_TYPE), dispatch.HandlerFunc(handleAdditionSysLingZhuInfo))
}

func handleAdditionSysLingZhuInfo(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSAdditionSysLingTongLingZhuInfo)
	lingtongId := csMsg.GetLingtongId()
	sysType, _ := additionsystypes.ConvertLingTongIdToAdditionSysType(int(lingtongId))
	//参数不对
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsys:灵珠信息请求类型,错误")
		return
	}
	err = additionSysLingZhuInfo(tpl, sysType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
				"error":    err,
			}).Warn("additionsys:灵珠信息请求类型,错误")
		return
	}

	return
}

func additionSysLingZhuInfo(pl player.Player, typ additionsystypes.AdditionSysType) (err error) {
	if !additionsyslogic.GetAdditionSysLingZhuFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("inventory:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	objList := []*playeradditionsys.PlayerAdditionSysLingZhuObject{}
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	lingZhuMap := additionsysManager.GetAdditionSysLingZhuMap(typ)
	for _, lingzhuObj := range lingZhuMap {
		// lingzhuObj := additionsysManager.GetAdditionSysLingZhu(typ, lingzhuObj.GetLingZhuType())
		objList = append(objList, lingzhuObj)
	}
	id, _ := typ.ConvertAdditionSysTypeToLingTongId()
	scMsg := pbutil.BuildSCAdditionSysLingZhuInfo(int32(id), objList)
	pl.SendMsg(scMsg)
	return
}
