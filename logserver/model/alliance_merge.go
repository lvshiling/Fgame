/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*AllianceMerge)(nil))
}

/*仙盟合并*/
type AllianceMerge struct {
	AllianceLogMsg `bson:",inline"`

	//被邀请仙盟id
	InviteAllianceId int64 `json:"inviteAllianceId"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *AllianceMerge) LogName() string {
	return "alliance_merge"
}
