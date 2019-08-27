package login

import (
	"fgame/fgame/client/session"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func AuthLoginServer(s session.Session, userName string) {
	csAccountLogin := &uipb.CSAccountLogin{}
	devicePlatform := int32(1)
	csAccountLogin.DevicePlatform = &devicePlatform
	platform := int32(1)
	csAccountLogin.Platform = &platform
	pcLoginData := &uipb.PCLoginData{}
	pcLoginData.Name = &userName
	csAccountLogin.PcLoginData = pcLoginData
	s.Send(csAccountLogin)
}

func AuthGameServer(s session.Session, serverId int32, token string) {
	csLogin := &uipb.CSLogin{}
	csLogin.Token = &token
	csLogin.ServerId = &serverId
	s.Send(csLogin)
}
