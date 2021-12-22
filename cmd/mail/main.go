package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gomail.v2"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	from         = MustEnv("FROM")
	smtp_host    = MustEnv("SMTP_HOST")
	smtp_port, _ = strconv.Atoi(MustEnv("SMTP_PORT"))
	smtp_user    = MustEnv("SMTP_USER")
	smtp_pass    = MustEnv("SMTP_PASS")
)

func main() {

	logFile, err := os.OpenFile("/home/pi/sunsetty/log.mail", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	stdout := log.New(mw, "", log.Ltime)

	if len(os.Args) != 2 {
		stdout.Println("usage: mail <full_path_image>")
		stdout.Fatal("Missing image file!")
	}

	picPath := os.Args[1]
	stdout.Printf("sending pic %s", picPath)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", from)
	m.SetHeader("Subject", "Buon pomeriggio!")
	m.SetBody("text/html", "Ciao Tom !<br?>Buon pomeriggio. Questa è la foto di oggi ")
	m.Attach(picPath)

	d := gomail.NewDialer(smtp_host, smtp_port, smtp_user, smtp_pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		stdout.Fatal(err)
	}

	stdout.Println("Done!")
}

func MustEnv(key string) string {
	s := os.Getenv(key)
	if s == "" {
		panic(fmt.Sprintf("Required env %s is missing", key))
	}

	return s
}
