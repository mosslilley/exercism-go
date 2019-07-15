/*

 */

package account

import (
	"sync"
)

// Account does stuff
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

// Close closes a bank account
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

// Balance returns the balance of a bank account
func (a *Account) Balance() (balance int64, ok bool) {
	if a.open {
		return a.amt, true
	}
	return a.amt, false
}

//Deposit deposits the amount into the account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.amt+amount >= 0 && a.open {
		a.amt += amount
		return a.amt, true
	}
	return a.amt, false
}
