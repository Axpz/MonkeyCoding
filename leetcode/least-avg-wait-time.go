package leetcode

import (
	"container/heap"
)

type Task struct {
	reqTime int
	durTime int
	index   int
}

type MinHeap []Task

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	iDurTime, jDurTime := (*h)[i].durTime, (*h)[j].durTime
	if iDurTime == jDurTime {
		iReqTime, jReqTime := (*h)[i].reqTime, (*h)[j].reqTime
		return iReqTime < jReqTime
	}

	return iDurTime < jDurTime
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Pop() any {
	if h.Len() <= 0 {
		return nil
	}

	item := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return item
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func calShortestJobFirst(reqTimes, durTimes []int) float64 {
	n := len(reqTimes)
	minHeap := &MinHeap{}

	heap.Init(minHeap)

	index := 0
	totalWaitTime := 0
	currTime := 0

	for index < n || minHeap.Len() > 0 {
		// 添加所有已请求的任务到堆中
		for index < n && reqTimes[index] <= currTime {
			heap.Push(minHeap, Task{
				reqTime: reqTimes[index],
				durTime: durTimes[index],
				index:   index,
			})
			index++
		}

		if minHeap.Len() == 0 {
			currTime = reqTimes[index]
			continue
		}

		task := heap.Pop(minHeap).(Task)
		waitTime := currTime - task.reqTime
		totalWaitTime += waitTime
		currTime += task.durTime
	}

	return float64(totalWaitTime) / float64(n)
}
