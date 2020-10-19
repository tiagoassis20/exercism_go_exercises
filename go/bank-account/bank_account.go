package account

import "sync"

type Account struct {
	balance int64
	input   chan operation
	isOpen  bool
	closing sync.Mutex
}

type operation struct {
	ret   chan int64
	value int64
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	a := &Account{
		balance: initialDeposit,
		input:   make(chan operation, 1),
		isOpen:  true,
	}
	go func(account *Account) {
		for val := range account.input {
			if val.value < 0 && val.value+account.balance < 0 {
				val.ret <- -1
			} else {
				account.balance += val.value
				val.ret <- account.balance
			}
		}
	}(a)
	return a
}

func (a *Account) Close() (payout int64, ok bool) {
	a.closing.Lock()
	defer a.closing.Unlock()
	if a.isOpen {
		a.isOpen = false
		close(a.input)
		return a.balance, true
	}
	return 0, a.isOpen
}

func (a *Account) Balance() (balance int64, ok bool) {
	if a.isOpen {
		return a.balance, true
	}
	return 0, false
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.isOpen {
		ch := make(chan int64)
		a.input <- operation{ret: ch, value: amount}
		newValue := <-ch
		if newValue >= 0 {
			return newValue, true
		}
	}
	return 0, false
}
