package AbstractFactory

type Notification interface {
	SendNotification(recipe string, message string)
}

type NotificationFactory interface {
	CreateNotification() Notification
}
