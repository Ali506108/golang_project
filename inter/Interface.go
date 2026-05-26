package inter

import (
	"context"
	"errors"
	"fmt"
)

var (
	errUnknowSenderType = errors.New("notifier : unknow sender type")
	errEmptyAddress     = errors.New("notifier : address cannot be empty")
)

type Caller interface {
	Call(number int) error
}

type Sender interface {
	Send(ctx context.Context, msg string) error
}

type IPhone struct {
	Caller
	Sender
}

type Email struct {
	Address string
}

func (e *Email) Send(ctx context.Context, msg string) error {
	if e.Address == "" {
		return errEmptyAddress
	}
	fmt.Printf("[Email] Message '%s' sent to %s\n", msg, e.Address)
	return nil
}

type Sms struct {
	ProvideNumber string
}

func (s *Sms) Send(ctx context.Context, msg string) error {
	if s.ProvideNumber == "" {
		return errEmptyAddress
	}
	fmt.Printf("[SMS] Message : '%s' sent to %s \n", msg, s.ProvideNumber)
	return nil
}

func Notify(ctx context.Context, s Sender, msg string) error {
	if s == nil {
		return errUnknowSenderType
	}

	if err := s.Send(ctx, msg); err != nil {
		return fmt.Errorf("Failed to send notification : %s ", err)
	}

	if err := s.Send(ctx, msg); err != nil {
		return fmt.Errorf("failed to send notification : %w ", err)

	}
	return nil
}

func AutoLog(s Sender) {
	switch v := s.(type) {
	case *Email:
		fmt.Printf("[Audit] type: Email , data is : %s\n", v.Address)
	case *Sms:
		fmt.Printf("[Audit] type : Sms phone number is : %s\n", v.ProvideNumber)
	default:
		fmt.Printf("[Audit] unknow data types")
	}
}
