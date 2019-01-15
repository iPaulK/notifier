package smtp

import (
	"errors"

	"github.com/iPaulK/notifier/providers"
	gomail "gopkg.in/gomail.v2"
)

const CODE = "smtp"

type Options struct {
	status   string
	username string
	password string
	host     string
	port     int
	protocol string
	from     string
	to       []string
}

type client struct {
	options *Options
	subject string
	body    string
}

var _ providers.ByEmail = &client{}

// New creates new client
func New(options *Options) (*client, error) {
	return &client{options: options}, nil
}

// Load init client
func Load() (*client, error) {
	options := &Options{}
	return New(options)
}

func (c client) Code() string {
	return CODE
}

func (c client) From(from string) providers.ByEmail {
	c.options.from = from
	return &c
}

func (c client) To(to string, cc ...string) providers.ByEmail {
	c.options.to = append([]string{to}, cc...)
	return &c
}

func (c client) SetSubject(subject string) providers.ByEmail {
	c.subject = subject
	return &c
}

func (c client) SetBody(body string) providers.ByEmail {
	c.body = body
	return &c
}

// Send sends email
func (c *client) Send() error {
	if len(c.options.to) == 0 {
		return errors.New("Missing to")
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", c.options.from)
	mail.SetHeader("To", c.options.to...)
	mail.SetHeader("Subject", c.subject)
	mail.SetBody("text/html", c.body)

	var d *gomail.Dialer
	if c.options.status != "disabled" && c.options.username != "" && c.options.password != "" {
		d = gomail.NewPlainDialer(c.options.host, c.options.port, c.options.username, c.options.password)
	} else {
		d = &gomail.Dialer{Host: c.options.host, Port: c.options.port}
	}
	return d.DialAndSend(mail)
}
