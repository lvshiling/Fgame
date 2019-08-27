package lang

const (
	BossTicketPlayerNoInPVPMap LangCode = BossTicket + iota
	BossTicketPlayerNoInWorld
	BossTicketBossBornNotice
)

var bossCallLangMap = map[LangCode]string{
	BossTicketPlayerNoInPVPMap: "新手村无法使用BOSS召唤券",
	BossTicketPlayerNoInWorld:  "玩家当前不在世界场景，无法使用BOSS召唤券",
	BossTicketBossBornNotice:   "%s吹响BOSS号角，召唤出了上古魔头%s（掉落：绝版时装，珍贵商城道具）— %s",
}

func init() {
	mergeLang(bossCallLangMap)
}
