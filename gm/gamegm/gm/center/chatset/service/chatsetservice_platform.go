package service

import (
	"fgame/fgame/gm/gamegm/common"
	constant "fgame/fgame/gm/gamegm/constant"
	gmError "fgame/fgame/gm/gamegm/error"
	chatSetmodel "fgame/fgame/gm/gamegm/gm/center/chatset/model"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

func (m *chatSetService) AddChatSetPlatform(p_platformid int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_teamVip int, p_teamPlayerLevel int) error {
	if p_platformid < 0 {
		return gmError.GetError(gmError.ErrorCodeChatSetEmpty)
	}

	exflag, err := m.existsChatSetPlatform(p_platformid)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeChatSetExist)
	}

	modelInfo := &chatSetmodel.PlatformChatSetInfo{
		PlatformId:       p_platformid,
		WorldVip:         p_worldVip,
		WorldPlayerLevel: p_worldPlayerLevel,
		PChatVip:         p_pChatVip,
		PChatPlayerLevel: p_pChatPlayerLevel,
		GuildVip:         p_guildVip,
		GuildPlayerLevel: p_guildPlayerLevel,
		TeamVip:          p_teamVip,
		TeamPlayerLevel:  p_teamPlayerLevel,
	}
	err = m.saveChatSetPlatform(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *chatSetService) UpdateChatSetPlatform(p_id int, p_platformid int, p_worldVip int, p_worldPlayerLevel int, p_pChatVip int, p_pChatPlayerLevel int, p_guildVip int, p_guildPlayerLevel int, p_teamVip int, p_teamPlayerLevel int) error {
	if p_platformid < 0 {
		return gmError.GetError(gmError.ErrorCodeChatSetEmpty)
	}

	exflag, err := m.existsChatSetWithIdPlatform(p_id, p_platformid)
	if err != nil {
		return err
	}
	if exflag {
		return gmError.GetError(gmError.ErrorCodeChatSetExist)
	}

	modelInfo := &chatSetmodel.PlatformChatSetInfo{
		Id:               p_id,
		PlatformId:       p_platformid,
		WorldVip:         p_worldVip,
		WorldPlayerLevel: p_worldPlayerLevel,
		PChatVip:         p_pChatVip,
		PChatPlayerLevel: p_pChatPlayerLevel,
		GuildVip:         p_guildVip,
		GuildPlayerLevel: p_guildPlayerLevel,
		TeamVip:          p_teamVip,
		TeamPlayerLevel:  p_teamPlayerLevel,
	}
	err = m.saveChatSetPlatform(modelInfo)
	if err != nil {
		return err
	}
	return nil
}

func (m *chatSetService) DeleteChatSetPlatform(p_chatSetid int64) error {
	now := timeutils.TimeToMillisecond(time.Now())
	errdb := m.db.DB().Table("t_platform_chatset").Where("id = ?", p_chatSetid).Update("deleteTime", now)
	if errdb.Error != nil {
		log.WithFields(log.Fields{
			"chatSetid": p_chatSetid,
			"error":     errdb.Error,
		}).Error("删除聊天配置失败")
		return errdb.Error
	}
	return nil
}

func (m *chatSetService) GetChatSetPlatformById(p_id int) (*chatSetmodel.PlatformChatSetInfo, error) {
	info := &chatSetmodel.PlatformChatSetInfo{}
	exdb := m.db.DB().Where("id = ? and deleteTime = 0", p_id).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_id,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSet")
		return nil, exdb.Error
	}
	return info, nil
}

func (m *chatSetService) GetChatSetPlatform(p_platformid int) (*chatSetmodel.PlatformChatSetInfo, error) {
	info := &chatSetmodel.PlatformChatSetInfo{}
	exdb := m.db.DB().Where("platformId = ? and deleteTime = 0", p_platformid).First(info)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_platformid,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSet")
		return nil, exdb.Error
	}
	return info, nil
}

func (m *chatSetService) GetChatSetListPlatform(p_platformid int, p_index int, p_platformList []int64) ([]*chatSetmodel.PlatformChatSetInfo, error) {
	return m.getChatSetListPlatform(p_platformid, p_index, p_platformList)
}

func (m *chatSetService) GetChatSetCountPlatform(p_platformid int, p_platformList []int64) (int, error) {
	return m.getChatSetCountPlatform(p_platformid, p_platformList)
}

func (m *chatSetService) existsChatSetPlatform(p_platformid int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform_chatset").Where("platformId = ? and deleteTime = 0", p_platformid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_platformid,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSetPlatform")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *chatSetService) saveChatSetPlatform(p_info *chatSetmodel.PlatformChatSetInfo) error {
	exdb := m.db.DB().Save(p_info)
	now := timeutils.TimeToMillisecond(time.Now())
	if p_info.Id == 0 {
		p_info.CreateTime = now
	} else {
		p_info.UpdateTime = now
	}
	if exdb.Error != nil {
		log.WithFields(log.Fields{
			"chatSetid": p_info.PlatformId,
			"error":     exdb.Error,
		}).Error("保存聊天配置失败")
		return exdb.Error
	}
	return nil
}

func (m *chatSetService) existsChatSetWithIdPlatform(p_chatSetid int, p_platformid int) (bool, error) {
	count := 0
	exdb := m.db.DB().Table("t_platform_chatset").Where("id != ? and platformId = ? and deleteTime = 0", p_chatSetid, p_platformid).Count(&count)
	if exdb.Error != nil && exdb.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"platformId": p_platformid,
			"ID":         p_chatSetid,
			"error":      exdb.Error,
		}).Error("查询聊天配置信息失败existsChatSetWithIdPlatform")
		return false, exdb.Error
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m *chatSetService) getChatSetListPlatform(p_platformid int, p_index int, p_platformList []int64) ([]*chatSetmodel.PlatformChatSetInfo, error) {
	offset := (p_index - 1) * constant.DefaultPageSize
	if offset < 0 {
		offset = 0
	}
	limit := constant.DefaultPageSize
	rst := make([]*chatSetmodel.PlatformChatSetInfo, 0)
	where := ""
	if p_platformid > 0 {
		where += fmt.Sprintf(" and platformId=%d", p_platformid)
	}

	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and platformId IN (%s)", common.CombinInt64Array(p_platformList))
	}

	exerr := m.db.DB().Where("deleteTime =0" + where).Order("platformId ASC").Offset(offset).Limit(limit).Find(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取聊天配置列表失败")
		return nil, exerr.Error
	}
	return rst, nil
}

func (m *chatSetService) getChatSetCountPlatform(p_platformid int, p_platformList []int64) (int, error) {
	rst := 0
	where := ""
	if p_platformid > 0 {
		where += fmt.Sprintf(" and platformId=%d", p_platformid)
	}
	if len(p_platformList) > 0 {
		where += fmt.Sprintf(" and platformId IN (%s)", common.CombinInt64Array(p_platformList))
	}

	exerr := m.db.DB().Table("t_platform_chatset").Where("deleteTime =0" + where).Count(&rst)
	if exerr.Error != nil && exerr.Error != gorm.ErrRecordNotFound {
		log.WithFields(log.Fields{
			"error": exerr.Error,
		}).Error("获取聊天配置列表失败")
		return 0, exerr.Error
	}
	return rst, nil
}
