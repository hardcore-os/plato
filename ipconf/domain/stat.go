package domain

import "math"

// 对于gateway网关机来说，存在不同时期加入进来的物理机，所以机器的配置是不同的，使用负载来衡量会导致偏差。
// 为更好的应对动态的机器配置变化，我们统计其剩余资源值，来衡量一个机器其是否更适合增加其负载。
// 这里的数值代表的是，此endpoint对应的机器其，自身剩余的资源指标。
type Stat struct {
	ConnectNum   float64 // 业务上，im gateway 总体持有的长连接数量 的剩余值
	MessageBytes float64 // 业务上，im gateway 每秒收发消息的总字节数 的剩余值
}

// 因此假设网络带宽将是系统瓶颈所在，那么哪台机器富余的带宽资源多，哪台机器的负载就是最轻的。
// TODO: 如何预估他的数量级？何时使用静态值衡量
// TODO: json 的解析失效
func (s *Stat) CalculateActiveSorce() float64 {
	return getGB(s.MessageBytes)
}

func (s *Stat) Avg(num float64) {
	s.ConnectNum /= num
	s.MessageBytes /= num
}
func (s *Stat) Clone() *Stat {
	newStat := &Stat{
		MessageBytes: s.MessageBytes,
		ConnectNum:   s.ConnectNum,
	}
	return newStat
}

func (s *Stat) Add(st *Stat) {
	if st == nil {
		return
	}
	s.ConnectNum += st.ConnectNum
	s.MessageBytes += st.MessageBytes
}

func (s *Stat) Sub(st *Stat) {
	if st == nil {
		return
	}
	s.ConnectNum -= st.ConnectNum
	s.MessageBytes -= st.MessageBytes
}

func getGB(m float64) float64 {
	return decimal(m / (1 << 30))
}
func decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}
func min(a, b, c float64) float64 {
	m := func(k, j float64) float64 {
		if k > j {
			return j
		}
		return k
	}
	return m(a, m(b, c))
}
func (s *Stat) CalculateStaticSorce() float64 {
	return s.ConnectNum
}
