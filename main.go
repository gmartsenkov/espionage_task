package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/bradhe/stopwatch"
)

func main() {
	start := stopwatch.Start()
	arr := []int{1, 2, 4, 5, 6, 7, 8, 9}
	for {
		shuffledArr := shuffle(arr[0:])
		a := arrToInt(shuffledArr[0:3])
		b := 30 + shuffledArr[3]
		result := arrToInt(shuffledArr[4:])
		fmt.Printf("%v * %v == %v \n", a, b, result)
		if a*b == result {
			watch := stopwatch.Stop(start)
			fmt.Printf("Milliseconds elapsed: %v\n", watch.Milliseconds())
			os.Exit(0)
		}
	}
}

func arrToInt(arr []int) int {
	result := 0
	base := float64(10)
	for i := 0; i < len(arr); i++ {
		result = result + int(float64(arr[i])*(math.Pow(base, float64(len(arr)-i-1))))
	}
	return int(result)
}

func shuffle(slice []int) []int {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
