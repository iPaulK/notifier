package providers

type ByEmail interface {
	Code() string
	From(from string) ByEmail
	SetSubject(subject string) ByEmail
	SetBody(body string) ByEmail
	To(to string, cc ...string) ByEmail
	Send() error
}

type ByChat interface {
	Code() string
	SetBody(body string) ByChat
	To(to string, cc ...string) ByChat
	Send() error
}
