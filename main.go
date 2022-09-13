package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Date(2021, 01, 19, 07, 35, 0, 0, time.Local)

	//now := time.Now()
	fmt.Println(now.Format("02/01/2006 15:04:05 -07"))
}
