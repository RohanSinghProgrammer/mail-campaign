package main

import "sync"

func main() {
	ch := make(chan Receiver)

	go loadRecipient("mails.csv", ch)

	var wg sync.WaitGroup
	workers := 6

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go emailWorker(i, ch, &wg)
	}

	wg.Wait()
}
