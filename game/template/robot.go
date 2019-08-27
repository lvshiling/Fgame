package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	coreutils "fgame/fgame/core/utils"
	playertypes "fgame/fgame/game/player/types"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"
)

//机器人
type RobotTemplate struct {
	*RobotTemplateVO
	professionList             []int32
	professionListWeights      []int64
	genderList                 []int32
	genderListWeights          []int64
	fashionList                []int32
	fashionListWeights         []int64
	weaponList                 []int32
	weaponListWeights          []int64
	wingList                   []int32
	wingListWeights            []int64
	mountList                  []int32
	mountListWeights           []int64
	fabaoList                  []int32
	fabaoListWeights           []int64
	shenfaList                 []int32
	shenfaListWeights          []int64
	xiantiList                 []int32
	xiantiListWeights          []int64
	fieldList                  []int32
	fieldListWeights           []int64
	titleList                  []int32
	titleListWeights           []int64
	jueXueList                 []int32
	jueXueListWeights          []int64
	soulList                   []int32
	soulListWeights            []int64
	lingTongList               []int32
	lingTongListWeights        []int64
	lingTongFashionList        []int32
	lingTongFashionListWeights []int64
	lingTongWeaponList         []int32
	lingTongWeaponListWeights  []int64
	lingTongWingList           []int32
	lingTongWingListWeights    []int64
	lingTongMountList          []int32
	lingTongMountListWeights   []int64
	lingTongFabaoList          []int32
	lingTongFabaoListWeights   []int64
	lingTongShenfaList         []int32
	lingTongShenfaListWeights  []int64
	lingTongXiantiList         []int32
	lingTongXiantiListWeights  []int64
	lingTongLingyuList         []int32
	lingTongLingyuListWeights  []int64
}

func (t *RobotTemplate) TemplateId() int {
	return t.Id
}

func (t *RobotTemplate) RandomLevel() int32 {
	return int32(mathutils.RandomRange(int(t.LevMin), int(t.LevMax)))
}

func (t *RobotTemplate) RandomRole() playertypes.RoleType {
	index := mathutils.RandomWeights(t.professionListWeights)
	return playertypes.RoleType(t.professionList[index])
}

func (t *RobotTemplate) RandomSex() playertypes.SexType {
	index := mathutils.RandomWeights(t.genderListWeights)
	return playertypes.SexType(t.genderList[index])
}
func (t *RobotTemplate) RandomFashion() int32 {
	index := mathutils.RandomWeights(t.fashionListWeights)
	return t.fashionList[index]
}

func (t *RobotTemplate) RandomWeapon() int32 {
	index := mathutils.RandomWeights(t.weaponListWeights)
	return t.weaponList[index]
}

func (t *RobotTemplate) RandomWing() int32 {
	if len(t.wingListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.wingListWeights)
	return t.wingList[index]
}

func (t *RobotTemplate) RandomMount() int32 {
	if len(t.mountListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.mountListWeights)
	return t.mountList[index]
}

func (t *RobotTemplate) RandomFabao() int32 {
	if len(t.fabaoListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.fabaoListWeights)
	return t.fabaoList[index]
}

func (t *RobotTemplate) RandomShenfa() int32 {
	if len(t.shenfaListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.shenfaListWeights)
	return t.shenfaList[index]
}

func (t *RobotTemplate) RandomXianti() int32 {
	if len(t.xiantiListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.xiantiListWeights)
	return t.xiantiList[index]
}

func (t *RobotTemplate) RandomField() int32 {
	if len(t.fieldListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.fieldListWeights)
	return t.fieldList[index]
}
func (t *RobotTemplate) RandomTitle() int32 {
	if len(t.titleListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.titleListWeights)
	return t.titleList[index]
}

func (t *RobotTemplate) RandomJueXue() int32 {
	if len(t.jueXueListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.jueXueListWeights)
	return t.jueXueList[index]
}

func (t *RobotTemplate) RandomSoul() int32 {
	if len(t.soulListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.soulListWeights)
	return t.soulList[index]
}

func (t *RobotTemplate) RandomLingTong() int32 {
	if len(t.lingTongListWeights) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongListWeights)
	return t.lingTongList[index]
}

func (t *RobotTemplate) RandomLingTongFashion() int32 {
	if len(t.lingTongFashionList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongFashionListWeights)
	return t.lingTongFashionList[index]
}
func (t *RobotTemplate) RandomLingTongWeapon() int32 {
	if len(t.lingTongWeaponList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongWeaponListWeights)
	return t.lingTongWeaponList[index]
}

func (t *RobotTemplate) RandomLingTongWing() int32 {
	if len(t.lingTongWingList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongWingListWeights)
	return t.lingTongWingList[index]
}

func (t *RobotTemplate) RandomLingTongMount() int32 {
	if len(t.lingTongMountList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongMountListWeights)
	return t.lingTongMountList[index]
}

func (t *RobotTemplate) RandomLingTongFabao() int32 {
	if len(t.lingTongFabaoList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongFabaoListWeights)
	return t.lingTongFabaoList[index]
}

func (t *RobotTemplate) RandomLingTongShenfa() int32 {
	if len(t.lingTongShenfaList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongShenfaListWeights)
	return t.lingTongShenfaList[index]
}

func (t *RobotTemplate) RandomLingTongXianti() int32 {
	if len(t.lingTongXiantiList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongXiantiListWeights)
	return t.lingTongXiantiList[index]
}

func (t *RobotTemplate) RandomLingTongLingyu() int32 {
	if len(t.lingTongLingyuList) == 0 {
		return 0
	}
	index := mathutils.RandomWeights(t.lingTongLingyuListWeights)
	return t.lingTongLingyuList[index]
}

func (t *RobotTemplate) PatchAfterCheck() {
}

func (t *RobotTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	tempProfessionList, err := coreutils.SplitAsIntArray(t.NeedProfession)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.NeedProfession)
		return template.NewTemplateFieldError("NeedProfession", err)
	}

	t.professionList = tempProfessionList
	for _, _ = range t.professionList {
		t.professionListWeights = append(t.professionListWeights, 1)
	}
	tempGenderList, err := coreutils.SplitAsIntArray(t.NeedGender)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.NeedGender)
		return template.NewTemplateFieldError("NeedGender", err)
	}
	t.genderList = tempGenderList
	for _, _ = range t.genderList {
		t.genderListWeights = append(t.genderListWeights, 1)
	}

	tempFashionList, err := coreutils.SplitAsIntArray(t.FashionId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FashionId)
		return template.NewTemplateFieldError("FashionId", err)
	}
	t.fashionList = tempFashionList
	for _, _ = range t.fashionList {
		t.fashionListWeights = append(t.fashionListWeights, 1)
	}
	tempWeaponList, err := coreutils.SplitAsIntArray(t.WeaponId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.WeaponId)
		return template.NewTemplateFieldError("WeaponId", err)
	}
	t.weaponList = tempWeaponList
	for _, _ = range t.weaponList {
		t.weaponListWeights = append(t.weaponListWeights, 1)
	}
	tempWingList, err := coreutils.SplitAsIntArray(t.WingId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.WingId)
		return template.NewTemplateFieldError("WingId", err)
	}
	t.wingList = tempWingList
	for _, _ = range t.wingList {
		t.wingListWeights = append(t.wingListWeights, 1)
	}
	tempMountList, err := coreutils.SplitAsIntArray(t.MountId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.MountId)
		return template.NewTemplateFieldError("MountId", err)
	}
	t.mountList = tempMountList
	for _, _ = range t.mountList {
		t.mountListWeights = append(t.mountListWeights, 1)
	}
	tempFabaoList, err := coreutils.SplitAsIntArray(t.FabaoId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FabaoId)
		return template.NewTemplateFieldError("FabaoId", err)
	}
	t.fabaoList = tempFabaoList
	for _, _ = range t.fabaoList {
		t.fabaoListWeights = append(t.fabaoListWeights, 1)
	}
	tempShenfaList, err := coreutils.SplitAsIntArray(t.ShenfaId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.ShenfaId)
		return template.NewTemplateFieldError("ShenfaId", err)
	}
	t.shenfaList = tempShenfaList
	for _, _ = range t.shenfaList {
		t.shenfaListWeights = append(t.shenfaListWeights, 1)
	}
	tempXianTiList, err := coreutils.SplitAsIntArray(t.XiantiId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.XiantiId)
		return template.NewTemplateFieldError("XiantiId", err)
	}
	t.xiantiList = tempXianTiList
	for _, _ = range t.xiantiList {
		t.xiantiListWeights = append(t.xiantiListWeights, 1)
	}
	tempFieldList, err := coreutils.SplitAsIntArray(t.FieldId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.FieldId)
		return template.NewTemplateFieldError("FieldId", err)
	}
	t.fieldList = tempFieldList
	for _, _ = range t.fieldList {
		t.fieldListWeights = append(t.fieldListWeights, 1)
	}
	tempTitleList, err := coreutils.SplitAsIntArray(t.TitleId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.TitleId)
		return template.NewTemplateFieldError("TitleId", err)
	}
	t.titleList = tempTitleList
	for _, _ = range t.titleList {
		t.titleListWeights = append(t.titleListWeights, 1)
	}
	tempJuexueList, err := coreutils.SplitAsIntArray(t.SkillJuexueId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.SkillJuexueId)
		return template.NewTemplateFieldError("SkillJuexueId", err)
	}
	t.jueXueList = tempJuexueList
	for _, _ = range t.jueXueList {
		t.jueXueListWeights = append(t.jueXueListWeights, 1)
	}
	tempSoulList, err := coreutils.SplitAsIntArray(t.SkillSoulId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.SkillSoulId)
		return template.NewTemplateFieldError("SkillSoulId", err)
	}
	t.soulList = tempSoulList
	for _, _ = range t.soulList {
		t.soulListWeights = append(t.soulListWeights, 1)
	}

	tempLingTongList, err := coreutils.SplitAsIntArray(t.LingTongId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongId)
		return template.NewTemplateFieldError("LingTongId", err)
	}
	t.lingTongList = tempLingTongList
	for _, _ = range t.lingTongList {
		t.lingTongListWeights = append(t.lingTongListWeights, 1)
	}

	tempLingTongFashionList, err := coreutils.SplitAsIntArray(t.LingTongFashionId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongFashionId)
		return template.NewTemplateFieldError("LingTongFashionId", err)
	}
	t.lingTongFashionList = tempLingTongFashionList
	for _, _ = range t.lingTongFashionList {
		t.lingTongFashionListWeights = append(t.lingTongFashionListWeights, 1)
	}

	tempLingTongWeaponList, err := coreutils.SplitAsIntArray(t.LingTongWeapon)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongWeapon)
		return template.NewTemplateFieldError("LingTongWeapon", err)
	}
	t.lingTongWeaponList = tempLingTongWeaponList
	for _, _ = range t.lingTongWeaponList {
		t.lingTongWeaponListWeights = append(t.lingTongWeaponListWeights, 1)
	}

	tempLingTongWingList, err := coreutils.SplitAsIntArray(t.LingTongWingId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongWingId)
		return template.NewTemplateFieldError("LingTongWingId", err)
	}
	t.lingTongWingList = tempLingTongWingList
	for _, _ = range t.lingTongWingList {
		t.lingTongWingListWeights = append(t.lingTongWingListWeights, 1)
	}

	tempLingTongMountList, err := coreutils.SplitAsIntArray(t.LingTongMountId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongMountId)
		return template.NewTemplateFieldError("LingTongMountId", err)
	}
	t.lingTongMountList = tempLingTongMountList
	for _, _ = range t.lingTongList {
		t.lingTongMountListWeights = append(t.lingTongMountListWeights, 1)
	}

	tempLingTongFabaoList, err := coreutils.SplitAsIntArray(t.LingTongFabaoId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongFabaoId)
		return template.NewTemplateFieldError("LingTongFabaoId", err)
	}
	t.lingTongFabaoList = tempLingTongFabaoList
	for _, _ = range t.lingTongFabaoList {
		t.lingTongFabaoListWeights = append(t.lingTongFabaoListWeights, 1)
	}

	tempLingTongShenfaList, err := coreutils.SplitAsIntArray(t.LingTongShenfaId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongShenfaId)
		return template.NewTemplateFieldError("LingTongShenfaId", err)
	}
	t.lingTongShenfaList = tempLingTongShenfaList
	for _, _ = range t.lingTongShenfaList {
		t.lingTongShenfaListWeights = append(t.lingTongShenfaListWeights, 1)
	}

	tempLingTongXiantiList, err := coreutils.SplitAsIntArray(t.LingTongXianTiId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongXianTiId)
		return template.NewTemplateFieldError("LingTongXianTiId", err)
	}
	t.lingTongXiantiList = tempLingTongXiantiList
	for _, _ = range t.lingTongXiantiList {
		t.lingTongXiantiListWeights = append(t.lingTongXiantiListWeights, 1)
	}

	tempLingTongLingYuList, err := coreutils.SplitAsIntArray(t.LingTongFieldId)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", t.LingTongFieldId)
		return template.NewTemplateFieldError("LingTongFieldId", err)
	}
	t.lingTongLingyuList = tempLingTongLingYuList
	for _, _ = range t.lingTongLingyuList {
		t.lingTongLingyuListWeights = append(t.lingTongLingyuListWeights, 1)
	}

	return nil
}

func (t *RobotTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()
	err = validator.MinValidate(float64(t.LevMin), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevMin)
		return template.NewTemplateFieldError("LevMin", err)
	}

	err = validator.MinValidate(float64(t.LevMax), float64(t.LevMin), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", t.LevMax)
		return template.NewTemplateFieldError("LevMax", err)
	}

	for _, profession := range t.professionList {
		if !playertypes.RoleType(profession).Valid() {
			err = fmt.Errorf("[%s] invalid", t.NeedProfession)
			return template.NewTemplateFieldError("NeedProfession", err)
		}
	}
	if len(t.professionList) <= 0 {
		err = fmt.Errorf("[%s] 至少需要1个选项", t.NeedProfession)
		return template.NewTemplateFieldError("NeedProfession", err)
	}
	for _, gender := range t.genderList {
		if !playertypes.SexType(gender).Valid() {
			err = fmt.Errorf("[%s] invalid", t.NeedGender)
			return template.NewTemplateFieldError("NeedGender", err)
		}
	}
	if len(t.genderList) <= 0 {
		err = fmt.Errorf("[%s] 至少需要1个选项", t.NeedGender)
		return template.NewTemplateFieldError("NeedGender", err)
	}
	for _, fashionId := range t.fashionList {
		tempFashionTemplate := template.GetTemplateService().Get(int(fashionId), (*FashionTemplate)(nil))
		if tempFashionTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.FashionId)
			return template.NewTemplateFieldError("FashionId", err)
		}
	}
	if len(t.fashionList) <= 0 {
		err = fmt.Errorf("[%s] 至少需要1个选项", t.FashionId)
		return template.NewTemplateFieldError("FashionId", err)
	}
	for _, weaponId := range t.weaponList {
		tempWeaponTemplate := template.GetTemplateService().Get(int(weaponId), (*WeaponTemplate)(nil))
		if tempWeaponTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.WeaponId)
			return template.NewTemplateFieldError("WeaponId", err)
		}
	}
	if len(t.weaponList) <= 0 {
		err = fmt.Errorf("[%s] 至少需要1个选项", t.WeaponId)
		return template.NewTemplateFieldError("WeaponId", err)
	}
	for _, wingId := range t.wingList {
		tempWingTemplate := template.GetTemplateService().Get(int(wingId), (*WingTemplate)(nil))
		if tempWingTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.WingId)
			return template.NewTemplateFieldError("WingId", err)
		}
	}
	for _, mountId := range t.mountList {
		tempMountTemplate := template.GetTemplateService().Get(int(mountId), (*MountTemplate)(nil))
		if tempMountTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.MountId)
			return template.NewTemplateFieldError("MountId", err)
		}
	}
	for _, fabaoId := range t.fabaoList {
		tempFaBaoTemplate := template.GetTemplateService().Get(int(fabaoId), (*FaBaoTemplate)(nil))
		if tempFaBaoTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.FabaoId)
			return template.NewTemplateFieldError("FabaoId", err)
		}
	}
	for _, shenfaId := range t.shenfaList {
		tempShenFaTemplate := template.GetTemplateService().Get(int(shenfaId), (*ShenfaTemplate)(nil))
		if tempShenFaTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.ShenfaId)
			return template.NewTemplateFieldError("ShenfaId", err)
		}
	}
	for _, xiantiId := range t.xiantiList {
		tempXiantiTemplate := template.GetTemplateService().Get(int(xiantiId), (*XianTiTemplate)(nil))
		if tempXiantiTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.XiantiId)
			return template.NewTemplateFieldError("XiantiId", err)
		}
	}
	for _, fieldId := range t.fieldList {
		tempLingyuTemplate := template.GetTemplateService().Get(int(fieldId), (*LingyuTemplate)(nil))
		if tempLingyuTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.FieldId)
			return template.NewTemplateFieldError("FieldId", err)
		}
	}
	for _, titleId := range t.titleList {
		tempTitleTemplate := template.GetTemplateService().Get(int(titleId), (*TitleTemplate)(nil))
		if tempTitleTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.TitleId)
			return template.NewTemplateFieldError("TitleId", err)
		}
	}
	for _, skillId := range t.jueXueList {
		tempSkillTemplate := template.GetTemplateService().Get(int(skillId), (*SkillTemplate)(nil))
		if tempSkillTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.SkillJuexueId)
			return template.NewTemplateFieldError("SkillJuexueId", err)
		}
		skillTemplate := tempSkillTemplate.(*SkillTemplate)
		if skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeJueXue {
			err = fmt.Errorf("[%s] invalid", t.SkillJuexueId)
			return template.NewTemplateFieldError("SkillJuexueId", err)
		}
	}
	for _, soulId := range t.soulList {
		tempSkillTemplate := template.GetTemplateService().Get(int(soulId), (*SkillTemplate)(nil))
		if tempSkillTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.SkillSoulId)
			return template.NewTemplateFieldError("SkillSoulId", err)
		}
		skillTemplate := tempSkillTemplate.(*SkillTemplate)
		if skillTemplate.GetSkillFirstType() != skilltypes.SkillFirstTypeGuHun {
			err = fmt.Errorf("[%s] invalid", t.SkillSoulId)
			return template.NewTemplateFieldError("SkillSoulId", err)
		}
	}

	for _, lingTongId := range t.lingTongList {
		tempLingTongTemplate := template.GetTemplateService().Get(int(lingTongId), (*LingTongTemplate)(nil))
		if tempLingTongTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongId)
			return template.NewTemplateFieldError("LingTongId", err)
		}
	}
	for _, lingTongFashionId := range t.lingTongFashionList {
		tempLingTongFashionTemplate := template.GetTemplateService().Get(int(lingTongFashionId), (*LingTongTemplate)(nil))
		if tempLingTongFashionTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongFashionId)
			return template.NewTemplateFieldError("LingTongFashionId", err)
		}
	}
	for _, lingTongWeaponId := range t.lingTongWeaponList {
		tempLingTongWeaponTemplate := template.GetTemplateService().Get(int(lingTongWeaponId), (*LingTongWeaponTemplate)(nil))
		if tempLingTongWeaponTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongWeapon)
			return template.NewTemplateFieldError("LingTongWeapon", err)
		}
	}
	for _, lingTongWingId := range t.lingTongWingList {
		tempLingTongWingTemplate := template.GetTemplateService().Get(int(lingTongWingId), (*LingTongWingTemplate)(nil))
		if tempLingTongWingTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongWingId)
			return template.NewTemplateFieldError("LingTongWingId", err)
		}
	}
	for _, lingTongMountId := range t.lingTongMountList {
		tempLingTongMountTemplate := template.GetTemplateService().Get(int(lingTongMountId), (*LingTongMountTemplate)(nil))
		if tempLingTongMountTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongMountId)
			return template.NewTemplateFieldError("LingTongMountId", err)
		}
	}
	for _, lingTongFabaoId := range t.lingTongFabaoList {
		tempLingTongFabaoTemplate := template.GetTemplateService().Get(int(lingTongFabaoId), (*LingTongFaBaoTemplate)(nil))
		if tempLingTongFabaoTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongFabaoId)
			return template.NewTemplateFieldError("lingTongFabaoId", err)
		}
	}
	for _, lingTongShenfaId := range t.lingTongShenfaList {
		tempLingTongTemplate := template.GetTemplateService().Get(int(lingTongShenfaId), (*LingTongShenFaTemplate)(nil))
		if tempLingTongTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongShenfaId)
			return template.NewTemplateFieldError("LingTongShenfaId", err)
		}
	}
	for _, lingTongXiantiId := range t.lingTongXiantiList {
		tempLingTongTemplate := template.GetTemplateService().Get(int(lingTongXiantiId), (*LingTongXianTiTemplate)(nil))
		if tempLingTongTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongXianTiId)
			return template.NewTemplateFieldError("LingTongXianTiId", err)
		}
	}
	for _, lingTongLingyuId := range t.lingTongLingyuList {
		tempLingTongTemplate := template.GetTemplateService().Get(int(lingTongLingyuId), (*LingTongLingYuTemplate)(nil))
		if tempLingTongTemplate == nil {
			err = fmt.Errorf("[%s] invalid", t.LingTongFieldId)
			return template.NewTemplateFieldError("LingTongFieldId", err)
		}
	}

	return
}

func (t *RobotTemplate) FileName() string {
	return "tb_robot.json"
}

func init() {
	template.Register((*RobotTemplate)(nil))
}
