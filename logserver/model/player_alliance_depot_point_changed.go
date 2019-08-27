/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerAllianceDepotPointChanged)(nil))
}

/*玩家仙盟仓库积分变化*/
type PlayerAllianceDepotPointChanged struct {
	PlayerLogMsg `bson:",inline"`

	//变化积分
	ChangedPoint int32 `json:"changedPoint"`

	//变化前的积分
	BeforePoint int32 `json:"beforePoint"`

	//当前的积分
	CurPoint int32 `json:"curPoint"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerAllianceDepotPointChanged) LogName() string {
	return "player_alliance_depot_point_changed"
}
