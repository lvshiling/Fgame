package lang

const (
	WeaponRepeatActive LangCode = WeaponBase + iota
	WeaponNotActiveNotEat
	WeaponEatDanReachedLimit
	WeaponNotActiveNotUpstar
	WeaponReacheFullStar
	WeaponUpstarFailure
	WeaponNotActiveNotAwaken
	WeaponRepeatAwaken
	WeaponAwakenNotStar
	WeaponNotActiveNotWear
	WeaponActivateNotice
	WeaponAwakenNotice
	WeaponActivateFail
)

var (
	weaponLangMap = map[LangCode]string{
		WeaponRepeatActive:       "兵魂已激活,无需激活",
		WeaponNotActiveNotEat:    "未激活的兵魂,无法食培养丹",
		WeaponEatDanReachedLimit: "食丹等级已达最大",
		WeaponNotActiveNotUpstar: "未激活的兵魂,无法升星",
		WeaponReacheFullStar:     "兵魂已满星",
		WeaponUpstarFailure:      "升星失败",
		WeaponNotActiveNotAwaken: "未激活的兵魂,无法觉醒",
		WeaponRepeatAwaken:       "兵魂重复觉醒",
		WeaponAwakenNotStar:      "星数不足,无法觉醒",
		WeaponNotActiveNotWear:   "未激活的兵魂,无法穿戴",
		WeaponActivateNotice:     "神兵天降，%s成功激活%s，战力飙升%s",
		WeaponAwakenNotice:       "神兵觉醒，%s成功将%s炼化光武，战力飙升%s，混元伤害翻倍！",
		WeaponActivateFail:       "兵魂激活失败，该兵魂不能手动激活",
	}
)

func init() {
	mergeLang(weaponLangMap)
}
