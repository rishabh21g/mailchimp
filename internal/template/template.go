package template

import (
	"bytes"
	"html/template"
	"log"

	"github.com/rishabh21g/mailchimp/internal/types"
)

func ExecuteTemplate(recp types.Recipient) (string, error) {

	temp, err := template.ParseFiles("../../internal/template/email.tmpl")
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
