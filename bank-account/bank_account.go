/*
Creates a bank account type which supports concurrency
*/

package account

import (
	"sync"
)

// Account type which holds the amount
type Account struct {
	sync.Mutex
	amt  int64
	open bool
}

// Open a new bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit >= 0 {
		a := Account{amt: initialDeposit, open: true}
		return &a
	}
	return nil
}

// Close a bank account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.open {
		payout = a.amt
		a.amt = 0
		a.open = false
		return payout, true
	}
	return 0, false
}

// Balance returns the balance of a bank account as long as it's open
func (a *Account) Balance() (balance int64, ok bool) {
	if a.open {
		return a.amt, true
	}
	return a.amt, false
}

//Deposit deposits the amount into the account as long as the balance doesn't go below 0
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.amt+amount >= 0 && a.open {
		a.amt += amount
		return a.amt, true
	}
	return a.amt, false
}
