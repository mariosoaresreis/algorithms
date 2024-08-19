package MergeIntervals

import "sort"

func Insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	mergedIntervals := make([]Interval, 0)
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	start := intervals[0].Start
	end := intervals[0].End

	for _, interval := range intervals {
		if end >= interval.Start {
			end = max(interval.End, end)
		} else {
			mergedIntervals = append(mergedIntervals, Interval{Start: start, End: end})
			start = interval.Start
			end = interval.End
		}
	}

	mergedIntervals = append(mergedIntervals, Interval{Start: start, End: end})

	return mergedIntervals
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
