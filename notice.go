package notifier

type Notice struct {
	Body    string
	Subject string
}

func (n *Notice) SetBody(body string) *Notice {
	n.Body = body
	return n
}

func (n *Notice) GetBody() string {
	return n.Body
}

func (n *Notice) SetSubject(subject string) *Notice {
	n.Subject = subject
	return n
}

func (n *Notice) GetSubject() string {
	return n.Subject
}
