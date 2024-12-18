package main

import (
	"21/challenges/day1"
	"21/challenges/day10"
	"21/challenges/day11"
	"21/challenges/day2"
	"21/challenges/day3"
	"21/challenges/day4"
	"21/challenges/day5"
	"21/challenges/day6"
	"21/challenges/day7"
	"21/challenges/day8"
	"21/challenges/day9"
	"flag"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetReportCaller(true)
	log.SetReportTimestamp(false)
	log.SetFormatter(log.TextFormatter)

	var mode string
	verbosePtr := flag.Bool("v", false, "Show debug logs")
	flag.Parse()
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	}

	if len(flag.Args()) < 1 {
		log.Fatalf("Need to specify challenge date, got args %v", flag.Args())
	}
	chal, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		log.Fatal("Challenge date must be an int", "invalid challenge", chal, "error", err)
	}

	if len(flag.Args()) != 2 {
		mode = "test"
	} else {
		mode = "final"
	}
	startTime := time.Now()
	switch chal {
	case 1:
		day1.Sol(mode)
	case 2:
		day2.Sol(mode)
	case 3:
		day3.Sol(mode)
	case 4:
		day4.Sol(mode)
	case 5:
		day5.Sol(mode)
	case 6:
		day6.Sol(mode)
	case 7:
		day7.Sol(mode)
	case 8:
		day8.Sol(mode)
	case 9:
		day9.Sol(mode)
	case 10:
		day10.Sol(mode)
	case 11:
		day11.Sol(mode)
	}
	executionTime := time.Since(startTime)
	defer log.Infof("Finished executing in %s", executionTime)
}
