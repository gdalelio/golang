package lessons

import (
	"fmt"
	"runtime"
	"time"
)

func GoOS() {

	fmt.Print("Go is running on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//other OSes
		fmt.Printf("%s.", os)

	}
	fmt.Println(time.Now().Weekday())

}
