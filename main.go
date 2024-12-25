package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

	"gopkg.in/gomail.v2"
)

func sendMailSimple(subject string, body string, to []string) {

	auth := smtp.PlainAuth(
		"",
		"nagabhushanbhandary524@gmail.com",
		"wztzhuutrjyatcjz",
		"smtp.gmail.com",
	)
	messagee := "Subject:" + subject + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"nagabhushanbhandary524@gmail.com",
		to,
		[]byte(messagee),
	)
	if err != nil {
		log.Fatal(err)
	}
}
func sendMailSimpleHTML(subject string, templatePath string, to []string) {
	//get html
	var body bytes.Buffer
	t, _ := template.ParseFiles(templatePath)

	t.Execute(&body, struct{ Name string }{Name: "bhushan"})
	auth := smtp.PlainAuth(
		"",

		"nagabhushanbhandary524@gmail.com",
		"wztzhuutrjyatcjz",
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	messagee := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"nagabhushanbhandary524@gmail.com",
		to,
		[]byte(messagee),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func sendGomail(templatePath string) {
	var body bytes.Buffer

	// Parse the template file
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Execute the template
	err = t.Execute(&body, struct{ Name string }{Name: "bhushan"})
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", "nagabhushanbhandary524@gmail.com")
	m.SetHeader("To", "nagabhushanbhandary524@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./paymentpage.png")

	// Configure the SMTP dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, "nagabhushanbhandary524@gmail.com", "wztzhuutrjyatcjz")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

func main() {
	// sendMailSimple("New sub",
	// 	"hii @ bhushan dont loose this  one time password  i.e 7899",
	// 	[]string{"nagabhushanbhandary524@gmail.com"},
	// )
	sendMailSimpleHTML("New sub",
		"./text.html",
		[]string{"nagabhushanbhandary524@gmail.com"},
	)
	sendGomail("./text.html")
}
