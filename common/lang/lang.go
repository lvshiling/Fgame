package lang

import (
	"sync"
)

type LangCode int32

//语言代码base
const (
	UnknownBase LangCode = iota * 1000
	ExceptionBase
	GMBase
	CommonBase
	ServerBase
	AccountLoginBase
	AccountBase
	PlayerBase
	SceneBase
	BattleBase
	InventoryBase // 10
	MiscBase
	DanBase
	MountBase
	WingBase
	BodyShieldBase
	SkillBase
	TitleBase
	QuestBase
	ShopBase
	FashionBase // 20
	WeaponBase
	PkBase
	SoulBase
	RealmBase
	FriendBase
	JueXueBase
	XinFaBase
	GemBase
	SynthesisBase
	EmailBase // 30
	SoulRuinsBase
	XianfuBase
	AllianceBase
	EmperorBase
	MoonloveBase
	ActivityBase
	AllianceSceneBase
	TeamBase
	SecretCardBase
	DragonBase //40
	CrossBase
	TransportationBase
	FourGodBase
	MarryBase
	WorldBossBase
	ShenfaBase
	LingyuBase
	GoldEquipBase
	OneArenaBase
	ArenaBase
	OpenActivityBase
	MajorBase
	ReliveBase
	AnqiBase
	HuiYuanBase
	TuLongBase
	CollectBase
	VipBase
	BossTicket
	FireworksBase
	ChargeBase
	ChatBase
	LianYuBase
	NewBase
	MassacreBase
	TowerBase
	MyBossBase
	GodSiegeBase
	SystemSkillBase
	CouponBase
	UnrealBoss
	FaBaoBase
	AdditionSysBase
	XueDunBase
	MaterialBase
	XianTiBase
	LivenessBase
	OutlandBossBase
	BaGuaBase
	SongBuTingBase
	TeamCopyBase
	TianShuBase
	QuizBase
	DianXingBase
	RankBase
	WardrobeBase
	TianMoBase
	ShiHunFanBase
	GuaJiBase
	LingTongDevBase
	LingTongBase
	LingTongFashionBase
	FeiShengBase
	ShenMoBase
	ChessBase
	HongBaoBase
	SupremeTitleBase
	ShengTanBase
	EquipBaoKuBase
	MingGeBase
	TuLongEquipBase
	ZhenFaBase
	ShenQiBase
	YingLingPuBase
	BabyBase
	TradeBase
	ShenYuBase
	XianTaoBase
	SystemCompensateBase
	HouseBase
	LongGongBase
	YuXiBase
	GuideBase
	WeekBase
	FuShiBase
	QiXueBase
	ChuangShiBase
	ArenapvpBase
	JieYiBase
	FeebackBase
	WushuangBase
	DingShiBase
	XianZunCardBase
	ShangguzhilingBase
	RingBase
)

type LangService struct {
	langMap map[LangCode]string
}

func (ls *LangService) ReadLang(code LangCode) string {
	lang, ok := ls.langMap[code]
	if ok {
		return lang
	}
	return ""
}

var (
	once sync.Once
	ls   *LangService
)

func init() {
	once.Do(func() {
		ls = &LangService{
			langMap: langMap,
		}
	})
}

func GetLangService() *LangService {
	return ls
}

var (
	langMap = make(map[LangCode]string)
)

func mergeLang(mergeLangMap map[LangCode]string) {
	for lc, s := range mergeLangMap {
		langMap[lc] = s
	}
}

func GetLangMap() map[LangCode]string {
	return langMap
}
