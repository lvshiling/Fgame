package types

type MoonloveEventType string

const (
	//解开双人
	EventTypeMoonloveDoubleRelease MoonloveEventType = "MoonloveDoubleRelease"
	//双人
	EventTypeMoonloveDoubleCombine MoonloveEventType = "MoonloveDoubleCombine"
)

type MoonloveDoubleReleaseEventData struct {
	player1 int64
	player2 int64
}

func (d *MoonloveDoubleReleaseEventData) GetPlayer1() int64 {
	return d.player1
}

func (d *MoonloveDoubleReleaseEventData) GetPlayer2() int64 {
	return d.player2
}

func CreateMoonloveDoubleReleaseEventData(player1 int64, player2 int64) *MoonloveDoubleReleaseEventData {
	return &MoonloveDoubleReleaseEventData{
		player1: player1,
		player2: player2,
	}
}

type MoonloveDoubleCombineEventData struct {
	player1 int64
	player2 int64
}

func (d *MoonloveDoubleCombineEventData) GetPlayer1() int64 {
	return d.player1
}

func (d *MoonloveDoubleCombineEventData) GetPlayer2() int64 {
	return d.player2
}

func CreateMoonloveDoubleCombineEventData(player1 int64, player2 int64) *MoonloveDoubleCombineEventData {
	return &MoonloveDoubleCombineEventData{
		player1: player1,
		player2: player2,
	}
}
