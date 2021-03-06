package lb

import "hash/crc32"

// hash 负载均衡：常常被用在“请求联系”中
// 基于 Hash 算法的负载均衡器
var _ Balancer = new(HashBalancer)

type HashBalancer struct {
}

func (h *HashBalancer) Next(key string, hosts []*ServerInstance) *ServerInstance {
	if len(hosts) == 0 {
		return nil
	}
	// hash
	count := crc32.ChecksumIEEE([]byte(key))
	//取模计算索引
	index := int(count) % len(hosts)
	//按照索引取出实例
	instance := hosts[index]
	return instance
}
