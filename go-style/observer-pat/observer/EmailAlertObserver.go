package observer

import (
	"fmt"
)

// EmailAlertObserverImpl is a concrete observer that sends an email alert
type EmailAlertObserverImpl struct {
	email string
}

// NewEmailAlertObserverImpl creates a new instance of EmailAlertObserverImpl
func NewEmailAlertObserverImpl(emailId string) *EmailAlertObserverImpl {
	return &EmailAlertObserverImpl{
		email: emailId,
	}
}

// Update is called when the observable changes
func (e *EmailAlertObserverImpl) Update() {
	sendMail(e.email, "Product is in stock, hurry up!")
}

// sendMail is a helper function to simulate sending an email
func sendMail(emailId string, message string) {
	fmt.Printf("Email sent to: %s\nMessage: %s\n", emailId, message)
}
