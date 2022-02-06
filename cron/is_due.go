package cron

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type ExpressionKey struct {
	value        string
	integer      bool
	integerValue int
	valid        bool
}

var maximumKeyValues = map[int]int{
	0: 59,
	1: 59,
	2: 23,
	3: 31,
	4: 6,
}

var keys = map[int]string{
	0: "minute",
	1: "hour",
	2: "day",
	3: "month",
	4: "weekday",
}

// IsDue is to check the cron expression date
func IsDue(expression string, date time.Time) bool {

	cron := make(map[string]ExpressionKey)

	for index, value := range strings.Split(expression, " ") {

		cronKey := ExpressionKey{}

		cronKey.value = value

		integerValue, err := strconv.Atoi(value)
		if err == nil {
			cronKey.integer = true
			cronKey.integerValue = integerValue
		}

		if cronKey.integer {
			if integerValue >= 0 && integerValue <= maximumKeyValues[index] {
				cronKey.valid = true
			}
		} else {
			if strings.Contains(value, "*") || strings.Contains(value, "-") || strings.Contains(value, ",") {
				cronKey.valid = true
			}
		}

		cron[keys[index]] = cronKey

	}

	for index, value := range cron {

		if !value.valid {
			return false
		}

		var timeAsInteger int

		switch index {
		case "minute":
			timeAsInteger = date.Minute()
		case "hour":
			timeAsInteger = date.Hour()
		case "day":
			timeAsInteger = date.Day()
		case "month":
			timeAsInteger = int(date.Month())
		case "weekday":
			timeAsInteger = int(date.Weekday())
		}

		if cron[index].integer {

			if cron[index].integerValue != timeAsInteger {
				return false
			}

		} else {

			if strings.Contains(cron[index].value, "*/") {
				part := strings.ReplaceAll(cron[index].value, "*/", "")
				integerValue, err := strconv.ParseFloat(part, 64)
				if err == nil {
					var calc float64
					calc = float64(timeAsInteger) / integerValue
					if calc != math.Trunc(calc) {
						return false
					}
				}
			}

			if strings.Contains(cron[index].value, "-") {
				split := strings.Split(cron[index].value, "-")
				var integerValues []int
				for _, value := range split {
					integerValue, err := strconv.Atoi(value)
					if err != nil {
						return false
					}
					integerValues = append(integerValues, integerValue)
				}
				if timeAsInteger <= integerValues[0] || timeAsInteger >= integerValues[1] {
					return false
				}
			}

			if strings.Contains(cron[index].value, ",") {
				var execute bool
				for _, value := range strings.Split(cron[index].value, ",") {
					integerValue, err := strconv.Atoi(value)
					if err != nil {
						return false
					}
					if integerValue == timeAsInteger {
						execute = true
					}
				}
				if !execute {
					return false
				}
			}

		}

	}

	return true

}
