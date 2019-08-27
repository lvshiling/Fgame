package proxy

//代理管理器
//一个session一个代理管理器
type ProxyManager struct {
	proxyMapOfMap map[string]map[string]Proxy
}

//添加代理
func (pm *ProxyManager) AddProxy(p Proxy) {
	proxyMap, ok := pm.proxyMapOfMap[p.Service()]
	if !ok {
		proxyMap = make(map[string]Proxy)
		pm.proxyMapOfMap[p.Service()] = proxyMap
	}
	_, ok = proxyMap[p.Id()]
	//TODO 处理重复加入
	if ok {
		panic("never reach here")
	}
	proxyMap[p.Id()] = p
	return
}

//移除代理
func (pm *ProxyManager) RemoveProxy(p Proxy) {
	proxyMap, ok := pm.proxyMapOfMap[p.Service()]
	if !ok {
		return
	}
	delete(proxyMap, p.Id())
}

//获取代理
func (pm *ProxyManager) GetProxy(service string, id string) Proxy {
	proxyMap, ok := pm.proxyMapOfMap[service]
	if !ok {
		return nil
	}
	p, ok := proxyMap[id]
	if !ok {
		return nil
	}
	return p
}

func NewProxyManager() *ProxyManager {
	pm := &ProxyManager{}
	pm.proxyMapOfMap = make(map[string]map[string]Proxy)
	return pm
}
