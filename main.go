package main

import (
	"log"
	"flag"
	"time"
	"github.com/lukechampine/hey"
)

func main()  {
	num := flag.Int("d", 5, "timebox duration")
	title := flag.String("t", "Default task", "timebox Title")
    flag.Parse()
	
	hey.Push(hey.Notification{
		Title:    *title,
		Body:     "Timebox started",
		AppName:  "Timebox",
		Duration: 0,
	})

	log.Println("timebox:", *num)
    time.Sleep(time.Duration(*num) * time.Second)

	hey.Push(hey.Notification{
		Title:    *title,
		Body:     "Timebox ended",
		AppName:  "Timebox",
		Duration: 0,
	})
}

