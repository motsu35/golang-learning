package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	defer fmt.Println("world!")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sqrt(float64(sum)), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("in two days.")
	default:
		fmt.Println("too far away :(")

	}

	fmt.Println("Hello")
}
