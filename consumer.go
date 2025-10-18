package main

import (
	"fmt"
	"sync"
)

func EmailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) error {
	defer wg.Done()
	for recipient := range ch {
		data := recipient
		fmt.Println(id, data)
	}

	return nil
}
