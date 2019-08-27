package logic

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/marry/marry"
	marryscene "fgame/fgame/game/marry/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/team/team"

	log "github.com/Sirupsen/logrus"
)

func CheckCanEnterScene(pl player.Player) bool {
	//玩家正在跨服中
	if pl.IsCross() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:正在跨服中")
		SendSystemMessage(pl, lang.PlayerInCross)
		return false
	}

	//正在3v3匹配
	teamId := pl.GetTeamId()
	if teamId != 0 {
		teamDataObj := team.GetTeamService().GetTeam(teamId)
		if teamDataObj != nil && teamDataObj.IsMatch() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:正在3v3匹配")
			SendSystemMessage(pl, lang.PlayerIn3v3Match)
			return false
		}

		if teamDataObj != nil && teamDataObj.IsCopyBattle() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:玩家当前处于组队副本战斗中,无法传送")
			SendSystemMessage(pl, lang.PlayerInTeamCopyBattle)
			return false
		}
	}

	//正在无间炼狱排队
	flag := pl.IsLianYuLineUp()
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:您当前正在无间炼狱排队中,无法进入副本")
		SendSystemMessage(pl, lang.PlayerInLianYuLineUp)
		return false
	}

	//正在神兽攻城排队
	flag = pl.IsGodSiegeLineUp()
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:您当前正在神兽攻城排队中,无法传送")
		SendSystemMessage(pl, lang.PlayerInGodSiegeLineUp)
		return false
	}

	//判断玩家当前是否处于婚宴游街
	marrySceneData := marry.GetMarryService().GetMarrySceneData()
	if marrySceneData.Status == marryscene.MarrySceneStatusCruise {
		if pl.GetId() == marrySceneData.PlayerId || pl.GetId() == marrySceneData.SpouseId {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:玩家当前处于婚礼游街状态,无法传送")
			SendSystemMessage(pl, lang.PlayerInMarryCruise)
			return false
		}
	}

	pls := pl.GetScene()
	if pls == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:玩家不在场景")
		SendSystemMessage(pl, lang.PlayerNoInScene)
		return false
	}

	//是否副本场景
	if pls.MapTemplate().IsFuBen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("player:当前处于副本场景")
		SendSystemMessage(pl, lang.PlayerInFuBen)
		return false
	}

	if !pls.MapTemplate().IfChangeScenePvp() {
		//战斗状态
		if pl.IsPvpBattle() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:当前处于战斗状态")
			SendSystemMessage(pl, lang.PlayerInPVP)
			return false
		}
	}

	return true
}
