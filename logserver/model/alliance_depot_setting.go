/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*AllianceDepotSetting)(nil))
}

/*仙盟仓库设置*/
type AllianceDepotSetting struct {
	AllianceLogMsg `bson:",inline"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *AllianceDepotSetting) LogName() string {
	return "alliance_depot_setting"
}
