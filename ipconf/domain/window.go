package domain

const (
	windowSize = 5 // 防止流量毛刺，取5个点的平均值
)

type stateWindow struct {
	stateQueue []*Stat
	statChan   chan *Stat
	sumStat    *Stat
	idx        int64
}

func newStateWindow() *stateWindow {
	return &stateWindow{
		stateQueue: make([]*Stat, windowSize),
		statChan:   make(chan *Stat),
		sumStat:    &Stat{},
	}
}

func (sw *stateWindow) getStat() *Stat {
	res := sw.sumStat.Clone() // 因为sumStat是指针，所以需要clone一份计算均值后返回
	res.Avg(windowSize)
	return res
}

func (sw *stateWindow) appendStat(s *Stat) {
	// 窗口大小没达到windowSize直接入窗，否则减去窗口最前面的state
	if sw.idx >= windowSize {
		sw.sumStat.Sub(sw.stateQueue[sw.idx%windowSize])
	}
	// 更新最新的stat
	sw.stateQueue[sw.idx%windowSize] = s
	// 计算最新的窗口和
	sw.sumStat.Add(s)
	sw.idx++
}
