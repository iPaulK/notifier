package viber

import "github.com/iPaulK/notifier/providers"

const CODE = "viber"

type Options struct {
	// TODO: add viber options
}

type client struct {
	options *Options
	body    string
}

var _ providers.ByChat = &client{}

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

func (c client) SetBody(body string) providers.ByChat {
	c.body = body
	return &c
}

func (c client) To(to string, cc ...string) providers.ByChat {
	// TODO: prepare to
	return &c
}

func (c *client) Send() error {
	// TODO: send to viber
	return nil
}
