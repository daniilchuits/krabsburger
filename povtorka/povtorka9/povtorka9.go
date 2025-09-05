package main

import "fmt"

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() float64
}

type SavingAccount struct {
	balance float64
}

type CreditAccount struct {
	balance float64
	limit   float64
}

func (c *CreditAccount) Deposit(amount float64) {
	c.balance += amount
	fmt.Printf("Пополнено на %.2f\n", amount)
}

func (s *SavingAccount) Deposit(amount float64) {
	s.balance += amount
	fmt.Printf("Пополнено на %.2f\n", amount)
}

func (c *CreditAccount) Withdraw(amount float64) {
	if amount > c.balance+c.limit {
		fmt.Printf("Недостаточно средств: баланс %.2f + лимит %.2f\n", c.balance, c.limit)
		return
	}
	c.balance -= amount
	fmt.Printf("Снято %.2f\n", amount)
}

func (s *SavingAccount) Withdraw(amount float64) {
	if amount > s.balance {
		fmt.Printf("Недостаточно средств, баланс %.2f\n", s.balance)
		return
	}
	s.balance -= amount
	fmt.Printf("Снято %.2f\n", amount)
}

func (c *CreditAccount) Balance() float64 {
	return c.balance
}

func (s *SavingAccount) Balance() float64 {
	return s.balance
}

func PrintBalance(a Account) {
	fmt.Printf("Текущий баланс: %.2f\n", a.Balance())
}

func main() {
	sav := &SavingAccount{}
	cred := &CreditAccount{limit: 200}

	accounts := []Account{sav, cred}

	// операции
	sav.Deposit(100)
	sav.Withdraw(50)
	PrintBalance(sav)

	cred.Deposit(50)
	cred.Withdraw(200)
	PrintBalance(cred)

	for _, acc := range accounts {
		fmt.Println("Финальный баланс:", acc.Balance())
	}
}
