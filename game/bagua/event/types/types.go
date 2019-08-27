package types

import "fgame/fgame/game/player"

type BaGuaEventType string

const (
	//夫妻助战决策
	EventTypeBaGuaPairInviteDeal BaGuaEventType = "BaGuaPairInviteDeal"
	//玩家取消夫妻助战邀请
	EventTypeBaGuaPairInviteCancle BaGuaEventType = "BaGuaPairInviteCancle"
	//配偶中途退出
	EventTypeBaGuaPairSpouseExit BaGuaEventType = "BaGuaPairSpouseExit"
	//邀请过期(邀请发出 配偶下线未回复)
	EventTypeBaGuaPairNoAnswer BaGuaEventType = "BaGuaPairNoAnswer"
	//夫妻助战闯关者掉线
	EventTypeBaGuaPairInviteOffonline BaGuaEventType = "BaGuaPairInviteOffonline"
	//八卦秘境挑战结束
	EventTypeBaGuaResult BaGuaEventType = "BaGuaResult"
)

type BaGuaPairInviteDealEventData struct {
	playerId int64
	spl      player.Player
	level    int32
	agree    bool
}

func (r *BaGuaPairInviteDealEventData) GetPlayerId() int64 {
	return r.playerId
}

func (r *BaGuaPairInviteDealEventData) GetAgree() bool {
	return r.agree
}

func (r *BaGuaPairInviteDealEventData) GetSpousePlayer() player.Player {
	return r.spl
}

func (r *BaGuaPairInviteDealEventData) GetLevel() int32 {
	return r.level
}

func CreateBaGuaPairInviteDealEventData(playerId int64, spl player.Player, level int32, agree bool) *BaGuaPairInviteDealEventData {
	d := &BaGuaPairInviteDealEventData{
		playerId: playerId,
		spl:      spl,
		level:    level,
		agree:    agree,
	}
	return d
}
