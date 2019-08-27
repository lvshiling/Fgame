package player

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/core/storage"
	playerentity "fgame/fgame/game/player/entity"
	"fgame/fgame/game/player/types"

	"github.com/pkg/errors"
)

//TODO 改成私有变量
type PlayerObject struct {
	player                   *Player
	Id                       int64
	UserId                   int64
	Name                     string
	ServerId                 int32
	OriginServerId           int32
	Role                     types.RoleType
	Sex                      types.SexType
	LastLoginTime            int64
	LastLogoutTime           int64
	OnlineTime               int64
	TotalOnlineTime          int64
	TodayOnlineTime          int64
	Forbid                   int32
	ForbidText               string
	ForbidName               string
	ForbidTime               int64
	ForbidEndTime            int64
	ForbidChat               int32
	ForbidChatText           string
	ForbidChatName           string
	ForbidChatTime           int64
	ForbidChatEndTime        int64
	IgnoreChat               int32
	IgnoreChatText           string
	IgnoreChatName           string
	IgnoreChatTime           int64
	IgnoreChatEndTime        int64
	OfflineTime              int64
	IsOpenVideo              int32
	PrivilegeType            types.PrivilegeType
	TotalChargeMoney         int64
	TotalChargeGold          int64
	TotalPrivilegeChargeGold int64
	Online                   int32
	GetNewReward             int32
	SystemCompensate         int32
	SdkType                  logintypes.SDKType
	Ip                       string
	TodayChargeMoney         int64
	YesterdayChargeMoney     int64
	ChargeTime               int64
	UpdateTime               int64
	CreateTime               int64
	DeleteTime               int64
}

func NewPlayerObject(p *Player) *PlayerObject {
	po := &PlayerObject{}
	po.player = p

	return po
}

func (po *PlayerObject) GetPlayerId() int64 {
	return po.Id
}

func (po *PlayerObject) GetDBId() int64 {
	return po.Id
}

func (po *PlayerObject) ToEntity() (e storage.Entity, err error) {
	pe := &playerentity.PlayerEntity{}
	pe.Id = po.Id
	pe.UserId = po.UserId
	pe.ServerId = po.ServerId
	pe.OriginServerId = po.OriginServerId
	pe.Name = po.Name
	pe.Role = int32(po.Role)
	pe.Sex = int32(po.Sex)
	pe.LastLoginTime = po.LastLoginTime
	pe.OnlineTime = po.OnlineTime
	pe.OfflineTime = po.OfflineTime
	pe.LastLogoutTime = po.LastLogoutTime
	pe.TotalOnlineTime = po.TotalOnlineTime
	pe.TodayOnlineTime = po.TodayOnlineTime
	pe.Forbid = po.Forbid
	pe.ForbidText = po.ForbidText
	pe.ForbidName = po.ForbidName
	pe.ForbidTime = po.ForbidTime
	pe.ForbidEndTime = po.ForbidEndTime
	pe.ForbidChat = po.ForbidChat
	pe.ForbidChatText = po.ForbidChatText
	pe.ForbidChatName = po.ForbidChatName
	pe.ForbidChatTime = po.ForbidChatTime
	pe.ForbidChatEndTime = po.ForbidChatEndTime
	pe.IgnoreChat = po.IgnoreChat
	pe.IgnoreChatText = po.IgnoreChatText
	pe.IgnoreChatName = po.IgnoreChatName
	pe.IgnoreChatTime = po.IgnoreChatTime
	pe.IgnoreChatEndTime = po.IgnoreChatEndTime
	pe.IsOpenVideo = po.IsOpenVideo
	pe.PrivilegeType = int32(po.PrivilegeType)
	pe.TotalChargeGold = po.TotalChargeGold
	pe.TotalChargeMoney = po.TotalChargeMoney
	pe.TotalPrivilegeChargeGold = po.TotalPrivilegeChargeGold
	pe.GetNewReward = po.GetNewReward
	pe.SystemCompensate = po.SystemCompensate
	pe.Online = po.Online
	pe.SdkType = int32(po.SdkType)
	pe.Ip = po.Ip
	pe.TodayChargeMoney = po.TodayChargeMoney
	pe.YesterdayChargeMoney = po.YesterdayChargeMoney
	pe.ChargeTime = po.ChargeTime
	pe.UpdateTime = po.UpdateTime
	pe.CreateTime = po.CreateTime
	pe.DeleteTime = po.DeleteTime
	e = pe
	return
}

func (po *PlayerObject) FromEntity(e storage.Entity) (err error) {
	pe, _ := e.(*playerentity.PlayerEntity)
	po.Id = pe.Id
	po.UserId = pe.UserId
	po.Name = pe.Name
	po.ServerId = pe.ServerId
	po.OriginServerId = pe.OriginServerId
	po.Role = types.RoleType(pe.Role)
	po.Sex = types.SexType(pe.Sex)
	po.LastLoginTime = pe.LastLoginTime
	po.OnlineTime = pe.OnlineTime
	po.OfflineTime = pe.OfflineTime
	po.LastLogoutTime = pe.LastLogoutTime
	po.TotalOnlineTime = pe.TotalOnlineTime
	po.TodayOnlineTime = pe.TodayOnlineTime
	po.Forbid = pe.Forbid
	po.ForbidText = pe.ForbidText
	po.ForbidName = pe.ForbidName
	po.ForbidTime = pe.ForbidTime
	po.ForbidChat = pe.ForbidChat
	po.ForbidEndTime = pe.ForbidEndTime
	po.ForbidChatText = pe.ForbidChatText
	po.ForbidChatName = pe.ForbidChatName
	po.ForbidChatTime = pe.ForbidChatTime
	po.ForbidChatEndTime = pe.ForbidChatEndTime
	po.IgnoreChat = pe.IgnoreChat
	po.IgnoreChatText = pe.IgnoreChatText
	po.IgnoreChatName = pe.IgnoreChatName
	po.IgnoreChatTime = pe.IgnoreChatTime
	po.IgnoreChatEndTime = pe.IgnoreChatEndTime
	po.IsOpenVideo = pe.IsOpenVideo
	po.PrivilegeType = types.PrivilegeType(pe.PrivilegeType)
	po.TotalChargeGold = pe.TotalChargeGold
	po.TotalChargeMoney = pe.TotalChargeMoney
	po.TotalPrivilegeChargeGold = pe.TotalPrivilegeChargeGold
	po.Online = pe.Online
	po.GetNewReward = pe.GetNewReward
	po.SystemCompensate = pe.SystemCompensate
	po.SdkType = logintypes.SDKType(pe.SdkType)
	po.Ip = pe.Ip
	po.TodayChargeMoney = pe.TodayChargeMoney
	po.YesterdayChargeMoney = pe.YesterdayChargeMoney
	po.ChargeTime = pe.ChargeTime
	po.UpdateTime = pe.UpdateTime
	po.CreateTime = pe.CreateTime
	po.DeleteTime = pe.DeleteTime
	return
}

func (po *PlayerObject) SetModified() {
	e, err := po.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Player"))
	}
	obj, _ := e.(types.PlayerDataEntity)
	if obj == nil {
		panic("never reach here")
	}
	po.player.AddChangedObject(obj)
	return
}
