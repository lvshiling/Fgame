package types

type ChuangShiSceneEventType string

const (
	//玩家进入战场事件
	EventTypePlayerEnterChuangShiScene ChuangShiSceneEventType = "PlayerEnterChuangShiScene"
	//城门破了
	EventTypeChuangShiSceneDoorBroke = "ChuangShiSceneDoorBroke"
	//活动结束
	EventTypeChuangShiSceneFinish = "ChuangShiSceneFinish"
	//定时奖励
	EventTypeChuangShiSceneTickRew = "ChuangShiSceneTickRew"

	//正在复活点占领
	EventTypeChuangShiSceneReliveOccupying = "ChuangShiSceneReliveOccupying"
	//停止复活点占领
	EventTypeChuangShiSceneReliveOccupyStop = "ChuangShiSceneReliveOccpyStop"
	//占领复活点完成
	EventTypeChuangShiSceneReliveOccupyFinish = "ChuangShiSceneReliveOccpyFinish"
	//原地复活次数变化
	EventTypeChuangShiSceneImmediate = "ChuangShiSceneImmediate"

	//初始化防护罩
	EventTypeChuangShiSceneInitProtect = "ChuangShiSceneInitProtect"
	//初始化玉玺
	EventTypeChuangShiSceneInitYuXi = "ChuangShiSceneInitYuXi"
)

type ReliveOccupyEventData struct {
	allianceId int64
	playerId   int64
}

func (d *ReliveOccupyEventData) GetChuangShiId() int64 {
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
