package sms

import (
	"context"

	"github.com/EsmaeilMazahery/wild/enums"
)

// IProvider is an interface to sms provider
type IProvider interface {
	//Send save Service data to the store
	Send(ctx context.Context, message string, receptor ...string) error
	//Verify finds a member by ID
	Verify(ctx context.Context, receptor string, token string, smsTemplate enums.SmsTemplate) error
}
