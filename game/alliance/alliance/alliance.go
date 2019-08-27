package alliance

import (
	"fgame/fgame/core/storage"
	allianceentity "fgame/fgame/game/alliance/entity"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

const (
	douShenNumLimit = 9
)

type AllianceObject struct {
	id                       int64
	serverId                 int32
	originServerId           int32
	name                     string
	notice                   string
	level                    int32
	jianShe                  int64
	huFu                     int64
	totalForce               int64
	mengzhuId                int64
	createId                 int64
	transportTimes           int32
	lastTransportRefreshTime int64
	isAutoAgree              int32
	isAutoRemoveDepot        int32
	maxRemoveZhuanSheng      int32
	maxRemoveQuality         itemtypes.ItemQualityType
	lastMergeTime            int64
	// campType                 chuangshitypes.ChuangShiCampType
	// lastCampChangedTime      int64
	updateTime int64
	createTime int64
	deleteTime int64
}

func createAllianceObject() *AllianceObject {
	o := &AllianceObject{}
	return o
}

func convertAllianceObjectToEntity(o *AllianceObject) (*allianceentity.AllianceEntity, error) {
	e := &allianceentity.AllianceEntity{
		Id:                       o.id,
		ServerId:                 o.serverId,
		OriginServerId:           o.originServerId,
		Name:                     o.name,
		Notice:                   o.notice,
		Level:                    o.level,
		HuFu:                     o.huFu,
		JianShe:                  o.jianShe,
		TotalForce:               o.totalForce,
		MengzhuId:                o.mengzhuId,
		CreateId:                 o.createId,
		TransportTimes:           o.transportTimes,
		LastTransportRefreshTime: o.lastTransportRefreshTime,
		IsAutoAgree:              o.isAutoAgree,
		IsAutoRemoveDepot:        o.isAutoRemoveDepot,
		MaxRemoveZhuanSheng:      o.maxRemoveZhuanSheng,
		MaxRemoveQuality:         int32(o.maxRemoveQuality),
		LastMergeTime:            o.lastMergeTime,
		// CampType:                 int32(o.campType),
		UpdateTime: o.updateTime,
		CreateTime: o.createTime,
		DeleteTime: o.deleteTime,
	}
	return e, nil
}

func (o *AllianceObject) GetId() int64 {
	return o.id
}

func (o *AllianceObject) GetDBId() int64 {
	return o.id
}

func (o *AllianceObject) GetName() string {
	return o.name
}

func (o *AllianceObject) GetNotice() string {
	return o.notice
}

func (o *AllianceObject) GetLevel() int32 {
	return o.level
}

func (o *AllianceObject) GetJianShe() int64 {
	return o.jianShe
}

func (o *AllianceObject) GetHuFu() int64 {
	return o.huFu
}

func (o *AllianceObject) GetTotalForce() int64 {
	return o.totalForce
}

func (o *AllianceObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *AllianceObject) GetMengzhuId() int64 {
	return o.mengzhuId
}

func (o *AllianceObject) GetIsAutoAgree() int32 {
	return o.isAutoAgree
}

// func (o *AllianceObject) GetCampType() chuangshitypes.ChuangShiCampType {
// 	return o.campType
// }

func (o *AllianceObject) GetIsAutoRemoveDepot() int32 {
	return o.isAutoRemoveDepot
}

func (o *AllianceObject) GetMaxRemoveZhuanSheng() int32 {
	return o.maxRemoveZhuanSheng
}

func (o *AllianceObject) GetMaxRemoveQuality() itemtypes.ItemQualityType {
	return o.maxRemoveQuality
}

func (o *AllianceObject) GetCreateId() int64 {
	return o.createId
}

func (o *AllianceObject) GetOriginServerId() int32 {
	return o.originServerId
}

func (o *AllianceObject) GetMergeAllianceTime() int64 {
	return o.lastMergeTime
}

func (o *AllianceObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertAllianceObjectToEntity(o)
	return e, err
}

func (o *AllianceObject) FromEntity(e storage.Entity) error {
	ae, _ := e.(*allianceentity.AllianceEntity)
	o.id = ae.Id
	o.serverId = ae.ServerId
	o.originServerId = ae.OriginServerId
	o.name = ae.Name
	o.notice = ae.Notice
	o.level = ae.Level
	o.huFu = ae.HuFu
	o.jianShe = ae.JianShe
	o.mengzhuId = ae.MengzhuId
	o.createId = ae.CreateId
	o.totalForce = ae.TotalForce
	o.transportTimes = ae.TransportTimes
	o.lastTransportRefreshTime = ae.LastTransportRefreshTime
	o.isAutoAgree = ae.IsAutoAgree
	o.isAutoRemoveDepot = ae.IsAutoRemoveDepot
	o.maxRemoveZhuanSheng = ae.MaxRemoveZhuanSheng
	o.lastMergeTime = ae.LastMergeTime
	o.maxRemoveQuality = itemtypes.ItemQualityType(ae.MaxRemoveQuality)
	// o.campType = chuangshitypes.ChuangShiCampType(ae.CampType)
	o.updateTime = ae.UpdateTime
	o.createTime = ae.CreateTime
	o.deleteTime = ae.DeleteTime
	return nil
}

func (o *AllianceObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Alliance"))
	}
	global.GetGame().GetGlobalUpdater().AddChangedObject(e)
	return
}

//排序
type AllianceMemberForceRankDataList []*AllianceMemberObject

func (c AllianceMemberForceRankDataList) Len() int {
	return len(c)
}

func (crl AllianceMemberForceRankDataList) Less(i, j int) bool {
	if crl[i].GetForce() == 0 {
		return true
	}
	if crl[j].GetForce() == 0 {
		return false
	}
	return crl[i].GetForce() < crl[j].GetForce()
}

func (c AllianceMemberForceRankDataList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//排序
type AllianceDouShenMemberRankDataList []*AllianceMemberObject

func (c AllianceDouShenMemberRankDataList) Len() int {
	return len(c)
}

func (crl AllianceDouShenMemberRankDataList) Less(i, j int) bool {
	if crl[i].lingyuId != crl[j].lingyuId {
		return crl[i].lingyuId < crl[j].lingyuId
	}
	if crl[i].GetForce() == 0 {
		return true
	}
	if crl[j].GetForce() == 0 {
		return false
	}
	return crl[i].GetForce() < crl[j].GetForce()
}

func (c AllianceDouShenMemberRankDataList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

//仙盟
type Alliance struct {
	allianceObject     *AllianceObject
	memberList         []*AllianceMemberObject     //成员列表
	applyList          []*AllianceJoinApplyObject  //申请列表
	managerList        []*AllianceMemberObject     //管理层
	logList            []*AllianceLogObject        //日志列表
	douShenList        []*AllianceMemberObject     //斗神列表
	invitationList     []*AllianceInvitationObject //邀请列表
	depotBag           *AllianceDepotBag           //仙盟仓库
	allianceBossObject *AllianceBossObject         //仙盟boss
}

func (al *Alliance) updateDouShenForceList(timestamp int64) (err error) {
	tempMemberList := make([]*AllianceMemberObject, len(al.GetMemberList()))
	copy(tempMemberList, al.GetMemberList())

	sort.Sort(sort.Reverse(AllianceDouShenMemberRankDataList(tempMemberList)))
	dousShenNum := len(al.memberList)
	if dousShenNum > douShenNumLimit {
		dousShenNum = douShenNumLimit
	}

	oldDouShenList := al.douShenList
	// 超过等级限制
	levelLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceDouShenLevelLimit)
	newList := make([]*AllianceMemberObject, 0, dousShenNum)
	for _, mem := range tempMemberList {
		if mem.level < levelLimit {
			continue
		}
		if mem.lingyuId <= 0 {
			continue
		}

		newList = append(newList, mem)
		if len(newList) < dousShenNum {
			continue
		}

		break
	}
	al.douShenList = newList

	// 新增的斗神成员列表
	var changedList []*AllianceMemberObject
	for _, newMem := range newList {
		isNew := true
		for _, oldMem := range oldDouShenList {
			if newMem.memberId == oldMem.memberId {
				isNew = false
			}
		}
		if isNew {
			changedList = append(changedList, newMem)
		}
	}

	if len(changedList) > 0 {
		//发送领域事件
		gameevent.Emit(allianceeventtypes.EventTypeAllianceDouShenLingYuChanged, al, changedList)
	}
	return
}

func (al *Alliance) GetAllianceId() int64 {
	return al.allianceObject.GetId()
}

func (al *Alliance) GetAllianceName() string {
	if merge.GetMergeService().GetMergeTime() > 0 {
		return fmt.Sprintf("s%d.%s", al.allianceObject.GetOriginServerId(), al.allianceObject.GetName())
	}
	return al.allianceObject.GetName()
}

func (al *Alliance) GetAllianceLevel() int32 {
	return al.allianceObject.GetLevel()
}

func (al *Alliance) GetIsAutoAgree() int32 {
	return al.allianceObject.GetIsAutoAgree()
}

func (al *Alliance) GetFuMengZhuId() int64 {
	for _, mem := range al.managerList {
		if mem.GetPosition() != alliancetypes.AlliancePositionFuMengZhu {
			continue
		}

		return mem.memberId
	}
	return 0
}

func (al *Alliance) GetAllianceMengZhuId() int64 {
	return al.allianceObject.GetMengzhuId()
}

func (al *Alliance) GetMengzhuName() string {
	_, mem := al.getMember(al.allianceObject.mengzhuId)
	if mem != nil {
		return mem.name
	}
	return ""
}

func (al *Alliance) GetMengzhuVip() int32 {
	_, mem := al.getMember(al.allianceObject.mengzhuId)
	if mem != nil {
		return mem.vip
	}
	return 0
}

func (al *Alliance) GetMengzhuWingId() int32 {
	info, _ := player.GetPlayerService().GetPlayerInfo(al.GetAllianceMengZhuId())
	if info != nil {
		return info.WingInfo.GetWingId()
	}
	return 0
}
func (al *Alliance) GetMengzhuWeaponId() int32 {
	info, _ := player.GetPlayerService().GetPlayerInfo(al.GetAllianceMengZhuId())
	if info != nil {
		return info.AllWeaponInfo.Wear
	}
	return 0
}
func (al *Alliance) GetMengzhuFashionId() int32 {
	info, _ := player.GetPlayerService().GetPlayerInfo(al.GetAllianceMengZhuId())
	if info != nil {
		return info.FashionId
	}
	return 0
}

func (al *Alliance) GetMengzhuSex() playertypes.SexType {
	_, mem := al.getMember(al.allianceObject.mengzhuId)
	if mem != nil {
		return mem.sex
	}
	return -1
}

func (al *Alliance) GetAllianceObject() *AllianceObject {
	return al.allianceObject
}

func (al *Alliance) GetMemberList() []*AllianceMemberObject {
	return al.memberList
}

func (al *Alliance) GetApplyList() []*AllianceJoinApplyObject {
	return al.applyList
}

func (al *Alliance) GetLogList() []*AllianceLogObject {
	return al.logList
}

func (al *Alliance) GetDouShenList() []*AllianceMemberObject {
	return al.douShenList
}

func (al *Alliance) NumOfMembers() int32 {
	return int32(len(al.memberList))
}

func (al *Alliance) GetAllManagers() []*AllianceMemberObject {
	return al.managerList
}

func (al *Alliance) GetManagerNum(pos alliancetypes.AlliancePosition) int32 {
	count := int32(0)
	for _, mem := range al.managerList {
		if mem.position != pos {
			continue
		}

		count += 1
	}

	return count
}

func (al *Alliance) GetNumOfManagers(position alliancetypes.AlliancePosition) int32 {
	num := int32(0)
	for _, memObj := range al.GetAllManagers() {
		if memObj.position == position {
			num += 1
		}
	}

	return num
}

func (al *Alliance) GetDepotItemList() []*AllianceDepotItemObject {
	return al.depotBag.GetItemList()
}

func (al *Alliance) GetDepotItemByIndex(index int32) *AllianceDepotItemObject {
	return al.depotBag.getByIndex(index)
}

func (al *Alliance) AddJianShe(jianShe int64) {
	if jianShe <= 0 {
		panic("增加建设度不能小于0")
	}

	now := global.GetGame().GetTimeService().Now()
	al.allianceObject.jianShe += jianShe
	al.checkExpAndLevel()
	al.allianceObject.updateTime = now
	al.allianceObject.SetModified()
}

func (al *Alliance) CostJianShe(jianShe int64) {
	if jianShe <= 0 {
		panic("减少建设度不能小于0")
	}

	// 判断创建仙盟版本
	versionType := checkAllianceVersion()

	newJianShe := al.GetAllianceObject().GetJianShe()
	newJianShe -= jianShe
	oldLevel := al.GetAllianceLevel()
	newLevel := oldLevel
	for newJianShe < 0 {
		newLevel -= 1
		maxJianShe, flag := getMaxExp(versionType, newLevel)
		if !flag {
			break
		}
		newJianShe += int64(maxJianShe)
	}

	al.allianceObject.jianShe = newJianShe
	al.allianceObject.level = newLevel

	al.allianceObject.SetModified()

	return

}

func (al *Alliance) checkExpAndLevel() {
	// 判断创建仙盟版本
	versionType := checkAllianceVersion()

	oldLevel := al.allianceObject.level
	newExp := al.allianceObject.jianShe
	newLevel := oldLevel
	for {
		exp, flag := getMaxExp(versionType, newLevel)
		if !flag {
			break
		}
		if newExp < int64(exp) {
			break
		}

		newLevel += 1
		newExp -= int64(exp)
	}
	if newLevel > oldLevel {
		al.allianceObject.level = newLevel
		al.allianceObject.jianShe = newExp

		gameevent.Emit(allianceeventtypes.EventTypeAllianceLevelChanged, al, nil)
	}
	return
}

func getMaxExp(version alliancetypes.AllianceVersionType, level int32) (exp int32, flag bool) {
	tempTemplateObject := alliancetemplate.GetAllianceTemplateService().GetAllianceTemplate(version, level)
	if tempTemplateObject == nil {
		return 0, false
	}

	if tempTemplateObject.GetNextLevelAllianceTemplate() == nil {
		return 0, false
	}
	exp = int32(tempTemplateObject.GetNextLevelAllianceTemplate().UnionBuild)
	flag = true
	return
}

func (al *Alliance) increaseHuFu() {
	now := global.GetGame().GetTimeService().Now()
	al.allianceObject.huFu += 1
	al.allianceObject.updateTime = now
	al.allianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceHuFuChanged, al, al.allianceObject.huFu)
}

func (al *Alliance) IsFull() bool {
	// 判断创建仙盟版本
	versionType := checkAllianceVersion()
	tem := alliancetemplate.GetAllianceTemplateService().GetAllianceTemplate(versionType, al.GetAllianceObject().level)
	maxNum := tem.UnionMax
	return al.NumOfMembers() >= maxNum
}

func (al *Alliance) IfFull(addNum int32) bool {
	// 判断创建仙盟版本
	versionType := checkAllianceVersion()
	tem := alliancetemplate.GetAllianceTemplateService().GetAllianceTemplate(versionType, al.GetAllianceObject().level)
	maxNum := tem.UnionMax
	return al.NumOfMembers()+addNum > maxNum
}

func (al *Alliance) getMember(memberId int64) (index int32, memObj *AllianceMemberObject) {
	for index, memObj := range al.memberList {
		if memberId == memObj.GetMemberId() {
			return int32(index), memObj
		}
	}
	return -1, nil

}

func (al *Alliance) addMember(memberObj *AllianceMemberObject) (flag bool) {
	_, tmemObj := al.getMember(memberObj.GetMemberId())
	if tmemObj != nil {
		return
	}
	flag = true
	al.memberList = append(al.memberList, memberObj)
	if memberObj.position != alliancetypes.AlliancePositionMember {
		al.addManagerMember(memberObj)
	}
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMemberJoin, al, memberObj)
	return
}

func (al *Alliance) removeMember(memberObj *AllianceMemberObject, isClearPlayerData bool) {
	index, tmemObj := al.getMember(memberObj.GetMemberId())
	if tmemObj == nil {
		return
	}
	if memberObj.position != alliancetypes.AlliancePositionMember {
		al.removeManagerMember(memberObj)
	}

	al.memberList = append(al.memberList[:index], al.memberList[index+1:]...)

	eventData := allianceeventtypes.CreateAllianceMemberExitEventData(memberObj.memberId, isClearPlayerData)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMemberExit, al, eventData)

	al.removeDouShen(memberObj)
}

func (al *Alliance) removeDouShen(memberObj *AllianceMemberObject) {
	index, douShen := al.getDouShenMember(memberObj.GetMemberId())
	if douShen == nil {
		return
	}

	al.douShenList = append(al.douShenList[:index], al.douShenList[index+1:]...)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDouShenMemberExit, al, memberObj)

	//当斗神殿为空，马上更新
	if len(al.douShenList) == 0 {
		now := global.GetGame().GetTimeService().Now()
		_ = timeutils.GetHourMs(now)
		al.updateDouShenForceList(now)
	}
}

func (al *Alliance) getManagerMember(memberId int64) (index int32, memObj *AllianceMemberObject) {
	for index, memObj := range al.managerList {
		if memberId == memObj.GetMemberId() {
			return int32(index), memObj
		}
	}
	return -1, nil
}

func (al *Alliance) getDouShenMember(memberId int64) (index int32, memObj *AllianceMemberObject) {
	for index, memObj := range al.douShenList {
		if memberId == memObj.GetMemberId() {
			return int32(index), memObj
		}
	}
	return -1, nil
}

func (al *Alliance) managerChanged(memberObj *AllianceMemberObject) {
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMemberPositionChanged, al, memberObj)

	_, manager := al.getManagerMember(memberObj.GetMemberId())
	if memberObj.GetPosition() == alliancetypes.AlliancePositionMember {
		if manager == nil {
			return
		}
		al.removeManagerMember(memberObj)
	} else {
		if manager != nil {
			return
		}
		al.addManagerMember(memberObj)
	}
}

func (al *Alliance) addManagerMember(memberObj *AllianceMemberObject) {
	if memberObj.position == alliancetypes.AlliancePositionMember {
		return
	}
	_, tmemObj := al.getManagerMember(memberObj.GetMemberId())
	if tmemObj != nil {
		return
	}
	al.managerList = append(al.managerList, memberObj)
}

func (al *Alliance) removeManagerMember(memberObj *AllianceMemberObject) {
	index, tmemObj := al.getManagerMember(memberObj.GetMemberId())
	if tmemObj == nil {
		return
	}
	al.managerList = append(al.managerList[:index], al.managerList[index+1:]...)
}

func (al *Alliance) getApply(applyId int64) (index int32, applyObj *AllianceJoinApplyObject) {
	for index, applyObj := range al.applyList {
		if applyId == applyObj.GetJoinId() {
			return int32(index), applyObj
		}
	}
	return -1, nil
}

func (al *Alliance) addApply(applyObj *AllianceJoinApplyObject) (flag bool) {
	_, tApplyObj := al.getApply(applyObj.GetJoinId())
	if tApplyObj != nil {
		return
	}
	flag = true
	al.applyList = append(al.applyList, applyObj)
	return
}

func (al *Alliance) addInvitation(invitationObj *AllianceInvitationObject) (flag bool) {
	_, tInvitationObj := al.getInvitation(invitationObj.GetInvitationId())
	if tInvitationObj != nil {
		return
	}
	flag = true
	al.invitationList = append(al.invitationList, invitationObj)
	return
}

func (al *Alliance) getInvitation(invitationId int64) (index int32, invitationObj *AllianceInvitationObject) {
	for index, invitationObj := range al.invitationList {
		if invitationId == invitationObj.GetInvitationId() {
			return int32(index), invitationObj
		}
	}
	return -1, nil
}

func (al *Alliance) removeInvitation(invitationObj *AllianceInvitationObject) {
	index, tInvitationObj := al.getInvitation(invitationObj.GetInvitationId())
	if tInvitationObj == nil {
		return
	}
	al.invitationList = append(al.invitationList[:index], al.invitationList[index+1:]...)
}

func (al *Alliance) removeApply(applyObj *AllianceJoinApplyObject) {
	index, tApplyObj := al.getApply(applyObj.GetJoinId())
	if tApplyObj == nil {
		return
	}
	al.applyList = append(al.applyList[:index], al.applyList[index+1:]...)
}

func (al *Alliance) updateForce() {
	totalForce := int64(0)
	for _, memObj := range al.memberList {
		totalForce += memObj.GetForce()
	}
	al.allianceObject.totalForce = totalForce
	al.allianceObject.SetModified()
}

func (al *Alliance) addLog(content string) {

	o := createAllianceLogObject()
	id, err := idutil.GetId()
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()

	o.id = id
	o.allianceId = al.GetAllianceId()
	o.content = content
	o.createTime = now
	o.SetModified()

	al.logList = append(al.logList, o)
	lenOfLog := int32(len(al.logList))
	maxLogSize := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAlliancelogLimit)
	if lenOfLog > maxLogSize {
		startIndex := lenOfLog - maxLogSize
		al.logList = al.logList[startIndex:]
	}
}

func (al *Alliance) appendLog(log *AllianceLogObject) {
	al.logList = append(al.logList, log)
}

func (al *Alliance) updateTransportTimes() error {
	err := al.refreshTimes()
	if err != nil {
		return err
	}

	al.allianceObject.transportTimes += 1
	al.allianceObject.SetModified()
	return nil
}

func (al *Alliance) hasEnoughTransportTimes() bool {
	al.refreshTimes()

	joinTimes := al.allianceObject.transportTimes
	maxTimes := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceTransportationTimes)
	if joinTimes >= maxTimes {
		return false
	}

	return true
}

func (al *Alliance) hasEnoughDepotSlot(itemData *droptemplate.DropItemData) bool {
	itemId := itemData.GetItemId()
	num := itemData.GetNum()
	level := itemData.GetLevel()
	bind := itemData.GetBindType()
	expireType := itemData.GetExpireType()
	// expireTime := itemData.GetExpireTime()

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return false
	}

	needSlot := al.depotBag.countNeedSlotOfItemLevel(itemId, num, level, bind, expireType, itemData.GetExpireTimestamp())
	if needSlot > al.depotBag.GetEmptySlots() {
		return false
	}
	return true
}

func (al *Alliance) GetAllianceBossObject() *AllianceBossObject {
	return al.allianceBossObject
}

//GM重置镖车次数
func (al *Alliance) GMResetTimes() {
	now := global.GetGame().GetTimeService().Now()
	preDay, _ := timeutils.PreDayOfTime(now)
	al.allianceObject.updateTime = preDay
	al.allianceObject.SetModified()

	al.refreshTimes()
}

func (al *Alliance) refreshTimes() (err error) {
	now := global.GetGame().GetTimeService().Now()
	lastUpdateTime := al.allianceObject.lastTransportRefreshTime

	flag, err := timeutils.IsSameFive(lastUpdateTime, now)
	if err != nil {
		return err
	}
	if !flag {
		al.allianceObject.transportTimes = 0
		al.allianceObject.lastTransportRefreshTime = now
		al.allianceObject.SetModified()
	}
	return
}

func (al *Alliance) ifCanAllianceRename() bool {
	now := global.GetGame().GetTimeService().Now()
	limitTime := alliancetemplate.GetAllianceTemplateService().GetAllianceConstantTemp().GetRenameLimitTime()
	return limitTime > now-al.allianceObject.lastMergeTime
}

func createAlliance(obj *AllianceObject) *Alliance {
	all := &Alliance{
		allianceObject: obj,
	}
	return all
}
