package template

import (
	"fgame/fgame/core/template"
	fashiontypes "fgame/fgame/game/fashion/types"
	wardrobetypes "fgame/fgame/game/wardrobe/types"
	weapontypes "fgame/fgame/game/weapon/types"
	"fmt"
)

type sysIdInfo struct {
	seqId     int32
	permanent bool
}

func newSysIdInfo(seqId int32, permanent bool) *sysIdInfo {
	d := &sysIdInfo{
		seqId:     seqId,
		permanent: permanent,
	}
	return d
}

func (t *sysIdInfo) GetSeqId() int32 {
	return t.seqId
}

func (t *sysIdInfo) GetIsPermanent() bool {
	return t.permanent
}

//衣橱配置
type YiChuSuitTemplate struct {
	*YiChuSuitTemplateVO
	wardrobeType       int32 //衣橱类型
	peiYangTemplate    *YiChuPeiYangTemplate
	sysIdMap           map[wardrobetypes.WardrobeSysType]*sysIdInfo
	peiYangTemplateMap map[int32]*YiChuPeiYangTemplate //衣橱培养map
	permanentNum       int32
}

func (t *YiChuSuitTemplate) TemplateId() int {
	return t.Id
}

func (t *YiChuSuitTemplate) GetType() int32 {
	return t.wardrobeType
}

func (t *YiChuSuitTemplate) GetSysIdMap() map[wardrobetypes.WardrobeSysType]*sysIdInfo {
	return t.sysIdMap
}

func (t *YiChuSuitTemplate) GetPermanentNum() int32 {
	return t.permanentNum
}

func (t *YiChuSuitTemplate) IfExist(sysType wardrobetypes.WardrobeSysType, seqId int32) (flag bool) {
	sysInfo, ok := t.sysIdMap[sysType]
	if !ok {
		return
	}
	if sysInfo.GetSeqId() != seqId {
		return
	}
	flag = true
	return
}

func (t *YiChuSuitTemplate) GetPeiYangByLevel(level int32) *YiChuPeiYangTemplate {
	if v, ok := t.peiYangTemplateMap[level]; ok {
		return v
	}
	return nil
}

func (t *YiChuSuitTemplate) PatchAfterCheck() {
	if t.peiYangTemplate != nil {
		t.peiYangTemplateMap = make(map[int32]*YiChuPeiYangTemplate)
		//赋值PeiYangTemplateMap
		for tempTempalte := t.peiYangTemplate; tempTempalte != nil; tempTempalte = tempTempalte.nextYiChuPeiYangTemplate {
			level := tempTempalte.Level
			t.peiYangTemplateMap[level] = tempTempalte
		}
	}
}

func (t *YiChuSuitTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	//验证 type
	t.wardrobeType = int32(t.Type)
	// if !t.wardrobeType.Valid() {
	// 	err = fmt.Errorf("[%d] invalid", t.Type)
	// 	err = template.NewTemplateFieldError("Type", err)
	// 	return
	// }

	//验证 peiyang_begin_id
	if t.PeiYangBeginId != 0 {
		to := template.GetTemplateService().Get(int(t.PeiYangBeginId), (*YiChuPeiYangTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.PeiYangBeginId)
			return template.NewTemplateFieldError("PeiYangBeginId", err)
		}

		peiYangTemplate, ok := to.(*YiChuPeiYangTemplate)
		if !ok {
			return fmt.Errorf("peiYangBeginId [%d] invalid", t.PeiYangBeginId)
		}
		t.peiYangTemplate = peiYangTemplate
	}

	t.sysIdMap = make(map[wardrobetypes.WardrobeSysType]*sysIdInfo)

	//验证坐骑id
	if t.MountId != 0 {
		to := template.GetTemplateService().Get(int(t.MountId), (*MountTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.MountId)
			err = template.NewTemplateFieldError("MountId", err)
			return
		}
		sysIdInfo := newSysIdInfo(t.MountId, true)
		t.sysIdMap[wardrobetypes.WardrobeSysTypeMount] = sysIdInfo
	}

	//验证战翼id
	if t.WingId != 0 {
		to := template.GetTemplateService().Get(int(t.WingId), (*WingTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.WingId)
			err = template.NewTemplateFieldError("WingId", err)
			return
		}
		sysIdInfo := newSysIdInfo(t.WingId, true)
		t.sysIdMap[wardrobetypes.WardrobeSysTypeWing] = sysIdInfo
	}

	//验证兵魂id
	if t.WeaponId != 0 {
		to := template.GetTemplateService().Get(int(t.WeaponId), (*WeaponTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.WeaponId)
			err = template.NewTemplateFieldError("WeaponId", err)
			return
		}
		weaponTemplate := to.(*WeaponTemplate)
		sysIdInfo := newSysIdInfo(t.WeaponId, false)
		if weaponTemplate.GetWeaponTag() != weapontypes.WeaponTagTypeTemp {
			sysIdInfo.permanent = true
		}
		t.sysIdMap[wardrobetypes.WardrobeSysTypeWeapon] = sysIdInfo
	}

	//验证法宝id
	if t.FaBaoId != 0 {
		to := template.GetTemplateService().Get(int(t.FaBaoId), (*FaBaoTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.FaBaoId)
			err = template.NewTemplateFieldError("FaBaoId", err)
			return
		}
		sysIdInfo := newSysIdInfo(t.FaBaoId, true)
		t.sysIdMap[wardrobetypes.WardrobeSysTypeFaBao] = sysIdInfo
	}

	//验证身法id
	if t.ShenFaId != 0 {
		to := template.GetTemplateService().Get(int(t.ShenFaId), (*ShenfaTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.ShenFaId)
			err = template.NewTemplateFieldError("ShenFaId", err)
			return
		}
		sysIdInfo := newSysIdInfo(t.ShenFaId, true)
		t.sysIdMap[wardrobetypes.WardrobeSysTypeShenFa] = sysIdInfo
	}

	//验证仙体id
	if t.XianTiId != 0 {
		to := template.GetTemplateService().Get(int(t.XianTiId), (*XianTiTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.XianTiId)
			err = template.NewTemplateFieldError("XianTiId", err)
			return
		}
		sysIdInfo := newSysIdInfo(t.XianTiId, true)
		t.sysIdMap[wardrobetypes.WardrobeSysTypeXianTi] = sysIdInfo
	}

	//验证领域id
	if t.FieldId != 0 {
		to := template.GetTemplateService().Get(int(t.FieldId), (*LingyuTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.FieldId)
			err = template.NewTemplateFieldError("FieldId", err)
			return
		}
		sysIdInfo := newSysIdInfo(t.FieldId, true)
		t.sysIdMap[wardrobetypes.WardrobeSysTypeField] = sysIdInfo
	}

	//验证时装id
	if t.FashionId != 0 {
		to := template.GetTemplateService().Get(int(t.FashionId), (*FashionTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.FashionId)
			err = template.NewTemplateFieldError("FashionId", err)
			return
		}
		fashionTemplate := to.(*FashionTemplate)
		sysIdInfo := newSysIdInfo(t.FashionId, false)
		if fashionTemplate.GetFashionType() != fashiontypes.FashionTypeEffective {
			sysIdInfo.permanent = true
		}
		t.sysIdMap[wardrobetypes.WardrobeSysTypeFashion] = sysIdInfo
	}

	//验证称号id
	if t.TitleId != 0 {
		to := template.GetTemplateService().Get(int(t.TitleId), (*TitleTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", t.TitleId)
			err = template.NewTemplateFieldError("TitleId", err)
			return
		}
		titleTemplate := to.(*TitleTemplate)
		sysIdInfo := newSysIdInfo(t.TitleId, false)
		if titleTemplate.Time == 0 {
			sysIdInfo.permanent = true
		}
		t.sysIdMap[wardrobetypes.WardrobeSysTypeTitle] = sysIdInfo
	}
	for _, sysIdInfo := range t.sysIdMap {
		if !sysIdInfo.GetIsPermanent() {
			continue
		}
		t.permanentNum++
	}

	return nil
}

func (t *YiChuSuitTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(t.FileName(), t.TemplateId(), err)
			return
		}
	}()

	return nil
}

func (t *YiChuSuitTemplate) GetEatPeiYangTemplate(curLevel int32, num int32) (peiYangTemplate *YiChuPeiYangTemplate, flag bool) {
	if num <= 0 {
		return
	}
	leftNum := num
	for level := curLevel + 1; leftNum > 0; level++ {
		peiYangTemplate, flag = t.peiYangTemplateMap[level]
		if !flag {
			return nil, false
		}
		leftNum -= peiYangTemplate.ItemCount
	}
	if leftNum != 0 {
		return nil, false
	}
	flag = true
	return
}

func (t *YiChuSuitTemplate) FileName() string {
	return "tb_yichu_suit.json"
}

func init() {
	template.Register((*YiChuSuitTemplate)(nil))
}
