package api

type State struct {
	NodeServer  *nodeServer             `json:"NodeServer"`
	Dispatch    map[string]*dispatch    `json:"Dispatch"`
	GateServer  map[string]*gateServer  `json:"GateServer"`
	GameServer  map[string]*gameServer  `json:"GameServer"`
	MultiServer map[string]*multiServer `json:"MultiServer"`
	MServer     map[string]*muipServer  `json:"MuipServer"`
}

type nodeServer struct {
	AppId string `json:"AppId"`
}
type dispatch struct {
	AppId string `json:"AppId"`
}
type gameServer struct {
	AppId     string `json:"AppId"`
	PlayerNum int64  `json:"PlayerNum"`
}
type gateServer struct {
	AppId     string `json:"AppId"`
	PlayerNum int64  `json:"PlayerNum"`
}
type multiServer struct {
	AppId string `json:"AppId"`
}
type muipServer struct {
	AppId string `json:"AppId"`
}

// func (a *muip.Api) State(c *gin.Context) {
// 	allService := a.muip.getAllService()
// 	state := &State{
// 		Dispatch:    make(map[string]*dispatch),
// 		GateServer:  make(map[string]*gateServer),
// 		GameServer:  make(map[string]*gameServer),
// 		MultiServer: make(map[string]*multiServer),
// 		MServer:     make(map[string]*muipServer),
// 	}
// 	for stype, serviceList := range allService {
// 		for _, service := range serviceList {
// 			switch stype {
// 			case spb.ServerType_SERVICE_DISPATCH:
// 				state.Dispatch[service.AppId] = &dispatch{AppId: service.AppId}
// 			case spb.ServerType_SERVICE_GATE:
// 				state.GateServer[service.AppId] = &gateServer{AppId: service.AppId, PlayerNum: service.PlayerNum}
// 			case spb.ServerType_SERVICE_GAME:
// 				state.GameServer[service.AppId] = &gameServer{AppId: service.AppId, PlayerNum: service.PlayerNum}
// 			case spb.ServerType_SERVICE_MULTI:
// 				state.MultiServer[service.AppId] = &multiServer{AppId: service.AppId}
// 			case spb.ServerType_SERVICE_MUIP:
// 				state.MServer[service.AppId] = &muipServer{AppId: service.AppId}
// 			}
// 		}
// 	}
// 	c.IndentedJSON(200, state)
// }
