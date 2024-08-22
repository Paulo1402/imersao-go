package main

import "time"


func worker(workerId int, data chan int) {
	println("Worker", workerId, "received", <-data)


	for x := range data {
		println("Worker", workerId, "received", x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	qtdWorkers := 3

	for i := range qtdWorkers {
		go worker(i, ch)
	}

	for i := range 15 {
		ch <- i
	}
}