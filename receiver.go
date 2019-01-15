package notifier

type Receiver struct {
	// To whom notification will be sent
	To []string

	// How this notification will be sent
	Method string
}

// NewReceiver creates new receiver
func NewReceiver(to []string, method string) *Receiver {
	return &Receiver{to, method}
}
