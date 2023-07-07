package main

import (
    "fmt"
    "time"
)

func f(id int) {
    for i := 0; i < 3; i++ {
        fmt.Println(id, ":", i)
    }
}

func status(msg string) {
    fmt.Println("Status check:", msg)
}

func main() {
	f(1)
    for i := 2; i < 10; i++ {
        go f(i)
    }

    go status("I am still running")
	
    time.Sleep(3 * time.Second)
    fmt.Println("done")
}
