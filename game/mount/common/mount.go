package common

import "fgame/fgame/game/mount/mount"

//坐骑皮肤信息
type MountSkinInfo struct {
	MountId int32 `json:"mountId"`
	Level   int32 `json:"level"`
	UpPro   int32 `json:"upPro"`
}

//坐骑信息
type MountInfo struct {
	AdvanceId   int              `json:"advanceId"`
	MountId     int32            `json:"mountId"`
	UnrealLevel int32            `json:"unrealLevel"`
	UnrealPro   int32            `json:"unrealPro"`
	CulLevel    int32            `json:"culLevel"`
	CulPro      int32            `json:"culPro"`
	SkinList    []*MountSkinInfo `json:"skinList"`
}

func (m *MountInfo) GetMountId() int32 {
	if m.MountId == 0 {
		mountTemplate := mount.GetMountService().GetMountNumber(int32(m.AdvanceId))
		if mountTemplate != nil {
			return int32(mountTemplate.TemplateId())
		}
	}
	return m.MountId
}
