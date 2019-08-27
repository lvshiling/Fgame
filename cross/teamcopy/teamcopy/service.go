package teamcopy

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/cross/player/player"
	teamcopyeventtypes "fgame/fgame/cross/teamcopy/event/types"
	teamscene "fgame/fgame/cross/teamcopy/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type TeamCopyService interface {
	Heartbeat()
	//开始战斗
	TeamCopyStartBattle(pl *player.Player, playerList []*BattleTeamMember) bool
	//获取组队副本场景数据
	GetTeamCopySceneData(pl *player.Player) teamscene.TeamCopySceneData
	//获取组队副本数据
	GetTeamCopyDataByPlayerId(playerId int64) *teamCopyData
	//场景完成
	TeamCopyFinish(sd teamscene.TeamCopySceneData)

	//成员上线
	TeamCopyMemberOnline(pl scene.Player)
	//成员下线
	TeamCopyMemberOffline(pl scene.Player)
	//成员放弃
	TeamCopyMemeberGiveUp(pl scene.Player)
	//成员退出
	TeamCopyMemeberExit(pl scene.Player)
}

type teamCopyService struct {
	rwm             sync.RWMutex
	teamCopyDataMap map[int64]*teamCopyData
	playerTeamMap   map[int64]*teamscene.TeamObject
}

func (s *teamCopyService) init() (err error) {
	s.teamCopyDataMap = make(map[int64]*teamCopyData)
	s.playerTeamMap = make(map[int64]*teamscene.TeamObject)
	return nil
}

func (s *teamCopyService) TeamCopyStartBattle(pl *player.Player, playerList []*BattleTeamMember) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if len(playerList) == 0 {
		log.WithFields(
			log.Fields{}).Warn("teamcopy:成员是空的")
		return false
	}

	for _, mem := range playerList {
		teamId := s.getPlayerTeamdId(mem.GetPlayerId())
		if teamId != 0 {
			log.WithFields(
				log.Fields{
					"teamId": teamId,
				}).Warn("teamcopy:组队副本已存在")
			return false
		}
		if !mem.GetTeamPurpose().Vaild() {
			log.WithFields(
				log.Fields{
					"teamId":      teamId,
					"teamPurpose": mem.GetTeamPurpose(),
				}).Warn("teamcopy:组队副本标识错误")
			return false
		}
	}

	teamId := s.getPlayerTeamdId(pl.GetId())
	if teamId != 0 {
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Warn("teamcopy:组队副本已存在")
		return false
	}

	//判断场景是否存在了
	sceneData := s.getTeamCopyData(teamId)
	if sceneData != nil {
		log.WithFields(
			log.Fields{
				"teamId": teamId,
			}).Warn("teamcopy:组队副本已存在")
		return false
	}

	//创建队伍
	teamObj := createTeamCopyTeamWithRobot(playerList)
	sceneData = s.createTeamCopySceneData(teamObj)
	s.teamInit(teamObj, sceneData)
	return true
}

func (s *teamCopyService) createTeamCopySceneData(teamObj *teamscene.TeamObject) (data *teamCopyData) {
	teamCopyData := s.getTeamCopyData(teamObj.GetTeamId())
	if teamCopyData != nil {
		return
	}
	data = createSceneData(teamObj)
	return
}

func (s *teamCopyService) teamInit(teamObj *teamscene.TeamObject, data *teamCopyData) {
	s.playerTeamInit(teamObj)
	s.teamCopyDataMap[teamObj.GetTeamId()] = data
}

func (s *teamCopyService) playerTeamInit(teamObj *teamscene.TeamObject) {
	for _, mem := range teamObj.GetMemberList() {
		s.playerTeamMap[mem.GetPlayerId()] = teamObj
	}
}

func (s *teamCopyService) getPlayerTeamdId(playerId int64) (teamId int64) {
	teamObj, ok := s.playerTeamMap[playerId]
	if !ok {
		return
	}
	return teamObj.GetTeamId()
}

func (s *teamCopyService) getTeamCopyData(teamId int64) *teamCopyData {
	teamCopyData, ok := s.teamCopyDataMap[teamId]
	if !ok {
		return nil
	}
	return teamCopyData
}

func (s *teamCopyService) GetTeamCopySceneData(pl *player.Player) (sceneData teamscene.TeamCopySceneData) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	teamId := s.getPlayerTeamdId(pl.GetId())
	if teamId == 0 {
		return nil
	}
	teamCopyData := s.getTeamCopyData(teamId)
	if teamCopyData == nil {
		return
	}
	return teamCopyData.sceneData
}

func (s *teamCopyService) GetTeamCopyDataByPlayerId(playerId int64) (data *teamCopyData) {
	teamId := s.getPlayerTeamdId(playerId)
	if teamId == 0 {
		return nil
	}
	data = s.getTeamCopyData(teamId)
	return
}

func (s *teamCopyService) removeTeamId(teamId int64) {
	teamCopyData := s.getTeamCopyData(teamId)
	if teamCopyData == nil {
		return
	}
	teamCopyData.init()
	delete(s.teamCopyDataMap, teamId)
}

func (s *teamCopyService) removeTeamPlayer(teamObj *teamscene.TeamObject) {
	teamCopyData := s.getTeamCopyData(teamObj.GetTeamId())
	if teamCopyData == nil {
		return
	}
	for _, mem := range teamObj.GetMemberList() {
		delete(s.playerTeamMap, mem.GetPlayerId())
	}
	return
}

func (s *teamCopyService) removePlayer(p scene.Player) {
	delete(s.playerTeamMap, p.GetId())
}

//场景完成
func (s *teamCopyService) TeamCopyFinish(sd teamscene.TeamCopySceneData) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	teamObj := sd.GetTeamObj()
	s.removeTeamPlayer(teamObj)
	s.removeTeamId(teamObj.GetTeamId())
}

//成员上线
func (s *teamCopyService) TeamCopyMemberOnline(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	teamId := s.getPlayerTeamdId(playerId)
	if teamId == 0 {
		return
	}
	sd := s.getTeamCopyData(teamId)
	if sd == nil {
		return
	}
	mem := sd.sceneData.GetMember(pl)
	if mem == nil {
		return
	}
	if mem.GetStatus() != teamscene.MemberStatusOffline {
		return
	}
	mem.SetStatus(teamscene.MemberStatusOnline)
	//玩家进场景去广播
	//gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyMemberOnline, pl, sd.sceneData)
}

//成员下线
func (s *teamCopyService) TeamCopyMemberOffline(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	teamId := s.getPlayerTeamdId(playerId)
	if teamId == 0 {
		return
	}
	sd := s.getTeamCopyData(teamId)
	if sd == nil {
		return
	}
	mem := sd.sceneData.GetMember(pl)
	if mem == nil {
		return
	}
	if mem.GetStatus() != teamscene.MemberStatusOnline {
		return
	}
	mem.SetStatus(teamscene.MemberStatusOffline)
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyMemberOffline, pl, sd.sceneData)
	s.checkTeamCopy(sd.sceneData)
}

//成员退出
func (s *teamCopyService) TeamCopyMemeberExit(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	teamId := s.getPlayerTeamdId(playerId)
	if teamId == 0 {
		return
	}
	sd := s.getTeamCopyData(teamId)
	if sd == nil {
		return
	}
	mem := sd.sceneData.GetMember(pl)
	if mem == nil {
		return
	}
	if mem.GetStatus() == teamscene.MemberStatusGoAway {
		return
	}
	//移除玩家
	s.removePlayer(pl)
	mem.SetStatus(teamscene.MemberStatusGoAway)
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyMemberExit, pl, sd.sceneData)
	s.checkTeamCopy(sd.sceneData)
}

//成员放弃
func (s *teamCopyService) TeamCopyMemeberGiveUp(pl scene.Player) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	playerId := pl.GetId()
	teamId := s.getPlayerTeamdId(playerId)
	if teamId == 0 {
		return
	}
	sd := s.getTeamCopyData(teamId)
	if sd == nil {
		return
	}
	mem := sd.sceneData.GetMember(pl)
	if mem == nil {
		return
	}
	if mem.GetStatus() == teamscene.MemberStatusFailed {
		return
	}
	//移除玩家
	s.removePlayer(pl)
	mem.SetStatus(teamscene.MemberStatusFailed)
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyMemberGiveUp, pl, sd.sceneData)
	s.checkTeamCopy(sd.sceneData)
}

func (s *teamCopyService) checkTeamCopy(sd teamscene.TeamCopySceneData) {
	if sd.IfAllLevel() {
		teamObj := sd.GetTeamObj()
		s.removeTeamPlayer(teamObj)
		s.removeTeamId(teamObj.GetTeamId())

		//所有玩家退出或放弃
		teamCopyEnd(sd)
	}
}

//结束场景
func teamCopyEnd(sd teamscene.TeamCopySceneData) {
	s := sd.GetScene()
	ctx := scene.WithScene(context.Background(), s)
	s.Post(message.NewScheduleMessage(onTeamCopyEnd, ctx, nil, nil))
}

func onTeamCopyEnd(ctx context.Context, result interface{}, err error) error {
	teamScene := scene.SceneInContext(ctx)
	sd := teamScene.SceneDelegate().(teamscene.TeamCopySceneData)
	sd.FinishScene()
	return nil
}

func (s *teamCopyService) Heartbeat() {

}

var (
	once sync.Once
	cs   *teamCopyService
)

func Init() (err error) {
	once.Do(func() {
		cs = &teamCopyService{}
		err = cs.init()
	})
	return err
}

func GetTeamCopyService() TeamCopyService {
	return cs
}
