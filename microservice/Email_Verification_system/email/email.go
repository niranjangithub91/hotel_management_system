package email

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

func Give_the_mail(t string) string {
	rand.Seed(time.Now().UnixNano())
	from := "xyz@email.com"
	password := "lsdbhjfdb"
	number := rand.Intn(9999)
	toList := []string{t}
	host := "smtp.gmail.com"
	port := "587"
	string_number := strconv.Itoa(number)
	fmt.Println(string_number)
	s := fmt.Sprintf("Your OTP  :- is %s \n", string_number)
	fmt.Println(s)
	body := []byte(string_number)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string_number)
	fmt.Println("Successfully sent mail to all user in toList")
	return string_number
}
