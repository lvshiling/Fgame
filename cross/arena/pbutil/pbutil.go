package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	arenascene "fgame/fgame/cross/arena/scene"
	"fgame/fgame/cross/player/player"
	arenatypes "fgame/fgame/game/arena/types"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
)

func BuildSCArenaFourGodCancelQueue(sList []scene.Scene) *uipb.SCArenaFourGodCancelQueue {
	scArenaFourGodCancelQueue := &uipb.SCArenaFourGodCancelQueue{}
	for _, s := range sList {
		sd := s.SceneDelegate().(arenascene.FourGodSceneData)
		scArenaFourGodCancelQueue.FourGodList = append(scArenaFourGodCancelQueue.FourGodList, BuildFourGodInfo(sd))
	}
	return scArenaFourGodCancelQueue
}

func BuildSCArenaFourGodQueue(fourGodType arenatypes.FourGodType, queueNum int32) *uipb.SCArenaFourGodQueue {
	scArenaFourGodQueue := &uipb.SCArenaFourGodQueue{}
	scArenaFourGodQueue.Num = &queueNum
	fourGodTypeInt := int32(fourGodType)
	scArenaFourGodQueue.FourGodType = &fourGodTypeInt
	return scArenaFourGodQueue
}

func BuildSCArenaSelectFourGod(fourGodType arenatypes.FourGodType) *uipb.SCArenaSelectFourGod {
	fourGodTypeInt := int32(fourGodType)
	scArenaSelectFourGod := &uipb.SCArenaSelectFourGod{}
	scArenaSelectFourGod.FourGodType = &fourGodTypeInt
	return scArenaSelectFourGod
}

func BuildArenaPlayerShowData(s scene.Scene, teamPlayer *arenascene.TeamMemberObject) *uipb.ArenaPlayerShowData {
	p := s.GetPlayer(teamPlayer.GetPlayerId())
	if p == nil {
		playerShowData := &uipb.ArenaPlayerShowData{}
		playerId := teamPlayer.GetPlayerId()
		playerShowData.PlayerId = &playerId
		online := int32(0)
		playerShowData.Online = &online
		return playerShowData
	}

	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	maxHP := p.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	playerShowData.MaxHp = &maxHP
	hp := p.GetHP()
	playerShowData.Hp = &hp
	online := int32(arenascene.MemberStatusOnline)
	playerShowData.Online = &online
	remainReliveTime := teamPlayer.GetReliveTime()
	playerShowData.RemainReliveTime = &remainReliveTime

	if p.IsDead() {
		isDead := int32(1)
		playerShowData.IsDead = &isDead
		deadTime := p.GetDeadTime()
		playerShowData.DeadTime = &deadTime
	} else {
		isDead := int32(0)
		playerShowData.IsDead = &isDead
	}

	return playerShowData
}

func BuildArenaPlayer(s scene.Scene, teamPlayer *arenascene.TeamMemberObject) *uipb.ArenaPlayer {
	p := &uipb.ArenaPlayer{}
	playerId := teamPlayer.GetPlayerId()
	p.PlayerId = &playerId
	name := teamPlayer.GetName()
	p.Name = &name
	role := int32(teamPlayer.GetRole())
	p.Role = &role
	sex := int32(teamPlayer.GetSex())
	p.Sex = &sex
	force := teamPlayer.GetForce()
	p.Force = &force
	level := teamPlayer.GetLevel()
	p.Level = &level
	fashionId := teamPlayer.GetFashionId()
	p.FashionId = &fashionId
	p.PlayerShowData = BuildArenaPlayerShowData(s, teamPlayer)
	return p
}

func BuildArenaTeam(s scene.Scene, teamObj *arenascene.ArenaTeam) *uipb.ArenaTeam {
	t := &uipb.ArenaTeam{}
	teamId := teamObj.GetTeam().GetTeamId()
	t.TeamId = &teamId
	level := teamObj.GetCurrent()
	t.CurrentLevel = &level
	for _, mem := range teamObj.GetTeam().GetMemberList() {
		t.PlayerList = append(t.PlayerList, BuildArenaPlayer(s, mem))
	}
	state := int32(teamObj.GetState())
	t.State = &state
	fourGodType := int32(teamObj.GetFourGodType())
	t.FourGodType = &fourGodType
	return t
}

func BuildSCArenaSceneInfo(s arenascene.ArenaSceneData) *uipb.SCArenaSceneInfo {
	t := &uipb.SCArenaSceneInfo{}
	t.Team1 = BuildArenaTeam(s.GetScene(), s.GetTeam1())
	t.Team2 = BuildArenaTeam(s.GetScene(), s.GetTeam2())
	lastTime := s.GetLastTime()
	t.LastTime = &lastTime
	state := int32(s.GetState())
	t.State = &state
	winnerTeam := s.GetWinnerTeam()
	winnerTeamId := int64(0)
	if winnerTeam != nil {
		winnerTeamId = winnerTeam.GetTeam().GetTeamId()
	}
	t.WinnerTeamId = &winnerTeamId
	treasurePlayerId := s.GetRandomTreasureId()
	t.GetTreasurePlayerId = &treasurePlayerId
	startTime := s.GetScene().GetStartTime()
	t.StartTime = &startTime
	return t
}

var (
	scArenaNextMatch = &uipb.SCArenaNextMatch{}
)

func BuildSCArenaNextMatch() *uipb.SCArenaNextMatch {

	return scArenaNextMatch
}

var (
	scArenaNextMatchBroadcast = &uipb.SCArenaNextMatchBroadcast{}
)

func BuildSCArenaNextMatchBroadcast() *uipb.SCArenaNextMatchBroadcast {
	return scArenaNextMatchBroadcast
}

func BuildFourGodInfo(s arenascene.FourGodSceneData) *uipb.FourGodInfo {
	fourGodTypeInt := int32(s.GetFourGodType())
	fourGodInfo := &uipb.FourGodInfo{}
	fourGodInfo.FourGodType = &fourGodTypeInt
	teamNum := int32(len(s.GetCurrentTeamList()))
	fourGodInfo.TeamNum = &teamNum
	queueNum := int32(len(s.GetCurrentTeamQueue()))
	fourGodInfo.QueueNum = &queueNum
	return fourGodInfo
}
func BuildSCArenaFourGodList(sList []scene.Scene) *uipb.SCArenaFourGodList {
	scArenaFourGodList := &uipb.SCArenaFourGodList{}
	for _, s := range sList {
		sd := s.SceneDelegate().(arenascene.FourGodSceneData)
		scArenaFourGodList.FourGodList = append(scArenaFourGodList.FourGodList, BuildFourGodInfo(sd))
	}
	return scArenaFourGodList
}

var (
	scArenaSceneStart = &uipb.SCArenaSceneStart{}
)

func BuildSCArenaSceneStart() *uipb.SCArenaSceneStart {

	return scArenaSceneStart
}

func BuildSCArenaSceneEnd(winnerId int64) *uipb.SCArenaSceneEnd {
	scArenaSceneEnd := &uipb.SCArenaSceneEnd{}
	scArenaSceneEnd.WinnerTeamId = &winnerId
	return scArenaSceneEnd
}

func BuildSCArenaFourGodSceneInfo(s arenascene.FourGodSceneData, t *arenascene.ArenaTeam) *uipb.SCArenaFourGodSceneInfo {
	scArenaFourGodSceneInfo := &uipb.SCArenaFourGodSceneInfo{}
	fourGodType := int32(s.GetFourGodType())
	scArenaFourGodSceneInfo.FourGodType = &fourGodType
	teamNum := int32(len(s.GetCurrentTeamList()))
	scArenaFourGodSceneInfo.TeamNum = &teamNum
	scArenaFourGodSceneInfo.Team = BuildArenaTeam(s.GetScene(), t)
	shengShouDead := s.GetArenaBoss().IsDead()
	scArenaFourGodSceneInfo.ShengShouDead = &shengShouDead
	if shengShouDead {
		shengShouDeadTime := s.GetArenaBoss().GetDeadTime()
		scArenaFourGodSceneInfo.ShengShouDeadTime = &shengShouDeadTime
	}
	shengShouHp := s.GetArenaBoss().GetHP()
	scArenaFourGodSceneInfo.ShengShouHp = &shengShouHp
	shengShouMaxHP := s.GetArenaBoss().GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	scArenaFourGodSceneInfo.ShengShouMaxHp = &shengShouMaxHP
	// expDead := s.GetExpTree().IsDead()
	//废弃: 经验树没了
	expDead := false
	scArenaFourGodSceneInfo.ExpDead = &expDead
	expDeadTime := int64(0)
	scArenaFourGodSceneInfo.ExpDeadTime = &expDeadTime
	// }
	return scArenaFourGodSceneInfo
}

func BuildArenaTeamSimpleInfo(t *arenascene.ArenaTeam) *uipb.ArenaTeamSimpleInfo {
	arenaTeamSimpleInfo := &uipb.ArenaTeamSimpleInfo{}
	teamName := t.GetTeam().GetTeamName()
	arenaTeamSimpleInfo.TeamName = &teamName
	teamId := t.GetTeam().GetTeamId()
	arenaTeamSimpleInfo.TeamId = &teamId
	totalReliveTimes := t.GetTeam().GetTotalReliveTime()
	arenaTeamSimpleInfo.TotalReliveTimes = &totalReliveTimes

	return arenaTeamSimpleInfo
}

func BuildSCArenaFourGodTeamInfoList(s arenascene.FourGodSceneData) *uipb.SCArenaFourGodTeamInfoList {
	scArenaFourGodTeamInfoList := &uipb.SCArenaFourGodTeamInfoList{}
	fourGodType := int32(s.GetFourGodType())
	for _, t := range s.GetCurrentTeamList() {
		scArenaFourGodTeamInfoList.TeamList = append(scArenaFourGodTeamInfoList.TeamList, BuildArenaTeamSimpleInfo(t))
	}

	scArenaFourGodTeamInfoList.FourGodType = &fourGodType
	return scArenaFourGodTeamInfoList
}

func BuildSCArenaFourGodQueueChanged(queueNum int32) *uipb.SCArenaFourGodQueueChanged {
	scArenaFourGodQueueChanged := &uipb.SCArenaFourGodQueueChanged{}
	scArenaFourGodQueueChanged.Num = &queueNum
	return scArenaFourGodQueueChanged
}

func BuildSCArenaFourGodTeamChanged(num int32) *uipb.SCArenaFourGodTeamChanged {
	scArenaFourGodTeamChanged := &uipb.SCArenaFourGodTeamChanged{}
	scArenaFourGodTeamChanged.TeamNum = &num
	return scArenaFourGodTeamChanged
}

func BuildSCArenaFourGodSceneCollecting(playerId int64, collectId int64) *uipb.SCArenaFourGodSceneCollecting {
	scArenaFourGodSceneCollecting := &uipb.SCArenaFourGodSceneCollecting{}
	scArenaFourGodSceneCollecting.PlayerId = &playerId
	scArenaFourGodSceneCollecting.CollectId = &collectId
	return scArenaFourGodSceneCollecting
}

func BuildSCArenaFourGodSceneCollect(playerId int64, collectId int64) *uipb.SCArenaFourGodSceneCollect {
	scArenaFourGodSceneCollect := &uipb.SCArenaFourGodSceneCollect{}
	scArenaFourGodSceneCollect.PlayerId = &playerId
	scArenaFourGodSceneCollect.CollectId = &collectId
	return scArenaFourGodSceneCollect
}

func BuildSCArenaFourGodSceneCollectStop(playerId int64, collectId int64) *uipb.SCArenaFourGodSceneCollectStop {
	scArenaFourGodSceneCollectStop := &uipb.SCArenaFourGodSceneCollectStop{}
	scArenaFourGodSceneCollectStop.PlayerId = &playerId
	scArenaFourGodSceneCollectStop.CollectId = &collectId
	return scArenaFourGodSceneCollectStop
}

var (
	scArenaFinish = &uipb.SCArenaFinish{}
)

func BuildSCArenaFinish() *uipb.SCArenaFinish {

	return scArenaFinish
}

func BuildSCPlayerArenaData(p scene.Player) *uipb.SCPlayerArenaData {
	scPlayerArenaData := &uipb.SCPlayerArenaData{}
	playerArenaData := &uipb.PlayerArenaData{}

	reliveTime := p.GetArenaReliveTime()
	playerArenaData.ReliveTime = &reliveTime
	winTime := p.GetArenaWinTime()
	playerArenaData.WinTime = &winTime
	scPlayerArenaData.PlayerArenaData = playerArenaData
	return scPlayerArenaData
}

func BuildSCArenaPlayerDataHPChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	hp := p.GetHP()
	playerShowData.Hp = &hp
	maxHP := p.GetMaxHP()
	playerShowData.MaxHp = &maxHP
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataMaxHPChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	hp := p.GetHP()
	playerShowData.Hp = &hp
	maxHP := p.GetMaxHP()
	playerShowData.MaxHp = &maxHP
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerReliveChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	reliveTime := p.GetArenaReliveTime()

	playerShowData.RemainReliveTime = &reliveTime

	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataEnterSceneChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	online := int32(arenascene.MemberStatusOnline)
	playerShowData.Online = &online
	hp := p.GetHP()
	playerShowData.Hp = &hp
	maxHP := p.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	playerShowData.MaxHp = &maxHP
	reliveTime := p.GetArenaReliveTime()
	playerShowData.RemainReliveTime = &reliveTime
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataOnlineChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	online := int32(arenascene.MemberStatusOnline)
	playerShowData.Online = &online
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataOfflineChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	online := int32(arenascene.MemberStatusOffline)
	playerShowData.Online = &online
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataExitChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	online := int32(arenascene.MemberStatusGoAway)
	playerShowData.Online = &online
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataFailedChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	online := int32(arenascene.MemberStatusFailed)
	playerShowData.Online = &online
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataDeadChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	dead := int32(1)
	playerShowData.IsDead = &dead
	deadTime := p.GetDeadTime()
	playerShowData.DeadTime = &deadTime
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

func BuildSCArenaPlayerDataRebornChanged(p scene.Player) *uipb.SCArenaPlayerDataChanged {
	scArenaPlayerDataChanged := &uipb.SCArenaPlayerDataChanged{}
	playerShowData := &uipb.ArenaPlayerShowData{}
	playerId := p.GetId()
	playerShowData.PlayerId = &playerId
	dead := int32(0)
	playerShowData.IsDead = &dead
	scArenaPlayerDataChanged.PlayerShowData = playerShowData
	return scArenaPlayerDataChanged
}

var (
	scFourGodSceneBossDead = &uipb.SCFourGodSceneBossDead{}
)

func BuildSCFourGodSceneBossDead() *uipb.SCFourGodSceneBossDead {

	return scFourGodSceneBossDead
}

var (
	scFourGodSceneBossReborn = &uipb.SCFourGodSceneBossReborn{}
)

func BuildSCFourGodSceneBossReborn() *uipb.SCFourGodSceneBossReborn {
	return scFourGodSceneBossReborn
}

var (
	scFourGodSceneExpTreeReborn = &uipb.SCFourGodSceneExpTreeReborn{}
)

func BuildSCFourGodSceneExpTreeReborn() *uipb.SCFourGodSceneExpTreeReborn {

	return scFourGodSceneExpTreeReborn
}

var (
	scFourGodSceneExpTreeDead = &uipb.SCFourGodSceneExpTreeDead{}
)

func BuildSCFourGodSceneExpTreeDead() *uipb.SCFourGodSceneExpTreeDead {

	return scFourGodSceneExpTreeDead
}

func BuildSCFourGodSceneBossHpChanged(n scene.NPC) *uipb.SCFourGodSceneBossHpChanged {
	scFourGodSceneBossHpChanged := &uipb.SCFourGodSceneBossHpChanged{}
	hp := n.GetHP()
	scFourGodSceneBossHpChanged.Hp = &hp
	return scFourGodSceneBossHpChanged
}

func BuildSCPlayerArenaDataReliveTimeChanged(p scene.Player) *uipb.SCPlayerArenaDataChanged {
	scPlayerArenaDataChanged := &uipb.SCPlayerArenaDataChanged{}
	playerArenaData := &uipb.PlayerArenaData{}

	reliveTime := p.GetArenaReliveTime()
	playerArenaData.ReliveTime = &reliveTime
	scPlayerArenaDataChanged.PlayerArenaData = playerArenaData
	return scPlayerArenaDataChanged
}

func BuildSCPlayerArenaDataWinTimeChanged(p *player.Player) *uipb.SCPlayerArenaDataChanged {
	scPlayerArenaDataChanged := &uipb.SCPlayerArenaDataChanged{}
	playerArenaData := &uipb.PlayerArenaData{}

	winTime := p.GetArenaWinTime()
	playerArenaData.WinTime = &winTime
	scPlayerArenaDataChanged.PlayerArenaData = playerArenaData
	return scPlayerArenaDataChanged
}
