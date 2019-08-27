package types

type JieYiInviteData struct {
	PlayerId  int64
	JieYiId   int64
	DaoJuType JieYiDaoJuType
	Token     JieYiTokenType
	Name      string
	PlName    string
	LeaveWord string
}
