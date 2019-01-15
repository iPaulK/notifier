package notifier

import (
	"fmt"
	"strings"

	"github.com/iPaulK/notifier/providers"
	"github.com/iPaulK/notifier/providers/smtp"
	"github.com/iPaulK/notifier/providers/viber"
)

type Notifier struct{}

// NewNotifier creates new notifier
func NewNotifier() *Notifier {
	return &Notifier{}
}

// SendToReceiver send notice to receiver
func (n *Notifier) SendToReceiver(notice Notice, receiver *Receiver) error {
	via, err := loadNotifier(receiver.Method)
	if err != nil {
		return err
	}

	switch v := via.(type) {
	case providers.ByEmail:
		layout := NewLayout().SetBody(notice.GetBody())
		body, err := layout.Render()
		if err != nil {
			return fmt.Errorf("failed to render mail. Reason: %s", err)
		}
		return v.To(receiver.To[0], receiver.To[1:]...).
			SetSubject(notice.GetSubject()).
			SetBody(body).
			Send()
	case providers.ByChat:
		return v.To(receiver.To[0], receiver.To[1:]...).
			SetBody(notice.GetBody()).
			Send()
	default:
		return fmt.Errorf(`invalid notifier "%s"`, receiver.Method)
	}
}

// loadNotifier load notifier by code
func loadNotifier(via string) (interface{}, error) {
	switch strings.ToLower(via) {
	case smtp.CODE:
		return smtp.Load()
	case viber.CODE:
		return viber.Load()
	}
	return nil, fmt.Errorf("unknown notifier %s", via)
}
