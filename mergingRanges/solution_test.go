package mergingRanges

import (
  "reflect"
  "sort"
  "testing"
)

func TestMergingRangesFuncTouching(tt *testing.T) {
  input := []meetingTime{{2, 4}, {1, 2}, {0, 1}}
  output := []meetingTime{{0, 4}}
  result, err := mergingRanges(input)
  if err != nil {
    tt.Error(err)
  }
  if !reflect.DeepEqual(result, output) {
    tt.Errorf("Expected: %v; Actual: %v.", output, result)
  }
}

func TestMergingRangesFuncTotalInclude(tt *testing.T) {
  input := []meetingTime{{1, 5}, {2, 3}}
  output := []meetingTime{{1, 5}}
  result, err := mergingRanges(input)
  if err != nil {
    tt.Error(err)
  }
  if !reflect.DeepEqual(result, output) {
    tt.Errorf("Expected: %v; Actual: %v.", output, result)
  }
}

func TestMergingRangesFuncTotalIncludeLargerSet(tt *testing.T) {
  input := []meetingTime{{1, 10}, {2, 6}, {3, 5}, {7, 9}}
  output := []meetingTime{{1, 10}}
  result, err := mergingRanges(input)
  if err != nil {
    tt.Error(err)
  }
  if !reflect.DeepEqual(result, output) {
    tt.Errorf("Expected: %v; Actual: %v.", output, result)
  }
}

func TestMergingRangesFuncMultiResult(tt *testing.T) {
  input := []meetingTime{{2, 4}, {2, 3}, {0, 1}}
  output := []meetingTime{{0, 1}, {2, 4}}
  result, err := mergingRanges(input)
  if err != nil {
    tt.Error(err)
  }
  // Sort the result to compare.
  sort.Sort(ByStartTime(result))
  if !reflect.DeepEqual(result, output) {
    tt.Errorf("Expected: %v; Actual: %v.", output, result)
  }
}

func TestMergingRangesFuncZeroTimeMeetings(tt *testing.T) {
  input := []meetingTime{{1, 1}, {0, 0}, {9, 9}}
  result, err := mergingRanges(input)
  if err != nil {
    tt.Error(err)
  }
  if len(result) != 0 {
    tt.Errorf("Expected to have zero result but got %v", result)
  }
}

func TestMergingRangesEmptyInput(tt *testing.T) {
  input := []meetingTime{}
  result, err := mergingRanges(input)
  if err != nil {
    tt.Error(err)
  }
  if len(result) != 0 {
    tt.Errorf("Expected to have zero result but got %v", result)
  }
}

func TestMergingRangesInvalidMeetingTime(tt *testing.T) {
  input := []meetingTime{{1, -1}}
  _, err := mergingRanges(input)
  if err == nil {
    tt.Errorf("Expecting error to occur since the input %v is not valid.",
      input)
  }
}
