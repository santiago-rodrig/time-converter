package main

import (
	"flag"
	"fmt"
	"time"
)

var hourStart = flag.Int("hs", 0, "Starting hour")
var hoursOffset = flag.Int("ho", 0, "Offset hours")

func main() {
	flag.Parse()
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), *hourStart, now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	future := start.Add(time.Duration(*hoursOffset) * time.Hour)
	fmt.Println(future.Format(time.RFC1123Z))
}
