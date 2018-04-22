package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func cpuFunc() (string, int, error) {
	out, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return "", 0, err
	}

	loads := strings.Split(string(out), " ")
	if len(loads) < 2 {
		return "", 0, fmt.Errorf("Unable to find 5 minute load average")
	}

	loadString := loads[1]
	load, err := strconv.ParseFloat(loadString, 64)
	if err != nil {
		return "", 0, err
	}

	cpus := 0.0
	out, err = ioutil.ReadFile("/proc/cpuinfo")
	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "processor") {
			cpus++
		}
	}

	level := 0
	if ratio := load / cpus; ratio < 1 {
		level = INFO
	} else if ratio < 2 {
		level = WARNING
	} else {
		level = ERROR
	}

	return loadString, level, nil
}
