package smtp

// this should be replaced with hermes later

import (
	"github.com/neel4os/warg/internal/common/config"
	"gopkg.in/gomail.v2"
)

type SendMail struct {
	dialer *gomail.Dialer
}

func NewSendMail() *SendMail {
	cfg := config.GetConfig()
	smtpcfg := cfg.SmtpConfig
	d := gomail.NewDialer(smtpcfg.Host, smtpcfg.Port, "", "")
	d.TLSConfig = nil
	d.SSL = false
	return &SendMail{
		dialer: d,
	}
}

func (s *SendMail) SendMail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "warg@admin.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
