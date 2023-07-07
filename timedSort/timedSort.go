package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	rand.Seed(time.Now().Unix())

	arr := rand.Perm(10000)
	fmt.Println("Before:", arr)

	sort.Ints(arr)
	fmt.Println("After:", arr)

	if sort.IntsAreSorted(arr) {
		fmt.Println("Successfully sorted!")
	} else {
		fmt.Println("Uh oh... Something went wrong!")
	}

	fmt.Printf("Program took %v\n", time.Since(start))
}