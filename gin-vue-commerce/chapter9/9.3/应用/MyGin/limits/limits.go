package limits

import (
	"golang.org/x/time/rate"
	"sync"
)

// 定义限流对象
type LimiterItem struct {
	// 需要限流的IP地址
	IP string
	// 限流对象
	Limter *rate.Limiter
}

// 定义限流对象集合
type Limiter struct {
	LmiterSlice []*LimiterItem
	Locker      sync.RWMutex
}

// 实例化限流对象集合
func NewFlowLimiter() *Limiter {
	return &Limiter{
		LmiterSlice: []*LimiterItem{},
		Locker:      sync.RWMutex{},
	}
}

// 获取服务或接口的限流对象
// serverName:服务或接口标识
// qps:服务或接口限流最大并发数
func (counter *Limiter) GetLimiter(ip string, qps float64) *rate.Limiter {
	// 从已有限流对象集合获取某个服务或接口的限流对象
	for _, item := range counter.LmiterSlice {
		if item.IP == ip {
			return item.Limter
		}
	}

	// 如果已有限流对象集合没有，则新建限流对象
	newLimiter := rate.NewLimiter(rate.Limit(qps), int(qps*3))
	item := &LimiterItem{
		IP:     ip,
		Limter: newLimiter,
	}
	// 新建限流对象写入已有限流对象集合
	counter.LmiterSlice = append(counter.LmiterSlice, item)
	counter.Locker.Lock()
	defer counter.Locker.Unlock()
	// 返回新建限流对象
	return newLimiter
}

// 初始化限流对象集合，便于调用
var LimiterHandler *Limiter

func init() {
	LimiterHandler = NewFlowLimiter()
}
