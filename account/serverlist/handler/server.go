package handler

// func init() {
// 	dispatch.Register(codec.MessageType(uipb.MessageType_CS_SERVER_LIST_TYPE), dispatch.HandlerFunc(handleServerList))
// }

// type server struct {
// 	Id   int32
// 	Name string
// 	Ip   string
// }

// var (
// 	serverList = []*server{
// 		&server{
// 			Id:   1,
// 			Name: "内网测试",
// 			Ip:   "192.168.1.13:3000",
// 		},
// 		&server{
// 			Id:   2,
// 			Name: "张荣昌测试",
// 			Ip:   "192.168.1.7:3000",
// 		},
// 		&server{
// 			Id:   3,
// 			Name: "杨耒州测试",
// 			Ip:   "192.168.1.20:3000",
// 		},
// 		&server{
// 			Id:   4,
// 			Name: "张荣昌调试",
// 			Ip:   "192.168.1.7:3001",
// 		},
// 		&server{
// 			Id:   5,
// 			Name: "外网测试",
// 			Ip:   "47.104.31.33:3000",
// 		},
// 		&server{
// 			Id:   6,
// 			Name: "姚中扬",
// 			Ip:   "192.168.1.2:3000",
// 		},
// 		&server{
// 			Id:   7,
// 			Name: "杨耒州调试",
// 			Ip:   "192.168.1.20:3001",
// 		},
// 	}
// )

// //获取服务器列表
// func handleServerList(s session.Session, msg interface{}) error {
// 	log.Debug("serverlist:处理服务器列表")

// 	scServerList := buildServerList()
// 	msgBytes, err := processor.Encode(scServerList)
// 	if err != nil {
// 		return err
// 	}
// 	err = s.Send(msgBytes)
// 	if err != nil {
// 		return err
// 	}
// 	log.Debug("serverlist:处理服务器列表完成")
// 	return nil
// }

// func buildServerList() *uipb.SCServerList {
// 	scServerList := &uipb.SCServerList{}
// 	for _, server := range global.GetGlobal().GetServerList() {
// 		s := &uipb.Server{}
// 		s.Id = &server.Id
// 		s.Name = &server.Name
// 		s.Ip = &server.Ip
// 		scServerList.Servers = append(scServerList.Servers, s)
// 	}
// 	return scServerList
// }
