package MergeIntervals

func InsertNewInterval(intervals []Interval, newInterval Interval) []Interval {

	if len(intervals) < 2 {
		return intervals
	}

	count := 0
	mergedIntervals := make([]Interval, 0)

	for count < len(intervals) && intervals[count].End < newInterval.Start {
		mergedIntervals = append(mergedIntervals, intervals[count])
		count++
	}

	for count < len(intervals) && intervals[count].Start <= newInterval.End {
		newInterval.Start = min(newInterval.Start, intervals[count].Start)
		newInterval.End = max(newInterval.End, intervals[count].End)
		count++
	}

	mergedIntervals = append(mergedIntervals, newInterval)

	for count < len(intervals) {
		mergedIntervals = append(mergedIntervals, intervals[count])
		count++
	}

	return mergedIntervals
}
