package main

import (
	"fmt"
	"time"
)

func main ()  {

	fmt.Print("Go runs on ")
	i:= 0
	switch i {

	case 0 : fmt.Print("windows")
	case 1 : fmt.Print("linux")
	}

	today:= time.Now().Weekday()
	fmt.Print("\n",time.Saturday)

	switch time.Saturday {

	case today + 0:
		fmt.Print("\n today")
	
	}

}
