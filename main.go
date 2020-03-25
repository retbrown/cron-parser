package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/retbrown/cron-parser/internal/cron"
)

func main() {
	cronStruct, err := cron.ParseString(os.Args[1:])
	if err != nil {
		fmt.Println("Cron error: " + err.Error())
		os.Exit(1)
	}

	fmt.Printf("%-14s%s\n", "minute", strings.Join(cronStruct.Minute, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(cronStruct.Hour, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(cronStruct.DayOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(cronStruct.Month, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(cronStruct.DayOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command", strings.Join(cronStruct.Command, " "))
}
