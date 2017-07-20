package main

import (
	"github.com/supme/directEmail"
	"time"
	"fmt"
)

func main() {
	email := directEmail.New()

	// if use socks5 proxy
	email.Ip = "socks://123.124.125.126:8080"

	// if use specified interface
	email.Ip = "192.168.0.10"

	// if use NAT
	email.MapIp = map[string]string {
		"192.168.0.10": "31.33.34.35",
	}

	// if left blank, then auto resolved (for socks use IP for connecting server)
	email.Host = "resolv.name.myhost.com"

	email.FromEmail = "sender@example.com"
	email.FromName = "Sender name"

	email.ToEmail = "user@example.com"
	email.ToName = "Reciver name"

	// add extra headers if need
	email.Header(fmt.Sprintf("Message-ID: <test_message_%d>", time.Now().Unix()))
	email.Header("Content-Language: ru")

	email.Subject = "Тест отправки email"

	email.Part(directEmail.TypeTextHTML, `
<h2>My email</h2>
<p>Текст моего сообщения</p>
	`)

	email.Part(directEmail.TypeTextPlain, `
My email
Текст моего сообщения
`)

	// attach file if need
	email.Attachment("/path/to/attach/file.jpg")


	email.Render()
	print("\n", string(email.GetRawMessage()), "\n\n\n")

	err := email.Send()
	if err != nil {
		print("Send email with error:", err.Error())
	}

}