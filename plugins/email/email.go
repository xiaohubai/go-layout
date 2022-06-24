package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/xiaohubai/go-layout/configs/global"
)

func Send(t, subject string, body string) (err error) {
	auth := smtp.PlainAuth("", global.Cfg.Email.From, global.Cfg.Email.Secret, global.Cfg.Email.Host)
	e := &email.Email{
		To:      strings.Split(t, ","),
		From:    global.Cfg.Email.From,
		Subject: subject,
		HTML:    []byte(body),
	}
	address := fmt.Sprintf("%s:%d", global.Cfg.Email.Host, global.Cfg.Email.Port)
	if global.Cfg.Email.IsSSL {
		err = e.SendWithTLS(address, auth, &tls.Config{ServerName: global.Cfg.Email.Host})
	} else {
		err = e.Send(address, auth)
	}
	return
}
