package cron

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description    string
		argsInput      []string
		expectedOutput *CronStruct
		expectedError  error
	}{
		{
			description:    "Too many arguments passed",
			argsInput:      []string{"*/15, 0, 1,15, *, 1-5, *, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid number of cron elements"),
		},
		{
			description:    "Too few arguments passed",
			argsInput:      []string{"*/15, 0, 1,15, *, 1-5"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid number of cron elements"),
		},
		{
			description: "Valid cron",
			argsInput:   []string{"*/15, 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "15", "30", "45"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    []string{"/usr/bin/find"},
			},
			expectedError: nil,
		},
		{
			description: "Valid cron 2",
			argsInput:   []string{"*/30, 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "30"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    []string{"/usr/bin/find"},
			},
			expectedError: nil,
		},
		{
			description: "Valid Day of week",
			argsInput:   []string{"*/30, 0, 1,15, *, MON-FRI, /usr/bin/find"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "30"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"0", "1", "2", "3", "4"},
				Command:    []string{"/usr/bin/find"},
			},
			expectedError: nil,
		},
		{
			description: "Valid Single Day of week",
			argsInput:   []string{"*/30, 0, 1,15, *, WED, /usr/bin/find"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "30"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"2"},
				Command:    []string{"/usr/bin/find"},
			},
			expectedError: nil,
		},
		{
			description: "Valid gap range Day of week",
			argsInput:   []string{"*/30, 0, 1,15, *, 4-0, /usr/bin/find"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "30"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"4", "5", "6", "0"},
				Command:    []string{"/usr/bin/find"},
			},
			expectedError: nil,
		},
		{
			description: "multi part command",
			argsInput:   []string{"*/30, 0, 1,15, *, 4-0, /usr/bin/find test file.txt"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "30"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"4", "5", "6", "0"},
				Command:    []string{"/usr/bin/find test file.txt"},
			},
			expectedError: nil,
		},
		{
			description: "Valid cron all star",
			argsInput:   []string{"*, *, *, *, *, /usr/bin/find"},
			expectedOutput: &CronStruct{
				Minute:     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"},
				Hour:       []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"},
				DayOfMonth: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
				Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				DayOfWeek:  []string{"0", "1", "2", "3", "4", "5", "6"},
				Command:    []string{"/usr/bin/find"},
			},
			expectedError: nil,
		},
		{
			description:    "Invalid range",
			argsInput:      []string{"10-1, 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid minute definition"),
		},
		{
			description:    "Invalid range 2",
			argsInput:      []string{"1-61, 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid minute definition"),
		},
		{
			description:    "Invalid range 3",
			argsInput:      []string{"0-30, 0, 0-15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid day of month definition"),
		},
		{
			description:    "Invalid step",
			argsInput:      []string{"*/30/2, 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid minute definition"),
		},
		{
			description:    "Invalid Minute",
			argsInput:      []string{"., 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid minute definition"),
		},
		{
			description:    "Invalid Minute Missing",
			argsInput:      []string{", 0, 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid minute definition"),
		},
		{
			description:    "Invalid Hour",
			argsInput:      []string{"*/15, ., 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid hour definition"),
		},
		{
			description:    "Invalid Hour Missing",
			argsInput:      []string{"*/15, , 1,15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid hour definition"),
		},
		{
			description:    "Invalid Day of Month",
			argsInput:      []string{"*/15, 0, 1.15, *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid day of month definition"),
		},
		{
			description:    "Invalid Day of Month Missing",
			argsInput:      []string{"*/15, 0, , *, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid day of month definition"),
		},
		{
			description:    "Invalid Month",
			argsInput:      []string{"*/15, 0, 1,15, 1.2, 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid month definition"),
		},
		{
			description:    "Invalid Month Missing",
			argsInput:      []string{"*/15, 0, 1,15, , 1-5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid month definition"),
		},
		{
			description:    "Invalid Day of Week",
			argsInput:      []string{"*/15, 0, 1,15, *, 1.5, /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid day of week definition"),
		},
		{
			description:    "Invalid Day of Week Missing",
			argsInput:      []string{"*/15, 0, 1,15, *, , /usr/bin/find"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid day of week definition"),
		},
		{
			description:    "Invalid Command",
			argsInput:      []string{"*/15, 0, 1,15, *, 1-5, *"},
			expectedOutput: nil,
			expectedError:  errors.New("Invalid command definition"),
		},
	}

	for _, tc := range tests {
		cronResponse, err := ParseString(tc.argsInput)

		assert.Equal(tc.expectedOutput, cronResponse, tc.description)
		assert.Equal(tc.expectedError, err, tc.description)
	}
}
