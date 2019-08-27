package marry

import (
	"fgame/fgame/common/lang"
	gamecommon "fgame/fgame/game/common/common"
)

var (
	//求婚者已下线
	ErrorMarryProposalIsNoOnline = gamecommon.CodeError(lang.MarryProposalIsNoOnline)
	//决策者已婚
	ErrorMarryDealIsMarried = gamecommon.CodeError(lang.MarryDealIsMarried)
	//离婚请求对方下线
	ErrorMarryDivorceNoOnline = gamecommon.CodeError(lang.MarryDivorceNoOnline)
	//有一方性别发生变化
	ErrorMarryDealIsSexChanged = gamecommon.CodeError(lang.MarryDealIsSexChanged)
	//本场次婚礼已被预定
	ErrorMarryWedIsBeScheduled = gamecommon.CodeError(lang.MarryWedIsBeScheduled)
	//您目前处于离婚状态
	ErrorMarryReserveIsDivorced = gamecommon.CodeError(lang.MarryReserveIsDivorced)
	//您目前不是以后状态
	ErrorMarryReserveNoMarried = gamecommon.CodeError(lang.MarryReserveNoMarried)
	//您的爱人目前不在线
	ErrorMarryPreWedSpouseNoOnline = gamecommon.CodeError(lang.MarryPreWedSpouseNoOnline)
	//预定婚期请求已失效
	ErrorMarryPreWedIsOverdue = gamecommon.CodeError(lang.MarryPreWedIsOverdue)
)
