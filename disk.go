package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"syscall"

	"github.com/shaardie/selfcheck/util"
)

func prettifySpace(space float64) string {
	units := []string{"B", "KB", "MB", "GB"}
	var unit string
	for _, u := range units {
		unit = u
		if space < 1024 {
			break
		}
		space = space / 1024
	}
	return fmt.Sprintf("%v%v", math.Round(space), unit)
}

func run(path string) func() (string, int, error) {
	return func() (string, int, error) {
		fs := syscall.Statfs_t{}
		err := syscall.Statfs(path, &fs)
		if err != nil {
			return "", 0, err
		}
		all := float64(fs.Blocks) * float64(fs.Bsize)
		free := float64(fs.Bfree) * float64(fs.Bsize)

		level := ERROR
		if ratio := free / all; ratio > 0.1 {
			level = INFO
		} else if ratio > 0.05 {
			level = WARNING
		}

		return fmt.Sprintf("%v free", prettifySpace(free)), level, nil
	}
}

func diskChecks() ([]util.Check, error) {
	checks := make([]util.Check, 0)
	out, err := ioutil.ReadFile("/proc/mounts")
	if err != nil {
		return checks, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		if !(strings.HasPrefix(line, "/dev/") || strings.HasPrefix(line, "tmpfs")) {
			continue
		}
		mountArray := strings.Split(line, " ")
		if len(mountArray) < 2 {
			return checks, fmt.Errorf("'%v' not valid", line)
		}
		mountpoint := mountArray[1]
		check := util.SimpleChecker(fmt.Sprintf("Mountpoint %v", mountpoint), "", run(mountpoint))
		checks = append(checks, check)
	}
	return checks, nil
}
