package main

import "fmt"

type NotificationService struct {
	notification Notification
}

type Notification interface {
	send(msg string) string
}

// Email ...
type EmailNotifier struct{}

func (e EmailNotifier) send(msg string) string {
	return fmt.Sprintf("Send <email> is message: %s", msg)
}

func (s NotificationService) SendNotification(msg string) string {
	return s.notification.send(msg)
}

func main() {
	notiSvc := NotificationService{
		notification: EmailNotifier{},
	}

	msg := notiSvc.SendNotification("Comfirm mail")
	fmt.Println(msg)
}
