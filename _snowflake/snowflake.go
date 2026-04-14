package _snowflake

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_time"
	"sync"
	"time"
)

type Snowflake struct {
	epochMs       int64 // 1580608922000
	nodeBit       uint8 // 10
	sequenceBit   uint8 // 12
	nodeId        int64
	maxNodeID     int64 // [0,1023]
	maxSequence   int64 // [0,4095]
	lastTimestamp int64
	sequence      int64
	nodeShift     uint8
	timeShift     uint8
	mu            sync.Mutex
}

func New(epochMs int64, nodeBit uint8, sequenceBit uint8, nodeId int64) *Snowflake {
	_interceptor.
		Insure(epochMs >= 0).
		Message("epochMs 不能小于 0").
		Data(map[string]interface{}{
			"epochMs": epochMs,
		}).
		Do()
	_interceptor.
		Insure(int(nodeBit)+int(sequenceBit) < 63).
		Message("nodeBit 与 sequenceBit 之和必须小于 63").
		Data(map[string]interface{}{
			"nodeBit":     nodeBit,
			"sequenceBit": sequenceBit,
		}).
		Do()
	maxNodeID := (int64(1) << nodeBit) - 1
	maxSequence := (int64(1) << sequenceBit) - 1
	_interceptor.
		Insure(nodeId >= 0 && nodeId <= maxNodeID).
		Message("nodeId 超出可配置范围").
		Data(map[string]interface{}{
			"nodeId":    nodeId,
			"maxNodeID": maxNodeID,
		}).
		Do()
	return &Snowflake{
		epochMs:       epochMs,
		nodeBit:       nodeBit,
		sequenceBit:   sequenceBit,
		nodeId:        nodeId,
		maxNodeID:     maxNodeID,
		maxSequence:   maxSequence,
		lastTimestamp: -1,
		nodeShift:     sequenceBit,
		timeShift:     nodeBit + sequenceBit,
	}
}
func (this *Snowflake) Id() int64 {
	return this.IdList(1)[0]
}
func (this *Snowflake) IdList(count int) []int64 {
	_interceptor.
		Insure(count > 0).
		Message("count必须大于0").
		Data(map[string]interface{}{
			"count": count,
		}).
		Do()
	this.mu.Lock()
	defer this.mu.Unlock()
	idList := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		idList = append(idList, this.nextIDLocked())
	}
	return idList
}
func (this *Snowflake) nextIDLocked() int64 {
	now := time.Now().UnixMilli()
	_interceptor.
		Insure(now >= this.epochMs).
		Message("当前时间早于起始时间").
		Data(map[string]interface{}{
			"now":     now,
			"epochMs": this.epochMs,
		}).
		Do()
	if this.lastTimestamp >= 0 && now < this.lastTimestamp {
		now = _time.WaitUntilMilli(this.lastTimestamp)
	}
	if now == this.lastTimestamp {
		this.sequence = (this.sequence + 1) & this.maxSequence
		if this.sequence == 0 {
			now = _time.WaitNextMilli(this.lastTimestamp)
		}
	} else {
		this.sequence = 0
	}
	this.lastTimestamp = now
	return ((now - this.epochMs) << this.timeShift) | (this.nodeId << this.nodeShift) | this.sequence
}
