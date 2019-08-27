package marry

// import (
// 	"fgame/fgame/game/center/center"
// 	gamecentertypes "fgame/fgame/game/center/types"
// 	marrytemplate "fgame/fgame/game/marry/template"
// 	marrytypes "fgame/fgame/game/marry/types"
// 	"sync"
// )

// //结婚收费设置的缓存服务数据
// type MarrySetService interface {
// 	ResetHouTaiType(htType marrytypes.MarryHoutaiType)
// 	GetHouTaiType() marrytypes.MarryHoutaiType
// }

// type marrySetService struct {
// 	//读写锁
// 	rwm sync.RWMutex
// }

// func (m *marrySetService) init() {
// 	marryKindType := center.GetCenterService().GetMarryKindType()
// 	houtaiType := marrytypes.MarryHoutaiTypeCommon
// 	switch marryKindType {
// 	case gamecentertypes.MarryPriceTypeCheap:
// 		houtaiType = marrytypes.MarryHoutaiTypeCheep
// 		break
// 	default:
// 		houtaiType = marrytypes.MarryHoutaiTypeCommon
// 	}
// 	marrytemplate.GetMarryTemplateService().SetHouTaiType(houtaiType)
// }

// func (m *marrySetService) ResetHouTaiType(htType marrytypes.MarryHoutaiType) {
// 	m.rwm.Lock()
// 	defer m.rwm.Unlock()
// 	marrytemplate.GetMarryTemplateService().SetHouTaiType(htType)
// }

// func (m *marrySetService) GetHouTaiType() marrytypes.MarryHoutaiType {
// 	m.rwm.Lock()
// 	defer m.rwm.Unlock()
// 	htType := marrytemplate.GetMarryTemplateService().GetHouTaiType()
// 	return htType
// }
