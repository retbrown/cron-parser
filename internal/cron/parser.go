package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errMinute     = fmt.Errorf("Invalid minute definition")
	errHour       = fmt.Errorf("Invalid hour definition")
	errDayOfMonth = fmt.Errorf("Invalid day of month definition")
	errMonth      = fmt.Errorf("Invalid month definition")
	errDayOfWeek  = fmt.Errorf("Invalid day of week definition")
	errCommand    = fmt.Errorf("Invalid command definition")
)

// CronStruct contains expanded times for a cron definition
type CronStruct struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    []string
}

// ParseString takes the provided cron string, parses and returns the expanded times the string will run. Or error if invalid
func ParseString(args []string) (*CronStruct, error) {

	if len(args) != 1 {
		return nil, fmt.Errorf("Invalid number of cron elements")
	}

	params := strings.Split(args[0], " ")

	if len(params) != 6 {
		return nil, fmt.Errorf("Invalid number of cron elements")
	}

	parsedMinute, err := parseTime(params[0], errMinute, 0, 60)
	if err != nil {
		return nil, err
	}

	parsedHour, err := parseTime(params[1], errHour, 0, 24)
	if err != nil {
		return nil, err
	}

	parsedDayOfMonth, err := parseTime(params[2], errDayOfMonth, 1, 32)
	if err != nil {
		return nil, err
	}

	parsedMonth, err := parseTime(params[3], errMonth, 1, 13)
	if err != nil {
		return nil, err
	}

	parsedDayOfWeek, err := parseTime(params[4], errDayOfWeek, 0, 7)
	if err != nil {
		return nil, err
	}

	parsedCommand, err := parseCommand(params[5])
	if err != nil {
		return nil, err
	}

	return &CronStruct{
		Minute:     parsedMinute,
		Hour:       parsedHour,
		DayOfMonth: parsedDayOfMonth,
		Month:      parsedMonth,
		DayOfWeek:  parsedDayOfWeek,
		Command:    parsedCommand,
	}, nil
}

func parseCommand(value string) ([]string, error) {
	found, _ := regexp.MatchString(`^[a-zA-Z 0-9\,\-\/\.]*$`, value)

	if !found {
		return nil, errCommand
	}

	return []string{value}, nil
}

func parseTime(value string, errValue error, min int64, max int64) ([]string, error) {
	value = strings.TrimRight(value, ",")

	found, _ := regexp.MatchString(`^[0-9\,\*\-\/]*$`, value)
	if !found {
		return nil, errValue
	}

	if len(value) == 0 {
		return nil, errValue
	}

	if value == "*" {
		return getValues(min, max), nil
	}

	if strings.Contains(value, "/") {
		allValues := getValues(min, max)

		values := strings.Split(value, "/")

		if len(values) > 2 {
			return nil, errValue
		}

		every, err := strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			return nil, errValue
		}

		returnValues := make([]string, 0)

		for i := int64(0); i < int64(len(allValues)); i = i + every {
			returnValues = append(returnValues, allValues[i])
		}

		return returnValues, nil
	}

	if strings.Contains(value, "-") {
		values := strings.Split(value, "-")

		minValue, err := strconv.ParseInt(values[0], 10, 64)
		maxValue, err := strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			return nil, errValue
		}

		if minValue > maxValue {
			return nil, errValue
		}

		if minValue < min {
			return nil, errValue
		}

		if maxValue+1 > max {
			return nil, errValue
		}

		return getValues(minValue, maxValue+1), nil
	}

	if strings.Contains(value, ",") {
		return strings.Split(value, ","), nil
	}

	return []string{value}, nil
}

func getValues(min int64, max int64) []string {
	values := make([]string, 0)
	for i := min; i < max; i++ {
		values = append(values, strconv.FormatInt(int64(i), 10))
	}

	return values
}
