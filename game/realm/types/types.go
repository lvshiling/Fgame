package types

type RealmResultType int32

const (
	//决策同意
	RealmResultTypeOk RealmResultType = 1 + iota
	//决策拒绝
	RealmResultTypeNo
)

const (
	levelBit   = 54
	curTimeMax = 1<<54 - 1
)

func Resolve(val int64) (level int32, curTime int64) {
	level = int32(val >> levelBit)
	curTime = curTimeMax - val&curTimeMax
	return
}

func Combine(level int32, curTime int64) (val int64) {
	levelValue := int64(level) << levelBit
	curTimeValue := curTimeMax - curTime
	return levelValue + curTimeValue
}

type RealmRankData struct {
	Name  string
	Level int32
}

func NewRealmRankData(name string, level int32) *RealmRankData {
	data := &RealmRankData{
		Name:  name,
		Level: level,
	}
	return data
}

type RealmInvite struct {
	PlayerId   int64
	PlayerName string
	SpouseId   int64
	SpouseName string
	Level      int32
	CreateTime int64
}

func NewRealmInvite(playerId int64, playerName string, spouseId int64, spouseName string, level int32, createTime int64) *RealmInvite {
	data := &RealmInvite{
		PlayerId:   playerId,
		PlayerName: playerName,
		SpouseId:   spouseId,
		SpouseName: spouseName,
		Level:      level,
		CreateTime: createTime,
	}
	return data
}

type RealmPair struct {
	playerId int64
	spouseId int64
}

func (r *RealmPair) GetPlayerId() int64 {
	return r.playerId
}

func (r *RealmPair) GetSpouseId() int64 {
	return r.spouseId
}

func NewRealmPair(playerId int64, spouseId int64) *RealmPair {
	data := &RealmPair{
		playerId: playerId,
		spouseId: spouseId,
	}
	return data
}
