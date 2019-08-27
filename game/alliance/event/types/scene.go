package types

type AllianceSceneEventType string

const (
	//玩家进入战场事件
	EventTypePlayerEnterAllianceScene AllianceSceneEventType = "PlayerEnterAllianceScene"
	//城门破了
	EventTypeAllianceSceneDoorBroke = "AllianceSceneDoorBroke"
	//活动结束
	EventTypeAllianceSceneFinish = "AllianceSceneFinish"
	//防守虎符改变
	EventTypeAllianceSceneDefendHuFuChanged = "AllianceSceneDefendHuFuChanged"
	//仙盟虎符改变
	EventTypeAllianceSceneHuFuChanged = "AllianceSceneHuFuChanged"
	//城战定时奖励
	EventTypeAllianceSceneTickRew = "AllianceSceneTickRew"

	//正在复活点占领
	EventTypeAllianceSceneReliveOccupying = "AllianceSceneReliveOccupying"
	//停止复活点占领
	EventTypeAllianceSceneReliveOccupyStop = "AllianceSceneReliveOccpyStop"
	//占领复活点完成
	EventTypeAllianceSceneReliveOccupyFinish = "AllianceSceneReliveOccpyFinish"
	//原地复活次数变化
	EventTypeAllianceSceneImmediate = "AllianceSceneImmediate"

	//初始化防护罩
	EventTypeAllianceSceneInitProtect = "AllianceSceneInitProtect"
	//初始化玉玺
	EventTypeAllianceSceneInitYuXi = "AllianceSceneInitYuXi"
)

type AllianceSceneHuFuChangedEventData struct {
	allianceId int64
	huFu       int64
}

func (d *AllianceSceneHuFuChangedEventData) GetAllianceId() int64 {
	return d.allianceId
}

func (d *AllianceSceneHuFuChangedEventData) GetHuFu() int64 {
	return d.huFu
}

func CreateAllianceSceneHuFuChangedEvent(allianceId int64, huFu int64) *AllianceSceneHuFuChangedEventData {
	d := &AllianceSceneHuFuChangedEventData{
		allianceId: allianceId,
		huFu:       huFu,
	}
	return d
}

type ReliveOccupyEventData struct {
	allianceId int64
	playerId   int64
}

func (d *ReliveOccupyEventData) GetAllianceId() int64 {
	return d.allianceId
}
func (d *ReliveOccupyEventData) GetPlayerId() int64 {
	return d.playerId
}

func CreateReliveOccupyEvent(allianceId int64, playerId int64) *ReliveOccupyEventData {
	d := &ReliveOccupyEventData{
		allianceId: allianceId,
		playerId:   playerId,
	}
	return d
}
