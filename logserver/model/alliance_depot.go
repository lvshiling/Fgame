/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*AllianceDepot)(nil))
}

/*仙盟仓库*/
type AllianceDepot struct {
	AllianceLogMsg `bson:",inline"`

	//物品的id
	ItemId int32 `json:"itemId"`

	//变化的物品数
	ChangedNum int32 `json:"changedNum"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *AllianceDepot) LogName() string {
	return "alliance_depot"
}
