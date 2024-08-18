package MergeIntervals

import (
	"sort"
)

type Interval struct {
	Start, End int
}

func Merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	mergedIntervals := []Interval{}
	start := intervals[0].Start
	end := intervals[0].End

	for _, interval := range intervals {
		if end >= interval.Start {
			end = max(end, interval.End)
		} else {
			mergedIntervals = append(mergedIntervals, Interval{start, end})
			start = interval.Start
			end = interval.End
		}
	}

	mergedIntervals = append(mergedIntervals, Interval{start, end})

	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
