package errors

import "fmt"

type BeginTransactionErr struct {
	err error
}

func NewBeginTransactionErr(err error) BeginTransactionErr {
	return BeginTransactionErr{
		err: err,
	}
}

func (err BeginTransactionErr) Error() string {
	return fmt.Sprintf("can't begin transaction: %w", err)
}

func (err BeginTransactionErr) Unwrap() error {
	return err
}
