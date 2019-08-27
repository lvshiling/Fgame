package player

import (
	"fgame/fgame/game/constant/constant"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"

	"fgame/fgame/game/xuedun/dao"

	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	gametemplate "fgame/fgame/game/template"
	xueduncommon "fgame/fgame/game/xuedun/common"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	xueduntemplate "fgame/fgame/game/xuedun/template"
)

//玩家血盾管理器
type PlayerXueDunDataManager struct {
	p player.Player
	//玩家血盾对象
	playerXueDunObject *PlayerXueDunObject
}

func (pmdm *PlayerXueDunDataManager) Player() player.Player {
	return pmdm.p
}

//加载
func (pmdm *PlayerXueDunDataManager) Load() (err error) {
	//加载玩家血盾信息
	xueDunEntity, err := dao.GetXueDunDao().GetXueDunEntity(pmdm.p.GetId())
	if err != nil {
		return
	}
	if xueDunEntity == nil {
		pmdm.initPlayerXueDunObject()
	} else {
		pmdm.playerXueDunObject = NewPlayerXueDunObject(pmdm.p)
		pmdm.playerXueDunObject.FromEntity(xueDunEntity)
	}

	return nil
}

//第一次初始化
func (pmdm *PlayerXueDunDataManager) initPlayerXueDunObject() {
	pmo := NewPlayerXueDunObject(pmdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pmo.id = id
	//生成id
	pmo.playerId = pmdm.p.GetId()
	pmo.number = 0
	pmo.blood = 0
	pmo.power = 0
	pmo.star = 0
	pmo.starNum = 0
	pmo.starPro = 0
	pmo.culLevel = 0
	pmo.culNum = 0
	pmo.culPro = 0
	pmo.power = 0
	pmo.isActive = 0
	pmo.createTime = now
	pmdm.playerXueDunObject = pmo
	pmo.SetModified()
}

//加载后
func (pmdm *PlayerXueDunDataManager) AfterLoad() (err error) {
	return nil
}

//血盾信息对象
func (pmdm *PlayerXueDunDataManager) GetXueDunInfo() *PlayerXueDunObject {
	return pmdm.playerXueDunObject
}

//获取当前阶数
func (pmdm *PlayerXueDunDataManager) GetXueDunNumber() int32 {
	return int32(pmdm.playerXueDunObject.number)
}

//心跳
func (pmdm *PlayerXueDunDataManager) Heartbeat() {

}

//是否满阶
func (pmdm *PlayerXueDunDataManager) IsFull() (bloodShieldTemplate *gametemplate.BloodShieldTemplate, flag bool) {
	isActive := pmdm.playerXueDunObject.GetIsActive()
	if !isActive {
		bloodShieldTemplate = xueduntemplate.GetXueDunTemplateService().GetXueDunActiveTemplate()
		return
	}
	number := pmdm.playerXueDunObject.GetNumber()
	curStar := pmdm.playerXueDunObject.GetStar()
	bloodShieldTemplate = xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(number, curStar)
	if bloodShieldTemplate.NextId == 0 {
		nextNumber := number + 1
		bloodShieldTemplate = xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(nextNumber, 1)
	} else {
		bloodShieldTemplate = xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(number, curStar+1)
	}
	if bloodShieldTemplate == nil {
		flag = true
	}
	return
}

//获取食丹等级上限
func (pmdm *PlayerXueDunDataManager) GetCulLevelLimit() (num int32) {
	number := pmdm.playerXueDunObject.GetNumber()
	curStar := pmdm.playerXueDunObject.GetStar()
	bloodShieldTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(number, curStar)
	if bloodShieldTemplate == nil {
		return
	}
	return bloodShieldTemplate.MedicinalLimit
}

func (pmdm *PlayerXueDunDataManager) EatCulDan(level int32) {
	if pmdm.playerXueDunObject.GetCulLevel() == level || level <= 0 {
		return
	}
	xueDunPeiYangTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunPeiYangTemplate(level)
	if xueDunPeiYangTemplate == nil {
		return
	}
	pmdm.playerXueDunObject.culLevel = level
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
	return
}

//血盾战斗力
func (pmdm *PlayerXueDunDataManager) XueDunPower(power int64) {
	if power <= 0 {
		return
	}
	if pmdm.playerXueDunObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXueDunObject.power = power
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
	return
}

//升阶
func (pmdm *PlayerXueDunDataManager) Upgrade(bloodShieldTemplate *gametemplate.BloodShieldTemplate, pro int32, sucess bool) (flag bool) {
	if bloodShieldTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	oldSkillId := int32(0)
	beforeNum := pmdm.playerXueDunObject.GetNumber()
	beforeStar := pmdm.playerXueDunObject.GetStar()

	isActive := pmdm.playerXueDunObject.GetIsActive()
	if sucess {
		if !isActive {
			pmdm.playerXueDunObject.isActive = 1
		} else {
			beforeTemplate := xueduntemplate.GetXueDunTemplateService().GetXueDunNumber(beforeNum, beforeStar)
			oldSkillId = beforeTemplate.SpellId
		}
		pmdm.playerXueDunObject.number = bloodShieldTemplate.Type
		pmdm.playerXueDunObject.star = bloodShieldTemplate.Star
		pmdm.playerXueDunObject.starNum = 0
		pmdm.playerXueDunObject.starPro = 0

		gameevent.Emit(xueduneventtypes.EventTypeXueDunUpgrade, pmdm.p, oldSkillId)

	} else {
		pmdm.playerXueDunObject.starNum++
		pmdm.playerXueDunObject.starPro += pro
	}
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
	if pmdm.playerXueDunObject.number != beforeNum {
		gameevent.Emit(xueduneventtypes.EventTypeXueDunNumberChanged, pmdm.p, nil)
	}
	flag = true
	return
}

func (pmdm *PlayerXueDunDataManager) XueDunSubBloodChanged(subBlood int64) (flag bool) {
	if subBlood < 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if pmdm.playerXueDunObject.blood < subBlood {
		return
	}
	pmdm.playerXueDunObject.blood -= subBlood
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
	flag = true
	gameevent.Emit(xueduneventtypes.EventTypeXueDunBloodChanged, pmdm.p, pmdm.playerXueDunObject.blood)
	return
}

func (pmdm *PlayerXueDunDataManager) XueDunBloodChanged(addBlood int64) {
	if addBlood <= 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	bloodLimit := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeXueDunBloodLimit))
	leftBlood := bloodLimit - addBlood
	if leftBlood < addBlood {
		addBlood = leftBlood
	}
	pmdm.playerXueDunObject.blood += addBlood
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
	gameevent.Emit(xueduneventtypes.EventTypeXueDunBloodChanged, pmdm.p, pmdm.playerXueDunObject.blood)
}

func (pmdm *PlayerXueDunDataManager) ToXueDunInfo() *xueduncommon.XueDunInfo {
	info := &xueduncommon.XueDunInfo{
		Number:   pmdm.playerXueDunObject.number,
		Star:     pmdm.playerXueDunObject.star,
		StarPro:  pmdm.playerXueDunObject.starPro,
		CulLevel: pmdm.playerXueDunObject.culLevel,
		CulPro:   pmdm.playerXueDunObject.culPro,
	}
	return info
}

//仅gm使用 设置血炼值
func (pmdm *PlayerXueDunDataManager) GmSetXueDunBlood(blood int64) {
	pmdm.playerXueDunObject.blood = blood
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
}

//仅gm使用 血盾进阶
func (pmdm *PlayerXueDunDataManager) GmSetXueDunNumber(number int32) {
	pmdm.playerXueDunObject.number = number
	pmdm.playerXueDunObject.star = 1
	pmdm.playerXueDunObject.starNum = 0
	pmdm.playerXueDunObject.starPro = 0
	pmdm.playerXueDunObject.culLevel = 0
	pmdm.playerXueDunObject.culPro = 0
	pmdm.playerXueDunObject.isActive = 1
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
	return
}

//仅gm使用 血盾重置
func (pmdm *PlayerXueDunDataManager) GmSetXueDunReset() {
	pmdm.playerXueDunObject.number = 0
	pmdm.playerXueDunObject.star = 0
	pmdm.playerXueDunObject.starNum = 0
	pmdm.playerXueDunObject.starPro = 0
	pmdm.playerXueDunObject.culLevel = 0
	pmdm.playerXueDunObject.culPro = 0
	pmdm.playerXueDunObject.isActive = 0
	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
}

//仅gm使用 血盾食草料等级
func (pmdm *PlayerXueDunDataManager) GmSetXueDunPeiYangLevel(level int32) {

	pmdm.playerXueDunObject.culLevel = level
	pmdm.playerXueDunObject.culNum = 0
	pmdm.playerXueDunObject.culPro = 0

	now := global.GetGame().GetTimeService().Now()
	pmdm.playerXueDunObject.updateTime = now
	pmdm.playerXueDunObject.SetModified()
}

func CreatePlayerXueDunDataManager(p player.Player) player.PlayerDataManager {
	pmdm := &PlayerXueDunDataManager{}
	pmdm.p = p
	return pmdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerXueDunDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerXueDunDataManager))
}
