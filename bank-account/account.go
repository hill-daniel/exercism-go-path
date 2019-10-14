package account

import (
	"sync"
)

// Account simulates a bank account supporting opening/closing, withdrawals, and deposits of money.
type Account struct {
	sync.Mutex
	balance int64
	open    bool
}

// Open creates a new account with an initial balance of initialDeposit.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		balance: initialDeposit,
		open:    true}
}

// Close closes the account and returns the current amount of money stored in the account as payout.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}

	payout = a.balance
	a.open = false
	return payout, true
}

// Balance returns the current amount of money stored in the account.
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}

	return a.balance, true
}

// Deposit stores the given amount of money into the account and returns the updated balance.
// Negative deposits are considered withdrawals and will succeed if the current balance
// is greater than or equal to the amount to be withdrawn.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}

	if amount < 0 && a.balance < -amount {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}
