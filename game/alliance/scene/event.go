package scene

type AllianceReliveCollectEventData struct {
	playerId   int64
	allianceId int64
}

func (d *AllianceReliveCollectEventData) GetAllianceId() int64 {
	return d.allianceId
}

func (d *AllianceReliveCollectEventData) GetPlayerId() int64 {
	return d.playerId
}

func CreateAllianceReliveCollectEventData(allianceId int64, playerId int64) *AllianceReliveCollectEventData {
	eventData := &AllianceReliveCollectEventData{}
	eventData.allianceId = allianceId
	eventData.playerId = playerId
	return eventData
}

type AllianceOccupyEventData struct {
	playerId   int64
	allianceId int64
}

func (d *AllianceOccupyEventData) GetAllianceId() int64 {
	return d.allianceId
}

func (d *AllianceOccupyEventData) GetPlayerId() int64 {
	return d.playerId
}

func CreateAllianceOccupyEventData(allianceId int64, playerId int64) *AllianceOccupyEventData {
	eventData := &AllianceOccupyEventData{}
	eventData.allianceId = allianceId
	eventData.playerId = playerId
	return eventData
}
