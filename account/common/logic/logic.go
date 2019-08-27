package logic

import (
	accountsession "fgame/fgame/account/session"
	"fgame/fgame/common/exception"
	"fgame/fgame/common/lang"
	"fgame/fgame/game/common/pbutil"
)

func SendSessionSystemMessage(gs accountsession.Session, code lang.LangCode, args ...string) {
	content := lang.GetLangService().ReadLang(code)
	scSystemMessage := pbutil.BuildSCSystemMessage(content, args...)
	gs.Send(scSystemMessage)
	return
}

func SendSessionExceptionMessage(gs accountsession.Session, code exception.ExceptionCode) {
	content := lang.GetLangService().ReadLang(code.LangCode())
	scException := pbutil.BuildSCException(content, code)
	gs.Send(scException)
	return
}
