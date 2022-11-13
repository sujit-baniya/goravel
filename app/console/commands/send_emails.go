package commands

import (
	"fmt"
	"github.com/sujit-baniya/framework/contracts/console"
	"github.com/sujit-baniya/framework/contracts/console/command"
)

type SendEmails struct {
}

// Signature The name and signature of the console command.
func (receiver *SendEmails) Signature() string {
	return "send:emails"
}

// Description The console command description.
func (receiver *SendEmails) Description() string {
	return "Send Emails"
}

// Extend The console command extend.
func (receiver *SendEmails) Extend() command.Extend {
	return command.Extend{
		Flags: []command.Flag{
			{
				Name:    "lang",
				Value:   "default",
				Aliases: []string{"l"},
				Usage:   "language for the greeting",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *SendEmails) Handle(ctx console.Context) error {
	fmt.Println(ctx.Option("lang"))
	return nil
}
