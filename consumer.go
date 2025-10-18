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
		formattedMsg := fmt.Sprintf("To: %s\r\nSUbject: Test Email\r\n\r\n%s\r\n", recipient.Email, "Just testing email")
		message := []byte(formattedMsg)
		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "rishabhiitm@zohomail.in", []string{recipient.Email}, message)
		if err != nil {
			log.Fatal(err)
			return err

		}
		time.Sleep(time.Microsecond * 50)
		fmt.Printf("Worker %d: Sending email to: %s\n", id, recipient.Email)

	}

	return nil
}
