package sms

import (
	"context"
	"fmt"

	"github.com/EsmaeilMazahery/wild/enums"
	"github.com/EsmaeilMazahery/wild/infrastructure/constant"
	"github.com/kavenegar/kavenegar-go"
)

// Smskavenegar stores users in memory
type Smskavenegar struct {
	sender string
	api    *kavenegar.Kavenegar
}

// NewSmskavenegar returns a new in-memory user store
func NewSmskavenegar() *Smskavenegar {
	return &Smskavenegar{
		api:    kavenegar.New(constant.KavenegarAPIKey),
		sender: "",
	}
}

//Send send sms to kavenegar
func (p *Smskavenegar) Send(ctx context.Context, message string, receptor ...string) error {

	if res, err := p.api.Message.Send(p.sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			return err
		case *kavenegar.HTTPError:
			return err
		default:
			return err
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}

	return nil
}

//Verify ...
func (p *Smskavenegar) Verify(ctx context.Context, receptor string, token string, smsTemplate enums.SmsTemplate) error {
	params := &kavenegar.VerifyLookupParam{}
	if res, err := p.api.Verify.Lookup(receptor, smsTemplate.String(), token, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			return err
		case *kavenegar.HTTPError:
			return err
		default:
			return err
		}
	} else {
		fmt.Println("MessageID 	= ", res.MessageID)
		fmt.Println("Status    	= ", res.Status)
	}
	return nil
}
