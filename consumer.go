package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Receiver, wg *sync.WaitGroup) {
	defer wg.Done()

	smtpHost := "localhost"
	smtpPort := "1025"
	for receiver := range ch {
		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Its for %s\r\n\r\n%s", receiver.Email, receiver.Name, "Just an campaign email")
		formattedMsg, err := execTemplate(receiver)
		if err != nil {
			// TODO: Add proper error handling  logging
			log.Printf("Worker %d, Facing issue with email %s", id, receiver.Email)
			continue
		}
		msg := []byte(formattedMsg)
		fmt.Printf("Worker %d, Sending email to %s", id, receiver.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "rohan@singh.com", []string{receiver.Email}, msg)
		if err != nil {
			log.Fatal("Error sending email!", err.Error())
		}
		// Delay to prevent from SMTP Server rate limiting
		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d, Send email to %s", id, receiver.Email)
	}
}
