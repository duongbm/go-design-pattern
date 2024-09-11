package AbstractFactory

import "fmt"

type EmailNotifier struct{}

func (e *EmailNotifier) SendNotification(recipient string, message string) {
	// logic send with email
	fmt.Println("Send email: ")
	fmt.Println("Recipient: ", recipient)
	fmt.Println("Message: ", message)
}

type EmailNotifierFactory struct{}

func (e *EmailNotifierFactory) CreateNotification() Notification {
	return &EmailNotifier{}
}
