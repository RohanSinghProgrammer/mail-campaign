package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Receiver, wg *sync.WaitGroup){
	defer wg.Done()
	for receiver := range ch {
		fmt.Println(id, receiver)
	}
}