package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	recipientChannel := make(chan Recipient) // unbuffered channel
	go LoadRecipients("./email.csv", recipientChannel)

	workers := 5
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go EmailWorker(i, recipientChannel, &wg)

	}
	wg.Wait()
	// time.Sleep(time.Second)       // temporary solution for completing the task of goroutines

}
