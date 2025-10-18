package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func EmailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) error {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		message, err := ExecuteTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker %d Error parsing template %s", id, recipient.Email)
			continue

		}
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "rishabhiitm@zohomail.in", []string{recipient.Email}, []byte(message))
		if err != nil {
			log.Fatal(err)
			return err

		}
		time.Sleep(time.Microsecond * 50)
		fmt.Printf("Worker %d: Sending email to: %s\n", id, recipient.Email)

	}

	return nil
}
