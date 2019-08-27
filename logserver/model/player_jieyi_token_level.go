/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerJieYiTokenLevel)(nil))
}

/*结义信物等级改变*/
type PlayerJieYiTokenLevel struct {
	PlayerLogMsg `bson:",inline"`

	//之前信物等级
	BeforeTokenLevel int32 `json:"beforeTokenLevel"`

	//当前信物等级
	CurTokenLevel int32 `json:"curTokenLevel"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerJieYiTokenLevel) LogName() string {
	return "player_jieyi_token_level"
}
