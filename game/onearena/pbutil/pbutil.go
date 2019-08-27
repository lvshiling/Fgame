package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/onearena/onearena"
	playeronearena "fgame/fgame/game/onearena/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

func BuildSCOneArenaGet(pl player.Player, oneArenaList []*onearena.OneArenaObject) *uipb.SCOneArenaGet {
	oneArenaGet := &uipb.SCOneArenaGet{}
	manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	for _, oneArenaObj := range oneArenaList {
		robTiem := int64(0)
		level := oneArenaObj.Level
		pos := oneArenaObj.Pos

		oneArenaRecord := manager.GetOneArenaRecord(level, pos)
		if oneArenaRecord != nil {
			robTiem = oneArenaRecord.RobTime
		}
		oneArenaGet.OneArenaList = append(oneArenaGet.OneArenaList, buildOneArena(oneArenaObj, robTiem))
	}

	oneArenaObj := manager.GetOneArena()
	oneArenaGet.KunSell = buildKunSell(oneArenaObj.KunSilver, oneArenaObj.KunBindGold)
	return oneArenaGet
}

func BuildSCOneArenaRob(result int32) *uipb.SCOneArenaRob {
	oneArenaRob := &uipb.SCOneArenaRob{}
	oneArenaRob.Result = &result
	return oneArenaRob
}

func BuildSCOneArenaRobbedPush(name string) *uipb.SCOneArenaRobbedPush {
	oneArenaRobbedPush := &uipb.SCOneArenaRobbedPush{}
	oneArenaRobbedPush.Name = &name
	return oneArenaRobbedPush
}

func BuildSCOneArenaSell(totalSilver int64, totalBindGold int64) *uipb.SCOneArenaSell {
	oneArenaSell := &uipb.SCOneArenaSell{}
	oneArenaSell.KunSell = buildKunSell(totalSilver, totalBindGold)
	return oneArenaSell
}

func BuilSCOneArenaRecord(recordList []*playeronearena.PlayerOneArenaRobbedObject) *uipb.SCOneArenaRecord {
	oneArenaRecord := &uipb.SCOneArenaRecord{}

	for _, record := range recordList {
		oneArenaRecord.RecordList = append(oneArenaRecord.RecordList, buildRecord(record))
	}

	return oneArenaRecord
}

func buildKunSell(totalSilver int64, totalBindGold int64) *uipb.OneArenaKunSell {
	oneArenaKunSell := &uipb.OneArenaKunSell{}
	oneArenaKunSell.TotalSilver = &totalSilver
	oneArenaKunSell.TotalBindGold = &totalBindGold
	return oneArenaKunSell
}

func buildRecord(record *playeronearena.PlayerOneArenaRobbedObject) *uipb.OneArenaRecord {
	oneArenaRecord := &uipb.OneArenaRecord{}

	robName := record.RobName
	robTime := record.RobTime

	sucess := true
	if record.Status == 2 {
		sucess = false
	}

	oneArenaRecord.RobName = &robName
	oneArenaRecord.RobTime = &robTime
	oneArenaRecord.Status = &sucess

	return oneArenaRecord
}

func BuildSCOneArenaRobResult(sucess bool, name string, level int32) *uipb.SCOneArenaRobResult {
	oneArenaRobResult := &uipb.SCOneArenaRobResult{}
	oneArenaRobResult.Sucess = &sucess
	oneArenaRobResult.Name = &name
	oneArenaRobResult.Level = &level
	return oneArenaRobResult
}

func BuildSCOneArenaRobot(spl scene.Player) *uipb.SCOneArenaRobot {
	scOneArenaRobot := &uipb.SCOneArenaRobot{}

	playerId := spl.GetId()
	name := spl.GetName()
	level := spl.GetLevel()
	hp := spl.GetHP()
	force := spl.GetForce()
	role := int32(spl.GetRole())
	sex := int32(spl.GetSex())

	scOneArenaRobot.PlayerId = &playerId
	scOneArenaRobot.Name = &name
	scOneArenaRobot.Level = &level
	scOneArenaRobot.Hp = &hp
	scOneArenaRobot.Force = &force
	scOneArenaRobot.Role = &role
	scOneArenaRobot.Sex = &sex
	return scOneArenaRobot

}

func buildOneArena(oneArenaObj *onearena.OneArenaObject, robTime int64) *uipb.OneArenaInfo {
	oneArenaInfo := &uipb.OneArenaInfo{}
	level := int32(oneArenaObj.Level)
	pos := oneArenaObj.Pos
	name := oneArenaObj.OwnerName
	isRobbing := oneArenaObj.IsRobbing
	lastTime := oneArenaObj.LastTime
	force := int64(-1)
	isVip := false
	if oneArenaObj.OwnerId != 0 {
		playerInfo, _ := player.GetPlayerService().GetPlayerInfo(oneArenaObj.OwnerId)
		if playerInfo != nil {
			force = playerInfo.Force
			vip := playerInfo.VipInfo.VipLevel
			if vip > 0 {
				isVip = true
			}
		}
	}
	oneArenaInfo.Level = &level
	oneArenaInfo.Pos = &pos
	oneArenaInfo.Name = &name
	oneArenaInfo.IsVip = &isVip
	oneArenaInfo.IsRobbing = &isRobbing
	oneArenaInfo.RobTime = &robTime
	oneArenaInfo.LastTime = &lastTime
	oneArenaInfo.Force = &force
	ownerId := oneArenaObj.OwnerId
	oneArenaInfo.OwnerId = &ownerId
	return oneArenaInfo
}
