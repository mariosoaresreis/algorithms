package MergeIntervals

import (
	"container/heap"
	"sort"
)

type Meeting struct {
	Start int
	End   int
}

type MeetingHeap []Meeting

func (h MeetingHeap) Len() int           { return len(h) }
func (h MeetingHeap) Less(i, j int) bool { return h[i].End < h[j].End }
func (h MeetingHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MeetingHeap) Push(x interface{}) {
	*h = append(*h, x.(Meeting))
}

func (h *MeetingHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func FindMinimumMeetingRooms2(meetings []Meeting) int {
	stack := make([]Meeting, 0)
	minRooms := 0

	if len(meetings) == 0 {
		return 0
	}

	// sort the meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})

	for i, _ := range meetings {
		for len(stack) > 0 && meetings[i].Start >= stack[0].End {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, meetings[i])

		if len(stack) > minRooms {
			minRooms = len(stack)
		}
	}

	return minRooms
}

func FindMinimumMeetingRooms(meetings []Meeting) int {
	if len(meetings) == 0 {
		return 0
	}

	// sort the meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})

	minRooms := 0
	minHeap := &MeetingHeap{}
	heap.Init(minHeap)

	for _, meeting := range meetings {
		// remove all meetings that have ended
		for minHeap.Len() > 0 && meeting.Start >= (*minHeap)[0].End {
			heap.Pop(minHeap)
		}
		// add the current meeting into the minHeap
		heap.Push(minHeap, meeting)
		// all active meeting are in the minHeap, so we need rooms for all of them.
		if minHeap.Len() > minRooms {
			minRooms = minHeap.Len()
		}
	}

	return minRooms
}
