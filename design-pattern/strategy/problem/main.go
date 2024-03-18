package strategy

import "fmt"

func main() {
	s := NotificationService{
		notifierType: "sms",
	}
	s.SendNotification("hello world")
}

type NotificationService struct {
	notifierType string
}

func (n NotificationService) SendNotification(message string) {
	if n.notifierType == "email" {
		fmt.Printf("Send message: %s (Sender: Email)\n", message)
	} else if n.notifierType == "sms" {
		fmt.Printf("Send message: %s (Sender: Sms)\n", message)

	}
}
