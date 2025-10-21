package main

import (
	"bytes"
	"html/template"
	"sync"
)

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

func execTemplate(r Receiver) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, r)
	if err != nil{
		return "", err
	}

	return tpl.String(), nil
}
