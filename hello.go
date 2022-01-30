package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func main() {
	fmt.Println("\nExercise One:")
	fmt.Println("Hello Go!")
	second()
	third()
	fourth()
	fifth()
	sixth()
	seventh()
	eighth()
	nine()
	ten()
	eleven()
}

func second() {
	fmt.Println("\nExercise Two:")
	fmt.Println("My favorite number is", rand.Intn(10))
}

func third() {
	fmt.Println("\nExercise Three:")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func fourth() {
	fmt.Println("\nExercise Four:")
	fmt.Println(math.Pi)
}

func add(x, y int) int {
	return x + y
}

func fifth() {
	fmt.Println("\nExercise Five:")
	fmt.Println("42 + 13 =", add(42, 13))
}

func swap(x, y string) (string, string) {
	return y, x
}

func sixth() {
	a, b := swap("hello", "world")
	fmt.Println("\nExercise Six:")
	fmt.Println(a, b)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func seventh() {
	fmt.Println("\nExercise Seven:")
	fmt.Println(split(17))
}

func eighth() {
	var c, python, java bool
	var i int
	fmt.Println("\nExercise Eight:")
	fmt.Println(i, c, python, java)
}

func nine() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println("\nExercise Nine:")
	fmt.Println(i, j, c, python, java)
}

func ten() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println("\nExercise Ten:")
	fmt.Println(i, j, k, c, python, java)
}

func eleven() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Println("\nExercise Eleven:")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
