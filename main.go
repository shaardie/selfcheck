package main

import (
	"fmt"
	"strings"

	"github.com/shaardie/selfcheck/util"

	"github.com/fatih/color"
)

const (
	INFO    = 1
	WARNING = 2
	ERROR   = 3
)

var (
	red    = color.New(color.FgRed)
	white  = color.New(color.FgWhite)
	yellow = color.New(color.FgYellow)
)

func log(level int, name string, result string, description string) {
	switch level {
	case INFO:
		white.Printf("INFO:")
	case WARNING:
		yellow.Printf("WARNING:")
	case ERROR:
		red.Printf("ERROR:")
	}
	fmt.Printf(" %v: %v\n", name, result)
	if description != "" {
		for _, line := range strings.Split(description, "\n") {
			fmt.Printf("  %v", line)
		}
	}
}

func main() {

	checks := []util.Check{
		util.SimpleChecker("CPU Load", "", cpuFunc),
		util.SimpleChecker("Uptime", "", uptimeFunc),
	}

	disk, err := diskChecks()
	if err != nil {
		fmt.Printf("Unable to do disk checks, %v", err)
	} else {
		checks = append(checks, disk...)
	}

	for _, check := range checks {
		name := check.Name()
		result, level, err := check.Run()
		if err != nil {
			fmt.Printf("Unable to run check %v, %v", name, err)
			continue
		}
		log(level, name, result, check.Description())
	}
}
