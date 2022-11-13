package jobs

import "fmt"

type SendEmails struct {
}

// Signature The name and signature of the job.
func (receiver *SendEmails) Signature() string {
	return "send_emails"
}

// Handle Execute the job.
func (receiver *SendEmails) Handle(args ...any) error {
	fmt.Println("I'm sending email")
	return nil
}
