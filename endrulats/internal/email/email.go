package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func EmailVerify(email, SiteToken string) error {

	type list struct {
		Address string
	}

	l := list{Address: email}

	fmt.Println(`Please go to "http://localhost:5002/loginemail/` + l.Address + `/` + SiteToken + `" to login.`)

	d := gomail.NewDialer("smtp.mail.yahoo.com", 587, "dacdac123321@yahoo.com", "Qwaszx1@")
	s, err := d.Dial()
	if err != nil {
		fmt.Println(err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "---")
	m.SetAddressHeader("To", l.Address, l.Address)
	m.SetHeader("Subject", "login verification")
	m.SetBody("text/html", fmt.Sprintf(`
	Hi `+l.Address+`

	Please go to "http://localhost:5002/loginemail/`+l.Address+`/`+SiteToken+`" to login.
	`))

	if err := gomail.Send(s, m); err != nil {

		fmt.Printf("Could not send email to %q: %v", l.Address, err)
		return err
	}
	m.Reset()

	return nil
}
