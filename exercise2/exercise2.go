package exercise2

import (
	"strconv"
	"strings"
)

func HowManyRounds(A string, B string) int {
	startTime := splitTime(A)
	endTime := splitTime(B)

	isOvernight := isOvernight(startTime, endTime)

	startTimeCleaned := cleanTimeOfStart(startTime)
	endTimeCleaned := cleanTimeOfFinish(endTime)

	if isOvernight && startTimeCleaned[0] == 0 {
		//For example:
		// first: startTime [23,56] and endTime [00,15]               - Overnight=True
		//     -> startTimeCleaned [00,00] and endTimeCleaned [00,15] - Overnight=False
		isOvernight = false
	}

	if !isOvernight && isStartTimeBiggerOrEqualThanEndTime(startTimeCleaned, endTimeCleaned) {
		//For example:
		// first: startTime [23,16] and endTime [23,17]
		//     -> startTimeCleaned [23,30] and endTimeCleaned [23,15]
		return 0
	}

	roundsCount := 0

	if isOvernight {
		roundsCount = roundsCount + roundsCounter(startTimeCleaned, []int{24, 00})

		startTimeCleaned = []int{0, 0}
	}

	roundsCount = roundsCount + roundsCounter(startTimeCleaned, endTimeCleaned)

	return roundsCount
}

func splitTime(time string) []int {
	split := strings.Split(time, ":")

	hour, _ := strconv.Atoi(split[0])
	minutes, _ := strconv.Atoi(split[1])

	return []int{hour, minutes}
}

func cleanTimeOfFinish(endTime []int) []int {
	minutes := endTime[1]

	if 0 < minutes && minutes < 15 {
		minutes = 0
	} else if 15 < minutes && minutes < 30 {
		minutes = 30
	} else if 45 < minutes && minutes < 60 {
		minutes = 45
	}

	return []int{endTime[0], minutes}
}

func cleanTimeOfStart(startTime []int) []int {
	hour := startTime[0]
	minutes := startTime[1]

	if 0 < minutes && minutes < 15 {
		minutes = 15
	} else if 15 < minutes && minutes < 30 {
		minutes = 30
	} else if 45 < minutes && minutes < 60 {
		minutes = 0
		hour = addOneHour(startTime[0])
	}

	return []int{hour, minutes}
}

func addOneHour(hour int) int {
	if hour == 23 {
		return 0
	} else {
		return hour + 1
	}
}

func isOvernight(startTime []int, endTime []int) bool {
	startTimeInMinutes := timeInMinutes(startTime)
	endTimeInMinutes := timeInMinutes(endTime)

	return endTimeInMinutes < startTimeInMinutes
}

// Is this function the same as isOvernight?
func isStartTimeBiggerOrEqualThanEndTime(startTime []int, endTime []int) bool {
	if startTime[0] >= endTime[0] { //startHour is bigger or equal than endHour
		return startTime[1] >= endTime[1] //startMinutes are bigger or equal than endMinutes
	}
	return false
}

func roundsCounter(startTime []int, endTime []int) int {
	return (timeInMinutes(endTime) - timeInMinutes(startTime)) / 15
}

func timeInMinutes(time []int) int {
	return time[0]*60 + time[1]
}
