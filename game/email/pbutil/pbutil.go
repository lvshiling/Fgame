package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	babypbutil "fgame/fgame/game/baby/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
	email "fgame/fgame/game/email/player"
)

func BuildSCEmailsGet(emailObjMap []*email.PlayerEmailObject) *uipb.SCEmailsGet {
	scEmailsGet := &uipb.SCEmailsGet{}
	if len(emailObjMap) > 0 {

		scEmailsGetArr := make([]*uipb.EmailInfo, 0, len(emailObjMap))
		for _, emailObj := range emailObjMap {
			scEmailsGetArr = append(scEmailsGetArr, buildEmailInfo(emailObj))
		}
		scEmailsGet.EmailInfo = scEmailsGetArr
	}
	return scEmailsGet
}

func buildEmailInfo(emailObject *email.PlayerEmailObject) *uipb.EmailInfo {
	emailId := emailObject.GetEmailId()
	isRead := emailObject.GetIsRead()
	isGetAttachment := emailObject.GetIsGetAttachment()
	title := emailObject.GetTitle()
	content := emailObject.GetContent()
	createTime := emailObject.GetCreateTime()
	attachmentInfo := emailObject.GetAttachmentInfo()

	scEmailInfo := &uipb.EmailInfo{}
	scEmailInfo.EmailId = &emailId
	scEmailInfo.IsRead = &isRead
	scEmailInfo.IsGetAttachment = &isGetAttachment
	scEmailInfo.Title = &title
	scEmailInfo.Content = &content
	scEmailInfo.CreateTime = &createTime
	if len(attachmentInfo) > 0 {
		scEmailInfo.AttachementInfo = buildAttachmentList(attachmentInfo)
	}

	return scEmailInfo
}

func buildAttachmentList(attachmentInfo []*droptemplate.DropItemData) (attachList []*uipb.AttachementInfo) {
	for _, itemData := range attachmentInfo {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()
		bindType := int32(itemData.GetBindType())
		upstar := itemData.GetUpstar()
		attrList := itemData.GetAttrList()
		isRandomAttr := itemData.GetIsRandomAttr()
		expireType := int32(itemData.GetExpireType())
		expireTime := itemData.GetExpireTimestamp()
		itemGetTime := itemData.GetItemGetTime()
		quality := itemData.GetQuality()
		danbei := itemData.GetDanbei()
		talentList := itemData.GetTalentList()
		sex := int32(itemData.GetSex())
		openLightLevel := int32(itemData.GetOpenLightLevel())
		openTimes := int32(itemData.GetOpenTimes())
		info := &uipb.AttachementInfo{}
		info.ItemId = &itemId
		info.Num = &num
		info.Level = &level
		info.Upstar = &upstar
		info.AttrList = attrList
		info.BindType = &bindType
		info.IsRandomAttr = &isRandomAttr
		info.ExpireType = &expireType
		info.ExpireTime = &expireTime
		info.ItemGetTime = &itemGetTime
		info.Quality = &quality
		info.Sex = &sex
		info.TalentList = babypbutil.BuildTalentInfoList(talentList)
		info.Danbei = &danbei

		info.OpenLightLevel = &openLightLevel
		info.OpenTimes = &openTimes
		attachList = append(attachList, info)
	}

	return attachList
}

func BuildSCReadEmail(emailId int64) *uipb.SCReadEmail {
	scReadEmail := &uipb.SCReadEmail{}
	scReadEmail.EmailId = &emailId

	return scReadEmail
}

func BuildSCDelHadReadEmail(hadDelEmailIdArr []int64) *uipb.SCDelHadReadEmail {
	scDelHadReadEmail := &uipb.SCDelHadReadEmail{}
	scDelHadReadEmail.EmailIdList = hadDelEmailIdArr

	return scDelHadReadEmail
}

func BuildSCDelEmail(emailId int64) *uipb.SCDelEmail {
	scDelEmail := &uipb.SCDelEmail{}
	scDelEmail.EmailId = &emailId

	return scDelEmail
}

func BuildSCGetAttachment(emailId int64, totalItemList []*droptemplate.DropItemData) *uipb.SCGetAttachment {
	scGetAttachment := &uipb.SCGetAttachment{}
	scGetAttachment.EmailId = &emailId
	scGetAttachment.AttachementInfo = buildAttachmentList(totalItemList)

	return scGetAttachment
}

func BuildSCGetAttachmentBatch(emailIdArr []int64, totalItemList []*droptemplate.DropItemData) *uipb.SCGetAttachmentBatch {
	scGetAttachmentBatch := &uipb.SCGetAttachmentBatch{}
	scGetAttachmentBatch.AttachementInfo = buildAttachmentList(totalItemList)
	scGetAttachmentBatch.EmailIdList = emailIdArr

	return scGetAttachmentBatch
}

func BuildSCAddEmail(emailObject *email.PlayerEmailObject) *uipb.SCAddEmail {
	scAddEmail := &uipb.SCAddEmail{}
	scAddEmail.EmailInfo = buildEmailInfo(emailObject)

	return scAddEmail
}
