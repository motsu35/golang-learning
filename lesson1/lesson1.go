package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

var c, python, java = true, false, "NO!"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("my favorite number is", rand.Intn(10))
	fmt.Println("i also like pi, but not the kind you eat!.", math.Pi)
	fmt.Println("if you like those check out this cool number", add(5, 10))
	a, b := swap("me", "swap")
	fmt.Println(a, b)
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	fmt.Println("================================")
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	fmt.Println("================================")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
