package lb

// load balance
// 负载均衡器接口
type Balancer interface {
	Next(key string, hosts []*ServerInstance) *ServerInstance
}
