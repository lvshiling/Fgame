package types

type BaGuaInvite struct {
	PlayerId   int64
	PlayerName string
	SpouseId   int64
	SpouseName string
	Level      int32
	CreateTime int64
}

func NewBaGuaInvite(playerId int64, playerName string, spouseId int64, spouseName string, level int32, createTime int64) *BaGuaInvite {
	data := &BaGuaInvite{
		PlayerId:   playerId,
		PlayerName: playerName,
		SpouseId:   spouseId,
		SpouseName: spouseName,
		Level:      level,
		CreateTime: createTime,
	}
	return data
}

type BaGuaPair struct {
	playerId int64
	spouseId int64
}

func (r *BaGuaPair) GetPlayerId() int64 {
	return r.playerId
}

func (r *BaGuaPair) GetSpouseId() int64 {
	return r.spouseId
}

func NewBaGuaPair(playerId int64, spouseId int64) *BaGuaPair {
	data := &BaGuaPair{
		playerId: playerId,
		spouseId: spouseId,
	}
	return data
}
