package logic

import (
	"fgame/fgame/game/player"
)

// import (
// 	"fgame/fgame/common/lang"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	chuangshidata "fgame/fgame/game/chuangshi/data"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	playerchuangshi "fgame/fgame/game/chuangshi/player"
// 	chuangshitypes "fgame/fgame/game/chuangshi/types"
// 	commomlogic "fgame/fgame/game/common/logic"
// 	"fgame/fgame/game/player"
// 	playerlogic "fgame/fgame/game/player/logic"
// 	playertypes "fgame/fgame/game/player/types"
// 	gametemplate "fgame/fgame/game/template"
// 	"fmt"

// 	log "github.com/Sirupsen/logrus"
// )

// // 神王竞选报名结果
// func ShenWangSignUpResult(pl player.Player) {

// 	signUpObj := chuangshi.GetChuangShiService().GetShenWangSignUp(pl.GetId())
// 	if signUpObj == nil {
// 		return
// 	}

// 	if signUpObj.IfSigning() {
// 		return
// 	}

// 	status := signUpObj.GetStatues()
// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangShiManager.SignUpResult(status)
// 	chuangshi.GetChuangShiService().ShenWangSignUpRemove(pl.GetId())

// 	scMsg := pbutil.BuildSCChuangShiShenWangBaoMing(int32(status))
// 	pl.SendMsg(scMsg)
// }

// // 神王竞选投票结果
// func ShenWangVoteResult(pl player.Player) {
// 	voteObj := chuangshi.GetChuangShiService().GetShenWangVote(pl.GetId())
// 	if voteObj == nil {
// 		return
// 	}

// 	if voteObj.IfVoting() {
// 		return
// 	}

// 	status := voteObj.GetStatues()
// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangShiManager.VoteResult(status)
// 	chuangshi.GetChuangShiService().ShenWangSignUpRemove(pl.GetId())

// 	scMsg := pbutil.BuildSCChuangShiShengWangTouPiao(int32(status), voteObj.GetSupportId())
// 	pl.SendMsg(scMsg)
// }

// // 城池建设结果
// func ChengFangJianSheResult(pl player.Player) {
// 	jianSheObj := chuangshi.GetChuangShiService().GetChengFangJianShe(pl.GetId())
// 	if jianSheObj == nil {
// 		return
// 	}

// 	if jianSheObj.IfProgressing() {
// 		return
// 	}

// 	chuangshi.GetChuangShiService().ChengFangJianSheRemove(pl.GetId())
// }

// // 创世之战升职判断
// func ChuangShiGuanZhiAdvance(curTimesNum int32, temp *gametemplate.ChuangShiGuanZhiTemplate) (success bool) {
// 	timesMin := temp.TimesMin
// 	timesMax := temp.TimesMax
// 	updateRate := temp.UpLevPercent
// 	_, _, success = commomlogic.GetStatusAndProgress(curTimesNum, 0, timesMin, timesMax, 0, 0, updateRate, 0)
// 	return
// }

// //加入阵营
// func PlayerJoinCamp(pl player.Player, campType chuangshitypes.ChuangShiCampType) (success bool) {
// 	chuangshiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangShiObj := chuangshiManager.GetPlayerChuangShiInfo()
// 	if chuangShiObj.IfJoinCamp() {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世加入阵营失败，已经存在阵营")
// 		return
// 	}

// 	flag := chuangshi.GetChuangShiService().JoinCamp(campType, chuangshidata.CareteMemberInfo(pl))
// 	if !flag {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 			}).Warnln("chuangshi:处理创世加入阵营失败")
// 		playerlogic.SendSystemMessage(pl, lang.ChuangShiJoinCampFailed)
// 		return
// 	}

// 	success = chuangshiManager.JoinCamp(campType)
// 	if !success {
// 		panic(fmt.Errorf("chuangshi:加入阵营应该成功"))
// 	}

// 	scMsg := pbutil.BuildSCChuangShiJoinCamp(int32(campType))
// 	pl.SendMsg(scMsg)
// 	return
// }

// func SendPlayerChuangShiInfo(pl player.Player) {
// 	chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
// 	chuangShiInfo := chuangShiManager.GetPlayerChuangShiInfo()
// 	chuangShiGuanZhiInfo := chuangShiManager.GetPlayerChuangShiGuanZhiInfo()
// 	signInfo := chuangShiManager.GetPlayerChuangShiSignInfo()
// 	voteInfo := chuangShiManager.GetPlayerChuangShiVoteInfo()

// 	campList := chuangshi.GetChuangShiService().GetCampList()
// 	scMsg := pbutil.BuildSCChuangShiInfo(campList, chuangShiInfo, chuangShiGuanZhiInfo, int32(signInfo.GetStatus()), int32(voteInfo.GetStatus()))
// 	pl.SendMsg(scMsg)
// }

func SendPlayerChuangShiInfo(pl player.Player) {
	// chuangShiInfo := playerchuangshi.NewPlayerChuangShiObject(pl)
	// chuangShiGuanZhiInfo := playerchuangshi.NewPlayerChuangShiGuanZhiObject(pl)
	// signInfo := playerchuangshi.NewPlayerChuangShiSignObject(pl)
	// voteInfo := playerchuangshi.NewPlayerChuangShiVoteObject(pl)
	// chuangShiInfo.SetCampType(chuangshitypes.ChuangShiCampTypePanGu)

	// scMsg := pbutil.BuildSCChuangShiPlayerInfoNotice(chuangShiInfo, chuangShiGuanZhiInfo, int32(signInfo.GetStatus()), int32(voteInfo.GetStatus()))
	// pl.SendMsg(scMsg)
}
