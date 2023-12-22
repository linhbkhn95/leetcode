package simplebank

type Bank struct {
	balances map[int]int64
}

func Constructor(balance []int64) *Bank {
	balances := make(map[int]int64, len(balance))
	for i, value := range balance {
		balances[i+1] = value
	}
	return &Bank{
		balances: balances,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	balance, ok := this.balances[account1]
	if !ok || balance < money {
		return false
	}
	this.balances[account1] -= money
	this.balances[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if _, ok := this.balances[account]; !ok {
		return false
	}
	this.balances[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	balance, ok := this.balances[account]
	if !ok || balance < money {
		return false
	}
	this.balances[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
