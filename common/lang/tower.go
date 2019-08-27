package lang

const (
	TowerBossDead LangCode = TowerBase + iota
	TowerBossReborn
	TowerNotOnDaBaoNotice
)

var (
	towerLangMap = map[LangCode]string{
		TowerBossDead:         "%s大杀四方，成功击杀BOSS%s",
		TowerBossReborn:       "%s出现在%s，击杀即可获得极品道具——%s",
		TowerNotOnDaBaoNotice: "您当前不处于打宝时间内，击杀怪物将不再获得任何收益！",
	}
)

func init() {
	mergeLang(towerLangMap)
}
