package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to go time")
	fmt.Println("****************************************************************************")
	presentTime := time.Now()
	fmt.Println(presentTime) // It gives the present time with a whole lot of other stuff
	fmt.Println("****************************************************************************")
	fmt.Println("After formatting")
	fmt.Println(presentTime.Format("02-01-2006")) // in the Format func you can pass a lot of parameters eg: 01-02-2006 gives the date in MM - DD - YYYY , the date 01-02-2006 is required to be passed as an argument for the formatting of date
	fmt.Println("****************************************************************************")
	fmt.Println("After day formatting")
	fmt.Println(presentTime.Format("02-01-2006 Monday")) // Monday needs to be passed as a day for go to print the day 
	fmt.Println("****************************************************************************")
	fmt.Println("After day and time formatting")
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) // the following time is needed to be passed to get the current time

	//If u want to crete a date there are also a whole lot of parameters year, time.montnh, rest of them are int day, hour, min, sec, nano sec, timezone
	fmt.Println("****************************************************************************")
	fmt.Println("Here is a created date")
	createdDate := time.Date(2024,time.November,21,21,0,0,0,time.UTC)
	fmt.Println(createdDate)

}
