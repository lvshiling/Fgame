package types

import "fgame/fgame/game/player"

type RealmEventType string

const (
	//夫妻助战决策
	EventTypeRealmPairInviteDeal RealmEventType = "RealmPairInviteDeal"
	//玩家取消夫妻助战邀请
	EventTypeRealmPairInviteCancle RealmEventType = "RealmPairInviteCancle"
	//配偶中途退出
	EventTypeRealmPairSpouseExit RealmEventType = "RealmPairSpouseExit"
	//邀请过期(邀请发出 配偶下线未回复)
	EventTypeRealmPairNoAnswer RealmEventType = "RealmPairNoAnswer"
	//夫妻助战闯关者掉线
	EventTypeRealmPairInviteOffonline RealmEventType = "RealmPairInviteOffonline"
	//天劫塔挑战结束
	EventTypeRealmResult RealmEventType = "RealmResult"
	//天劫塔排行榜第一变化
	//EventTypeRealmFirstChange RealmEventType = "RealmFirstChange"
)

type RealmPairInviteDealEventData struct {
	playerId int64
	spl      player.Player
	level    int32
	agree    bool
}

func (r *RealmPairInviteDealEventData) GetPlayerId() int64 {
	return r.playerId
}

func (r *RealmPairInviteDealEventData) GetSpousePlayer() player.Player {
	return r.spl
}

func (r *RealmPairInviteDealEventData) GetAgree() bool {
	return r.agree
}

func (r *RealmPairInviteDealEventData) GetLevel() int32 {
	return r.level
}

func CreateRealmPairInviteDealEventData(playerId int64, spl player.Player, level int32, agree bool) *RealmPairInviteDealEventData {
	d := &RealmPairInviteDealEventData{
		playerId: playerId,
		level:    level,
		spl:      spl,
		agree:    agree,
	}
	return d
}
