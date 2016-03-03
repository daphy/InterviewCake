package mergingRanges

import (
  "fmt"
  "sort"
)

type meetingTime struct {
  startTime int64
  endTime   int64
}

// Define sort interface.
type ByStartTime []meetingTime

func (times ByStartTime) Len() int {
  return len(times)
}

func (times ByStartTime) Less(i, j int) bool {
  return times[i].startTime < times[j].startTime
}

func (times ByStartTime) Swap(i, j int) {
  times[i], times[j] = times[j], times[i]
}

func mergingRanges(inputTimes []meetingTime) (
  mergedTimes []meetingTime, err error) {
  sort.Sort(ByStartTime(inputTimes))

  curStartTime := int64(-1)
  curEndTime := int64(-1)
  for _, oneMeetingTime := range inputTimes {
    if oneMeetingTime.endTime < oneMeetingTime.startTime {
      err = fmt.Errorf("Invalid meeting time %v; end time can not be before "+
        "start time.", oneMeetingTime)
      return
    }
    // If this time starts *after* the current condensed meeting time, this
    // record will start a new condensed time info. In this case, save the
    // current info, and restart a new one.
    if oneMeetingTime.startTime > curEndTime {
      if curEndTime > curStartTime {
        mergedTimes = append(mergedTimes,
          meetingTime{curStartTime, curEndTime})
      }
      // If curEndTime <= curStart time, this condensed meeting time isn't
      // meaningful; we will not record it.
      curEndTime = oneMeetingTime.endTime
      curStartTime = oneMeetingTime.startTime
      continue
    }
    // If this meeting times starts within the current condensed time, we will
    // need to see if we should update the current end time to include this
    // time.
    if oneMeetingTime.startTime <= curEndTime {
      if oneMeetingTime.endTime > curEndTime {
        curEndTime = oneMeetingTime.endTime
      }
      continue
    }
    panic("This line should never be reached.")
  }

  // We need to see if there's anything to add to the mergedTimes.
  if curEndTime > curStartTime {
    mergedTimes = append(mergedTimes, meetingTime{curStartTime, curEndTime})
  }

  return
}
