package AbstractFactory

import "fmt"

type SMSNotifier struct{}

func (s *SMSNotifier) SendNotification(recipient string, message string) {
	fmt.Println("SMS sender")
	fmt.Println("phone", recipient)
	fmt.Println("message", message)
}

type SMSNotifierFactory struct{}

func (s *SMSNotifierFactory) CreateNotification() Notification {
	return &SMSNotifier{}
}
