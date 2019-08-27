package types

type MajorInvite struct {
	PlayerId   int64
	PlayerName string
	SpouseId   int64
	SpouseName string
	FuBenType  MajorType
	FuBenId    int32
	CreateTime int64
}

func NewMajorInvite(playerId int64, playerName string, spouseId int64, spouseName string, createTime int64, majorType MajorType, fubenId int32) *MajorInvite {
	data := &MajorInvite{
		PlayerId:   playerId,
		PlayerName: playerName,
		SpouseId:   spouseId,
		SpouseName: spouseName,
		FuBenType:  majorType,
		FuBenId:    fubenId,
		CreateTime: createTime,
	}
	return data
}
