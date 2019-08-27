/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerJieYiToken)(nil))
}

/*结义信物改变*/
type PlayerJieYiToken struct {
	PlayerLogMsg `bson:",inline"`

	//之前信物类型
	BeforeToken string `json:"beforeToken"`

	//当前信物类型
	CurToken string `json:"curToken"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerJieYiToken) LogName() string {
	return "player_jieyi_token"
}
