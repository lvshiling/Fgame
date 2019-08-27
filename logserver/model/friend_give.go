/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*FriendGive)(nil))
}

/*赠送好友装备*/
type FriendGive struct {
	PlayerLogMsg `bson:",inline"`

	//变更原因编号
	Reason int32 `json:"reason"`

	//变更原因
	ReasonText string `json:"reasonText"`
}

func (c *FriendGive) LogName() string {
	return "friend_give"
}
