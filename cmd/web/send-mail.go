package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/fxfrancky/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	// A function that will execute in the background and listen to the MailChannel
	go func() {
		for {
			msg := <-app.MainChan
			sendMsg(msg)
		}

	}()
}

func sendMsg(m models.MailData) {

	// serverHost := flag.String("serverhost", "", "Server host")
	// // serverPort := flag.String("serverport", "", "Server Port")
	// serverUsername := flag.String("serverusername", "", "Server Username")
	// serverPassword := flag.String("serverpassword", "", "Server Password")
	// flag.Parse()

	server := mail.NewSMTPClient()
	// server.Host = "localhost"
	server.Host = app.ServerHost
	// server.Host = *serverHost
	// server.Port = 1025
	server.Port = 587
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	// For Production
	server.Username = app.ServerUsername
	server.Password = app.ServerPassword
	server.Encryption = mail.EncryptionSTARTTLS
	// server.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// server.Port

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal("Error During the connection", err.Error())
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailTemplate := string(data)
		// Replace [%body%] int the template by m.Content
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal("Error Trying to Send the message", err.Error())
	} else {
		log.Println("Email Sent!")
	}
}
