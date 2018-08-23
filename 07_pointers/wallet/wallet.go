package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds error when funds are insufficient
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin defines a bitcoin
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct defines a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit will add amt to balance
func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

// Withdraw will substract amt to balance
func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}

// Balance returns balnce value
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
