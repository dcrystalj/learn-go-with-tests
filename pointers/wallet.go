package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("some error")

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(d Bitcoin) {
	w.balance += d
}

func (w *Wallet) Withdraw(d Bitcoin) error {
	if w.balance-d < 0 {
		return ErrInsufficientFunds
	}
	w.balance -= d
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
