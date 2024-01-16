package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SequenceService struct {
	mutex     sync.Mutex
	bufferA   *Segment
	bufferB   *Segment
	isSwapped bool
}

type Segment struct {
	minID int64
	maxID int64
	curID int64
	ready bool
}

// Call updateSegmentInBackground with the correct step value
func NewSequenceService(start int64, step int64) *SequenceService {
	ss := &SequenceService{
		bufferA:   &Segment{minID: start, maxID: start + step - 1, curID: start, ready: true},                  // maxID should be start + step - 1
		bufferB:   &Segment{minID: start + step, maxID: start + 2*step - 1, curID: start + step, ready: false}, // maxID should be start + 2*step - 1
		isSwapped: false,
	}
	go ss.updateSegmentInBackground(step) // Pass the step value to the background function
	return ss
}

// 模拟从ZK中获取号段
var globalSequenceID int64 = 3000000000

func (ss *SequenceService) getSequenceIDFromExternalService(step int64) (int64, int64, error) {
	time.Sleep(1 * time.Second) // Simulate network delay
	// Get the current sequence ID and atomically update the global sequence ID
	initSequenceID := atomic.LoadInt64(&globalSequenceID)
	atomic.AddInt64(&globalSequenceID, step)
	// fmt.Println(initSequenceID, globalSequenceID)
	// The new range will start from the current sequence ID and end at current sequence ID + step - 1
	return initSequenceID, initSequenceID + step - 1, nil
}

func (ss *SequenceService) updateSegmentInBackground(step int64) {
	for {
		time.Sleep(1 * time.Second) // Sleep for a shorter duration to update more frequently
		ss.mutex.Lock()
		var curSegment *Segment
		if ss.isSwapped {
			curSegment = ss.bufferA
		} else {
			curSegment = ss.bufferB
		}
		if !curSegment.ready {
			minID, maxID, err := ss.getSequenceIDFromExternalService(step)
			if err != nil {
				fmt.Println("Error:", err)
				ss.mutex.Unlock()
				continue
			}
			curSegment.minID = minID
			curSegment.maxID = maxID
			curSegment.curID = minID // Start allocating from minID
			curSegment.ready = true
		} else {
			// Set the current segment's ready flag to false so that it will be updated in the next iteration
			curSegment.ready = false
		}
		ss.mutex.Unlock()
	}
}
func (ss *SequenceService) getId() (int64, error) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock() // Use defer to ensure mutex is always unlocked

	var curSegment *Segment
	if !ss.isSwapped {
		curSegment = ss.bufferA
	} else {
		curSegment = ss.bufferB
	}

	if curSegment.curID <= curSegment.maxID {
		id := curSegment.curID
		curSegment.curID++
		return id, nil
	}

	// Current segment is exhausted, switch to the other buffer if it's ready
	otherSegment := ss.bufferA
	if !ss.isSwapped {
		otherSegment = ss.bufferB
	}

	if otherSegment.ready && otherSegment.curID <= otherSegment.maxID {
		ss.isSwapped = !ss.isSwapped
		curSegment.ready = false // Mark the current segment as not ready
		id := otherSegment.curID
		otherSegment.curID++
		return id, nil
	}

	// If the other segment is not ready, return an error or wait for it to become ready
	return 0, fmt.Errorf("no ID available, waiting for segment to become ready")
}
func main() {
	ss := NewSequenceService(3000000000, 1000)
	var id int64
	var err error
	for i := 0; i < 2000; i++ {
		id, err = ss.getId()
		if err != nil {
			fmt.Println("Error:", err)
			break // Exit the loop if we run out of IDs
		} else {
			//fmt.Println(id)
		}
	}
	fmt.Println(id)
}

/*
总结: 这段代码是双缓冲获取序列号的一个框架代码
利用chatgpt写一个demo，但是还是有点问题，后面再排查定位下
*/
