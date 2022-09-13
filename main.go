package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Date(2021, 01, 19, 07, 35, 0, 0, time.Local)

	now := time.Now()

	date := "30/01/2019 15:27:00 -03"
	parsed, err := time.Parse("02/01/2006 15:04:05 -07", date)
	if err != nil {
		panic(err)
	}

	next := parsed.Add(15 * 24 * time.Hour)
	//before := parsed.Add(-15 * 24 * time.Hour)

	fmt.Println(now.Format("02/01/2006 15:04:05 -07"))
	fmt.Println(next.Format("02/01/2006 15:04:05 -07"))

	fmt.Println(next.Equal(now))
	fmt.Println(next.Before(now))
	fmt.Println(next.After(now))

}
