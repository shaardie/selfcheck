package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func uptimeFunc() (string, int, error) {
	out, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return "", 0, err
	}

	uptimeList := strings.Split(string(out), " ")

	uptimeString := uptimeList[0]
	up, err := strconv.ParseFloat(uptimeString, 64)
	if err != nil {
		return "", 0, err
	}

	level := ERROR
	if up < 30*60*60 {
		level = INFO
	} else if up < 90*60*60 {
		level = WARNING
	}

	return fmt.Sprintf("%vs", uptimeString), level, nil

}
