package main

import (
	"bytes"
	"html/template"
	"log"
)

func ExecuteTemplate(recp Recipient) (string, error) {
	temp, err := template.ParseFiles("email.tmpl")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var tpl bytes.Buffer
	err = temp.Execute(&tpl, recp)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tpl.String(), nil
}
