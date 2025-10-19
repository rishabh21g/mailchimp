package main

import (
	"sync"

	"github.com/rishabh21g/mailchimp/internal/services"
	"github.com/rishabh21g/mailchimp/internal/types"
)

func main() {
	var wg sync.WaitGroup
	recipientChannel := make(chan types.Recipient) // unbuffered channel
	go services.LoadRecipients("../../internal/data/email.csv", recipientChannel)

	workers := 5
	for i := 0; i <= workers; i++ {
		wg.Add(1)
		go services.EmailWorker(i, recipientChannel, &wg)

	}
	wg.Wait()
	// time.Sleep(time.Second)       // temporary solution for completing the task of goroutines

}
